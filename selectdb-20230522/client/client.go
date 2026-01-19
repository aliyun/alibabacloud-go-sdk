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
	client.Endpoint, _err = client.GetEndpoint(dara.String("selectdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Applies for a public endpoint for an ApsaraDB for SelectDB instance.
//
// @param request - AllocateInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocateInstancePublicConnectionResponse
func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionStringPrefix) {
		query["ConnectionStringPrefix"] = request.ConnectionStringPrefix
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
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
		Action:      dara.String("AllocateInstancePublicConnection"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a public endpoint for an ApsaraDB for SelectDB instance.
//
// @param request - AllocateInstancePublicConnectionRequest
//
// @return AllocateInstancePublicConnectionResponse
func (client *Client) AllocateInstancePublicConnection(request *AllocateInstancePublicConnectionRequest) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.AllocateInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 资源转组
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源转组
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SelectDB实例创建前检查
//
// @param request - CheckCreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckCreateDBInstanceResponse
func (client *Client) CheckCreateDBInstanceWithOptions(request *CheckCreateDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CheckCreateDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheSize) {
		query["CacheSize"] = request.CacheSize
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
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
		Action:      dara.String("CheckCreateDBInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckCreateDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SelectDB实例创建前检查
//
// @param request - CheckCreateDBInstanceRequest
//
// @return CheckCreateDBInstanceResponse
func (client *Client) CheckCreateDBInstance(request *CheckCreateDBInstanceRequest) (_result *CheckCreateDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckCreateDBInstanceResponse{}
	_body, _err := client.CheckCreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 判断指定 IP 是否已经存在于网络白名单组
//
// @param request - CheckIpExistsInSecurityIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckIpExistsInSecurityIpListResponse
func (client *Client) CheckIpExistsInSecurityIpListWithOptions(request *CheckIpExistsInSecurityIpListRequest, runtime *dara.RuntimeOptions) (_result *CheckIpExistsInSecurityIpListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckIpExistsInSecurityIpList"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckIpExistsInSecurityIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 判断指定 IP 是否已经存在于网络白名单组
//
// @param request - CheckIpExistsInSecurityIpListRequest
//
// @return CheckIpExistsInSecurityIpListResponse
func (client *Client) CheckIpExistsInSecurityIpList(request *CheckIpExistsInSecurityIpListRequest) (_result *CheckIpExistsInSecurityIpListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckIpExistsInSecurityIpListResponse{}
	_body, _err := client.CheckIpExistsInSecurityIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查服务关联角色
//
// @param request - CheckServiceLinkedRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleResponse
func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *dara.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckServiceLinkedRole"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查服务关联角色
//
// @param request - CheckServiceLinkedRoleRequest
//
// @return CheckServiceLinkedRoleResponse
func (client *Client) CheckServiceLinkedRole(request *CheckServiceLinkedRoleRequest) (_result *CheckServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CheckServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cluster in an ApsaraDB for SelectDB instance.
//
// Description:
//
// > : For an instance that uses the pay-as-you-go billing method, you can create only pay-as-you-go clusters. For an instance that uses the subscription billing method, you can create pay-as-you-go or subscription clusters.
//
// @param request - CreateDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBClusterResponse
func (client *Client) CreateDBClusterWithOptions(request *CreateDBClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheSize) {
		query["CacheSize"] = request.CacheSize
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClusterNodeCount) {
		query["ClusterNodeCount"] = request.ClusterNodeCount
	}

	if !dara.IsNil(request.ClusterNodeType) {
		query["ClusterNodeType"] = request.ClusterNodeType
	}

	if !dara.IsNil(request.DBClusterClass) {
		query["DBClusterClass"] = request.DBClusterClass
	}

	if !dara.IsNil(request.DBClusterDescription) {
		query["DBClusterDescription"] = request.DBClusterDescription
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBCluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cluster in an ApsaraDB for SelectDB instance.
//
// Description:
//
// > : For an instance that uses the pay-as-you-go billing method, you can create only pay-as-you-go clusters. For an instance that uses the subscription billing method, you can create pay-as-you-go or subscription clusters.
//
// @param request - CreateDBClusterRequest
//
// @return CreateDBClusterResponse
func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (_result *CreateDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBClusterResponse{}
	_body, _err := client.CreateDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a binding relationship for clusters. If the zone-redundant storage (ZRS) deployment method is used, you can create a binding relationship between two clusters.
//
// Description:
//
// This operation is supported only for instances that use the zone-redundant storage (ZRS) feature and meet the following requirements:
//
//   - The instance clusters reside in different zones.
//
//   - The billing method of the instance clusters is consistent.
//
// @param request - CreateDBClusterBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBClusterBindingResponse
func (client *Client) CreateDBClusterBindingWithOptions(request *CreateDBClusterBindingRequest, runtime *dara.RuntimeOptions) (_result *CreateDBClusterBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBClusterIdBak) {
		query["DBClusterIdBak"] = request.DBClusterIdBak
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBClusterBinding"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBClusterBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a binding relationship for clusters. If the zone-redundant storage (ZRS) deployment method is used, you can create a binding relationship between two clusters.
//
// Description:
//
// This operation is supported only for instances that use the zone-redundant storage (ZRS) feature and meet the following requirements:
//
//   - The instance clusters reside in different zones.
//
//   - The billing method of the instance clusters is consistent.
//
// @param request - CreateDBClusterBindingRequest
//
// @return CreateDBClusterBindingResponse
func (client *Client) CreateDBClusterBinding(request *CreateDBClusterBindingRequest) (_result *CreateDBClusterBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBClusterBindingResponse{}
	_body, _err := client.CreateDBClusterBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an ApsaraDB for SelectDB instance.
//
// @param tmpReq - CreateDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstanceWithOptions(tmpReq *CreateDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDBInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MultiZone) {
		request.MultiZoneShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiZone, dara.String("MultiZone"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AddVPCIPs) {
		query["AddVPCIPs"] = request.AddVPCIPs
	}

	if !dara.IsNil(request.CacheSize) {
		query["CacheSize"] = request.CacheSize
	}

	if !dara.IsNil(request.ChargeType) {
		query["ChargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClusterNodeCount) {
		query["ClusterNodeCount"] = request.ClusterNodeCount
	}

	if !dara.IsNil(request.ClusterNodeType) {
		query["ClusterNodeType"] = request.ClusterNodeType
	}

	if !dara.IsNil(request.ConfigPatternType) {
		query["ConfigPatternType"] = request.ConfigPatternType
	}

	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceClass) {
		query["DBInstanceClass"] = request.DBInstanceClass
	}

	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DeployScheme) {
		query["DeployScheme"] = request.DeployScheme
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.MultiZoneShrink) {
		query["MultiZone"] = request.MultiZoneShrink
	}

	if !dara.IsNil(request.Period) {
		query["Period"] = request.Period
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	if !dara.IsNil(request.UsedTime) {
		query["UsedTime"] = request.UsedTime
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDBInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ApsaraDB for SelectDB instance.
//
// @param request - CreateDBInstanceRequest
//
// @return CreateDBInstanceResponse
func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (_result *CreateDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scheduled scaling rule.
//
// @param request - CreateElasticRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateElasticRuleResponse
func (client *Client) CreateElasticRuleWithOptions(request *CreateElasticRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateElasticRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterClass) {
		query["ClusterClass"] = request.ClusterClass
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.ElasticRuleStartTime) {
		query["ElasticRuleStartTime"] = request.ElasticRuleStartTime
	}

	if !dara.IsNil(request.ExecutionPeriod) {
		query["ExecutionPeriod"] = request.ExecutionPeriod
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
		Action:      dara.String("CreateElasticRule"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateElasticRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled scaling rule.
//
// @param request - CreateElasticRuleRequest
//
// @return CreateElasticRuleResponse
func (client *Client) CreateElasticRule(request *CreateElasticRuleRequest) (_result *CreateElasticRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateElasticRuleResponse{}
	_body, _err := client.CreateElasticRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service-linked role for ApsaraDB for SelectDB.
//
// @param request - CreateServiceLinkedRoleForSelectDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleForSelectDBResponse
func (client *Client) CreateServiceLinkedRoleForSelectDBWithOptions(request *CreateServiceLinkedRoleForSelectDBRequest, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleForSelectDBResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRoleForSelectDB"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleForSelectDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service-linked role for ApsaraDB for SelectDB.
//
// @param request - CreateServiceLinkedRoleForSelectDBRequest
//
// @return CreateServiceLinkedRoleForSelectDBResponse
func (client *Client) CreateServiceLinkedRoleForSelectDB(request *CreateServiceLinkedRoleForSelectDBRequest) (_result *CreateServiceLinkedRoleForSelectDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleForSelectDBResponse{}
	_body, _err := client.CreateServiceLinkedRoleForSelectDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建虚拟集群
//
// @param request - CreateVirtualClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirtualClusterResponse
func (client *Client) CreateVirtualClusterWithOptions(request *CreateVirtualClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateVirtualClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveClusterId) {
		query["ActiveClusterId"] = request.ActiveClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StandbyClusterId) {
		query["StandbyClusterId"] = request.StandbyClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVirtualCluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVirtualClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建虚拟集群
//
// @param request - CreateVirtualClusterRequest
//
// @return CreateVirtualClusterResponse
func (client *Client) CreateVirtualCluster(request *CreateVirtualClusterRequest) (_result *CreateVirtualClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVirtualClusterResponse{}
	_body, _err := client.CreateVirtualClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a cluster from an instance.
//
// @param request - DeleteDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterResponse
func (client *Client) DeleteDBClusterWithOptions(request *DeleteDBClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBCluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a cluster from an instance.
//
// @param request - DeleteDBClusterRequest
//
// @return DeleteDBClusterResponse
func (client *Client) DeleteDBCluster(request *DeleteDBClusterRequest) (_result *DeleteDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBClusterResponse{}
	_body, _err := client.DeleteDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the binding relationship between two clusters for mutual backup.
//
// @param request - DeleteDBClusterBindingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBClusterBindingResponse
func (client *Client) DeleteDBClusterBindingWithOptions(request *DeleteDBClusterBindingRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBClusterBindingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBClusterIdBak) {
		query["DBClusterIdBak"] = request.DBClusterIdBak
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBClusterBinding"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBClusterBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the binding relationship between two clusters for mutual backup.
//
// @param request - DeleteDBClusterBindingRequest
//
// @return DeleteDBClusterBindingResponse
func (client *Client) DeleteDBClusterBinding(request *DeleteDBClusterBindingRequest) (_result *DeleteDBClusterBindingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBClusterBindingResponse{}
	_body, _err := client.DeleteDBClusterBindingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除DB实例
//
// @param request - DeleteDBInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDBInstanceResponse
func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDBInstance"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除DB实例
//
// @param request - DeleteDBInstanceRequest
//
// @return DeleteDBInstanceResponse
func (client *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (_result *DeleteDBInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DeleteDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scheduled scaling rule.
//
// @param request - DeleteElasticRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteElasticRuleResponse
func (client *Client) DeleteElasticRuleWithOptions(request *DeleteElasticRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteElasticRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteElasticRule"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteElasticRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scheduled scaling rule.
//
// @param request - DeleteElasticRuleRequest
//
// @return DeleteElasticRuleResponse
func (client *Client) DeleteElasticRule(request *DeleteElasticRuleRequest) (_result *DeleteElasticRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteElasticRuleResponse{}
	_body, _err := client.DeleteElasticRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除虚拟集群
//
// @param request - DeleteVirtualClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirtualClusterResponse
func (client *Client) DeleteVirtualClusterWithOptions(request *DeleteVirtualClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteVirtualClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVirtualCluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVirtualClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除虚拟集群
//
// @param request - DeleteVirtualClusterRequest
//
// @return DeleteVirtualClusterResponse
func (client *Client) DeleteVirtualCluster(request *DeleteVirtualClusterRequest) (_result *DeleteVirtualClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVirtualClusterResponse{}
	_body, _err := client.DeleteVirtualClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all instance specifications.
//
// @param request - DescribeAllDBInstanceClassRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAllDBInstanceClassResponse
func (client *Client) DescribeAllDBInstanceClassWithOptions(request *DescribeAllDBInstanceClassRequest, runtime *dara.RuntimeOptions) (_result *DescribeAllDBInstanceClassResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAllDBInstanceClass"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAllDBInstanceClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all instance specifications.
//
// @param request - DescribeAllDBInstanceClassRequest
//
// @return DescribeAllDBInstanceClassResponse
func (client *Client) DescribeAllDBInstanceClass(request *DescribeAllDBInstanceClassRequest) (_result *DescribeAllDBInstanceClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAllDBInstanceClassResponse{}
	_body, _err := client.DescribeAllDBInstanceClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration information about a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - DescribeDBClusterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterConfigResponse
func (client *Client) DescribeDBClusterConfigWithOptions(request *DescribeDBClusterConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterConfig"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration information about a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - DescribeDBClusterConfigRequest
//
// @return DescribeDBClusterConfigResponse
func (client *Client) DescribeDBClusterConfig(request *DescribeDBClusterConfigRequest) (_result *DescribeDBClusterConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterConfigResponse{}
	_body, _err := client.DescribeDBClusterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration change logs of a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - DescribeDBClusterConfigChangeLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterConfigChangeLogsResponse
func (client *Client) DescribeDBClusterConfigChangeLogsWithOptions(request *DescribeDBClusterConfigChangeLogsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterConfigChangeLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterConfigChangeLogs"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterConfigChangeLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration change logs of a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - DescribeDBClusterConfigChangeLogsRequest
//
// @return DescribeDBClusterConfigChangeLogsResponse
func (client *Client) DescribeDBClusterConfigChangeLogs(request *DescribeDBClusterConfigChangeLogsRequest) (_result *DescribeDBClusterConfigChangeLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterConfigChangeLogsResponse{}
	_body, _err := client.DescribeDBClusterConfigChangeLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取集群的各规格缓存限制
//
// @param request - DescribeDBClusterStorageLimitationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBClusterStorageLimitationResponse
func (client *Client) DescribeDBClusterStorageLimitationWithOptions(request *DescribeDBClusterStorageLimitationRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBClusterStorageLimitationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBClusterStorageLimitation"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBClusterStorageLimitationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取集群的各规格缓存限制
//
// @param request - DescribeDBClusterStorageLimitationRequest
//
// @return DescribeDBClusterStorageLimitationResponse
func (client *Client) DescribeDBClusterStorageLimitation(request *DescribeDBClusterStorageLimitationRequest) (_result *DescribeDBClusterStorageLimitationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBClusterStorageLimitationResponse{}
	_body, _err := client.DescribeDBClusterStorageLimitationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about an ApsaraDB for SelectDB instance.
//
// @param request - DescribeDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceAttributeResponse
func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("DescribeDBInstanceAttribute"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about an ApsaraDB for SelectDB instance.
//
// @param request - DescribeDBInstanceAttributeRequest
//
// @return DescribeDBInstanceAttributeResponse
func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DescribeDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the network information about an ApsaraDB for SelectDB instance.
//
// @param request - DescribeDBInstanceNetInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstanceNetInfoResponse
func (client *Client) DescribeDBInstanceNetInfoWithOptions(request *DescribeDBInstanceNetInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("DescribeDBInstanceNetInfo"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the network information about an ApsaraDB for SelectDB instance.
//
// @param request - DescribeDBInstanceNetInfoRequest
//
// @return DescribeDBInstanceNetInfoResponse
func (client *Client) DescribeDBInstanceNetInfo(request *DescribeDBInstanceNetInfoRequest) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DescribeDBInstanceNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about ApsaraDB for SelectDB instances.
//
// @param tmpReq - DescribeDBInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDBInstancesResponse
func (client *Client) DescribeDBInstancesWithOptions(tmpReq *DescribeDBInstancesRequest, runtime *dara.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeDBInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceDescription) {
		query["DBInstanceDescription"] = request.DBInstanceDescription
	}

	if !dara.IsNil(request.DBInstanceIds) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !dara.IsNil(request.DBInstanceStatus) {
		query["DBInstanceStatus"] = request.DBInstanceStatus
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

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDBInstances"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about ApsaraDB for SelectDB instances.
//
// @param request - DescribeDBInstancesRequest
//
// @return DescribeDBInstancesResponse
func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (_result *DescribeDBInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DescribeDBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries scheduled scaling rules.
//
// @param request - DescribeElasticRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeElasticRulesResponse
func (client *Client) DescribeElasticRulesWithOptions(request *DescribeElasticRulesRequest, runtime *dara.RuntimeOptions) (_result *DescribeElasticRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeElasticRules"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeElasticRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries scheduled scaling rules.
//
// @param request - DescribeElasticRulesRequest
//
// @return DescribeElasticRulesResponse
func (client *Client) DescribeElasticRules(request *DescribeElasticRulesRequest) (_result *DescribeElasticRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeElasticRulesResponse{}
	_body, _err := client.DescribeElasticRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries available regions and zones.
//
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
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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

// Summary:
//
// Queries available regions and zones.
//
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
// Queries the IP addresses in the whitelists of an ApsaraDB for SelectDB instance.
//
// @param request - DescribeSecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityIPListResponse
func (client *Client) DescribeSecurityIPListWithOptions(request *DescribeSecurityIPListRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityIPListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("DescribeSecurityIPList"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityIPListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses in the whitelists of an ApsaraDB for SelectDB instance.
//
// @param request - DescribeSecurityIPListRequest
//
// @return DescribeSecurityIPListResponse
func (client *Client) DescribeSecurityIPList(request *DescribeSecurityIPListRequest) (_result *DescribeSecurityIPListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityIPListResponse{}
	_body, _err := client.DescribeSecurityIPListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeVSwitches
//
// @param request - DescribeVSwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVSwitchesResponse
func (client *Client) DescribeVSwitchesWithOptions(request *DescribeVSwitchesRequest, runtime *dara.RuntimeOptions) (_result *DescribeVSwitchesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
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
		Action:      dara.String("DescribeVSwitches"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeVSwitches
//
// @param request - DescribeVSwitchesRequest
//
// @return DescribeVSwitchesResponse
func (client *Client) DescribeVSwitches(request *DescribeVSwitchesRequest) (_result *DescribeVSwitchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeVSwitchesResponse{}
	_body, _err := client.DescribeVSwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DescribeZones
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DescribeZones
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uses the scheduled scaling policy.
//
// @param request - EnDisableScalingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnDisableScalingRulesResponse
func (client *Client) EnDisableScalingRulesWithOptions(request *EnDisableScalingRulesRequest, runtime *dara.RuntimeOptions) (_result *EnDisableScalingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScalingRulesEnable) {
		query["ScalingRulesEnable"] = request.ScalingRulesEnable
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnDisableScalingRules"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnDisableScalingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uses the scheduled scaling policy.
//
// @param request - EnDisableScalingRulesRequest
//
// @return EnDisableScalingRulesResponse
func (client *Client) EnDisableScalingRules(request *EnDisableScalingRulesRequest) (_result *EnDisableScalingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnDisableScalingRulesResponse{}
	_body, _err := client.EnDisableScalingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the pricing for creating a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - GetCreateBEClusterInquiryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCreateBEClusterInquiryResponse
func (client *Client) GetCreateBEClusterInquiryWithOptions(request *GetCreateBEClusterInquiryRequest, runtime *dara.RuntimeOptions) (_result *GetCreateBEClusterInquiryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCreateBEClusterInquiry"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCreateBEClusterInquiryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the pricing for creating a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - GetCreateBEClusterInquiryRequest
//
// @return GetCreateBEClusterInquiryResponse
func (client *Client) GetCreateBEClusterInquiry(request *GetCreateBEClusterInquiryRequest) (_result *GetCreateBEClusterInquiryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCreateBEClusterInquiryResponse{}
	_body, _err := client.GetCreateBEClusterInquiryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the pricing for changing the specifications of a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - GetModifyBEClusterInquiryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModifyBEClusterInquiryResponse
func (client *Client) GetModifyBEClusterInquiryWithOptions(request *GetModifyBEClusterInquiryRequest, runtime *dara.RuntimeOptions) (_result *GetModifyBEClusterInquiryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModifyBEClusterInquiry"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModifyBEClusterInquiryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the pricing for changing the specifications of a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - GetModifyBEClusterInquiryRequest
//
// @return GetModifyBEClusterInquiryResponse
func (client *Client) GetModifyBEClusterInquiry(request *GetModifyBEClusterInquiryRequest) (_result *GetModifyBEClusterInquiryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetModifyBEClusterInquiryResponse{}
	_body, _err := client.GetModifyBEClusterInquiryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - ModifyBEClusterAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyBEClusterAttributeResponse
func (client *Client) ModifyBEClusterAttributeWithOptions(request *ModifyBEClusterAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyBEClusterAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.InstanceAttributeType) {
		query["InstanceAttributeType"] = request.InstanceAttributeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyBEClusterAttribute"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyBEClusterAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - ModifyBEClusterAttributeRequest
//
// @return ModifyBEClusterAttributeResponse
func (client *Client) ModifyBEClusterAttribute(request *ModifyBEClusterAttributeRequest) (_result *ModifyBEClusterAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyBEClusterAttributeResponse{}
	_body, _err := client.ModifyBEClusterAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 集群变配
//
// @param request - ModifyDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterResponse
func (client *Client) ModifyDBClusterWithOptions(request *ModifyDBClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CacheSize) {
		query["CacheSize"] = request.CacheSize
	}

	if !dara.IsNil(request.ClusterNodeCount) {
		query["ClusterNodeCount"] = request.ClusterNodeCount
	}

	if !dara.IsNil(request.ClusterNodeType) {
		query["ClusterNodeType"] = request.ClusterNodeType
	}

	if !dara.IsNil(request.DBClusterClass) {
		query["DBClusterClass"] = request.DBClusterClass
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScaleMax) {
		query["ScaleMax"] = request.ScaleMax
	}

	if !dara.IsNil(request.ScaleMin) {
		query["ScaleMin"] = request.ScaleMin
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBCluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 集群变配
//
// @param request - ModifyDBClusterRequest
//
// @return ModifyDBClusterResponse
func (client *Client) ModifyDBCluster(request *ModifyDBClusterRequest) (_result *ModifyDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterResponse{}
	_body, _err := client.ModifyDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - ModifyDBClusterConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBClusterConfigResponse
func (client *Client) ModifyDBClusterConfigWithOptions(request *ModifyDBClusterConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBClusterConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigKey) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.ParallelOperation) {
		query["ParallelOperation"] = request.ParallelOperation
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBClusterConfig"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBClusterConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - ModifyDBClusterConfigRequest
//
// @return ModifyDBClusterConfigResponse
func (client *Client) ModifyDBClusterConfig(request *ModifyDBClusterConfigRequest) (_result *ModifyDBClusterConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBClusterConfigResponse{}
	_body, _err := client.ModifyDBClusterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window or description of an ApsaraDB for SelectDB instance.
//
// @param request - ModifyDBInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDBInstanceAttributeResponse
func (client *Client) ModifyDBInstanceAttributeWithOptions(request *ModifyDBInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDBInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.InstanceAttributeType) {
		query["InstanceAttributeType"] = request.InstanceAttributeType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDBInstanceAttribute"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDBInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window or description of an ApsaraDB for SelectDB instance.
//
// @param request - ModifyDBInstanceAttributeRequest
//
// @return ModifyDBInstanceAttributeResponse
func (client *Client) ModifyDBInstanceAttribute(request *ModifyDBInstanceAttributeRequest) (_result *ModifyDBInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDBInstanceAttributeResponse{}
	_body, _err := client.ModifyDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a scheduled scaling rule.
//
// @param request - ModifyElasticRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyElasticRuleResponse
func (client *Client) ModifyElasticRuleWithOptions(request *ModifyElasticRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyElasticRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterClass) {
		query["ClusterClass"] = request.ClusterClass
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DbInstanceId) {
		query["DbInstanceId"] = request.DbInstanceId
	}

	if !dara.IsNil(request.ElasticRuleStartTime) {
		query["ElasticRuleStartTime"] = request.ElasticRuleStartTime
	}

	if !dara.IsNil(request.ExecutionPeriod) {
		query["ExecutionPeriod"] = request.ExecutionPeriod
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyElasticRule"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyElasticRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a scheduled scaling rule.
//
// @param request - ModifyElasticRuleRequest
//
// @return ModifyElasticRuleResponse
func (client *Client) ModifyElasticRule(request *ModifyElasticRuleRequest) (_result *ModifyElasticRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyElasticRuleResponse{}
	_body, _err := client.ModifyElasticRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the IP addresses in a whitelist of an ApsaraDB for SelectDB instance.
//
// @param request - ModifySecurityIPListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecurityIPListResponse
func (client *Client) ModifySecurityIPListWithOptions(request *ModifySecurityIPListRequest, runtime *dara.RuntimeOptions) (_result *ModifySecurityIPListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ModifyMode) {
		query["ModifyMode"] = request.ModifyMode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityIPList) {
		query["SecurityIPList"] = request.SecurityIPList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecurityIPList"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecurityIPListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the IP addresses in a whitelist of an ApsaraDB for SelectDB instance.
//
// @param request - ModifySecurityIPListRequest
//
// @return ModifySecurityIPListResponse
func (client *Client) ModifySecurityIPList(request *ModifySecurityIPListRequest) (_result *ModifySecurityIPListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySecurityIPListResponse{}
	_body, _err := client.ModifySecurityIPListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改虚拟集群
//
// @param request - ModifyVirtualClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyVirtualClusterResponse
func (client *Client) ModifyVirtualClusterWithOptions(request *ModifyVirtualClusterRequest, runtime *dara.RuntimeOptions) (_result *ModifyVirtualClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActiveClusterId) {
		query["ActiveClusterId"] = request.ActiveClusterId
	}

	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StandbyClusterId) {
		query["StandbyClusterId"] = request.StandbyClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyVirtualCluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyVirtualClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改虚拟集群
//
// @param request - ModifyVirtualClusterRequest
//
// @return ModifyVirtualClusterResponse
func (client *Client) ModifyVirtualCluster(request *ModifyVirtualClusterRequest) (_result *ModifyVirtualClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyVirtualClusterResponse{}
	_body, _err := client.ModifyVirtualClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an ApsaraDB for SelectDB instance.
//
// @param request - ReleaseInstancePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstancePublicConnectionResponse
func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionString) {
		query["ConnectionString"] = request.ConnectionString
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("ReleaseInstancePublicConnection"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of an ApsaraDB for SelectDB instance.
//
// @param request - ReleaseInstancePublicConnectionRequest
//
// @return ReleaseInstancePublicConnectionResponse
func (client *Client) ReleaseInstancePublicConnection(request *ReleaseInstancePublicConnectionRequest) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.ReleaseInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of an account for an ApsaraDB for SelectDB instance.
//
// @param request - ResetAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *dara.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetAccountPassword"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of an account for an ApsaraDB for SelectDB instance.
//
// @param request - ResetAccountPasswordRequest
//
// @return ResetAccountPasswordResponse
func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - RestartDBClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDBClusterResponse
func (client *Client) RestartDBClusterWithOptions(request *RestartDBClusterRequest, runtime *dara.RuntimeOptions) (_result *RestartDBClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.ParallelOperation) {
		query["ParallelOperation"] = request.ParallelOperation
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		body["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDBCluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDBClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a cluster in an ApsaraDB for SelectDB instance.
//
// @param request - RestartDBClusterRequest
//
// @return RestartDBClusterResponse
func (client *Client) RestartDBCluster(request *RestartDBClusterRequest) (_result *RestartDBClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartDBClusterResponse{}
	_body, _err := client.RestartDBClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂停后恢复集群
//
// @param request - StartBEClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBEClusterResponse
func (client *Client) StartBEClusterWithOptions(request *StartBEClusterRequest, runtime *dara.RuntimeOptions) (_result *StartBEClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("StartBECluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartBEClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停后恢复集群
//
// @param request - StartBEClusterRequest
//
// @return StartBEClusterResponse
func (client *Client) StartBECluster(request *StartBEClusterRequest) (_result *StartBEClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartBEClusterResponse{}
	_body, _err := client.StartBEClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 暂停BE集群
//
// @param request - StopBEClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopBEClusterResponse
func (client *Client) StopBEClusterWithOptions(request *StopBEClusterRequest, runtime *dara.RuntimeOptions) (_result *StopBEClusterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBClusterId) {
		query["DBClusterId"] = request.DBClusterId
	}

	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
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
		Action:      dara.String("StopBECluster"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopBEClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 暂停BE集群
//
// @param request - StopBEClusterRequest
//
// @return StopBEClusterResponse
func (client *Client) StopBECluster(request *StopBEClusterRequest) (_result *StopBEClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopBEClusterResponse{}
	_body, _err := client.StopBEClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 资源打用户标签
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
		Version:     dara.String("2023-05-22"),
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
// 资源打用户标签
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
// 资源去除用户标签
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
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2023-05-22"),
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
// 资源去除用户标签
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
// # UpgradeDBInstanceDeployScheme
//
// @param tmpReq - UpgradeDBInstanceDeploySchemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceDeploySchemeResponse
func (client *Client) UpgradeDBInstanceDeploySchemeWithOptions(tmpReq *UpgradeDBInstanceDeploySchemeRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBInstanceDeploySchemeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpgradeDBInstanceDeploySchemeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MultiZone) {
		request.MultiZoneShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultiZone, dara.String("MultiZone"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.MultiZoneShrink) {
		query["MultiZone"] = request.MultiZoneShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceDeployScheme"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBInstanceDeploySchemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpgradeDBInstanceDeployScheme
//
// @param request - UpgradeDBInstanceDeploySchemeRequest
//
// @return UpgradeDBInstanceDeploySchemeResponse
func (client *Client) UpgradeDBInstanceDeployScheme(request *UpgradeDBInstanceDeploySchemeRequest) (_result *UpgradeDBInstanceDeploySchemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeDBInstanceDeploySchemeResponse{}
	_body, _err := client.UpgradeDBInstanceDeploySchemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the database engine version of an ApsaraDB for SelectDB instance.
//
// @param request - UpgradeDBInstanceEngineVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeDBInstanceEngineVersionResponse
func (client *Client) UpgradeDBInstanceEngineVersionWithOptions(request *UpgradeDBInstanceEngineVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeDBInstanceEngineVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBInstanceId) {
		query["DBInstanceId"] = request.DBInstanceId
	}

	if !dara.IsNil(request.EngineVersion) {
		query["EngineVersion"] = request.EngineVersion
	}

	if !dara.IsNil(request.ParallelOperation) {
		query["ParallelOperation"] = request.ParallelOperation
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SwitchTimeMode) {
		query["SwitchTimeMode"] = request.SwitchTimeMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeDBInstanceEngineVersion"),
		Version:     dara.String("2023-05-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeDBInstanceEngineVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the database engine version of an ApsaraDB for SelectDB instance.
//
// @param request - UpgradeDBInstanceEngineVersionRequest
//
// @return UpgradeDBInstanceEngineVersionResponse
func (client *Client) UpgradeDBInstanceEngineVersion(request *UpgradeDBInstanceEngineVersionRequest) (_result *UpgradeDBInstanceEngineVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeDBInstanceEngineVersionResponse{}
	_body, _err := client.UpgradeDBInstanceEngineVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
