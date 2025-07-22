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
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("expressconnectrouter"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Enables log delivery for flow logs.
//
// @param request - ActivateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateFlowLogResponse
func (client *Client) ActivateFlowLogWithOptions(request *ActivateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *ActivateFlowLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FlowLogId) {
		body["FlowLogId"] = request.FlowLogId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateFlowLog"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateFlowLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables log delivery for flow logs.
//
// @param request - ActivateFlowLogRequest
//
// @return ActivateFlowLogResponse
func (client *Client) ActivateFlowLog(request *ActivateFlowLogRequest) (_result *ActivateFlowLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActivateFlowLogResponse{}
	_body, _err := client.ActivateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a virtual border router (VBR) with an Express Connect router (ECR).
//
// Description:
//
// Before you call the **AttachExpressConnectRouterChildInstance*	- operation to associate a VBR with an ECR, make sure that the ECR is in the **Active*	- state.
//
// @param request - AttachExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachExpressConnectRouterChildInstanceResponse
func (client *Client) AttachExpressConnectRouterChildInstanceWithOptions(request *AttachExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChildInstanceId) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceOwnerId) {
		body["ChildInstanceOwnerId"] = request.ChildInstanceOwnerId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		body["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		body["ChildInstanceType"] = request.ChildInstanceType
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

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachExpressConnectRouterChildInstance"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a virtual border router (VBR) with an Express Connect router (ECR).
//
// Description:
//
// Before you call the **AttachExpressConnectRouterChildInstance*	- operation to associate a VBR with an ECR, make sure that the ECR is in the **Active*	- state.
//
// @param request - AttachExpressConnectRouterChildInstanceRequest
//
// @return AttachExpressConnectRouterChildInstanceResponse
func (client *Client) AttachExpressConnectRouterChildInstance(request *AttachExpressConnectRouterChildInstanceRequest) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.AttachExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks the Cloud Data Transfer (CDT) service required to add a region to an Express Connect router (ECR).
//
// @param request - CheckAddRegionToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAddRegionToExpressConnectRouterResponse
func (client *Client) CheckAddRegionToExpressConnectRouterWithOptions(request *CheckAddRegionToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FreshRegionId) {
		body["FreshRegionId"] = request.FreshRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckAddRegionToExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckAddRegionToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the Cloud Data Transfer (CDT) service required to add a region to an Express Connect router (ECR).
//
// @param request - CheckAddRegionToExpressConnectRouterRequest
//
// @return CheckAddRegionToExpressConnectRouterResponse
func (client *Client) CheckAddRegionToExpressConnectRouter(request *CheckAddRegionToExpressConnectRouterRequest) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckAddRegionToExpressConnectRouterResponse{}
	_body, _err := client.CheckAddRegionToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Express Connect Router (ECR).
//
// Description:
//
// After you create an ECR, it enters the **Active*	- state.
//
// @param request - CreateExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressConnectRouterResponse
func (client *Client) CreateExpressConnectRouterWithOptions(request *CreateExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *CreateExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlibabaSideAsn) {
		body["AlibabaSideAsn"] = request.AlibabaSideAsn
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

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
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
		Action:      dara.String("CreateExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Express Connect Router (ECR).
//
// Description:
//
// After you create an ECR, it enters the **Active*	- state.
//
// @param request - CreateExpressConnectRouterRequest
//
// @return CreateExpressConnectRouterResponse
func (client *Client) CreateExpressConnectRouter(request *CreateExpressConnectRouterRequest) (_result *CreateExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExpressConnectRouterResponse{}
	_body, _err := client.CreateExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a virtual private cloud (VPC) or a transit router (TR) with an Express Connect router (ECR).
//
// @param request - CreateExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressConnectRouterAssociationResponse
func (client *Client) CreateExpressConnectRouterAssociationWithOptions(request *CreateExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedPrefixes) {
		body["AllowedPrefixes"] = request.AllowedPrefixes
	}

	if !dara.IsNil(request.AllowedPrefixesMode) {
		body["AllowedPrefixesMode"] = request.AllowedPrefixesMode
	}

	if !dara.IsNil(request.AssociationRegionId) {
		body["AssociationRegionId"] = request.AssociationRegionId
	}

	if !dara.IsNil(request.CenId) {
		body["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CreateAttachment) {
		body["CreateAttachment"] = request.CreateAttachment
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.TransitRouterId) {
		body["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.TransitRouterOwnerId) {
		body["TransitRouterOwnerId"] = request.TransitRouterOwnerId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.VpcOwnerId) {
		body["VpcOwnerId"] = request.VpcOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExpressConnectRouterAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a virtual private cloud (VPC) or a transit router (TR) with an Express Connect router (ECR).
//
// @param request - CreateExpressConnectRouterAssociationRequest
//
// @return CreateExpressConnectRouterAssociationResponse
func (client *Client) CreateExpressConnectRouterAssociation(request *CreateExpressConnectRouterAssociationRequest) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExpressConnectRouterAssociationResponse{}
	_body, _err := client.CreateExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a flow log and enables log delivery.
//
// @param request - CreateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowLogResponse
func (client *Client) CreateFlowLogWithOptions(request *CreateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *CreateFlowLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.LogStoreName) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SamplingRate) {
		query["SamplingRate"] = request.SamplingRate
	}

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

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FlowLogName) {
		body["FlowLogName"] = request.FlowLogName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFlowLog"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a flow log and enables log delivery.
//
// @param request - CreateFlowLogRequest
//
// @return CreateFlowLogResponse
func (client *Client) CreateFlowLog(request *CreateFlowLogRequest) (_result *CreateFlowLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFlowLogResponse{}
	_body, _err := client.CreateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables log delivery.
//
// @param request - DeactivateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateFlowLogResponse
func (client *Client) DeactivateFlowLogWithOptions(request *DeactivateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *DeactivateFlowLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FlowLogId) {
		body["FlowLogId"] = request.FlowLogId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeactivateFlowLog"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeactivateFlowLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables log delivery.
//
// @param request - DeactivateFlowLogRequest
//
// @return DeactivateFlowLogResponse
func (client *Client) DeactivateFlowLog(request *DeactivateFlowLogRequest) (_result *DeactivateFlowLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeactivateFlowLogResponse{}
	_body, _err := client.DeactivateFlowLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect router (ECR).
//
// Description:
//
// Take note of the following items:
//
//   - Before you call this operation, make sure that all resources are disassociated from the ECR.
//
//   - You can delete only ECRs that are in the **Active*	- state.
//
// @param request - DeleteExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressConnectRouterResponse
func (client *Client) DeleteExpressConnectRouterWithOptions(request *DeleteExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DeleteExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect router (ECR).
//
// Description:
//
// Take note of the following items:
//
//   - Before you call this operation, make sure that all resources are disassociated from the ECR.
//
//   - You can delete only ECRs that are in the **Active*	- state.
//
// @param request - DeleteExpressConnectRouterRequest
//
// @return DeleteExpressConnectRouterResponse
func (client *Client) DeleteExpressConnectRouter(request *DeleteExpressConnectRouterRequest) (_result *DeleteExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExpressConnectRouterResponse{}
	_body, _err := client.DeleteExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates an Express Connect router (ECR) from a virtual private cloud (VPC) or a transit router (TR).
//
// @param request - DeleteExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressConnectRouterAssociationResponse
func (client *Client) DeleteExpressConnectRouterAssociationWithOptions(request *DeleteExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeleteAttachment) {
		body["DeleteAttachment"] = request.DeleteAttachment
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExpressConnectRouterAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an Express Connect router (ECR) from a virtual private cloud (VPC) or a transit router (TR).
//
// @param request - DeleteExpressConnectRouterAssociationRequest
//
// @return DeleteExpressConnectRouterAssociationResponse
func (client *Client) DeleteExpressConnectRouterAssociation(request *DeleteExpressConnectRouterAssociationRequest) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteExpressConnectRouterAssociationResponse{}
	_body, _err := client.DeleteExpressConnectRouterAssociationWithOptions(request, runtime)
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
// @param request - DeleteFlowlogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowlogResponse
func (client *Client) DeleteFlowlogWithOptions(request *DeleteFlowlogRequest, runtime *dara.RuntimeOptions) (_result *DeleteFlowlogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFlowlog"),
		Version:     dara.String("2023-09-01"),
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
// Queries the route entries that are disabled on an Express Connect router (ECR).
//
// @param request - DescribeDisabledExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisabledExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeDisabledExpressConnectRouterRouteEntriesWithOptions(request *DescribeDisabledExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDisabledExpressConnectRouterRouteEntries"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route entries that are disabled on an Express Connect router (ECR).
//
// @param request - DescribeDisabledExpressConnectRouterRouteEntriesRequest
//
// @return DescribeDisabledExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeDisabledExpressConnectRouterRouteEntries(request *DescribeDisabledExpressConnectRouterRouteEntriesRequest) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDisabledExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DescribeDisabledExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of Express Connect routers (ECRs).
//
// @param request - DescribeExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterResponse
func (client *Client) DescribeExpressConnectRouterWithOptions(request *DescribeExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
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
		Action:      dara.String("DescribeExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Express Connect routers (ECRs).
//
// @param request - DescribeExpressConnectRouterRequest
//
// @return DescribeExpressConnectRouterResponse
func (client *Client) DescribeExpressConnectRouter(request *DescribeExpressConnectRouterRequest) (_result *DescribeExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterResponse{}
	_body, _err := client.DescribeExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical route prefixes of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterAllowedPrefixHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterAllowedPrefixHistoryResponse
func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistoryWithOptions(request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterAllowedPrefixHistory"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical route prefixes of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterAllowedPrefixHistoryRequest
//
// @return DescribeExpressConnectRouterAllowedPrefixHistoryResponse
func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistory(request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterAllowedPrefixHistoryResponse{}
	_body, _err := client.DescribeExpressConnectRouterAllowedPrefixHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the virtual private clouds (VPCs) and transit routers (TRs) associated with an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterAssociationResponse
func (client *Client) DescribeExpressConnectRouterAssociationWithOptions(request *DescribeExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.AssociationNodeType) {
		body["AssociationNodeType"] = request.AssociationNodeType
	}

	if !dara.IsNil(request.AssociationRegionId) {
		body["AssociationRegionId"] = request.AssociationRegionId
	}

	if !dara.IsNil(request.CenId) {
		body["CenId"] = request.CenId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TransitRouterId) {
		body["TransitRouterId"] = request.TransitRouterId
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual private clouds (VPCs) and transit routers (TRs) associated with an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterAssociationRequest
//
// @return DescribeExpressConnectRouterAssociationResponse
func (client *Client) DescribeExpressConnectRouterAssociation(request *DescribeExpressConnectRouterAssociationRequest) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterAssociationResponse{}
	_body, _err := client.DescribeExpressConnectRouterAssociationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the virtual border routers (VBRs) that are associated with an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterChildInstanceResponse
func (client *Client) DescribeExpressConnectRouterChildInstanceWithOptions(request *DescribeExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ChildInstanceId) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceRegionId) {
		body["ChildInstanceRegionId"] = request.ChildInstanceRegionId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterChildInstance"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual border routers (VBRs) that are associated with an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterChildInstanceRequest
//
// @return DescribeExpressConnectRouterChildInstanceResponse
func (client *Client) DescribeExpressConnectRouterChildInstance(request *DescribeExpressConnectRouterChildInstanceRequest) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.DescribeExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cross-region forwarding modes of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterInterRegionTransitModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) DescribeExpressConnectRouterInterRegionTransitModeWithOptions(request *DescribeExpressConnectRouterInterRegionTransitModeRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterInterRegionTransitMode"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cross-region forwarding modes of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterInterRegionTransitModeRequest
//
// @return DescribeExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) DescribeExpressConnectRouterInterRegionTransitMode(request *DescribeExpressConnectRouterInterRegionTransitModeRequest) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.DescribeExpressConnectRouterInterRegionTransitModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of regions in which resources related to an Express Connect router (ECR) are deployed.
//
// @param request - DescribeExpressConnectRouterRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterRegionResponse
func (client *Client) DescribeExpressConnectRouterRegionWithOptions(request *DescribeExpressConnectRouterRegionRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterRegion"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of regions in which resources related to an Express Connect router (ECR) are deployed.
//
// @param request - DescribeExpressConnectRouterRegionRequest
//
// @return DescribeExpressConnectRouterRegionResponse
func (client *Client) DescribeExpressConnectRouterRegion(request *DescribeExpressConnectRouterRegionRequest) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterRegionResponse{}
	_body, _err := client.DescribeExpressConnectRouterRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the route entries of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeExpressConnectRouterRouteEntriesWithOptions(request *DescribeExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AsPath) {
		body["AsPath"] = request.AsPath
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Community) {
		body["Community"] = request.Community
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NexthopInstanceId) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	if !dara.IsNil(request.QueryRegionId) {
		body["QueryRegionId"] = request.QueryRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeExpressConnectRouterRouteEntries"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route entries of an Express Connect router (ECR).
//
// @param request - DescribeExpressConnectRouterRouteEntriesRequest
//
// @return DescribeExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeExpressConnectRouterRouteEntries(request *DescribeExpressConnectRouterRouteEntriesRequest) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DescribeExpressConnectRouterRouteEntriesWithOptions(request, runtime)
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
// @param request - DescribeFlowLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFlowLogsResponse
func (client *Client) DescribeFlowLogsWithOptions(request *DescribeFlowLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeFlowLogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.FlowLogName) {
		query["FlowLogName"] = request.FlowLogName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LogStoreName) {
		query["LogStoreName"] = request.LogStoreName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

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

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFlowLogs"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFlowLogsResponse{}
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
// @param request - DescribeFlowLogsRequest
//
// @return DescribeFlowLogsResponse
func (client *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (_result *DescribeFlowLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeFlowLogsResponse{}
	_body, _err := client.DescribeFlowLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network instances whose permissions are granted to an Express Connect router (ECR).
//
// @param request - DescribeInstanceGrantedToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceGrantedToExpressConnectRouterResponse
func (client *Client) DescribeInstanceGrantedToExpressConnectRouterWithOptions(request *DescribeInstanceGrantedToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CallerType) {
		body["CallerType"] = request.CallerType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceOwnerId) {
		body["InstanceOwnerId"] = request.InstanceOwnerId
	}

	if !dara.IsNil(request.InstanceRegionId) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagModels) {
		body["TagModels"] = request.TagModels
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceGrantedToExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network instances whose permissions are granted to an Express Connect router (ECR).
//
// @param request - DescribeInstanceGrantedToExpressConnectRouterRequest
//
// @return DescribeInstanceGrantedToExpressConnectRouterResponse
func (client *Client) DescribeInstanceGrantedToExpressConnectRouter(request *DescribeInstanceGrantedToExpressConnectRouterRequest) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeInstanceGrantedToExpressConnectRouterResponse{}
	_body, _err := client.DescribeInstanceGrantedToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a virtual border router (VBR) from an Express Connect router (ECR).
//
// Description:
//
// Before you call the **DetachExpressConnectRouterChildInstance*	- operation to uninstall a VBR from an ECR, make sure that the ECR is in the **Active*	- state.
//
// @param request - DetachExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachExpressConnectRouterChildInstanceResponse
func (client *Client) DetachExpressConnectRouterChildInstanceWithOptions(request *DetachExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChildInstanceId) {
		body["ChildInstanceId"] = request.ChildInstanceId
	}

	if !dara.IsNil(request.ChildInstanceType) {
		body["ChildInstanceType"] = request.ChildInstanceType
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachExpressConnectRouterChildInstance"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a virtual border router (VBR) from an Express Connect router (ECR).
//
// Description:
//
// Before you call the **DetachExpressConnectRouterChildInstance*	- operation to uninstall a VBR from an ECR, make sure that the ECR is in the **Active*	- state.
//
// @param request - DetachExpressConnectRouterChildInstanceRequest
//
// @return DetachExpressConnectRouterChildInstanceResponse
func (client *Client) DetachExpressConnectRouterChildInstance(request *DetachExpressConnectRouterChildInstanceRequest) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.DetachExpressConnectRouterChildInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables route entries of an Express Connect router (ECR).
//
// @param request - DisableExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableExpressConnectRouterRouteEntriesResponse
func (client *Client) DisableExpressConnectRouterRouteEntriesWithOptions(request *DisableExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.NexthopInstanceId) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableExpressConnectRouterRouteEntries"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables route entries of an Express Connect router (ECR).
//
// @param request - DisableExpressConnectRouterRouteEntriesRequest
//
// @return DisableExpressConnectRouterRouteEntriesResponse
func (client *Client) DisableExpressConnectRouterRouteEntries(request *DisableExpressConnectRouterRouteEntriesRequest) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.DisableExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables route entries of an Express Connect router (ECR).
//
// @param request - EnableExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableExpressConnectRouterRouteEntriesResponse
func (client *Client) EnableExpressConnectRouterRouteEntriesWithOptions(request *EnableExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DestinationCidrBlock) {
		body["DestinationCidrBlock"] = request.DestinationCidrBlock
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.NexthopInstanceId) {
		body["NexthopInstanceId"] = request.NexthopInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableExpressConnectRouterRouteEntries"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables route entries of an Express Connect router (ECR).
//
// @param request - EnableExpressConnectRouterRouteEntriesRequest
//
// @return EnableExpressConnectRouterRouteEntriesResponse
func (client *Client) EnableExpressConnectRouterRouteEntries(request *EnableExpressConnectRouterRouteEntriesRequest) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableExpressConnectRouterRouteEntriesResponse{}
	_body, _err := client.EnableExpressConnectRouterRouteEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect router (ECR) and disassociates the virtual private cloud (VPC), transit router (TR), and virtual border router (VBR) associated with the ECR.
//
// Description:
//
//	  If you forcefully delete an ECR, all the resources associated with the ECR are disassociated at a time. Make sure that the disassociation does not affect the stability of your business.
//
//		- You can delete only ECRs that are in the **Active*	- state.
//
// @param request - ForceDeleteExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForceDeleteExpressConnectRouterResponse
func (client *Client) ForceDeleteExpressConnectRouterWithOptions(request *ForceDeleteExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForceDeleteExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForceDeleteExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect router (ECR) and disassociates the virtual private cloud (VPC), transit router (TR), and virtual border router (VBR) associated with the ECR.
//
// Description:
//
//	  If you forcefully delete an ECR, all the resources associated with the ECR are disassociated at a time. Make sure that the disassociation does not affect the stability of your business.
//
//		- You can delete only ECRs that are in the **Active*	- state.
//
// @param request - ForceDeleteExpressConnectRouterRequest
//
// @return ForceDeleteExpressConnectRouterResponse
func (client *Client) ForceDeleteExpressConnectRouter(request *ForceDeleteExpressConnectRouterRequest) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ForceDeleteExpressConnectRouterResponse{}
	_body, _err := client.ForceDeleteExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions on a virtual private cloud (VPC) or a virtual border router (VBR) to an Express Connect router (ECR) of another account.
//
// Description:
//
// When you associate a network instance of another account with an ECR, you must grant permissions on the network instance to the ECR.
//
// @param request - GrantInstanceToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantInstanceToExpressConnectRouterResponse
func (client *Client) GrantInstanceToExpressConnectRouterWithOptions(request *GrantInstanceToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.EcrOwnerAliUid) {
		body["EcrOwnerAliUid"] = request.EcrOwnerAliUid
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceRegionId) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantInstanceToExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantInstanceToExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on a virtual private cloud (VPC) or a virtual border router (VBR) to an Express Connect router (ECR) of another account.
//
// Description:
//
// When you associate a network instance of another account with an ECR, you must grant permissions on the network instance to the ECR.
//
// @param request - GrantInstanceToExpressConnectRouterRequest
//
// @return GrantInstanceToExpressConnectRouterResponse
func (client *Client) GrantInstanceToExpressConnectRouter(request *GrantInstanceToExpressConnectRouterRequest) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantInstanceToExpressConnectRouterResponse{}
	_body, _err := client.GrantInstanceToExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of regions in which the Express Connect router (ECR) feature is activated.
//
// @param request - ListExpressConnectRouterSupportedRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExpressConnectRouterSupportedRegionResponse
func (client *Client) ListExpressConnectRouterSupportedRegionWithOptions(request *ListExpressConnectRouterSupportedRegionRequest, runtime *dara.RuntimeOptions) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.NodeType) {
		body["NodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExpressConnectRouterSupportedRegion"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExpressConnectRouterSupportedRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of regions in which the Express Connect router (ECR) feature is activated.
//
// @param request - ListExpressConnectRouterSupportedRegionRequest
//
// @return ListExpressConnectRouterSupportedRegionResponse
func (client *Client) ListExpressConnectRouterSupportedRegion(request *ListExpressConnectRouterSupportedRegionRequest) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExpressConnectRouterSupportedRegionResponse{}
	_body, _err := client.ListExpressConnectRouterSupportedRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags that are added to an Express Connect router (ECR).
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2023-09-01"),
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
// Queries a list of tags that are added to an Express Connect router (ECR).
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
// Modifies the properties such as the name of an Express Connect router (ECR).
//
// Description:
//
// You can modify only properties of ECRs in the **Active*	- state.
//
// @param request - ModifyExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterResponse
func (client *Client) ModifyExpressConnectRouterWithOptions(request *ModifyExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties such as the name of an Express Connect router (ECR).
//
// Description:
//
// You can modify only properties of ECRs in the **Active*	- state.
//
// @param request - ModifyExpressConnectRouterRequest
//
// @return ModifyExpressConnectRouterResponse
func (client *Client) ModifyExpressConnectRouter(request *ModifyExpressConnectRouterRequest) (_result *ModifyExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterResponse{}
	_body, _err := client.ModifyExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the route prefixes of a virtual private cloud (VPC) or a transit router (TR) that is associated with an Express Connect router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationAllowedPrefixRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterAssociationAllowedPrefixResponse
func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefixWithOptions(request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedPrefixes) {
		body["AllowedPrefixes"] = request.AllowedPrefixes
	}

	if !dara.IsNil(request.AllowedPrefixesMode) {
		body["AllowedPrefixesMode"] = request.AllowedPrefixesMode
	}

	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.OwnerAccount) {
		body["OwnerAccount"] = request.OwnerAccount
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouterAssociationAllowedPrefix"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the route prefixes of a virtual private cloud (VPC) or a transit router (TR) that is associated with an Express Connect router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationAllowedPrefixRequest
//
// @return ModifyExpressConnectRouterAssociationAllowedPrefixResponse
func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefix(request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterAssociationAllowedPrefixResponse{}
	_body, _err := client.ModifyExpressConnectRouterAssociationAllowedPrefixWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the cross-region forwarding mode of an Express Connect router (ECR).
//
// @param request - ModifyExpressConnectRouterInterRegionTransitModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) ModifyExpressConnectRouterInterRegionTransitModeWithOptions(request *ModifyExpressConnectRouterInterRegionTransitModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.TransitModeList) {
		body["TransitModeList"] = request.TransitModeList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouterInterRegionTransitMode"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the cross-region forwarding mode of an Express Connect router (ECR).
//
// @param request - ModifyExpressConnectRouterInterRegionTransitModeRequest
//
// @return ModifyExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) ModifyExpressConnectRouterInterRegionTransitMode(request *ModifyExpressConnectRouterInterRegionTransitModeRequest) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyExpressConnectRouterInterRegionTransitModeResponse{}
	_body, _err := client.ModifyExpressConnectRouterInterRegionTransitModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name, description, sampling rate, and sampling interval.
//
// @param request - ModifyFlowLogAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFlowLogAttributeResponse
func (client *Client) ModifyFlowLogAttributeWithOptions(request *ModifyFlowLogAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyFlowLogAttributeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.SamplingRate) {
		query["SamplingRate"] = request.SamplingRate
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.FlowLogName) {
		body["FlowLogName"] = request.FlowLogName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFlowLogAttribute"),
		Version:     dara.String("2023-09-01"),
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
// Modifies the name, description, sampling rate, and sampling interval.
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
// Modifies the resource group to which an Express Connect router (ECR) belongs.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.NewResourceGroupId) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2023-09-01"),
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
// Modifies the resource group to which an Express Connect router (ECR) belongs.
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
// Revokes permissions on a virtual private cloud (VPC) or a virtual border router (VBR) from an Express Connect router (ECR) owned by another account.
//
// @param request - RevokeInstanceFromExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromExpressConnectRouterResponse
func (client *Client) RevokeInstanceFromExpressConnectRouterWithOptions(request *RevokeInstanceFromExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	if !dara.IsNil(request.EcrOwnerAliUid) {
		body["EcrOwnerAliUid"] = request.EcrOwnerAliUid
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceRegionId) {
		body["InstanceRegionId"] = request.InstanceRegionId
	}

	if !dara.IsNil(request.InstanceType) {
		body["InstanceType"] = request.InstanceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeInstanceFromExpressConnectRouter"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeInstanceFromExpressConnectRouterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on a virtual private cloud (VPC) or a virtual border router (VBR) from an Express Connect router (ECR) owned by another account.
//
// @param request - RevokeInstanceFromExpressConnectRouterRequest
//
// @return RevokeInstanceFromExpressConnectRouterResponse
func (client *Client) RevokeInstanceFromExpressConnectRouter(request *RevokeInstanceFromExpressConnectRouterRequest) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeInstanceFromExpressConnectRouterResponse{}
	_body, _err := client.RevokeInstanceFromExpressConnectRouterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes the forwarding bandwidth limit between regions for an Express Connect router (ECR).
//
// Description:
//
// Updates are allowed only when the ECR is in the **Active*	- state.
//
// @param request - SynchronizeExpressConnectRouterInterRegionBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeExpressConnectRouterInterRegionBandwidthResponse
func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidthWithOptions(request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest, runtime *dara.RuntimeOptions) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EcrId) {
		body["EcrId"] = request.EcrId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SynchronizeExpressConnectRouterInterRegionBandwidth"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the forwarding bandwidth limit between regions for an Express Connect router (ECR).
//
// Description:
//
// Updates are allowed only when the ECR is in the **Active*	- state.
//
// @param request - SynchronizeExpressConnectRouterInterRegionBandwidthRequest
//
// @return SynchronizeExpressConnectRouterInterRegionBandwidthResponse
func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidth(request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SynchronizeExpressConnectRouterInterRegionBandwidthResponse{}
	_body, _err := client.SynchronizeExpressConnectRouterInterRegionBandwidthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to an Express Connect router (ECR). You can add tags to only one ECR each time you call this operation. You can add multiple tags at a time.
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
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DryRun) {
		body["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2023-09-01"),
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
// Adds tags to an Express Connect router (ECR). You can add tags to only one ECR each time you call this operation. You can add multiple tags at a time.
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
// Removes tags from an Express Connect router (ECR).
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

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		body["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2023-09-01"),
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
// Removes tags from an Express Connect router (ECR).
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
