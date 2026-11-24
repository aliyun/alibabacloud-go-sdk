// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Enables log delivery for a flow log.
//
// @param request - ActivateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateFlowLogResponse
func (client *Client) ActivateFlowLogWithContext(ctx context.Context, request *ActivateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *ActivateFlowLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a child instance (VBR) to an Express Connect Router (ECR).
//
// Description:
//
// Before you invoke the **AttachExpressConnectRouterChildInstance*	- operation to load a child instance (VBR) to an Express Connect Router (ECR), make sure that the Express Connect Router (ECR) is in the **Active*	- state.
//
// @param request - AttachExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachExpressConnectRouterChildInstanceResponse
func (client *Client) AttachExpressConnectRouterChildInstanceWithContext(ctx context.Context, request *AttachExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *AttachExpressConnectRouterChildInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks the Cloud Data Transfer (CDT) service required to add a specified region to an Express Connect Router (ECR).
//
// @param request - CheckAddRegionToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckAddRegionToExpressConnectRouterResponse
func (client *Client) CheckAddRegionToExpressConnectRouterWithContext(ctx context.Context, request *CheckAddRegionToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *CheckAddRegionToExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// After you create an Express Connect Router (ECR) instance, the instance status is **Active**.
//
// @param request - CreateExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressConnectRouterResponse
func (client *Client) CreateExpressConnectRouterWithContext(ctx context.Context, request *CreateExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *CreateExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an Express Connect Router (ECR) with a virtual private cloud (VPC) or a transit router (TR).
//
// @param request - CreateExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExpressConnectRouterAssociationResponse
func (client *Client) CreateExpressConnectRouterAssociationWithContext(ctx context.Context, request *CreateExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *CreateExpressConnectRouterAssociationResponse, _err error) {
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

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a flow log and enables log delivery by default.
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

	if !dara.IsNil(request.TargetSlsRegionId) {
		query["TargetSlsRegionId"] = request.TargetSlsRegionId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an association between an Express Connect Router (ECR) and a route map.
//
// @param request - CreateRouteMapAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRouteMapAssociationResponse
func (client *Client) CreateRouteMapAssociationWithContext(ctx context.Context, request *CreateRouteMapAssociationRequest, runtime *dara.RuntimeOptions) (_result *CreateRouteMapAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionIdList) {
		query["RegionIdList"] = request.RegionIdList
	}

	if !dara.IsNil(request.RouteMapId) {
		query["RouteMapId"] = request.RouteMapId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Action:      dara.String("CreateRouteMapAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRouteMapAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops log delivery for a flow log.
//
// @param request - DeactivateFlowLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeactivateFlowLogResponse
func (client *Client) DeactivateFlowLogWithContext(ctx context.Context, request *DeactivateFlowLogRequest, runtime *dara.RuntimeOptions) (_result *DeactivateFlowLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Express Connect Router (ECR).
//
// Description:
//
// Before you invoke the DeleteExpressConnectRouter operation to delete an Express Connect Router (ECR) instance, note the following:
//
// - Before you delete an Express Connect Router (ECR) instance, you must delete the resources associated with the Express Connect Router (ECR) instance.
//
// - You can delete only Express Connect Router (ECR) instances that are in the **Active*	- state.
//
// @param request - DeleteExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressConnectRouterResponse
func (client *Client) DeleteExpressConnectRouterWithContext(ctx context.Context, request *DeleteExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DeleteExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates an Express Connect Router (ECR) from a virtual private cloud (VPC) or a transit router.
//
// @param request - DeleteExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExpressConnectRouterAssociationResponse
func (client *Client) DeleteExpressConnectRouterAssociationWithContext(ctx context.Context, request *DeleteExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *DeleteExpressConnectRouterAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	if !dara.IsNil(request.FlowLogId) {
		query["FlowLogId"] = request.FlowLogId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the association between an Express Connect Router (ECR) and a route policy.
//
// @param request - DeleteRouteMapAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRouteMapAssociationResponse
func (client *Client) DeleteRouteMapAssociationWithContext(ctx context.Context, request *DeleteRouteMapAssociationRequest, runtime *dara.RuntimeOptions) (_result *DeleteRouteMapAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		query["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Action:      dara.String("DeleteRouteMapAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRouteMapAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the disabled route entry of an Express Connect Router (ECR).
//
// @param request - DescribeDisabledExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDisabledExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeDisabledExpressConnectRouterRouteEntriesWithContext(ctx context.Context, request *DescribeDisabledExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDisabledExpressConnectRouterRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of Express Connect Router (ECR) instances.
//
// @param request - DescribeExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterResponse
func (client *Client) DescribeExpressConnectRouterWithContext(ctx context.Context, request *DescribeExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the history of prefix routes for an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterAllowedPrefixHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterAllowedPrefixHistoryResponse
func (client *Client) DescribeExpressConnectRouterAllowedPrefixHistoryWithContext(ctx context.Context, request *DescribeExpressConnectRouterAllowedPrefixHistoryRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterAllowedPrefixHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual private clouds (VPCs) and transit routers (TRs) associated with an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterAssociationResponse
func (client *Client) DescribeExpressConnectRouterAssociationWithContext(ctx context.Context, request *DescribeExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterAssociationResponse, _err error) {
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

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the virtual border routers (VBRs) associated with an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterChildInstanceResponse
func (client *Client) DescribeExpressConnectRouterChildInstanceWithContext(ctx context.Context, request *DescribeExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterChildInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cross-region forwarding mode of an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterInterRegionTransitModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) DescribeExpressConnectRouterInterRegionTransitModeWithContext(ctx context.Context, request *DescribeExpressConnectRouterInterRegionTransitModeRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions where Express Connect Router (ECR) resources are located.
//
// @param request - DescribeExpressConnectRouterRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterRegionResponse
func (client *Client) DescribeExpressConnectRouterRegionWithContext(ctx context.Context, request *DescribeExpressConnectRouterRegionRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route entries of an Express Connect Router (ECR).
//
// @param request - DescribeExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeExpressConnectRouterRouteEntriesResponse
func (client *Client) DescribeExpressConnectRouterRouteEntriesWithContext(ctx context.Context, request *DescribeExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DescribeExpressConnectRouterRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network instances authorized to an Express Connect Router (ECR).
//
// @param request - DescribeInstanceGrantedToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceGrantedToExpressConnectRouterResponse
func (client *Client) DescribeInstanceGrantedToExpressConnectRouterWithContext(ctx context.Context, request *DescribeInstanceGrantedToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *DescribeInstanceGrantedToExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a virtual border router (VBR) from an Express Connect Router (ECR).
//
// Description:
//
// Before you call the **DetachExpressConnectRouterChildInstance*	- operation to detach a VBR instance from an Express Connect Router (ECR), ensure that the ECR is in the **Active*	- state.
//
// @param request - DetachExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachExpressConnectRouterChildInstanceResponse
func (client *Client) DetachExpressConnectRouterChildInstanceWithContext(ctx context.Context, request *DetachExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *DetachExpressConnectRouterChildInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a route entry of an Express Connect Router (ECR).
//
// @param request - DisableExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableExpressConnectRouterRouteEntriesResponse
func (client *Client) DisableExpressConnectRouterRouteEntriesWithContext(ctx context.Context, request *DisableExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *DisableExpressConnectRouterRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables route entries for an Express Connect Router (ECR).
//
// @param request - EnableExpressConnectRouterRouteEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableExpressConnectRouterRouteEntriesResponse
func (client *Client) EnableExpressConnectRouterRouteEntriesWithContext(ctx context.Context, request *EnableExpressConnectRouterRouteEntriesRequest, runtime *dara.RuntimeOptions) (_result *EnableExpressConnectRouterRouteEntriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Force deletes an Express Connect Router (ECR) and detaches its associated virtual private clouds (VPCs), Transit Routers (TRs), and virtual border routers (VBRs).
//
// Description:
//
// - Force deleting an Express Connect Router (ECR) instance detaches all associated resources. Before you proceed, confirm that this operation does not disrupt your services.
//
// - You can delete only ECR instances that are in the **Active*	- state.
//
// @param request - ForceDeleteExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForceDeleteExpressConnectRouterResponse
func (client *Client) ForceDeleteExpressConnectRouterWithContext(ctx context.Context, request *ForceDeleteExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *ForceDeleteExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants a virtual private cloud (VPC) or virtual border router (VBR) authorization to a cross-account Express Connect Router (ECR).
//
// Description:
//
// To associate a cross-account network instance with an Express Connect Router (ECR) instance, you must grant the network instance authorization to the Express Connect Router (ECR) instance.
//
// @param request - GrantInstanceToExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantInstanceToExpressConnectRouterResponse
func (client *Client) GrantInstanceToExpressConnectRouterWithContext(ctx context.Context, request *GrantInstanceToExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *GrantInstanceToExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions that support Express Connect Router (ECR).
//
// @param request - ListExpressConnectRouterSupportedRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExpressConnectRouterSupportedRegionResponse
func (client *Client) ListExpressConnectRouterSupportedRegionWithContext(ctx context.Context, request *ListExpressConnectRouterSupportedRegionRequest, runtime *dara.RuntimeOptions) (_result *ListExpressConnectRouterSupportedRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.NodeType) {
		body["NodeType"] = request.NodeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the association between an Express Connect Router (ECR) and a route map.
//
// @param request - ListRouteMapAssociationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRouteMapAssociationsResponse
func (client *Client) ListRouteMapAssociationsWithContext(ctx context.Context, request *ListRouteMapAssociationsRequest, runtime *dara.RuntimeOptions) (_result *ListRouteMapAssociationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		query["AssociationId"] = request.AssociationId
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

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Action:      dara.String("ListRouteMapAssociations"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRouteMapAssociationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of tags bound to an Express Connect Router (ECR).
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
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the properties of an Express Connect Router (ECR), such as its name.
//
// Description:
//
// You can modify an Express Connect Router (ECR) instance only when it is in the **Active*	- state.
//
// @param request - ModifyExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterResponse
func (client *Client) ModifyExpressConnectRouterWithContext(ctx context.Context, request *ModifyExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a child instance associated with an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterAssociationResponse
func (client *Client) ModifyExpressConnectRouterAssociationWithContext(ctx context.Context, request *ModifyExpressConnectRouterAssociationRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		body["AssociationId"] = request.AssociationId
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouterAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the prefixes for a virtual private cloud (VPC) or a transit router that is attached to an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterAssociationAllowedPrefixRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterAssociationAllowedPrefixResponse
func (client *Client) ModifyExpressConnectRouterAssociationAllowedPrefixWithContext(ctx context.Context, request *ModifyExpressConnectRouterAssociationAllowedPrefixRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterAssociationAllowedPrefixResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a child instance (VBR) associated with an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterChildInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterChildInstanceResponse
func (client *Client) ModifyExpressConnectRouterChildInstanceWithContext(ctx context.Context, request *ModifyExpressConnectRouterChildInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterChildInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyExpressConnectRouterChildInstance"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyExpressConnectRouterChildInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the inter-region forwarding mode of an Express Connect Router (ECR).
//
// @param request - ModifyExpressConnectRouterInterRegionTransitModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyExpressConnectRouterInterRegionTransitModeResponse
func (client *Client) ModifyExpressConnectRouterInterRegionTransitModeWithContext(ctx context.Context, request *ModifyExpressConnectRouterInterRegionTransitModeRequest, runtime *dara.RuntimeOptions) (_result *ModifyExpressConnectRouterInterRegionTransitModeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name, description, sampling rate, and sampling interval of a flow log.
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

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which an Express Connect Router (ECR) belongs.
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
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the authorization of a virtual private cloud (VPC) or virtual border router (VBR) on a cross-account Express Connect Router (ECR).
//
// @param request - RevokeInstanceFromExpressConnectRouterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeInstanceFromExpressConnectRouterResponse
func (client *Client) RevokeInstanceFromExpressConnectRouterWithContext(ctx context.Context, request *RevokeInstanceFromExpressConnectRouterRequest, runtime *dara.RuntimeOptions) (_result *RevokeInstanceFromExpressConnectRouterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the inter-region bandwidth limits for an Express Connect Router (ECR).
//
// Description:
//
// You can update an ECR instance only when it is in the **Active*	- state.
//
// @param request - SynchronizeExpressConnectRouterInterRegionBandwidthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SynchronizeExpressConnectRouterInterRegionBandwidthResponse
func (client *Client) SynchronizeExpressConnectRouterInterRegionBandwidthWithContext(ctx context.Context, request *SynchronizeExpressConnectRouterInterRegionBandwidthRequest, runtime *dara.RuntimeOptions) (_result *SynchronizeExpressConnectRouterInterRegionBandwidthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds tags to an Express Connect Router (ECR). Each call to this operation binds tags to a single ECR instance, but you can bind multiple tags at a time.
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
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds tags from an Express Connect Router (ECR).
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
	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the association between an Express Connect Router (ECR) and a routing policy.
//
// @param request - UpdateRouteMapAssociationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRouteMapAssociationResponse
func (client *Client) UpdateRouteMapAssociationWithContext(ctx context.Context, request *UpdateRouteMapAssociationRequest, runtime *dara.RuntimeOptions) (_result *UpdateRouteMapAssociationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociationId) {
		query["AssociationId"] = request.AssociationId
	}

	if !dara.IsNil(request.RegionIdList) {
		query["RegionIdList"] = request.RegionIdList
	}

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
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
		Action:      dara.String("UpdateRouteMapAssociation"),
		Version:     dara.String("2023-09-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRouteMapAssociationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
