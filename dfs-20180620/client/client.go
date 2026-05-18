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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dfs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
func (client *Client) AttachVscMountPointWithOptions(tmpReq *AttachVscMountPointRequest, runtime *dara.RuntimeOptions) (_result *AttachVscMountPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 挂载VSC挂载点
//
// Description:
//
// ***
//
// @param request - AttachVscMountPointRequest
//
// @return AttachVscMountPointResponse
func (client *Client) AttachVscMountPoint(request *AttachVscMountPointRequest) (_result *AttachVscMountPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachVscMountPointResponse{}
	_body, _err := client.AttachVscMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AttachVscToMountPointsWithOptions(tmpReq *AttachVscToMountPointsRequest, runtime *dara.RuntimeOptions) (_result *AttachVscToMountPointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - AttachVscToMountPointsRequest
//
// @return AttachVscToMountPointsResponse
func (client *Client) AttachVscToMountPoints(request *AttachVscToMountPointsRequest) (_result *AttachVscToMountPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachVscToMountPointsResponse{}
	_body, _err := client.AttachVscToMountPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - BindVscMountPointAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindVscMountPointAliasResponse
func (client *Client) BindVscMountPointAliasWithOptions(request *BindVscMountPointAliasRequest, runtime *dara.RuntimeOptions) (_result *BindVscMountPointAliasResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindVscMountPointAliasRequest
//
// @return BindVscMountPointAliasResponse
func (client *Client) BindVscMountPointAlias(request *BindVscMountPointAliasRequest) (_result *BindVscMountPointAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindVscMountPointAliasResponse{}
	_body, _err := client.BindVscMountPointAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessGroupResponse
func (client *Client) CreateAccessGroupWithOptions(request *CreateAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAccessGroupRequest
//
// @return CreateAccessGroupResponse
func (client *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (_result *CreateAccessGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessGroupResponse{}
	_body, _err := client.CreateAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccessRuleResponse
func (client *Client) CreateAccessRuleWithOptions(request *CreateAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAccessRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAccessRuleRequest
//
// @return CreateAccessRuleResponse
func (client *Client) CreateAccessRule(request *CreateAccessRuleRequest) (_result *CreateAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccessRuleResponse{}
	_body, _err := client.CreateAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileSystemResponse
func (client *Client) CreateFileSystemWithOptions(request *CreateFileSystemRequest, runtime *dara.RuntimeOptions) (_result *CreateFileSystemResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateFileSystemRequest
//
// @return CreateFileSystemResponse
func (client *Client) CreateFileSystem(request *CreateFileSystemRequest) (_result *CreateFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFileSystemResponse{}
	_body, _err := client.CreateFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMountPointResponse
func (client *Client) CreateMountPointWithOptions(request *CreateMountPointRequest, runtime *dara.RuntimeOptions) (_result *CreateMountPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMountPointRequest
//
// @return CreateMountPointResponse
func (client *Client) CreateMountPoint(request *CreateMountPointRequest) (_result *CreateMountPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMountPointResponse{}
	_body, _err := client.CreateMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateQosPolicyWithOptions(tmpReq *CreateQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateQosPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateQosPolicyRequest
//
// @return CreateQosPolicyResponse
func (client *Client) CreateQosPolicy(request *CreateQosPolicyRequest) (_result *CreateQosPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQosPolicyResponse{}
	_body, _err := client.CreateQosPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateUserGroupsMappingWithOptions(tmpReq *CreateUserGroupsMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupsMappingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreateUserGroupsMappingRequest
//
// @return CreateUserGroupsMappingResponse
func (client *Client) CreateUserGroupsMapping(request *CreateUserGroupsMappingRequest) (_result *CreateUserGroupsMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserGroupsMappingResponse{}
	_body, _err := client.CreateUserGroupsMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - CreateVscMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVscMountPointResponse
func (client *Client) CreateVscMountPointWithOptions(tmpReq *CreateVscMountPointRequest, runtime *dara.RuntimeOptions) (_result *CreateVscMountPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateVscMountPointRequest
//
// @return CreateVscMountPointResponse
func (client *Client) CreateVscMountPoint(request *CreateVscMountPointRequest) (_result *CreateVscMountPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVscMountPointResponse{}
	_body, _err := client.CreateVscMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessGroupResponse
func (client *Client) DeleteAccessGroupWithOptions(request *DeleteAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAccessGroupRequest
//
// @return DeleteAccessGroupResponse
func (client *Client) DeleteAccessGroup(request *DeleteAccessGroupRequest) (_result *DeleteAccessGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessGroupResponse{}
	_body, _err := client.DeleteAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccessRuleResponse
func (client *Client) DeleteAccessRuleWithOptions(request *DeleteAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccessRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAccessRuleRequest
//
// @return DeleteAccessRuleResponse
func (client *Client) DeleteAccessRule(request *DeleteAccessRuleRequest) (_result *DeleteAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccessRuleResponse{}
	_body, _err := client.DeleteAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileSystemResponse
func (client *Client) DeleteFileSystemWithOptions(request *DeleteFileSystemRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileSystemResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFileSystemRequest
//
// @return DeleteFileSystemResponse
func (client *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (_result *DeleteFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFileSystemResponse{}
	_body, _err := client.DeleteFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMountPointResponse
func (client *Client) DeleteMountPointWithOptions(request *DeleteMountPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteMountPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMountPointRequest
//
// @return DeleteMountPointResponse
func (client *Client) DeleteMountPoint(request *DeleteMountPointRequest) (_result *DeleteMountPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMountPointResponse{}
	_body, _err := client.DeleteMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteQosPolicyWithOptions(request *DeleteQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteQosPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteQosPolicyResponse
func (client *Client) DeleteQosPolicy(request *DeleteQosPolicyRequest) (_result *DeleteQosPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQosPolicyResponse{}
	_body, _err := client.DeleteQosPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteUserGroupsMappingWithOptions(tmpReq *DeleteUserGroupsMappingRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupsMappingResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DeleteUserGroupsMappingRequest
//
// @return DeleteUserGroupsMappingResponse
func (client *Client) DeleteUserGroupsMapping(request *DeleteUserGroupsMappingRequest) (_result *DeleteUserGroupsMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserGroupsMappingResponse{}
	_body, _err := client.DeleteUserGroupsMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteVscMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVscMountPointResponse
func (client *Client) DeleteVscMountPointWithOptions(request *DeleteVscMountPointRequest, runtime *dara.RuntimeOptions) (_result *DeleteVscMountPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteVscMountPointRequest
//
// @return DeleteVscMountPointResponse
func (client *Client) DeleteVscMountPoint(request *DeleteVscMountPointRequest) (_result *DeleteVscMountPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVscMountPointResponse{}
	_body, _err := client.DeleteVscMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeMountPointsVscAttachInfoWithOptions(tmpReq *DescribeMountPointsVscAttachInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeMountPointsVscAttachInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - DescribeMountPointsVscAttachInfoRequest
//
// @return DescribeMountPointsVscAttachInfoResponse
func (client *Client) DescribeMountPointsVscAttachInfo(request *DescribeMountPointsVscAttachInfoRequest) (_result *DescribeMountPointsVscAttachInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMountPointsVscAttachInfoResponse{}
	_body, _err := client.DescribeMountPointsVscAttachInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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
// 查询VSC挂载信息
//
// @param request - DescribeVscMountPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVscMountPointsResponse
func (client *Client) DescribeVscMountPointsWithOptions(request *DescribeVscMountPointsRequest, runtime *dara.RuntimeOptions) (_result *DescribeVscMountPointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeVscMountPointsResponse
func (client *Client) DescribeVscMountPoints(request *DescribeVscMountPointsRequest) (_result *DescribeVscMountPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVscMountPointsResponse{}
	_body, _err := client.DescribeVscMountPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量卸载VSC 挂载点
//
// @param tmpReq - DetachVscFromMountPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachVscFromMountPointsResponse
func (client *Client) DetachVscFromMountPointsWithOptions(tmpReq *DetachVscFromMountPointsRequest, runtime *dara.RuntimeOptions) (_result *DetachVscFromMountPointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachVscFromMountPointsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DetachInfos) {
		request.DetachInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DetachInfos, dara.String("DetachInfos"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DetachInfosShrink) {
		query["DetachInfos"] = request.DetachInfosShrink
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
		Action:      dara.String("DetachVscFromMountPoints"),
		Version:     dara.String("2018-06-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachVscFromMountPointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量卸载VSC 挂载点
//
// @param request - DetachVscFromMountPointsRequest
//
// @return DetachVscFromMountPointsResponse
func (client *Client) DetachVscFromMountPoints(request *DetachVscFromMountPointsRequest) (_result *DetachVscFromMountPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachVscFromMountPointsResponse{}
	_body, _err := client.DetachVscFromMountPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - DetachVscMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachVscMountPointResponse
func (client *Client) DetachVscMountPointWithOptions(tmpReq *DetachVscMountPointRequest, runtime *dara.RuntimeOptions) (_result *DetachVscMountPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DetachVscMountPointRequest
//
// @return DetachVscMountPointResponse
func (client *Client) DetachVscMountPoint(request *DetachVscMountPointRequest) (_result *DetachVscMountPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachVscMountPointResponse{}
	_body, _err := client.DetachVscMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessGroupResponse
func (client *Client) GetAccessGroupWithOptions(request *GetAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *GetAccessGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAccessGroupRequest
//
// @return GetAccessGroupResponse
func (client *Client) GetAccessGroup(request *GetAccessGroupRequest) (_result *GetAccessGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessGroupResponse{}
	_body, _err := client.GetAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessRuleResponse
func (client *Client) GetAccessRuleWithOptions(request *GetAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAccessRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAccessRuleRequest
//
// @return GetAccessRuleResponse
func (client *Client) GetAccessRule(request *GetAccessRuleRequest) (_result *GetAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccessRuleResponse{}
	_body, _err := client.GetAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileSystemResponse
func (client *Client) GetFileSystemWithOptions(request *GetFileSystemRequest, runtime *dara.RuntimeOptions) (_result *GetFileSystemResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetFileSystemRequest
//
// @return GetFileSystemResponse
func (client *Client) GetFileSystem(request *GetFileSystemRequest) (_result *GetFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileSystemResponse{}
	_body, _err := client.GetFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMountPointResponse
func (client *Client) GetMountPointWithOptions(request *GetMountPointRequest, runtime *dara.RuntimeOptions) (_result *GetMountPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMountPointRequest
//
// @return GetMountPointResponse
func (client *Client) GetMountPoint(request *GetMountPointRequest) (_result *GetMountPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMountPointResponse{}
	_body, _err := client.GetMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegionResponse
func (client *Client) GetRegionWithOptions(request *GetRegionRequest, runtime *dara.RuntimeOptions) (_result *GetRegionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRegionRequest
//
// @return GetRegionResponse
func (client *Client) GetRegion(request *GetRegionRequest) (_result *GetRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRegionResponse{}
	_body, _err := client.GetRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAccessGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessGroupsResponse
func (client *Client) ListAccessGroupsWithOptions(request *ListAccessGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListAccessGroupsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAccessGroupsRequest
//
// @return ListAccessGroupsResponse
func (client *Client) ListAccessGroups(request *ListAccessGroupsRequest) (_result *ListAccessGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccessGroupsResponse{}
	_body, _err := client.ListAccessGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAccessRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccessRulesResponse
func (client *Client) ListAccessRulesWithOptions(request *ListAccessRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAccessRulesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAccessRulesRequest
//
// @return ListAccessRulesResponse
func (client *Client) ListAccessRules(request *ListAccessRulesRequest) (_result *ListAccessRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccessRulesResponse{}
	_body, _err := client.ListAccessRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListFederationsWithOptions(request *ListFederationsRequest, runtime *dara.RuntimeOptions) (_result *ListFederationsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListFederationsResponse
func (client *Client) ListFederations(request *ListFederationsRequest) (_result *ListFederationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFederationsResponse{}
	_body, _err := client.ListFederationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListFileSystemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileSystemsResponse
func (client *Client) ListFileSystemsWithOptions(request *ListFileSystemsRequest, runtime *dara.RuntimeOptions) (_result *ListFileSystemsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListFileSystemsRequest
//
// @return ListFileSystemsResponse
func (client *Client) ListFileSystems(request *ListFileSystemsRequest) (_result *ListFileSystemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFileSystemsResponse{}
	_body, _err := client.ListFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListMountPointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMountPointsResponse
func (client *Client) ListMountPointsWithOptions(request *ListMountPointsRequest, runtime *dara.RuntimeOptions) (_result *ListMountPointsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListMountPointsRequest
//
// @return ListMountPointsResponse
func (client *Client) ListMountPoints(request *ListMountPointsRequest) (_result *ListMountPointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMountPointsResponse{}
	_body, _err := client.ListMountPointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListQosPoliciesWithOptions(request *ListQosPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListQosPoliciesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListQosPoliciesResponse
func (client *Client) ListQosPolicies(request *ListQosPoliciesRequest) (_result *ListQosPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQosPoliciesResponse{}
	_body, _err := client.ListQosPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListUserGroupsMappingsWithOptions(request *ListUserGroupsMappingsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsMappingsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListUserGroupsMappingsResponse
func (client *Client) ListUserGroupsMappings(request *ListUserGroupsMappingsRequest) (_result *ListUserGroupsMappingsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsMappingsResponse{}
	_body, _err := client.ListUserGroupsMappingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAccessGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessGroupResponse
func (client *Client) ModifyAccessGroupWithOptions(request *ModifyAccessGroupRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessGroupResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAccessGroupRequest
//
// @return ModifyAccessGroupResponse
func (client *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (_result *ModifyAccessGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccessGroupResponse{}
	_body, _err := client.ModifyAccessGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAccessRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAccessRuleResponse
func (client *Client) ModifyAccessRuleWithOptions(request *ModifyAccessRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyAccessRuleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyAccessRuleRequest
//
// @return ModifyAccessRuleResponse
func (client *Client) ModifyAccessRule(request *ModifyAccessRuleRequest) (_result *ModifyAccessRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAccessRuleResponse{}
	_body, _err := client.ModifyAccessRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyFileSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFileSystemResponse
func (client *Client) ModifyFileSystemWithOptions(request *ModifyFileSystemRequest, runtime *dara.RuntimeOptions) (_result *ModifyFileSystemResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyFileSystemRequest
//
// @return ModifyFileSystemResponse
func (client *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (_result *ModifyFileSystemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyFileSystemResponse{}
	_body, _err := client.ModifyFileSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyMountPointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMountPointResponse
func (client *Client) ModifyMountPointWithOptions(request *ModifyMountPointRequest, runtime *dara.RuntimeOptions) (_result *ModifyMountPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyMountPointRequest
//
// @return ModifyMountPointResponse
func (client *Client) ModifyMountPoint(request *ModifyMountPointRequest) (_result *ModifyMountPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMountPointResponse{}
	_body, _err := client.ModifyMountPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ModifyQosPolicyWithOptions(request *ModifyQosPolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyQosPolicyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ModifyQosPolicyResponse
func (client *Client) ModifyQosPolicy(request *ModifyQosPolicyRequest) (_result *ModifyQosPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyQosPolicyResponse{}
	_body, _err := client.ModifyQosPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
