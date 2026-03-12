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
	client.Endpoint, _err = client.GetEndpoint(dara.String("starrocks"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 新建网关
//
// @param request - AddGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewayResponse
func (client *Client) AddGatewayWithOptions(request *AddGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeNodeNumber) {
		query["FeNodeNumber"] = request.FeNodeNumber
	}

	if !dara.IsNil(request.GatewayName) {
		query["GatewayName"] = request.GatewayName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGateway"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/gateway/add"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新建网关
//
// @param request - AddGatewayRequest
//
// @return AddGatewayResponse
func (client *Client) AddGateway(request *AddGatewayRequest) (_result *AddGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddGatewayResponse{}
	_body, _err := client.AddGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to modify the resource group of a Serverless StarRocks instance.
//
// @param request - ChangeResourceGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
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

	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceGroup/change"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// This interface is used to modify the resource group of a Serverless StarRocks instance.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建StarRocks集群
//
// @param request - CreateInstanceV1Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceV1Response
func (client *Client) CreateInstanceV1WithOptions(request *CreateInstanceV1Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceV1Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdminPassword) {
		body["AdminPassword"] = request.AdminPassword
	}

	if !dara.IsNil(request.AgentNodeGroup) {
		body["AgentNodeGroup"] = request.AgentNodeGroup
	}

	if !dara.IsNil(request.AutoPay) {
		body["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BackendNodeGroups) {
		body["BackendNodeGroups"] = request.BackendNodeGroups
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DlfCatalogName) {
		body["DlfCatalogName"] = request.DlfCatalogName
	}

	if !dara.IsNil(request.DlfCatalogType) {
		body["DlfCatalogType"] = request.DlfCatalogType
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.Encrypted) {
		body["Encrypted"] = request.Encrypted
	}

	if !dara.IsNil(request.FrontendNodeGroups) {
		body["FrontendNodeGroups"] = request.FrontendNodeGroups
	}

	if !dara.IsNil(request.GatewayType) {
		body["GatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.KmsKeyId) {
		body["KmsKeyId"] = request.KmsKeyId
	}

	if !dara.IsNil(request.LinkedRamUserName) {
		body["LinkedRamUserName"] = request.LinkedRamUserName
	}

	if !dara.IsNil(request.ObserverNodeGroups) {
		body["ObserverNodeGroups"] = request.ObserverNodeGroups
	}

	if !dara.IsNil(request.OssAccessingRoleName) {
		body["OssAccessingRoleName"] = request.OssAccessingRoleName
	}

	if !dara.IsNil(request.PackageType) {
		body["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.PayType) {
		body["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PrincipalType) {
		body["PrincipalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		body["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.RamUserId) {
		body["RamUserId"] = request.RamUserId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.RunMode) {
		body["RunMode"] = request.RunMode
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	if !dara.IsNil(request.VSwitches) {
		body["VSwitches"] = request.VSwitches
	}

	if !dara.IsNil(request.Version) {
		body["Version"] = request.Version
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceV1"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/cluster/createV1"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceV1Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建StarRocks集群
//
// @param request - CreateInstanceV1Request
//
// @return CreateInstanceV1Response
func (client *Client) CreateInstanceV1(request *CreateInstanceV1Request) (_result *CreateInstanceV1Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceV1Response{}
	_body, _err := client.CreateInstanceV1WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This interface is used to create the AliyunServiceRoleForEMRStarRocks role for users.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRoleWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceLinkedRole"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/user/create_default_role"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to create the AliyunServiceRoleForEMRStarRocks role for users.
//
// @return CreateServiceLinkedRoleResponse
func (client *Client) CreateServiceLinkedRole() (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除网关
//
// @param request - DeleteGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithOptions(request *DeleteGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGateway"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/gateway/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除网关
//
// @param request - DeleteGatewayRequest
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGateway(request *DeleteGatewayRequest) (_result *DeleteGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例配置历史
//
// @param request - DescribeConfigHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeConfigHistoryResponse
func (client *Client) DescribeConfigHistoryWithOptions(request *DescribeConfigHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeConfigHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EffectStatuses) {
		query["EffectStatuses"] = request.EffectStatuses
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NeedTotal) {
		query["NeedTotal"] = request.NeedTotal
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeConfigHistory"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/config/describeConfigHistory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeConfigHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例配置历史
//
// @param request - DescribeConfigHistoryRequest
//
// @return DescribeConfigHistoryResponse
func (client *Client) DescribeConfigHistory(request *DescribeConfigHistoryRequest) (_result *DescribeConfigHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeConfigHistoryResponse{}
	_body, _err := client.DescribeConfigHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例配置
//
// @param request - DescribeInstanceConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceConfigsResponse
func (client *Client) DescribeInstanceConfigsWithOptions(request *DescribeInstanceConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInstanceConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowModify) {
		query["AllowModify"] = request.AllowModify
	}

	if !dara.IsNil(request.ConfigKey) {
		query["ConfigKey"] = request.ConfigKey
	}

	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NeedTotal) {
		query["NeedTotal"] = request.NeedTotal
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceConfigs"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/config/describeInstanceConfigs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例配置
//
// @param request - DescribeInstanceConfigsRequest
//
// @return DescribeInstanceConfigsResponse
func (client *Client) DescribeInstanceConfigs(request *DescribeInstanceConfigsRequest) (_result *DescribeInstanceConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceConfigsResponse{}
	_body, _err := client.DescribeInstanceConfigsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation is used to query Serverless StarRocks instances, supporting filtering based on instance name or tags and other information.
//
// @param tmpReq - DescribeInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstancesWithOptions(tmpReq *DescribeInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.InstanceStatus) {
		query["InstanceStatus"] = request.InstanceStatus
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

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstances"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/describeInstances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This operation is used to query Serverless StarRocks instances, supporting filtering based on instance name or tags and other information.
//
// @param request - DescribeInstancesRequest
//
// @return DescribeInstancesResponse
func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节点组信息
//
// @param request - DescribeNodeGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeGroupsResponse
func (client *Client) DescribeNodeGroupsWithOptions(request *DescribeNodeGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeNodeGroupsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ComponentType) {
		body["componentType"] = request.ComponentType
	}

	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeGroupIds) {
		body["nodeGroupIds"] = request.NodeGroupIds
	}

	if !dara.IsNil(request.NodeGroupName) {
		body["nodeGroupName"] = request.NodeGroupName
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNodeGroups"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/nodegroup/describeNodeGroups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNodeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点组信息
//
// @param request - DescribeNodeGroupsRequest
//
// @return DescribeNodeGroupsResponse
func (client *Client) DescribeNodeGroups(request *DescribeNodeGroupsRequest) (_result *DescribeNodeGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeNodeGroupsResponse{}
	_body, _err := client.DescribeNodeGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # StarRocks关闭SSL
//
// @param request - DisableSSLConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSSLConnectionResponse
func (client *Client) DisableSSLConnectionWithOptions(request *DisableSSLConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DisableSSLConnectionResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableSSLConnection"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/disableSSLConnection"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableSSLConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StarRocks关闭SSL
//
// @param request - DisableSSLConnectionRequest
//
// @return DisableSSLConnectionResponse
func (client *Client) DisableSSLConnection(request *DisableSSLConnectionRequest) (_result *DisableSSLConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableSSLConnectionResponse{}
	_body, _err := client.DisableSSLConnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # StarRocks开启SSL
//
// @param request - EnableSSLConnectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSSLConnectionResponse
func (client *Client) EnableSSLConnectionWithOptions(request *EnableSSLConnectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableSSLConnectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomSSLCertificate) {
		body["CustomSSLCertificate"] = request.CustomSSLCertificate
	}

	if !dara.IsNil(request.EnableCustom) {
		body["EnableCustom"] = request.EnableCustom
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Renewal) {
		body["Renewal"] = request.Renewal
	}

	if !dara.IsNil(request.SslKeyPassword) {
		body["SslKeyPassword"] = request.SslKeyPassword
	}

	if !dara.IsNil(request.SslKeystorePassword) {
		body["SslKeystorePassword"] = request.SslKeystorePassword
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableSSLConnection"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/enableSSLConnection"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableSSLConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # StarRocks开启SSL
//
// @param request - EnableSSLConnectionRequest
//
// @return EnableSSLConnectionResponse
func (client *Client) EnableSSLConnection(request *EnableSSLConnectionRequest) (_result *EnableSSLConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableSSLConnectionResponse{}
	_body, _err := client.EnableSSLConnectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取StarRocks集群实例的特性开关
//
// @param request - GetInstanceFeatureGateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceFeatureGateResponse
func (client *Client) GetInstanceFeatureGateWithOptions(request *GetInstanceFeatureGateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceFeatureGateResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceFeatureGate"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/features/featureGate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceFeatureGateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取StarRocks集群实例的特性开关
//
// @param request - GetInstanceFeatureGateRequest
//
// @return GetInstanceFeatureGateResponse
func (client *Client) GetInstanceFeatureGate(request *GetInstanceFeatureGateRequest) (_result *GetInstanceFeatureGateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceFeatureGateResponse{}
	_body, _err := client.GetInstanceFeatureGateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 默认网关开启内网SLB
//
// @param request - IsolateLeaderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsolateLeaderResponse
func (client *Client) IsolateLeaderWithOptions(request *IsolateLeaderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *IsolateLeaderResponse, _err error) {
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

	if !dara.IsNil(request.IsolateLeader) {
		query["IsolateLeader"] = request.IsolateLeader
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IsolateLeader"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/gateway/isolateLeader"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IsolateLeaderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 默认网关开启内网SLB
//
// @param request - IsolateLeaderRequest
//
// @return IsolateLeaderResponse
func (client *Client) IsolateLeader(request *IsolateLeaderRequest) (_result *IsolateLeaderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IsolateLeaderResponse{}
	_body, _err := client.IsolateLeaderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取网关列表
//
// @param request - ListGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayResponse
func (client *Client) ListGatewayWithOptions(request *ListGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewayResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGateway"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/gateway/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网关列表
//
// @param request - ListGatewayRequest
//
// @return ListGatewayResponse
func (client *Client) ListGateway(request *ListGatewayRequest) (_result *ListGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewayResponse{}
	_body, _err := client.ListGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例的付费类型
//
// @param request - ModifyChargeTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyChargeTypeResponse
func (client *Client) ModifyChargeTypeWithOptions(request *ModifyChargeTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyChargeTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenew) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BillingInstanceIds) {
		query["BillingInstanceIds"] = request.BillingInstanceIds
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyChargeType"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/cluster/modifyChargeType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyChargeTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例的付费类型
//
// @param request - ModifyChargeTypeRequest
//
// @return ModifyChargeTypeResponse
func (client *Client) ModifyChargeType(request *ModifyChargeTypeRequest) (_result *ModifyChargeTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyChargeTypeResponse{}
	_body, _err := client.ModifyChargeTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the number of CUs for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
//   - You can modify the number of CUs for a warehouse of only StarRocks instances of Standard Edition.
//
//   - You can increase the number of disks only for warehouses of the standard specifications.
//
//   - The instance must be in the Running state.
//
// After you modify the number of CUs for a warehouse, the billing of CUs has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged based on the number of CUs.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the number of CUs before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyCuRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCuResponse
func (client *Client) ModifyCuWithOptions(request *ModifyCuRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyCuResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FastMode) {
		query["FastMode"] = request.FastMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCu"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifyCu"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the number of CUs for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
//   - You can modify the number of CUs for a warehouse of only StarRocks instances of Standard Edition.
//
//   - You can increase the number of disks only for warehouses of the standard specifications.
//
//   - The instance must be in the Running state.
//
// After you modify the number of CUs for a warehouse, the billing of CUs has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged based on the number of CUs.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the number of CUs before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyCuRequest
//
// @return ModifyCuResponse
func (client *Client) ModifyCu(request *ModifyCuRequest) (_result *ModifyCuResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyCuResponse{}
	_body, _err := client.ModifyCuWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a precheck before you modify the number of CUs for a warehouse.
//
// @param request - ModifyCuPreCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyCuPreCheckResponse
func (client *Client) ModifyCuPreCheckWithOptions(request *ModifyCuPreCheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyCuPreCheckResponse, _err error) {
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

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyCuPreCheck"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifyCuPreCheck"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyCuPreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a precheck before you modify the number of CUs for a warehouse.
//
// @param request - ModifyCuPreCheckRequest
//
// @return ModifyCuPreCheckResponse
func (client *Client) ModifyCuPreCheck(request *ModifyCuPreCheckRequest) (_result *ModifyCuPreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyCuPreCheckResponse{}
	_body, _err := client.ModifyCuPreCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Increases the number of disks for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
//   - You can increase the number of disks only for StarRocks instances of Standard Edition.
//
//   - You can increase the number of disks only for warehouses of the standard specifications.
//
//   - The instance must be in the Running state.
//
// After you increase the number of disks for a warehouse, the billing of disks has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk type.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the number of disks before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskNumberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskNumberResponse
func (client *Client) ModifyDiskNumberWithOptions(request *ModifyDiskNumberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyDiskNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FastMode) {
		query["FastMode"] = request.FastMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDiskNumber"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifyDiskNumber"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDiskNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Increases the number of disks for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
//   - You can increase the number of disks only for StarRocks instances of Standard Edition.
//
//   - You can increase the number of disks only for warehouses of the standard specifications.
//
//   - The instance must be in the Running state.
//
// After you increase the number of disks for a warehouse, the billing of disks has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk type.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the number of disks before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskNumberRequest
//
// @return ModifyDiskNumberResponse
func (client *Client) ModifyDiskNumber(request *ModifyDiskNumberRequest) (_result *ModifyDiskNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDiskNumberResponse{}
	_body, _err := client.ModifyDiskNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the disk performance level for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
//   - You can modify the disk performance level only for StarRocks instances of Standard Edition.
//
//   - You can modify the disk performance level only for warehouses of the standard specifications.
//
//   - The instance must be in the Running state.
//
//   - You cannot downgrade the performance level to PL0.
//
//   - The performance level of an Enterprise SSD (ESSD) is limited by the ESSD disk size. If you cannot upgrade the performance level of an ESSD, expand the ESSD and try again. For more information, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
//
// After the disk performance level is changed, the billing of the disk has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk type.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the disk performance level before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskPerformanceLevelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskPerformanceLevelResponse
func (client *Client) ModifyDiskPerformanceLevelWithOptions(request *ModifyDiskPerformanceLevelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyDiskPerformanceLevelResponse, _err error) {
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

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDiskPerformanceLevel"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifyDiskPerformanceLevel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDiskPerformanceLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the disk performance level for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://www.alibabacloud.com/en/product/ecs?_p_lc=1&spm=openapi-amp.newDocPublishment.0.0.47c9281fkIZGiB#pricing) of EMR Serverless StarRocks instances.
//
// Before you call this operation, take note of the following items:
//
//   - You can modify the disk performance level only for StarRocks instances of Standard Edition.
//
//   - You can modify the disk performance level only for warehouses of the standard specifications.
//
//   - The instance must be in the Running state.
//
//   - You cannot downgrade the performance level to PL0.
//
//   - The performance level of an Enterprise SSD (ESSD) is limited by the ESSD disk size. If you cannot upgrade the performance level of an ESSD, expand the ESSD and try again. For more information, see [Enterprise SSDs](https://help.aliyun.com/document_detail/122389.html).
//
// After the disk performance level is changed, the billing of the disk has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk type.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the disk performance level before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskPerformanceLevelRequest
//
// @return ModifyDiskPerformanceLevelResponse
func (client *Client) ModifyDiskPerformanceLevel(request *ModifyDiskPerformanceLevelRequest) (_result *ModifyDiskPerformanceLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDiskPerformanceLevelResponse{}
	_body, _err := client.ModifyDiskPerformanceLevelWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Expands the disk size for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
//   - You can expand the disk size only for StarRocks instances of Standard Edition.
//
//   - You can expand the disk size only for warehouses of the standard specifications.
//
//   - The instance must be in the Running state.
//
// After you expand the disk size, the billing of disks has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk size.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the disk size before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskSizeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskSizeResponse
func (client *Client) ModifyDiskSizeWithOptions(request *ModifyDiskSizeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyDiskSizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FastMode) {
		query["FastMode"] = request.FastMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDiskSize"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifyDiskSize"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDiskSizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Expands the disk size for a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
//   - You can expand the disk size only for StarRocks instances of Standard Edition.
//
//   - You can expand the disk size only for warehouses of the standard specifications.
//
//   - The instance must be in the Running state.
//
// After you expand the disk size, the billing of disks has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged for the disk based on the new disk size.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the disk size before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyDiskSizeRequest
//
// @return ModifyDiskSizeResponse
func (client *Client) ModifyDiskSize(request *ModifyDiskSizeRequest) (_result *ModifyDiskSizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDiskSizeResponse{}
	_body, _err := client.ModifyDiskSizeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改计算组的节点磁盘类型
//
// @param request - ModifyDiskTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDiskTypeResponse
func (client *Client) ModifyDiskTypeWithOptions(request *ModifyDiskTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyDiskTypeResponse, _err error) {
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

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.TargetDiskType) {
		query["TargetDiskType"] = request.TargetDiskType
	}

	if !dara.IsNil(request.TargetPerformanceLevel) {
		query["TargetPerformanceLevel"] = request.TargetPerformanceLevel
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDiskType"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifyDiskType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDiskTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改计算组的节点磁盘类型
//
// @param request - ModifyDiskTypeRequest
//
// @return ModifyDiskTypeResponse
func (client *Client) ModifyDiskType(request *ModifyDiskTypeRequest) (_result *ModifyDiskTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyDiskTypeResponse{}
	_body, _err := client.ModifyDiskTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例配置
//
// @param request - ModifyInstanceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceConfigResponse
func (client *Client) ModifyInstanceConfigWithOptions(request *ModifyInstanceConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyInstanceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddConfigList) {
		query["AddConfigList"] = request.AddConfigList
	}

	if !dara.IsNil(request.ConfigList) {
		query["ConfigList"] = request.ConfigList
	}

	if !dara.IsNil(request.DeleteConfigList) {
		query["DeleteConfigList"] = request.DeleteConfigList
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigsToAdd) {
		body["configsToAdd"] = request.ConfigsToAdd
	}

	if !dara.IsNil(request.ConfigsToDelete) {
		body["configsToDelete"] = request.ConfigsToDelete
	}

	if !dara.IsNil(request.ConfigsToUpdate) {
		body["configsToUpdate"] = request.ConfigsToUpdate
	}

	if !dara.IsNil(request.FastMode) {
		body["fastMode"] = request.FastMode
	}

	if !dara.IsNil(request.Restart) {
		body["restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceConfig"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/config/modifyInstanceConfig"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例配置
//
// @param request - ModifyInstanceConfigRequest
//
// @return ModifyInstanceConfigResponse
func (client *Client) ModifyInstanceConfig(request *ModifyInstanceConfigRequest) (_result *ModifyInstanceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyInstanceConfigResponse{}
	_body, _err := client.ModifyInstanceConfigWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 配置变更预检查，返回此次变更需要重启的计算组ID
//
// @param request - ModifyInstanceConfigPreCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceConfigPreCheckResponse
func (client *Client) ModifyInstanceConfigPreCheckWithOptions(request *ModifyInstanceConfigPreCheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyInstanceConfigPreCheckResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigsToAdd) {
		body["configsToAdd"] = request.ConfigsToAdd
	}

	if !dara.IsNil(request.ConfigsToDelete) {
		body["configsToDelete"] = request.ConfigsToDelete
	}

	if !dara.IsNil(request.ConfigsToUpdate) {
		body["configsToUpdate"] = request.ConfigsToUpdate
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceConfigPreCheck"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/config/modifyInstanceConfigPreCheck"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceConfigPreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置变更预检查，返回此次变更需要重启的计算组ID
//
// @param request - ModifyInstanceConfigPreCheckRequest
//
// @return ModifyInstanceConfigPreCheckResponse
func (client *Client) ModifyInstanceConfigPreCheck(request *ModifyInstanceConfigPreCheckRequest) (_result *ModifyInstanceConfigPreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyInstanceConfigPreCheckResponse{}
	_body, _err := client.ModifyInstanceConfigPreCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the number of nodes in a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
//   - You can modify the number of nodes in a warehouse of only StarRocks instances of Standard Edition.
//
//   - The instance must be in the Running state.
//
//   - The number of frontend nodes (FEs) cannot be an even number, and you cannot reduce the number of FE nodes.
//
// After you modify the number of nodes in a warehouse, the billing of nodes has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged based on the number of nodes.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the number of nodes before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyNodeNumberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodeNumberResponse
func (client *Client) ModifyNodeNumberWithOptions(request *ModifyNodeNumberRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyNodeNumberResponse, _err error) {
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

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.Parallelism) {
		query["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.TerminationGracePeriodSeconds) {
		query["TerminationGracePeriodSeconds"] = request.TerminationGracePeriodSeconds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNodeNumber"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifyNodeNumber"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNodeNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the number of nodes in a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of EMR Serverless StarRocks instances. Before you call this operation, take note of the following items:
//
//   - You can modify the number of nodes in a warehouse of only StarRocks instances of Standard Edition.
//
//   - The instance must be in the Running state.
//
//   - The number of frontend nodes (FEs) cannot be an even number, and you cannot reduce the number of FE nodes.
//
// After you modify the number of nodes in a warehouse, the billing of nodes has the following changes:
//
//   - Pay-as-you-go StarRocks instances: You are charged based on the number of nodes.
//
//   - Subscription StarRocks instances: You are charged additionally based on the price difference between the number of nodes before and after the change and the remaining days of the billing cycle. The billing cycle starts from 00:00 the next day until the end of the subscription period.
//
// @param request - ModifyNodeNumberRequest
//
// @return ModifyNodeNumberResponse
func (client *Client) ModifyNodeNumber(request *ModifyNodeNumberRequest) (_result *ModifyNodeNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyNodeNumberResponse{}
	_body, _err := client.ModifyNodeNumberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a precheck before you modify the number of nodes in a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// @param request - ModifyNodeNumberPreCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyNodeNumberPreCheckResponse
func (client *Client) ModifyNodeNumberPreCheckWithOptions(request *ModifyNodeNumberPreCheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyNodeNumberPreCheckResponse, _err error) {
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

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyNodeNumberPreCheck"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifyNodeNumberPreCheck"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyNodeNumberPreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a precheck before you modify the number of nodes in a warehouse of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// @param request - ModifyNodeNumberPreCheckRequest
//
// @return ModifyNodeNumberPreCheckResponse
func (client *Client) ModifyNodeNumberPreCheck(request *ModifyNodeNumberPreCheckRequest) (_result *ModifyNodeNumberPreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyNodeNumberPreCheckResponse{}
	_body, _err := client.ModifyNodeNumberPreCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of an E-MapReduce (EMR) Serverless StarRocks instance that the versions that you can upgrade to. The versions of a StarRocks instance include the major version and minor version. You can view the major version and minor version of a StarRocks instance in the Version Information section of the Instance Details tab in the EMR console. You can call this operation to query the minor versions or major versions that the versions that you can upgrade to.
//
// @param request - QueryUpgradableVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUpgradableVersionsResponse
func (client *Client) QueryUpgradableVersionsWithOptions(request *QueryUpgradableVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryUpgradableVersionsResponse, _err error) {
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

	if !dara.IsNil(request.Minor) {
		query["Minor"] = request.Minor
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUpgradableVersions"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/queryUpgradableVersions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUpgradableVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of an E-MapReduce (EMR) Serverless StarRocks instance that the versions that you can upgrade to. The versions of a StarRocks instance include the major version and minor version. You can view the major version and minor version of a StarRocks instance in the Version Information section of the Instance Details tab in the EMR console. You can call this operation to query the minor versions or major versions that the versions that you can upgrade to.
//
// @param request - QueryUpgradableVersionsRequest
//
// @return QueryUpgradableVersionsResponse
func (client *Client) QueryUpgradableVersions(request *QueryUpgradableVersionsRequest) (_result *QueryUpgradableVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryUpgradableVersionsResponse{}
	_body, _err := client.QueryUpgradableVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go E-MapReduce (EMR) Serverless StarRocks instance. To unsubscribe from a subscription instance, go to the Unsubscribe page of the Expenses and Costs console.
//
// Description:
//
// *
//
// **Warning:*	- After an instance is released, all physical resources used by the instance are recycled. Relevant data is erased and cannot be restored.
//
// @param request - ReleaseInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstance"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/cluster/release"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go E-MapReduce (EMR) Serverless StarRocks instance. To unsubscribe from a subscription instance, go to the Unsubscribe page of the Expenses and Costs console.
//
// Description:
//
// *
//
// **Warning:*	- After an instance is released, all physical resources used by the instance are recycled. Relevant data is erased and cannot be restored.
//
// @param request - ReleaseInstanceRequest
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// This operation is an asynchronous operation. After you call this operation to restart a StarRocks instance, the operation sets the status of the instance to Restarting and begins the restart process. When the status of the instance changes to Running, the instance is restarted.
//
// @param request - RestartInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FastMode) {
		query["FastMode"] = request.FastMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartInstance"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/restartCluster"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an E-MapReduce (EMR) Serverless StarRocks instance.
//
// Description:
//
// This operation is an asynchronous operation. After you call this operation to restart a StarRocks instance, the operation sets the status of the instance to Restarting and begins the restart process. When the status of the instance changes to Running, the instance is restarted.
//
// @param request - RestartInstanceRequest
//
// @return RestartInstanceResponse
func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启指定的node group
//
// @param request - RestartNodeGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartNodeGroupResponse
func (client *Client) RestartNodeGroupWithOptions(request *RestartNodeGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartNodeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FastMode) {
		query["FastMode"] = request.FastMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartNodeGroup"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/nodegroup/restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启指定的node group
//
// @param request - RestartNodeGroupRequest
//
// @return RestartNodeGroupResponse
func (client *Client) RestartNodeGroup(request *RestartNodeGroupRequest) (_result *RestartNodeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartNodeGroupResponse{}
	_body, _err := client.RestartNodeGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启集群中的节点
//
// @param request - RestartNodesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartNodesResponse
func (client *Client) RestartNodesWithOptions(request *RestartNodesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartNodesResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.RestartNodeGroups) {
		body["RestartNodeGroups"] = request.RestartNodeGroups
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartNodes"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/restart/restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启集群中的节点
//
// @param request - RestartNodesRequest
//
// @return RestartNodesResponse
func (client *Client) RestartNodes(request *RestartNodesRequest) (_result *RestartNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartNodesResponse{}
	_body, _err := client.RestartNodesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从备份中恢复实例
//
// @param request - RestoreInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreInstanceResponse
func (client *Client) RestoreInstanceWithOptions(request *RestoreInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestoreInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdminPassword) {
		body["AdminPassword"] = request.AdminPassword
	}

	if !dara.IsNil(request.AutoRenew) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.BackupTaskId) {
		body["BackupTaskId"] = request.BackupTaskId
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceName) {
		body["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.PayType) {
		body["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		body["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		body["Tags"] = request.Tags
	}

	if !dara.IsNil(request.VSwitches) {
		body["VSwitches"] = request.VSwitches
	}

	if !dara.IsNil(request.VpcId) {
		body["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestoreInstance"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/restore/restoreInstance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestoreInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从备份中恢复实例
//
// @param request - RestoreInstanceRequest
//
// @return RestoreInstanceResponse
func (client *Client) RestoreInstance(request *RestoreInstanceRequest) (_result *RestoreInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestoreInstanceResponse{}
	_body, _err := client.RestoreInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 该接口用于恢复来自openlake自动停机的实例。
//
// @param request - ResumeInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstanceWithOptions(request *ResumeInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeInstance"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/lifecycle/resumeInstance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 该接口用于恢复来自openlake自动停机的实例。
//
// @param request - ResumeInstanceRequest
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstance(request *ResumeInstanceRequest) (_result *ResumeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 回滚正在进行中的配置修改
//
// @param request - RollbackConfigModificationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackConfigModificationResponse
func (client *Client) RollbackConfigModificationWithOptions(request *RollbackConfigModificationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RollbackConfigModificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigHistoryId) {
		query["ConfigHistoryId"] = request.ConfigHistoryId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Restart) {
		query["Restart"] = request.Restart
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackConfigModification"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/config/rollbackConfigModification"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackConfigModificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 回滚正在进行中的配置修改
//
// @param request - RollbackConfigModificationRequest
//
// @return RollbackConfigModificationResponse
func (client *Client) RollbackConfigModification(request *RollbackConfigModificationRequest) (_result *RollbackConfigModificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RollbackConfigModificationResponse{}
	_body, _err := client.RollbackConfigModificationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a tag to a resource.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
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
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/tags"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Adds a tag to a resource.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 公网SLB开关
//
// @param request - TogglePublicSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TogglePublicSlbResponse
func (client *Client) TogglePublicSlbWithOptions(request *TogglePublicSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TogglePublicSlbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnablePublicSlb) {
		query["EnablePublicSlb"] = request.EnablePublicSlb
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TogglePublicSlb"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/gateway/togglePublicSlb"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TogglePublicSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 公网SLB开关
//
// @param request - TogglePublicSlbRequest
//
// @return TogglePublicSlbResponse
func (client *Client) TogglePublicSlb(request *TogglePublicSlbRequest) (_result *TogglePublicSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TogglePublicSlbResponse{}
	_body, _err := client.TogglePublicSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from specified resources.
//
// @param tmpReq - UnTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResourcesWithOptions(tmpReq *UnTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UnTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("ResourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TagKey) {
		request.TagKeyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKey, dara.String("TagKey"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["ResourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKeyShrink) {
		query["TagKey"] = request.TagKeyShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnTagResources"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/tags"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from specified resources.
//
// @param request - UnTagResourcesRequest
//
// @return UnTagResourcesResponse
func (client *Client) UnTagResources(request *UnTagResourcesRequest) (_result *UnTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnTagResourcesResponse{}
	_body, _err := client.UnTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新网关
//
// @param request - UpdateGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayResponse
func (client *Client) UpdateGatewayWithOptions(request *UpdateGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeNodeNumber) {
		query["FeNodeNumber"] = request.FeNodeNumber
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayName) {
		query["GatewayName"] = request.GatewayName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGateway"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/gateway/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新网关
//
// @param request - UpdateGatewayRequest
//
// @return UpdateGatewayResponse
func (client *Client) UpdateGateway(request *UpdateGatewayRequest) (_result *UpdateGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayResponse{}
	_body, _err := client.UpdateGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// @param request - UpdateInstanceNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceNameWithOptions(request *UpdateInstanceNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterName) {
		query["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceName"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/cluster/update_name"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name of an E-MapReduce (EMR) Serverless StarRocks instance.
//
// @param request - UpdateInstanceNameRequest
//
// @return UpdateInstanceNameResponse
func (client *Client) UpdateInstanceName(request *UpdateInstanceNameRequest) (_result *UpdateInstanceNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceNameResponse{}
	_body, _err := client.UpdateInstanceNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the version of an E-MapReduce (EMR) Serverless StarRocks instance. The versions of a StarRocks instance include the major version and minor version. You can view the major version and minor version of a StarRocks instance in the Version Information section of the Instance Details tab in the EMR console. This operation can be used to upgrade the minor version or major version of a StarRocks instance. You can call the QueryUpgradableVersions operation to query the versions that you can upgrade to.
//
// Description:
//
// The instance must be in the Running state when you call this operation.
//
// @param request - UpgradeVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeVersionResponse
func (client *Client) UpgradeVersionWithOptions(request *UpgradeVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FastMode) {
		query["FastMode"] = request.FastMode
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Minor) {
		query["Minor"] = request.Minor
	}

	if !dara.IsNil(request.TargetVersion) {
		query["TargetVersion"] = request.TargetVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeVersion"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/upgradeVersion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the version of an E-MapReduce (EMR) Serverless StarRocks instance. The versions of a StarRocks instance include the major version and minor version. You can view the major version and minor version of a StarRocks instance in the Version Information section of the Instance Details tab in the EMR console. This operation can be used to upgrade the minor version or major version of a StarRocks instance. You can call the QueryUpgradableVersions operation to query the versions that you can upgrade to.
//
// Description:
//
// The instance must be in the Running state when you call this operation.
//
// @param request - UpgradeVersionRequest
//
// @return UpgradeVersionResponse
func (client *Client) UpgradeVersion(request *UpgradeVersionRequest) (_result *UpgradeVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeVersionResponse{}
	_body, _err := client.UpgradeVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
