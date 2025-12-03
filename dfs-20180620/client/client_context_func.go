// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 挂载VSC挂载点
//
// Description:
//
// ***
//
// @param tmpReq - AttachVscMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachVscMountPointResponse
func (client *Client) AttachVscMountPointWithContext(ctx context.Context, tmpReq *AttachVscMountPointRequest, runtime *dara.RuntimeOptions) (_result *AttachVscMountPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachVscMountPointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VscIds) {
		request.VscIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VscIds, dara.String("VscIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.MountPointId) {
		query["MountPointId"] = request.MountPointId
	}

	if !dara.IsNil(request.UseAssumeRoleChkServerPerm) {
		query["UseAssumeRoleChkServerPerm"] = request.UseAssumeRoleChkServerPerm
	}

	if !dara.IsNil(request.VscIdsShrink) {
		query["VscIds"] = request.VscIdsShrink
	}

	if !dara.IsNil(request.VscName) {
		query["VscName"] = request.VscName
	}

	if !dara.IsNil(request.VscType) {
		query["VscType"] = request.VscType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachVscMountPoint"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachVscMountPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量挂载VSC挂载点
//
// @param tmpReq - AttachVscToMountPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachVscToMountPointsResponse
func (client *Client) AttachVscToMountPointsWithContext(ctx context.Context, tmpReq *AttachVscToMountPointsRequest, runtime *dara.RuntimeOptions) (_result *AttachVscToMountPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachVscToMountPointsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AttachInfos) {
		request.AttachInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AttachInfos, dara.String("AttachInfos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachInfosShrink) {
		query["AttachInfos"] = request.AttachInfosShrink
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.UseAssumeRoleChkServerPerm) {
		query["UseAssumeRoleChkServerPerm"] = request.UseAssumeRoleChkServerPerm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachVscToMountPoints"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachVscToMountPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindVscMountPointAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindVscMountPointAliasResponse
func (client *Client) BindVscMountPointAliasWithContext(ctx context.Context, request *BindVscMountPointAliasRequest, runtime *dara.RuntimeOptions) (_result *BindVscMountPointAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasPrefix) {
		query["AliasPrefix"] = request.AliasPrefix
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MountPointId) {
		query["MountPointId"] = request.MountPointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindVscMountPointAlias"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindVscMountPointAliasResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessGroupResponse
func (client *Client) CreateAccessGroupWithContext(ctx context.Context, request *CreateAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessGroup"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessRuleResponse
func (client *Client) CreateAccessRuleWithContext(ctx context.Context, request *CreateAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.NetworkSegment) {
		query["NetworkSegment"] = request.NetworkSegment
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RWAccessType) {
		query["RWAccessType"] = request.RWAccessType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccessRule"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileSystemResponse
func (client *Client) CreateFileSystemWithContext(ctx context.Context, request *CreateFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataRedundancyType) {
		query["DataRedundancyType"] = request.DataRedundancyType
	}

	if !dara.IsNil(request.DedicatedClusterId) {
		query["DedicatedClusterId"] = request.DedicatedClusterId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemName) {
		query["FileSystemName"] = request.FileSystemName
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.PartitionNumber) {
		query["PartitionNumber"] = request.PartitionNumber
	}

	if !dara.IsNil(request.ProtocolType) {
		query["ProtocolType"] = request.ProtocolType
	}

	if !dara.IsNil(request.ProvisionedThroughputInMiBps) {
		query["ProvisionedThroughputInMiBps"] = request.ProvisionedThroughputInMiBps
	}

	if !dara.IsNil(request.SpaceCapacity) {
		query["SpaceCapacity"] = request.SpaceCapacity
	}

	if !dara.IsNil(request.StorageSetName) {
		query["StorageSetName"] = request.StorageSetName
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.ThroughputMode) {
		query["ThroughputMode"] = request.ThroughputMode
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFileSystem"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMountPointResponse
func (client *Client) CreateMountPointWithContext(ctx context.Context, request *CreateMountPointRequest, runtime *dara.RuntimeOptions) (_result *CreateMountPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.UsePerformanceMode) {
		query["UsePerformanceMode"] = request.UsePerformanceMode
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
		Action:      dara.String("CreateMountPoint"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMountPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateQosPolicy
//
// @param tmpReq - CreateQosPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQosPolicyResponse
func (client *Client) CreateQosPolicyWithContext(ctx context.Context, tmpReq *CreateQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateQosPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateQosPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FlowIds) {
		request.FlowIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FlowIds, dara.String("FlowIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ReqTags) {
		request.ReqTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReqTags, dara.String("ReqTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ZoneIds) {
		request.ZoneIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ZoneIds, dara.String("ZoneIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FederationId) {
		query["FederationId"] = request.FederationId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FlowIdsShrink) {
		query["FlowIds"] = request.FlowIdsShrink
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MaxIOBandWidth) {
		query["MaxIOBandWidth"] = request.MaxIOBandWidth
	}

	if !dara.IsNil(request.MaxIOps) {
		query["MaxIOps"] = request.MaxIOps
	}

	if !dara.IsNil(request.MaxMetaQps) {
		query["MaxMetaQps"] = request.MaxMetaQps
	}

	if !dara.IsNil(request.ReqTagsShrink) {
		query["ReqTags"] = request.ReqTagsShrink
	}

	if !dara.IsNil(request.ZoneIdsShrink) {
		query["ZoneIds"] = request.ZoneIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQosPolicy"),
		Version:     dara.String("2018-06-20"),
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
// 创建 ugo
//
// @param tmpReq - CreateUserGroupsMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupsMappingResponse
func (client *Client) CreateUserGroupsMappingWithContext(ctx context.Context, tmpReq *CreateUserGroupsMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupsMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUserGroupsMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GroupNames) {
		request.GroupNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupNames, dara.String("GroupNames"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserGroupsMapping"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserGroupsMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - CreateVscMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVscMountPointResponse
func (client *Client) CreateVscMountPointWithContext(ctx context.Context, tmpReq *CreateVscMountPointRequest, runtime *dara.RuntimeOptions) (_result *CreateVscMountPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateVscMountPointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVscMountPoint"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVscMountPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessGroupResponse
func (client *Client) DeleteAccessGroupWithContext(ctx context.Context, request *DeleteAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessGroup"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessRuleResponse
func (client *Client) DeleteAccessRuleWithContext(ctx context.Context, request *DeleteAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccessRule"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileSystemResponse
func (client *Client) DeleteFileSystemWithContext(ctx context.Context, request *DeleteFileSystemRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFileSystem"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMountPointResponse
func (client *Client) DeleteMountPointWithContext(ctx context.Context, request *DeleteMountPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteMountPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MountPointId) {
		query["MountPointId"] = request.MountPointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMountPoint"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMountPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteQosPolicy
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
	if !dara.IsNil(request.FederationId) {
		query["FederationId"] = request.FederationId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.QosPolicyId) {
		query["QosPolicyId"] = request.QosPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQosPolicy"),
		Version:     dara.String("2018-06-20"),
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
// 删除 ugo
//
// @param tmpReq - DeleteUserGroupsMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupsMappingResponse
func (client *Client) DeleteUserGroupsMappingWithContext(ctx context.Context, tmpReq *DeleteUserGroupsMappingRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupsMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteUserGroupsMappingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.GroupNames) {
		request.GroupNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupNames, dara.String("GroupNames"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroupsMapping"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserGroupsMappingResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteVscMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVscMountPointResponse
func (client *Client) DeleteVscMountPointWithContext(ctx context.Context, request *DeleteVscMountPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteVscMountPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MountPointId) {
		query["MountPointId"] = request.MountPointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVscMountPoint"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVscMountPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量挂载VSC挂载点
//
// @param tmpReq - DescribeMountPointsVscAttachInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMountPointsVscAttachInfoResponse
func (client *Client) DescribeMountPointsVscAttachInfoWithContext(ctx context.Context, tmpReq *DescribeMountPointsVscAttachInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountPointsVscAttachInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeMountPointsVscAttachInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueryInfos) {
		request.QueryInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueryInfos, dara.String("QueryInfos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.QueryInfosShrink) {
		query["QueryInfos"] = request.QueryInfosShrink
	}

	if !dara.IsNil(request.UseAssumeRoleChkServerPerm) {
		query["UseAssumeRoleChkServerPerm"] = request.UseAssumeRoleChkServerPerm
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMountPointsVscAttachInfo"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMountPointsVscAttachInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2018-06-20"),
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
// 查询VSC挂载信息
//
// @param request - DescribeVscMountPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVscMountPointsResponse
func (client *Client) DescribeVscMountPointsWithContext(ctx context.Context, request *DescribeVscMountPointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVscMountPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MountPointId) {
		query["MountPointId"] = request.MountPointId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VscId) {
		query["VscId"] = request.VscId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVscMountPoints"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVscMountPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param tmpReq - DetachVscMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachVscMountPointResponse
func (client *Client) DetachVscMountPointWithContext(ctx context.Context, tmpReq *DetachVscMountPointRequest, runtime *dara.RuntimeOptions) (_result *DetachVscMountPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachVscMountPointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InstanceIds) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, dara.String("InstanceIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VscIds) {
		request.VscIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VscIds, dara.String("VscIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.InstanceIdsShrink) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !dara.IsNil(request.MountPointId) {
		query["MountPointId"] = request.MountPointId
	}

	if !dara.IsNil(request.UseAssumeRoleChkServerPerm) {
		query["UseAssumeRoleChkServerPerm"] = request.UseAssumeRoleChkServerPerm
	}

	if !dara.IsNil(request.VscIdsShrink) {
		query["VscIds"] = request.VscIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachVscMountPoint"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachVscMountPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessGroupResponse
func (client *Client) GetAccessGroupWithContext(ctx context.Context, request *GetAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *GetAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessGroup"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessRuleResponse
func (client *Client) GetAccessRuleWithContext(ctx context.Context, request *GetAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessRule"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileSystemResponse
func (client *Client) GetFileSystemWithContext(ctx context.Context, request *GetFileSystemRequest, runtime *dara.RuntimeOptions) (_result *GetFileSystemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileSystem"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMountPointResponse
func (client *Client) GetMountPointWithContext(ctx context.Context, request *GetMountPointRequest, runtime *dara.RuntimeOptions) (_result *GetMountPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MountPointId) {
		query["MountPointId"] = request.MountPointId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMountPoint"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMountPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionResponse
func (client *Client) GetRegionWithContext(ctx context.Context, request *GetRegionRequest, runtime *dara.RuntimeOptions) (_result *GetRegionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegion"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAccessGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessGroupsResponse
func (client *Client) ListAccessGroupsWithContext(ctx context.Context, request *ListAccessGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListAccessGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.StartOffset) {
		query["StartOffset"] = request.StartOffset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessGroups"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessGroupsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAccessRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessRulesResponse
func (client *Client) ListAccessRulesWithContext(ctx context.Context, request *ListAccessRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAccessRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.StartOffset) {
		query["StartOffset"] = request.StartOffset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccessRules"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccessRulesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Federation
//
// @param request - ListFederationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFederationsResponse
func (client *Client) ListFederationsWithContext(ctx context.Context, request *ListFederationsRequest, runtime *dara.RuntimeOptions) (_result *ListFederationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederationId) {
		query["FederationId"] = request.FederationId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFederations"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFederationsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileSystemsResponse
func (client *Client) ListFileSystemsWithContext(ctx context.Context, request *ListFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *ListFileSystemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.StartOffset) {
		query["StartOffset"] = request.StartOffset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFileSystems"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFileSystemsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMountPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMountPointsResponse
func (client *Client) ListMountPointsWithContext(ctx context.Context, request *ListMountPointsRequest, runtime *dara.RuntimeOptions) (_result *ListMountPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.StartOffset) {
		query["StartOffset"] = request.StartOffset
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMountPoints"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMountPointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询QosPolicies
//
// @param request - ListQosPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQosPoliciesResponse
func (client *Client) ListQosPoliciesWithContext(ctx context.Context, request *ListQosPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListQosPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FederationId) {
		query["FederationId"] = request.FederationId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQosPolicies"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQosPoliciesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// list ugm
//
// @param request - ListUserGroupsMappingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsMappingsResponse
func (client *Client) ListUserGroupsMappingsWithContext(ctx context.Context, request *ListUserGroupsMappingsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsMappingsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilesystemId) {
		query["FilesystemId"] = request.FilesystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroupsMappings"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsMappingsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessGroupResponse
func (client *Client) ModifyAccessGroupWithContext(ctx context.Context, request *ModifyAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.AccessGroupName) {
		query["AccessGroupName"] = request.AccessGroupName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccessGroup"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessRuleResponse
func (client *Client) ModifyAccessRuleWithContext(ctx context.Context, request *ModifyAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.AccessRuleId) {
		query["AccessRuleId"] = request.AccessRuleId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RWAccessType) {
		query["RWAccessType"] = request.RWAccessType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAccessRule"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFileSystemResponse
func (client *Client) ModifyFileSystemWithContext(ctx context.Context, request *ModifyFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
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

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.FileSystemName) {
		query["FileSystemName"] = request.FileSystemName
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.ProvisionedThroughputInMiBps) {
		query["ProvisionedThroughputInMiBps"] = request.ProvisionedThroughputInMiBps
	}

	if !dara.IsNil(request.SpaceCapacity) {
		query["SpaceCapacity"] = request.SpaceCapacity
	}

	if !dara.IsNil(request.ThroughputMode) {
		query["ThroughputMode"] = request.ThroughputMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyFileSystem"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMountPointResponse
func (client *Client) ModifyMountPointWithContext(ctx context.Context, request *ModifyMountPointRequest, runtime *dara.RuntimeOptions) (_result *ModifyMountPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessGroupId) {
		query["AccessGroupId"] = request.AccessGroupId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MountPointId) {
		query["MountPointId"] = request.MountPointId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMountPoint"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMountPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ModifyQosPolicy
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

	if !dara.IsNil(request.FederationId) {
		query["FederationId"] = request.FederationId
	}

	if !dara.IsNil(request.FileSystemId) {
		query["FileSystemId"] = request.FileSystemId
	}

	if !dara.IsNil(request.InputRegionId) {
		query["InputRegionId"] = request.InputRegionId
	}

	if !dara.IsNil(request.MaxIOBandWidth) {
		query["MaxIOBandWidth"] = request.MaxIOBandWidth
	}

	if !dara.IsNil(request.MaxIOps) {
		query["MaxIOps"] = request.MaxIOps
	}

	if !dara.IsNil(request.MaxMetaQps) {
		query["MaxMetaQps"] = request.MaxMetaQps
	}

	if !dara.IsNil(request.QosPolicyId) {
		query["QosPolicyId"] = request.QosPolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyQosPolicy"),
		Version:     dara.String("2018-06-20"),
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
