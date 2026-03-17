// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Activates a Smart Access Gateway (SAG) device.
//
// @param request - ActivateSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateSmartAccessGatewayResponse
func (client *Client) ActivateSmartAccessGatewayWithContext(ctx context.Context, request *ActivateSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *ActivateSmartAccessGatewayResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a flow log.
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
		Version:     dara.String("2018-03-13"),
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
// Creates an access control list (ACL) rule.
//
// @param request - AddACLRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddACLRuleResponse
func (client *Client) AddACLRuleWithContext(ctx context.Context, request *AddACLRuleRequest, runtime *dara.RuntimeOptions) (_result *AddACLRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestCidr) {
		query["DestCidr"] = request.DestCidr
	}

	if !dara.IsNil(request.DestPortRange) {
		query["DestPortRange"] = request.DestPortRange
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DpiGroupIds) {
		query["DpiGroupIds"] = request.DpiGroupIds
	}

	if !dara.IsNil(request.DpiSignatureIds) {
		query["DpiSignatureIds"] = request.DpiSignatureIds
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
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

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
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

	if !dara.IsNil(request.SourceCidr) {
		query["SourceCidr"] = request.SourceCidr
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddACLRule"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddACLRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a destination network address translation (DNAT) entry to a Smart Access Gateway (SAG) instance.
//
// @param request - AddDnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDnatEntryResponse
func (client *Client) AddDnatEntryWithContext(ctx context.Context, request *AddDnatEntryRequest, runtime *dara.RuntimeOptions) (_result *AddDnatEntryResponse, _err error) {
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

	if !dara.IsNil(request.InternalIp) {
		query["InternalIp"] = request.InternalIp
	}

	if !dara.IsNil(request.InternalPort) {
		query["InternalPort"] = request.InternalPort
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
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

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDnatEntry"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds DNS forwarding configurations to an SCG5000 or SCG5000-5G instance. The device version must be 3.4.2 or later.
//
// @param request - AddSmartAccessGatewayDnsForwardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSmartAccessGatewayDnsForwardResponse
func (client *Client) AddSmartAccessGatewayDnsForwardWithContext(ctx context.Context, request *AddSmartAccessGatewayDnsForwardRequest, runtime *dara.RuntimeOptions) (_result *AddSmartAccessGatewayDnsForwardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.MasterIp) {
		query["MasterIp"] = request.MasterIp
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OutboundPortIndex) {
		query["OutboundPortIndex"] = request.OutboundPortIndex
	}

	if !dara.IsNil(request.OutboundPortName) {
		query["OutboundPortName"] = request.OutboundPortName
	}

	if !dara.IsNil(request.OutboundPortType) {
		query["OutboundPortType"] = request.OutboundPortType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	if !dara.IsNil(request.SlaveIp) {
		query["SlaveIp"] = request.SlaveIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSmartAccessGatewayDnsForward"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSmartAccessGatewayDnsForwardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to add a source network address translation (SNAT) entry to a Smart Access Gateway (SAG) instance.
//
// @param request - AddSnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSnatEntryResponse
func (client *Client) AddSnatEntryWithContext(ctx context.Context, request *AddSnatEntryRequest, runtime *dara.RuntimeOptions) (_result *AddSnatEntryResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SnatIp) {
		query["SnatIp"] = request.SnatIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddSnatEntry"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddSnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an access control list (ACL) with a Smart Access Gateway (SAG) instance.
//
// @param request - AssociateACLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateACLResponse
func (client *Client) AssociateACLWithContext(ctx context.Context, request *AssociateACLRequest, runtime *dara.RuntimeOptions) (_result *AssociateACLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateACL"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateACLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a flow log with a Smart Access Gateway (SAG) instance.
//
// @param request - AssociateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateFlowLogResponse
func (client *Client) AssociateFlowLogWithContext(ctx context.Context, request *AssociateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *AssociateFlowLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateFlowLog"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateFlowLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies a Quality of Service (QoS) policy to a Smart Access Gateway (SAG) instance.
//
// @param request - AssociateQosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateQosResponse
func (client *Client) AssociateQosWithContext(ctx context.Context, request *AssociateQosRequest, runtime *dara.RuntimeOptions) (_result *AssociateQosResponse, _err error) {
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

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateQos"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateQosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a bandwidth plan for application acceleration with a Smart Access Gateway (SAG) instance.
//
// Description:
//
// Before you associate a bandwidth plan for application acceleration with an SAG instance, make sure that the following requirements are met:
//
//   - If you want to associate a bandwidth plan for application acceleration with an SAG CPE instance, the version of the SAG device associated with the SAG CPE instance must be 2.4.0 or later.
//
//     If the version of the SAG device is earlier than 2.4.0, update it. For more information, see [Update the version of an SAG device](https://help.aliyun.com/document_detail/163948.html).
//
//   - If you want to associate a bandwidth plan for application acceleration with an SAG app instance, the version of the SAG app must be 2.4.0 or later.
//
//     If the version of the SAG app is earlier than 2.4.0, update it. For more information, see [Install the SAG app](https://help.aliyun.com/document_detail/102544.html).
//
//   - The SAG instance to be associated is in the available state.
//
// @param request - AssociateSmartAGWithApplicationBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociateSmartAGWithApplicationBandwidthPackageResponse
func (client *Client) AssociateSmartAGWithApplicationBandwidthPackageWithContext(ctx context.Context, request *AssociateSmartAGWithApplicationBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *AssociateSmartAGWithApplicationBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationBandwidthPackageId) {
		query["ApplicationBandwidthPackageId"] = request.ApplicationBandwidthPackageId
	}

	if !dara.IsNil(request.AssociateConfigs) {
		query["AssociateConfigs"] = request.AssociateConfigs
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociateSmartAGWithApplicationBandwidthPackage"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociateSmartAGWithApplicationBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Smart Access Gateway (SAG) device with an SAG instance.
//
// @param request - BindSerialNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindSerialNumberResponse
func (client *Client) BindSerialNumberWithContext(ctx context.Context, request *BindSerialNumberRequest, runtime *dara.RuntimeOptions) (_result *BindSerialNumberResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindSerialNumber"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindSerialNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Smart Access Gateway (SAG) instance with a Cloud Connect Network (CCN) instance.
//
// @param request - BindSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindSmartAccessGatewayResponse
func (client *Client) BindSmartAccessGatewayWithContext(ctx context.Context, request *BindSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *BindSmartAccessGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnId) {
		query["CcnId"] = request.CcnId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGUid) {
		query["SmartAGUid"] = request.SmartAGUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Smart Access Gateway (SAG) instance with a virtual border router (VBR).
//
// @param request - BindVbrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindVbrResponse
func (client *Client) BindVbrWithContext(ctx context.Context, request *BindVbrRequest, runtime *dara.RuntimeOptions) (_result *BindVbrResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGUid) {
		query["SmartAGUid"] = request.SmartAGUid
	}

	if !dara.IsNil(request.VbrId) {
		query["VbrId"] = request.VbrId
	}

	if !dara.IsNil(request.VbrRegionId) {
		query["VbrRegionId"] = request.VbrRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindVbr"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindVbrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of a virtual customer-premises equipment (vCPE) device of Smart Access Gateway (SAG).
//
// @param request - ClearSagCipherRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearSagCipherResponse
func (client *Client) ClearSagCipherWithContext(ctx context.Context, request *ClearSagCipherRequest, runtime *dara.RuntimeOptions) (_result *ClearSagCipherResponse, _err error) {
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

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.SnNumber) {
		query["SnNumber"] = request.SnNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearSagCipher"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearSagCipherResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Clears the routable IP addresses of a Smart Access Gateway (SAG) instance.
//
// @param request - ClearSagRouteableAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearSagRouteableAddressResponse
func (client *Client) ClearSagRouteableAddressWithContext(ctx context.Context, request *ClearSagRouteableAddressRequest, runtime *dara.RuntimeOptions) (_result *ClearSagRouteableAddressResponse, _err error) {
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

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearSagRouteableAddress"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearSagRouteableAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL).
//
// @param request - CreateACLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateACLResponse
func (client *Client) CreateACLWithContext(ctx context.Context, request *CreateACLRequest, runtime *dara.RuntimeOptions) (_result *CreateACLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
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
		Action:      dara.String("CreateACL"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateACLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Cloud Connect Network (CCN) instance.
//
// Description:
//
// CCN is a matrix consisting of Alibaba Cloud distributed access gateways. It is an important component of Smart Access Gateway (SAG). After you associate an SAG instance with a CCN instance, the SAG instance connects the private networks associated with Alibaba Cloud. For more information, see [Overview of Cloud Connect Network](https://help.aliyun.com/document_detail/93667.html).
//
// @param request - CreateCloudConnectNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCloudConnectNetworkResponse
func (client *Client) CreateCloudConnectNetworkWithContext(ctx context.Context, request *CreateCloudConnectNetworkRequest, runtime *dara.RuntimeOptions) (_result *CreateCloudConnectNetworkResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	if !dara.IsNil(request.SnatCidrBlock) {
		query["SnatCidrBlock"] = request.SnatCidrBlock
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCloudConnectNetwork"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCloudConnectNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an enterprise code.
//
// @param request - CreateEnterpriseCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseCodeResponse
func (client *Client) CreateEnterpriseCodeWithContext(ctx context.Context, request *CreateEnterpriseCodeRequest, runtime *dara.RuntimeOptions) (_result *CreateEnterpriseCodeResponse, _err error) {
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

	if !dara.IsNil(request.EnterpriseCode) {
		query["EnterpriseCode"] = request.EnterpriseCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnterpriseCode"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnterpriseCodeResponse{}
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
// @param request - CreateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowLogResponse
func (client *Client) CreateFlowLogWithContext(ctx context.Context, request *CreateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveAging) {
		query["ActiveAging"] = request.ActiveAging
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InactiveAging) {
		query["InactiveAging"] = request.InactiveAging
	}

	if !dara.IsNil(request.LogstoreName) {
		query["LogstoreName"] = request.LogstoreName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetflowServerIp) {
		query["NetflowServerIp"] = request.NetflowServerIp
	}

	if !dara.IsNil(request.NetflowServerPort) {
		query["NetflowServerPort"] = request.NetflowServerPort
	}

	if !dara.IsNil(request.NetflowVersion) {
		query["NetflowVersion"] = request.NetflowVersion
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
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

	if !dara.IsNil(request.SlsRegionId) {
		query["SlsRegionId"] = request.SlsRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowLog"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a health check for a Smart Access Gateway (SAG) instance.
//
// @param request - CreateHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHealthCheckResponse
func (client *Client) CreateHealthCheckWithContext(ctx context.Context, request *CreateHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *CreateHealthCheckResponse, _err error) {
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

	if !dara.IsNil(request.DstIpAddr) {
		query["DstIpAddr"] = request.DstIpAddr
	}

	if !dara.IsNil(request.DstPort) {
		query["DstPort"] = request.DstPort
	}

	if !dara.IsNil(request.FailCountThreshold) {
		query["FailCountThreshold"] = request.FailCountThreshold
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

	if !dara.IsNil(request.ProbeCount) {
		query["ProbeCount"] = request.ProbeCount
	}

	if !dara.IsNil(request.ProbeInterval) {
		query["ProbeInterval"] = request.ProbeInterval
	}

	if !dara.IsNil(request.ProbeTimeout) {
		query["ProbeTimeout"] = request.ProbeTimeout
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

	if !dara.IsNil(request.RttFailThreshold) {
		query["RttFailThreshold"] = request.RttFailThreshold
	}

	if !dara.IsNil(request.RttThreshold) {
		query["RttThreshold"] = request.RttThreshold
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SrcIpAddr) {
		query["SrcIpAddr"] = request.SrcIpAddr
	}

	if !dara.IsNil(request.SrcPort) {
		query["SrcPort"] = request.SrcPort
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHealthCheck"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHealthCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a probing task for a Smart Access Gateway (SAG) device.
//
// Description:
//
//	  Only SAG-1000 devices whose software version is 2.7.0 or later support the probing feature.
//
//		- The SAG instance must have the deep packet inspection (DPI) feature enabled. You can call the [SetAdvancedMonitorState](https://help.aliyun.com/document_detail/476404.html) operation to enable or disable the DPI feature.
//
// @param request - CreateProbeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProbeTaskResponse
func (client *Client) CreateProbeTaskWithContext(ctx context.Context, request *CreateProbeTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateProbeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.PacketNumber) {
		query["PacketNumber"] = request.PacketNumber
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ProbeTaskSourceAddress) {
		query["ProbeTaskSourceAddress"] = request.ProbeTaskSourceAddress
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProbeTask"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProbeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a quality of service (QoS) policy.
//
// @param request - CreateQosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQosResponse
func (client *Client) CreateQosWithContext(ctx context.Context, request *CreateQosRequest, runtime *dara.RuntimeOptions) (_result *CreateQosResponse, _err error) {
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

	if !dara.IsNil(request.QosDescription) {
		query["QosDescription"] = request.QosDescription
	}

	if !dara.IsNil(request.QosName) {
		query["QosName"] = request.QosName
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
		Action:      dara.String("CreateQos"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a traffic throttling rule for a quality of service (QoS) policy.
//
// @param request - CreateQosCarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQosCarResponse
func (client *Client) CreateQosCarWithContext(ctx context.Context, request *CreateQosCarRequest, runtime *dara.RuntimeOptions) (_result *CreateQosCarResponse, _err error) {
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

	if !dara.IsNil(request.LimitType) {
		query["LimitType"] = request.LimitType
	}

	if !dara.IsNil(request.MaxBandwidthAbs) {
		query["MaxBandwidthAbs"] = request.MaxBandwidthAbs
	}

	if !dara.IsNil(request.MaxBandwidthPercent) {
		query["MaxBandwidthPercent"] = request.MaxBandwidthPercent
	}

	if !dara.IsNil(request.MinBandwidthAbs) {
		query["MinBandwidthAbs"] = request.MinBandwidthAbs
	}

	if !dara.IsNil(request.MinBandwidthPercent) {
		query["MinBandwidthPercent"] = request.MinBandwidthPercent
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

	if !dara.IsNil(request.PercentSourceType) {
		query["PercentSourceType"] = request.PercentSourceType
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
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
		Action:      dara.String("CreateQosCar"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQosCarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a traffic classification rule for a quality of service (QoS) policy.
//
// Description:
//
// ## Prerequisites
//
// A traffic throttling rule is created. For more information, see [CreateQosCar](https://help.aliyun.com/document_detail/131806.html).
//
// @param request - CreateQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQosPolicyResponse
func (client *Client) CreateQosPolicyWithContext(ctx context.Context, request *CreateQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateQosPolicyResponse, _err error) {
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

	if !dara.IsNil(request.DestCidr) {
		query["DestCidr"] = request.DestCidr
	}

	if !dara.IsNil(request.DestPortRange) {
		query["DestPortRange"] = request.DestPortRange
	}

	if !dara.IsNil(request.DpiGroupIds) {
		query["DpiGroupIds"] = request.DpiGroupIds
	}

	if !dara.IsNil(request.DpiSignatureIds) {
		query["DpiSignatureIds"] = request.DpiSignatureIds
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
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

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
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

	if !dara.IsNil(request.SourceCidr) {
		query["SourceCidr"] = request.SourceCidr
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQosPolicy"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQosPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a subinterface for an Express Connect circuit.
//
// @param request - CreateSagExpressConnectInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSagExpressConnectInterfaceResponse
func (client *Client) CreateSagExpressConnectInterfaceWithContext(ctx context.Context, request *CreateSagExpressConnectInterfaceRequest, runtime *dara.RuntimeOptions) (_result *CreateSagExpressConnectInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Vlan) {
		query["Vlan"] = request.Vlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSagExpressConnectInterface"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSagExpressConnectInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a static route to a Smart Access Gateway (SAG) instance.
//
// @param request - CreateSagStaticRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSagStaticRouteResponse
func (client *Client) CreateSagStaticRouteWithContext(ctx context.Context, request *CreateSagStaticRouteRequest, runtime *dara.RuntimeOptions) (_result *CreateSagStaticRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidr) {
		query["DestinationCidr"] = request.DestinationCidr
	}

	if !dara.IsNil(request.NextHop) {
		query["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Vlan) {
		query["Vlan"] = request.Vlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSagStaticRoute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSagStaticRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a service address for a Smart Access Gateway (SAG) device.
//
// @param request - CreateServiceAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceAddressResponse
func (client *Client) CreateServiceAddressWithContext(ctx context.Context, request *CreateServiceAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceAddress"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Smart Access Gateway (SAG) CPE or vCPE instance.
//
// @param request - CreateSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmartAccessGatewayResponse
func (client *Client) CreateSmartAccessGatewayWithContext(ctx context.Context, request *CreateSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *CreateSmartAccessGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlreadyHaveSag) {
		query["AlreadyHaveSag"] = request.AlreadyHaveSag
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.BuyerMessage) {
		query["BuyerMessage"] = request.BuyerMessage
	}

	if !dara.IsNil(request.CPEVersion) {
		query["CPEVersion"] = request.CPEVersion
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HaType) {
		query["HaType"] = request.HaType
	}

	if !dara.IsNil(request.HardWareSpec) {
		query["HardWareSpec"] = request.HardWareSpec
	}

	if !dara.IsNil(request.MaxBandWidth) {
		query["MaxBandWidth"] = request.MaxBandWidth
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

	if !dara.IsNil(request.ReceiverAddress) {
		query["ReceiverAddress"] = request.ReceiverAddress
	}

	if !dara.IsNil(request.ReceiverCity) {
		query["ReceiverCity"] = request.ReceiverCity
	}

	if !dara.IsNil(request.ReceiverCountry) {
		query["ReceiverCountry"] = request.ReceiverCountry
	}

	if !dara.IsNil(request.ReceiverDistrict) {
		query["ReceiverDistrict"] = request.ReceiverDistrict
	}

	if !dara.IsNil(request.ReceiverEmail) {
		query["ReceiverEmail"] = request.ReceiverEmail
	}

	if !dara.IsNil(request.ReceiverMobile) {
		query["ReceiverMobile"] = request.ReceiverMobile
	}

	if !dara.IsNil(request.ReceiverName) {
		query["ReceiverName"] = request.ReceiverName
	}

	if !dara.IsNil(request.ReceiverPhone) {
		query["ReceiverPhone"] = request.ReceiverPhone
	}

	if !dara.IsNil(request.ReceiverState) {
		query["ReceiverState"] = request.ReceiverState
	}

	if !dara.IsNil(request.ReceiverTown) {
		query["ReceiverTown"] = request.ReceiverTown
	}

	if !dara.IsNil(request.ReceiverZip) {
		query["ReceiverZip"] = request.ReceiverZip
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
		Action:      dara.String("CreateSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a client account.
//
// @param request - CreateSmartAccessGatewayClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmartAccessGatewayClientUserResponse
func (client *Client) CreateSmartAccessGatewayClientUserWithContext(ctx context.Context, request *CreateSmartAccessGatewayClientUserRequest, runtime *dara.RuntimeOptions) (_result *CreateSmartAccessGatewayClientUserResponse, _err error) {
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

	if !dara.IsNil(request.ClientIp) {
		query["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserMail) {
		query["UserMail"] = request.UserMail
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmartAccessGatewayClientUser"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmartAccessGatewayClientUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Smart Access Gateway (SAG) app instance.
//
// @param request - CreateSmartAccessGatewaySoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSmartAccessGatewaySoftwareResponse
func (client *Client) CreateSmartAccessGatewaySoftwareWithContext(ctx context.Context, request *CreateSmartAccessGatewaySoftwareRequest, runtime *dara.RuntimeOptions) (_result *CreateSmartAccessGatewaySoftwareResponse, _err error) {
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

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.DataPlan) {
		query["DataPlan"] = request.DataPlan
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.UserCount) {
		query["UserCount"] = request.UserCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSmartAccessGatewaySoftware"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSmartAccessGatewaySoftwareResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a flow log.
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
		Version:     dara.String("2018-03-13"),
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
// Deletes an access control list (ACL).
//
// @param request - DeleteACLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteACLResponse
func (client *Client) DeleteACLWithContext(ctx context.Context, request *DeleteACLRequest, runtime *dara.RuntimeOptions) (_result *DeleteACLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
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
		Action:      dara.String("DeleteACL"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteACLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a rule from an ACL.
//
// @param request - DeleteACLRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteACLRuleResponse
func (client *Client) DeleteACLRuleWithContext(ctx context.Context, request *DeleteACLRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteACLRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AcrId) {
		query["AcrId"] = request.AcrId
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
		Action:      dara.String("DeleteACLRule"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteACLRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Cloud Connect Network (CCN) instance.
//
// Description:
//
// >  Make sure that the CCN instance you want to delete is not associated with a Smart Access Gateway (SAG) instance or a Cloud Enterprise Network (CEN) instance.
//
// @param request - DeleteCloudConnectNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCloudConnectNetworkResponse
func (client *Client) DeleteCloudConnectNetworkWithContext(ctx context.Context, request *DeleteCloudConnectNetworkRequest, runtime *dara.RuntimeOptions) (_result *DeleteCloudConnectNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnId) {
		query["CcnId"] = request.CcnId
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
		Action:      dara.String("DeleteCloudConnectNetwork"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCloudConnectNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a DNAT entry from a Smart Access Gateway (SAG) instance.
//
// @param request - DeleteDnatEntryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDnatEntryResponse
func (client *Client) DeleteDnatEntryWithContext(ctx context.Context, request *DeleteDnatEntryRequest, runtime *dara.RuntimeOptions) (_result *DeleteDnatEntryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DnatEntryId) {
		query["DnatEntryId"] = request.DnatEntryId
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

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDnatEntry"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDnatEntryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified enterprise code.
//
// Description:
//
// Before you call this operation, take note of the following rules:
//
//   - You cannot delete default enterprise codes.
//
//     To delete a default enterprise code, change it to a custom enterprise code and then delete it. For more information, see [UpdateEnterpriseCode](https://help.aliyun.com/document_detail/197700.html).
//
//   - You cannot delete enterprise codes that are associated with a Smart Access Gateway (SAG) APP instance.
//
//     To delete an enterprise code that is associated with an SAG APP instance, associate the SAG APP instance with another enterprise code, and then delete the enterprise code. For more information, see [UpdateSmartAGEnterpriseCode](https://help.aliyun.com/document_detail/197701.html).
//
// @param request - DeleteEnterpriseCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnterpriseCodeResponse
func (client *Client) DeleteEnterpriseCodeWithContext(ctx context.Context, request *DeleteEnterpriseCodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnterpriseCodeResponse, _err error) {
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

	if !dara.IsNil(request.EnterpriseCode) {
		query["EnterpriseCode"] = request.EnterpriseCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnterpriseCode"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnterpriseCodeResponse{}
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
// @param request - DeleteFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowLogResponse
func (client *Client) DeleteFlowLogWithContext(ctx context.Context, request *DeleteFlowLogRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DeleteFlowLog"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to delete a health check instance.
//
// @param request - DeleteHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHealthCheckResponse
func (client *Client) DeleteHealthCheckWithContext(ctx context.Context, request *DeleteHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *DeleteHealthCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HcInstanceId) {
		query["HcInstanceId"] = request.HcInstanceId
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
		Action:      dara.String("DeleteHealthCheck"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHealthCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a probe task.
//
// @param request - DeleteProbeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProbeTaskResponse
func (client *Client) DeleteProbeTaskWithContext(ctx context.Context, request *DeleteProbeTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteProbeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProbeTaskId) {
		query["ProbeTaskId"] = request.ProbeTaskId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProbeTask"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProbeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Quality of Service (QoS) policy.
//
// @param request - DeleteQosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQosResponse
func (client *Client) DeleteQosWithContext(ctx context.Context, request *DeleteQosRequest, runtime *dara.RuntimeOptions) (_result *DeleteQosResponse, _err error) {
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

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
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
		Action:      dara.String("DeleteQos"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a speed limiting rule of a Quality of Service (QoS) policy.
//
// @param request - DeleteQosCarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQosCarResponse
func (client *Client) DeleteQosCarWithContext(ctx context.Context, request *DeleteQosCarRequest, runtime *dara.RuntimeOptions) (_result *DeleteQosCarResponse, _err error) {
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

	if !dara.IsNil(request.QosCarId) {
		query["QosCarId"] = request.QosCarId
	}

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
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
		Action:      dara.String("DeleteQosCar"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQosCarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quintuple rule of a Quality of Service (QoS) policy.
//
// @param request - DeleteQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQosPolicyResponse
func (client *Client) DeleteQosPolicyWithContext(ctx context.Context, request *DeleteQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteQosPolicyResponse, _err error) {
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

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
	}

	if !dara.IsNil(request.QosPolicyId) {
		query["QosPolicyId"] = request.QosPolicyId
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
		Action:      dara.String("DeleteQosPolicy"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQosPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a route advertisement policy.
//
// @param request - DeleteRouteDistributionStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteDistributionStrategyResponse
func (client *Client) DeleteRouteDistributionStrategyWithContext(ctx context.Context, request *DeleteRouteDistributionStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteDistributionStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestCidrBlock) {
		query["DestCidrBlock"] = request.DestCidrBlock
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

	if !dara.IsNil(request.RouteSource) {
		query["RouteSource"] = request.RouteSource
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRouteDistributionStrategy"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteDistributionStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a subinterface from a leased line port.
//
// @param request - DeleteSagExpressConnectInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSagExpressConnectInterfaceResponse
func (client *Client) DeleteSagExpressConnectInterfaceWithContext(ctx context.Context, request *DeleteSagExpressConnectInterfaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteSagExpressConnectInterfaceResponse, _err error) {
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

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Vlan) {
		query["Vlan"] = request.Vlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSagExpressConnectInterface"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSagExpressConnectInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to delete a static route.
//
// @param request - DeleteSagStaticRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSagStaticRouteResponse
func (client *Client) DeleteSagStaticRouteWithContext(ctx context.Context, request *DeleteSagStaticRouteRequest, runtime *dara.RuntimeOptions) (_result *DeleteSagStaticRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidr) {
		query["DestinationCidr"] = request.DestinationCidr
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Vlan) {
		query["Vlan"] = request.Vlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSagStaticRoute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSagStaticRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service address from a Smart Access Gateway (SAG) device.
//
// @param request - DeleteServiceAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceAddressResponse
func (client *Client) DeleteServiceAddressWithContext(ctx context.Context, request *DeleteServiceAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteServiceAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceAddress"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Smart Access Gateway (SAG) instance.
//
// Description:
//
//	  The SAG instance that you want to delete is an SAG CPE instance or an SAG vCPE instance.
//
//		- The SAG instance that you want to delete is locked due to overdue payments.
//
//		- The SAG instance that you want to delete is not associated with a Cloud Connect Network (CCN) instance or a virtual border router (VBR). If the SAG instance is associated with a CCN instance or a VBR, dissociate the SAG instance from the CCN instance or VBR first. For more information, see [Detach a network](https://help.aliyun.com/document_detail/164903.html).
//
// @param request - DeleteSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmartAccessGatewayResponse
func (client *Client) DeleteSmartAccessGatewayWithContext(ctx context.Context, request *DeleteSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmartAccessGatewayResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a client account from a Smart Access Gateway (SAG) app instance.
//
// @param request - DeleteSmartAccessGatewayClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmartAccessGatewayClientUserResponse
func (client *Client) DeleteSmartAccessGatewayClientUserWithContext(ctx context.Context, request *DeleteSmartAccessGatewayClientUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmartAccessGatewayClientUserResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSmartAccessGatewayClientUser"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSmartAccessGatewayClientUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables DNS forwarding for SCG5000 or SCG5000-5G devices whose software version is 3.4.2 or later.
//
// @param request - DeleteSmartAccessGatewayDnsForwardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmartAccessGatewayDnsForwardResponse
func (client *Client) DeleteSmartAccessGatewayDnsForwardWithContext(ctx context.Context, request *DeleteSmartAccessGatewayDnsForwardRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmartAccessGatewayDnsForwardResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSmartAccessGatewayDnsForward"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSmartAccessGatewayDnsForwardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes an SNAT entry from a Smart Access Gateway (SAG) instance.
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
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSnatEntry"),
		Version:     dara.String("2018-03-13"),
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
// Queries the information about an access control list (ACL).
//
// @param request - DescribeACLAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeACLAttributeResponse
func (client *Client) DescribeACLAttributeWithContext(ctx context.Context, request *DescribeACLAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeACLAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeACLAttribute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeACLAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access control lists (ACLs) in a specified region.
//
// @param request - DescribeACLsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeACLsResponse
func (client *Client) DescribeACLsWithContext(ctx context.Context, request *DescribeACLsRequest, runtime *dara.RuntimeOptions) (_result *DescribeACLsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclIds) {
		query["AclIds"] = request.AclIds
	}

	if !dara.IsNil(request.AclType) {
		query["AclType"] = request.AclType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeACLs"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeACLsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Smart Access Gateway (SAG) instances in a region that can be associated with a Cloud Connect Network (CCN) instance.
//
// @param request - DescribeBindableSmartAccessGatewaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBindableSmartAccessGatewaysResponse
func (client *Client) DescribeBindableSmartAccessGatewaysWithContext(ctx context.Context, request *DescribeBindableSmartAccessGatewaysRequest, runtime *dara.RuntimeOptions) (_result *DescribeBindableSmartAccessGatewaysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnId) {
		query["CcnId"] = request.CcnId
	}

	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBindableSmartAccessGateways"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBindableSmartAccessGatewaysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the DNS settings of a Smart Access Gateway (SAG) instance associated with SAG app.
//
// @param request - DescribeClientUserDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientUserDNSResponse
func (client *Client) DescribeClientUserDNSWithContext(ctx context.Context, request *DescribeClientUserDNSRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientUserDNSResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClientUserDNS"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClientUserDNSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Cloud Connect Network (CCN) instances that you have created in a specific region.
//
// @param request - DescribeCloudConnectNetworksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudConnectNetworksResponse
func (client *Client) DescribeCloudConnectNetworksWithContext(ctx context.Context, request *DescribeCloudConnectNetworksRequest, runtime *dara.RuntimeOptions) (_result *DescribeCloudConnectNetworksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnId) {
		query["CcnId"] = request.CcnId
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCloudConnectNetworks"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCloudConnectNetworksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the automatic upgrade policy of a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeDeviceAutoUpgradePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceAutoUpgradePolicyResponse
func (client *Client) DescribeDeviceAutoUpgradePolicyWithContext(ctx context.Context, request *DescribeDeviceAutoUpgradePolicyRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceAutoUpgradePolicyResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.VersionType) {
		query["VersionType"] = request.VersionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceAutoUpgradePolicy"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceAutoUpgradePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries DNAT entries that are associated with a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeDnatEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDnatEntriesResponse
func (client *Client) DescribeDnatEntriesWithContext(ctx context.Context, request *DescribeDnatEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDnatEntriesResponse, _err error) {
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

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDnatEntries"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDnatEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Smart Access Gateway (SAG) instances that are associated with a specified flow log.
//
// @param request - DescribeFlowLogSagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowLogSagsResponse
func (client *Client) DescribeFlowLogSagsWithContext(ctx context.Context, request *DescribeFlowLogSagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowLogSagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowLogSags"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowLogSagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries flow logs in a specified region.
//
// @param request - DescribeFlowLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowLogsResponse
func (client *Client) DescribeFlowLogsWithContext(ctx context.Context, request *DescribeFlowLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowLogsResponse, _err error) {
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

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.FlowLogName) {
		query["FlowLogName"] = request.FlowLogName
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowLogs"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowLogsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization information about a Cloud Connect Network (CCN) instance.
//
// @param request - DescribeGrantRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantRulesResponse
func (client *Client) DescribeGrantRulesWithContext(ctx context.Context, request *DescribeGrantRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociatedCcnId) {
		query["AssociatedCcnId"] = request.AssociatedCcnId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGrantRules"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGrantRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permission information about a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeGrantSagRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantSagRulesResponse
func (client *Client) DescribeGrantSagRulesWithContext(ctx context.Context, request *DescribeGrantSagRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantSagRulesResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGrantSagRules"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGrantSagRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries authorization information about Smart Access Gateway (SAG) instances and cross-account virtual border routers (VBRs).
//
// @param request - DescribeGrantSagVbrRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGrantSagVbrRulesResponse
func (client *Client) DescribeGrantSagVbrRulesWithContext(ctx context.Context, request *DescribeGrantSagVbrRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeGrantSagVbrRulesResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.VbrInstanceId) {
		query["VbrInstanceId"] = request.VbrInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeGrantSagVbrRules"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeGrantSagVbrRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about a health check instance.
//
// @param request - DescribeHealthCheckAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHealthCheckAttributeResponse
func (client *Client) DescribeHealthCheckAttributeWithContext(ctx context.Context, request *DescribeHealthCheckAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeHealthCheckAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HcInstanceId) {
		query["HcInstanceId"] = request.HcInstanceId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHealthCheckAttribute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHealthCheckAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries health checks that are associated with a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeHealthChecksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeHealthChecksResponse
func (client *Client) DescribeHealthChecksWithContext(ctx context.Context, request *DescribeHealthChecksRequest, runtime *dara.RuntimeOptions) (_result *DescribeHealthChecksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HcInstanceId) {
		query["HcInstanceId"] = request.HcInstanceId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeHealthChecks"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeHealthChecksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries traffic throttling rules of a quality of service (QoS) policy.
//
// @param request - DescribeQosCarsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQosCarsResponse
func (client *Client) DescribeQosCarsWithContext(ctx context.Context, request *DescribeQosCarsRequest, runtime *dara.RuntimeOptions) (_result *DescribeQosCarsResponse, _err error) {
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

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
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

	if !dara.IsNil(request.QosCarId) {
		query["QosCarId"] = request.QosCarId
	}

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
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
		Action:      dara.String("DescribeQosCars"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQosCarsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality of service (QoS) rules that contain 5-tuples.
//
// @param request - DescribeQosPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQosPoliciesResponse
func (client *Client) DescribeQosPoliciesWithContext(ctx context.Context, request *DescribeQosPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DescribeQosPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
	}

	if !dara.IsNil(request.QosPolicyId) {
		query["QosPolicyId"] = request.QosPolicyId
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
		Action:      dara.String("DescribeQosPolicies"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQosPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries quality of service (QoS) policies in a specified region.
//
// @param request - DescribeQosesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQosesResponse
func (client *Client) DescribeQosesWithContext(ctx context.Context, request *DescribeQosesRequest, runtime *dara.RuntimeOptions) (_result *DescribeQosesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QosIds) {
		query["QosIds"] = request.QosIds
	}

	if !dara.IsNil(request.QosName) {
		query["QosName"] = request.QosName
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
		Action:      dara.String("DescribeQoses"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQosesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries available regions.
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
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
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
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2018-03-13"),
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
// Queries route advertisement policies.
//
// @param request - DescribeRouteDistributionStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRouteDistributionStrategiesResponse
func (client *Client) DescribeRouteDistributionStrategiesWithContext(ctx context.Context, request *DescribeRouteDistributionStrategiesRequest, runtime *dara.RuntimeOptions) (_result *DescribeRouteDistributionStrategiesResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRouteDistributionStrategies"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRouteDistributionStrategiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSAGDeviceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSAGDeviceInfoResponse
func (client *Client) DescribeSAGDeviceInfoWithContext(ctx context.Context, request *DescribeSAGDeviceInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeSAGDeviceInfoResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSAGDeviceInfo"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSAGDeviceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the DNS servers used by a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSagCurrentDnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagCurrentDnsResponse
func (client *Client) DescribeSagCurrentDnsWithContext(ctx context.Context, request *DescribeSagCurrentDnsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagCurrentDnsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagCurrentDns"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagCurrentDnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the top 10 Smart Access Gateway (SAG) instances that have the highest packet loss rates in a specific region.
//
// @param request - DescribeSagDropTopNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagDropTopNResponse
func (client *Client) DescribeSagDropTopNWithContext(ctx context.Context, request *DescribeSagDropTopNRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagDropTopNResponse, _err error) {
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

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagDropTopN"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagDropTopNResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sub-interfaces added to an Express Connect circuit port.
//
// @param request - DescribeSagExpressConnectInterfaceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagExpressConnectInterfaceListResponse
func (client *Client) DescribeSagExpressConnectInterfaceListWithContext(ctx context.Context, request *DescribeSagExpressConnectInterfaceListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagExpressConnectInterfaceListResponse, _err error) {
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

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagExpressConnectInterfaceList"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagExpressConnectInterfaceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global routing protocol.
//
// @param request - DescribeSagGlobalRouteProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagGlobalRouteProtocolResponse
func (client *Client) DescribeSagGlobalRouteProtocolWithContext(ctx context.Context, request *DescribeSagGlobalRouteProtocolRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagGlobalRouteProtocolResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagGlobalRouteProtocol"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagGlobalRouteProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the high availability (HA) configuration of a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSagHaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagHaResponse
func (client *Client) DescribeSagHaWithContext(ctx context.Context, request *DescribeSagHaRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagHaResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagHa"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagHaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the LAN port settings of a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSagLanListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagLanListResponse
func (client *Client) DescribeSagLanListWithContext(ctx context.Context, request *DescribeSagLanListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagLanListResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagLanList"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagLanListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the management port settings of a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSagManagementPortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagManagementPortResponse
func (client *Client) DescribeSagManagementPortWithContext(ctx context.Context, request *DescribeSagManagementPortRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagManagementPortResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagManagementPort"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagManagementPortResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of clients connected to Alibaba Cloud through a Smart Access Gateway (SAG) app instance.
//
// @param request - DescribeSagOnlineClientStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagOnlineClientStatisticsResponse
func (client *Client) DescribeSagOnlineClientStatisticsWithContext(ctx context.Context, request *DescribeSagOnlineClientStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagOnlineClientStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGIds) {
		query["SmartAGIds"] = request.SmartAGIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagOnlineClientStatistics"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagOnlineClientStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the information of a physical port.
//
// @param request - DescribeSagPortListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagPortListResponse
func (client *Client) DescribeSagPortListWithContext(ctx context.Context, request *DescribeSagPortListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagPortListResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagPortList"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagPortListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ports for which the specified routing protocol is enabled.
//
// @param request - DescribeSagPortRouteProtocolListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagPortRouteProtocolListResponse
func (client *Client) DescribeSagPortRouteProtocolListWithContext(ctx context.Context, request *DescribeSagPortRouteProtocolListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagPortRouteProtocolListResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagPortRouteProtocolList"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagPortRouteProtocolListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries remote logon information about a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSagRemoteAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagRemoteAccessResponse
func (client *Client) DescribeSagRemoteAccessWithContext(ctx context.Context, request *DescribeSagRemoteAccessRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagRemoteAccessResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagRemoteAccess"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagRemoteAccessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routes of a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSagRouteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagRouteListResponse
func (client *Client) DescribeSagRouteListWithContext(ctx context.Context, request *DescribeSagRouteListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagRouteListResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagRouteList"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagRouteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of Border Gateway Protocol (BGP) dynamic routing.
//
// @param request - DescribeSagRouteProtocolBgpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagRouteProtocolBgpResponse
func (client *Client) DescribeSagRouteProtocolBgpWithContext(ctx context.Context, request *DescribeSagRouteProtocolBgpRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagRouteProtocolBgpResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagRouteProtocolBgp"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagRouteProtocolBgpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the OSPF settings of a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSagRouteProtocolOspfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagRouteProtocolOspfResponse
func (client *Client) DescribeSagRouteProtocolOspfWithContext(ctx context.Context, request *DescribeSagRouteProtocolOspfRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagRouteProtocolOspfResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagRouteProtocolOspf"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagRouteProtocolOspfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the static routes of a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSagStaticRouteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagStaticRouteListResponse
func (client *Client) DescribeSagStaticRouteListWithContext(ctx context.Context, request *DescribeSagStaticRouteListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagStaticRouteListResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagStaticRouteList"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagStaticRouteListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the top 10 Smart Access Gateway (SAG) instances that have the highest data transfer rates in a specific region.
//
// @param request - DescribeSagTrafficTopNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagTrafficTopNResponse
func (client *Client) DescribeSagTrafficTopNWithContext(ctx context.Context, request *DescribeSagTrafficTopNRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagTrafficTopNResponse, _err error) {
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

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagTrafficTopN"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagTrafficTopNResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the DNS servers used by a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSagUserDnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagUserDnsResponse
func (client *Client) DescribeSagUserDnsWithContext(ctx context.Context, request *DescribeSagUserDnsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagUserDnsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagUserDns"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagUserDnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a specified virtual border router (VBR) is associated with a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSagVbrRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagVbrRelationsResponse
func (client *Client) DescribeSagVbrRelationsWithContext(ctx context.Context, request *DescribeSagVbrRelationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagVbrRelationsResponse, _err error) {
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

	if !dara.IsNil(request.VbrInstanceIds) {
		query["VbrInstanceIds"] = request.VbrInstanceIds
	}

	if !dara.IsNil(request.VbrRegionId) {
		query["VbrRegionId"] = request.VbrRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagVbrRelations"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagVbrRelationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the 4G subscriber identity module (SIM) card used by the WAN port of a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSagWan4GRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagWan4GResponse
func (client *Client) DescribeSagWan4GWithContext(ctx context.Context, request *DescribeSagWan4GRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagWan4GResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagWan4G"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagWan4GResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the WAN port settings of a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSagWanListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagWanListResponse
func (client *Client) DescribeSagWanListWithContext(ctx context.Context, request *DescribeSagWanListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagWanListResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagWanList"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagWanListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SNAT settings of the WAN port of a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSagWanSnatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagWanSnatResponse
func (client *Client) DescribeSagWanSnatWithContext(ctx context.Context, request *DescribeSagWanSnatRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagWanSnatResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagWanSnat"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagWanSnatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the Wi-Fi settings of a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSagWifiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSagWifiResponse
func (client *Client) DescribeSagWifiWithContext(ctx context.Context, request *DescribeSagWifiRequest, runtime *dara.RuntimeOptions) (_result *DescribeSagWifiResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSagWifi"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSagWifiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSmartAccessGatewayAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmartAccessGatewayAttributeResponse
func (client *Client) DescribeSmartAccessGatewayAttributeWithContext(ctx context.Context, request *DescribeSmartAccessGatewayAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmartAccessGatewayAttributeResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmartAccessGatewayAttribute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmartAccessGatewayAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries client accounts added to a Smart Access Gateway (SAG) app instance.
//
// @param request - DescribeSmartAccessGatewayClientUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmartAccessGatewayClientUsersResponse
func (client *Client) DescribeSmartAccessGatewayClientUsersWithContext(ctx context.Context, request *DescribeSmartAccessGatewayClientUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmartAccessGatewayClientUsersResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserMail) {
		query["UserMail"] = request.UserMail
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmartAccessGatewayClientUsers"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmartAccessGatewayClientUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the high availability (HA) settings of a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSmartAccessGatewayHaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmartAccessGatewayHaResponse
func (client *Client) DescribeSmartAccessGatewayHaWithContext(ctx context.Context, request *DescribeSmartAccessGatewayHaRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmartAccessGatewayHaResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmartAccessGatewayHa"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmartAccessGatewayHaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the version of a Smart Access Gateway (SAG) device.
//
// @param request - DescribeSmartAccessGatewayVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmartAccessGatewayVersionsResponse
func (client *Client) DescribeSmartAccessGatewayVersionsWithContext(ctx context.Context, request *DescribeSmartAccessGatewayVersionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmartAccessGatewayVersionsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.VersionType) {
		query["VersionType"] = request.VersionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmartAccessGatewayVersions"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmartAccessGatewayVersionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Smart Access Gateway (SAG) instances.
//
// @param request - DescribeSmartAccessGatewaysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSmartAccessGatewaysResponse
func (client *Client) DescribeSmartAccessGatewaysWithContext(ctx context.Context, request *DescribeSmartAccessGatewaysRequest, runtime *dara.RuntimeOptions) (_result *DescribeSmartAccessGatewaysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclIds) {
		query["AclIds"] = request.AclIds
	}

	if !dara.IsNil(request.AssociatedCcnId) {
		query["AssociatedCcnId"] = request.AssociatedCcnId
	}

	if !dara.IsNil(request.AssociatedCcnName) {
		query["AssociatedCcnName"] = request.AssociatedCcnName
	}

	if !dara.IsNil(request.BusinessState) {
		query["BusinessState"] = request.BusinessState
	}

	if !dara.IsNil(request.CanAssociateQos) {
		query["CanAssociateQos"] = request.CanAssociateQos
	}

	if !dara.IsNil(request.HardwareType) {
		query["HardwareType"] = request.HardwareType
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGIds) {
		query["SmartAGIds"] = request.SmartAGIds
	}

	if !dara.IsNil(request.SoftwareVersion) {
		query["SoftwareVersion"] = request.SoftwareVersion
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UnboundAclIds) {
		query["UnboundAclIds"] = request.UnboundAclIds
	}

	if !dara.IsNil(request.VersionComparator) {
		query["VersionComparator"] = request.VersionComparator
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSmartAccessGateways"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSmartAccessGatewaysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query source network address translation (SNAT) entries associated with a Smart Access Gateway (SAG) instance.
//
// @param request - DescribeSnatEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnatEntriesResponse
func (client *Client) DescribeSnatEntriesWithContext(ctx context.Context, request *DescribeSnatEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnatEntriesResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnatEntries"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnatEntriesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Smart Access Gateway (SAG) instances that are not associated with a flow log in a specific region.
//
// @param request - DescribeUnbindFlowLogSagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUnbindFlowLogSagsResponse
func (client *Client) DescribeUnbindFlowLogSagsWithContext(ctx context.Context, request *DescribeUnbindFlowLogSagsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUnbindFlowLogSagsResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUnbindFlowLogSags"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUnbindFlowLogSagsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the amount of data transfer generated by each client account of a Smart Access Gateway (SAG) app instance.
//
// @param request - DescribeUserFlowStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserFlowStatisticsResponse
func (client *Client) DescribeUserFlowStatisticsWithContext(ctx context.Context, request *DescribeUserFlowStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserFlowStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.StatisticsDate) {
		query["StatisticsDate"] = request.StatisticsDate
	}

	if !dara.IsNil(request.UserNames) {
		query["UserNames"] = request.UserNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserFlowStatistics"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserFlowStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of clients that are connected to Alibaba Cloud through a specific Smart Access Gateway (SAG) app instance.
//
// @param request - DescribeUserOnlineClientStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserOnlineClientStatisticsResponse
func (client *Client) DescribeUserOnlineClientStatisticsWithContext(ctx context.Context, request *DescribeUserOnlineClientStatisticsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserOnlineClientStatisticsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserNames) {
		query["UserNames"] = request.UserNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserOnlineClientStatistics"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserOnlineClientStatisticsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the connection information about a client based on the ID of the Smart Access Gateway (SAG) APP instance and username of the client account.
//
// @param request - DescribeUserOnlineClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserOnlineClientsResponse
func (client *Client) DescribeUserOnlineClientsWithContext(ctx context.Context, request *DescribeUserOnlineClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserOnlineClientsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserOnlineClients"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserOnlineClientsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enables diagnostics for Smart Access Gateway (SAG) devices
//
// @param request - DiagnoseSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DiagnoseSmartAccessGatewayResponse
func (client *Client) DiagnoseSmartAccessGatewayWithContext(ctx context.Context, request *DiagnoseSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *DiagnoseSmartAccessGatewayResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DiagnoseSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DiagnoseSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables the deep packet inspection (DPI) feature for a Smart Access Gateway (SAG) instance.
//
// @param request - DisableSmartAGDpiMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSmartAGDpiMonitorResponse
func (client *Client) DisableSmartAGDpiMonitorWithContext(ctx context.Context, request *DisableSmartAGDpiMonitorRequest, runtime *dara.RuntimeOptions) (_result *DisableSmartAGDpiMonitorResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSmartAGDpiMonitor"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSmartAGDpiMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a client account of a Smart Access Gateway (SAG) app instance.
//
// @param request - DisableSmartAccessGatewayUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSmartAccessGatewayUserResponse
func (client *Client) DisableSmartAccessGatewayUserWithContext(ctx context.Context, request *DisableSmartAccessGatewayUserRequest, runtime *dara.RuntimeOptions) (_result *DisableSmartAccessGatewayUserResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSmartAccessGatewayUser"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSmartAccessGatewayUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an access control list (ACL) from a Smart Access Gateway (SAG) instance.
//
// @param request - DisassociateACLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateACLResponse
func (client *Client) DisassociateACLWithContext(ctx context.Context, request *DisassociateACLRequest, runtime *dara.RuntimeOptions) (_result *DisassociateACLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateACL"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateACLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a Smart Access Gateway (SAG) instance from a flow log.
//
// @param request - DisassociateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateFlowLogResponse
func (client *Client) DisassociateFlowLogWithContext(ctx context.Context, request *DisassociateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *DisassociateFlowLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateFlowLog"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateFlowLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a Smart Access Gateway (SAG) instance from a Quality of Service (QoS) policy.
//
// @param request - DisassociateQosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisassociateQosResponse
func (client *Client) DisassociateQosWithContext(ctx context.Context, request *DisassociateQosRequest, runtime *dara.RuntimeOptions) (_result *DisassociateQosResponse, _err error) {
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

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisassociateQos"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisassociateQosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a diagnosis report for a Smart Access Gateway (SAG) device.
//
// @param request - DiscribeSmartAccessGatewayDiagnosisReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DiscribeSmartAccessGatewayDiagnosisReportResponse
func (client *Client) DiscribeSmartAccessGatewayDiagnosisReportWithContext(ctx context.Context, request *DiscribeSmartAccessGatewayDiagnosisReportRequest, runtime *dara.RuntimeOptions) (_result *DiscribeSmartAccessGatewayDiagnosisReportResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DiscribeSmartAccessGatewayDiagnosisReport"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DiscribeSmartAccessGatewayDiagnosisReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a Smart Access Gateway (SAG) instance from a bandwidth plan for application acceleration.
//
// Description:
//
// When you call **DissociateSmartAGFromApplicationBandwidthPackage**, you can set the **SmartAGId*	- parameter to specify an SAG instance, or set the **SmartAGIdList*	- to specify multiple SAG instances.
//
// @param request - DissociateSmartAGFromApplicationBandwidthPackageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DissociateSmartAGFromApplicationBandwidthPackageResponse
func (client *Client) DissociateSmartAGFromApplicationBandwidthPackageWithContext(ctx context.Context, request *DissociateSmartAGFromApplicationBandwidthPackageRequest, runtime *dara.RuntimeOptions) (_result *DissociateSmartAGFromApplicationBandwidthPackageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationBandwidthPackageId) {
		query["ApplicationBandwidthPackageId"] = request.ApplicationBandwidthPackageId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGIdList) {
		query["SmartAGIdList"] = request.SmartAGIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DissociateSmartAGFromApplicationBandwidthPackage"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DissociateSmartAGFromApplicationBandwidthPackageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downgrades the bandwidth of a Smart Access Gateway (SAG) instance.
//
// @param request - DowngradeSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DowngradeSmartAccessGatewayResponse
func (client *Client) DowngradeSmartAccessGatewayWithContext(ctx context.Context, request *DowngradeSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *DowngradeSmartAccessGatewayResponse, _err error) {
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

	if !dara.IsNil(request.BandWidthSpec) {
		query["BandWidthSpec"] = request.BandWidthSpec
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DowngradeSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DowngradeSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Decreases the quota of client accounts that can be connected to a Smart Access Gateway (SAG) app instance.
//
// @param request - DowngradeSmartAccessGatewaySoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DowngradeSmartAccessGatewaySoftwareResponse
func (client *Client) DowngradeSmartAccessGatewaySoftwareWithContext(ctx context.Context, request *DowngradeSmartAccessGatewaySoftwareRequest, runtime *dara.RuntimeOptions) (_result *DowngradeSmartAccessGatewaySoftwareResponse, _err error) {
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

	if !dara.IsNil(request.DataPlan) {
		query["DataPlan"] = request.DataPlan
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserCount) {
		query["UserCount"] = request.UserCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DowngradeSmartAccessGatewaySoftware"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DowngradeSmartAccessGatewaySoftwareResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables monitoring based on the deep packet inspection (DPI) feature for a Smart Access Gateway (SAG) instance.
//
// Description:
//
//	  The DPI feature is enabled for the SAG instance. For more information, see [UpdateSmartAGDpiAttribute](https://help.aliyun.com/document_detail/196146.html).
//
//		- Log Service is activated. For more information, see [Quick Start](https://help.aliyun.com/document_detail/54604.html).
//
// @param request - EnableSmartAGDpiMonitorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSmartAGDpiMonitorResponse
func (client *Client) EnableSmartAGDpiMonitorWithContext(ctx context.Context, request *EnableSmartAGDpiMonitorRequest, runtime *dara.RuntimeOptions) (_result *EnableSmartAGDpiMonitorResponse, _err error) {
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

	if !dara.IsNil(request.SlsLogStore) {
		query["SlsLogStore"] = request.SlsLogStore
	}

	if !dara.IsNil(request.SlsProjectName) {
		query["SlsProjectName"] = request.SlsProjectName
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSmartAGDpiMonitor"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSmartAGDpiMonitorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates a client account of a Smart Access Gateway (SAG) app instance.
//
// @param request - EnableSmartAccessGatewayUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSmartAccessGatewayUserResponse
func (client *Client) EnableSmartAccessGatewayUserWithContext(ctx context.Context, request *EnableSmartAccessGatewayUserRequest, runtime *dara.RuntimeOptions) (_result *EnableSmartAccessGatewayUserResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSmartAccessGatewayUser"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSmartAccessGatewayUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration errors of the deep packet inspection (DPI) feature in an access control list (ACL).
//
// Description:
//
//	  An application-aware ACL is created. For more information, see [AddACLRule](https://help.aliyun.com/document_detail/114012.html).
//
//		- The application-aware ACL is applied to a Smart Access Gateway (SAG) instance. For more information, see [AssociateACL](https://help.aliyun.com/document_detail/114009.html).
//
// @param request - GetAclAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAclAttributeResponse
func (client *Client) GetAclAttributeWithContext(ctx context.Context, request *GetAclAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetAclAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAclAttribute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAclAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the deep packet inspection (DPI) feature of a Smart Access Gateway (SAG) instance.
//
// @param request - GetAdvancedMonitorStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdvancedMonitorStateResponse
func (client *Client) GetAdvancedMonitorStateWithContext(ctx context.Context, request *GetAdvancedMonitorStateRequest, runtime *dara.RuntimeOptions) (_result *GetAdvancedMonitorStateResponse, _err error) {
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

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdvancedMonitorState"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdvancedMonitorStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of Cloud Connect Network (CCN) instances that can be created within the current account in a region.
//
// @param request - GetCloudConnectNetworkUseLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCloudConnectNetworkUseLimitResponse
func (client *Client) GetCloudConnectNetworkUseLimitWithContext(ctx context.Context, request *GetCloudConnectNetworkUseLimitRequest, runtime *dara.RuntimeOptions) (_result *GetCloudConnectNetworkUseLimitResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCloudConnectNetworkUseLimit"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCloudConnectNetworkUseLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a quality of service (QoS) policy.
//
// @param request - GetQosAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQosAttributeResponse
func (client *Client) GetQosAttributeWithContext(ctx context.Context, request *GetQosAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetQosAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQosAttribute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQosAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of the deep packet inspection (DPI) feature for a Smart Access Gateway (SAG) instance.
//
// @param request - GetSmartAGDpiAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmartAGDpiAttributeResponse
func (client *Client) GetSmartAGDpiAttributeWithContext(ctx context.Context, request *GetSmartAGDpiAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetSmartAGDpiAttributeResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmartAGDpiAttribute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmartAGDpiAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of Smart Access Gateway (SAG) instances that you can purchase.
//
// @param request - GetSmartAccessGatewayUseLimitRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmartAccessGatewayUseLimitResponse
func (client *Client) GetSmartAccessGatewayUseLimitWithContext(ctx context.Context, request *GetSmartAccessGatewayUseLimitRequest, runtime *dara.RuntimeOptions) (_result *GetSmartAccessGatewayUseLimitResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmartAccessGatewayUseLimit"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmartAccessGatewayUseLimitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a Cloud Enterprise Network (CEN) instance to communicate with a Cloud Connect Network (CCN) instance that belongs to another Alibaba Cloud account.
//
// @param request - GrantInstanceToCbnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantInstanceToCbnResponse
func (client *Client) GrantInstanceToCbnWithContext(ctx context.Context, request *GrantInstanceToCbnRequest, runtime *dara.RuntimeOptions) (_result *GrantInstanceToCbnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnInstanceId) {
		query["CcnInstanceId"] = request.CcnInstanceId
	}

	if !dara.IsNil(request.CenInstanceId) {
		query["CenInstanceId"] = request.CenInstanceId
	}

	if !dara.IsNil(request.CenUid) {
		query["CenUid"] = request.CenUid
	}

	if !dara.IsNil(request.GrantTrafficService) {
		query["GrantTrafficService"] = request.GrantTrafficService
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
		Action:      dara.String("GrantInstanceToCbn"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantInstanceToCbnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a Smart Access Gateway (SAG) instance to communicate with a Cloud Connect Network (CCN) instance that belongs to another Alibaba Cloud account.
//
// @param request - GrantSagInstanceToCcnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantSagInstanceToCcnResponse
func (client *Client) GrantSagInstanceToCcnWithContext(ctx context.Context, request *GrantSagInstanceToCcnRequest, runtime *dara.RuntimeOptions) (_result *GrantSagInstanceToCcnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnInstanceId) {
		query["CcnInstanceId"] = request.CcnInstanceId
	}

	if !dara.IsNil(request.CcnUid) {
		query["CcnUid"] = request.CcnUid
	}

	if !dara.IsNil(request.GrantTrafficService) {
		query["GrantTrafficService"] = request.GrantTrafficService
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantSagInstanceToCcn"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantSagInstanceToCcnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Acquires permissions from a virtual border router (VBR) under another account.
//
// @param request - GrantSagInstanceToVbrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantSagInstanceToVbrResponse
func (client *Client) GrantSagInstanceToVbrWithContext(ctx context.Context, request *GrantSagInstanceToVbrRequest, runtime *dara.RuntimeOptions) (_result *GrantSagInstanceToVbrResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.VbrInstanceId) {
		query["VbrInstanceId"] = request.VbrInstanceId
	}

	if !dara.IsNil(request.VbrRegionId) {
		query["VbrRegionId"] = request.VbrRegionId
	}

	if !dara.IsNil(request.VbrUid) {
		query["VbrUid"] = request.VbrUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantSagInstanceToVbr"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantSagInstanceToVbrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Closes a connection based on the Smart Access Gateway (SAG) app instance ID and username.
//
// @param request - KickOutClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return KickOutClientsResponse
func (client *Client) KickOutClientsWithContext(ctx context.Context, request *KickOutClientsRequest, runtime *dara.RuntimeOptions) (_result *KickOutClientsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("KickOutClients"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &KickOutClientsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries network qualities of endpoints for Smart Access Gateway (SAG) instances.
//
// @param request - ListAccessPointNetworkQualitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessPointNetworkQualitiesResponse
func (client *Client) ListAccessPointNetworkQualitiesWithContext(ctx context.Context, request *ListAccessPointNetworkQualitiesRequest, runtime *dara.RuntimeOptions) (_result *ListAccessPointNetworkQualitiesResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessPointNetworkQualities"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessPointNetworkQualitiesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access points in a region.
//
// Description:
//
// An access point connects Smart Access Gateway (SAG) instances to Cloud Connect Network (CCN). Some regions may contain more than one access point. After an SAG instance is associated with a CCN instance, the system selects the nearest access point to establish connections to Alibaba Cloud. This topic describes how to query access points in a specific region. For more information about the areas that support CCN, see [Introduction to CCN](https://help.aliyun.com/document_detail/93667.html).
//
// @param request - ListAccessPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessPointsResponse
func (client *Client) ListAccessPointsWithContext(ctx context.Context, request *ListAccessPointsRequest, runtime *dara.RuntimeOptions) (_result *ListAccessPointsResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessPoints"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the service addresses of a Smart Access Gateway (SAG) device.
//
// @param request - ListAvailableServiceAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableServiceAddressResponse
func (client *Client) ListAvailableServiceAddressWithContext(ctx context.Context, request *ListAvailableServiceAddressRequest, runtime *dara.RuntimeOptions) (_result *ListAvailableServiceAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableServiceAddress"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableServiceAddressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries configuration errors of the deep packet inspection (DPI) feature.
//
// Description:
//
// ## Background information
//
// If you have configured an application-aware access control list (ACL) or a quality of service (QoS) policy and associated it with a Smart Access Gateway (SAG) instance, you can call this operation to query whether the ACL rules or 5-tuples in the QoS policy are applied to the SAG instance. If settings are not applied to the SAG instance, the error information is returned.
//
// @param request - ListDpiConfigErrorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDpiConfigErrorResponse
func (client *Client) ListDpiConfigErrorWithContext(ctx context.Context, request *ListDpiConfigErrorRequest, runtime *dara.RuntimeOptions) (_result *ListDpiConfigErrorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DpiConfigType) {
		query["DpiConfigType"] = request.DpiConfigType
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

	if !dara.IsNil(request.RuleInstanceId) {
		query["RuleInstanceId"] = request.RuleInstanceId
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDpiConfigError"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDpiConfigErrorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about application groups supported by Smart Access Gateway (SAG) instances in a specified region.
//
// @param request - ListDpiGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDpiGroupsResponse
func (client *Client) ListDpiGroupsWithContext(ctx context.Context, request *ListDpiGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListDpiGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DpiGroupIds) {
		query["DpiGroupIds"] = request.DpiGroupIds
	}

	if !dara.IsNil(request.DpiGroupNames) {
		query["DpiGroupNames"] = request.DpiGroupNames
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
		Action:      dara.String("ListDpiGroups"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDpiGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an application or an application group in a region, or about the applications supported by Smart Access Gateway (SAG) in a region.
//
// Description:
//
// This operation supports the following features:
//
//   - Queries the information about all applications supported by the SAG instance in a specified region.
//
//   - Queries the information about an application by application ID in a specified region.
//
//   - Queries the information about an application by application name in a specified region.
//
//   - Queries the information about an application group by group ID in a specified region.
//
// If this is the first time you call this operation, we recommend that you query all applications supported by the SAG instance in the specified region by region ID. Then, you can query the information about a specified application.
//
// @param request - ListDpiSignaturesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDpiSignaturesResponse
func (client *Client) ListDpiSignaturesWithContext(ctx context.Context, request *ListDpiSignaturesRequest, runtime *dara.RuntimeOptions) (_result *ListDpiSignaturesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DpiGroupId) {
		query["DpiGroupId"] = request.DpiGroupId
	}

	if !dara.IsNil(request.DpiSignatureIds) {
		query["DpiSignatureIds"] = request.DpiSignatureIds
	}

	if !dara.IsNil(request.DpiSignatureNames) {
		query["DpiSignatureNames"] = request.DpiSignatureNames
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
		Action:      dara.String("ListDpiSignatures"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDpiSignaturesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries enterprise codes.
//
// @param request - ListEnterpriseCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseCodeResponse
func (client *Client) ListEnterpriseCodeWithContext(ctx context.Context, request *ListEnterpriseCodeRequest, runtime *dara.RuntimeOptions) (_result *ListEnterpriseCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseCode) {
		query["EnterpriseCode"] = request.EnterpriseCode
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
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
		Action:      dara.String("ListEnterpriseCode"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a probe task.
//
// @param request - ListProbeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProbeTaskResponse
func (client *Client) ListProbeTaskWithContext(ctx context.Context, request *ListProbeTaskRequest, runtime *dara.RuntimeOptions) (_result *ListProbeTaskResponse, _err error) {
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

	if !dara.IsNil(request.ProbeTaskId) {
		query["ProbeTaskId"] = request.ProbeTaskId
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.SagName) {
		query["SagName"] = request.SagName
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProbeTask"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProbeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries features that are not supported by a specified Smart Access Gateway (SAG) device.
//
// @param request - ListSmartAGApiUnsupportedFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSmartAGApiUnsupportedFeatureResponse
func (client *Client) ListSmartAGApiUnsupportedFeatureWithContext(ctx context.Context, request *ListSmartAGApiUnsupportedFeatureRequest, runtime *dara.RuntimeOptions) (_result *ListSmartAGApiUnsupportedFeatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OpenApiName) {
		query["OpenApiName"] = request.OpenApiName
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSmartAGApiUnsupportedFeature"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSmartAGApiUnsupportedFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about Smart Access Gateway (SAG) instances within specific access points in a specific region.
//
// @param request - ListSmartAGByAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSmartAGByAccessPointResponse
func (client *Client) ListSmartAGByAccessPointWithContext(ctx context.Context, request *ListSmartAGByAccessPointRequest, runtime *dara.RuntimeOptions) (_result *ListSmartAGByAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
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

	if !dara.IsNil(request.SmartAGStatus) {
		query["SmartAGStatus"] = request.SmartAGStatus
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSmartAGByAccessPoint"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSmartAGByAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of an access control list (ACL).
//
// @param request - ModifyACLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyACLResponse
func (client *Client) ModifyACLWithContext(ctx context.Context, request *ModifyACLRequest, runtime *dara.RuntimeOptions) (_result *ModifyACLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
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
		Action:      dara.String("ModifyACL"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyACLResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an access control list (ACL) rule.
//
// @param request - ModifyACLRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyACLRuleResponse
func (client *Client) ModifyACLRuleWithContext(ctx context.Context, request *ModifyACLRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyACLRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclId) {
		query["AclId"] = request.AclId
	}

	if !dara.IsNil(request.AcrId) {
		query["AcrId"] = request.AcrId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DestCidr) {
		query["DestCidr"] = request.DestCidr
	}

	if !dara.IsNil(request.DestPortRange) {
		query["DestPortRange"] = request.DestPortRange
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.DpiGroupIds) {
		query["DpiGroupIds"] = request.DpiGroupIds
	}

	if !dara.IsNil(request.DpiSignatureIds) {
		query["DpiSignatureIds"] = request.DpiSignatureIds
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
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

	if !dara.IsNil(request.Policy) {
		query["Policy"] = request.Policy
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
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

	if !dara.IsNil(request.SourceCidr) {
		query["SourceCidr"] = request.SourceCidr
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyACLRule"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyACLRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the DNS settings of a Smart Access Gateway (SAG) app instance.
//
// @param request - ModifyClientUserDNSRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyClientUserDNSResponse
func (client *Client) ModifyClientUserDNSWithContext(ctx context.Context, request *ModifyClientUserDNSRequest, runtime *dara.RuntimeOptions) (_result *ModifyClientUserDNSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppDNS) {
		query["AppDNS"] = request.AppDNS
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecoveredDNS) {
		query["RecoveredDNS"] = request.RecoveredDNS
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyClientUserDNS"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyClientUserDNSResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a Cloud Connect Network (CCN) instance.
//
// @param request - ModifyCloudConnectNetworkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCloudConnectNetworkResponse
func (client *Client) ModifyCloudConnectNetworkWithContext(ctx context.Context, request *ModifyCloudConnectNetworkRequest, runtime *dara.RuntimeOptions) (_result *ModifyCloudConnectNetworkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnId) {
		query["CcnId"] = request.CcnId
	}

	if !dara.IsNil(request.CidrBlock) {
		query["CidrBlock"] = request.CidrBlock
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InterworkingStatus) {
		query["InterworkingStatus"] = request.InterworkingStatus
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
		Action:      dara.String("ModifyCloudConnectNetwork"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCloudConnectNetworkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the automatic upgrade policy of a Smart Access Gateway (SAG) device.
//
// @param request - ModifyDeviceAutoUpgradePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDeviceAutoUpgradePolicyResponse
func (client *Client) ModifyDeviceAutoUpgradePolicyWithContext(ctx context.Context, request *ModifyDeviceAutoUpgradePolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDeviceAutoUpgradePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CronExpression) {
		query["CronExpression"] = request.CronExpression
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.TimeZone) {
		query["TimeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.UpgradeType) {
		query["UpgradeType"] = request.UpgradeType
	}

	if !dara.IsNil(request.VersionType) {
		query["VersionType"] = request.VersionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDeviceAutoUpgradePolicy"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDeviceAutoUpgradePolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of a flow log.
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
	if !dara.IsNil(request.ActiveAging) {
		query["ActiveAging"] = request.ActiveAging
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.InactiveAging) {
		query["InactiveAging"] = request.InactiveAging
	}

	if !dara.IsNil(request.LogstoreName) {
		query["LogstoreName"] = request.LogstoreName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NetflowServerIp) {
		query["NetflowServerIp"] = request.NetflowServerIp
	}

	if !dara.IsNil(request.NetflowServerPort) {
		query["NetflowServerPort"] = request.NetflowServerPort
	}

	if !dara.IsNil(request.NetflowVersion) {
		query["NetflowVersion"] = request.NetflowVersion
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
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

	if !dara.IsNil(request.SlsRegionId) {
		query["SlsRegionId"] = request.SlsRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFlowLogAttribute"),
		Version:     dara.String("2018-03-13"),
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
// Modifies a health check.
//
// @param request - ModifyHealthCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHealthCheckResponse
func (client *Client) ModifyHealthCheckWithContext(ctx context.Context, request *ModifyHealthCheckRequest, runtime *dara.RuntimeOptions) (_result *ModifyHealthCheckResponse, _err error) {
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

	if !dara.IsNil(request.DstIpAddr) {
		query["DstIpAddr"] = request.DstIpAddr
	}

	if !dara.IsNil(request.DstPort) {
		query["DstPort"] = request.DstPort
	}

	if !dara.IsNil(request.FailCountThreshold) {
		query["FailCountThreshold"] = request.FailCountThreshold
	}

	if !dara.IsNil(request.HcInstanceId) {
		query["HcInstanceId"] = request.HcInstanceId
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

	if !dara.IsNil(request.ProbeCount) {
		query["ProbeCount"] = request.ProbeCount
	}

	if !dara.IsNil(request.ProbeInterval) {
		query["ProbeInterval"] = request.ProbeInterval
	}

	if !dara.IsNil(request.ProbeTimeout) {
		query["ProbeTimeout"] = request.ProbeTimeout
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

	if !dara.IsNil(request.RttFailThreshold) {
		query["RttFailThreshold"] = request.RttFailThreshold
	}

	if !dara.IsNil(request.RttThreshold) {
		query["RttThreshold"] = request.RttThreshold
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SrcIpAddr) {
		query["SrcIpAddr"] = request.SrcIpAddr
	}

	if !dara.IsNil(request.SrcPort) {
		query["SrcPort"] = request.SrcPort
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHealthCheck"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHealthCheckResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to modify a quality of service (QoS) policy, for example, its name.
//
// @param request - ModifyQosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyQosResponse
func (client *Client) ModifyQosWithContext(ctx context.Context, request *ModifyQosRequest, runtime *dara.RuntimeOptions) (_result *ModifyQosResponse, _err error) {
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

	if !dara.IsNil(request.QosDescription) {
		query["QosDescription"] = request.QosDescription
	}

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
	}

	if !dara.IsNil(request.QosName) {
		query["QosName"] = request.QosName
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
		Action:      dara.String("ModifyQos"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyQosResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a traffic throttling rule in a quality of service (QoS) policy.
//
// @param request - ModifyQosCarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyQosCarResponse
func (client *Client) ModifyQosCarWithContext(ctx context.Context, request *ModifyQosCarRequest, runtime *dara.RuntimeOptions) (_result *ModifyQosCarResponse, _err error) {
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

	if !dara.IsNil(request.LimitType) {
		query["LimitType"] = request.LimitType
	}

	if !dara.IsNil(request.MaxBandwidthAbs) {
		query["MaxBandwidthAbs"] = request.MaxBandwidthAbs
	}

	if !dara.IsNil(request.MaxBandwidthPercent) {
		query["MaxBandwidthPercent"] = request.MaxBandwidthPercent
	}

	if !dara.IsNil(request.MinBandwidthAbs) {
		query["MinBandwidthAbs"] = request.MinBandwidthAbs
	}

	if !dara.IsNil(request.MinBandwidthPercent) {
		query["MinBandwidthPercent"] = request.MinBandwidthPercent
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

	if !dara.IsNil(request.PercentSourceType) {
		query["PercentSourceType"] = request.PercentSourceType
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.QosCarId) {
		query["QosCarId"] = request.QosCarId
	}

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
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
		Action:      dara.String("ModifyQosCar"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyQosCarResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a traffic classification rule of a quality of service (QoS) policy.
//
// @param request - ModifyQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyQosPolicyResponse
func (client *Client) ModifyQosPolicyWithContext(ctx context.Context, request *ModifyQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyQosPolicyResponse, _err error) {
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

	if !dara.IsNil(request.DestCidr) {
		query["DestCidr"] = request.DestCidr
	}

	if !dara.IsNil(request.DestPortRange) {
		query["DestPortRange"] = request.DestPortRange
	}

	if !dara.IsNil(request.DpiGroupIds) {
		query["DpiGroupIds"] = request.DpiGroupIds
	}

	if !dara.IsNil(request.DpiSignatureIds) {
		query["DpiSignatureIds"] = request.DpiSignatureIds
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IpProtocol) {
		query["IpProtocol"] = request.IpProtocol
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

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.QosId) {
		query["QosId"] = request.QosId
	}

	if !dara.IsNil(request.QosPolicyId) {
		query["QosPolicyId"] = request.QosPolicyId
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

	if !dara.IsNil(request.SourceCidr) {
		query["SourceCidr"] = request.SourceCidr
	}

	if !dara.IsNil(request.SourcePortRange) {
		query["SourcePortRange"] = request.SourcePortRange
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyQosPolicy"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyQosPolicyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the route advertisement policy for a CIDR block.
//
// @param request - ModifyRouteDistributionStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyRouteDistributionStrategyResponse
func (client *Client) ModifyRouteDistributionStrategyWithContext(ctx context.Context, request *ModifyRouteDistributionStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyRouteDistributionStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestCidrBlock) {
		query["DestCidrBlock"] = request.DestCidrBlock
	}

	if !dara.IsNil(request.HcInstanceId) {
		query["HcInstanceId"] = request.HcInstanceId
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

	if !dara.IsNil(request.RouteDistribution) {
		query["RouteDistribution"] = request.RouteDistribution
	}

	if !dara.IsNil(request.RouteSource) {
		query["RouteSource"] = request.RouteSource
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyRouteDistributionStrategy"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyRouteDistributionStrategyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the password that is used to log on to a smart access gateway (SAG) device.
//
// @param request - ModifySAGAdminPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySAGAdminPasswordResponse
func (client *Client) ModifySAGAdminPasswordWithContext(ctx context.Context, request *ModifySAGAdminPasswordRequest, runtime *dara.RuntimeOptions) (_result *ModifySAGAdminPasswordResponse, _err error) {
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

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySAGAdminPassword"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySAGAdminPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the sub-interface information of an Express Connect circuit port.
//
// @param request - ModifySagExpressConnectInterfaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagExpressConnectInterfaceResponse
func (client *Client) ModifySagExpressConnectInterfaceWithContext(ctx context.Context, request *ModifySagExpressConnectInterfaceRequest, runtime *dara.RuntimeOptions) (_result *ModifySagExpressConnectInterfaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Vlan) {
		query["Vlan"] = request.Vlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagExpressConnectInterface"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagExpressConnectInterfaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the routing protocol of a Smart Access Gateway (SAG) instance.
//
// @param request - ModifySagGlobalRouteProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagGlobalRouteProtocolResponse
func (client *Client) ModifySagGlobalRouteProtocolWithContext(ctx context.Context, request *ModifySagGlobalRouteProtocolRequest, runtime *dara.RuntimeOptions) (_result *ModifySagGlobalRouteProtocolResponse, _err error) {
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

	if !dara.IsNil(request.RouteProtocol) {
		query["RouteProtocol"] = request.RouteProtocol
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagGlobalRouteProtocol"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagGlobalRouteProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the high availability (HA) settings of a Smart Access Gateway (SAG) device.
//
// @param request - ModifySagHaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagHaResponse
func (client *Client) ModifySagHaWithContext(ctx context.Context, request *ModifySagHaRequest, runtime *dara.RuntimeOptions) (_result *ModifySagHaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.VirtualIp) {
		query["VirtualIp"] = request.VirtualIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagHa"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagHaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to modify the LAN port configurations of a Smart Access Gateway (SAG) device.
//
// @param request - ModifySagLanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagLanResponse
func (client *Client) ModifySagLanWithContext(ctx context.Context, request *ModifySagLanRequest, runtime *dara.RuntimeOptions) (_result *ModifySagLanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndIp) {
		query["EndIp"] = request.EndIp
	}

	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	if !dara.IsNil(request.IPType) {
		query["IPType"] = request.IPType
	}

	if !dara.IsNil(request.Lease) {
		query["Lease"] = request.Lease
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.StartIp) {
		query["StartIp"] = request.StartIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagLan"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagLanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the settings of a Smart Access Gateway (SAG) device port.
//
// @param request - ModifySagManagementPortRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagManagementPortResponse
func (client *Client) ModifySagManagementPortWithContext(ctx context.Context, request *ModifySagManagementPortRequest, runtime *dara.RuntimeOptions) (_result *ModifySagManagementPortResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Gateway) {
		query["Gateway"] = request.Gateway
	}

	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagManagementPort"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagManagementPortResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the role of a port.
//
// Description:
//
// >  If you modify the role of a port, the current role configurations of the port are deleted.
//
// @param request - ModifySagPortRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagPortRoleResponse
func (client *Client) ModifySagPortRoleWithContext(ctx context.Context, request *ModifySagPortRoleRequest, runtime *dara.RuntimeOptions) (_result *ModifySagPortRoleResponse, _err error) {
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

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagPortRole"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagPortRoleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to modify the routing protocol of a port.
//
// @param request - ModifySagPortRouteProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagPortRouteProtocolResponse
func (client *Client) ModifySagPortRouteProtocolWithContext(ctx context.Context, request *ModifySagPortRouteProtocolRequest, runtime *dara.RuntimeOptions) (_result *ModifySagPortRouteProtocolResponse, _err error) {
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

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoteAs) {
		query["RemoteAs"] = request.RemoteAs
	}

	if !dara.IsNil(request.RemoteIp) {
		query["RemoteIp"] = request.RemoteIp
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteProtocol) {
		query["RouteProtocol"] = request.RouteProtocol
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Vlan) {
		query["Vlan"] = request.Vlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagPortRouteProtocol"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagPortRouteProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the remote access IP address of a Smart Access Gateway (SAG) device.
//
// @param request - ModifySagRemoteAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagRemoteAccessResponse
func (client *Client) ModifySagRemoteAccessWithContext(ctx context.Context, request *ModifySagRemoteAccessRequest, runtime *dara.RuntimeOptions) (_result *ModifySagRemoteAccessResponse, _err error) {
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

	if !dara.IsNil(request.RemoteAccessIp) {
		query["RemoteAccessIp"] = request.RemoteAccessIp
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagRemoteAccess"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagRemoteAccessResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Border Gateway Protocol (BGP) settings of a Smart Access Gateway (SAG) device.
//
// @param request - ModifySagRouteProtocolBgpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagRouteProtocolBgpResponse
func (client *Client) ModifySagRouteProtocolBgpWithContext(ctx context.Context, request *ModifySagRouteProtocolBgpRequest, runtime *dara.RuntimeOptions) (_result *ModifySagRouteProtocolBgpResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HoldTime) {
		query["HoldTime"] = request.HoldTime
	}

	if !dara.IsNil(request.KeepAlive) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !dara.IsNil(request.LocalAs) {
		query["LocalAs"] = request.LocalAs
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

	if !dara.IsNil(request.RouterId) {
		query["RouterId"] = request.RouterId
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagRouteProtocolBgp"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagRouteProtocolBgpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Open Shortest Path First (OSPF) settings of a Smart Access Gateway (SAG) instance.
//
// @param request - ModifySagRouteProtocolOspfRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagRouteProtocolOspfResponse
func (client *Client) ModifySagRouteProtocolOspfWithContext(ctx context.Context, request *ModifySagRouteProtocolOspfRequest, runtime *dara.RuntimeOptions) (_result *ModifySagRouteProtocolOspfResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AreaId) {
		query["AreaId"] = request.AreaId
	}

	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.DeadTime) {
		query["DeadTime"] = request.DeadTime
	}

	if !dara.IsNil(request.HelloTime) {
		query["HelloTime"] = request.HelloTime
	}

	if !dara.IsNil(request.Md5Key) {
		query["Md5Key"] = request.Md5Key
	}

	if !dara.IsNil(request.Md5KeyId) {
		query["Md5KeyId"] = request.Md5KeyId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
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

	if !dara.IsNil(request.RouterId) {
		query["RouterId"] = request.RouterId
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagRouteProtocolOspf"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagRouteProtocolOspfResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a static route.
//
// @param request - ModifySagStaticRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagStaticRouteResponse
func (client *Client) ModifySagStaticRouteWithContext(ctx context.Context, request *ModifySagStaticRouteRequest, runtime *dara.RuntimeOptions) (_result *ModifySagStaticRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationCidr) {
		query["DestinationCidr"] = request.DestinationCidr
	}

	if !dara.IsNil(request.NextHop) {
		query["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Vlan) {
		query["Vlan"] = request.Vlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagStaticRoute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagStaticRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the DNS servers used by a Smart Access Gateway (SAG) instance.
//
// @param request - ModifySagUserDnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagUserDnsResponse
func (client *Client) ModifySagUserDnsWithContext(ctx context.Context, request *ModifySagUserDnsRequest, runtime *dara.RuntimeOptions) (_result *ModifySagUserDnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MasterDns) {
		query["MasterDns"] = request.MasterDns
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

	if !dara.IsNil(request.SlaveDns) {
		query["SlaveDns"] = request.SlaveDns
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagUserDns"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagUserDnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the WAN port configurations of a Smart Access Gateway (SAG) device.
//
// Description:
//
// We recommend that you understand the functionality of a WAN port before you modify its settings. For more information, see [Configure a WAN port](https://help.aliyun.com/document_detail/163955.html).
//
// @param request - ModifySagWanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagWanResponse
func (client *Client) ModifySagWanWithContext(ctx context.Context, request *ModifySagWanRequest, runtime *dara.RuntimeOptions) (_result *ModifySagWanResponse, _err error) {
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

	if !dara.IsNil(request.Gateway) {
		query["Gateway"] = request.Gateway
	}

	if !dara.IsNil(request.IP) {
		query["IP"] = request.IP
	}

	if !dara.IsNil(request.IPType) {
		query["IPType"] = request.IPType
	}

	if !dara.IsNil(request.ISP) {
		query["ISP"] = request.ISP
	}

	if !dara.IsNil(request.Mask) {
		query["Mask"] = request.Mask
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagWan"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagWanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to modify the SNAT configurations of a WAN port of a Smart Access Gateway (SAG) device.
//
// @param request - ModifySagWanSnatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagWanSnatResponse
func (client *Client) ModifySagWanSnatWithContext(ctx context.Context, request *ModifySagWanSnatRequest, runtime *dara.RuntimeOptions) (_result *ModifySagWanSnatResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	if !dara.IsNil(request.Snat) {
		query["Snat"] = request.Snat
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagWanSnat"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagWanSnatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to modify the Wi-Fi settings of a Smart Access Gateway (SAG) device.
//
// @param request - ModifySagWifiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySagWifiResponse
func (client *Client) ModifySagWifiWithContext(ctx context.Context, request *ModifySagWifiRequest, runtime *dara.RuntimeOptions) (_result *ModifySagWifiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.Bandwidth) {
		query["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.EncryptAlgorithm) {
		query["EncryptAlgorithm"] = request.EncryptAlgorithm
	}

	if !dara.IsNil(request.IsAuth) {
		query["IsAuth"] = request.IsAuth
	}

	if !dara.IsNil(request.IsBroadcast) {
		query["IsBroadcast"] = request.IsBroadcast
	}

	if !dara.IsNil(request.IsEnable) {
		query["IsEnable"] = request.IsEnable
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
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

	if !dara.IsNil(request.SSID) {
		query["SSID"] = request.SSID
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySagWifi"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySagWifiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a Smart Access Gateway (SAG) instance.
//
// @param request - ModifySmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmartAccessGatewayResponse
func (client *Client) ModifySmartAccessGatewayWithContext(ctx context.Context, request *ModifySmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *ModifySmartAccessGatewayResponse, _err error) {
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

	if !dara.IsNil(request.EnableSoftwareConnectionAudit) {
		query["EnableSoftwareConnectionAudit"] = request.EnableSoftwareConnectionAudit
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

	if !dara.IsNil(request.Position) {
		query["Position"] = request.Position
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

	if !dara.IsNil(request.RoutingStrategy) {
		query["RoutingStrategy"] = request.RoutingStrategy
	}

	if !dara.IsNil(request.SecurityLockThreshold) {
		query["SecurityLockThreshold"] = request.SecurityLockThreshold
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maximum bandwidth and email address of a client account on a Smart Access Gateway (SAG) app instance.
//
// @param request - ModifySmartAccessGatewayClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmartAccessGatewayClientUserResponse
func (client *Client) ModifySmartAccessGatewayClientUserWithContext(ctx context.Context, request *ModifySmartAccessGatewayClientUserRequest, runtime *dara.RuntimeOptions) (_result *ModifySmartAccessGatewayClientUserResponse, _err error) {
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

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySmartAccessGatewayClientUser"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySmartAccessGatewayClientUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a Smart Access Gateway (SAG) instance.
//
// @param request - ModifySmartAccessGatewayUpBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySmartAccessGatewayUpBandwidthResponse
func (client *Client) ModifySmartAccessGatewayUpBandwidthWithContext(ctx context.Context, request *ModifySmartAccessGatewayUpBandwidthRequest, runtime *dara.RuntimeOptions) (_result *ModifySmartAccessGatewayUpBandwidthResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UpBandwidth4G) {
		query["UpBandwidth4G"] = request.UpBandwidth4G
	}

	if !dara.IsNil(request.UpBandwidthWan) {
		query["UpBandwidthWan"] = request.UpBandwidthWan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySmartAccessGatewayUpBandwidth"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySmartAccessGatewayUpBandwidthResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves Smart Access Gateway (SAG) resources from one resource group to another.
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
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2018-03-13"),
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
// Probes the network connectivity between a Smart Access Gateway (SAG) instance and an access point.
//
// Description:
//
// You can call this operation to test the connectivity between an SAG instance and a specified access point.
//
//   - If the SAG instance can connect to the access point, the ID of the request is returned in this operation.
//
//   - If the SAG instance cannot connect to the access point, the ID of the request and corresponding error messages are returned in this operation.
//
// @param request - ProbeAccessPointNetworkQualityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProbeAccessPointNetworkQualityResponse
func (client *Client) ProbeAccessPointNetworkQualityWithContext(ctx context.Context, request *ProbeAccessPointNetworkQualityRequest, runtime *dara.RuntimeOptions) (_result *ProbeAccessPointNetworkQualityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointIds) {
		query["AccessPointIds"] = request.AccessPointIds
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProbeAccessPointNetworkQuality"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProbeAccessPointNetworkQualityResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a Smart Access Gateway (SAG) instance.
//
// @param request - RebootSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootSmartAccessGatewayResponse
func (client *Client) RebootSmartAccessGatewayWithContext(ctx context.Context, request *RebootSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *RebootSmartAccessGatewayResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password that a client account uses to log on to the Smart Access Gateway (SAG) app.
//
// @param request - ResetSmartAccessGatewayClientUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSmartAccessGatewayClientUserPasswordResponse
func (client *Client) ResetSmartAccessGatewayClientUserPasswordWithContext(ctx context.Context, request *ResetSmartAccessGatewayClientUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetSmartAccessGatewayClientUserPasswordResponse, _err error) {
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

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetSmartAccessGatewayClientUserPassword"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetSmartAccessGatewayClientUserPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disallows a Cloud Enterprise Network (CEN) instance to communicate with a Cloud Connect Network (CCN) instance.
//
// @param request - RevokeInstanceFromCbnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromCbnResponse
func (client *Client) RevokeInstanceFromCbnWithContext(ctx context.Context, request *RevokeInstanceFromCbnRequest, runtime *dara.RuntimeOptions) (_result *RevokeInstanceFromCbnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnInstanceId) {
		query["CcnInstanceId"] = request.CcnInstanceId
	}

	if !dara.IsNil(request.CenInstanceId) {
		query["CenInstanceId"] = request.CenInstanceId
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
		Action:      dara.String("RevokeInstanceFromCbn"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeInstanceFromCbnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions that a Smart Access Gateway (SAG) instance has on virtual border routers (VBRs) in a different Alibaba Cloud account.
//
// @param request - RevokeInstanceFromVbrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromVbrResponse
func (client *Client) RevokeInstanceFromVbrWithContext(ctx context.Context, request *RevokeInstanceFromVbrRequest, runtime *dara.RuntimeOptions) (_result *RevokeInstanceFromVbrResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.VbrInstanceId) {
		query["VbrInstanceId"] = request.VbrInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeInstanceFromVbr"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeInstanceFromVbrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disallows a Smart Access Gateway (SAG) instance to communicate with a Cloud Connect Network (CCN) instance.
//
// @param request - RevokeSagInstanceFromCcnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeSagInstanceFromCcnResponse
func (client *Client) RevokeSagInstanceFromCcnWithContext(ctx context.Context, request *RevokeSagInstanceFromCcnRequest, runtime *dara.RuntimeOptions) (_result *RevokeSagInstanceFromCcnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnInstanceId) {
		query["CcnInstanceId"] = request.CcnInstanceId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeSagInstanceFromCcn"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeSagInstanceFromCcnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables roaming for the SAG app to allow clients to access Alibaba Cloud across regions.
//
// Description:
//
// Before you call `RoamClientUser`, we recommend that you read and understand the features and usage notes of roaming. For more information, see [Configure roaming on clients](https://help.aliyun.com/document_detail/177220.html).
//
// @param request - RoamClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RoamClientUserResponse
func (client *Client) RoamClientUserWithContext(ctx context.Context, request *RoamClientUserRequest, runtime *dara.RuntimeOptions) (_result *RoamClientUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OriginRegionId) {
		query["OriginRegionId"] = request.OriginRegionId
	}

	if !dara.IsNil(request.OriginSmartAGId) {
		query["OriginSmartAGId"] = request.OriginSmartAGId
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

	if !dara.IsNil(request.TargetSmartAGId) {
		query["TargetSmartAGId"] = request.TargetSmartAGId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RoamClientUser"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RoamClientUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the deep packet inspection (DPI) feature for a Smart Access Gateway (SAG) instance.
//
// @param request - SetAdvancedMonitorStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAdvancedMonitorStateResponse
func (client *Client) SetAdvancedMonitorStateWithContext(ctx context.Context, request *SetAdvancedMonitorStateRequest, runtime *dara.RuntimeOptions) (_result *SetAdvancedMonitorStateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAdvancedMonitorState"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAdvancedMonitorStateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the settings of a Smart Access Gateway (SAG) device to the associated SAG instance.
//
// Description:
//
// *SynchronizeSmartAGWebConfig*	- is an asynchronous operation. After you send a request, the request ID is returned but the operation is still being performed in the system background. You can call the [DescribeSAGDeviceInfo](https://help.aliyun.com/document_detail/164279.html) operation to query the status of an SAG device.
//
//   - If an SAG device is in the **Synchronizing*	- state, the settings of the SAG device are being synchronized to the associated SAG instance.
//
//   - If an SAG device is in the **Synchronized*	- state, the settings of the SAG device are synchronized to the associated SAG instance.
//
// @param request - SynchronizeSmartAGWebConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeSmartAGWebConfigResponse
func (client *Client) SynchronizeSmartAGWebConfigWithContext(ctx context.Context, request *SynchronizeSmartAGWebConfigRequest, runtime *dara.RuntimeOptions) (_result *SynchronizeSmartAGWebConfigResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGSn) {
		query["SmartAGSn"] = request.SmartAGSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SynchronizeSmartAGWebConfig"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SynchronizeSmartAGWebConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a Smart Access Gateway (SAG) device from the associated SAG instance.
//
// @param request - UnbindSerialNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSerialNumberResponse
func (client *Client) UnbindSerialNumberWithContext(ctx context.Context, request *UnbindSerialNumberRequest, runtime *dara.RuntimeOptions) (_result *UnbindSerialNumberResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindSerialNumber"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindSerialNumberResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a Smart Access Gateway (SAG) instance from the associated Cloud Connect Network (CCN) instance.
//
// @param request - UnbindSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindSmartAccessGatewayResponse
func (client *Client) UnbindSmartAccessGatewayWithContext(ctx context.Context, request *UnbindSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *UnbindSmartAccessGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CcnId) {
		query["CcnId"] = request.CcnId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGUid) {
		query["SmartAGUid"] = request.SmartAGUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a Smart Access Gateway (SAG) instance from a virtual border router (VBR).
//
// @param request - UnbindVbrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindVbrResponse
func (client *Client) UnbindVbrWithContext(ctx context.Context, request *UnbindVbrRequest, runtime *dara.RuntimeOptions) (_result *UnbindVbrResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.SmartAGUid) {
		query["SmartAGUid"] = request.SmartAGUid
	}

	if !dara.IsNil(request.VbrId) {
		query["VbrId"] = request.VbrId
	}

	if !dara.IsNil(request.VbrRegionId) {
		query["VbrRegionId"] = request.VbrRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindVbr"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindVbrResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks a Smart Access Gateway (SAG) instance.
//
// @param request - UnlockSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockSmartAccessGatewayResponse
func (client *Client) UnlockSmartAccessGatewayWithContext(ctx context.Context, request *UnlockSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *UnlockSmartAccessGatewayResponse, _err error) {
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnlockSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnlockSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a specified enterprise code.
//
// @param request - UpdateEnterpriseCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnterpriseCodeResponse
func (client *Client) UpdateEnterpriseCodeWithContext(ctx context.Context, request *UpdateEnterpriseCodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateEnterpriseCodeResponse, _err error) {
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

	if !dara.IsNil(request.EnterpriseCode) {
		query["EnterpriseCode"] = request.EnterpriseCode
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnterpriseCode"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnterpriseCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a probe task.
//
// @param request - UpdateProbeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProbeTaskResponse
func (client *Client) UpdateProbeTaskWithContext(ctx context.Context, request *UpdateProbeTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateProbeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.PacketNumber) {
		query["PacketNumber"] = request.PacketNumber
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.ProbeTaskId) {
		query["ProbeTaskId"] = request.ProbeTaskId
	}

	if !dara.IsNil(request.ProbeTaskSourceAddress) {
		query["ProbeTaskSourceAddress"] = request.ProbeTaskSourceAddress
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagId) {
		query["SagId"] = request.SagId
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProbeTask"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProbeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the access point of a Smart Access Gateway (SAG) instance.
//
// Description:
//
// ## Prerequisites
//
// Before you call this operation, you can call the [ListAccessPoints](https://help.aliyun.com/document_detail/183876.html) operation to view the switchable access points of the SAG instance.
//
// @param request - UpdateSmartAGAccessPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAGAccessPointResponse
func (client *Client) UpdateSmartAGAccessPointWithContext(ctx context.Context, request *UpdateSmartAGAccessPointRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAGAccessPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPointId) {
		query["AccessPointId"] = request.AccessPointId
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAGAccessPoint"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAGAccessPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the deep packet inspection (DPI) feature for a Smart Access Gateway (SAG) instance.
//
// @param request - UpdateSmartAGDpiAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAGDpiAttributeResponse
func (client *Client) UpdateSmartAGDpiAttributeWithContext(ctx context.Context, request *UpdateSmartAGDpiAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAGDpiAttributeResponse, _err error) {
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

	if !dara.IsNil(request.DpiEnabled) {
		query["DpiEnabled"] = request.DpiEnabled
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAGDpiAttribute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAGDpiAttributeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a Smart Access Gateway (SAG) APP instance with another enterprise code.
//
// @param request - UpdateSmartAGEnterpriseCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAGEnterpriseCodeResponse
func (client *Client) UpdateSmartAGEnterpriseCodeWithContext(ctx context.Context, request *UpdateSmartAGEnterpriseCodeRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAGEnterpriseCodeResponse, _err error) {
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

	if !dara.IsNil(request.EnterpriseCode) {
		query["EnterpriseCode"] = request.EnterpriseCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAGEnterpriseCode"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAGEnterpriseCodeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maximum bandwidth for application acceleration of client accounts on Smart Access Gateway (SAG) app.
//
// Description:
//
// Before you set a maximum bandwidth value for a client, take note of the following rules:
//
//   - The maximum bandwidth value of a client cannot exceed that of the SAG app instance to which the client belongs.
//
//   - If you have not set maximum bandwidth values for clients that belong to an SAG app instance, and the maximum bandwidth value of the instance is less than 5 Mbit/s, for example, 4 Mbit/s, the maximum bandwidth value of each client that belongs to the SAG app instance is 4 Mbit/s by default.
//
//   - If you have not set maximum bandwidth values for clients that belong to an SAG app instance, and the maximum bandwidth value of the instance is 5 Mbit/s or higher, the maximum bandwidth value of each client that belongs to the SAG app instance is 5 Mbit/s by default.
//
// @param request - UpdateSmartAGUserAccelerateConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAGUserAccelerateConfigResponse
func (client *Client) UpdateSmartAGUserAccelerateConfigWithContext(ctx context.Context, request *UpdateSmartAGUserAccelerateConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAGUserAccelerateConfigResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAGUserAccelerateConfig"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAGUserAccelerateConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the password that is used to log on to an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can modify passwords that are used to log on to only SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - UpdateSmartAccessGatewayAdminPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayAdminPasswordResponse
func (client *Client) UpdateSmartAccessGatewayAdminPasswordWithContext(ctx context.Context, request *UpdateSmartAccessGatewayAdminPasswordRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayAdminPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayAdminPassword"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayAdminPasswordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the BGP configurations of an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can modify the BGP configuration only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - UpdateSmartAccessGatewayBgpRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayBgpRouteResponse
func (client *Client) UpdateSmartAccessGatewayBgpRouteWithContext(ctx context.Context, request *UpdateSmartAccessGatewayBgpRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayBgpRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.HoldTime) {
		query["HoldTime"] = request.HoldTime
	}

	if !dara.IsNil(request.KeepAlive) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !dara.IsNil(request.LocalAs) {
		query["LocalAs"] = request.LocalAs
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.RouterId) {
		query["RouterId"] = request.RouterId
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayBgpRoute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayBgpRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the DNS configuration of an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can modify the DNS configuration only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - UpdateSmartAccessGatewayDnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayDnsResponse
func (client *Client) UpdateSmartAccessGatewayDnsWithContext(ctx context.Context, request *UpdateSmartAccessGatewayDnsRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayDnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.MasterDns) {
		query["MasterDns"] = request.MasterDns
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	if !dara.IsNil(request.SlaveDns) {
		query["SlaveDns"] = request.SlaveDns
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayDns"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayDnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the DNS forwarding configurations for a Smart Access Gateway (SAG) SCG5000 or SCG5000-5G instance. The version of the device must be 3.4.2 or later.
//
// @param request - UpdateSmartAccessGatewayDnsForwardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayDnsForwardResponse
func (client *Client) UpdateSmartAccessGatewayDnsForwardWithContext(ctx context.Context, request *UpdateSmartAccessGatewayDnsForwardRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayDnsForwardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MasterIp) {
		query["MasterIp"] = request.MasterIp
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.OutboundPortIndex) {
		query["OutboundPortIndex"] = request.OutboundPortIndex
	}

	if !dara.IsNil(request.OutboundPortName) {
		query["OutboundPortName"] = request.OutboundPortName
	}

	if !dara.IsNil(request.OutboundPortType) {
		query["OutboundPortType"] = request.OutboundPortType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	if !dara.IsNil(request.SlaveIp) {
		query["SlaveIp"] = request.SlaveIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayDnsForward"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayDnsForwardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the global routing protocol of an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can modify the global routing protocol only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - UpdateSmartAccessGatewayGlobalRouteProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayGlobalRouteProtocolResponse
func (client *Client) UpdateSmartAccessGatewayGlobalRouteProtocolWithContext(ctx context.Context, request *UpdateSmartAccessGatewayGlobalRouteProtocolRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayGlobalRouteProtocolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.RouteProtocol) {
		query["RouteProtocol"] = request.RouteProtocol
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayGlobalRouteProtocol"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayGlobalRouteProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the OSPF configurations for an SAG SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// @param request - UpdateSmartAccessGatewayOspfRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayOspfRouteResponse
func (client *Client) UpdateSmartAccessGatewayOspfRouteWithContext(ctx context.Context, request *UpdateSmartAccessGatewayOspfRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayOspfRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AreaId) {
		query["AreaId"] = request.AreaId
	}

	if !dara.IsNil(request.AuthenticationType) {
		query["AuthenticationType"] = request.AuthenticationType
	}

	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.DeadTime) {
		query["DeadTime"] = request.DeadTime
	}

	if !dara.IsNil(request.HelloTime) {
		query["HelloTime"] = request.HelloTime
	}

	if !dara.IsNil(request.InterfaceName) {
		query["InterfaceName"] = request.InterfaceName
	}

	if !dara.IsNil(request.Md5Key) {
		query["Md5Key"] = request.Md5Key
	}

	if !dara.IsNil(request.Md5KeyId) {
		query["Md5KeyId"] = request.Md5KeyId
	}

	if !dara.IsNil(request.Networks) {
		query["Networks"] = request.Networks
	}

	if !dara.IsNil(request.OspfCost) {
		query["OspfCost"] = request.OspfCost
	}

	if !dara.IsNil(request.OspfNetworkType) {
		query["OspfNetworkType"] = request.OspfNetworkType
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RedistributeProtocol) {
		query["RedistributeProtocol"] = request.RedistributeProtocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.RouterId) {
		query["RouterId"] = request.RouterId
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayOspfRoute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayOspfRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the port protocol of an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can modify the port protocol only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - UpdateSmartAccessGatewayPortRouteProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayPortRouteProtocolResponse
func (client *Client) UpdateSmartAccessGatewayPortRouteProtocolWithContext(ctx context.Context, request *UpdateSmartAccessGatewayPortRouteProtocolRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayPortRouteProtocolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.PortName) {
		query["PortName"] = request.PortName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RemoteAs) {
		query["RemoteAs"] = request.RemoteAs
	}

	if !dara.IsNil(request.RemoteIp) {
		query["RemoteIp"] = request.RemoteIp
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.RouteProtocol) {
		query["RouteProtocol"] = request.RouteProtocol
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	if !dara.IsNil(request.Vlan) {
		query["Vlan"] = request.Vlan
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayPortRouteProtocol"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayPortRouteProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a SAG device to a later version.
//
// @param request - UpdateSmartAccessGatewayVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayVersionResponse
func (client *Client) UpdateSmartAccessGatewayVersionWithContext(ctx context.Context, request *UpdateSmartAccessGatewayVersionRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayVersionResponse, _err error) {
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

	if !dara.IsNil(request.SerialNumber) {
		query["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.VersionCode) {
		query["VersionCode"] = request.VersionCode
	}

	if !dara.IsNil(request.VersionType) {
		query["VersionType"] = request.VersionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayVersion"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the Source Network Address Translation (SNAT) configuration of the WAN port on an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can modify the SNAT configuration of the WAN port only on SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - UpdateSmartAccessGatewayWanSnatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSmartAccessGatewayWanSnatResponse
func (client *Client) UpdateSmartAccessGatewayWanSnatWithContext(ctx context.Context, request *UpdateSmartAccessGatewayWanSnatRequest, runtime *dara.RuntimeOptions) (_result *UpdateSmartAccessGatewayWanSnatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	if !dara.IsNil(request.Snat) {
		query["Snat"] = request.Snat
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSmartAccessGatewayWanSnat"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSmartAccessGatewayWanSnatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Increases the bandwidth of a Smart Access Gateway (SAG) instance.
//
// @param request - UpgradeSmartAccessGatewayRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeSmartAccessGatewayResponse
func (client *Client) UpgradeSmartAccessGatewayWithContext(ctx context.Context, request *UpgradeSmartAccessGatewayRequest, runtime *dara.RuntimeOptions) (_result *UpgradeSmartAccessGatewayResponse, _err error) {
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

	if !dara.IsNil(request.BandWidthSpec) {
		query["BandWidthSpec"] = request.BandWidthSpec
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeSmartAccessGateway"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeSmartAccessGatewayResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Increases the maximum number of client accounts supported by a Smart Access Gateway (SAG) app instance.
//
// @param request - UpgradeSmartAccessGatewaySoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeSmartAccessGatewaySoftwareResponse
func (client *Client) UpgradeSmartAccessGatewaySoftwareWithContext(ctx context.Context, request *UpgradeSmartAccessGatewaySoftwareRequest, runtime *dara.RuntimeOptions) (_result *UpgradeSmartAccessGatewaySoftwareResponse, _err error) {
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

	if !dara.IsNil(request.DataPlan) {
		query["DataPlan"] = request.DataPlan
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

	if !dara.IsNil(request.SmartAGId) {
		query["SmartAGId"] = request.SmartAGId
	}

	if !dara.IsNil(request.UserCount) {
		query["UserCount"] = request.UserCount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeSmartAccessGatewaySoftware"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeSmartAccessGatewaySoftwareResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the BGP configuration of an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can query the BGP configuration only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - ViewSmartAccessGatewayBgpRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewSmartAccessGatewayBgpRouteResponse
func (client *Client) ViewSmartAccessGatewayBgpRouteWithContext(ctx context.Context, request *ViewSmartAccessGatewayBgpRouteRequest, runtime *dara.RuntimeOptions) (_result *ViewSmartAccessGatewayBgpRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewSmartAccessGatewayBgpRoute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewSmartAccessGatewayBgpRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the DNS configuration of an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can query the DNS configuration only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - ViewSmartAccessGatewayDnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewSmartAccessGatewayDnsResponse
func (client *Client) ViewSmartAccessGatewayDnsWithContext(ctx context.Context, request *ViewSmartAccessGatewayDnsRequest, runtime *dara.RuntimeOptions) (_result *ViewSmartAccessGatewayDnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewSmartAccessGatewayDns"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewSmartAccessGatewayDnsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the DNS forwarding list of a Smart Access Gateway (SAG) SCG5000 or SCG5000-5G instance. The version of the device must be 3.4.2 or later.
//
// @param request - ViewSmartAccessGatewayDnsForwardsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewSmartAccessGatewayDnsForwardsResponse
func (client *Client) ViewSmartAccessGatewayDnsForwardsWithContext(ctx context.Context, request *ViewSmartAccessGatewayDnsForwardsRequest, runtime *dara.RuntimeOptions) (_result *ViewSmartAccessGatewayDnsForwardsResponse, _err error) {
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

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewSmartAccessGatewayDnsForwards"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewSmartAccessGatewayDnsForwardsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global routing protocol of an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can query the global protocol only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - ViewSmartAccessGatewayGlobalRouteProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewSmartAccessGatewayGlobalRouteProtocolResponse
func (client *Client) ViewSmartAccessGatewayGlobalRouteProtocolWithContext(ctx context.Context, request *ViewSmartAccessGatewayGlobalRouteProtocolRequest, runtime *dara.RuntimeOptions) (_result *ViewSmartAccessGatewayGlobalRouteProtocolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewSmartAccessGatewayGlobalRouteProtocol"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewSmartAccessGatewayGlobalRouteProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the BGP configuration of an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later and has OSPF enabled.
//
// Description:
//
// You can query the BGP configuration only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later and have OSPF enabled.
//
// @param request - ViewSmartAccessGatewayOspfRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewSmartAccessGatewayOspfRouteResponse
func (client *Client) ViewSmartAccessGatewayOspfRouteWithContext(ctx context.Context, request *ViewSmartAccessGatewayOspfRouteRequest, runtime *dara.RuntimeOptions) (_result *ViewSmartAccessGatewayOspfRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewSmartAccessGatewayOspfRoute"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewSmartAccessGatewayOspfRouteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ports that have routing protocols enabled on an SAG SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// @param request - ViewSmartAccessGatewayPortRouteProtocolRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewSmartAccessGatewayPortRouteProtocolResponse
func (client *Client) ViewSmartAccessGatewayPortRouteProtocolWithContext(ctx context.Context, request *ViewSmartAccessGatewayPortRouteProtocolRequest, runtime *dara.RuntimeOptions) (_result *ViewSmartAccessGatewayPortRouteProtocolResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewSmartAccessGatewayPortRouteProtocol"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewSmartAccessGatewayPortRouteProtocolResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route details about an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can query the route details only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - ViewSmartAccessGatewayRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewSmartAccessGatewayRoutesResponse
func (client *Client) ViewSmartAccessGatewayRoutesWithContext(ctx context.Context, request *ViewSmartAccessGatewayRoutesRequest, runtime *dara.RuntimeOptions) (_result *ViewSmartAccessGatewayRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewSmartAccessGatewayRoutes"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewSmartAccessGatewayRoutesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Source Network Address Translation (SNAT) configuration of the WAN port on an SCG5000 or SCG5000-5G device whose version is 3.4.2 or later.
//
// Description:
//
// You can query the SNAT configuration of the WAN port only of SCG5000 and SCG5000-5G devices whose version is 3.4.2 or later.
//
// @param request - ViewSmartAccessGatewayWanSnatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewSmartAccessGatewayWanSnatResponse
func (client *Client) ViewSmartAccessGatewayWanSnatWithContext(ctx context.Context, request *ViewSmartAccessGatewayWanSnatRequest, runtime *dara.RuntimeOptions) (_result *ViewSmartAccessGatewayWanSnatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CrossAccount) {
		query["CrossAccount"] = request.CrossAccount
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceUid) {
		query["ResourceUid"] = request.ResourceUid
	}

	if !dara.IsNil(request.SagInsId) {
		query["SagInsId"] = request.SagInsId
	}

	if !dara.IsNil(request.SagSn) {
		query["SagSn"] = request.SagSn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewSmartAccessGatewayWanSnat"),
		Version:     dara.String("2018-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewSmartAccessGatewayWanSnatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
