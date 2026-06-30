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
	client.EndpointMap = map[string]*string{
		"us-west-1":             dara.String("cbn.aliyuncs.com"),
		"us-east-1":             dara.String("cbn.aliyuncs.com"),
		"na-south-1":            dara.String("cbn.aliyuncs.com"),
		"me-central-1":          dara.String("cbn.aliyuncs.com"),
		"eu-west-1":             dara.String("cbn.aliyuncs.com"),
		"eu-central-1":          dara.String("cbn.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("cbn.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("cbn.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("cbn.aliyuncs.com"),
		"cn-shenzhen":           dara.String("cbn.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("cbn.aliyuncs.com"),
		"cn-shanghai":           dara.String("cbn.aliyuncs.com"),
		"cn-qingdao":            dara.String("cbn.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("cbn.aliyuncs.com"),
		"cn-huhehaote":          dara.String("cbn.aliyuncs.com"),
		"cn-hongkong":           dara.String("cbn.aliyuncs.com"),
		"cn-heyuan":             dara.String("cbn.aliyuncs.com"),
		"cn-hangzhou":           dara.String("cbn.aliyuncs.com"),
		"cn-guangzhou":          dara.String("cbn.aliyuncs.com"),
		"cn-chengdu":            dara.String("cbn.aliyuncs.com"),
		"cn-beijing-finance-1":  dara.String("cbn.aliyuncs.com"),
		"cn-beijing":            dara.String("cbn.aliyuncs.com"),
		"ap-southeast-6":        dara.String("cbn.aliyuncs.com"),
		"ap-southeast-5":        dara.String("cbn.aliyuncs.com"),
		"ap-southeast-3":        dara.String("cbn.aliyuncs.com"),
		"ap-southeast-2":        dara.String("cbn.aliyuncs.com"),
		"ap-southeast-1":        dara.String("cbn.aliyuncs.com"),
		"ap-south-1":            dara.String("cbn.aliyuncs.com"),
		"ap-northeast-1":        dara.String("cbn.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("cbn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Enables a flow log. After the flow log is enabled, the system collects traffic information about a specified resource.
//
// Description:
//
// - A flow log is enabled by default after creation. If the flow log was stopped, call this operation to re-enable it.
//
// - `ActiveFlowLog` is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//   - If a flow log is in the **Modifying*	- state, the flow log is being enabled. In this case, you can query the flow log but cannot perform other operations.
//
//   - If a flow log is in the **Active*	- state, the flow log is enabled.
//
// @param request - ActiveFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveFlowLogResponse
func (client *Client) ActiveFlowLogWithOptions(request *ActiveFlowLogRequest, runtime *dara.RuntimeOptions) (_result *ActiveFlowLogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a flow log. After the flow log is enabled, the system collects traffic information about a specified resource.
//
// Description:
//
// - A flow log is enabled by default after creation. If the flow log was stopped, call this operation to re-enable it.
//
// - `ActiveFlowLog` is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//   - If a flow log is in the **Modifying*	- state, the flow log is being enabled. In this case, you can query the flow log but cannot perform other operations.
//
//   - If a flow log is in the **Active*	- state, the flow log is enabled.
//
// @param request - ActiveFlowLogRequest
//
// @return ActiveFlowLogResponse
func (client *Client) ActiveFlowLog(request *ActiveFlowLogRequest) (_result *ActiveFlowLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActiveFlowLogResponse{}
	_body, _err := client.ActiveFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - If a traffic classification rule is in the **Creating*	- state, the traffic classification rule is being created. In this case, you can query the traffic classification rule but cannot perform other operations.
//
// - If a traffic classification rule is in the **Active*	- state, the traffic classification rule is added to the traffic marking policy.
//
// @param request - AddTrafficMatchRuleToTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTrafficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) AddTrafficMatchRuleToTrafficMarkingPolicyWithOptions(request *AddTrafficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddTrafficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - If a traffic classification rule is in the **Creating*	- state, the traffic classification rule is being created. In this case, you can query the traffic classification rule but cannot perform other operations.
//
// - If a traffic classification rule is in the **Active*	- state, the traffic classification rule is added to the traffic marking policy.
//
// @param request - AddTrafficMatchRuleToTrafficMarkingPolicyRequest
//
// @return AddTrafficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) AddTrafficMatchRuleToTrafficMarkingPolicy(request *AddTrafficMatchRuleToTrafficMarkingPolicyRequest) (_result *AddTrafficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTrafficMatchRuleToTrafficMarkingPolicyResponse{}
	_body, _err := client.AddTrafficMatchRuleToTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AddTraficMatchRuleToTrafficMarkingPolicyWithOptions(request *AddTraficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *AddTraficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AddTraficMatchRuleToTrafficMarkingPolicyResponse
// Deprecated
func (client *Client) AddTraficMatchRuleToTrafficMarkingPolicy(request *AddTraficMatchRuleToTrafficMarkingPolicyRequest) (_result *AddTraficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTraficMatchRuleToTrafficMarkingPolicyResponse{}
	_body, _err := client.AddTraficMatchRuleToTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a bandwidth plan to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can attach multiple bandwidth plans to a CEN instance. However, you cannot attach more than one bandwidth plan for the same connected areas.
//
// For example, if a bandwidth plan for connections within the Chinese mainland is attached to a CEN instance, you cannot attach another bandwidth plan for the same connected areas. However, you can attach a bandwidth plan for connections between the Chinese mainland and North America.
//
// @param request - AssociateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateCenBandwidthPackageResponse
func (client *Client) AssociateCenBandwidthPackageWithOptions(request *AssociateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *AssociateCenBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a bandwidth plan to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// You can attach multiple bandwidth plans to a CEN instance. However, you cannot attach more than one bandwidth plan for the same connected areas.
//
// For example, if a bandwidth plan for connections within the Chinese mainland is attached to a CEN instance, you cannot attach another bandwidth plan for the same connected areas. However, you can attach a bandwidth plan for connections between the Chinese mainland and North America.
//
// @param request - AssociateCenBandwidthPackageRequest
//
// @return AssociateCenBandwidthPackageResponse
func (client *Client) AssociateCenBandwidthPackage(request *AssociateCenBandwidthPackageRequest) (_result *AssociateCenBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateCenBandwidthPackageResponse{}
	_body, _err := client.AssociateCenBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can create a forwarding association.
//
// Description:
//
// After you create a network instance connection, you must associate it with the route table of an Enterprise Edition transit router. After the association is created, the Enterprise Edition transit router forwards traffic from the network instance based on the routes in the associated route table. Before you call this operation, take note of the following:
//
// - Only route tables of Enterprise Edition transit routers support associations. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - A network instance connection can be associated with only one route table of an Enterprise Edition transit router.
//
// - The **AssociateTransitRouterAttachmentWithRouteTable*	- operation is asynchronous. After you send a request, the system returns a **Request ID**, but the association is not immediately created. The system creates the association in the background. You can call the **ListTransitRouterRouteTableAssociations*	- operation to query the status of the association.
//
//   - If the association is in the **Associating*	- state, it is being created. In this state, you can only query the association and cannot perform other operations.
//
//   - If the association is in the **Active*	- state, it is successfully created.
//
// @param request - AssociateTransitRouterAttachmentWithRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateTransitRouterAttachmentWithRouteTableResponse
func (client *Client) AssociateTransitRouterAttachmentWithRouteTableWithOptions(request *AssociateTransitRouterAttachmentWithRouteTableRequest, runtime *dara.RuntimeOptions) (_result *AssociateTransitRouterAttachmentWithRouteTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can create a forwarding association.
//
// Description:
//
// After you create a network instance connection, you must associate it with the route table of an Enterprise Edition transit router. After the association is created, the Enterprise Edition transit router forwards traffic from the network instance based on the routes in the associated route table. Before you call this operation, take note of the following:
//
// - Only route tables of Enterprise Edition transit routers support associations. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - A network instance connection can be associated with only one route table of an Enterprise Edition transit router.
//
// - The **AssociateTransitRouterAttachmentWithRouteTable*	- operation is asynchronous. After you send a request, the system returns a **Request ID**, but the association is not immediately created. The system creates the association in the background. You can call the **ListTransitRouterRouteTableAssociations*	- operation to query the status of the association.
//
//   - If the association is in the **Associating*	- state, it is being created. In this state, you can only query the association and cannot perform other operations.
//
//   - If the association is in the **Active*	- state, it is successfully created.
//
// @param request - AssociateTransitRouterAttachmentWithRouteTableRequest
//
// @return AssociateTransitRouterAttachmentWithRouteTableResponse
func (client *Client) AssociateTransitRouterAttachmentWithRouteTable(request *AssociateTransitRouterAttachmentWithRouteTableRequest) (_result *AssociateTransitRouterAttachmentWithRouteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateTransitRouterAttachmentWithRouteTableResponse{}
	_body, _err := client.AssociateTransitRouterAttachmentWithRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AssociateTransitRouterMulticastDomainWithOptions(request *AssociateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *AssociateTransitRouterMulticastDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AssociateTransitRouterMulticastDomainResponse
func (client *Client) AssociateTransitRouterMulticastDomain(request *AssociateTransitRouterMulticastDomainRequest) (_result *AssociateTransitRouterMulticastDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AssociateTransitRouterMulticastDomainResponse{}
	_body, _err := client.AssociateTransitRouterMulticastDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - For more information about how to grant CEN permissions on virtual private clouds (VPCs) that belong to another Alibaba Cloud account, see [GrantInstanceToCen](https://help.aliyun.com/document_detail/126224.html).
//
// - For more information about how to grant CEN permissions on Cloud Connect Network (CCN) instances that belong to another Alibaba Cloud account, see [GrantInstanceToCbn](https://help.aliyun.com/document_detail/126141.html).
//
// - By default, you cannot grant permissions on virtual border routers (VBRs) that belong to another Alibaba Cloud account to a CEN instance. If you need to use this feature, contact your account manager.
//
// @param request - AttachCenChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachCenChildInstanceResponse
func (client *Client) AttachCenChildInstanceWithOptions(request *AttachCenChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachCenChildInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - For more information about how to grant CEN permissions on virtual private clouds (VPCs) that belong to another Alibaba Cloud account, see [GrantInstanceToCen](https://help.aliyun.com/document_detail/126224.html).
//
// - For more information about how to grant CEN permissions on Cloud Connect Network (CCN) instances that belong to another Alibaba Cloud account, see [GrantInstanceToCbn](https://help.aliyun.com/document_detail/126141.html).
//
// - By default, you cannot grant permissions on virtual border routers (VBRs) that belong to another Alibaba Cloud account to a CEN instance. If you need to use this feature, contact your account manager.
//
// @param request - AttachCenChildInstanceRequest
//
// @return AttachCenChildInstanceResponse
func (client *Client) AttachCenChildInstance(request *AttachCenChildInstanceRequest) (_result *AttachCenChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachCenChildInstanceResponse{}
	_body, _err := client.AttachCenChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckTransitRouterServiceWithOptions(request *CheckTransitRouterServiceRequest, runtime *dara.RuntimeOptions) (_result *CheckTransitRouterServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckTransitRouterServiceResponse
func (client *Client) CheckTransitRouterService(request *CheckTransitRouterServiceRequest) (_result *CheckTransitRouterServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckTransitRouterServiceResponse{}
	_body, _err := client.CheckTransitRouterServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A Cloud Enterprise Network (CEN) instance is the fundamental resource for managing an integrated network. A CEN instance manages a network that can span one or more regions. Before you enable communication between network instances, you must call the CreateCen operation to create a CEN instance.
//
// Description:
//
// *CreateCen*	- is an asynchronous operation. After a request is sent, the system returns a CEN instance ID, but the CEN instance is not created immediately. The creation task runs in the background. You can call the **DescribeCens*	- operation to query the status of the CEN instance.
//
// - If a CEN instance is in the **Creating*	- status, it is being created. In this status, you can only query the instance. You cannot perform other operations.
//
// - If a CEN instance is in the **Active*	- status, the instance is created.
//
// @param request - CreateCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenResponse
func (client *Client) CreateCenWithOptions(request *CreateCenRequest, runtime *dara.RuntimeOptions) (_result *CreateCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A Cloud Enterprise Network (CEN) instance is the fundamental resource for managing an integrated network. A CEN instance manages a network that can span one or more regions. Before you enable communication between network instances, you must call the CreateCen operation to create a CEN instance.
//
// Description:
//
// *CreateCen*	- is an asynchronous operation. After a request is sent, the system returns a CEN instance ID, but the CEN instance is not created immediately. The creation task runs in the background. You can call the **DescribeCens*	- operation to query the status of the CEN instance.
//
// - If a CEN instance is in the **Creating*	- status, it is being created. In this status, you can only query the instance. You cannot perform other operations.
//
// - If a CEN instance is in the **Active*	- status, the instance is created.
//
// @param request - CreateCenRequest
//
// @return CreateCenResponse
func (client *Client) CreateCen(request *CreateCenRequest) (_result *CreateCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCenResponse{}
	_body, _err := client.CreateCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// To connect network instances in different regions, you must purchase a bandwidth plan. You can call the CreateCenBandwidthPackage operation to create a bandwidth plan.
//
// Description:
//
// - When you create a bandwidth plan, you must specify the connected areas. A connected area is a collection of one or more Alibaba Cloud regions. You must select the connected areas based on the regions that you want to connect. For more information about the relationship between areas and regions, see [Purchase a bandwidth plan](https://help.aliyun.com/document_detail/181560.html).
//
// - For more information about billing, see [Billing](https://help.aliyun.com/document_detail/189836.html).
//
// - **CreateCenBandwidthPackage*	- is an asynchronous operation. After you send a request, the system returns a bandwidth plan ID. The bandwidth plan is created in the background. You can call the **DescribeCenBandwidthPackages*	- operation to query the status of the bandwidth plan. The bandwidth plan is successfully created when its status changes to **Idle*	- or **InUse**.
//
// @param request - CreateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenBandwidthPackageResponse
func (client *Client) CreateCenBandwidthPackageWithOptions(request *CreateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *CreateCenBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// To connect network instances in different regions, you must purchase a bandwidth plan. You can call the CreateCenBandwidthPackage operation to create a bandwidth plan.
//
// Description:
//
// - When you create a bandwidth plan, you must specify the connected areas. A connected area is a collection of one or more Alibaba Cloud regions. You must select the connected areas based on the regions that you want to connect. For more information about the relationship between areas and regions, see [Purchase a bandwidth plan](https://help.aliyun.com/document_detail/181560.html).
//
// - For more information about billing, see [Billing](https://help.aliyun.com/document_detail/189836.html).
//
// - **CreateCenBandwidthPackage*	- is an asynchronous operation. After you send a request, the system returns a bandwidth plan ID. The bandwidth plan is created in the background. You can call the **DescribeCenBandwidthPackages*	- operation to query the status of the bandwidth plan. The bandwidth plan is successfully created when its status changes to **Idle*	- or **InUse**.
//
// @param request - CreateCenBandwidthPackageRequest
//
// @return CreateCenBandwidthPackageResponse
func (client *Client) CreateCenBandwidthPackage(request *CreateCenBandwidthPackageRequest) (_result *CreateCenBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCenBandwidthPackageResponse{}
	_body, _err := client.CreateCenBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the CreateCenChildInstanceRouteEntryToAttachment operation to create a route for a network instance connected to an Enterprise Edition transit router. The next hop of the route points to the transit router in the current region.
//
// Description:
//
// - You can create routes only for Virtual Private Cloud (VPC) and virtual border router (VBR) instances that are connected to an Enterprise Edition transit router.
//
// - The next hop of the route is the **transit router connection*	- (the network instance connection) by default and cannot be modified.
//
// - **CreateCenChildInstanceRouteEntryToAttachment*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- and creates the route in the background. The route is not created immediately. You can call the **DescribeRouteEntryList*	- operation for the VPC to query the status of the route.
//
//   - If the route is in the **Pending*	- state, it is being created. During this time, you can only query the route and cannot perform other operations.
//
//   - If the route is in the **Available*	- state, the route is created.
//
// @param request - CreateCenChildInstanceRouteEntryToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenChildInstanceRouteEntryToAttachmentResponse
func (client *Client) CreateCenChildInstanceRouteEntryToAttachmentWithOptions(request *CreateCenChildInstanceRouteEntryToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CreateCenChildInstanceRouteEntryToAttachment operation to create a route for a network instance connected to an Enterprise Edition transit router. The next hop of the route points to the transit router in the current region.
//
// Description:
//
// - You can create routes only for Virtual Private Cloud (VPC) and virtual border router (VBR) instances that are connected to an Enterprise Edition transit router.
//
// - The next hop of the route is the **transit router connection*	- (the network instance connection) by default and cannot be modified.
//
// - **CreateCenChildInstanceRouteEntryToAttachment*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- and creates the route in the background. The route is not created immediately. You can call the **DescribeRouteEntryList*	- operation for the VPC to query the status of the route.
//
//   - If the route is in the **Pending*	- state, it is being created. During this time, you can only query the route and cannot perform other operations.
//
//   - If the route is in the **Available*	- state, the route is created.
//
// @param request - CreateCenChildInstanceRouteEntryToAttachmentRequest
//
// @return CreateCenChildInstanceRouteEntryToAttachmentResponse
func (client *Client) CreateCenChildInstanceRouteEntryToAttachment(request *CreateCenChildInstanceRouteEntryToAttachmentRequest) (_result *CreateCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCenChildInstanceRouteEntryToAttachmentResponse{}
	_body, _err := client.CreateCenChildInstanceRouteEntryToAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCenChildInstanceRouteEntryToCenWithOptions(request *CreateCenChildInstanceRouteEntryToCenRequest, runtime *dara.RuntimeOptions) (_result *CreateCenChildInstanceRouteEntryToCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCenChildInstanceRouteEntryToCenResponse
func (client *Client) CreateCenChildInstanceRouteEntryToCen(request *CreateCenChildInstanceRouteEntryToCenRequest) (_result *CreateCenChildInstanceRouteEntryToCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.CreateCenChildInstanceRouteEntryToCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a quality of service (QoS) policy for an inter-region connection on an Enterprise Edition transit router.
//
// Description:
//
// - Only inter-region connections created on Enterprise Edition transit routers support QoS policies.
//
// - Traffic scheduling applies only to outbound traffic on Enterprise Edition transit routers.
//
//	For example, you create an inter-region connection between the China (Hangzhou) and China (Qingdao) regions, and create a QoS policy for the transit router in the China (Hangzhou) region. In this case, the QoS policy can ensure bandwidth for network traffic from the China (Hangzhou) region to the China (Qingdao) region. However, the QoS policy does not apply to network traffic from the China (Qingdao) region to the China (Hangzhou) region.
//
// - **CreateCenInterRegionTrafficQosPolicy*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of the task.
//
//   - If the QoS policy is in the **Creating*	- state, the QoS policy is being created. You can query the QoS policy but cannot perform other operations on the QoS policy.
//
//   - If the QoS policy is in the **Active*	- state, the QoS policy is created.
//
// ### Prerequisites
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
func (client *Client) CreateCenInterRegionTrafficQosPolicyWithOptions(request *CreateCenInterRegionTrafficQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateCenInterRegionTrafficQosPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - Only inter-region connections created on Enterprise Edition transit routers support QoS policies.
//
// - Traffic scheduling applies only to outbound traffic on Enterprise Edition transit routers.
//
//	For example, you create an inter-region connection between the China (Hangzhou) and China (Qingdao) regions, and create a QoS policy for the transit router in the China (Hangzhou) region. In this case, the QoS policy can ensure bandwidth for network traffic from the China (Hangzhou) region to the China (Qingdao) region. However, the QoS policy does not apply to network traffic from the China (Qingdao) region to the China (Hangzhou) region.
//
// - **CreateCenInterRegionTrafficQosPolicy*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of the task.
//
//   - If the QoS policy is in the **Creating*	- state, the QoS policy is being created. You can query the QoS policy but cannot perform other operations on the QoS policy.
//
//   - If the QoS policy is in the **Active*	- state, the QoS policy is created.
//
// ### Prerequisites
//
// Before you call the **CreateCenInterRegionTrafficQosPolicy*	- operation, make sure that the following requirements are met:
//
// - An inter-region connection is created. For more information, see [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html).
//
// - A traffic marking policy is created. For more information, see [CreateTrafficMarkingPolicy](https://help.aliyun.com/document_detail/419025.html).
//
// @param request - CreateCenInterRegionTrafficQosPolicyRequest
//
// @return CreateCenInterRegionTrafficQosPolicyResponse
func (client *Client) CreateCenInterRegionTrafficQosPolicy(request *CreateCenInterRegionTrafficQosPolicyRequest) (_result *CreateCenInterRegionTrafficQosPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCenInterRegionTrafficQosPolicyResponse{}
	_body, _err := client.CreateCenInterRegionTrafficQosPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateCenInterRegionTrafficQosQueueWithOptions(request *CreateCenInterRegionTrafficQosQueueRequest, runtime *dara.RuntimeOptions) (_result *CreateCenInterRegionTrafficQosQueueResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateCenInterRegionTrafficQosQueueResponse
func (client *Client) CreateCenInterRegionTrafficQosQueue(request *CreateCenInterRegionTrafficQosQueueRequest) (_result *CreateCenInterRegionTrafficQosQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCenInterRegionTrafficQosQueueResponse{}
	_body, _err := client.CreateCenInterRegionTrafficQosQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - If a routing policy is in the **Creating*	- state, the routing policy is being created. In this case, you can query the routing policy but cannot perform other operations.
//
// - If a routing policy is in the **Active*	- state, the routing policy is created.
//
// @param request - CreateCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCenRouteMapResponse
func (client *Client) CreateCenRouteMapWithOptions(request *CreateCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *CreateCenRouteMapResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - If a routing policy is in the **Creating*	- state, the routing policy is being created. In this case, you can query the routing policy but cannot perform other operations.
//
// - If a routing policy is in the **Active*	- state, the routing policy is created.
//
// @param request - CreateCenRouteMapRequest
//
// @return CreateCenRouteMapResponse
func (client *Client) CreateCenRouteMap(request *CreateCenRouteMapRequest) (_result *CreateCenRouteMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCenRouteMapResponse{}
	_body, _err := client.CreateCenRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a flow log.
//
// Description:
//
// You can use flow logs to capture traffic that is transmitted over transit router instances and network instance connections. Network instance connections include inter-region connections, VPC connections, VPN connections, ECR connections, and VBR connections. Before you create a flow log, note the following:
//
// - Only Enterprise Edition transit routers support flow logs.
//
// - For inter-region connections, flow logs capture only outbound traffic from the transit router. Inbound traffic is not captured.
//
//	For example, an Elastic Compute Service (ECS) instance in the US (Silicon Valley) region accesses an ECS instance in the US (Virginia) region through Cloud Enterprise Network (CEN). If you create a flow log for the transit router in the US (Virginia) region, you can view messages sent from the ECS instance in the US (Virginia) region to the ECS instance in the US (Silicon Valley) region in the Simple Log Service console. However, you cannot view messages sent from the ECS instance in the US (Silicon Valley) region to the ECS instance in the US (Virginia) region. To view these messages, you must also create a flow log for the transit router in the US (Silicon Valley) region.
//
// - When a flow log captures traffic of a VPC connection, it captures only traffic transmitted over the transit router elastic network interface (ENI). To capture traffic transmitted over other ENIs in the VPC, see [VPC flow log overview](https://help.aliyun.com/document_detail/127150.html).
//
// - `CreateFlowlog` is an asynchronous operation. After you send a request, the system returns a flow log ID. However, the flow log is not immediately created. The system creates the flow log in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//   - If a flow log is in the **Creating*	- state, it is being created. In this state, you can only query the flow log.
//
//   - If a flow log is in the **Active*	- state, it is created.
//
// ### Prerequisites
//
// Before you create a flow log for a resource, make sure that the resource has been created. To create a resource, see the following topics:
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
func (client *Client) CreateFlowlogWithOptions(request *CreateFlowlogRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowlogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can use flow logs to capture traffic that is transmitted over transit router instances and network instance connections. Network instance connections include inter-region connections, VPC connections, VPN connections, ECR connections, and VBR connections. Before you create a flow log, note the following:
//
// - Only Enterprise Edition transit routers support flow logs.
//
// - For inter-region connections, flow logs capture only outbound traffic from the transit router. Inbound traffic is not captured.
//
//	For example, an Elastic Compute Service (ECS) instance in the US (Silicon Valley) region accesses an ECS instance in the US (Virginia) region through Cloud Enterprise Network (CEN). If you create a flow log for the transit router in the US (Virginia) region, you can view messages sent from the ECS instance in the US (Virginia) region to the ECS instance in the US (Silicon Valley) region in the Simple Log Service console. However, you cannot view messages sent from the ECS instance in the US (Silicon Valley) region to the ECS instance in the US (Virginia) region. To view these messages, you must also create a flow log for the transit router in the US (Silicon Valley) region.
//
// - When a flow log captures traffic of a VPC connection, it captures only traffic transmitted over the transit router elastic network interface (ENI). To capture traffic transmitted over other ENIs in the VPC, see [VPC flow log overview](https://help.aliyun.com/document_detail/127150.html).
//
// - `CreateFlowlog` is an asynchronous operation. After you send a request, the system returns a flow log ID. However, the flow log is not immediately created. The system creates the flow log in the background. You can call the `DescribeFlowlogs` operation to query the status of a flow log.
//
//   - If a flow log is in the **Creating*	- state, it is being created. In this state, you can only query the flow log.
//
//   - If a flow log is in the **Active*	- state, it is created.
//
// ### Prerequisites
//
// Before you create a flow log for a resource, make sure that the resource has been created. To create a resource, see the following topics:
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
// @return CreateFlowlogResponse
func (client *Client) CreateFlowlog(request *CreateFlowlogRequest) (_result *CreateFlowlogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlowlogResponse{}
	_body, _err := client.CreateFlowlogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a traffic marking policy. A traffic marking policy captures network traffic based on traffic classification rules and marks the traffic with the Differentiated Services Code Point (DSCP) values that you specify.
//
// Description:
//
// - Only Enterprise Edition transit routers support traffic marking policies.
//
// - **CreateTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a traffic marking policy ID and runs the task in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of a traffic marking policy.
//
//   - If a traffic marking policy is in the **Creating*	- state, the traffic marking policy is being created. You can query the traffic marking policy but cannot perform other operations.
//
//   - If a traffic marking policy is in the **Active*	- state, the traffic marking policy is created.
//
// @param request - CreateTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrafficMarkingPolicyResponse
func (client *Client) CreateTrafficMarkingPolicyWithOptions(request *CreateTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateTrafficMarkingPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - Only Enterprise Edition transit routers support traffic marking policies.
//
// - **CreateTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a traffic marking policy ID and runs the task in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of a traffic marking policy.
//
//   - If a traffic marking policy is in the **Creating*	- state, the traffic marking policy is being created. You can query the traffic marking policy but cannot perform other operations.
//
//   - If a traffic marking policy is in the **Active*	- state, the traffic marking policy is created.
//
// @param request - CreateTrafficMarkingPolicyRequest
//
// @return CreateTrafficMarkingPolicyResponse
func (client *Client) CreateTrafficMarkingPolicy(request *CreateTrafficMarkingPolicyRequest) (_result *CreateTrafficMarkingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTrafficMarkingPolicyResponse{}
	_body, _err := client.CreateTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an aggregate route.
//
// Description:
//
// After you add an aggregate route to the route table of an Enterprise Edition transit router, the transit router propagates the aggregate route only to the route tables of VPC instances that are associated with the transit router route table and have route synchronization enabled.
//
// Before you create an aggregate route, make sure that the following requirements are met. Otherwise, the Enterprise Edition transit router does not propagate the aggregate route to the route tables of VPC instances:
//
// - The VPC instance is associated with the route table of the Enterprise Edition transit router. For more information, see [AssociateTransitRouterAttachmentWithRouteTable](https://help.aliyun.com/document_detail/261242.html).
//
// - Route synchronization is enabled for the VPC instance. For more information, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261358.html).
//
// @param tmpReq - CreateTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouteTableAggregationResponse
func (client *Client) CreateTransitRouteTableAggregationWithOptions(tmpReq *CreateTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouteTableAggregationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// After you add an aggregate route to the route table of an Enterprise Edition transit router, the transit router propagates the aggregate route only to the route tables of VPC instances that are associated with the transit router route table and have route synchronization enabled.
//
// Before you create an aggregate route, make sure that the following requirements are met. Otherwise, the Enterprise Edition transit router does not propagate the aggregate route to the route tables of VPC instances:
//
// - The VPC instance is associated with the route table of the Enterprise Edition transit router. For more information, see [AssociateTransitRouterAttachmentWithRouteTable](https://help.aliyun.com/document_detail/261242.html).
//
// - Route synchronization is enabled for the VPC instance. For more information, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261358.html).
//
// @param request - CreateTransitRouteTableAggregationRequest
//
// @return CreateTransitRouteTableAggregationResponse
func (client *Client) CreateTransitRouteTableAggregation(request *CreateTransitRouteTableAggregationRequest) (_result *CreateTransitRouteTableAggregationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouteTableAggregationResponse{}
	_body, _err := client.CreateTransitRouteTableAggregationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the CreateTransitRouter operation to create an Enterprise Edition transit router instance.
//
// Description:
//
// - You can call the **CreateTransitRouter*	- operation to create an Enterprise Edition transit router instance. Enterprise Edition transit routers are available only in some regions. For more information about the supported regions, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - **CreateTransitRouter*	- is an asynchronous operation. After you send a request, the system returns an Enterprise Edition transit router instance ID, but the instance is still being created in the background. You can call the [ListTransitRouters](https://help.aliyun.com/document_detail/261219.html) operation to query the status of the Enterprise Edition transit router instance.
//
//   - If an Enterprise Edition transit router instance is in the **Creating*	- state, you can only query the instance and cannot perform other operations.
//
//   - If an Enterprise Edition transit router instance is in the **Active*	- state, the instance has been created.
//
// - You can create only one transit router instance in each region for a CEN instance.
//
// @param tmpReq - CreateTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterResponse
func (client *Client) CreateTransitRouterWithOptions(tmpReq *CreateTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CreateTransitRouter operation to create an Enterprise Edition transit router instance.
//
// Description:
//
// - You can call the **CreateTransitRouter*	- operation to create an Enterprise Edition transit router instance. Enterprise Edition transit routers are available only in some regions. For more information about the supported regions, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - **CreateTransitRouter*	- is an asynchronous operation. After you send a request, the system returns an Enterprise Edition transit router instance ID, but the instance is still being created in the background. You can call the [ListTransitRouters](https://help.aliyun.com/document_detail/261219.html) operation to query the status of the Enterprise Edition transit router instance.
//
//   - If an Enterprise Edition transit router instance is in the **Creating*	- state, you can only query the instance and cannot perform other operations.
//
//   - If an Enterprise Edition transit router instance is in the **Active*	- state, the instance has been created.
//
// - You can create only one transit router instance in each region for a CEN instance.
//
// @param request - CreateTransitRouterRequest
//
// @return CreateTransitRouterResponse
func (client *Client) CreateTransitRouter(request *CreateTransitRouterRequest) (_result *CreateTransitRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterResponse{}
	_body, _err := client.CreateTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A transit router CIDR block is a custom CIDR block that you can create for a transit router. This CIDR block is similar to a CIDR block that is used to assign an IP address to a router\\"s loopback interface. Call the CreateTransitRouterCidr operation to create a CIDR block for a transit router.
//
// Description:
//
// A transit router CIDR block is a custom CIDR block that you can create for a transit router. It is similar to a CIDR block used to assign an IP address to a router\\"s loopback interface. A transit router CIDR block is used to allocate IP addresses to network instance connections. For more information, see [Transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// The **CreateTransitRouterCidr*	- operation is used to add a CIDR block to a transit router only after the transit router is created.
//
// Before you create a transit router CIDR block, note the following information:
//
// - Only Enterprise Edition transit routers support CIDR blocks.
//
// - For more information about the limits on transit router CIDR blocks, see [Limits on transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// - A transit router supports up to five CIDR blocks. The subnet mask of each CIDR block must be 16 to 24 bits in length.
//
// - You cannot create CIDR blocks that are within 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16, or their subnets.
//
// - Each CIDR block must not conflict with the CIDR blocks of interconnected network instances in the Cloud Enterprise Network (CEN) instance.
//
// - Each CIDR block must be unique within the same CEN instance.
//
// - After you add a CIDR block to a transit router and create the first VPN connection on it, the system automatically allocates three CIDR blocks from the specified CIDR block. These three CIDR blocks are reserved by the system for creating VPN connections. The system then allocates IP addresses to IPsec connections from the remaining CIDR blocks.
//
//	You can call the [ListTransitRouterCidrAllocation](https://help.aliyun.com/document_detail/464173.html) operation to query the CIDR blocks that are reserved by the system or allocated to IPsec connections.
//
// @param request - CreateTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterCidrResponse
func (client *Client) CreateTransitRouterCidrWithOptions(request *CreateTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterCidrResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A transit router CIDR block is a custom CIDR block that you can create for a transit router. This CIDR block is similar to a CIDR block that is used to assign an IP address to a router\\"s loopback interface. Call the CreateTransitRouterCidr operation to create a CIDR block for a transit router.
//
// Description:
//
// A transit router CIDR block is a custom CIDR block that you can create for a transit router. It is similar to a CIDR block used to assign an IP address to a router\\"s loopback interface. A transit router CIDR block is used to allocate IP addresses to network instance connections. For more information, see [Transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// The **CreateTransitRouterCidr*	- operation is used to add a CIDR block to a transit router only after the transit router is created.
//
// Before you create a transit router CIDR block, note the following information:
//
// - Only Enterprise Edition transit routers support CIDR blocks.
//
// - For more information about the limits on transit router CIDR blocks, see [Limits on transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// - A transit router supports up to five CIDR blocks. The subnet mask of each CIDR block must be 16 to 24 bits in length.
//
// - You cannot create CIDR blocks that are within 100.64.0.0/10, 224.0.0.0/4, 127.0.0.0/8, or 169.254.0.0/16, or their subnets.
//
// - Each CIDR block must not conflict with the CIDR blocks of interconnected network instances in the Cloud Enterprise Network (CEN) instance.
//
// - Each CIDR block must be unique within the same CEN instance.
//
// - After you add a CIDR block to a transit router and create the first VPN connection on it, the system automatically allocates three CIDR blocks from the specified CIDR block. These three CIDR blocks are reserved by the system for creating VPN connections. The system then allocates IP addresses to IPsec connections from the remaining CIDR blocks.
//
//	You can call the [ListTransitRouterCidrAllocation](https://help.aliyun.com/document_detail/464173.html) operation to query the CIDR blocks that are reserved by the system or allocated to IPsec connections.
//
// @param request - CreateTransitRouterCidrRequest
//
// @return CreateTransitRouterCidrResponse
func (client *Client) CreateTransitRouterCidr(request *CreateTransitRouterCidrRequest) (_result *CreateTransitRouterCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterCidrResponse{}
	_body, _err := client.CreateTransitRouterCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the CreateTransitRouterEcrAttachment operation to connect an Express Connect Router (ECR) instance to a transit router in the same region.
//
// Description:
//
// - Only Enterprise Edition transit routers support ECR connections.
//
// - You can create an ECR connection on an Enterprise Edition transit router in one of the following two ways:
//
//   - If you have an Enterprise Edition transit router instance in the destination region, you can create an ECR connection by specifying **EcrId**, **RegionId**, and **TransitRouterId**.
//
//   - If you do not have an Enterprise Edition transit router instance in the destination region, you can create an ECR connection by specifying **EcrId**, **CenId**, and **RegionId**. When you create the ECR connection, the system automatically creates an Enterprise Edition transit router instance for you.
//
// - CreateTransitRouterEcrAttachment is an asynchronous operation. After you send a request, the system returns an ECR connection ID, but the connection is created in the background. You can call the ListTransitRouterEcrAttachments operation to query the status of the ECR connection.
//
//   - If an ECR connection is in the **Attaching*	- status, the connection is being created. In this status, you can only query the connection and cannot perform other operations.
//
//   - If an ECR connection is in the **Attached*	- status, the connection is created.
//
// - By default, after an ECR connection is created, it is not associated with any route table of the Enterprise Edition transit router for route learning or forwarding.
//
//	After the ECR connection is associated with a route table of the Enterprise Edition transit router for [route learning](https://help.aliyun.com/document_detail/468300.html), the system automatically propagates the routes from the ECR instance to the route table of the Enterprise Edition transit router.
//
// - After an ECR connection is created, the system automatically propagates routes from the route table of the Enterprise Edition transit router associated with the ECR connection to the route table of the ECR instance.
//
// ### Prerequisites
//
// - The Alibaba Cloud accounts that own the Enterprise Edition transit router and the ECR instance must belong to the same enterprise.
//
// - An Enterprise Edition transit router can connect to ECR instances that belong to the same account or different accounts. Before you create a cross-account ECR connection, you must obtain authorization from the account that owns the ECR instance. For more information, see [Authorize a cross-account network instance](https://help.aliyun.com/document_detail/181553.html).
//
// - **Before you call this operation to create an ECR connection, you must call the [CreateExpressConnectRouterAssociation](https://help.aliyun.com/document_detail/2712082.html) operation to create an association between the ECR instance and the Enterprise Edition transit router instance.**
//
//	**When you call the DeleteTransitRouterEcrAttachment operation to force delete an ECR connection, the system also deletes the association between the ECR instance and the Enterprise Edition transit router instance. You do not need to delete the association separately.**
//
// @param request - CreateTransitRouterEcrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterEcrAttachmentResponse
func (client *Client) CreateTransitRouterEcrAttachmentWithOptions(request *CreateTransitRouterEcrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterEcrAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the CreateTransitRouterEcrAttachment operation to connect an Express Connect Router (ECR) instance to a transit router in the same region.
//
// Description:
//
// - Only Enterprise Edition transit routers support ECR connections.
//
// - You can create an ECR connection on an Enterprise Edition transit router in one of the following two ways:
//
//   - If you have an Enterprise Edition transit router instance in the destination region, you can create an ECR connection by specifying **EcrId**, **RegionId**, and **TransitRouterId**.
//
//   - If you do not have an Enterprise Edition transit router instance in the destination region, you can create an ECR connection by specifying **EcrId**, **CenId**, and **RegionId**. When you create the ECR connection, the system automatically creates an Enterprise Edition transit router instance for you.
//
// - CreateTransitRouterEcrAttachment is an asynchronous operation. After you send a request, the system returns an ECR connection ID, but the connection is created in the background. You can call the ListTransitRouterEcrAttachments operation to query the status of the ECR connection.
//
//   - If an ECR connection is in the **Attaching*	- status, the connection is being created. In this status, you can only query the connection and cannot perform other operations.
//
//   - If an ECR connection is in the **Attached*	- status, the connection is created.
//
// - By default, after an ECR connection is created, it is not associated with any route table of the Enterprise Edition transit router for route learning or forwarding.
//
//	After the ECR connection is associated with a route table of the Enterprise Edition transit router for [route learning](https://help.aliyun.com/document_detail/468300.html), the system automatically propagates the routes from the ECR instance to the route table of the Enterprise Edition transit router.
//
// - After an ECR connection is created, the system automatically propagates routes from the route table of the Enterprise Edition transit router associated with the ECR connection to the route table of the ECR instance.
//
// ### Prerequisites
//
// - The Alibaba Cloud accounts that own the Enterprise Edition transit router and the ECR instance must belong to the same enterprise.
//
// - An Enterprise Edition transit router can connect to ECR instances that belong to the same account or different accounts. Before you create a cross-account ECR connection, you must obtain authorization from the account that owns the ECR instance. For more information, see [Authorize a cross-account network instance](https://help.aliyun.com/document_detail/181553.html).
//
// - **Before you call this operation to create an ECR connection, you must call the [CreateExpressConnectRouterAssociation](https://help.aliyun.com/document_detail/2712082.html) operation to create an association between the ECR instance and the Enterprise Edition transit router instance.**
//
//	**When you call the DeleteTransitRouterEcrAttachment operation to force delete an ECR connection, the system also deletes the association between the ECR instance and the Enterprise Edition transit router instance. You do not need to delete the association separately.**
//
// @param request - CreateTransitRouterEcrAttachmentRequest
//
// @return CreateTransitRouterEcrAttachmentResponse
func (client *Client) CreateTransitRouterEcrAttachment(request *CreateTransitRouterEcrAttachmentRequest) (_result *CreateTransitRouterEcrAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterEcrAttachmentResponse{}
	_body, _err := client.CreateTransitRouterEcrAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A multicast domain defines the scope of a multicast network in a region. Only resources within the multicast domain can send and receive multicast traffic. You can call the CreateTransitRouterMulticastDomain operation to create a multicast domain.
//
// Description:
//
// Before you call this operation, note the following:
//
// - Make sure that you have created an Enterprise Edition transit router in the region where you want to create the multicast network and enabled the multicast feature for the transit router. For more information, see [CreateTransitRouter](https://help.aliyun.com/document_detail/261169.html).
//
//	If you created an Enterprise Edition transit router before you requested multicast resources, you cannot enable the multicast feature for the transit router. You must delete the current Enterprise Edition transit router and create a new one. For more information about how to delete an Enterprise Edition transit router, see [DeleteTransitRouter](https://help.aliyun.com/document_detail/261218.html).
//
// - When you call the **CreateTransitRouterMulticastDomain*	- operation, if you specify **CenId*	- and **RegionId**, you do not need to specify **TransitRouterId**. If you specify **TransitRouterId**, you do not need to specify **CenId*	- or **RegionId**.
//
// @param request - CreateTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterMulticastDomainResponse
func (client *Client) CreateTransitRouterMulticastDomainWithOptions(request *CreateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterMulticastDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A multicast domain defines the scope of a multicast network in a region. Only resources within the multicast domain can send and receive multicast traffic. You can call the CreateTransitRouterMulticastDomain operation to create a multicast domain.
//
// Description:
//
// Before you call this operation, note the following:
//
// - Make sure that you have created an Enterprise Edition transit router in the region where you want to create the multicast network and enabled the multicast feature for the transit router. For more information, see [CreateTransitRouter](https://help.aliyun.com/document_detail/261169.html).
//
//	If you created an Enterprise Edition transit router before you requested multicast resources, you cannot enable the multicast feature for the transit router. You must delete the current Enterprise Edition transit router and create a new one. For more information about how to delete an Enterprise Edition transit router, see [DeleteTransitRouter](https://help.aliyun.com/document_detail/261218.html).
//
// - When you call the **CreateTransitRouterMulticastDomain*	- operation, if you specify **CenId*	- and **RegionId**, you do not need to specify **TransitRouterId**. If you specify **TransitRouterId**, you do not need to specify **CenId*	- or **RegionId**.
//
// @param request - CreateTransitRouterMulticastDomainRequest
//
// @return CreateTransitRouterMulticastDomainResponse
func (client *Client) CreateTransitRouterMulticastDomain(request *CreateTransitRouterMulticastDomainRequest) (_result *CreateTransitRouterMulticastDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterMulticastDomainResponse{}
	_body, _err := client.CreateTransitRouterMulticastDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After network instances (VPC, VBR, IPsec connection) are connected to a transit router, you need to create an inter-region connection to enable communication between network instances in different regions. You can call the CreateTransitRouterPeerAttachment operation to create an inter-region connection for an Enterprise Edition transit router instance.
//
// Description:
//
// - Enterprise Edition transit routers allow you to allocate bandwidth resources to inter-region connections using the following methods:
//
//   - **From bandwidth plan**:
//
//     You must purchase a bandwidth plan and then allocate bandwidth resources from the plan to inter-region connections. For more information about how to purchase a bandwidth plan, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html).
//
//   - **Pay-by-traffic**:
//
//     You can set a maximum bandwidth value for an inter-region connection. You are then charged based on the amount of data transferred over the connection. For more information about billing, see [Inter-region traffic](https://help.aliyun.com/document_detail/337827.html).
//
// - The **CreateTransitRouterPeerAttachment*	- operation is asynchronous. After you call this operation, the system returns an inter-region connection ID. However, the inter-region connection is not created immediately. The creation task runs in the background. You can call the **ListTransitRouterPeerAttachments*	- operation to query the status of the inter-region connection.
//
//   - When the inter-region connection is in the **Attaching*	- state, the connection is being created. In this state, you can only query the inter-region connection. You cannot perform other operations.
//
//   - When the inter-region connection is in the **Attached*	- state, the connection creation is complete.
//
// @param request - CreateTransitRouterPeerAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterPeerAttachmentResponse
func (client *Client) CreateTransitRouterPeerAttachmentWithOptions(request *CreateTransitRouterPeerAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterPeerAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After network instances (VPC, VBR, IPsec connection) are connected to a transit router, you need to create an inter-region connection to enable communication between network instances in different regions. You can call the CreateTransitRouterPeerAttachment operation to create an inter-region connection for an Enterprise Edition transit router instance.
//
// Description:
//
// - Enterprise Edition transit routers allow you to allocate bandwidth resources to inter-region connections using the following methods:
//
//   - **From bandwidth plan**:
//
//     You must purchase a bandwidth plan and then allocate bandwidth resources from the plan to inter-region connections. For more information about how to purchase a bandwidth plan, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html).
//
//   - **Pay-by-traffic**:
//
//     You can set a maximum bandwidth value for an inter-region connection. You are then charged based on the amount of data transferred over the connection. For more information about billing, see [Inter-region traffic](https://help.aliyun.com/document_detail/337827.html).
//
// - The **CreateTransitRouterPeerAttachment*	- operation is asynchronous. After you call this operation, the system returns an inter-region connection ID. However, the inter-region connection is not created immediately. The creation task runs in the background. You can call the **ListTransitRouterPeerAttachments*	- operation to query the status of the inter-region connection.
//
//   - When the inter-region connection is in the **Attaching*	- state, the connection is being created. In this state, you can only query the inter-region connection. You cannot perform other operations.
//
//   - When the inter-region connection is in the **Attached*	- state, the connection creation is complete.
//
// @param request - CreateTransitRouterPeerAttachmentRequest
//
// @return CreateTransitRouterPeerAttachmentResponse
func (client *Client) CreateTransitRouterPeerAttachment(request *CreateTransitRouterPeerAttachmentRequest) (_result *CreateTransitRouterPeerAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterPeerAttachmentResponse{}
	_body, _err := client.CreateTransitRouterPeerAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTransitRouterPrefixListAssociationWithOptions(request *CreateTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterPrefixListAssociationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateTransitRouterPrefixListAssociationResponse
func (client *Client) CreateTransitRouterPrefixListAssociation(request *CreateTransitRouterPrefixListAssociationRequest) (_result *CreateTransitRouterPrefixListAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterPrefixListAssociationResponse{}
	_body, _err := client.CreateTransitRouterPrefixListAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a route entry in the route table of an Enterprise Edition transit router.
//
// Description:
//
// *CreateTransitRouterRouteEntry*	- is an asynchronous operation. After you send a request, the system returns a route entry ID. The route entry is created in the background. You can call the **ListTransitRouterRouteEntries*	- operation to query the status of the route entry.
//
// - If a route entry is in the **Creating*	- state, the route entry is being created. In this state, you can only query the route entry and cannot perform other operations.
//
// - If a route entry is in the **Active*	- state, the route entry has been created.
//
// @param request - CreateTransitRouterRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterRouteEntryResponse
func (client *Client) CreateTransitRouterRouteEntryWithOptions(request *CreateTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterRouteEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route entry in the route table of an Enterprise Edition transit router.
//
// Description:
//
// *CreateTransitRouterRouteEntry*	- is an asynchronous operation. After you send a request, the system returns a route entry ID. The route entry is created in the background. You can call the **ListTransitRouterRouteEntries*	- operation to query the status of the route entry.
//
// - If a route entry is in the **Creating*	- state, the route entry is being created. In this state, you can only query the route entry and cannot perform other operations.
//
// - If a route entry is in the **Active*	- state, the route entry has been created.
//
// @param request - CreateTransitRouterRouteEntryRequest
//
// @return CreateTransitRouterRouteEntryResponse
func (client *Client) CreateTransitRouterRouteEntry(request *CreateTransitRouterRouteEntryRequest) (_result *CreateTransitRouterRouteEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterRouteEntryResponse{}
	_body, _err := client.CreateTransitRouterRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls the CreateTransitRouterRouteTable operation to create a custom route table for an Enterprise Edition transit router.
//
// Description:
//
// - You can create custom route tables only for Enterprise Edition transit routers. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - **CreateTransitRouterRouteTable*	- is an asynchronous operation. After you send a request, a route table ID is returned, but the route table is still being created in the background. You can call the **ListTransitRouterRouteTables*	- operation to query the status of a route table.
//
//   - If a route table is in the **Creating*	- state, it is being created. In this state, you can only query the route table and cannot perform other operations.
//
//   - If a route table is in the **Active*	- state, the route table is created.
//
// @param request - CreateTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterRouteTableResponse
func (client *Client) CreateTransitRouterRouteTableWithOptions(request *CreateTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterRouteTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls the CreateTransitRouterRouteTable operation to create a custom route table for an Enterprise Edition transit router.
//
// Description:
//
// - You can create custom route tables only for Enterprise Edition transit routers. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is Cloud Enterprise Network?](https://help.aliyun.com/document_detail/181681.html).
//
// - **CreateTransitRouterRouteTable*	- is an asynchronous operation. After you send a request, a route table ID is returned, but the route table is still being created in the background. You can call the **ListTransitRouterRouteTables*	- operation to query the status of a route table.
//
//   - If a route table is in the **Creating*	- state, it is being created. In this state, you can only query the route table and cannot perform other operations.
//
//   - If a route table is in the **Active*	- state, the route table is created.
//
// @param request - CreateTransitRouterRouteTableRequest
//
// @return CreateTransitRouterRouteTableResponse
func (client *Client) CreateTransitRouterRouteTable(request *CreateTransitRouterRouteTableRequest) (_result *CreateTransitRouterRouteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterRouteTableResponse{}
	_body, _err := client.CreateTransitRouterRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Connects a virtual border router (VBR) with a transit router in the same region.
//
// Description:
//
// - For information about the regions and zones supported by Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
// - You can create a VBR connection with or without an Enterprise Edition transit router:
//
//   - If you already have an Enterprise Edition transit router in the target region, specify the **VbrId**, **RegionId**, and **TransitRouterId*	- parameters.
//
//   - If you do not have an Enterprise Edition transit router in the target region, specify the **VbrId**, **CenId**, and **RegionId&#x20;**&#x70;arameters, and the system will automatically create an Enterprise Edition transit router when executing the operation.
//
// - This operation is executed asynchronously. After receiving a request, the system returns a VBR connection ID before the VBR connection is fully ready, and it continues the creation task in the backend. You can call **ListTransitRouterVbrAttachments*	- to check whether the connection has been created.
//
//   - If the VBR connection is in the **Attaching*	- state, it hasn\\"t been created. In this case, you can query information about the connection but cannot perform other operations on it.
//
//   - If the VBR connection is in the **Attached*	- state, the creation task has been completed.
//
// - The transit router and VBR can be in the same or different Alibaba Cloud accounts. In a cross-account scenario, both accounts must belong to the same enterprise, and you need to [grant the required permissions on the VBR to the transit router](https://help.aliyun.com/document_detail/181553.html).
//
// - A newly created VBR connection is not in route learning or associated forwarding correlations with any route table on the transit router.
//
// @param request - CreateTransitRouterVbrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVbrAttachmentResponse
func (client *Client) CreateTransitRouterVbrAttachmentWithOptions(request *CreateTransitRouterVbrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVbrAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Connects a virtual border router (VBR) with a transit router in the same region.
//
// Description:
//
// - For information about the regions and zones supported by Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
// - You can create a VBR connection with or without an Enterprise Edition transit router:
//
//   - If you already have an Enterprise Edition transit router in the target region, specify the **VbrId**, **RegionId**, and **TransitRouterId*	- parameters.
//
//   - If you do not have an Enterprise Edition transit router in the target region, specify the **VbrId**, **CenId**, and **RegionId&#x20;**&#x70;arameters, and the system will automatically create an Enterprise Edition transit router when executing the operation.
//
// - This operation is executed asynchronously. After receiving a request, the system returns a VBR connection ID before the VBR connection is fully ready, and it continues the creation task in the backend. You can call **ListTransitRouterVbrAttachments*	- to check whether the connection has been created.
//
//   - If the VBR connection is in the **Attaching*	- state, it hasn\\"t been created. In this case, you can query information about the connection but cannot perform other operations on it.
//
//   - If the VBR connection is in the **Attached*	- state, the creation task has been completed.
//
// - The transit router and VBR can be in the same or different Alibaba Cloud accounts. In a cross-account scenario, both accounts must belong to the same enterprise, and you need to [grant the required permissions on the VBR to the transit router](https://help.aliyun.com/document_detail/181553.html).
//
// - A newly created VBR connection is not in route learning or associated forwarding correlations with any route table on the transit router.
//
// @param request - CreateTransitRouterVbrAttachmentRequest
//
// @return CreateTransitRouterVbrAttachmentResponse
func (client *Client) CreateTransitRouterVbrAttachment(request *CreateTransitRouterVbrAttachmentRequest) (_result *CreateTransitRouterVbrAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterVbrAttachmentResponse{}
	_body, _err := client.CreateTransitRouterVbrAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a virtual private cloud (VPC) to a transit router. Once connected, the transit router enables private network communication.
//
// Description:
//
// - You can create a VPC connection for an Enterprise Edition transit router in one of two ways:
//
//   - If you have an Enterprise Edition transit router in the target region, you can create a VPC connection by specifying **VpcId**, **ZoneMappings.N.VSwitchId**, **ZoneMappings.N.ZoneId**, **TransitRouterId**, and **RegionId**.
//
//   - If you do not have an Enterprise Edition transit router in the target region, you can create a VPC connection by specifying **VpcId**, **ZoneMappings.N.VSwitchId**, **ZoneMappings.N.ZoneId**, **CenId**, and **RegionId**. When you create the VPC connection, the system automatically creates an Enterprise Edition transit router.
//
// - **CreateTransitRouterVpcAttachment*	- is an asynchronous operation. After you send a request, the system returns a VPC connection ID, and the connection is created in the background. Call the [ListTransitRouterVpcAttachments](https://help.aliyun.com/document_detail/261222.html) operation to query the status of the VPC connection.
//
//   - The **Attaching*	- state indicates that the VPC connection is being created. In this state, you can only query the VPC connection.
//
//   - The **Attached*	- state indicates that the VPC connection has been created.
//
// - By default, a newly created VPC connection is not associated with any transit router route table for route learning or forwarding.
//
// ### Prerequisites
//
// Before you call this API operation to create a VPC connection, ensure the following prerequisites are met:
//
// - The VPC must have at least one vSwitch in a zone that supports Enterprise Edition transit routers. The vSwitch must have at least one available IP address. For more information about supported regions and zones, see [Regions and zones that support Enterprise Edition transit routers](https://help.aliyun.com/document_detail/181681.html).
//
// - To connect a cross-account network instance, you must first have the required permissions. For more information, see [Grant permissions on a network instance that belongs to another account](https://help.aliyun.com/document_detail/181553.html).
//
// - Creating a VPC connection incurs costs. Ensure you understand the billing rules. For more information, see [Billing](https://help.aliyun.com/document_detail/189836.html).
//
// @param tmpReq - CreateTransitRouterVpcAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVpcAttachmentResponse
func (client *Client) CreateTransitRouterVpcAttachmentWithOptions(tmpReq *CreateTransitRouterVpcAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVpcAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a virtual private cloud (VPC) to a transit router. Once connected, the transit router enables private network communication.
//
// Description:
//
// - You can create a VPC connection for an Enterprise Edition transit router in one of two ways:
//
//   - If you have an Enterprise Edition transit router in the target region, you can create a VPC connection by specifying **VpcId**, **ZoneMappings.N.VSwitchId**, **ZoneMappings.N.ZoneId**, **TransitRouterId**, and **RegionId**.
//
//   - If you do not have an Enterprise Edition transit router in the target region, you can create a VPC connection by specifying **VpcId**, **ZoneMappings.N.VSwitchId**, **ZoneMappings.N.ZoneId**, **CenId**, and **RegionId**. When you create the VPC connection, the system automatically creates an Enterprise Edition transit router.
//
// - **CreateTransitRouterVpcAttachment*	- is an asynchronous operation. After you send a request, the system returns a VPC connection ID, and the connection is created in the background. Call the [ListTransitRouterVpcAttachments](https://help.aliyun.com/document_detail/261222.html) operation to query the status of the VPC connection.
//
//   - The **Attaching*	- state indicates that the VPC connection is being created. In this state, you can only query the VPC connection.
//
//   - The **Attached*	- state indicates that the VPC connection has been created.
//
// - By default, a newly created VPC connection is not associated with any transit router route table for route learning or forwarding.
//
// ### Prerequisites
//
// Before you call this API operation to create a VPC connection, ensure the following prerequisites are met:
//
// - The VPC must have at least one vSwitch in a zone that supports Enterprise Edition transit routers. The vSwitch must have at least one available IP address. For more information about supported regions and zones, see [Regions and zones that support Enterprise Edition transit routers](https://help.aliyun.com/document_detail/181681.html).
//
// - To connect a cross-account network instance, you must first have the required permissions. For more information, see [Grant permissions on a network instance that belongs to another account](https://help.aliyun.com/document_detail/181553.html).
//
// - Creating a VPC connection incurs costs. Ensure you understand the billing rules. For more information, see [Billing](https://help.aliyun.com/document_detail/189836.html).
//
// @param request - CreateTransitRouterVpcAttachmentRequest
//
// @return CreateTransitRouterVpcAttachmentResponse
func (client *Client) CreateTransitRouterVpcAttachment(request *CreateTransitRouterVpcAttachmentRequest) (_result *CreateTransitRouterVpcAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterVpcAttachmentResponse{}
	_body, _err := client.CreateTransitRouterVpcAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// A transit router supports IPsec-VPN connections, allowing a data center to connect to the transit router and communicate with other networks. You can call the CreateTransitRouterVpnAttachment operation to create a VPN connection.
//
// Description:
//
// - After you create a VPN connection, the connection does not learn routes from or associate with any route table of the transit router by default.
//
// - When you call the `CreateTransitRouterVpnAttachment` operation, you do not need to specify **TransitRouterId*	- if you specify **CenId*	- and **RegionId**. You do not need to specify **CenId*	- if you specify **TransitRouterId*	- and **RegionId**.
//
// ### Prerequisites
//
// - Before you create a VPN connection, make sure that you have created an IPsec-VPN connection in the same region as the transit router instance and that the IPsec-VPN connection is not attached to any resource. For more information, see [CreateVpnAttachment](https://help.aliyun.com/document_detail/442455.html).
//
// - If the transit router instance needs to connect to a cross-account IPsec-VPN connection, make sure that the transit router instance is granted the required permissions on the IPsec-VPN connection. For more information, see [GrantInstanceToTransitRouter](https://help.aliyun.com/document_detail/417520.html).
//
// - Before you create a VPN connection, make sure that you have configured a CIDR block for the transit router. For more information, see [CreateTransitRouterCidr](https://help.aliyun.com/document_detail/468230.html).
//
// @param request - CreateTransitRouterVpnAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTransitRouterVpnAttachmentResponse
func (client *Client) CreateTransitRouterVpnAttachmentWithOptions(request *CreateTransitRouterVpnAttachmentRequest, runtime *dara.RuntimeOptions) (_result *CreateTransitRouterVpnAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// A transit router supports IPsec-VPN connections, allowing a data center to connect to the transit router and communicate with other networks. You can call the CreateTransitRouterVpnAttachment operation to create a VPN connection.
//
// Description:
//
// - After you create a VPN connection, the connection does not learn routes from or associate with any route table of the transit router by default.
//
// - When you call the `CreateTransitRouterVpnAttachment` operation, you do not need to specify **TransitRouterId*	- if you specify **CenId*	- and **RegionId**. You do not need to specify **CenId*	- if you specify **TransitRouterId*	- and **RegionId**.
//
// ### Prerequisites
//
// - Before you create a VPN connection, make sure that you have created an IPsec-VPN connection in the same region as the transit router instance and that the IPsec-VPN connection is not attached to any resource. For more information, see [CreateVpnAttachment](https://help.aliyun.com/document_detail/442455.html).
//
// - If the transit router instance needs to connect to a cross-account IPsec-VPN connection, make sure that the transit router instance is granted the required permissions on the IPsec-VPN connection. For more information, see [GrantInstanceToTransitRouter](https://help.aliyun.com/document_detail/417520.html).
//
// - Before you create a VPN connection, make sure that you have configured a CIDR block for the transit router. For more information, see [CreateTransitRouterCidr](https://help.aliyun.com/document_detail/468230.html).
//
// @param request - CreateTransitRouterVpnAttachmentRequest
//
// @return CreateTransitRouterVpnAttachmentResponse
func (client *Client) CreateTransitRouterVpnAttachment(request *CreateTransitRouterVpnAttachmentRequest) (_result *CreateTransitRouterVpnAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTransitRouterVpnAttachmentResponse{}
	_body, _err := client.CreateTransitRouterVpnAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a flow log. A disabled flow log no longer captures network traffic.
//
// Description:
//
// This operation is executed asynchronously. After receiving a request, the system returns a **request ID*	- before it finishes disabling the flow log. The task is continued in the backend. You can call `DescribeFlowlogs` to check whether the flow log has been disabled.
//
// - If the flow log is in the **Modifying*	- state, the task is still in progress. In this case, you can query information about the flow log but cannot perform other operations on it.
//
// - If the flow log is in the **Inactive*	- state, the flow log has been disabled.
//
// @param request - DeactiveFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactiveFlowLogResponse
func (client *Client) DeactiveFlowLogWithOptions(request *DeactiveFlowLogRequest, runtime *dara.RuntimeOptions) (_result *DeactiveFlowLogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a flow log. A disabled flow log no longer captures network traffic.
//
// Description:
//
// This operation is executed asynchronously. After receiving a request, the system returns a **request ID*	- before it finishes disabling the flow log. The task is continued in the backend. You can call `DescribeFlowlogs` to check whether the flow log has been disabled.
//
// - If the flow log is in the **Modifying*	- state, the task is still in progress. In this case, you can query information about the flow log but cannot perform other operations on it.
//
// - If the flow log is in the **Inactive*	- state, the flow log has been disabled.
//
// @param request - DeactiveFlowLogRequest
//
// @return DeactiveFlowLogResponse
func (client *Client) DeactiveFlowLog(request *DeactiveFlowLogRequest) (_result *DeactiveFlowLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeactiveFlowLogResponse{}
	_body, _err := client.DeactiveFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - If the CEN instance is in the **Deleting*	- state, the CEN instance is being deleted. In this case, you can query the CEN instance but cannot perform other operations.
//
// - If the CEN instance cannot be found, the CEN instance is deleted.
//
// ### [](#)Prerequisites
//
// The CEN instance that you want to delete is not associated with a bandwidth plan, and the transit router associated with the CEN instance does not have a network instance connection or a custom route table.
//
// - For more information about how to detach a network instance, see the following topics:
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
// - For more information about how to delete custom route tables from an Enterprise Edition transit router, see [DeleteTransitRouterRouteTable](https://help.aliyun.com/document_detail/261235.html).
//
// - For more information about how to disassociate a bandwidth plan from a CEN instance, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// @param request - DeleteCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenResponse
func (client *Client) DeleteCenWithOptions(request *DeleteCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - If the CEN instance is in the **Deleting*	- state, the CEN instance is being deleted. In this case, you can query the CEN instance but cannot perform other operations.
//
// - If the CEN instance cannot be found, the CEN instance is deleted.
//
// ### [](#)Prerequisites
//
// The CEN instance that you want to delete is not associated with a bandwidth plan, and the transit router associated with the CEN instance does not have a network instance connection or a custom route table.
//
// - For more information about how to detach a network instance, see the following topics:
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
// - For more information about how to delete custom route tables from an Enterprise Edition transit router, see [DeleteTransitRouterRouteTable](https://help.aliyun.com/document_detail/261235.html).
//
// - For more information about how to disassociate a bandwidth plan from a CEN instance, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// @param request - DeleteCenRequest
//
// @return DeleteCenResponse
func (client *Client) DeleteCen(request *DeleteCenRequest) (_result *DeleteCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCenResponse{}
	_body, _err := client.DeleteCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a bandwidth plan.
//
// Description:
//
// <props="china">
//
// - Before you delete a bandwidth plan, ensure that it is detached from the Cloud Enterprise Network (CEN) instance. For more information, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// - If you want to delete a prepay bandwidth plan, you must go to the [Order Center](https://usercenter2.aliyun.com/refund/refund) to unsubscribe from the bandwidth plan. If you have questions about unsubscription, see [Unsubscription rules](https://help.aliyun.com/zh/user-center/user-guide/unsubscription-rules#p-1qo-3ce-m7z). This operation does not support deleting bandwidth plans that use the subscription billing method.
//
// <props="intl">
//
// Before you delete a bandwidth plan, ensure that it is detached from the Cloud Enterprise Network (CEN) instance. For more information, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// @param request - DeleteCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenBandwidthPackageResponse
func (client *Client) DeleteCenBandwidthPackageWithOptions(request *DeleteCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a bandwidth plan.
//
// Description:
//
// <props="china">
//
// - Before you delete a bandwidth plan, ensure that it is detached from the Cloud Enterprise Network (CEN) instance. For more information, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// - If you want to delete a prepay bandwidth plan, you must go to the [Order Center](https://usercenter2.aliyun.com/refund/refund) to unsubscribe from the bandwidth plan. If you have questions about unsubscription, see [Unsubscription rules](https://help.aliyun.com/zh/user-center/user-guide/unsubscription-rules#p-1qo-3ce-m7z). This operation does not support deleting bandwidth plans that use the subscription billing method.
//
// <props="intl">
//
// Before you delete a bandwidth plan, ensure that it is detached from the Cloud Enterprise Network (CEN) instance. For more information, see [UnassociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65935.html).
//
// @param request - DeleteCenBandwidthPackageRequest
//
// @return DeleteCenBandwidthPackageResponse
func (client *Client) DeleteCenBandwidthPackage(request *DeleteCenBandwidthPackageRequest) (_result *DeleteCenBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCenBandwidthPackageResponse{}
	_body, _err := client.DeleteCenBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DeleteCenChildInstanceRouteEntryToAttachment operation to delete a route entry from a network instance that is connected to an Enterprise Edition transit router.
//
// Description:
//
// - You can delete route entries from Virtual Private Cloud (VPC) instances and virtual border router (VBR) instances only if the next hop of the route entry is a **transit router connection*	- (a network instance connection).
//
// - **DeleteCenChildInstanceRouteEntryToAttachment*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the route entry is still being deleted in the background. You can call the **DescribeRouteEntryList*	- operation for VPC to query the status of the route entry.
//
//   - If a route entry is in the **Deleting*	- state, it is being deleted. During this time, you can only query the route entry and cannot perform other operations on it.
//
//   - If you cannot find the specified route entry, it has been deleted.
//
// @param request - DeleteCenChildInstanceRouteEntryToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenChildInstanceRouteEntryToAttachmentResponse
func (client *Client) DeleteCenChildInstanceRouteEntryToAttachmentWithOptions(request *DeleteCenChildInstanceRouteEntryToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DeleteCenChildInstanceRouteEntryToAttachment operation to delete a route entry from a network instance that is connected to an Enterprise Edition transit router.
//
// Description:
//
// - You can delete route entries from Virtual Private Cloud (VPC) instances and virtual border router (VBR) instances only if the next hop of the route entry is a **transit router connection*	- (a network instance connection).
//
// - **DeleteCenChildInstanceRouteEntryToAttachment*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the route entry is still being deleted in the background. You can call the **DescribeRouteEntryList*	- operation for VPC to query the status of the route entry.
//
//   - If a route entry is in the **Deleting*	- state, it is being deleted. During this time, you can only query the route entry and cannot perform other operations on it.
//
//   - If you cannot find the specified route entry, it has been deleted.
//
// @param request - DeleteCenChildInstanceRouteEntryToAttachmentRequest
//
// @return DeleteCenChildInstanceRouteEntryToAttachmentResponse
func (client *Client) DeleteCenChildInstanceRouteEntryToAttachment(request *DeleteCenChildInstanceRouteEntryToAttachmentRequest) (_result *DeleteCenChildInstanceRouteEntryToAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCenChildInstanceRouteEntryToAttachmentResponse{}
	_body, _err := client.DeleteCenChildInstanceRouteEntryToAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DeleteCenChildInstanceRouteEntryToCen operation to delete a route from a network instance.
//
// Description:
//
// - The DeleteCenChildInstanceRouteEntryToCen operation is not available by default. To use this operation, <props="china">[submit a ticket](https://selfservice.console.aliyun.com/ticket/category/cbn/today)<props="intl">[submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
// - You cannot use the DeleteCenChildInstanceRouteEntryToCen operation to delete routes from a network instance that is attached to an Enterprise Edition transit router.
//
// @param request - DeleteCenChildInstanceRouteEntryToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenChildInstanceRouteEntryToCenResponse
func (client *Client) DeleteCenChildInstanceRouteEntryToCenWithOptions(request *DeleteCenChildInstanceRouteEntryToCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenChildInstanceRouteEntryToCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DeleteCenChildInstanceRouteEntryToCen operation to delete a route from a network instance.
//
// Description:
//
// - The DeleteCenChildInstanceRouteEntryToCen operation is not available by default. To use this operation, <props="china">[submit a ticket](https://selfservice.console.aliyun.com/ticket/category/cbn/today)<props="intl">[submit a ticket](https://workorder-intl.console.aliyun.com/#/ticket/createIndex).
//
// - You cannot use the DeleteCenChildInstanceRouteEntryToCen operation to delete routes from a network instance that is attached to an Enterprise Edition transit router.
//
// @param request - DeleteCenChildInstanceRouteEntryToCenRequest
//
// @return DeleteCenChildInstanceRouteEntryToCenResponse
func (client *Client) DeleteCenChildInstanceRouteEntryToCen(request *DeleteCenChildInstanceRouteEntryToCenRequest) (_result *DeleteCenChildInstanceRouteEntryToCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCenChildInstanceRouteEntryToCenResponse{}
	_body, _err := client.DeleteCenChildInstanceRouteEntryToCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a quality of service (QoS) policy.
//
// Description:
//
// - Before you delete a QoS policy, you must delete all queues in the QoS policy except the default queue. For more information, see [DeleteCenInterRegionTrafficQosQueue](https://help.aliyun.com/document_detail/419062.html).
//
// - **DeleteCenInterRegionTrafficQosPolicy*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of a QoS policy.
//
//   - If a QoS policy is in the **Deleting*	- state, the QoS policy is being deleted. You can query the QoS policy but cannot perform other operations.
//
//   - If a QoS policy cannot be found, the QoS policy is deleted.
//
// @param request - DeleteCenInterRegionTrafficQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenInterRegionTrafficQosPolicyResponse
func (client *Client) DeleteCenInterRegionTrafficQosPolicyWithOptions(request *DeleteCenInterRegionTrafficQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenInterRegionTrafficQosPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - Before you delete a QoS policy, you must delete all queues in the QoS policy except the default queue. For more information, see [DeleteCenInterRegionTrafficQosQueue](https://help.aliyun.com/document_detail/419062.html).
//
// - **DeleteCenInterRegionTrafficQosPolicy*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of a QoS policy.
//
//   - If a QoS policy is in the **Deleting*	- state, the QoS policy is being deleted. You can query the QoS policy but cannot perform other operations.
//
//   - If a QoS policy cannot be found, the QoS policy is deleted.
//
// @param request - DeleteCenInterRegionTrafficQosPolicyRequest
//
// @return DeleteCenInterRegionTrafficQosPolicyResponse
func (client *Client) DeleteCenInterRegionTrafficQosPolicy(request *DeleteCenInterRegionTrafficQosPolicyRequest) (_result *DeleteCenInterRegionTrafficQosPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCenInterRegionTrafficQosPolicyResponse{}
	_body, _err := client.DeleteCenInterRegionTrafficQosPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a queue from a quality of service (QoS) policy.
//
// Description:
//
// - You cannot delete the default queue.
//
// - **DeleteCenInterRegionTrafficQosQueue*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of a queue. If a queue cannot be found, the queue is deleted.
//
// @param request - DeleteCenInterRegionTrafficQosQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenInterRegionTrafficQosQueueResponse
func (client *Client) DeleteCenInterRegionTrafficQosQueueWithOptions(request *DeleteCenInterRegionTrafficQosQueueRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenInterRegionTrafficQosQueueResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - You cannot delete the default queue.
//
// - **DeleteCenInterRegionTrafficQosQueue*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListCenInterRegionTrafficQosPolicies*	- operation to query the status of a queue. If a queue cannot be found, the queue is deleted.
//
// @param request - DeleteCenInterRegionTrafficQosQueueRequest
//
// @return DeleteCenInterRegionTrafficQosQueueResponse
func (client *Client) DeleteCenInterRegionTrafficQosQueue(request *DeleteCenInterRegionTrafficQosQueueRequest) (_result *DeleteCenInterRegionTrafficQosQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCenInterRegionTrafficQosQueueResponse{}
	_body, _err := client.DeleteCenInterRegionTrafficQosQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The DeleteCenRouteMap operation deletes a specified routing policy.
//
// Description:
//
// `DeleteCenRouteMap` is an asynchronous operation. After you call this operation, the system returns a request ID. The routing policy is then deleted in the background. You can call the `DescribeCenRouteMaps` operation to query the status of the routing policy.
//
// - If a routing policy is in the **Deleting*	- state, it is being deleted, and you can only perform query operations on it.
//
// - If the routing policy cannot be found when you call `DescribeCenRouteMaps`, the policy has been deleted.
//
// @param request - DeleteCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCenRouteMapResponse
func (client *Client) DeleteCenRouteMapWithOptions(request *DeleteCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *DeleteCenRouteMapResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The DeleteCenRouteMap operation deletes a specified routing policy.
//
// Description:
//
// `DeleteCenRouteMap` is an asynchronous operation. After you call this operation, the system returns a request ID. The routing policy is then deleted in the background. You can call the `DescribeCenRouteMaps` operation to query the status of the routing policy.
//
// - If a routing policy is in the **Deleting*	- state, it is being deleted, and you can only perform query operations on it.
//
// - If the routing policy cannot be found when you call `DescribeCenRouteMaps`, the policy has been deleted.
//
// @param request - DeleteCenRouteMapRequest
//
// @return DeleteCenRouteMapResponse
func (client *Client) DeleteCenRouteMap(request *DeleteCenRouteMapRequest) (_result *DeleteCenRouteMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCenRouteMapResponse{}
	_body, _err := client.DeleteCenRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a flow log.
//
// Description:
//
// This operation is executed asynchronously. After receiving a request, the system returns a **request ID*	- before it finishes deleting the flow log. The task is continued in the backend. You can call `DescribeFlowlogs` to check whether the flow log has been deleted.
//
// - If the flow log is in the **Deleting*	- state, it is still being deleted. In this case, you can query information about the flow log but cannot perform other operations on it.
//
// - If the `DescribeFlowlogs` call fails because the flow log is not found, it has been deleted.
//
// @param request - DeleteFlowlogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlogWithOptions(request *DeleteFlowlogRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowlogResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// This operation is executed asynchronously. After receiving a request, the system returns a **request ID*	- before it finishes deleting the flow log. The task is continued in the backend. You can call `DescribeFlowlogs` to check whether the flow log has been deleted.
//
// - If the flow log is in the **Deleting*	- state, it is still being deleted. In this case, you can query information about the flow log but cannot perform other operations on it.
//
// - If the `DescribeFlowlogs` call fails because the flow log is not found, it has been deleted.
//
// @param request - DeleteFlowlogRequest
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlog(request *DeleteFlowlogRequest) (_result *DeleteFlowlogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFlowlogResponse{}
	_body, _err := client.DeleteFlowlogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteRouteServiceInCenWithOptions(request *DeleteRouteServiceInCenRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteServiceInCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteRouteServiceInCenResponse
func (client *Client) DeleteRouteServiceInCen(request *DeleteRouteServiceInCenRequest) (_result *DeleteRouteServiceInCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRouteServiceInCenResponse{}
	_body, _err := client.DeleteRouteServiceInCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteTrafficMarkingPolicyWithOptions(request *DeleteTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteTrafficMarkingPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteTrafficMarkingPolicyResponse
func (client *Client) DeleteTrafficMarkingPolicy(request *DeleteTrafficMarkingPolicyRequest) (_result *DeleteTrafficMarkingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTrafficMarkingPolicyResponse{}
	_body, _err := client.DeleteTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an aggregate route.
//
// Description:
//
// - Before you delete an aggregate route, make sure that your network has a redundant route to prevent service interruptions.
//
// - After an aggregate route is deleted, the aggregate route is automatically withdrawn from virtual private clouds (VPCs). Specific routes that fall within the aggregate route are advertised to the VPCs.
//
// @param request - DeleteTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouteTableAggregationResponse
func (client *Client) DeleteTransitRouteTableAggregationWithOptions(request *DeleteTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouteTableAggregationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - Before you delete an aggregate route, make sure that your network has a redundant route to prevent service interruptions.
//
// - After an aggregate route is deleted, the aggregate route is automatically withdrawn from virtual private clouds (VPCs). Specific routes that fall within the aggregate route are advertised to the VPCs.
//
// @param request - DeleteTransitRouteTableAggregationRequest
//
// @return DeleteTransitRouteTableAggregationResponse
func (client *Client) DeleteTransitRouteTableAggregation(request *DeleteTransitRouteTableAggregationRequest) (_result *DeleteTransitRouteTableAggregationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouteTableAggregationResponse{}
	_body, _err := client.DeleteTransitRouteTableAggregationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - If a transit router is in the **Deleting*	- state, the transit router is being deleted. In this case, you can query the transit router but cannot perform other operations.
//
// - If a transit router cannot be found, the transit router is deleted.
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
func (client *Client) DeleteTransitRouterWithOptions(request *DeleteTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - If a transit router is in the **Deleting*	- state, the transit router is being deleted. In this case, you can query the transit router but cannot perform other operations.
//
// - If a transit router cannot be found, the transit router is deleted.
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
// @return DeleteTransitRouterResponse
func (client *Client) DeleteTransitRouter(request *DeleteTransitRouterRequest) (_result *DeleteTransitRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterResponse{}
	_body, _err := client.DeleteTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a CIDR block from a transit router.
//
// Description:
//
// You cannot delete a CIDR block from which IP addresses have been allocated.
//
// @param request - DeleteTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterCidrResponse
func (client *Client) DeleteTransitRouterCidrWithOptions(request *DeleteTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterCidrResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You cannot delete a CIDR block from which IP addresses have been allocated.
//
// @param request - DeleteTransitRouterCidrRequest
//
// @return DeleteTransitRouterCidrResponse
func (client *Client) DeleteTransitRouterCidr(request *DeleteTransitRouterCidrRequest) (_result *DeleteTransitRouterCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterCidrResponse{}
	_body, _err := client.DeleteTransitRouterCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteTransitRouterEcrAttachmentWithOptions(request *DeleteTransitRouterEcrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterEcrAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteTransitRouterEcrAttachmentResponse
func (client *Client) DeleteTransitRouterEcrAttachment(request *DeleteTransitRouterEcrAttachmentRequest) (_result *DeleteTransitRouterEcrAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterEcrAttachmentResponse{}
	_body, _err := client.DeleteTransitRouterEcrAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DeleteTransitRouterMulticastDomain operation to delete a multicast domain.
//
// Description:
//
// Before you delete a multicast domain, ensure that the following requirements are met:
//
// - The multicast domain is not associated with any vSwitch. For more information, see [DisassociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429774.html).
//
// - No multicast source or member exists in the multicast domain. For more information, see [DeregisterTransitRouterMulticastGroupSources](https://help.aliyun.com/document_detail/429776.html) and [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/429779.html).
//
// - The multicast domain is not associated with another multicast domain as a member. To disassociate the domains, delete the member from the other multicast domain. For more information, see [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/429779.html).
//
// - Ensure that you enter the correct parameter values when you call the operation. If you enter an incorrect parameter value, a request ID is returned, but the multicast domain is not deleted.
//
// @param request - DeleteTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterMulticastDomainResponse
func (client *Client) DeleteTransitRouterMulticastDomainWithOptions(request *DeleteTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterMulticastDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DeleteTransitRouterMulticastDomain operation to delete a multicast domain.
//
// Description:
//
// Before you delete a multicast domain, ensure that the following requirements are met:
//
// - The multicast domain is not associated with any vSwitch. For more information, see [DisassociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429774.html).
//
// - No multicast source or member exists in the multicast domain. For more information, see [DeregisterTransitRouterMulticastGroupSources](https://help.aliyun.com/document_detail/429776.html) and [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/429779.html).
//
// - The multicast domain is not associated with another multicast domain as a member. To disassociate the domains, delete the member from the other multicast domain. For more information, see [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/429779.html).
//
// - Ensure that you enter the correct parameter values when you call the operation. If you enter an incorrect parameter value, a request ID is returned, but the multicast domain is not deleted.
//
// @param request - DeleteTransitRouterMulticastDomainRequest
//
// @return DeleteTransitRouterMulticastDomainResponse
func (client *Client) DeleteTransitRouterMulticastDomain(request *DeleteTransitRouterMulticastDomainRequest) (_result *DeleteTransitRouterMulticastDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterMulticastDomainResponse{}
	_body, _err := client.DeleteTransitRouterMulticastDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DeleteTransitRouterPeerAttachment operation to delete an inter-region connection from an Enterprise Edition transit router.
//
// Description:
//
// The **DeleteTransitRouterPeerAttachment*	- operation is asynchronous. After you send a request, the system returns a **RequestId**, but the inter-region connection is not immediately deleted. The system deletes the connection in the background. You can call the **ListTransitRouterPeerAttachments*	- operation to query the status of the inter-region connection.
//
// - If an inter-region connection is in the **Detaching*	- state, it is being deleted. In this state, you can only query the connection and cannot perform other operations.
//
// - If the specified inter-region connection is not found, the connection has been deleted.
//
// Make sure that you specify valid parameter values when you call the **DeleteTransitRouterPeerAttachment*	- operation. If you specify an invalid parameter value, the system returns a **RequestId*	- but does not delete the inter-region connection.
//
// @param request - DeleteTransitRouterPeerAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterPeerAttachmentResponse
func (client *Client) DeleteTransitRouterPeerAttachmentWithOptions(request *DeleteTransitRouterPeerAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterPeerAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DeleteTransitRouterPeerAttachment operation to delete an inter-region connection from an Enterprise Edition transit router.
//
// Description:
//
// The **DeleteTransitRouterPeerAttachment*	- operation is asynchronous. After you send a request, the system returns a **RequestId**, but the inter-region connection is not immediately deleted. The system deletes the connection in the background. You can call the **ListTransitRouterPeerAttachments*	- operation to query the status of the inter-region connection.
//
// - If an inter-region connection is in the **Detaching*	- state, it is being deleted. In this state, you can only query the connection and cannot perform other operations.
//
// - If the specified inter-region connection is not found, the connection has been deleted.
//
// Make sure that you specify valid parameter values when you call the **DeleteTransitRouterPeerAttachment*	- operation. If you specify an invalid parameter value, the system returns a **RequestId*	- but does not delete the inter-region connection.
//
// @param request - DeleteTransitRouterPeerAttachmentRequest
//
// @return DeleteTransitRouterPeerAttachmentResponse
func (client *Client) DeleteTransitRouterPeerAttachment(request *DeleteTransitRouterPeerAttachmentRequest) (_result *DeleteTransitRouterPeerAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterPeerAttachmentResponse{}
	_body, _err := client.DeleteTransitRouterPeerAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteTransitRouterPrefixListAssociationWithOptions(request *DeleteTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterPrefixListAssociationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteTransitRouterPrefixListAssociationResponse
func (client *Client) DeleteTransitRouterPrefixListAssociation(request *DeleteTransitRouterPrefixListAssociationRequest) (_result *DeleteTransitRouterPrefixListAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterPrefixListAssociationResponse{}
	_body, _err := client.DeleteTransitRouterPrefixListAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DeleteTransitRouterRouteEntry operation to delete a static route of the Blackhole or Attachment type from the route table of an Enterprise Edition transit router.
//
// Description:
//
// Before you call this operation, note the following:
//
// - If you delete a route entry by specifying **TransitRouterRouteEntryId**, you do not need to specify the **TransitRouterRouteTableId*	- or **TransitRouterRouteEntryDestinationCidrBlock*	- parameters. These parameters are mutually exclusive.
//
// - If you do not specify **TransitRouterRouteEntryId**, you must specify the required parameters based on the next hop type:
//
//   - To delete a blackhole route, specify the **TransitRouterRouteTableId**, **TransitRouterRouteEntryDestinationCidrBlock**, and **TransitRouterRouteEntryNextHopType*	- parameters.
//
//   - To delete a route that is not a blackhole route, specify the **TransitRouterRouteTableId**, **TransitRouterRouteEntryDestinationCidrBlock**, **TransitRouterRouteEntryNextHopType**, and **TransitRouterRouteEntryNextHopId*	- parameters.
//
// - **DeleteTransitRouterRouteEntry*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the route entry is not immediately deleted. The system deletes the route entry in the background. You can call the **ListTransitRouterRouteEntries*	- operation to query the status of the route entry.
//
//   - If a route entry is in the **Deleting*	- state, it is being deleted. You can only query the route entry and cannot perform other operations.
//
//   - If the specified route entry cannot be found, the route entry has been deleted.
//
// ### Limits
//
// This operation deletes only static routes. It cannot delete routes that are automatically learned by the system. To query the type of a route entry, call the [ListTransitRouterRouteEntries](https://help.aliyun.com/document_detail/260941.html) operation.
//
// @param request - DeleteTransitRouterRouteEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterRouteEntryResponse
func (client *Client) DeleteTransitRouterRouteEntryWithOptions(request *DeleteTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterRouteEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DeleteTransitRouterRouteEntry operation to delete a static route of the Blackhole or Attachment type from the route table of an Enterprise Edition transit router.
//
// Description:
//
// Before you call this operation, note the following:
//
// - If you delete a route entry by specifying **TransitRouterRouteEntryId**, you do not need to specify the **TransitRouterRouteTableId*	- or **TransitRouterRouteEntryDestinationCidrBlock*	- parameters. These parameters are mutually exclusive.
//
// - If you do not specify **TransitRouterRouteEntryId**, you must specify the required parameters based on the next hop type:
//
//   - To delete a blackhole route, specify the **TransitRouterRouteTableId**, **TransitRouterRouteEntryDestinationCidrBlock**, and **TransitRouterRouteEntryNextHopType*	- parameters.
//
//   - To delete a route that is not a blackhole route, specify the **TransitRouterRouteTableId**, **TransitRouterRouteEntryDestinationCidrBlock**, **TransitRouterRouteEntryNextHopType**, and **TransitRouterRouteEntryNextHopId*	- parameters.
//
// - **DeleteTransitRouterRouteEntry*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the route entry is not immediately deleted. The system deletes the route entry in the background. You can call the **ListTransitRouterRouteEntries*	- operation to query the status of the route entry.
//
//   - If a route entry is in the **Deleting*	- state, it is being deleted. You can only query the route entry and cannot perform other operations.
//
//   - If the specified route entry cannot be found, the route entry has been deleted.
//
// ### Limits
//
// This operation deletes only static routes. It cannot delete routes that are automatically learned by the system. To query the type of a route entry, call the [ListTransitRouterRouteEntries](https://help.aliyun.com/document_detail/260941.html) operation.
//
// @param request - DeleteTransitRouterRouteEntryRequest
//
// @return DeleteTransitRouterRouteEntryResponse
func (client *Client) DeleteTransitRouterRouteEntry(request *DeleteTransitRouterRouteEntryRequest) (_result *DeleteTransitRouterRouteEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterRouteEntryResponse{}
	_body, _err := client.DeleteTransitRouterRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DeleteTransitRouterRouteTable operation to delete a custom route table of an Enterprise Edition transit router.
//
// Description:
//
// - You cannot delete the default route table of an Enterprise Edition transit router.
//
// - **DeleteTransitRouterRouteTable*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**. The custom route table is not deleted immediately because the system deletes the route table in the background. You can call the **ListTransitRouterRouteTables*	- operation to query the status of the custom route table.
//
//   - If a custom route table is in the Deleting state, the route table is being deleted. In this state, you can only query the route table. You cannot perform other operations.
//
//   - If the specified custom route table cannot be found, the route table has been deleted.
//
// @param request - DeleteTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterRouteTableResponse
func (client *Client) DeleteTransitRouterRouteTableWithOptions(request *DeleteTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterRouteTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DeleteTransitRouterRouteTable operation to delete a custom route table of an Enterprise Edition transit router.
//
// Description:
//
// - You cannot delete the default route table of an Enterprise Edition transit router.
//
// - **DeleteTransitRouterRouteTable*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**. The custom route table is not deleted immediately because the system deletes the route table in the background. You can call the **ListTransitRouterRouteTables*	- operation to query the status of the custom route table.
//
//   - If a custom route table is in the Deleting state, the route table is being deleted. In this state, you can only query the route table. You cannot perform other operations.
//
//   - If the specified custom route table cannot be found, the route table has been deleted.
//
// @param request - DeleteTransitRouterRouteTableRequest
//
// @return DeleteTransitRouterRouteTableResponse
func (client *Client) DeleteTransitRouterRouteTable(request *DeleteTransitRouterRouteTableRequest) (_result *DeleteTransitRouterRouteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterRouteTableResponse{}
	_body, _err := client.DeleteTransitRouterRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virtual border router (VBR) connection for an Enterprise Edition transit router.
//
// Description:
//
// The **DeleteTransitRouterVbrAttachment*	- operation is asynchronous. The system returns a **RequestId**, while running the deletion task in the background. You can call the **ListTransitRouterVbrAttachments*	- operation to query the status of the VBR connection.
//
// - When the VBR connection is in the **Detaching*	- state, the VBR is being deleted. You can only query the VBR connection but cannot perform other operations.
//
// - If a VBR connection cannot be found, the VBR connection is deleted.
//
// Before you call the DeleteTransitRouterVbrAttachment operation, make sure that all request parameters are valid. If a parameter is invalid, the system returns a request ID and does not delete the VBR connection.
//
// @param request - DeleteTransitRouterVbrAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVbrAttachmentResponse
func (client *Client) DeleteTransitRouterVbrAttachmentWithOptions(request *DeleteTransitRouterVbrAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVbrAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virtual border router (VBR) connection for an Enterprise Edition transit router.
//
// Description:
//
// The **DeleteTransitRouterVbrAttachment*	- operation is asynchronous. The system returns a **RequestId**, while running the deletion task in the background. You can call the **ListTransitRouterVbrAttachments*	- operation to query the status of the VBR connection.
//
// - When the VBR connection is in the **Detaching*	- state, the VBR is being deleted. You can only query the VBR connection but cannot perform other operations.
//
// - If a VBR connection cannot be found, the VBR connection is deleted.
//
// Before you call the DeleteTransitRouterVbrAttachment operation, make sure that all request parameters are valid. If a parameter is invalid, the system returns a request ID and does not delete the VBR connection.
//
// @param request - DeleteTransitRouterVbrAttachmentRequest
//
// @return DeleteTransitRouterVbrAttachmentResponse
func (client *Client) DeleteTransitRouterVbrAttachment(request *DeleteTransitRouterVbrAttachmentRequest) (_result *DeleteTransitRouterVbrAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterVbrAttachmentResponse{}
	_body, _err := client.DeleteTransitRouterVbrAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - If a VPC connection is in the **Detaching*	- state, the VPC connection is being deleted. You can query the VPC connection but cannot perform other operations.
//
// - If a VPC connection cannot be found, it is deleted.
//
// ## Prerequisites
//
// Before you delete a VPC connection, make sure that the following requirements are met:
//
// - No associated forwarding correlation is established between the VPC connection and the route tables of the Enterprise Edition transit router. For more information about how to delete an associated forwarding correlation, see [DissociateTransitRouterAttachmentFromRouteTable](https://help.aliyun.com/document_detail/260944.html).
//
// - No route learning correlation is established between the VPC connection and the route tables of the Enterprise Edition transit router. For more information about how to delete a route learning correlation, see [DisableTransitRouterRouteTablePropagation](https://help.aliyun.com/document_detail/260945.html).
//
// - The route table of the VPC does not contain routes that point to the VPC connection. For more information about how to delete routes from a VPC route table, see [DeleteRouteEntry](https://help.aliyun.com/document_detail/36013.html).
//
// - The route tables of the Enterprise Edition transit router do not contain a custom route entry whose next hop is the network instance connection. For more information about how to delete custom routes from the route tables of an Enterprise Edition transit router, see [DeleteTransitRouterRouteEntry](https://help.aliyun.com/document_detail/261240.html).
//
// - The route tables of the Enterprise Edition transit router do not contain a route that is generated from a prefix list and the next hop is the VPC connection. You can delete such routes by disassociating the route table from the prefix list. For more information, see [DeleteTransitRouterPrefixListAssociation](https://help.aliyun.com/document_detail/445486.html).
//
// @param request - DeleteTransitRouterVpcAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVpcAttachmentResponse
func (client *Client) DeleteTransitRouterVpcAttachmentWithOptions(request *DeleteTransitRouterVpcAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVpcAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - If a VPC connection is in the **Detaching*	- state, the VPC connection is being deleted. You can query the VPC connection but cannot perform other operations.
//
// - If a VPC connection cannot be found, it is deleted.
//
// ## Prerequisites
//
// Before you delete a VPC connection, make sure that the following requirements are met:
//
// - No associated forwarding correlation is established between the VPC connection and the route tables of the Enterprise Edition transit router. For more information about how to delete an associated forwarding correlation, see [DissociateTransitRouterAttachmentFromRouteTable](https://help.aliyun.com/document_detail/260944.html).
//
// - No route learning correlation is established between the VPC connection and the route tables of the Enterprise Edition transit router. For more information about how to delete a route learning correlation, see [DisableTransitRouterRouteTablePropagation](https://help.aliyun.com/document_detail/260945.html).
//
// - The route table of the VPC does not contain routes that point to the VPC connection. For more information about how to delete routes from a VPC route table, see [DeleteRouteEntry](https://help.aliyun.com/document_detail/36013.html).
//
// - The route tables of the Enterprise Edition transit router do not contain a custom route entry whose next hop is the network instance connection. For more information about how to delete custom routes from the route tables of an Enterprise Edition transit router, see [DeleteTransitRouterRouteEntry](https://help.aliyun.com/document_detail/261240.html).
//
// - The route tables of the Enterprise Edition transit router do not contain a route that is generated from a prefix list and the next hop is the VPC connection. You can delete such routes by disassociating the route table from the prefix list. For more information, see [DeleteTransitRouterPrefixListAssociation](https://help.aliyun.com/document_detail/445486.html).
//
// @param request - DeleteTransitRouterVpcAttachmentRequest
//
// @return DeleteTransitRouterVpcAttachmentResponse
func (client *Client) DeleteTransitRouterVpcAttachment(request *DeleteTransitRouterVpcAttachmentRequest) (_result *DeleteTransitRouterVpcAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterVpcAttachmentResponse{}
	_body, _err := client.DeleteTransitRouterVpcAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a VPN connection.
//
// Description:
//
// When you call the **DeleteTransitRouterVpnAttachment*	- operation, ensure that the parameter values are valid. If you specify invalid parameters, the system returns a **RequestId*	- but does not delete the VPN connection.
//
// @param request - DeleteTransitRouterVpnAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTransitRouterVpnAttachmentResponse
func (client *Client) DeleteTransitRouterVpnAttachmentWithOptions(request *DeleteTransitRouterVpnAttachmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteTransitRouterVpnAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VPN connection.
//
// Description:
//
// When you call the **DeleteTransitRouterVpnAttachment*	- operation, ensure that the parameter values are valid. If you specify invalid parameters, the system returns a **RequestId*	- but does not delete the VPN connection.
//
// @param request - DeleteTransitRouterVpnAttachmentRequest
//
// @return DeleteTransitRouterVpnAttachmentResponse
func (client *Client) DeleteTransitRouterVpnAttachment(request *DeleteTransitRouterVpnAttachmentRequest) (_result *DeleteTransitRouterVpnAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTransitRouterVpnAttachmentResponse{}
	_body, _err := client.DeleteTransitRouterVpnAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If a multicast member no longer needs to receive multicast traffic, you can call the DeregisterTransitRouterMulticastGroupMembers operation to delete the member from the multicast group.
//
// Description:
//
// `DeregisterTransitRouterMulticastGroupMembers` is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the multicast member is not deleted immediately. The deletion task runs in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of the multicast member.
//
// - If a multicast member is in the **Deregistering*	- state, it is being deleted. In this state, you can only query the member. You cannot perform other operations.
//
// - If the `ListTransitRouterMulticastGroups` operation does not find the multicast member in the multicast domain, the member is deleted.
//
// Ensure that you enter correct parameter values when you call the DeregisterTransitRouterMulticastGroupMembers operation. If you enter an incorrect parameter, the operation still returns a RequestId but does not delete the multicast member.
//
// @param request - DeregisterTransitRouterMulticastGroupMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterTransitRouterMulticastGroupMembersResponse
func (client *Client) DeregisterTransitRouterMulticastGroupMembersWithOptions(request *DeregisterTransitRouterMulticastGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *DeregisterTransitRouterMulticastGroupMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If a multicast member no longer needs to receive multicast traffic, you can call the DeregisterTransitRouterMulticastGroupMembers operation to delete the member from the multicast group.
//
// Description:
//
// `DeregisterTransitRouterMulticastGroupMembers` is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the multicast member is not deleted immediately. The deletion task runs in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of the multicast member.
//
// - If a multicast member is in the **Deregistering*	- state, it is being deleted. In this state, you can only query the member. You cannot perform other operations.
//
// - If the `ListTransitRouterMulticastGroups` operation does not find the multicast member in the multicast domain, the member is deleted.
//
// Ensure that you enter correct parameter values when you call the DeregisterTransitRouterMulticastGroupMembers operation. If you enter an incorrect parameter, the operation still returns a RequestId but does not delete the multicast member.
//
// @param request - DeregisterTransitRouterMulticastGroupMembersRequest
//
// @return DeregisterTransitRouterMulticastGroupMembersResponse
func (client *Client) DeregisterTransitRouterMulticastGroupMembers(request *DeregisterTransitRouterMulticastGroupMembersRequest) (_result *DeregisterTransitRouterMulticastGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeregisterTransitRouterMulticastGroupMembersResponse{}
	_body, _err := client.DeregisterTransitRouterMulticastGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If a multicast source no longer needs to send multicast traffic, you can call the DeregisterTransitRouterMulticastGroupSources operation to remove the multicast source from the multicast group.
//
// Description:
//
// `DeregisterTransitRouterMulticastGroupSources` is an asynchronous operation. After you send a request, the system returns a `RequestId`, but the multicast source is not deleted immediately. The system deletes the multicast source in the background. You can call `ListTransitRouterMulticastGroups` to query the status of the multicast source.
//
// - If a multicast source is in the `Deregistering` state, it is being deleted. In this state, you can only query the multicast source. You cannot perform other operations.
//
// - If you cannot find the multicast source in the multicast domain when you call `ListTransitRouterMulticastGroups`, the multicast source has been deleted.
//
// Ensure that you specify correct parameter values when you call the DeregisterTransitRouterMulticastGroupSources operation. If you specify an incorrect parameter, the operation returns a RequestId but does not delete the multicast source.
//
// @param request - DeregisterTransitRouterMulticastGroupSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterTransitRouterMulticastGroupSourcesResponse
func (client *Client) DeregisterTransitRouterMulticastGroupSourcesWithOptions(request *DeregisterTransitRouterMulticastGroupSourcesRequest, runtime *dara.RuntimeOptions) (_result *DeregisterTransitRouterMulticastGroupSourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If a multicast source no longer needs to send multicast traffic, you can call the DeregisterTransitRouterMulticastGroupSources operation to remove the multicast source from the multicast group.
//
// Description:
//
// `DeregisterTransitRouterMulticastGroupSources` is an asynchronous operation. After you send a request, the system returns a `RequestId`, but the multicast source is not deleted immediately. The system deletes the multicast source in the background. You can call `ListTransitRouterMulticastGroups` to query the status of the multicast source.
//
// - If a multicast source is in the `Deregistering` state, it is being deleted. In this state, you can only query the multicast source. You cannot perform other operations.
//
// - If you cannot find the multicast source in the multicast domain when you call `ListTransitRouterMulticastGroups`, the multicast source has been deleted.
//
// Ensure that you specify correct parameter values when you call the DeregisterTransitRouterMulticastGroupSources operation. If you specify an incorrect parameter, the operation returns a RequestId but does not delete the multicast source.
//
// @param request - DeregisterTransitRouterMulticastGroupSourcesRequest
//
// @return DeregisterTransitRouterMulticastGroupSourcesResponse
func (client *Client) DeregisterTransitRouterMulticastGroupSources(request *DeregisterTransitRouterMulticastGroupSourcesRequest) (_result *DeregisterTransitRouterMulticastGroupSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeregisterTransitRouterMulticastGroupSourcesResponse{}
	_body, _err := client.DeregisterTransitRouterMulticastGroupSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCenAttachedChildInstanceAttributeWithOptions(request *DescribeCenAttachedChildInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenAttachedChildInstanceAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCenAttachedChildInstanceAttributeResponse
func (client *Client) DescribeCenAttachedChildInstanceAttribute(request *DescribeCenAttachedChildInstanceAttributeRequest) (_result *DescribeCenAttachedChildInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenAttachedChildInstanceAttributeResponse{}
	_body, _err := client.DescribeCenAttachedChildInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - You can query all the network instances that are attached to a CEN instance by setting the `CenId` parameter.
//
// - You can query the network instances that are attached to a CEN instance in a specified region by setting the `CenId` and `ChildInstanceRegionId` parameters.
//
// - You can query a specified type of network instances that are attached to a CEN instance by setting the `CenId` and `ChildInstanceType` parameters.
//
// @param request - DescribeCenAttachedChildInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenAttachedChildInstancesResponse
func (client *Client) DescribeCenAttachedChildInstancesWithOptions(request *DescribeCenAttachedChildInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenAttachedChildInstancesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - You can query all the network instances that are attached to a CEN instance by setting the `CenId` parameter.
//
// - You can query the network instances that are attached to a CEN instance in a specified region by setting the `CenId` and `ChildInstanceRegionId` parameters.
//
// - You can query a specified type of network instances that are attached to a CEN instance by setting the `CenId` and `ChildInstanceType` parameters.
//
// @param request - DescribeCenAttachedChildInstancesRequest
//
// @return DescribeCenAttachedChildInstancesResponse
func (client *Client) DescribeCenAttachedChildInstances(request *DescribeCenAttachedChildInstancesRequest) (_result *DescribeCenAttachedChildInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenAttachedChildInstancesResponse{}
	_body, _err := client.DescribeCenAttachedChildInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCenBandwidthPackagesWithOptions(request *DescribeCenBandwidthPackagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenBandwidthPackagesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCenBandwidthPackagesResponse
func (client *Client) DescribeCenBandwidthPackages(request *DescribeCenBandwidthPackagesRequest) (_result *DescribeCenBandwidthPackagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenBandwidthPackagesResponse{}
	_body, _err := client.DescribeCenBandwidthPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeCenChildInstanceRouteEntriesWithOptions(request *DescribeCenChildInstanceRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenChildInstanceRouteEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeCenChildInstanceRouteEntriesResponse
func (client *Client) DescribeCenChildInstanceRouteEntries(request *DescribeCenChildInstanceRouteEntriesRequest) (_result *DescribeCenChildInstanceRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenChildInstanceRouteEntriesResponse{}
	_body, _err := client.DescribeCenChildInstanceRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the remaining bandwidth of a specified bandwidth plan.
//
// @param request - DescribeCenGeographicSpanRemainingBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenGeographicSpanRemainingBandwidthResponse
func (client *Client) DescribeCenGeographicSpanRemainingBandwidthWithOptions(request *DescribeCenGeographicSpanRemainingBandwidthRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenGeographicSpanRemainingBandwidthResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the remaining bandwidth of a specified bandwidth plan.
//
// @param request - DescribeCenGeographicSpanRemainingBandwidthRequest
//
// @return DescribeCenGeographicSpanRemainingBandwidthResponse
func (client *Client) DescribeCenGeographicSpanRemainingBandwidth(request *DescribeCenGeographicSpanRemainingBandwidthRequest) (_result *DescribeCenGeographicSpanRemainingBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenGeographicSpanRemainingBandwidthResponse{}
	_body, _err := client.DescribeCenGeographicSpanRemainingBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeCenGeographicSpans operation to query the connected areas supported by Cloud Enterprise Network (CEN).
//
// @param request - DescribeCenGeographicSpansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenGeographicSpansResponse
func (client *Client) DescribeCenGeographicSpansWithOptions(request *DescribeCenGeographicSpansRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenGeographicSpansResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeCenGeographicSpans operation to query the connected areas supported by Cloud Enterprise Network (CEN).
//
// @param request - DescribeCenGeographicSpansRequest
//
// @return DescribeCenGeographicSpansResponse
func (client *Client) DescribeCenGeographicSpans(request *DescribeCenGeographicSpansRequest) (_result *DescribeCenGeographicSpansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenGeographicSpansResponse{}
	_body, _err := client.DescribeCenGeographicSpansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DescribeCenInterRegionBandwidthLimits operation to query the bandwidth limits for inter-region communication.
//
// @param request - DescribeCenInterRegionBandwidthLimitsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenInterRegionBandwidthLimitsResponse
func (client *Client) DescribeCenInterRegionBandwidthLimitsWithOptions(request *DescribeCenInterRegionBandwidthLimitsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenInterRegionBandwidthLimitsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DescribeCenInterRegionBandwidthLimits operation to query the bandwidth limits for inter-region communication.
//
// @param request - DescribeCenInterRegionBandwidthLimitsRequest
//
// @return DescribeCenInterRegionBandwidthLimitsResponse
func (client *Client) DescribeCenInterRegionBandwidthLimits(request *DescribeCenInterRegionBandwidthLimitsRequest) (_result *DescribeCenInterRegionBandwidthLimitsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenInterRegionBandwidthLimitsResponse{}
	_body, _err := client.DescribeCenInterRegionBandwidthLimitsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the PrivateZone service configurations of a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenPrivateZoneRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenPrivateZoneRoutesResponse
func (client *Client) DescribeCenPrivateZoneRoutesWithOptions(request *DescribeCenPrivateZoneRoutesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenPrivateZoneRoutesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the PrivateZone service configurations of a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenPrivateZoneRoutesRequest
//
// @return DescribeCenPrivateZoneRoutesResponse
func (client *Client) DescribeCenPrivateZoneRoutes(request *DescribeCenPrivateZoneRoutesRequest) (_result *DescribeCenPrivateZoneRoutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenPrivateZoneRoutesResponse{}
	_body, _err := client.DescribeCenPrivateZoneRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries route entries in a specified region of a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenRegionDomainRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenRegionDomainRouteEntriesResponse
func (client *Client) DescribeCenRegionDomainRouteEntriesWithOptions(request *DescribeCenRegionDomainRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenRegionDomainRouteEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries route entries in a specified region of a Cloud Enterprise Network (CEN) instance.
//
// @param request - DescribeCenRegionDomainRouteEntriesRequest
//
// @return DescribeCenRegionDomainRouteEntriesResponse
func (client *Client) DescribeCenRegionDomainRouteEntries(request *DescribeCenRegionDomainRouteEntriesRequest) (_result *DescribeCenRegionDomainRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenRegionDomainRouteEntriesResponse{}
	_body, _err := client.DescribeCenRegionDomainRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeCenRouteMaps operation to query the configurations of routing policies.
//
// @param request - DescribeCenRouteMapsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenRouteMapsResponse
func (client *Client) DescribeCenRouteMapsWithOptions(request *DescribeCenRouteMapsRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenRouteMapsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeCenRouteMaps operation to query the configurations of routing policies.
//
// @param request - DescribeCenRouteMapsRequest
//
// @return DescribeCenRouteMapsResponse
func (client *Client) DescribeCenRouteMaps(request *DescribeCenRouteMapsRequest) (_result *DescribeCenRouteMapsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenRouteMapsResponse{}
	_body, _err := client.DescribeCenRouteMapsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the health check information for a virtual border router (VBR) in a specified region.
//
// @param request - DescribeCenVbrHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCenVbrHealthCheckResponse
func (client *Client) DescribeCenVbrHealthCheckWithOptions(request *DescribeCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *DescribeCenVbrHealthCheckResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the health check information for a virtual border router (VBR) in a specified region.
//
// @param request - DescribeCenVbrHealthCheckRequest
//
// @return DescribeCenVbrHealthCheckResponse
func (client *Client) DescribeCenVbrHealthCheck(request *DescribeCenVbrHealthCheckRequest) (_result *DescribeCenVbrHealthCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCenVbrHealthCheckResponse{}
	_body, _err := client.DescribeCenVbrHealthCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeCens operation to query information about Cloud Enterprise Network (CEN) instances that belong to your Alibaba Cloud account. This information includes the status of the instances, whether IPv6 is enabled, and a list of attached bandwidth plans.
//
// @param request - DescribeCensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCensResponse
func (client *Client) DescribeCensWithOptions(request *DescribeCensRequest, runtime *dara.RuntimeOptions) (_result *DescribeCensResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeCens operation to query information about Cloud Enterprise Network (CEN) instances that belong to your Alibaba Cloud account. This information includes the status of the instances, whether IPv6 is enabled, and a list of attached bandwidth plans.
//
// @param request - DescribeCensRequest
//
// @return DescribeCensResponse
func (client *Client) DescribeCens(request *DescribeCensRequest) (_result *DescribeCensResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCensResponse{}
	_body, _err := client.DescribeCensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the DescribeChildInstanceRegions operation to query the regions where you can attach network instances to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// The regions that CEN supports vary based on the network instance type. You can specify the `ProductType` parameter to query the regions that CEN supports for a specific type of network instance. If you do not specify the `ProductType` parameter, the system queries the regions supported for all network instance types by default.
//
// @param request - DescribeChildInstanceRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChildInstanceRegionsResponse
func (client *Client) DescribeChildInstanceRegionsWithOptions(request *DescribeChildInstanceRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeChildInstanceRegionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the DescribeChildInstanceRegions operation to query the regions where you can attach network instances to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// The regions that CEN supports vary based on the network instance type. You can specify the `ProductType` parameter to query the regions that CEN supports for a specific type of network instance. If you do not specify the `ProductType` parameter, the system queries the regions supported for all network instance types by default.
//
// @param request - DescribeChildInstanceRegionsRequest
//
// @return DescribeChildInstanceRegionsResponse
func (client *Client) DescribeChildInstanceRegions(request *DescribeChildInstanceRegionsRequest) (_result *DescribeChildInstanceRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeChildInstanceRegionsResponse{}
	_body, _err := client.DescribeChildInstanceRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeFlowlogsWithOptions(request *DescribeFlowlogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowlogsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeFlowlogsResponse
func (client *Client) DescribeFlowlogs(request *DescribeFlowlogsRequest) (_result *DescribeFlowlogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowlogsResponse{}
	_body, _err := client.DescribeFlowlogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeGeographicRegionMembershipWithOptions(request *DescribeGeographicRegionMembershipRequest, runtime *dara.RuntimeOptions) (_result *DescribeGeographicRegionMembershipResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeGeographicRegionMembershipResponse
func (client *Client) DescribeGeographicRegionMembership(request *DescribeGeographicRegionMembershipRequest) (_result *DescribeGeographicRegionMembershipResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGeographicRegionMembershipResponse{}
	_body, _err := client.DescribeGeographicRegionMembershipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cross-account network instances that have been authorized for a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// Calling the DescribeGrantRulesToCen operation with invalid parameters returns a **RequestId*	- but provides no information about the cross-account network instances that the CEN instance is permitted to access.
//
// @param request - DescribeGrantRulesToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantRulesToCenResponse
func (client *Client) DescribeGrantRulesToCenWithOptions(request *DescribeGrantRulesToCenRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantRulesToCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cross-account network instances that have been authorized for a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// Calling the DescribeGrantRulesToCen operation with invalid parameters returns a **RequestId*	- but provides no information about the cross-account network instances that the CEN instance is permitted to access.
//
// @param request - DescribeGrantRulesToCenRequest
//
// @return DescribeGrantRulesToCenResponse
func (client *Client) DescribeGrantRulesToCen(request *DescribeGrantRulesToCenRequest) (_result *DescribeGrantRulesToCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGrantRulesToCenResponse{}
	_body, _err := client.DescribeGrantRulesToCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions that a network instance has on a Cloud Enterprise Network (CEN) instance owned by another Alibaba Cloud account. This operation returns details such as the main account that owns the CEN instance and the payer for the network instance.
//
// @param request - DescribeGrantRulesToResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantRulesToResourceResponse
func (client *Client) DescribeGrantRulesToResourceWithOptions(request *DescribeGrantRulesToResourceRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantRulesToResourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions that a network instance has on a Cloud Enterprise Network (CEN) instance owned by another Alibaba Cloud account. This operation returns details such as the main account that owns the CEN instance and the payer for the network instance.
//
// @param request - DescribeGrantRulesToResourceRequest
//
// @return DescribeGrantRulesToResourceResponse
func (client *Client) DescribeGrantRulesToResource(request *DescribeGrantRulesToResourceRequest) (_result *DescribeGrantRulesToResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeGrantRulesToResourceResponse{}
	_body, _err := client.DescribeGrantRulesToResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the routes of virtual private clouds (VPCs) and virtual border routers (VBRs) are advertised to the Cloud Enterprise Network (CEN) instance to which the VCPs and VBRs are attached, the instance type of the next hop of each route, and whether advertised routes can be withdrawn.
//
// @param request - DescribePublishedRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePublishedRouteEntriesResponse
func (client *Client) DescribePublishedRouteEntriesWithOptions(request *DescribePublishedRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribePublishedRouteEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the routes of virtual private clouds (VPCs) and virtual border routers (VBRs) are advertised to the Cloud Enterprise Network (CEN) instance to which the VCPs and VBRs are attached, the instance type of the next hop of each route, and whether advertised routes can be withdrawn.
//
// @param request - DescribePublishedRouteEntriesRequest
//
// @return DescribePublishedRouteEntriesResponse
func (client *Client) DescribePublishedRouteEntries(request *DescribePublishedRouteEntriesRequest) (_result *DescribePublishedRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePublishedRouteEntriesResponse{}
	_body, _err := client.DescribePublishedRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The DescribeRouteConflict operation queries conflicting routes in a network instance.
//
// @param request - DescribeRouteConflictRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteConflictResponse
func (client *Client) DescribeRouteConflictWithOptions(request *DescribeRouteConflictRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteConflictResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The DescribeRouteConflict operation queries conflicting routes in a network instance.
//
// @param request - DescribeRouteConflictRequest
//
// @return DescribeRouteConflictResponse
func (client *Client) DescribeRouteConflict(request *DescribeRouteConflictRequest) (_result *DescribeRouteConflictResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRouteConflictResponse{}
	_body, _err := client.DescribeRouteConflictWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of Alibaba Cloud services that are configured in a Basic Edition transit router.
//
// @param request - DescribeRouteServicesInCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteServicesInCenResponse
func (client *Client) DescribeRouteServicesInCenWithOptions(request *DescribeRouteServicesInCenRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteServicesInCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of Alibaba Cloud services that are configured in a Basic Edition transit router.
//
// @param request - DescribeRouteServicesInCenRequest
//
// @return DescribeRouteServicesInCenResponse
func (client *Client) DescribeRouteServicesInCen(request *DescribeRouteServicesInCenRequest) (_result *DescribeRouteServicesInCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRouteServicesInCenResponse{}
	_body, _err := client.DescribeRouteServicesInCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the aggregate routes in the route table of an Enterprise Edition transit router.
//
// Description:
//
// You can specify the **TransitRouteTableId*	- and **TransitRouteTableAggregationCidr*	- parameters to query a specific aggregate route. If you specify only the **TransitRouteTableId*	- parameter, all aggregate routes in the route table of the Enterprise Edition transit router are queried by default.
//
// @param request - DescribeTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTransitRouteTableAggregationResponse
func (client *Client) DescribeTransitRouteTableAggregationWithOptions(request *DescribeTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransitRouteTableAggregationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// You can specify the **TransitRouteTableId*	- and **TransitRouteTableAggregationCidr*	- parameters to query a specific aggregate route. If you specify only the **TransitRouteTableId*	- parameter, all aggregate routes in the route table of the Enterprise Edition transit router are queried by default.
//
// @param request - DescribeTransitRouteTableAggregationRequest
//
// @return DescribeTransitRouteTableAggregationResponse
func (client *Client) DescribeTransitRouteTableAggregation(request *DescribeTransitRouteTableAggregationRequest) (_result *DescribeTransitRouteTableAggregationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTransitRouteTableAggregationResponse{}
	_body, _err := client.DescribeTransitRouteTableAggregationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeTransitRouteTableAggregationDetailWithOptions(request *DescribeTransitRouteTableAggregationDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeTransitRouteTableAggregationDetailResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeTransitRouteTableAggregationDetailResponse
func (client *Client) DescribeTransitRouteTableAggregationDetail(request *DescribeTransitRouteTableAggregationDetailRequest) (_result *DescribeTransitRouteTableAggregationDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeTransitRouteTableAggregationDetailResponse{}
	_body, _err := client.DescribeTransitRouteTableAggregationDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches a network instance from a Cloud Enterprise Network (CEN) transit router.
//
// Description:
//
// The transit router must be a Basic Edition transit router.
//
// @param request - DetachCenChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachCenChildInstanceResponse
func (client *Client) DetachCenChildInstanceWithOptions(request *DetachCenChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DetachCenChildInstanceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a network instance from a Cloud Enterprise Network (CEN) transit router.
//
// Description:
//
// The transit router must be a Basic Edition transit router.
//
// @param request - DetachCenChildInstanceRequest
//
// @return DetachCenChildInstanceResponse
func (client *Client) DetachCenChildInstance(request *DetachCenChildInstanceRequest) (_result *DetachCenChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachCenChildInstanceResponse{}
	_body, _err := client.DetachCenChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DisableCenVbrHealthCheckWithOptions(request *DisableCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *DisableCenVbrHealthCheckResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DisableCenVbrHealthCheckResponse
func (client *Client) DisableCenVbrHealthCheck(request *DisableCenVbrHealthCheckRequest) (_result *DisableCenVbrHealthCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableCenVbrHealthCheckResponse{}
	_body, _err := client.DisableCenVbrHealthCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - If a route learning correlation is in the **Disabling*	- state, the route learning correlation is being deleted. You can query the route learning correlation but cannot perform other operations.
//
// - If a route learning correlation cannot be found, the route learning correlation is deleted.
//
// @param request - DisableTransitRouterRouteTablePropagationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableTransitRouterRouteTablePropagationResponse
func (client *Client) DisableTransitRouterRouteTablePropagationWithOptions(request *DisableTransitRouterRouteTablePropagationRequest, runtime *dara.RuntimeOptions) (_result *DisableTransitRouterRouteTablePropagationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - If a route learning correlation is in the **Disabling*	- state, the route learning correlation is being deleted. You can query the route learning correlation but cannot perform other operations.
//
// - If a route learning correlation cannot be found, the route learning correlation is deleted.
//
// @param request - DisableTransitRouterRouteTablePropagationRequest
//
// @return DisableTransitRouterRouteTablePropagationResponse
func (client *Client) DisableTransitRouterRouteTablePropagation(request *DisableTransitRouterRouteTablePropagationRequest) (_result *DisableTransitRouterRouteTablePropagationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableTransitRouterRouteTablePropagationResponse{}
	_body, _err := client.DisableTransitRouterRouteTablePropagationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the DisassociateTransitRouterMulticastDomain operation to dissociate a vSwitch from a multicast domain.
//
// Description:
//
// - Before you dissociate a vSwitch from a multicast domain, make sure that no multicast source or member exists on the vSwitch. For more information about how to delete a multicast source and a multicast member, see [DeregisterTransitRouterMulticastGroupSources](https://help.aliyun.com/document_detail/468416.html) and [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/468409.html).
//
// - If you provide invalid parameters, the system returns a request ID but does not dissociate the vSwitch from the multicast domain.
//
// - **DisassociateTransitRouterMulticastDomain*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the vSwitch is not immediately dissociated from the multicast domain. The system runs the dissociation task in the background. You can call the **ListTransitRouterMulticastDomainAssociations*	- operation to query the association status of the vSwitch and the multicast domain.
//
//   - If the association status is **Dissociating**, the vSwitch is being dissociated from the multicast domain. In this state, you can only query the vSwitch. You cannot perform other operations.
//
//   - If the vSwitch cannot be found in the multicast domain, the vSwitch has been successfully dissociated from the multicast domain.
//
// @param request - DisassociateTransitRouterMulticastDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateTransitRouterMulticastDomainResponse
func (client *Client) DisassociateTransitRouterMulticastDomainWithOptions(request *DisassociateTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *DisassociateTransitRouterMulticastDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the DisassociateTransitRouterMulticastDomain operation to dissociate a vSwitch from a multicast domain.
//
// Description:
//
// - Before you dissociate a vSwitch from a multicast domain, make sure that no multicast source or member exists on the vSwitch. For more information about how to delete a multicast source and a multicast member, see [DeregisterTransitRouterMulticastGroupSources](https://help.aliyun.com/document_detail/468416.html) and [DeregisterTransitRouterMulticastGroupMembers](https://help.aliyun.com/document_detail/468409.html).
//
// - If you provide invalid parameters, the system returns a request ID but does not dissociate the vSwitch from the multicast domain.
//
// - **DisassociateTransitRouterMulticastDomain*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the vSwitch is not immediately dissociated from the multicast domain. The system runs the dissociation task in the background. You can call the **ListTransitRouterMulticastDomainAssociations*	- operation to query the association status of the vSwitch and the multicast domain.
//
//   - If the association status is **Dissociating**, the vSwitch is being dissociated from the multicast domain. In this state, you can only query the vSwitch. You cannot perform other operations.
//
//   - If the vSwitch cannot be found in the multicast domain, the vSwitch has been successfully dissociated from the multicast domain.
//
// @param request - DisassociateTransitRouterMulticastDomainRequest
//
// @return DisassociateTransitRouterMulticastDomainResponse
func (client *Client) DisassociateTransitRouterMulticastDomain(request *DisassociateTransitRouterMulticastDomainRequest) (_result *DisassociateTransitRouterMulticastDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisassociateTransitRouterMulticastDomainResponse{}
	_body, _err := client.DisassociateTransitRouterMulticastDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dissociates a network instance connection from a route table.
//
// Description:
//
// *DissociateTransitRouterAttachmentFromRouteTable*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the operation runs in the background. The network instance connection is not immediately dissociated from the route table. To query the status of the association, call the **ListTransitRouterRouteTableAssociations*	- operation.
//
// - If the association status is **Dissociating**, the network instance connection is being dissociated from the route table. In this state, you can only query the association. You cannot perform other operations.
//
// - If the **ListTransitRouterRouteTableAssociations*	- operation does not return information about the association, the network instance connection has been dissociated.
//
// @param request - DissociateTransitRouterAttachmentFromRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateTransitRouterAttachmentFromRouteTableResponse
func (client *Client) DissociateTransitRouterAttachmentFromRouteTableWithOptions(request *DissociateTransitRouterAttachmentFromRouteTableRequest, runtime *dara.RuntimeOptions) (_result *DissociateTransitRouterAttachmentFromRouteTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// *DissociateTransitRouterAttachmentFromRouteTable*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the operation runs in the background. The network instance connection is not immediately dissociated from the route table. To query the status of the association, call the **ListTransitRouterRouteTableAssociations*	- operation.
//
// - If the association status is **Dissociating**, the network instance connection is being dissociated from the route table. In this state, you can only query the association. You cannot perform other operations.
//
// - If the **ListTransitRouterRouteTableAssociations*	- operation does not return information about the association, the network instance connection has been dissociated.
//
// @param request - DissociateTransitRouterAttachmentFromRouteTableRequest
//
// @return DissociateTransitRouterAttachmentFromRouteTableResponse
func (client *Client) DissociateTransitRouterAttachmentFromRouteTable(request *DissociateTransitRouterAttachmentFromRouteTableRequest) (_result *DissociateTransitRouterAttachmentFromRouteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DissociateTransitRouterAttachmentFromRouteTableResponse{}
	_body, _err := client.DissociateTransitRouterAttachmentFromRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the health check feature for a virtual border router (VBR) or modifies the health check configuration of a VBR. Health checks help you promptly detect faulty Express Connect circuits.
//
// Description:
//
// You can configure a health check for a VBR instance to monitor the connection status of the Express Connect circuit between your data center and Alibaba Cloud. This helps you promptly identify issues.
//
// Before you use the health check feature, note the following:
//
// - If your VBR instance uses static routing, create a static route in the data center connected to the VBR instance after you configure the health check. Set the destination CIDR block of the static route to the source IP address of the health check, the subnet mask to 32 bits, and the next hop to the Alibaba Cloud-side IP address of the VBR instance.
//
// - If your VBR instance uses the Border Gateway Protocol (BGP) dynamic routing protocol, you do not need to create a route in the data center.
//
// - **EnableCenVbrHealthCheck*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the health check is not yet created or modified because the task is still running in the background. You can call the **DescribeCenVbrHealthCheck*	- operation to query the health check configuration. If the configuration is returned, this indicates that the health check is created or modified.
//
// @param request - EnableCenVbrHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCenVbrHealthCheckResponse
func (client *Client) EnableCenVbrHealthCheckWithOptions(request *EnableCenVbrHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *EnableCenVbrHealthCheckResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the health check feature for a virtual border router (VBR) or modifies the health check configuration of a VBR. Health checks help you promptly detect faulty Express Connect circuits.
//
// Description:
//
// You can configure a health check for a VBR instance to monitor the connection status of the Express Connect circuit between your data center and Alibaba Cloud. This helps you promptly identify issues.
//
// Before you use the health check feature, note the following:
//
// - If your VBR instance uses static routing, create a static route in the data center connected to the VBR instance after you configure the health check. Set the destination CIDR block of the static route to the source IP address of the health check, the subnet mask to 32 bits, and the next hop to the Alibaba Cloud-side IP address of the VBR instance.
//
// - If your VBR instance uses the Border Gateway Protocol (BGP) dynamic routing protocol, you do not need to create a route in the data center.
//
// - **EnableCenVbrHealthCheck*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the health check is not yet created or modified because the task is still running in the background. You can call the **DescribeCenVbrHealthCheck*	- operation to query the health check configuration. If the configuration is returned, this indicates that the health check is created or modified.
//
// @param request - EnableCenVbrHealthCheckRequest
//
// @return EnableCenVbrHealthCheckResponse
func (client *Client) EnableCenVbrHealthCheck(request *EnableCenVbrHealthCheckRequest) (_result *EnableCenVbrHealthCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableCenVbrHealthCheckResponse{}
	_body, _err := client.EnableCenVbrHealthCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - You can create route learning correlations only on Enterprise Edition transit routers. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
// - **EnableTransitRouterRouteTablePropagation*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterRouteTablePropagations*	- operation to query the route learning status between a network instance connection and a route table.
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
func (client *Client) EnableTransitRouterRouteTablePropagationWithOptions(request *EnableTransitRouterRouteTablePropagationRequest, runtime *dara.RuntimeOptions) (_result *EnableTransitRouterRouteTablePropagationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - You can create route learning correlations only on Enterprise Edition transit routers. For more information about the regions and zones that support Enterprise Edition transit routers, see [What is CEN?](https://help.aliyun.com/document_detail/181681.html)
//
// - **EnableTransitRouterRouteTablePropagation*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTransitRouterRouteTablePropagations*	- operation to query the route learning status between a network instance connection and a route table.
//
//   - **Enabling*	- indicates that a route learning correlation is being created between the network instance connection and route table. You can query the route learning correlation but cannot perform other operations.
//
//   - **Active*	- indicates that the route learning correlation is created between the network instance connection and route table.
//
// @param request - EnableTransitRouterRouteTablePropagationRequest
//
// @return EnableTransitRouterRouteTablePropagationResponse
func (client *Client) EnableTransitRouterRouteTablePropagation(request *EnableTransitRouterRouteTablePropagationRequest) (_result *EnableTransitRouterRouteTablePropagationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableTransitRouterRouteTablePropagationResponse{}
	_body, _err := client.EnableTransitRouterRouteTablePropagationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants a transit router permissions on network instances that belong to another Alibaba Cloud account. To connect a transit router of Account B to a network instance of Account A, you must use Account A to grant permissions to the transit router of Account B.
//
// Description:
//
// - The `GrantInstanceToTransitRouter` operation can be used to grant transit routers permissions on network instances that belong to other Alibaba Cloud accounts, including virtual private clouds (VPCs), virtual border routers (VBRs), IPsec-VPN connections, and Express Connect Router (ECRs).
//
//	To grant transit routers permissions on Cloud Connect Network (CCN) instances, call the [GrantInstanceToCbn](https://help.aliyun.com/document_detail/126141.html) operation.
//
// - Before you call `GrantInstanceToTransitRouter`, take note of the billing rules, permission limits, and prerequisites on permission management of transit routers. For more information, see [Acquire permissions to connect to a network instance that belongs to another account](https://help.aliyun.com/document_detail/181553.html).
//
// - Before you grant a transit router permissions on a network instance, make sure that the following requirements are met:
//
//	The account to which the network instance belongs and the account to which the transit router belongs are of the same type.
//
//	The ID of the Alibaba Cloud account to which the transit router belongs is obtained.
//
//	The ID of the Cloud Enterprise Network (CEN) instance to which the Enterprise Edition transit router belongs is obtained.
//
//	Before you grant a transit router permissions on a VBR, contact your account manager to acquire permissions on the VBR.
//
//	Before you grant a transit router permissions on an IPsec-VPN connection, make sure that the IPsec-VPN connection is not associated with a resource.
//
//	If the IPsec-VPN connection is attached to a VPN gateway, the IPsec-VPN connection cannot be attached to transit routers within the same account or different accounts.
//
//	If the IPsec-VPN connection is attached to a transit router, detach the IPsec-VPN connection from the transit router. For more information, see [Delete a network instance connection](https://help.aliyun.com/document_detail/181554.html).
//
// @param request - GrantInstanceToTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantInstanceToTransitRouterResponse
func (client *Client) GrantInstanceToTransitRouterWithOptions(request *GrantInstanceToTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *GrantInstanceToTransitRouterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - The `GrantInstanceToTransitRouter` operation can be used to grant transit routers permissions on network instances that belong to other Alibaba Cloud accounts, including virtual private clouds (VPCs), virtual border routers (VBRs), IPsec-VPN connections, and Express Connect Router (ECRs).
//
//	To grant transit routers permissions on Cloud Connect Network (CCN) instances, call the [GrantInstanceToCbn](https://help.aliyun.com/document_detail/126141.html) operation.
//
// - Before you call `GrantInstanceToTransitRouter`, take note of the billing rules, permission limits, and prerequisites on permission management of transit routers. For more information, see [Acquire permissions to connect to a network instance that belongs to another account](https://help.aliyun.com/document_detail/181553.html).
//
// - Before you grant a transit router permissions on a network instance, make sure that the following requirements are met:
//
//	The account to which the network instance belongs and the account to which the transit router belongs are of the same type.
//
//	The ID of the Alibaba Cloud account to which the transit router belongs is obtained.
//
//	The ID of the Cloud Enterprise Network (CEN) instance to which the Enterprise Edition transit router belongs is obtained.
//
//	Before you grant a transit router permissions on a VBR, contact your account manager to acquire permissions on the VBR.
//
//	Before you grant a transit router permissions on an IPsec-VPN connection, make sure that the IPsec-VPN connection is not associated with a resource.
//
//	If the IPsec-VPN connection is attached to a VPN gateway, the IPsec-VPN connection cannot be attached to transit routers within the same account or different accounts.
//
//	If the IPsec-VPN connection is attached to a transit router, detach the IPsec-VPN connection from the transit router. For more information, see [Delete a network instance connection](https://help.aliyun.com/document_detail/181554.html).
//
// @param request - GrantInstanceToTransitRouterRequest
//
// @return GrantInstanceToTransitRouterResponse
func (client *Client) GrantInstanceToTransitRouter(request *GrantInstanceToTransitRouterRequest) (_result *GrantInstanceToTransitRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantInstanceToTransitRouterResponse{}
	_body, _err := client.GrantInstanceToTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the routes that point to a network instance connection. The routes are retrieved from the route table of a network instance that is attached to an Enterprise Edition transit router.
//
// Description:
//
// Ensure that you specify valid parameter values when you call the ListCenChildInstanceRouteEntriesToAttachment operation. If you specify an invalid parameter, the system returns a request ID but does not return the routes of the network instance that is connected to the Enterprise Edition transit router.
//
// @param request - ListCenChildInstanceRouteEntriesToAttachmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenChildInstanceRouteEntriesToAttachmentResponse
func (client *Client) ListCenChildInstanceRouteEntriesToAttachmentWithOptions(request *ListCenChildInstanceRouteEntriesToAttachmentRequest, runtime *dara.RuntimeOptions) (_result *ListCenChildInstanceRouteEntriesToAttachmentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routes that point to a network instance connection. The routes are retrieved from the route table of a network instance that is attached to an Enterprise Edition transit router.
//
// Description:
//
// Ensure that you specify valid parameter values when you call the ListCenChildInstanceRouteEntriesToAttachment operation. If you specify an invalid parameter, the system returns a request ID but does not return the routes of the network instance that is connected to the Enterprise Edition transit router.
//
// @param request - ListCenChildInstanceRouteEntriesToAttachmentRequest
//
// @return ListCenChildInstanceRouteEntriesToAttachmentResponse
func (client *Client) ListCenChildInstanceRouteEntriesToAttachment(request *ListCenChildInstanceRouteEntriesToAttachmentRequest) (_result *ListCenChildInstanceRouteEntriesToAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCenChildInstanceRouteEntriesToAttachmentResponse{}
	_body, _err := client.ListCenChildInstanceRouteEntriesToAttachmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListCenInterRegionTrafficQosPolicies operation to query Quality of Service (QoS) policies.
//
// Description:
//
// Take note of the following items when you call the **ListCenInterRegionTrafficQosPolicies*	- operation:
//
// - You must specify at least one of the **TransitRouterId*	- and **TrafficQosPolicyId*	- parameters.
//
// - If you do not specify the **TrafficQosPolicyId*	- parameter, the operation returns information about the QoS policy based on the values of the **TransitRouterId**, **TransitRouterAttachmentId**, **TrafficQosPolicyName**, and **TrafficQosPolicyDescription*	- parameters. In this case, information about the queues in the policy is not returned, and the **TrafficQosQueues*	- field is not included in the response.
//
// - If you specify the **TrafficQosPolicyId*	- parameter, the operation returns information about the QoS policy and its queues. The **TrafficQosQueues*	- field is included in the response. If the **TrafficQosQueues*	- field is an empty array, only the default queue exists in the QoS policy.
//
// - Make sure that you enter valid parameter values. If you enter an invalid parameter, the system returns a RequestId but does not return information about the QoS policy.
//
// @param request - ListCenInterRegionTrafficQosPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenInterRegionTrafficQosPoliciesResponse
func (client *Client) ListCenInterRegionTrafficQosPoliciesWithOptions(request *ListCenInterRegionTrafficQosPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListCenInterRegionTrafficQosPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListCenInterRegionTrafficQosPolicies operation to query Quality of Service (QoS) policies.
//
// Description:
//
// Take note of the following items when you call the **ListCenInterRegionTrafficQosPolicies*	- operation:
//
// - You must specify at least one of the **TransitRouterId*	- and **TrafficQosPolicyId*	- parameters.
//
// - If you do not specify the **TrafficQosPolicyId*	- parameter, the operation returns information about the QoS policy based on the values of the **TransitRouterId**, **TransitRouterAttachmentId**, **TrafficQosPolicyName**, and **TrafficQosPolicyDescription*	- parameters. In this case, information about the queues in the policy is not returned, and the **TrafficQosQueues*	- field is not included in the response.
//
// - If you specify the **TrafficQosPolicyId*	- parameter, the operation returns information about the QoS policy and its queues. The **TrafficQosQueues*	- field is included in the response. If the **TrafficQosQueues*	- field is an empty array, only the default queue exists in the QoS policy.
//
// - Make sure that you enter valid parameter values. If you enter an invalid parameter, the system returns a RequestId but does not return information about the QoS policy.
//
// @param request - ListCenInterRegionTrafficQosPoliciesRequest
//
// @return ListCenInterRegionTrafficQosPoliciesResponse
func (client *Client) ListCenInterRegionTrafficQosPolicies(request *ListCenInterRegionTrafficQosPoliciesRequest) (_result *ListCenInterRegionTrafficQosPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCenInterRegionTrafficQosPoliciesResponse{}
	_body, _err := client.ListCenInterRegionTrafficQosPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the queues of a Quality of Service (QoS) policy.
//
// Description:
//
// When you call this operation, you must specify at least one of the **TransitRouterId**, **TrafficQosPolicyId**, or **TrafficQosQueueId*	- parameters.
//
// Ensure that you specify valid parameter values. If you specify an invalid parameter, the system returns a **RequestId*	- but does not return information about the QoS policy.
//
// @param request - ListCenInterRegionTrafficQosQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCenInterRegionTrafficQosQueuesResponse
func (client *Client) ListCenInterRegionTrafficQosQueuesWithOptions(request *ListCenInterRegionTrafficQosQueuesRequest, runtime *dara.RuntimeOptions) (_result *ListCenInterRegionTrafficQosQueuesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the queues of a Quality of Service (QoS) policy.
//
// Description:
//
// When you call this operation, you must specify at least one of the **TransitRouterId**, **TrafficQosPolicyId**, or **TrafficQosQueueId*	- parameters.
//
// Ensure that you specify valid parameter values. If you specify an invalid parameter, the system returns a **RequestId*	- but does not return information about the QoS policy.
//
// @param request - ListCenInterRegionTrafficQosQueuesRequest
//
// @return ListCenInterRegionTrafficQosQueuesResponse
func (client *Client) ListCenInterRegionTrafficQosQueues(request *ListCenInterRegionTrafficQosQueuesRequest) (_result *ListCenInterRegionTrafficQosQueuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCenInterRegionTrafficQosQueuesResponse{}
	_body, _err := client.ListCenInterRegionTrafficQosQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListGrantVSwitchEnis operation to query which elastic network interfaces (ENIs) in a virtual private cloud (VPC) can serve as multicast sources or members.
//
// Description:
//
// Before you call the `ListGrantVSwitchEnis` operation, make sure that the VPC is connected to a Cloud Enterprise Network (CEN) instance. For more information, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261358.html).
//
// @param request - ListGrantVSwitchEnisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGrantVSwitchEnisResponse
func (client *Client) ListGrantVSwitchEnisWithOptions(request *ListGrantVSwitchEnisRequest, runtime *dara.RuntimeOptions) (_result *ListGrantVSwitchEnisResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListGrantVSwitchEnis operation to query which elastic network interfaces (ENIs) in a virtual private cloud (VPC) can serve as multicast sources or members.
//
// Description:
//
// Before you call the `ListGrantVSwitchEnis` operation, make sure that the VPC is connected to a Cloud Enterprise Network (CEN) instance. For more information, see [CreateTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261358.html).
//
// @param request - ListGrantVSwitchEnisRequest
//
// @return ListGrantVSwitchEnisResponse
func (client *Client) ListGrantVSwitchEnis(request *ListGrantVSwitchEnisRequest) (_result *ListGrantVSwitchEnisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGrantVSwitchEnisResponse{}
	_body, _err := client.ListGrantVSwitchEnisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListGrantVSwitchesToCen operation to query vSwitches in a cross-account VPC that is connected to a CEN instance.
//
// Description:
//
// Before you call the `ListGrantVSwitchesToCen` operation, make sure that the CEN instance has been granted permissions on the cross-account VPC. For more information, see [GrantInstanceToCen](https://help.aliyun.com/document_detail/126224.html).
//
// @param request - ListGrantVSwitchesToCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGrantVSwitchesToCenResponse
func (client *Client) ListGrantVSwitchesToCenWithOptions(request *ListGrantVSwitchesToCenRequest, runtime *dara.RuntimeOptions) (_result *ListGrantVSwitchesToCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListGrantVSwitchesToCen operation to query vSwitches in a cross-account VPC that is connected to a CEN instance.
//
// Description:
//
// Before you call the `ListGrantVSwitchesToCen` operation, make sure that the CEN instance has been granted permissions on the cross-account VPC. For more information, see [GrantInstanceToCen](https://help.aliyun.com/document_detail/126224.html).
//
// @param request - ListGrantVSwitchesToCenRequest
//
// @return ListGrantVSwitchesToCenResponse
func (client *Client) ListGrantVSwitchesToCen(request *ListGrantVSwitchesToCenRequest) (_result *ListGrantVSwitchesToCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListGrantVSwitchesToCenResponse{}
	_body, _err := client.ListGrantVSwitchesToCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// When you call the ListTagResources operation, you must specify at least one of the **ResourceId.N*	- and **Tag.N.Key*	- request parameters.
//
// - If you specify only **ResourceId.N**, the system queries the tags that are attached to the specified CEN instance.
//
// - If you specify only **Tag.N.Key**, the system queries all CEN instances that are associated with the specified tag key.
//
// - If you specify both **ResourceId.N*	- and **Tag.N.Key**, the system queries for tags that match the specified tag key and are attached to the specified CEN instance.
//
//   - The specified CEN instance must be associated with the specified tag key. Otherwise, an empty result is returned.
//
//   - If you specify multiple tag keys, the relationship between the tag keys is **AND**.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// When you call the ListTagResources operation, you must specify at least one of the **ResourceId.N*	- and **Tag.N.Key*	- request parameters.
//
// - If you specify only **ResourceId.N**, the system queries the tags that are attached to the specified CEN instance.
//
// - If you specify only **Tag.N.Key**, the system queries all CEN instances that are associated with the specified tag key.
//
// - If you specify both **ResourceId.N*	- and **Tag.N.Key**, the system queries for tags that match the specified tag key and are attached to the specified CEN instance.
//
//   - The specified CEN instance must be associated with the specified tag key. Otherwise, an empty result is returned.
//
//   - If you specify multiple tag keys, the relationship between the tag keys is **AND**.
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
// You can call the ListTrafficMarkingPolicies operation to query details about traffic marking policies, such as their status and priority.
//
// Description:
//
// When you call the **ListTrafficMarkingPolicies*	- operation:
//
// - You must specify at least one of the **TransitRouterId*	- and **TrafficMarkingPolicyId*	- parameters.
//
// - If you do not specify the **TrafficMarkingPolicyId*	- parameter, the operation returns only information about the traffic marking policy based on the **TransitRouterId**, **TrafficMarkingPolicyName**, and **TrafficMarkingPolicyDescription*	- parameters. Information about traffic classification rules is not returned. The **TrafficMatchRules*	- field is not included in the response.
//
// - If you specify the **TrafficMarkingPolicyId*	- parameter, the operation returns information about the traffic marking policy and its traffic classification rules. The **TrafficMatchRules*	- field is included in the response. If the **TrafficMatchRules*	- field is an empty array, no traffic classification rules are configured for the policy.
//
// @param request - ListTrafficMarkingPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrafficMarkingPoliciesResponse
func (client *Client) ListTrafficMarkingPoliciesWithOptions(request *ListTrafficMarkingPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListTrafficMarkingPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListTrafficMarkingPolicies operation to query details about traffic marking policies, such as their status and priority.
//
// Description:
//
// When you call the **ListTrafficMarkingPolicies*	- operation:
//
// - You must specify at least one of the **TransitRouterId*	- and **TrafficMarkingPolicyId*	- parameters.
//
// - If you do not specify the **TrafficMarkingPolicyId*	- parameter, the operation returns only information about the traffic marking policy based on the **TransitRouterId**, **TrafficMarkingPolicyName**, and **TrafficMarkingPolicyDescription*	- parameters. Information about traffic classification rules is not returned. The **TrafficMatchRules*	- field is not included in the response.
//
// - If you specify the **TrafficMarkingPolicyId*	- parameter, the operation returns information about the traffic marking policy and its traffic classification rules. The **TrafficMatchRules*	- field is included in the response. If the **TrafficMatchRules*	- field is an empty array, no traffic classification rules are configured for the policy.
//
// @param request - ListTrafficMarkingPoliciesRequest
//
// @return ListTrafficMarkingPoliciesResponse
func (client *Client) ListTrafficMarkingPolicies(request *ListTrafficMarkingPoliciesRequest) (_result *ListTrafficMarkingPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTrafficMarkingPoliciesResponse{}
	_body, _err := client.ListTrafficMarkingPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the zones that are available for an Enterprise Edition transit router in a specified region.
//
// Description:
//
// - You can call the **ListTransitRouterAvailableResource*	- operation to query regular zones or zones that support the multicast feature for an Enterprise Edition transit router in a specified region.
//
//   - If you do not set the **SupportMulticast*	- parameter to **true**, the system queries only the regular zones supported by the Enterprise Edition transit router.
//
//   - If you set the **SupportMulticast*	- parameter to **true**, the system queries only the zones that support the multicast feature for the Enterprise Edition transit router.
//
// - On May 31, 2022, Cloud Enterprise Network (CEN) upgraded the connection pattern for Enterprise Edition transit routers and Virtual Private Clouds (VPCs). After the upgrade, you do not need to specify a primary and a secondary zone when you connect an Enterprise Edition transit router to a VPC instance. Instead, you can specify one or more zones.
//
//   - If your Enterprise Edition transit router has not been upgraded, you must specify a primary and a secondary zone when you connect the transit router to a VPC instance. After you call the **ListTransitRouterAvailableResource*	- operation, you can retrieve information about the primary and secondary zones from the **MasterZones*	- and **SlaveZones*	- parameters.
//
//   - If your Enterprise Edition transit router has been upgraded, you can specify any zone when you connect the transit router to a VPC instance. After you call the **ListTransitRouterAvailableResource*	- operation, you can retrieve information about the supported zones from the **AvailableZones*	- parameter.
//
// For more information about the upgrade for Enterprise Edition transit routers, see [Upgrade of the VPC connection pattern for Enterprise Edition transit routers](https://help.aliyun.com/document_detail/434191.html).
//
// @param request - ListTransitRouterAvailableResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterAvailableResourceResponse
func (client *Client) ListTransitRouterAvailableResourceWithOptions(request *ListTransitRouterAvailableResourceRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterAvailableResourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the zones that are available for an Enterprise Edition transit router in a specified region.
//
// Description:
//
// - You can call the **ListTransitRouterAvailableResource*	- operation to query regular zones or zones that support the multicast feature for an Enterprise Edition transit router in a specified region.
//
//   - If you do not set the **SupportMulticast*	- parameter to **true**, the system queries only the regular zones supported by the Enterprise Edition transit router.
//
//   - If you set the **SupportMulticast*	- parameter to **true**, the system queries only the zones that support the multicast feature for the Enterprise Edition transit router.
//
// - On May 31, 2022, Cloud Enterprise Network (CEN) upgraded the connection pattern for Enterprise Edition transit routers and Virtual Private Clouds (VPCs). After the upgrade, you do not need to specify a primary and a secondary zone when you connect an Enterprise Edition transit router to a VPC instance. Instead, you can specify one or more zones.
//
//   - If your Enterprise Edition transit router has not been upgraded, you must specify a primary and a secondary zone when you connect the transit router to a VPC instance. After you call the **ListTransitRouterAvailableResource*	- operation, you can retrieve information about the primary and secondary zones from the **MasterZones*	- and **SlaveZones*	- parameters.
//
//   - If your Enterprise Edition transit router has been upgraded, you can specify any zone when you connect the transit router to a VPC instance. After you call the **ListTransitRouterAvailableResource*	- operation, you can retrieve information about the supported zones from the **AvailableZones*	- parameter.
//
// For more information about the upgrade for Enterprise Edition transit routers, see [Upgrade of the VPC connection pattern for Enterprise Edition transit routers](https://help.aliyun.com/document_detail/434191.html).
//
// @param request - ListTransitRouterAvailableResourceRequest
//
// @return ListTransitRouterAvailableResourceResponse
func (client *Client) ListTransitRouterAvailableResource(request *ListTransitRouterAvailableResourceRequest) (_result *ListTransitRouterAvailableResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterAvailableResourceResponse{}
	_body, _err := client.ListTransitRouterAvailableResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTransitRouterCidrWithOptions(request *ListTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterCidrResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTransitRouterCidrResponse
func (client *Client) ListTransitRouterCidr(request *ListTransitRouterCidrRequest) (_result *ListTransitRouterCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterCidrResponse{}
	_body, _err := client.ListTransitRouterCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the allocation details of a CIDR block.
//
// @param request - ListTransitRouterCidrAllocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterCidrAllocationResponse
func (client *Client) ListTransitRouterCidrAllocationWithOptions(request *ListTransitRouterCidrAllocationRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterCidrAllocationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the allocation details of a CIDR block.
//
// @param request - ListTransitRouterCidrAllocationRequest
//
// @return ListTransitRouterCidrAllocationResponse
func (client *Client) ListTransitRouterCidrAllocation(request *ListTransitRouterCidrAllocationRequest) (_result *ListTransitRouterCidrAllocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterCidrAllocationResponse{}
	_body, _err := client.ListTransitRouterCidrAllocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTransitRouterEcrAttachmentsWithOptions(request *ListTransitRouterEcrAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterEcrAttachmentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTransitRouterEcrAttachmentsResponse
func (client *Client) ListTransitRouterEcrAttachments(request *ListTransitRouterEcrAttachmentsRequest) (_result *ListTransitRouterEcrAttachmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterEcrAttachmentsResponse{}
	_body, _err := client.ListTransitRouterEcrAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListTransitRouterMulticastDomainAssociations operation to query the associations between multicast domains and vSwitches.
//
// Description:
//
// - When you call this operation, you must specify either the **TransitRouterMulticastDomainId*	- or **TransitRouterAttachmentId*	- request parameter. If you specify **TransitRouterAttachmentId**, the system queries the vSwitches that are associated with the multicast domain in the VPC. If you specify **TransitRouterMulticastDomainId**, the system queries the vSwitches that are associated with the specified multicast domain.
//
// - When you call the **ListTransitRouterMulticastDomainAssociations*	- operation, you must provide valid parameter values. If you provide an invalid parameter, the system returns a **RequestId*	- but does not return the association between the multicast domain and the vSwitch.
//
// @param request - ListTransitRouterMulticastDomainAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainAssociationsResponse
func (client *Client) ListTransitRouterMulticastDomainAssociationsWithOptions(request *ListTransitRouterMulticastDomainAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainAssociationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListTransitRouterMulticastDomainAssociations operation to query the associations between multicast domains and vSwitches.
//
// Description:
//
// - When you call this operation, you must specify either the **TransitRouterMulticastDomainId*	- or **TransitRouterAttachmentId*	- request parameter. If you specify **TransitRouterAttachmentId**, the system queries the vSwitches that are associated with the multicast domain in the VPC. If you specify **TransitRouterMulticastDomainId**, the system queries the vSwitches that are associated with the specified multicast domain.
//
// - When you call the **ListTransitRouterMulticastDomainAssociations*	- operation, you must provide valid parameter values. If you provide an invalid parameter, the system returns a **RequestId*	- but does not return the association between the multicast domain and the vSwitch.
//
// @param request - ListTransitRouterMulticastDomainAssociationsRequest
//
// @return ListTransitRouterMulticastDomainAssociationsResponse
func (client *Client) ListTransitRouterMulticastDomainAssociations(request *ListTransitRouterMulticastDomainAssociationsRequest) (_result *ListTransitRouterMulticastDomainAssociationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterMulticastDomainAssociationsResponse{}
	_body, _err := client.ListTransitRouterMulticastDomainAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// After a VPC instance is connected to an Enterprise Edition transit router, you can call the ListTransitRouterMulticastDomainVSwitches operation to query information about vSwitches that are attached to multicast domains in the virtual private cloud (VPC) instance.
//
// Description:
//
// When you call the ListTransitRouterMulticastDomainVSwitches operation, ensure that the parameter values are correct. If you enter incorrect parameter values, the operation returns a RequestId but does not display information about vSwitches that are attached to multicast domains in the VPC instance.
//
// @param request - ListTransitRouterMulticastDomainVSwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainVSwitchesResponse
func (client *Client) ListTransitRouterMulticastDomainVSwitchesWithOptions(request *ListTransitRouterMulticastDomainVSwitchesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainVSwitchesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// After a VPC instance is connected to an Enterprise Edition transit router, you can call the ListTransitRouterMulticastDomainVSwitches operation to query information about vSwitches that are attached to multicast domains in the virtual private cloud (VPC) instance.
//
// Description:
//
// When you call the ListTransitRouterMulticastDomainVSwitches operation, ensure that the parameter values are correct. If you enter incorrect parameter values, the operation returns a RequestId but does not display information about vSwitches that are attached to multicast domains in the VPC instance.
//
// @param request - ListTransitRouterMulticastDomainVSwitchesRequest
//
// @return ListTransitRouterMulticastDomainVSwitchesResponse
func (client *Client) ListTransitRouterMulticastDomainVSwitches(request *ListTransitRouterMulticastDomainVSwitchesRequest) (_result *ListTransitRouterMulticastDomainVSwitchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterMulticastDomainVSwitchesResponse{}
	_body, _err := client.ListTransitRouterMulticastDomainVSwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the ListTransitRouterMulticastDomains operation to query information about multicast domains, such as their statuses, IDs, and descriptions.
//
// Description:
//
// - You must specify both RegionId and CenId. If you specify only RegionId, no information about multicast domains is returned. You can also specify TransitRouterId or TransitRouterMulticastDomainId individually.
//
// - Ensure that you specify valid parameter values. If you specify an invalid parameter, the system returns a **RequestId*	- but does not return the details of the multicast domain.
//
// @param request - ListTransitRouterMulticastDomainsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastDomainsResponse
func (client *Client) ListTransitRouterMulticastDomainsWithOptions(request *ListTransitRouterMulticastDomainsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastDomainsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the ListTransitRouterMulticastDomains operation to query information about multicast domains, such as their statuses, IDs, and descriptions.
//
// Description:
//
// - You must specify both RegionId and CenId. If you specify only RegionId, no information about multicast domains is returned. You can also specify TransitRouterId or TransitRouterMulticastDomainId individually.
//
// - Ensure that you specify valid parameter values. If you specify an invalid parameter, the system returns a **RequestId*	- but does not return the details of the multicast domain.
//
// @param request - ListTransitRouterMulticastDomainsRequest
//
// @return ListTransitRouterMulticastDomainsResponse
func (client *Client) ListTransitRouterMulticastDomains(request *ListTransitRouterMulticastDomainsRequest) (_result *ListTransitRouterMulticastDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterMulticastDomainsResponse{}
	_body, _err := client.ListTransitRouterMulticastDomainsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about multicast members and sources in a multicast domain.
//
// Description:
//
// You can call the `ListTransitRouterMulticastGroups` operation to query information about multicast members and sources. These are collectively referred to as multicast resources.
//
// - If you specify the **GroupIpAddress*	- parameter, you can query information about the multicast resources in a specific multicast group.
//
// - If you specify the **VSwitchIds*	- parameter, you can query information about the multicast resources on specific vSwitches.
//
// - If you specify the **PeerTransitRouterMulticastDomains*	- parameter, you can query information about cross-region multicast resources.
//
// - If you specify the **ResourceType*	- parameter, you can query information about multicast resources of a specific resource type.
//
// - If you specify the **ResourceId*	- parameter, you can query information about the multicast resources on a specific resource.
//
// - If you specify only the **TransitRouterMulticastDomainId*	- parameter, you can query information about all multicast resources in the multicast domain.
//
// @param request - ListTransitRouterMulticastGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterMulticastGroupsResponse
func (client *Client) ListTransitRouterMulticastGroupsWithOptions(request *ListTransitRouterMulticastGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterMulticastGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about multicast members and sources in a multicast domain.
//
// Description:
//
// You can call the `ListTransitRouterMulticastGroups` operation to query information about multicast members and sources. These are collectively referred to as multicast resources.
//
// - If you specify the **GroupIpAddress*	- parameter, you can query information about the multicast resources in a specific multicast group.
//
// - If you specify the **VSwitchIds*	- parameter, you can query information about the multicast resources on specific vSwitches.
//
// - If you specify the **PeerTransitRouterMulticastDomains*	- parameter, you can query information about cross-region multicast resources.
//
// - If you specify the **ResourceType*	- parameter, you can query information about multicast resources of a specific resource type.
//
// - If you specify the **ResourceId*	- parameter, you can query information about the multicast resources on a specific resource.
//
// - If you specify only the **TransitRouterMulticastDomainId*	- parameter, you can query information about all multicast resources in the multicast domain.
//
// @param request - ListTransitRouterMulticastGroupsRequest
//
// @return ListTransitRouterMulticastGroupsResponse
func (client *Client) ListTransitRouterMulticastGroups(request *ListTransitRouterMulticastGroupsRequest) (_result *ListTransitRouterMulticastGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterMulticastGroupsResponse{}
	_body, _err := client.ListTransitRouterMulticastGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the ListTransitRouterPeerAttachments operation to query details about the inter-region connections of an Enterprise Edition transit router.
//
// Description:
//
// You can query information about the inter-region connections of an Enterprise Edition transit router in one of the following ways:
//
// - Query by the ID of the Enterprise Edition transit router instance.
//
// - Query by the ID of the Cloud Enterprise Network (CEN) instance and the region ID of the transit router instance.
//
// @param request - ListTransitRouterPeerAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterPeerAttachmentsResponse
func (client *Client) ListTransitRouterPeerAttachmentsWithOptions(request *ListTransitRouterPeerAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterPeerAttachmentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the ListTransitRouterPeerAttachments operation to query details about the inter-region connections of an Enterprise Edition transit router.
//
// Description:
//
// You can query information about the inter-region connections of an Enterprise Edition transit router in one of the following ways:
//
// - Query by the ID of the Enterprise Edition transit router instance.
//
// - Query by the ID of the Cloud Enterprise Network (CEN) instance and the region ID of the transit router instance.
//
// @param request - ListTransitRouterPeerAttachmentsRequest
//
// @return ListTransitRouterPeerAttachmentsResponse
func (client *Client) ListTransitRouterPeerAttachments(request *ListTransitRouterPeerAttachmentsRequest) (_result *ListTransitRouterPeerAttachmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterPeerAttachmentsResponse{}
	_body, _err := client.ListTransitRouterPeerAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the prefix list associations for the route table of an Enterprise Edition transit router.
//
// @param request - ListTransitRouterPrefixListAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterPrefixListAssociationResponse
func (client *Client) ListTransitRouterPrefixListAssociationWithOptions(request *ListTransitRouterPrefixListAssociationRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterPrefixListAssociationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the prefix list associations for the route table of an Enterprise Edition transit router.
//
// @param request - ListTransitRouterPrefixListAssociationRequest
//
// @return ListTransitRouterPrefixListAssociationResponse
func (client *Client) ListTransitRouterPrefixListAssociation(request *ListTransitRouterPrefixListAssociationRequest) (_result *ListTransitRouterPrefixListAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterPrefixListAssociationResponse{}
	_body, _err := client.ListTransitRouterPrefixListAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTransitRouterRouteEntriesWithOptions(request *ListTransitRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTransitRouterRouteEntriesResponse
func (client *Client) ListTransitRouterRouteEntries(request *ListTransitRouterRouteEntriesRequest) (_result *ListTransitRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterRouteEntriesResponse{}
	_body, _err := client.ListTransitRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the ListTransitRouterRouteTableAssociations operation to query the forwarding associations for a route table of an Enterprise Edition transit router or for a network instance connection.
//
// Description:
//
// When you call the **ListTransitRouterRouteTableAssociations*	- operation, you must specify at least one of the following request parameters: **TransitRouterRouteTableId*	- and **TransitRouterAttachmentId**.
//
// - If you specify only **TransitRouterRouteTableId**, the operation queries the network instance connections that are associated with the specified route table of the Enterprise Edition transit router.
//
// - If you specify only **TransitRouterAttachmentId**, the operation queries the route tables of the Enterprise Edition transit router that are associated with the specified network instance connection.
//
// - If you specify both **TransitRouterRouteTableId*	- and **TransitRouterAttachmentId**, the operation queries the forwarding association between the network instance connection and the route table of the Enterprise Edition transit router.
//
//   - If a forwarding association exists between the network instance connection and the route table, the details of the forwarding association are returned.
//
//   - If no forwarding association exists between the network instance connection and the route table, the **TransitRouterAssociations*	- array is empty in the response.
//
// Make sure that you provide valid parameter values when you call the **ListTransitRouterRouteTableAssociations*	- operation.
//
// If you provide an invalid parameter, the system returns a **RequestId*	- but does not query the forwarding associations for the route table of the Enterprise Edition transit router or the network instance connection.
//
// @param request - ListTransitRouterRouteTableAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTableAssociationsResponse
func (client *Client) ListTransitRouterRouteTableAssociationsWithOptions(request *ListTransitRouterRouteTableAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTableAssociationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the ListTransitRouterRouteTableAssociations operation to query the forwarding associations for a route table of an Enterprise Edition transit router or for a network instance connection.
//
// Description:
//
// When you call the **ListTransitRouterRouteTableAssociations*	- operation, you must specify at least one of the following request parameters: **TransitRouterRouteTableId*	- and **TransitRouterAttachmentId**.
//
// - If you specify only **TransitRouterRouteTableId**, the operation queries the network instance connections that are associated with the specified route table of the Enterprise Edition transit router.
//
// - If you specify only **TransitRouterAttachmentId**, the operation queries the route tables of the Enterprise Edition transit router that are associated with the specified network instance connection.
//
// - If you specify both **TransitRouterRouteTableId*	- and **TransitRouterAttachmentId**, the operation queries the forwarding association between the network instance connection and the route table of the Enterprise Edition transit router.
//
//   - If a forwarding association exists between the network instance connection and the route table, the details of the forwarding association are returned.
//
//   - If no forwarding association exists between the network instance connection and the route table, the **TransitRouterAssociations*	- array is empty in the response.
//
// Make sure that you provide valid parameter values when you call the **ListTransitRouterRouteTableAssociations*	- operation.
//
// If you provide an invalid parameter, the system returns a **RequestId*	- but does not query the forwarding associations for the route table of the Enterprise Edition transit router or the network instance connection.
//
// @param request - ListTransitRouterRouteTableAssociationsRequest
//
// @return ListTransitRouterRouteTableAssociationsResponse
func (client *Client) ListTransitRouterRouteTableAssociations(request *ListTransitRouterRouteTableAssociationsRequest) (_result *ListTransitRouterRouteTableAssociationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterRouteTableAssociationsResponse{}
	_body, _err := client.ListTransitRouterRouteTableAssociationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the route propagations for the route table of an Enterprise Edition transit router.
//
// @param request - ListTransitRouterRouteTablePropagationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTablePropagationsResponse
func (client *Client) ListTransitRouterRouteTablePropagationsWithOptions(request *ListTransitRouterRouteTablePropagationsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTablePropagationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the route propagations for the route table of an Enterprise Edition transit router.
//
// @param request - ListTransitRouterRouteTablePropagationsRequest
//
// @return ListTransitRouterRouteTablePropagationsResponse
func (client *Client) ListTransitRouterRouteTablePropagations(request *ListTransitRouterRouteTablePropagationsRequest) (_result *ListTransitRouterRouteTablePropagationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterRouteTablePropagationsResponse{}
	_body, _err := client.ListTransitRouterRouteTablePropagationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the ListTransitRouterRouteTables operation to query the route tables associated with an Enterprise Edition transit router.
//
// @param request - ListTransitRouterRouteTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterRouteTablesResponse
func (client *Client) ListTransitRouterRouteTablesWithOptions(request *ListTransitRouterRouteTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterRouteTablesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the ListTransitRouterRouteTables operation to query the route tables associated with an Enterprise Edition transit router.
//
// @param request - ListTransitRouterRouteTablesRequest
//
// @return ListTransitRouterRouteTablesResponse
func (client *Client) ListTransitRouterRouteTables(request *ListTransitRouterRouteTablesRequest) (_result *ListTransitRouterRouteTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterRouteTablesResponse{}
	_body, _err := client.ListTransitRouterRouteTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about virtual border router (VBR) connections for an Enterprise Edition transit router. The returned information includes the total number of connections, connection status, connection ID, and the payer for the network instance.
//
// Description:
//
// You can query the VBR connections of an Enterprise Edition transit router in one of the following ways:
//
// - Query all VBR connections of an Enterprise Edition transit router by specifying the ID of the transit router.
//
// - Query all VBR connections of an Enterprise Edition transit router by specifying the ID of the Cloud Enterprise Network (CEN) instance and the region ID of the transit router.
//
// - Query a specific VBR connection by specifying the connection ID in the TransitRouterAttachmentId parameter.
//
// @param request - ListTransitRouterVbrAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVbrAttachmentsResponse
func (client *Client) ListTransitRouterVbrAttachmentsWithOptions(request *ListTransitRouterVbrAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVbrAttachmentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about virtual border router (VBR) connections for an Enterprise Edition transit router. The returned information includes the total number of connections, connection status, connection ID, and the payer for the network instance.
//
// Description:
//
// You can query the VBR connections of an Enterprise Edition transit router in one of the following ways:
//
// - Query all VBR connections of an Enterprise Edition transit router by specifying the ID of the transit router.
//
// - Query all VBR connections of an Enterprise Edition transit router by specifying the ID of the Cloud Enterprise Network (CEN) instance and the region ID of the transit router.
//
// - Query a specific VBR connection by specifying the connection ID in the TransitRouterAttachmentId parameter.
//
// @param request - ListTransitRouterVbrAttachmentsRequest
//
// @return ListTransitRouterVbrAttachmentsResponse
func (client *Client) ListTransitRouterVbrAttachments(request *ListTransitRouterVbrAttachmentsRequest) (_result *ListTransitRouterVbrAttachmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterVbrAttachmentsResponse{}
	_body, _err := client.ListTransitRouterVbrAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about virtual private cloud (VPC) connections on an Enterprise Edition transit router, including the connection status, billing method, and zone, as well as the vSwitches and elastic network interfaces (ENIs) of the connected VPC.
//
// Description:
//
// You can query the VPC connections of an Enterprise Edition transit router in one of the following ways:
//
// - Specify the ID of an Enterprise Edition transit router to query all its VPC connections.
//
// - Specify the ID of a Cloud Enterprise Network (CEN) instance and the region ID of the Enterprise Edition transit router to query all its VPC connections.
//
// - Specify the ID of a region that contains an Enterprise Edition transit router to query all VPC connections in that region.
//
// @param request - ListTransitRouterVpcAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVpcAttachmentsResponse
func (client *Client) ListTransitRouterVpcAttachmentsWithOptions(request *ListTransitRouterVpcAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVpcAttachmentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about virtual private cloud (VPC) connections on an Enterprise Edition transit router, including the connection status, billing method, and zone, as well as the vSwitches and elastic network interfaces (ENIs) of the connected VPC.
//
// Description:
//
// You can query the VPC connections of an Enterprise Edition transit router in one of the following ways:
//
// - Specify the ID of an Enterprise Edition transit router to query all its VPC connections.
//
// - Specify the ID of a Cloud Enterprise Network (CEN) instance and the region ID of the Enterprise Edition transit router to query all its VPC connections.
//
// - Specify the ID of a region that contains an Enterprise Edition transit router to query all VPC connections in that region.
//
// @param request - ListTransitRouterVpcAttachmentsRequest
//
// @return ListTransitRouterVpcAttachmentsResponse
func (client *Client) ListTransitRouterVpcAttachments(request *ListTransitRouterVpcAttachmentsRequest) (_result *ListTransitRouterVpcAttachmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterVpcAttachmentsResponse{}
	_body, _err := client.ListTransitRouterVpcAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListTransitRouterVpnAttachments operation to query information about VPN connections, such as their status, IPsec-VPN connection ID, and billing method.
//
// Description:
//
// The ListTransitRouterVpnAttachments operation supports the following query methods:
//
// - Enter only **TransitRouterAttachmentId*	- to query a specific VPN connection.
//
// - Enter only **TransitRouterId*	- to query all VPN connections that are associated with the specified transit router.
//
// - Enter **CenId*	- and **RegionId*	- to query VPN connections in a specific region of the specified Cloud Enterprise Network (CEN) instance.
//
// When you call the **ListTransitRouterVpnAttachments*	- operation, make sure that you enter valid parameter values. If you enter invalid parameters, the system returns a **RequestId*	- but does not return the queried VPN connection information.
//
// @param request - ListTransitRouterVpnAttachmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRouterVpnAttachmentsResponse
func (client *Client) ListTransitRouterVpnAttachmentsWithOptions(request *ListTransitRouterVpnAttachmentsRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRouterVpnAttachmentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListTransitRouterVpnAttachments operation to query information about VPN connections, such as their status, IPsec-VPN connection ID, and billing method.
//
// Description:
//
// The ListTransitRouterVpnAttachments operation supports the following query methods:
//
// - Enter only **TransitRouterAttachmentId*	- to query a specific VPN connection.
//
// - Enter only **TransitRouterId*	- to query all VPN connections that are associated with the specified transit router.
//
// - Enter **CenId*	- and **RegionId*	- to query VPN connections in a specific region of the specified Cloud Enterprise Network (CEN) instance.
//
// When you call the **ListTransitRouterVpnAttachments*	- operation, make sure that you enter valid parameter values. If you enter invalid parameters, the system returns a **RequestId*	- but does not return the queried VPN connection information.
//
// @param request - ListTransitRouterVpnAttachmentsRequest
//
// @return ListTransitRouterVpnAttachmentsResponse
func (client *Client) ListTransitRouterVpnAttachments(request *ListTransitRouterVpnAttachmentsRequest) (_result *ListTransitRouterVpnAttachmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRouterVpnAttachmentsResponse{}
	_body, _err := client.ListTransitRouterVpnAttachmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ListTransitRouters operation to query information about transit routers in a Cloud Enterprise Network (CEN) instance, such as the instance type, status, instance ID, and whether the multicast feature is enabled.
//
// Description:
//
// When you call this operation to query transit routers in a CEN instance, you can specify the **RegionId*	- and **TransitRouterId*	- parameters. Note the following information about these parameters:
//
// - If you do not specify **RegionId*	- or **TransitRouterId**, all transit routers in the CEN instance are queried.
//
// - If you specify only **RegionId**, the transit routers in the specified region are queried.
//
// - If you specify only **TransitRouterId**, the specified transit router is queried.
//
// @param request - ListTransitRoutersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTransitRoutersResponse
func (client *Client) ListTransitRoutersWithOptions(request *ListTransitRoutersRequest, runtime *dara.RuntimeOptions) (_result *ListTransitRoutersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ListTransitRouters operation to query information about transit routers in a Cloud Enterprise Network (CEN) instance, such as the instance type, status, instance ID, and whether the multicast feature is enabled.
//
// Description:
//
// When you call this operation to query transit routers in a CEN instance, you can specify the **RegionId*	- and **TransitRouterId*	- parameters. Note the following information about these parameters:
//
// - If you do not specify **RegionId*	- or **TransitRouterId**, all transit routers in the CEN instance are queried.
//
// - If you specify only **RegionId**, the transit routers in the specified region are queried.
//
// - If you specify only **TransitRouterId**, the specified transit router is queried.
//
// @param request - ListTransitRoutersRequest
//
// @return ListTransitRoutersResponse
func (client *Client) ListTransitRouters(request *ListTransitRoutersRequest) (_result *ListTransitRoutersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTransitRoutersResponse{}
	_body, _err := client.ListTransitRoutersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ModifyCenAttribute operation modifies the name and description of a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// *ModifyCenAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the CEN instance is not modified immediately. The system modifies the instance in the background. You can call the **DescribeCens*	- operation to query the status of the CEN instance.
//
// - If a CEN instance is in the **Modifying*	- state, the modification is in progress. In this state, you can only query the instance and cannot perform other operations.
//
// - If a CEN instance is in the **Active*	- state, the modification is complete.
//
// @param request - ModifyCenAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenAttributeResponse
func (client *Client) ModifyCenAttributeWithOptions(request *ModifyCenAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ModifyCenAttribute operation modifies the name and description of a Cloud Enterprise Network (CEN) instance.
//
// Description:
//
// *ModifyCenAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the CEN instance is not modified immediately. The system modifies the instance in the background. You can call the **DescribeCens*	- operation to query the status of the CEN instance.
//
// - If a CEN instance is in the **Modifying*	- state, the modification is in progress. In this state, you can only query the instance and cannot perform other operations.
//
// - If a CEN instance is in the **Active*	- state, the modification is complete.
//
// @param request - ModifyCenAttributeRequest
//
// @return ModifyCenAttributeResponse
func (client *Client) ModifyCenAttribute(request *ModifyCenAttributeRequest) (_result *ModifyCenAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCenAttributeResponse{}
	_body, _err := client.ModifyCenAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ModifyCenBandwidthPackageAttribute operation to modify the name and description of a bandwidth plan.
//
// @param request - ModifyCenBandwidthPackageAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenBandwidthPackageAttributeResponse
func (client *Client) ModifyCenBandwidthPackageAttributeWithOptions(request *ModifyCenBandwidthPackageAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenBandwidthPackageAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyCenBandwidthPackageAttribute operation to modify the name and description of a bandwidth plan.
//
// @param request - ModifyCenBandwidthPackageAttributeRequest
//
// @return ModifyCenBandwidthPackageAttributeResponse
func (client *Client) ModifyCenBandwidthPackageAttribute(request *ModifyCenBandwidthPackageAttributeRequest) (_result *ModifyCenBandwidthPackageAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCenBandwidthPackageAttributeResponse{}
	_body, _err := client.ModifyCenBandwidthPackageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyCenBandwidthPackageSpecWithOptions(request *ModifyCenBandwidthPackageSpecRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenBandwidthPackageSpecResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyCenBandwidthPackageSpecResponse
func (client *Client) ModifyCenBandwidthPackageSpec(request *ModifyCenBandwidthPackageSpecRequest) (_result *ModifyCenBandwidthPackageSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCenBandwidthPackageSpecResponse{}
	_body, _err := client.ModifyCenBandwidthPackageSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - **Modifying**: indicates that the system is modifying the routing policy. You can only query the routing policy, but cannot perform other operations.
//
// - **Active**: indicates that the routing policy is modified.
//
// @param request - ModifyCenRouteMapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCenRouteMapResponse
func (client *Client) ModifyCenRouteMapWithOptions(request *ModifyCenRouteMapRequest, runtime *dara.RuntimeOptions) (_result *ModifyCenRouteMapResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - **Modifying**: indicates that the system is modifying the routing policy. You can only query the routing policy, but cannot perform other operations.
//
// - **Active**: indicates that the routing policy is modified.
//
// @param request - ModifyCenRouteMapRequest
//
// @return ModifyCenRouteMapResponse
func (client *Client) ModifyCenRouteMap(request *ModifyCenRouteMapRequest) (_result *ModifyCenRouteMapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyCenRouteMapResponse{}
	_body, _err := client.ModifyCenRouteMapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and capture window of a flow log.
//
// Description:
//
// This operation is executed asynchronously. After receiving a request, the system returns a **request ID*	- before it finishes modifying the flow log. The task is continued in the backend. You can call `DescribeFlowlogs` to check whether the task has been completed.
//
// - If the flow log is in the **Modifying*	- state, it is still being modified. In this case, you can query information about the flow log but cannot perform other operations on it.
//
// - If the flow log is in the **Active*	- state, the modification task has been completed.
//
// @param request - ModifyFlowLogAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttributeWithOptions(request *ModifyFlowLogAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyFlowLogAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// This operation is executed asynchronously. After receiving a request, the system returns a **request ID*	- before it finishes modifying the flow log. The task is continued in the backend. You can call `DescribeFlowlogs` to check whether the task has been completed.
//
// - If the flow log is in the **Modifying*	- state, it is still being modified. In this case, you can query information about the flow log but cannot perform other operations on it.
//
// - If the flow log is in the **Active*	- state, the modification task has been completed.
//
// @param request - ModifyFlowLogAttributeRequest
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttribute(request *ModifyFlowLogAttributeRequest) (_result *ModifyFlowLogAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFlowLogAttributeResponse{}
	_body, _err := client.ModifyFlowLogAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation modifies the payer for a cross-account network instance connection to a transit router.
//
// Description:
//
// You can use this operation to modify the payer for a cross-account connection to a transit router, but only if the connected network instance is a Virtual Private Cloud (VPC), virtual border router (VBR), or IPsec instance.
//
// @param request - ModifyGrantInstanceToTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGrantInstanceToTransitRouterResponse
func (client *Client) ModifyGrantInstanceToTransitRouterWithOptions(request *ModifyGrantInstanceToTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *ModifyGrantInstanceToTransitRouterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation modifies the payer for a cross-account network instance connection to a transit router.
//
// Description:
//
// You can use this operation to modify the payer for a cross-account connection to a transit router, but only if the connected network instance is a Virtual Private Cloud (VPC), virtual border router (VBR), or IPsec instance.
//
// @param request - ModifyGrantInstanceToTransitRouterRequest
//
// @return ModifyGrantInstanceToTransitRouterResponse
func (client *Client) ModifyGrantInstanceToTransitRouter(request *ModifyGrantInstanceToTransitRouterRequest) (_result *ModifyGrantInstanceToTransitRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyGrantInstanceToTransitRouterResponse{}
	_body, _err := client.ModifyGrantInstanceToTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a stream classification rule.
//
// @param request - ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) ModifyTrafficMatchRuleToTrafficMarkingPolicyWithOptions(request *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a stream classification rule.
//
// @param request - ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest
//
// @return ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse
func (client *Client) ModifyTrafficMatchRuleToTrafficMarkingPolicy(request *ModifyTrafficMatchRuleToTrafficMarkingPolicyRequest) (_result *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTrafficMatchRuleToTrafficMarkingPolicyResponse{}
	_body, _err := client.ModifyTrafficMatchRuleToTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyTransitRouteTableAggregationWithOptions(tmpReq *ModifyTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouteTableAggregationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ModifyTransitRouteTableAggregationRequest
//
// @return ModifyTransitRouteTableAggregationResponse
func (client *Client) ModifyTransitRouteTableAggregation(request *ModifyTransitRouteTableAggregationRequest) (_result *ModifyTransitRouteTableAggregationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTransitRouteTableAggregationResponse{}
	_body, _err := client.ModifyTransitRouteTableAggregationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the ModifyTransitRouterCidr operation to modify a CIDR block of a transit router.
//
// Description:
//
// - Before you modify a CIDR block of a transit router, review the [Limits on transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// - You cannot modify a CIDR block if IP addresses have already been allocated from it.
//
// - The **ModifyTransitRouterCidr*	- operation is synchronous if you do not change the **PublishCidrRoute*	- parameter. The modification takes effect immediately.
//
// - The **ModifyTransitRouterCidr*	- operation is asynchronous if you change the **PublishCidrRoute*	- parameter. The system returns a **RequestId**, but the CIDR block is not immediately modified. The modification task runs in the background. You can call the **ListTransitRouterCidr*	- operation to query the status of the modification.
//
//   - If the CIDR block information has not changed, the modification is in progress.
//
//   - If the CIDR block information is updated, the modification is successful.
//
// @param request - ModifyTransitRouterCidrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTransitRouterCidrResponse
func (client *Client) ModifyTransitRouterCidrWithOptions(request *ModifyTransitRouterCidrRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouterCidrResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the ModifyTransitRouterCidr operation to modify a CIDR block of a transit router.
//
// Description:
//
// - Before you modify a CIDR block of a transit router, review the [Limits on transit router CIDR blocks](https://help.aliyun.com/document_detail/462635.html).
//
// - You cannot modify a CIDR block if IP addresses have already been allocated from it.
//
// - The **ModifyTransitRouterCidr*	- operation is synchronous if you do not change the **PublishCidrRoute*	- parameter. The modification takes effect immediately.
//
// - The **ModifyTransitRouterCidr*	- operation is asynchronous if you change the **PublishCidrRoute*	- parameter. The system returns a **RequestId**, but the CIDR block is not immediately modified. The modification task runs in the background. You can call the **ListTransitRouterCidr*	- operation to query the status of the modification.
//
//   - If the CIDR block information has not changed, the modification is in progress.
//
//   - If the CIDR block information is updated, the modification is successful.
//
// @param request - ModifyTransitRouterCidrRequest
//
// @return ModifyTransitRouterCidrResponse
func (client *Client) ModifyTransitRouterCidr(request *ModifyTransitRouterCidrRequest) (_result *ModifyTransitRouterCidrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTransitRouterCidrResponse{}
	_body, _err := client.ModifyTransitRouterCidrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyTransitRouterMulticastDomainWithOptions(request *ModifyTransitRouterMulticastDomainRequest, runtime *dara.RuntimeOptions) (_result *ModifyTransitRouterMulticastDomainResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyTransitRouterMulticastDomainResponse
func (client *Client) ModifyTransitRouterMulticastDomain(request *ModifyTransitRouterMulticastDomainRequest) (_result *ModifyTransitRouterMulticastDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTransitRouterMulticastDomainResponse{}
	_body, _err := client.ModifyTransitRouterMulticastDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the MoveResourceGroup operation to move a Cloud Enterprise Network (CEN) instance or a bandwidth plan to a different resource group.
//
// Description:
//
// By default, Cloud Enterprise Network (CEN) instances and bandwidth plans belong to the default resource group. You can call the `MoveResourceGroup` operation to move a CEN instance or a bandwidth plan to a different resource group.
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the MoveResourceGroup operation to move a Cloud Enterprise Network (CEN) instance or a bandwidth plan to a different resource group.
//
// Description:
//
// By default, Cloud Enterprise Network (CEN) instances and bandwidth plans belong to the default resource group. You can call the `MoveResourceGroup` operation to move a CEN instance or a bandwidth plan to a different resource group.
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
func (client *Client) OpenTransitRouterServiceWithOptions(request *OpenTransitRouterServiceRequest, runtime *dara.RuntimeOptions) (_result *OpenTransitRouterServiceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OpenTransitRouterServiceResponse
func (client *Client) OpenTransitRouterService(request *OpenTransitRouterServiceRequest) (_result *OpenTransitRouterServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenTransitRouterServiceResponse{}
	_body, _err := client.OpenTransitRouterServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the PublishRouteEntries operation to advertise routes from a virtual private cloud (VPC) or a virtual border router (VBR) to a Cloud Enterprise Network (CEN) instance. If no route conflicts occur, other network instances attached to the CEN instance can learn the advertised routes.
//
// Description:
//
// The following table lists the default advertising status for different types of routes in CEN. You can call the PublishRouteEntries operation to advertise routes that are not advertised to CEN by default.
//
// | Route                                                | Instance of the route | Advertised to CEN by default |
//
// | ---------------------------------------------------- | --------------------- | ---------------------------- |
//
// | Route to an ECS instance                             | VPC                   | No                           |
//
// | Route to a VPN Gateway                               | VPC                   | No                           |
//
// | Route to a high availability (HA) virtual IP address | VPC                   | No                           |
//
// | Route to a router interface                          | VPC                   | No                           |
//
// | Route to an Elastic Network Interface (ENI)          | VPC                   | No                           |
//
// | Route to an IPv6 Gateway                             | VPC                   | No                           |
//
// | Route to a NAT Gateway                               | VPC                   | No                           |
//
// | VPC system route                                     | VPC                   | Yes                          |
//
// | Route to a data center                               | VBR                   | Yes                          |
//
// | Border Gateway Protocol (BGP) route                  | VBR                   | Yes                          |
//
// @param request - PublishRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRouteEntriesResponse
func (client *Client) PublishRouteEntriesWithOptions(request *PublishRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *PublishRouteEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the PublishRouteEntries operation to advertise routes from a virtual private cloud (VPC) or a virtual border router (VBR) to a Cloud Enterprise Network (CEN) instance. If no route conflicts occur, other network instances attached to the CEN instance can learn the advertised routes.
//
// Description:
//
// The following table lists the default advertising status for different types of routes in CEN. You can call the PublishRouteEntries operation to advertise routes that are not advertised to CEN by default.
//
// | Route                                                | Instance of the route | Advertised to CEN by default |
//
// | ---------------------------------------------------- | --------------------- | ---------------------------- |
//
// | Route to an ECS instance                             | VPC                   | No                           |
//
// | Route to a VPN Gateway                               | VPC                   | No                           |
//
// | Route to a high availability (HA) virtual IP address | VPC                   | No                           |
//
// | Route to a router interface                          | VPC                   | No                           |
//
// | Route to an Elastic Network Interface (ENI)          | VPC                   | No                           |
//
// | Route to an IPv6 Gateway                             | VPC                   | No                           |
//
// | Route to a NAT Gateway                               | VPC                   | No                           |
//
// | VPC system route                                     | VPC                   | Yes                          |
//
// | Route to a data center                               | VBR                   | Yes                          |
//
// | Border Gateway Protocol (BGP) route                  | VBR                   | Yes                          |
//
// @param request - PublishRouteEntriesRequest
//
// @return PublishRouteEntriesResponse
func (client *Client) PublishRouteEntries(request *PublishRouteEntriesRequest) (_result *PublishRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishRouteEntriesResponse{}
	_body, _err := client.PublishRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Re-advertises an aggregate route.
//
// Description:
//
// For aggregate routes that failed to be advertised or were partially advertised, fix the route issue and call the **RefreshTransitRouteTableAggregation*	- operation to re-advertise the aggregate routes to virtual private clouds (VPCs). If you use the following solutions, the aggregate route is automatically advertised without manual operations:
//
// - Delete associated forwarding correlations
//
// - Disable route synchronization
//
// - Delete the VPC route table
//
// - Delete the aggregate route
//
// You can call the **DescribeTransitRouteTableAggregationDetail*	- operation to view the advertisement status of the aggregate route.
//
// @param request - RefreshTransitRouteTableAggregationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshTransitRouteTableAggregationResponse
func (client *Client) RefreshTransitRouteTableAggregationWithOptions(request *RefreshTransitRouteTableAggregationRequest, runtime *dara.RuntimeOptions) (_result *RefreshTransitRouteTableAggregationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Description:
//
// For aggregate routes that failed to be advertised or were partially advertised, fix the route issue and call the **RefreshTransitRouteTableAggregation*	- operation to re-advertise the aggregate routes to virtual private clouds (VPCs). If you use the following solutions, the aggregate route is automatically advertised without manual operations:
//
// - Delete associated forwarding correlations
//
// - Disable route synchronization
//
// - Delete the VPC route table
//
// - Delete the aggregate route
//
// You can call the **DescribeTransitRouteTableAggregationDetail*	- operation to view the advertisement status of the aggregate route.
//
// @param request - RefreshTransitRouteTableAggregationRequest
//
// @return RefreshTransitRouteTableAggregationResponse
func (client *Client) RefreshTransitRouteTableAggregation(request *RefreshTransitRouteTableAggregationRequest) (_result *RefreshTransitRouteTableAggregationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshTransitRouteTableAggregationResponse{}
	_body, _err := client.RefreshTransitRouteTableAggregationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - If you specify a value for the **NetworkInterfaceIds*	- parameter, an ENI in the current region is to be specified as a multicast member. Make sure that the ENI and vSwitch are associated with the multicast group. For more information, see [AssociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429778.html).
//
// - If you specify a value for the **PeerTransitRouterMulticastDomains**, a multicast member in a multicast group that belongs to another region but has the same IP address as the current multicast group is to be specified as a multicast member for the current multicast group. Make sure that an inter-region connection is established between the regions. For more information, see [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html).
//
//	For example, you created Multicast Group 1 in Multicast Domain 1, which is in the China (Hangzhou) region. You created Multicast Group 2 in Multicast Domain 2, which is in the China (Shanghai) region. Multicast Group 1 and Multicast Group 2 use the same multicast IP address, and Multicast Member 2 is in Multicast Group 2 in the China (Shanghai) region. If you call the `RegisterTransitRouterMulticastGroupMembers` operation to add multicast members to Multicast Group 1 in the China (Hangzhou) region and set **PeerTransitRouterMulticastDomains*	- to the ID of Multicast Group 2, which is in the China (Shanghai) region, Multicast Member 2, which is in Multicast Domain 2 in the China (Shanghai) region is added to Multicast Group 1 in the China (Hangzhou) region.
//
// - `RegisterTransitRouterMulticastGroupMembers` is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of a multicast member.
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
func (client *Client) RegisterTransitRouterMulticastGroupMembersWithOptions(request *RegisterTransitRouterMulticastGroupMembersRequest, runtime *dara.RuntimeOptions) (_result *RegisterTransitRouterMulticastGroupMembersResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - If you specify a value for the **NetworkInterfaceIds*	- parameter, an ENI in the current region is to be specified as a multicast member. Make sure that the ENI and vSwitch are associated with the multicast group. For more information, see [AssociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429778.html).
//
// - If you specify a value for the **PeerTransitRouterMulticastDomains**, a multicast member in a multicast group that belongs to another region but has the same IP address as the current multicast group is to be specified as a multicast member for the current multicast group. Make sure that an inter-region connection is established between the regions. For more information, see [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html).
//
//	For example, you created Multicast Group 1 in Multicast Domain 1, which is in the China (Hangzhou) region. You created Multicast Group 2 in Multicast Domain 2, which is in the China (Shanghai) region. Multicast Group 1 and Multicast Group 2 use the same multicast IP address, and Multicast Member 2 is in Multicast Group 2 in the China (Shanghai) region. If you call the `RegisterTransitRouterMulticastGroupMembers` operation to add multicast members to Multicast Group 1 in the China (Hangzhou) region and set **PeerTransitRouterMulticastDomains*	- to the ID of Multicast Group 2, which is in the China (Shanghai) region, Multicast Member 2, which is in Multicast Domain 2 in the China (Shanghai) region is added to Multicast Group 1 in the China (Hangzhou) region.
//
// - `RegisterTransitRouterMulticastGroupMembers` is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call the `ListTransitRouterMulticastGroups` operation to query the status of a multicast member.
//
//   - If the multicast member is in the **Registering**, the multicast member is being created. In this case, you can query the multicast member but cannot perform other operations on the multicast member.
//
//   - If the multicast member is in the **Registered*	- state, the multicast member is created.
//
// @param request - RegisterTransitRouterMulticastGroupMembersRequest
//
// @return RegisterTransitRouterMulticastGroupMembersResponse
func (client *Client) RegisterTransitRouterMulticastGroupMembers(request *RegisterTransitRouterMulticastGroupMembersRequest) (_result *RegisterTransitRouterMulticastGroupMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterTransitRouterMulticastGroupMembersResponse{}
	_body, _err := client.RegisterTransitRouterMulticastGroupMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Use the RegisterTransitRouterMulticastGroupSources operation to create a multicast source. A multicast source enables one-to-many communication.
//
// Description:
//
// - You can specify only an Elastic Network Interface (ENI) as a multicast source.
//
// - `RegisterTransitRouterMulticastGroupSources` is an asynchronous operation. After you send a request, the system returns a **RequestId**. The multicast source is created in the background and is not immediately available. You can call the `ListTransitRouterMulticastGroups` operation to query the status of the multicast source.
//
//   - If a multicast source is in the **Registering*	- status, the multicast source is being created. In this status, you can only query the multicast source.
//
//   - If a multicast source is in the **Registered*	- status, the multicast source is created.
//
// ### Prerequisites
//
// Before you call `RegisterTransitRouterMulticastGroupSources`, ensure that the vSwitch to which the ENI belongs is associated with the multicast domain. For more information, see [AssociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429778.html).
//
// @param request - RegisterTransitRouterMulticastGroupSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterTransitRouterMulticastGroupSourcesResponse
func (client *Client) RegisterTransitRouterMulticastGroupSourcesWithOptions(request *RegisterTransitRouterMulticastGroupSourcesRequest, runtime *dara.RuntimeOptions) (_result *RegisterTransitRouterMulticastGroupSourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use the RegisterTransitRouterMulticastGroupSources operation to create a multicast source. A multicast source enables one-to-many communication.
//
// Description:
//
// - You can specify only an Elastic Network Interface (ENI) as a multicast source.
//
// - `RegisterTransitRouterMulticastGroupSources` is an asynchronous operation. After you send a request, the system returns a **RequestId**. The multicast source is created in the background and is not immediately available. You can call the `ListTransitRouterMulticastGroups` operation to query the status of the multicast source.
//
//   - If a multicast source is in the **Registering*	- status, the multicast source is being created. In this status, you can only query the multicast source.
//
//   - If a multicast source is in the **Registered*	- status, the multicast source is created.
//
// ### Prerequisites
//
// Before you call `RegisterTransitRouterMulticastGroupSources`, ensure that the vSwitch to which the ENI belongs is associated with the multicast domain. For more information, see [AssociateTransitRouterMulticastDomain](https://help.aliyun.com/document_detail/429778.html).
//
// @param request - RegisterTransitRouterMulticastGroupSourcesRequest
//
// @return RegisterTransitRouterMulticastGroupSourcesResponse
func (client *Client) RegisterTransitRouterMulticastGroupSources(request *RegisterTransitRouterMulticastGroupSourcesRequest) (_result *RegisterTransitRouterMulticastGroupSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterTransitRouterMulticastGroupSourcesResponse{}
	_body, _err := client.RegisterTransitRouterMulticastGroupSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes specified traffic classification rules from a traffic marking policy.
//
// Description:
//
// - When you call **RemoveTrafficMatchRuleFromTrafficMarkingPolicy**, take note of the following rules:
//
//   - If you specify the ID of a traffic classification rule in the **TrafficMarkRuleIds*	- parameter, the specified traffic classification rule is deleted.
//
//   - If you do not specify a traffic classification rule ID in the **TrafficMarkRuleIds*	- parameter, no operation is performed after you call this operation.
//
//     If you want to delete a traffic classification rule, you must specify the rule ID before you call this operation.
//
// - **RemoveTrafficMatchRuleFromTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of a traffic classification rule.
//
//   - If a traffic classification rule is in the **Deleting*	- state, the traffic classification rule is being deleted. In this case, you can query the traffic classification rule but cannot perform other operations.
//
//   - If a traffic classification rule cannot be found, the traffic classification rule is deleted.
//
// @param request - RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse
func (client *Client) RemoveTrafficMatchRuleFromTrafficMarkingPolicyWithOptions(request *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - When you call **RemoveTrafficMatchRuleFromTrafficMarkingPolicy**, take note of the following rules:
//
//   - If you specify the ID of a traffic classification rule in the **TrafficMarkRuleIds*	- parameter, the specified traffic classification rule is deleted.
//
//   - If you do not specify a traffic classification rule ID in the **TrafficMarkRuleIds*	- parameter, no operation is performed after you call this operation.
//
//     If you want to delete a traffic classification rule, you must specify the rule ID before you call this operation.
//
// - **RemoveTrafficMatchRuleFromTrafficMarkingPolicy*	- is an asynchronous operation. After you send a request, the system returns a **request ID*	- and runs the task in the background. You can call the **ListTrafficMarkingPolicies*	- operation to query the status of a traffic classification rule.
//
//   - If a traffic classification rule is in the **Deleting*	- state, the traffic classification rule is being deleted. In this case, you can query the traffic classification rule but cannot perform other operations.
//
//   - If a traffic classification rule cannot be found, the traffic classification rule is deleted.
//
// @param request - RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest
//
// @return RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse
func (client *Client) RemoveTrafficMatchRuleFromTrafficMarkingPolicy(request *RemoveTrafficMatchRuleFromTrafficMarkingPolicyRequest) (_result *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponse{}
	_body, _err := client.RemoveTrafficMatchRuleFromTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicyWithOptions(request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest, runtime *dara.RuntimeOptions) (_result *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse
// Deprecated
func (client *Client) RemoveTraficMatchRuleFromTrafficMarkingPolicy(request *RemoveTraficMatchRuleFromTrafficMarkingPolicyRequest) (_result *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTraficMatchRuleFromTrafficMarkingPolicyResponse{}
	_body, _err := client.RemoveTraficMatchRuleFromTrafficMarkingPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replaces the route table that is associated with a network instance connection.
//
// Description:
//
// - You can replace the route table that is associated with a network instance connection only if the network instance connection is created by an Enterprise Edition transit router.
//
// - **ReplaceTransitRouterRouteTableAssociation*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the operation continues to run in the background. You can call the **ListTransitRouterRouteTableAssociations*	- operation to query the status of the association.
//
//   - If the association status is **Replacing**, the route table is being replaced. In this state, you can only query the association and cannot perform other operations.
//
//   - If the association status is **Active**, the route table has been replaced.
//
// @param request - ReplaceTransitRouterRouteTableAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceTransitRouterRouteTableAssociationResponse
func (client *Client) ReplaceTransitRouterRouteTableAssociationWithOptions(request *ReplaceTransitRouterRouteTableAssociationRequest, runtime *dara.RuntimeOptions) (_result *ReplaceTransitRouterRouteTableAssociationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replaces the route table that is associated with a network instance connection.
//
// Description:
//
// - You can replace the route table that is associated with a network instance connection only if the network instance connection is created by an Enterprise Edition transit router.
//
// - **ReplaceTransitRouterRouteTableAssociation*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the operation continues to run in the background. You can call the **ListTransitRouterRouteTableAssociations*	- operation to query the status of the association.
//
//   - If the association status is **Replacing**, the route table is being replaced. In this state, you can only query the association and cannot perform other operations.
//
//   - If the association status is **Active**, the route table has been replaced.
//
// @param request - ReplaceTransitRouterRouteTableAssociationRequest
//
// @return ReplaceTransitRouterRouteTableAssociationResponse
func (client *Client) ReplaceTransitRouterRouteTableAssociation(request *ReplaceTransitRouterRouteTableAssociationRequest) (_result *ReplaceTransitRouterRouteTableAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReplaceTransitRouterRouteTableAssociationResponse{}
	_body, _err := client.ReplaceTransitRouterRouteTableAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - This operation is supported only by Basic Edition transit routers. An on-premises network associated with a VBR can use CEN to access only a cloud service that is deployed in the same region.
//
//	For example, if cloud services are deployed in the China (Beijing) region, only on-premises networks connected to VBRs in the China (Beijing) region can access the cloud services.
//
// - **ResolveAndRouteServiceInCen*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call **DescribeRouteServicesInCen*	- to query the status of a cloud service.
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
// - The VBR or CCN instance to which your on-premises network is connected is attached to a CEN instance.
//
// - A VPC that is deployed in the same region as the cloud service is attached to the CEN instance. For more information, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/65902.html).
//
// @param request - ResolveAndRouteServiceInCenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResolveAndRouteServiceInCenResponse
func (client *Client) ResolveAndRouteServiceInCenWithOptions(request *ResolveAndRouteServiceInCenRequest, runtime *dara.RuntimeOptions) (_result *ResolveAndRouteServiceInCenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - This operation is supported only by Basic Edition transit routers. An on-premises network associated with a VBR can use CEN to access only a cloud service that is deployed in the same region.
//
//	For example, if cloud services are deployed in the China (Beijing) region, only on-premises networks connected to VBRs in the China (Beijing) region can access the cloud services.
//
// - **ResolveAndRouteServiceInCen*	- is an asynchronous operation. After a request is sent, the system returns a **request ID*	- and runs the task in the background. You can call **DescribeRouteServicesInCen*	- to query the status of a cloud service.
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
// - The VBR or CCN instance to which your on-premises network is connected is attached to a CEN instance.
//
// - A VPC that is deployed in the same region as the cloud service is attached to the CEN instance. For more information, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/65902.html).
//
// @param request - ResolveAndRouteServiceInCenRequest
//
// @return ResolveAndRouteServiceInCenResponse
func (client *Client) ResolveAndRouteServiceInCen(request *ResolveAndRouteServiceInCenRequest) (_result *ResolveAndRouteServiceInCenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResolveAndRouteServiceInCenResponse{}
	_body, _err := client.ResolveAndRouteServiceInCenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - For more information about how to detach VPCs from Enterprise Edition transit routers, see [DeleteTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261220.html).
//
// - For more information about how to detach VBRs from Enterprise Edition transit routers, see [DeleteTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/261223.html).
//
// - For more information about how to detach IPsec-VPN connections from Enterprise Edition transit routers, see [DeleteTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/443992.html).
//
// - For more information about how to detach ECRs from Enterprise Edition transit routers, see [DeleteTransitRouterEcrAttachment](https://help.aliyun.com/document_detail/443992.html).
//
// - For more information about how to detach network instances from Basic Edition transit routers, see [DetachCenChildInstance](https://help.aliyun.com/document_detail/65915.html).
//
// @param request - RevokeInstanceFromTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromTransitRouterResponse
func (client *Client) RevokeInstanceFromTransitRouterWithOptions(request *RevokeInstanceFromTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *RevokeInstanceFromTransitRouterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - For more information about how to detach VPCs from Enterprise Edition transit routers, see [DeleteTransitRouterVpcAttachment](https://help.aliyun.com/document_detail/261220.html).
//
// - For more information about how to detach VBRs from Enterprise Edition transit routers, see [DeleteTransitRouterVbrAttachment](https://help.aliyun.com/document_detail/261223.html).
//
// - For more information about how to detach IPsec-VPN connections from Enterprise Edition transit routers, see [DeleteTransitRouterVpnAttachment](https://help.aliyun.com/document_detail/443992.html).
//
// - For more information about how to detach ECRs from Enterprise Edition transit routers, see [DeleteTransitRouterEcrAttachment](https://help.aliyun.com/document_detail/443992.html).
//
// - For more information about how to detach network instances from Basic Edition transit routers, see [DetachCenChildInstance](https://help.aliyun.com/document_detail/65915.html).
//
// @param request - RevokeInstanceFromTransitRouterRequest
//
// @return RevokeInstanceFromTransitRouterResponse
func (client *Client) RevokeInstanceFromTransitRouter(request *RevokeInstanceFromTransitRouterRequest) (_result *RevokeInstanceFromTransitRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeInstanceFromTransitRouterResponse{}
	_body, _err := client.RevokeInstanceFromTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call the RoutePrivateZoneInCenToVpc operation to configure the PrivateZone service.
//
// Description:
//
// Alibaba Cloud DNS PrivateZone is a private Domain Name System (DNS) resolution and management service that is based on a Virtual Private Cloud (VPC). After a virtual border router (VBR) instance or a Cloud Connect Network (CCN) instance is attached to a Cloud Enterprise Network (CEN) instance, the associated on-premises network can access the PrivateZone service through the CEN instance.
//
// - An on-premises network that is associated with a VBR or CCN instance can access the PrivateZone service only in the same region.
//
//	For example, if the PrivateZone service is in the China (Beijing) region, only on-premises networks that are associated with VBR instances in the China (Beijing) region or with CCN instances in the Chinese mainland can access the PrivateZone service.
//
// - The **RoutePrivateZoneInCenToVpc*	- operation is asynchronous. After a request is sent, the system returns a **RequestId**. The configuration is then added in the background. You can call the **DescribeCenPrivateZoneRoutes*	- operation to query the status of the PrivateZone service.
//
//   - If the PrivateZone service is in the **Creating*	- state, the configuration is being added. In this state, you can only query the configuration and cannot perform other operations.
//
//   - If the PrivateZone service is in the **Active*	- state, the configuration is complete.
//
//   - If the PrivateZone service is in the **Failed*	- state, the configuration failed.
//
// #### Prerequisites
//
// Before you call the **RoutePrivateZoneInCenToVpc*	- operation, make sure that the following conditions are met:
//
// - The PrivateZone service is deployed. For more information, see [Quick Start for Alibaba Cloud DNS PrivateZone](https://help.aliyun.com/document_detail/64627.html).
//
// - The VPC instance associated with the PrivateZone service and the VBR or CCN instance in the access region are attached to the same CEN instance. For more information, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/65902.html).
//
// - If an on-premises network connects to Alibaba Cloud through a CCN instance, and the CCN, VPC, and CEN instances belong to different accounts, the CCN instance must be authorized first. For more information, see [Cloud Connect Network authorization](https://help.aliyun.com/document_detail/106674.html).
//
// @param request - RoutePrivateZoneInCenToVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RoutePrivateZoneInCenToVpcResponse
func (client *Client) RoutePrivateZoneInCenToVpcWithOptions(request *RoutePrivateZoneInCenToVpcRequest, runtime *dara.RuntimeOptions) (_result *RoutePrivateZoneInCenToVpcResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call the RoutePrivateZoneInCenToVpc operation to configure the PrivateZone service.
//
// Description:
//
// Alibaba Cloud DNS PrivateZone is a private Domain Name System (DNS) resolution and management service that is based on a Virtual Private Cloud (VPC). After a virtual border router (VBR) instance or a Cloud Connect Network (CCN) instance is attached to a Cloud Enterprise Network (CEN) instance, the associated on-premises network can access the PrivateZone service through the CEN instance.
//
// - An on-premises network that is associated with a VBR or CCN instance can access the PrivateZone service only in the same region.
//
//	For example, if the PrivateZone service is in the China (Beijing) region, only on-premises networks that are associated with VBR instances in the China (Beijing) region or with CCN instances in the Chinese mainland can access the PrivateZone service.
//
// - The **RoutePrivateZoneInCenToVpc*	- operation is asynchronous. After a request is sent, the system returns a **RequestId**. The configuration is then added in the background. You can call the **DescribeCenPrivateZoneRoutes*	- operation to query the status of the PrivateZone service.
//
//   - If the PrivateZone service is in the **Creating*	- state, the configuration is being added. In this state, you can only query the configuration and cannot perform other operations.
//
//   - If the PrivateZone service is in the **Active*	- state, the configuration is complete.
//
//   - If the PrivateZone service is in the **Failed*	- state, the configuration failed.
//
// #### Prerequisites
//
// Before you call the **RoutePrivateZoneInCenToVpc*	- operation, make sure that the following conditions are met:
//
// - The PrivateZone service is deployed. For more information, see [Quick Start for Alibaba Cloud DNS PrivateZone](https://help.aliyun.com/document_detail/64627.html).
//
// - The VPC instance associated with the PrivateZone service and the VBR or CCN instance in the access region are attached to the same CEN instance. For more information, see [AttachCenChildInstance](https://help.aliyun.com/document_detail/65902.html).
//
// - If an on-premises network connects to Alibaba Cloud through a CCN instance, and the CCN, VPC, and CEN instances belong to different accounts, the CCN instance must be authorized first. For more information, see [Cloud Connect Network authorization](https://help.aliyun.com/document_detail/106674.html).
//
// @param request - RoutePrivateZoneInCenToVpcRequest
//
// @return RoutePrivateZoneInCenToVpcResponse
func (client *Client) RoutePrivateZoneInCenToVpc(request *RoutePrivateZoneInCenToVpcRequest) (_result *RoutePrivateZoneInCenToVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RoutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.RoutePrivateZoneInCenToVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the SetCenInterRegionBandwidthLimit operation to set, modify, or delete the inter-region bandwidth between two regions in a bandwidth plan for a Basic Edition transit router.
//
// Description:
//
// The target Cloud Enterprise Network (CEN) instance must be associated with a bandwidth plan. For more information, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html) and [AssociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65934.html).
//
// ### Limits
//
// The target Cloud Enterprise Network (CEN) instance already has a bandwidth plan. For more information, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html) and [AssociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65934.html).
//
// You can call the **SetCenInterRegionBandwidthLimit*	- API to set, modify, or delete the bandwidth for inter-region communication:
//
// - The **SetCenInterRegionBandwidthLimit*	- operation supports setting, modifying, or deleting the inter-region communication bandwidth for Basic Edition transit routers only.
//
// - You cannot modify the inter-region communication bandwidth if bandwidth multiplexing is enabled for the inter-region connection.
//
// ### Limits
//
// - The maximum bandwidth for an inter-region communication cannot exceed the peak bandwidth of the bandwidth plan.
//
// - The total bandwidth of all inter-region communications in a bandwidth plan cannot exceed the peak bandwidth of the plan.
//
// - You cannot modify the inter-region communication bandwidth if bandwidth multiplexing is enabled for the inter-region connection.
//
// - You can use the **SetCenInterRegionBandwidthLimit*	- API operation to set, modify, or delete the inter-region communication bandwidth for Basic Edition transit routers only.
//
//	To set, modify, or delete the inter-region communication bandwidth for an Enterprise Edition transit router, use the [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html), [UpdateTransitRouterPeerAttachmentAttribute](https://help.aliyun.com/document_detail/261229.html), and [DeleteTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261227.html) API operations.
//
// @param request - SetCenInterRegionBandwidthLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetCenInterRegionBandwidthLimitResponse
func (client *Client) SetCenInterRegionBandwidthLimitWithOptions(request *SetCenInterRegionBandwidthLimitRequest, runtime *dara.RuntimeOptions) (_result *SetCenInterRegionBandwidthLimitResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the SetCenInterRegionBandwidthLimit operation to set, modify, or delete the inter-region bandwidth between two regions in a bandwidth plan for a Basic Edition transit router.
//
// Description:
//
// The target Cloud Enterprise Network (CEN) instance must be associated with a bandwidth plan. For more information, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html) and [AssociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65934.html).
//
// ### Limits
//
// The target Cloud Enterprise Network (CEN) instance already has a bandwidth plan. For more information, see [CreateCenBandwidthPackage](https://help.aliyun.com/document_detail/65919.html) and [AssociateCenBandwidthPackage](https://help.aliyun.com/document_detail/65934.html).
//
// You can call the **SetCenInterRegionBandwidthLimit*	- API to set, modify, or delete the bandwidth for inter-region communication:
//
// - The **SetCenInterRegionBandwidthLimit*	- operation supports setting, modifying, or deleting the inter-region communication bandwidth for Basic Edition transit routers only.
//
// - You cannot modify the inter-region communication bandwidth if bandwidth multiplexing is enabled for the inter-region connection.
//
// ### Limits
//
// - The maximum bandwidth for an inter-region communication cannot exceed the peak bandwidth of the bandwidth plan.
//
// - The total bandwidth of all inter-region communications in a bandwidth plan cannot exceed the peak bandwidth of the plan.
//
// - You cannot modify the inter-region communication bandwidth if bandwidth multiplexing is enabled for the inter-region connection.
//
// - You can use the **SetCenInterRegionBandwidthLimit*	- API operation to set, modify, or delete the inter-region communication bandwidth for Basic Edition transit routers only.
//
//	To set, modify, or delete the inter-region communication bandwidth for an Enterprise Edition transit router, use the [CreateTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261363.html), [UpdateTransitRouterPeerAttachmentAttribute](https://help.aliyun.com/document_detail/261229.html), and [DeleteTransitRouterPeerAttachment](https://help.aliyun.com/document_detail/261227.html) API operations.
//
// @param request - SetCenInterRegionBandwidthLimitRequest
//
// @return SetCenInterRegionBandwidthLimitResponse
func (client *Client) SetCenInterRegionBandwidthLimit(request *SetCenInterRegionBandwidthLimitRequest) (_result *SetCenInterRegionBandwidthLimitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetCenInterRegionBandwidthLimitResponse{}
	_body, _err := client.SetCenInterRegionBandwidthLimitWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and attaches tags to resources.
//
// Description:
//
// - A tag consists of a tag key and a tag value. Both the tag key and tag value are required.
//
// - If you attach multiple tags to a Cloud Enterprise Network (CEN) instance, the tag keys must be unique for that instance.
//
// - You can attach a maximum of 20 tags to a CEN instance.
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - A tag consists of a tag key and a tag value. Both the tag key and tag value are required.
//
// - If you attach multiple tags to a Cloud Enterprise Network (CEN) instance, the tag keys must be unique for that instance.
//
// - You can attach a maximum of 20 tags to a CEN instance.
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
// Temporarily upgrades a subscription bandwidth plan of Cloud Enterprise Network (CEN).
//
// Description:
//
// Subscription bandwidth packages support temporary upgrade. You can increase the bandwidth for a specified period, responding to traffic fluctuations.
//
// The minimum upgrade interval supported is 3 hours. After payment is completed, the bandwidth is immediately upgraded without affecting the service.
//
// > After the specified time window ends, the bandwidth limit is restored to the original value. If the actual bandwidth exceeds the limit, packets may be dropped due to network traffic throttling. Plan your upgrade window and match the bandwidth peak to your needs.
//
// - Currently, the temporary upgrade feature is not enabled by default. To use it, contact your account manager.
//
// - Pay-as-you-go and expired subscription bandwidth packages do not support the temporary upgrade feature.
//
// - The **TempUpgradeCenBandwidthPackageSpec*	- operation is asynchronous. The system first returns a **RequestId**, while running the upgrade task in the background. Call the **DescribeCenBandwidthPackages*	- API to query the bandwidth package specifications. When they match your request, the upgrade is complete.
//
// @param request - TempUpgradeCenBandwidthPackageSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TempUpgradeCenBandwidthPackageSpecResponse
func (client *Client) TempUpgradeCenBandwidthPackageSpecWithOptions(request *TempUpgradeCenBandwidthPackageSpecRequest, runtime *dara.RuntimeOptions) (_result *TempUpgradeCenBandwidthPackageSpecResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Temporarily upgrades a subscription bandwidth plan of Cloud Enterprise Network (CEN).
//
// Description:
//
// Subscription bandwidth packages support temporary upgrade. You can increase the bandwidth for a specified period, responding to traffic fluctuations.
//
// The minimum upgrade interval supported is 3 hours. After payment is completed, the bandwidth is immediately upgraded without affecting the service.
//
// > After the specified time window ends, the bandwidth limit is restored to the original value. If the actual bandwidth exceeds the limit, packets may be dropped due to network traffic throttling. Plan your upgrade window and match the bandwidth peak to your needs.
//
// - Currently, the temporary upgrade feature is not enabled by default. To use it, contact your account manager.
//
// - Pay-as-you-go and expired subscription bandwidth packages do not support the temporary upgrade feature.
//
// - The **TempUpgradeCenBandwidthPackageSpec*	- operation is asynchronous. The system first returns a **RequestId**, while running the upgrade task in the background. Call the **DescribeCenBandwidthPackages*	- API to query the bandwidth package specifications. When they match your request, the upgrade is complete.
//
// @param request - TempUpgradeCenBandwidthPackageSpecRequest
//
// @return TempUpgradeCenBandwidthPackageSpecResponse
func (client *Client) TempUpgradeCenBandwidthPackageSpec(request *TempUpgradeCenBandwidthPackageSpecRequest) (_result *TempUpgradeCenBandwidthPackageSpecResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TempUpgradeCenBandwidthPackageSpecResponse{}
	_body, _err := client.TempUpgradeCenBandwidthPackageSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UnassociateCenBandwidthPackage operation to detach a bandwidth plan from a Cloud Enterprise Network (CEN) instance. After you detach the bandwidth plan, you can attach it to another CEN instance.
//
// Description:
//
// Before you call this operation, ensure that no cross-region bandwidth is configured for the bandwidth plan. To delete the cross-region bandwidth, see [SetCenInterRegionBandwidthLimit](https://help.aliyun.com/document_detail/65942.html).
//
// @param request - UnassociateCenBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnassociateCenBandwidthPackageResponse
func (client *Client) UnassociateCenBandwidthPackageWithOptions(request *UnassociateCenBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *UnassociateCenBandwidthPackageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UnassociateCenBandwidthPackage operation to detach a bandwidth plan from a Cloud Enterprise Network (CEN) instance. After you detach the bandwidth plan, you can attach it to another CEN instance.
//
// Description:
//
// Before you call this operation, ensure that no cross-region bandwidth is configured for the bandwidth plan. To delete the cross-region bandwidth, see [SetCenInterRegionBandwidthLimit](https://help.aliyun.com/document_detail/65942.html).
//
// @param request - UnassociateCenBandwidthPackageRequest
//
// @return UnassociateCenBandwidthPackageResponse
func (client *Client) UnassociateCenBandwidthPackage(request *UnassociateCenBandwidthPackageRequest) (_result *UnassociateCenBandwidthPackageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnassociateCenBandwidthPackageResponse{}
	_body, _err := client.UnassociateCenBandwidthPackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a connection to PrivateZone.
//
// Description:
//
// The **UnroutePrivateZoneInCenToVpc*	- operation is asynchronous. The system returns a **RequestId**, while the system runs the deletion task in the background. Call the **DescribeCenPrivateZoneRoutes*	- operation to query the PrivateZone status.
//
// - The **Deleting*	- state indicates the PrivateZone connection is being deleted. You can only perform the query operation.
//
// - When the specified PrivateZone connection is not found, it has been deleted.
//
// If the PrivateZone connection has an access region that is a Cloud Connect Network (CCN) region, you must first delete the PrivateZone connection for the CCN region before you delete the PrivateZone connections for other regions.
//
// @param request - UnroutePrivateZoneInCenToVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnroutePrivateZoneInCenToVpcResponse
func (client *Client) UnroutePrivateZoneInCenToVpcWithOptions(request *UnroutePrivateZoneInCenToVpcRequest, runtime *dara.RuntimeOptions) (_result *UnroutePrivateZoneInCenToVpcResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a connection to PrivateZone.
//
// Description:
//
// The **UnroutePrivateZoneInCenToVpc*	- operation is asynchronous. The system returns a **RequestId**, while the system runs the deletion task in the background. Call the **DescribeCenPrivateZoneRoutes*	- operation to query the PrivateZone status.
//
// - The **Deleting*	- state indicates the PrivateZone connection is being deleted. You can only perform the query operation.
//
// - When the specified PrivateZone connection is not found, it has been deleted.
//
// If the PrivateZone connection has an access region that is a Cloud Connect Network (CCN) region, you must first delete the PrivateZone connection for the CCN region before you delete the PrivateZone connections for other regions.
//
// @param request - UnroutePrivateZoneInCenToVpcRequest
//
// @return UnroutePrivateZoneInCenToVpcResponse
func (client *Client) UnroutePrivateZoneInCenToVpc(request *UnroutePrivateZoneInCenToVpcRequest) (_result *UnroutePrivateZoneInCenToVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnroutePrivateZoneInCenToVpcResponse{}
	_body, _err := client.UnroutePrivateZoneInCenToVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// Modifies the name and description of a quality of service (QoS) policy.
//
// @param request - UpdateCenInterRegionTrafficQosPolicyAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCenInterRegionTrafficQosPolicyAttributeResponse
func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttributeWithOptions(request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateCenInterRegionTrafficQosPolicyAttributeResponse
func (client *Client) UpdateCenInterRegionTrafficQosPolicyAttribute(request *UpdateCenInterRegionTrafficQosPolicyAttributeRequest) (_result *UpdateCenInterRegionTrafficQosPolicyAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCenInterRegionTrafficQosPolicyAttributeResponse{}
	_body, _err := client.UpdateCenInterRegionTrafficQosPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateCenInterRegionTrafficQosQueueAttributeWithOptions(request *UpdateCenInterRegionTrafficQosQueueAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateCenInterRegionTrafficQosQueueAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateCenInterRegionTrafficQosQueueAttributeResponse
func (client *Client) UpdateCenInterRegionTrafficQosQueueAttribute(request *UpdateCenInterRegionTrafficQosQueueAttributeRequest) (_result *UpdateCenInterRegionTrafficQosQueueAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCenInterRegionTrafficQosQueueAttributeResponse{}
	_body, _err := client.UpdateCenInterRegionTrafficQosQueueAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateTrafficMarkingPolicyAttributeWithOptions(request *UpdateTrafficMarkingPolicyAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTrafficMarkingPolicyAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateTrafficMarkingPolicyAttributeResponse
func (client *Client) UpdateTrafficMarkingPolicyAttribute(request *UpdateTrafficMarkingPolicyAttributeRequest) (_result *UpdateTrafficMarkingPolicyAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTrafficMarkingPolicyAttributeResponse{}
	_body, _err := client.UpdateTrafficMarkingPolicyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpdateTransitRouter operation to modify the name and description of a TransitRouter instance.
//
// Description:
//
// *UpdateTransitRouter*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**. The TransitRouter instance is not immediately modified because the modification task runs in the background. You can call the **ListTransitRouters*	- operation to query the status of the TransitRouter instance.
//
// - If a TransitRouter instance is in the **Modifying*	- state, the instance is being modified. In this state, you can only query the instance. You cannot perform other operations.
//
// - If a TransitRouter instance is in the **Active*	- state, the modification is complete.
//
// @param request - UpdateTransitRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterResponse
func (client *Client) UpdateTransitRouterWithOptions(request *UpdateTransitRouterRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpdateTransitRouter operation to modify the name and description of a TransitRouter instance.
//
// Description:
//
// *UpdateTransitRouter*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**. The TransitRouter instance is not immediately modified because the modification task runs in the background. You can call the **ListTransitRouters*	- operation to query the status of the TransitRouter instance.
//
// - If a TransitRouter instance is in the **Modifying*	- state, the instance is being modified. In this state, you can only query the instance. You cannot perform other operations.
//
// - If a TransitRouter instance is in the **Active*	- state, the modification is complete.
//
// @param request - UpdateTransitRouterRequest
//
// @return UpdateTransitRouterResponse
func (client *Client) UpdateTransitRouter(request *UpdateTransitRouterRequest) (_result *UpdateTransitRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterResponse{}
	_body, _err := client.UpdateTransitRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name and description of an ECR attachment for an Enterprise Edition Transit Router.
//
// Description:
//
// `UpdateTransitRouterEcrAttachmentAttribute` is an asynchronous call. After you send a request, the system returns a request ID, but the ECR attachment is not modified immediately. The modification task runs in the background. You can call `ListTransitRouterEcrAttachments` to query the status of the ECR attachment.
//
// If an ECR attachment is in the `Modifying` state, you can only query the attachment and cannot perform other operations on it. When the attachment enters the `Attached` state, the modification is complete.
//
// @param request - UpdateTransitRouterEcrAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterEcrAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterEcrAttachmentAttributeWithOptions(request *UpdateTransitRouterEcrAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterEcrAttachmentAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of an ECR attachment for an Enterprise Edition Transit Router.
//
// Description:
//
// `UpdateTransitRouterEcrAttachmentAttribute` is an asynchronous call. After you send a request, the system returns a request ID, but the ECR attachment is not modified immediately. The modification task runs in the background. You can call `ListTransitRouterEcrAttachments` to query the status of the ECR attachment.
//
// If an ECR attachment is in the `Modifying` state, you can only query the attachment and cannot perform other operations on it. When the attachment enters the `Attached` state, the modification is complete.
//
// @param request - UpdateTransitRouterEcrAttachmentAttributeRequest
//
// @return UpdateTransitRouterEcrAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterEcrAttachmentAttribute(request *UpdateTransitRouterEcrAttachmentAttributeRequest) (_result *UpdateTransitRouterEcrAttachmentAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterEcrAttachmentAttributeResponse{}
	_body, _err := client.UpdateTransitRouterEcrAttachmentAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
// - If an inter-region connection is in the **Modifying*	- state, the inter-region connection is being modified. You can query the inter-region connection but cannot perform other operations.
//
// - If an inter-region connection is in the **Attached*	- state, the inter-region connection is modified.
//
// @param request - UpdateTransitRouterPeerAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterPeerAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterPeerAttachmentAttributeWithOptions(request *UpdateTransitRouterPeerAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterPeerAttachmentAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// - If an inter-region connection is in the **Modifying*	- state, the inter-region connection is being modified. You can query the inter-region connection but cannot perform other operations.
//
// - If an inter-region connection is in the **Attached*	- state, the inter-region connection is modified.
//
// @param request - UpdateTransitRouterPeerAttachmentAttributeRequest
//
// @return UpdateTransitRouterPeerAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterPeerAttachmentAttribute(request *UpdateTransitRouterPeerAttachmentAttributeRequest) (_result *UpdateTransitRouterPeerAttachmentAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterPeerAttachmentAttributeResponse{}
	_body, _err := client.UpdateTransitRouterPeerAttachmentAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateTransitRouterRouteEntryWithOptions(request *UpdateTransitRouterRouteEntryRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterRouteEntryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateTransitRouterRouteEntryResponse
func (client *Client) UpdateTransitRouterRouteEntry(request *UpdateTransitRouterRouteEntryRequest) (_result *UpdateTransitRouterRouteEntryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterRouteEntryResponse{}
	_body, _err := client.UpdateTransitRouterRouteEntryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the UpdateTransitRouterRouteTable operation to modify the name and description of a route table for an Enterprise Edition transit router, or to enable or disable multi-region equal-cost multi-path (ECMP) routing.
//
// @param request - UpdateTransitRouterRouteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterRouteTableResponse
func (client *Client) UpdateTransitRouterRouteTableWithOptions(request *UpdateTransitRouterRouteTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterRouteTableResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the UpdateTransitRouterRouteTable operation to modify the name and description of a route table for an Enterprise Edition transit router, or to enable or disable multi-region equal-cost multi-path (ECMP) routing.
//
// @param request - UpdateTransitRouterRouteTableRequest
//
// @return UpdateTransitRouterRouteTableResponse
func (client *Client) UpdateTransitRouterRouteTable(request *UpdateTransitRouterRouteTableRequest) (_result *UpdateTransitRouterRouteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterRouteTableResponse{}
	_body, _err := client.UpdateTransitRouterRouteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and automatic route advertising setting for a virtual border router (VBR) connection on an Enterprise Edition transit router.
//
// Description:
//
// *UpdateTransitRouterVbrAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the operation is still in progress in the background. You can call the **ListTransitRouterVbrAttachments*	- operation to query the status of the VBR connection.
//
// - If a VBR connection is in the **Modifying*	- state, you can only query the VBR connection and cannot perform other operations.
//
// - If a VBR connection is in the **Attached*	- state, the modification is complete.
//
// @param request - UpdateTransitRouterVbrAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVbrAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVbrAttachmentAttributeWithOptions(request *UpdateTransitRouterVbrAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVbrAttachmentAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, and automatic route advertising setting for a virtual border router (VBR) connection on an Enterprise Edition transit router.
//
// Description:
//
// *UpdateTransitRouterVbrAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId**, but the operation is still in progress in the background. You can call the **ListTransitRouterVbrAttachments*	- operation to query the status of the VBR connection.
//
// - If a VBR connection is in the **Modifying*	- state, you can only query the VBR connection and cannot perform other operations.
//
// - If a VBR connection is in the **Attached*	- state, the modification is complete.
//
// @param request - UpdateTransitRouterVbrAttachmentAttributeRequest
//
// @return UpdateTransitRouterVbrAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVbrAttachmentAttribute(request *UpdateTransitRouterVbrAttachmentAttributeRequest) (_result *UpdateTransitRouterVbrAttachmentAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterVbrAttachmentAttributeResponse{}
	_body, _err := client.UpdateTransitRouterVbrAttachmentAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a VPC connection on an Enterprise Edition transit router and to control automatic route advertising to the VPC.
//
// Description:
//
// *UpdateTransitRouterVpcAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- and completes the modification in the background. To query the status of the VPC connection, call **ListTransitRouterVpcAttachments**.
//
// - If a VPC connection is in the **Modifying*	- state, you can only query it.
//
// - If a VPC connection is in the **Attached*	- state, the modification is complete.
//
// @param tmpReq - UpdateTransitRouterVpcAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVpcAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVpcAttachmentAttributeWithOptions(tmpReq *UpdateTransitRouterVpcAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpcAttachmentAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a VPC connection on an Enterprise Edition transit router and to control automatic route advertising to the VPC.
//
// Description:
//
// *UpdateTransitRouterVpcAttachmentAttribute*	- is an asynchronous operation. After you send a request, the system returns a **RequestId*	- and completes the modification in the background. To query the status of the VPC connection, call **ListTransitRouterVpcAttachments**.
//
// - If a VPC connection is in the **Modifying*	- state, you can only query it.
//
// - If a VPC connection is in the **Attached*	- state, the modification is complete.
//
// @param request - UpdateTransitRouterVpcAttachmentAttributeRequest
//
// @return UpdateTransitRouterVpcAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVpcAttachmentAttribute(request *UpdateTransitRouterVpcAttachmentAttributeRequest) (_result *UpdateTransitRouterVpcAttachmentAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterVpcAttachmentAttributeResponse{}
	_body, _err := client.UpdateTransitRouterVpcAttachmentAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateTransitRouterVpcAttachmentZonesWithOptions(request *UpdateTransitRouterVpcAttachmentZonesRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpcAttachmentZonesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateTransitRouterVpcAttachmentZonesResponse
func (client *Client) UpdateTransitRouterVpcAttachmentZones(request *UpdateTransitRouterVpcAttachmentZonesRequest) (_result *UpdateTransitRouterVpcAttachmentZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterVpcAttachmentZonesResponse{}
	_body, _err := client.UpdateTransitRouterVpcAttachmentZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a VPN connection attached to an Enterprise Edition Transit Router. You can modify the connection\\"s name, description, and automatic route publishing setting.
//
// @param request - UpdateTransitRouterVpnAttachmentAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTransitRouterVpnAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVpnAttachmentAttributeWithOptions(request *UpdateTransitRouterVpnAttachmentAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTransitRouterVpnAttachmentAttributeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a VPN connection attached to an Enterprise Edition Transit Router. You can modify the connection\\"s name, description, and automatic route publishing setting.
//
// @param request - UpdateTransitRouterVpnAttachmentAttributeRequest
//
// @return UpdateTransitRouterVpnAttachmentAttributeResponse
func (client *Client) UpdateTransitRouterVpnAttachmentAttribute(request *UpdateTransitRouterVpnAttachmentAttributeRequest) (_result *UpdateTransitRouterVpnAttachmentAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTransitRouterVpnAttachmentAttributeResponse{}
	_body, _err := client.UpdateTransitRouterVpnAttachmentAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the WithdrawPublishedRouteEntries operation to revoke routes published from a Virtual Private Cloud (VPC) or Virtual Border Router (VBR) instance to Cloud Enterprise Network.
//
// @param request - WithdrawPublishedRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WithdrawPublishedRouteEntriesResponse
func (client *Client) WithdrawPublishedRouteEntriesWithOptions(request *WithdrawPublishedRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *WithdrawPublishedRouteEntriesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the WithdrawPublishedRouteEntries operation to revoke routes published from a Virtual Private Cloud (VPC) or Virtual Border Router (VBR) instance to Cloud Enterprise Network.
//
// @param request - WithdrawPublishedRouteEntriesRequest
//
// @return WithdrawPublishedRouteEntriesResponse
func (client *Client) WithdrawPublishedRouteEntries(request *WithdrawPublishedRouteEntriesRequest) (_result *WithdrawPublishedRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &WithdrawPublishedRouteEntriesResponse{}
	_body, _err := client.WithdrawPublishedRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
