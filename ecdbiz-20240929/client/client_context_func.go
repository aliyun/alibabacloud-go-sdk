// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Activates the commercial edition rights under the current account. First, query the rights by calling ListRights, and then select the target region, office network, image, and user. A Success response indicates that the request is processed. To check the resource creation result, query the RightsActiveStatus, RightsActiveFailureCause, and RightsResourceId parameters returned by ListRights.
//
// Description:
//
// Activates the commercial edition rights under the current account. First, query the rights by calling ListRights, and then select the target region, office network, image, and user. A Success response indicates that the request is processed. To check the resource creation result, query the RightsActiveStatus, RightsActiveFailureCause, and RightsResourceId parameters returned by ListRights.
//
// @param request - ActiveRightsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveRightsResponse
func (client *Client) ActiveRightsWithContext(ctx context.Context, request *ActiveRightsRequest, runtime *dara.RuntimeOptions) (_result *ActiveRightsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataDiskSize) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.PolicyGroupId) {
		query["PolicyGroupId"] = request.PolicyGroupId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RootDiskSize) {
		query["RootDiskSize"] = request.RootDiskSize
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	if !dara.IsNil(request.WyidList) {
		query["WyidList"] = request.WyidList
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.CustomImageId) {
		body["CustomImageId"] = request.CustomImageId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Hibernate) {
		body["Hibernate"] = request.Hibernate
	}

	if !dara.IsNil(request.OfficeSiteId) {
		body["OfficeSiteId"] = request.OfficeSiteId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.RightsAssignedUser) {
		bodyFlat["RightsAssignedUser"] = request.RightsAssignedUser
	}

	if !dara.IsNil(request.RightsGpuSpec) {
		body["RightsGpuSpec"] = request.RightsGpuSpec
	}

	if !dara.IsNil(request.RightsIds) {
		bodyFlat["RightsIds"] = request.RightsIds
	}

	if !dara.IsNil(request.RightsResourceArea) {
		body["RightsResourceArea"] = request.RightsResourceArea
	}

	if !dara.IsNil(request.RightsResourceImageOs) {
		body["RightsResourceImageOs"] = request.RightsResourceImageOs
	}

	if !dara.IsNil(request.RightsResourceRegion) {
		body["RightsResourceRegion"] = request.RightsResourceRegion
	}

	if !dara.IsNil(request.RightsResourceType) {
		body["RightsResourceType"] = request.RightsResourceType
	}

	if !dara.IsNil(request.RightsSpec) {
		body["RightsSpec"] = request.RightsSpec
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.VSwitchId) {
		body["VSwitchId"] = request.VSwitchId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActiveRights"),
		Version:     dara.String("2024-09-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveRightsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the commercial edition rights accessible to the current caller identity by using paging. Use this operation to select rights before activation. After you call ActiveRights, you can query the processing results by using RightsActiveStatus, RightsActiveFailureCause, and RightsResourceId.
//
// Description:
//
// Queries the commercial edition rights accessible to the current caller identity by using paging. Use this operation to select rights before activation. After you call ActiveRights, you can query the processing results by using RightsActiveStatus, RightsActiveFailureCause, and RightsResourceId.
//
// @param request - ListRightsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRightsResponse
func (client *Client) ListRightsWithContext(ctx context.Context, request *ListRightsRequest, runtime *dara.RuntimeOptions) (_result *ListRightsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessSceneCode) {
		query["BusinessSceneCode"] = request.BusinessSceneCode
	}

	if !dara.IsNil(request.RightsIds) {
		query["RightsIds"] = request.RightsIds
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AiDesktop) {
		body["AiDesktop"] = request.AiDesktop
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.OrderType) {
		body["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.Overseas) {
		body["Overseas"] = request.Overseas
	}

	if !dara.IsNil(request.RightsActiveStatus) {
		body["RightsActiveStatus"] = request.RightsActiveStatus
	}

	if !dara.IsNil(request.RightsAssignedUser) {
		body["RightsAssignedUser"] = request.RightsAssignedUser
	}

	if !dara.IsNil(request.RightsPeriod) {
		body["RightsPeriod"] = request.RightsPeriod
	}

	if !dara.IsNil(request.RightsResourceId) {
		body["RightsResourceId"] = request.RightsResourceId
	}

	if !dara.IsNil(request.RightsResourceType) {
		body["RightsResourceType"] = request.RightsResourceType
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.RightsResourceTypes) {
		bodyFlat["RightsResourceTypes"] = request.RightsResourceTypes
	}

	if !dara.IsNil(request.RightsSpec) {
		body["RightsSpec"] = request.RightsSpec
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRights"),
		Version:     dara.String("2024-09-29"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRightsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
