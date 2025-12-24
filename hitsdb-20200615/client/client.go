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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("hitsdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Changes a resource group to another.
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
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2020-06-15"),
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
// Changes a resource group to another.
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

// @param request - CheckLdpsColumnarIndexStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckLdpsColumnarIndexStatusResponse
func (client *Client) CheckLdpsColumnarIndexStatusWithOptions(request *CheckLdpsColumnarIndexStatusRequest, runtime *dara.RuntimeOptions) (_result *CheckLdpsColumnarIndexStatusResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckLdpsColumnarIndexStatus"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckLdpsColumnarIndexStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckLdpsColumnarIndexStatusRequest
//
// @return CheckLdpsColumnarIndexStatusResponse
func (client *Client) CheckLdpsColumnarIndexStatus(request *CheckLdpsColumnarIndexStatusRequest) (_result *CheckLdpsColumnarIndexStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckLdpsColumnarIndexStatusResponse{}
	_body, _err := client.CheckLdpsColumnarIndexStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建弹性伸缩配置
//
// @param tmpReq - CreateAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoScalingConfigResponse
func (client *Client) CreateAutoScalingConfigWithOptions(tmpReq *CreateAutoScalingConfigRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoScalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAutoScalingConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ScaleRuleList) {
		request.ScaleRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScaleRuleList, dara.String("ScaleRuleList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.EffectiveTimeEnd) {
		query["EffectiveTimeEnd"] = request.EffectiveTimeEnd
	}

	if !dara.IsNil(request.EffectiveTimeStart) {
		query["EffectiveTimeStart"] = request.EffectiveTimeStart
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodesMax) {
		query["NodesMax"] = request.NodesMax
	}

	if !dara.IsNil(request.NodesMin) {
		query["NodesMin"] = request.NodesMin
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

	if !dara.IsNil(request.ScaleRuleListShrink) {
		query["ScaleRuleList"] = request.ScaleRuleListShrink
	}

	if !dara.IsNil(request.ScaleType) {
		query["ScaleType"] = request.ScaleType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SpecId) {
		query["SpecId"] = request.SpecId
	}

	if !dara.IsNil(request.StorageCapacityMax) {
		query["StorageCapacityMax"] = request.StorageCapacityMax
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoScalingConfig"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoScalingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建弹性伸缩配置
//
// @param request - CreateAutoScalingConfigRequest
//
// @return CreateAutoScalingConfigResponse
func (client *Client) CreateAutoScalingConfig(request *CreateAutoScalingConfigRequest) (_result *CreateAutoScalingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAutoScalingConfigResponse{}
	_body, _err := client.CreateAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建弹性伸缩规则
//
// @param request - CreateAutoScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoScalingRuleResponse
func (client *Client) CreateAutoScalingRuleWithOptions(request *CreateAutoScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateAutoScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ObservationWindow) {
		query["ObservationWindow"] = request.ObservationWindow
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
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

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.ScaleInStep) {
		query["ScaleInStep"] = request.ScaleInStep
	}

	if !dara.IsNil(request.ScaleOutStep) {
		query["ScaleOutStep"] = request.ScaleOutStep
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TargetMetric) {
		query["TargetMetric"] = request.TargetMetric
	}

	if !dara.IsNil(request.TargetNodes) {
		query["TargetNodes"] = request.TargetNodes
	}

	if !dara.IsNil(request.ThresholdLower) {
		query["ThresholdLower"] = request.ThresholdLower
	}

	if !dara.IsNil(request.ThresholdUpper) {
		query["ThresholdUpper"] = request.ThresholdUpper
	}

	if !dara.IsNil(request.TriggerCronExpr) {
		query["TriggerCronExpr"] = request.TriggerCronExpr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAutoScalingRule"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAutoScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建弹性伸缩规则
//
// @param request - CreateAutoScalingRuleRequest
//
// @return CreateAutoScalingRuleResponse
func (client *Client) CreateAutoScalingRule(request *CreateAutoScalingRuleRequest) (_result *CreateAutoScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAutoScalingRuleResponse{}
	_body, _err := client.CreateAutoScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLdpsComputeGroupResponse
func (client *Client) CreateLdpsComputeGroupWithOptions(request *CreateLdpsComputeGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateLdpsComputeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Properties) {
		query["Properties"] = request.Properties
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLdpsComputeGroup"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLdpsComputeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateLdpsComputeGroupRequest
//
// @return CreateLdpsComputeGroupResponse
func (client *Client) CreateLdpsComputeGroup(request *CreateLdpsComputeGroupRequest) (_result *CreateLdpsComputeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLdpsComputeGroupResponse{}
	_body, _err := client.CreateLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Lindorm instance.
//
// Description:
//
// You must select at least one engine when you create a Lindorm instance. For more information about how to select the storage type and engine type when you create a Lindorm instance, see [Select engine types](https://help.aliyun.com/document_detail/181971.html) and [Select storage types](https://help.aliyun.com/document_detail/174643.html).
//
// @param request - CreateLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLindormInstanceResponse
func (client *Client) CreateLindormInstanceWithOptions(request *CreateLindormInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateLindormInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArbiterVSwitchId) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !dara.IsNil(request.ArbiterZoneId) {
		query["ArbiterZoneId"] = request.ArbiterZoneId
	}

	if !dara.IsNil(request.ArchVersion) {
		query["ArchVersion"] = request.ArchVersion
	}

	if !dara.IsNil(request.AutoRenewDuration) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.AutoRenewal) {
		query["AutoRenewal"] = request.AutoRenewal
	}

	if !dara.IsNil(request.ColdStorage) {
		query["ColdStorage"] = request.ColdStorage
	}

	if !dara.IsNil(request.CoreSingleStorage) {
		query["CoreSingleStorage"] = request.CoreSingleStorage
	}

	if !dara.IsNil(request.CoreSpec) {
		query["CoreSpec"] = request.CoreSpec
	}

	if !dara.IsNil(request.DiskCategory) {
		query["DiskCategory"] = request.DiskCategory
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.FilestoreNum) {
		query["FilestoreNum"] = request.FilestoreNum
	}

	if !dara.IsNil(request.FilestoreSpec) {
		query["FilestoreSpec"] = request.FilestoreSpec
	}

	if !dara.IsNil(request.InstanceAlias) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !dara.IsNil(request.InstanceStorage) {
		query["InstanceStorage"] = request.InstanceStorage
	}

	if !dara.IsNil(request.LindormNum) {
		query["LindormNum"] = request.LindormNum
	}

	if !dara.IsNil(request.LindormSpec) {
		query["LindormSpec"] = request.LindormSpec
	}

	if !dara.IsNil(request.LogDiskCategory) {
		query["LogDiskCategory"] = request.LogDiskCategory
	}

	if !dara.IsNil(request.LogNum) {
		query["LogNum"] = request.LogNum
	}

	if !dara.IsNil(request.LogSingleStorage) {
		query["LogSingleStorage"] = request.LogSingleStorage
	}

	if !dara.IsNil(request.LogSpec) {
		query["LogSpec"] = request.LogSpec
	}

	if !dara.IsNil(request.LtsNum) {
		query["LtsNum"] = request.LtsNum
	}

	if !dara.IsNil(request.LtsSpec) {
		query["LtsSpec"] = request.LtsSpec
	}

	if !dara.IsNil(request.MultiZoneCombination) {
		query["MultiZoneCombination"] = request.MultiZoneCombination
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PrimaryVSwitchId) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !dara.IsNil(request.PrimaryZoneId) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SolrNum) {
		query["SolrNum"] = request.SolrNum
	}

	if !dara.IsNil(request.SolrSpec) {
		query["SolrSpec"] = request.SolrSpec
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !dara.IsNil(request.StandbyZoneId) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !dara.IsNil(request.StreamNum) {
		query["StreamNum"] = request.StreamNum
	}

	if !dara.IsNil(request.StreamSpec) {
		query["StreamSpec"] = request.StreamSpec
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TsdbNum) {
		query["TsdbNum"] = request.TsdbNum
	}

	if !dara.IsNil(request.TsdbSpec) {
		query["TsdbSpec"] = request.TsdbSpec
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLindormInstance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Lindorm instance.
//
// Description:
//
// You must select at least one engine when you create a Lindorm instance. For more information about how to select the storage type and engine type when you create a Lindorm instance, see [Select engine types](https://help.aliyun.com/document_detail/181971.html) and [Select storage types](https://help.aliyun.com/document_detail/174643.html).
//
// @param request - CreateLindormInstanceRequest
//
// @return CreateLindormInstanceResponse
func (client *Client) CreateLindormInstance(request *CreateLindormInstanceRequest) (_result *CreateLindormInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLindormInstanceResponse{}
	_body, _err := client.CreateLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Lindorm实例
//
// @param request - CreateLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLindormV2InstanceResponse
func (client *Client) CreateLindormV2InstanceWithOptions(request *CreateLindormV2InstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateLindormV2InstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArbiterVSwitchId) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !dara.IsNil(request.ArbiterZoneId) {
		query["ArbiterZoneId"] = request.ArbiterZoneId
	}

	if !dara.IsNil(request.ArchVersion) {
		query["ArchVersion"] = request.ArchVersion
	}

	if !dara.IsNil(request.AutoRenewDuration) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !dara.IsNil(request.AutoRenewal) {
		query["AutoRenewal"] = request.AutoRenewal
	}

	if !dara.IsNil(request.CapacityStorageSize) {
		query["CapacityStorageSize"] = request.CapacityStorageSize
	}

	if !dara.IsNil(request.CloudStorageSize) {
		query["CloudStorageSize"] = request.CloudStorageSize
	}

	if !dara.IsNil(request.CloudStorageType) {
		query["CloudStorageType"] = request.CloudStorageType
	}

	if !dara.IsNil(request.ClusterMode) {
		query["ClusterMode"] = request.ClusterMode
	}

	if !dara.IsNil(request.ClusterPattern) {
		query["ClusterPattern"] = request.ClusterPattern
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EnableCapacityStorage) {
		query["EnableCapacityStorage"] = request.EnableCapacityStorage
	}

	if !dara.IsNil(request.EngineList) {
		query["EngineList"] = request.EngineList
	}

	if !dara.IsNil(request.InstanceAlias) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PrimaryVSwitchId) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !dara.IsNil(request.PrimaryZoneId) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !dara.IsNil(request.StandbyZoneId) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !dara.IsNil(request.VPCId) {
		query["VPCId"] = request.VPCId
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLindormV2Instance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLindormV2InstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Lindorm实例
//
// @param request - CreateLindormV2InstanceRequest
//
// @return CreateLindormV2InstanceResponse
func (client *Client) CreateLindormV2Instance(request *CreateLindormV2InstanceRequest) (_result *CreateLindormV2InstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLindormV2InstanceResponse{}
	_body, _err := client.CreateLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoScalingConfigResponse
func (client *Client) DeleteAutoScalingConfigWithOptions(request *DeleteAutoScalingConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoScalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoScalingConfig"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoScalingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAutoScalingConfigRequest
//
// @return DeleteAutoScalingConfigResponse
func (client *Client) DeleteAutoScalingConfig(request *DeleteAutoScalingConfigRequest) (_result *DeleteAutoScalingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutoScalingConfigResponse{}
	_body, _err := client.DeleteAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAutoScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoScalingRuleResponse
func (client *Client) DeleteAutoScalingRuleWithOptions(request *DeleteAutoScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteAutoScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAutoScalingRule"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAutoScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAutoScalingRuleRequest
//
// @return DeleteAutoScalingRuleResponse
func (client *Client) DeleteAutoScalingRule(request *DeleteAutoScalingRuleRequest) (_result *DeleteAutoScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAutoScalingRuleResponse{}
	_body, _err := client.DeleteAutoScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteCustomResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomResourceResponse
func (client *Client) DeleteCustomResourceWithOptions(request *DeleteCustomResourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomResourceResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomResource"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteCustomResourceRequest
//
// @return DeleteCustomResourceResponse
func (client *Client) DeleteCustomResource(request *DeleteCustomResourceRequest) (_result *DeleteCustomResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomResourceResponse{}
	_body, _err := client.DeleteCustomResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLdpsComputeGroupResponse
func (client *Client) DeleteLdpsComputeGroupWithOptions(request *DeleteLdpsComputeGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteLdpsComputeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLdpsComputeGroup"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLdpsComputeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteLdpsComputeGroupRequest
//
// @return DeleteLdpsComputeGroupResponse
func (client *Client) DeleteLdpsComputeGroup(request *DeleteLdpsComputeGroupRequest) (_result *DeleteLdpsComputeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLdpsComputeGroupResponse{}
	_body, _err := client.DeleteLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeployLdpsSemiManagedComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployLdpsSemiManagedComponentResponse
func (client *Client) DeployLdpsSemiManagedComponentWithOptions(request *DeployLdpsSemiManagedComponentRequest, runtime *dara.RuntimeOptions) (_result *DeployLdpsSemiManagedComponentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComponentName) {
		query["ComponentName"] = request.ComponentName
	}

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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployLdpsSemiManagedComponent"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployLdpsSemiManagedComponentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeployLdpsSemiManagedComponentRequest
//
// @return DeployLdpsSemiManagedComponentResponse
func (client *Client) DeployLdpsSemiManagedComponent(request *DeployLdpsSemiManagedComponentRequest) (_result *DeployLdpsSemiManagedComponentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployLdpsSemiManagedComponentResponse{}
	_body, _err := client.DeployLdpsSemiManagedComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the regions supported by Lindorm.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2020-06-15"),
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

// Summary:
//
// Obtains the regions supported by Lindorm.
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

// @param request - GetAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoScalingConfigResponse
func (client *Client) GetAutoScalingConfigWithOptions(request *GetAutoScalingConfigRequest, runtime *dara.RuntimeOptions) (_result *GetAutoScalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoScalingConfig"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoScalingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAutoScalingConfigRequest
//
// @return GetAutoScalingConfigResponse
func (client *Client) GetAutoScalingConfig(request *GetAutoScalingConfigRequest) (_result *GetAutoScalingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoScalingConfigResponse{}
	_body, _err := client.GetAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAutoScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoScalingRuleResponse
func (client *Client) GetAutoScalingRuleWithOptions(request *GetAutoScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *GetAutoScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoScalingRule"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAutoScalingRuleRequest
//
// @return GetAutoScalingRuleResponse
func (client *Client) GetAutoScalingRule(request *GetAutoScalingRuleRequest) (_result *GetAutoScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAutoScalingRuleResponse{}
	_body, _err := client.GetAutoScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetClientSourceIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientSourceIpResponse
func (client *Client) GetClientSourceIpWithOptions(request *GetClientSourceIpRequest, runtime *dara.RuntimeOptions) (_result *GetClientSourceIpResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClientSourceIp"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientSourceIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetClientSourceIpRequest
//
// @return GetClientSourceIpResponse
func (client *Client) GetClientSourceIp(request *GetClientSourceIpRequest) (_result *GetClientSourceIpResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClientSourceIpResponse{}
	_body, _err := client.GetClientSourceIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetEngineDefaultAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEngineDefaultAuthResponse
func (client *Client) GetEngineDefaultAuthWithOptions(request *GetEngineDefaultAuthRequest, runtime *dara.RuntimeOptions) (_result *GetEngineDefaultAuthResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEngineDefaultAuth"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEngineDefaultAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetEngineDefaultAuthRequest
//
// @return GetEngineDefaultAuthResponse
func (client *Client) GetEngineDefaultAuth(request *GetEngineDefaultAuthRequest) (_result *GetEngineDefaultAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEngineDefaultAuthResponse{}
	_body, _err := client.GetEngineDefaultAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the whitelists configured for a Lindorm instance.
//
// @param request - GetInstanceIpWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceIpWhiteListResponse
func (client *Client) GetInstanceIpWhiteListWithOptions(request *GetInstanceIpWhiteListRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceIpWhiteListResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceIpWhiteList"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the whitelists configured for a Lindorm instance.
//
// @param request - GetInstanceIpWhiteListRequest
//
// @return GetInstanceIpWhiteListResponse
func (client *Client) GetInstanceIpWhiteList(request *GetInstanceIpWhiteListRequest) (_result *GetInstanceIpWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceIpWhiteListResponse{}
	_body, _err := client.GetInstanceIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetInstanceSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceSecurityGroupsResponse
func (client *Client) GetInstanceSecurityGroupsWithOptions(request *GetInstanceSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceSecurityGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceSecurityGroups"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetInstanceSecurityGroupsRequest
//
// @return GetInstanceSecurityGroupsResponse
func (client *Client) GetInstanceSecurityGroups(request *GetInstanceSecurityGroupsRequest) (_result *GetInstanceSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceSecurityGroupsResponse{}
	_body, _err := client.GetInstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetInstanceSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceSummaryResponse
func (client *Client) GetInstanceSummaryWithOptions(request *GetInstanceSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceSummaryResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceSummary"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetInstanceSummaryRequest
//
// @return GetInstanceSummaryResponse
func (client *Client) GetInstanceSummary(request *GetInstanceSummaryRequest) (_result *GetInstanceSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceSummaryResponse{}
	_body, _err := client.GetInstanceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLdpsComputeGroupResponse
func (client *Client) GetLdpsComputeGroupWithOptions(request *GetLdpsComputeGroupRequest, runtime *dara.RuntimeOptions) (_result *GetLdpsComputeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLdpsComputeGroup"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLdpsComputeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetLdpsComputeGroupRequest
//
// @return GetLdpsComputeGroupResponse
func (client *Client) GetLdpsComputeGroup(request *GetLdpsComputeGroupRequest) (_result *GetLdpsComputeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLdpsComputeGroupResponse{}
	_body, _err := client.GetLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLdpsNamespacedQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLdpsNamespacedQuotaResponse
func (client *Client) GetLdpsNamespacedQuotaWithOptions(request *GetLdpsNamespacedQuotaRequest, runtime *dara.RuntimeOptions) (_result *GetLdpsNamespacedQuotaResponse, _err error) {
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

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLdpsNamespacedQuota"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLdpsNamespacedQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetLdpsNamespacedQuotaRequest
//
// @return GetLdpsNamespacedQuotaResponse
func (client *Client) GetLdpsNamespacedQuota(request *GetLdpsNamespacedQuotaRequest) (_result *GetLdpsNamespacedQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLdpsNamespacedQuotaResponse{}
	_body, _err := client.GetLdpsNamespacedQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLdpsResourceCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLdpsResourceCostResponse
func (client *Client) GetLdpsResourceCostWithOptions(request *GetLdpsResourceCostRequest, runtime *dara.RuntimeOptions) (_result *GetLdpsResourceCostResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLdpsResourceCost"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLdpsResourceCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetLdpsResourceCostRequest
//
// @return GetLdpsResourceCostResponse
func (client *Client) GetLdpsResourceCost(request *GetLdpsResourceCostRequest) (_result *GetLdpsResourceCostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLdpsResourceCostResponse{}
	_body, _err := client.GetLdpsResourceCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Lindorm引擎配置
//
// @param request - GetLindormEngineConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormEngineConfigResponse
func (client *Client) GetLindormEngineConfigWithOptions(request *GetLindormEngineConfigRequest, runtime *dara.RuntimeOptions) (_result *GetLindormEngineConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormEngineConfig"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormEngineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Lindorm引擎配置
//
// @param request - GetLindormEngineConfigRequest
//
// @return GetLindormEngineConfigResponse
func (client *Client) GetLindormEngineConfig(request *GetLindormEngineConfigRequest) (_result *GetLindormEngineConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormEngineConfigResponse{}
	_body, _err := client.GetLindormEngineConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of each storage type in a Lindorm instance.
//
// Description:
//
// If the version of the underlying storage engine in a Lindorm cluster is 4.1.9 or later, the storage usage values returned for the LStorageUsageList parameter prevail.
//
// @param request - GetLindormFsUsedDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormFsUsedDetailResponse
func (client *Client) GetLindormFsUsedDetailWithOptions(request *GetLindormFsUsedDetailRequest, runtime *dara.RuntimeOptions) (_result *GetLindormFsUsedDetailResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormFsUsedDetail"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormFsUsedDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of each storage type in a Lindorm instance.
//
// Description:
//
// If the version of the underlying storage engine in a Lindorm cluster is 4.1.9 or later, the storage usage values returned for the LStorageUsageList parameter prevail.
//
// @param request - GetLindormFsUsedDetailRequest
//
// @return GetLindormFsUsedDetailResponse
func (client *Client) GetLindormFsUsedDetail(request *GetLindormFsUsedDetailRequest) (_result *GetLindormFsUsedDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormFsUsedDetailResponse{}
	_body, _err := client.GetLindormFsUsedDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the detailed information about a Lindorm instance, including the instance type, billing method, and VPC.
//
// @param request - GetLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormInstanceResponse
func (client *Client) GetLindormInstanceWithOptions(request *GetLindormInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetLindormInstanceResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormInstance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the detailed information about a Lindorm instance, including the instance type, billing method, and VPC.
//
// @param request - GetLindormInstanceRequest
//
// @return GetLindormInstanceResponse
func (client *Client) GetLindormInstance(request *GetLindormInstanceRequest) (_result *GetLindormInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormInstanceResponse{}
	_body, _err := client.GetLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the engine types supported by the specified Lindorm instance.
//
// @param request - GetLindormInstanceEngineListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormInstanceEngineListResponse
func (client *Client) GetLindormInstanceEngineListWithOptions(request *GetLindormInstanceEngineListRequest, runtime *dara.RuntimeOptions) (_result *GetLindormInstanceEngineListResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormInstanceEngineList"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormInstanceEngineListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the engine types supported by the specified Lindorm instance.
//
// @param request - GetLindormInstanceEngineListRequest
//
// @return GetLindormInstanceEngineListResponse
func (client *Client) GetLindormInstanceEngineList(request *GetLindormInstanceEngineListRequest) (_result *GetLindormInstanceEngineListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormInstanceEngineListResponse{}
	_body, _err := client.GetLindormInstanceEngineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the instances that meet the specified conditions.
//
// @param request - GetLindormInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormInstanceListResponse
func (client *Client) GetLindormInstanceListWithOptions(request *GetLindormInstanceListRequest, runtime *dara.RuntimeOptions) (_result *GetLindormInstanceListResponse, _err error) {
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

	if !dara.IsNil(request.QueryStr) {
		query["QueryStr"] = request.QueryStr
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.ServiceType) {
		query["ServiceType"] = request.ServiceType
	}

	if !dara.IsNil(request.SupportEngine) {
		query["SupportEngine"] = request.SupportEngine
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormInstanceList"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instances that meet the specified conditions.
//
// @param request - GetLindormInstanceListRequest
//
// @return GetLindormInstanceListResponse
func (client *Client) GetLindormInstanceList(request *GetLindormInstanceListRequest) (_result *GetLindormInstanceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormInstanceListResponse{}
	_body, _err := client.GetLindormInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例详情
//
// @param request - GetLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2InstanceResponse
func (client *Client) GetLindormV2InstanceWithOptions(request *GetLindormV2InstanceRequest, runtime *dara.RuntimeOptions) (_result *GetLindormV2InstanceResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormV2Instance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormV2InstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例详情
//
// @param request - GetLindormV2InstanceRequest
//
// @return GetLindormV2InstanceResponse
func (client *Client) GetLindormV2Instance(request *GetLindormV2InstanceRequest) (_result *GetLindormV2InstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormV2InstanceResponse{}
	_body, _err := client.GetLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询新架构实例详情
//
// @param request - GetLindormV2InstanceDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2InstanceDetailsResponse
func (client *Client) GetLindormV2InstanceDetailsWithOptions(request *GetLindormV2InstanceDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetLindormV2InstanceDetailsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormV2InstanceDetails"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormV2InstanceDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询新架构实例详情
//
// @param request - GetLindormV2InstanceDetailsRequest
//
// @return GetLindormV2InstanceDetailsResponse
func (client *Client) GetLindormV2InstanceDetails(request *GetLindormV2InstanceDetailsRequest) (_result *GetLindormV2InstanceDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormV2InstanceDetailsResponse{}
	_body, _err := client.GetLindormV2InstanceDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLindormV2InstanceEngineListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2InstanceEngineListResponse
func (client *Client) GetLindormV2InstanceEngineListWithOptions(request *GetLindormV2InstanceEngineListRequest, runtime *dara.RuntimeOptions) (_result *GetLindormV2InstanceEngineListResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormV2InstanceEngineList"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormV2InstanceEngineListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetLindormV2InstanceEngineListRequest
//
// @return GetLindormV2InstanceEngineListResponse
func (client *Client) GetLindormV2InstanceEngineList(request *GetLindormV2InstanceEngineListRequest) (_result *GetLindormV2InstanceEngineListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormV2InstanceEngineListResponse{}
	_body, _err := client.GetLindormV2InstanceEngineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询新架构实例详情
//
// @param request - GetLindormV2InstanceForTerraformRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2InstanceForTerraformResponse
func (client *Client) GetLindormV2InstanceForTerraformWithOptions(request *GetLindormV2InstanceForTerraformRequest, runtime *dara.RuntimeOptions) (_result *GetLindormV2InstanceForTerraformResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormV2InstanceForTerraform"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormV2InstanceForTerraformResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询新架构实例详情
//
// @param request - GetLindormV2InstanceForTerraformRequest
//
// @return GetLindormV2InstanceForTerraformResponse
func (client *Client) GetLindormV2InstanceForTerraform(request *GetLindormV2InstanceForTerraformRequest) (_result *GetLindormV2InstanceForTerraformResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormV2InstanceForTerraformResponse{}
	_body, _err := client.GetLindormV2InstanceForTerraformWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询新架构实例安全组信息
//
// @param request - GetLindormV2InstanceSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2InstanceSecurityGroupsResponse
func (client *Client) GetLindormV2InstanceSecurityGroupsWithOptions(request *GetLindormV2InstanceSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *GetLindormV2InstanceSecurityGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormV2InstanceSecurityGroups"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormV2InstanceSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询新架构实例安全组信息
//
// @param request - GetLindormV2InstanceSecurityGroupsRequest
//
// @return GetLindormV2InstanceSecurityGroupsResponse
func (client *Client) GetLindormV2InstanceSecurityGroups(request *GetLindormV2InstanceSecurityGroupsRequest) (_result *GetLindormV2InstanceSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormV2InstanceSecurityGroupsResponse{}
	_body, _err := client.GetLindormV2InstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLindormV2StorageUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2StorageUsageResponse
func (client *Client) GetLindormV2StorageUsageWithOptions(request *GetLindormV2StorageUsageRequest, runtime *dara.RuntimeOptions) (_result *GetLindormV2StorageUsageResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormV2StorageUsage"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormV2StorageUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetLindormV2StorageUsageRequest
//
// @return GetLindormV2StorageUsageResponse
func (client *Client) GetLindormV2StorageUsage(request *GetLindormV2StorageUsageRequest) (_result *GetLindormV2StorageUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormV2StorageUsageResponse{}
	_body, _err := client.GetLindormV2StorageUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取流引擎信息
//
// @param request - GetLindormV2StreamEngineInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2StreamEngineInfoResponse
func (client *Client) GetLindormV2StreamEngineInfoWithOptions(request *GetLindormV2StreamEngineInfoRequest, runtime *dara.RuntimeOptions) (_result *GetLindormV2StreamEngineInfoResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLindormV2StreamEngineInfo"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLindormV2StreamEngineInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取流引擎信息
//
// @param request - GetLindormV2StreamEngineInfoRequest
//
// @return GetLindormV2StreamEngineInfoResponse
func (client *Client) GetLindormV2StreamEngineInfo(request *GetLindormV2StreamEngineInfoRequest) (_result *GetLindormV2StreamEngineInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLindormV2StreamEngineInfoResponse{}
	_body, _err := client.GetLindormV2StreamEngineInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询弹性伸缩配置
//
// @param tmpReq - ListAutoScalingConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoScalingConfigsResponse
func (client *Client) ListAutoScalingConfigsWithOptions(tmpReq *ListAutoScalingConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListAutoScalingConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAutoScalingConfigsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ScaleTypes) {
		request.ScaleTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScaleTypes, dara.String("ScaleTypes"), dara.String("json"))
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScaleTypesShrink) {
		query["ScaleTypes"] = request.ScaleTypesShrink
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAutoScalingConfigs"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutoScalingConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询弹性伸缩配置
//
// @param request - ListAutoScalingConfigsRequest
//
// @return ListAutoScalingConfigsResponse
func (client *Client) ListAutoScalingConfigs(request *ListAutoScalingConfigsRequest) (_result *ListAutoScalingConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAutoScalingConfigsResponse{}
	_body, _err := client.ListAutoScalingConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询伸缩记录
//
// @param tmpReq - ListAutoScalingRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoScalingRecordsResponse
func (client *Client) ListAutoScalingRecordsWithOptions(tmpReq *ListAutoScalingRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListAutoScalingRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAutoScalingRecordsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ScaleTypes) {
		request.ScaleTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScaleTypes, dara.String("ScaleTypes"), dara.String("json"))
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

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
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

	if !dara.IsNil(request.ScaleTypesShrink) {
		query["ScaleTypes"] = request.ScaleTypesShrink
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAutoScalingRecords"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutoScalingRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询伸缩记录
//
// @param request - ListAutoScalingRecordsRequest
//
// @return ListAutoScalingRecordsResponse
func (client *Client) ListAutoScalingRecords(request *ListAutoScalingRecordsRequest) (_result *ListAutoScalingRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAutoScalingRecordsResponse{}
	_body, _err := client.ListAutoScalingRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAutoScalingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoScalingRulesResponse
func (client *Client) ListAutoScalingRulesWithOptions(request *ListAutoScalingRulesRequest, runtime *dara.RuntimeOptions) (_result *ListAutoScalingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAutoScalingRules"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAutoScalingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListAutoScalingRulesRequest
//
// @return ListAutoScalingRulesResponse
func (client *Client) ListAutoScalingRules(request *ListAutoScalingRulesRequest) (_result *ListAutoScalingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAutoScalingRulesResponse{}
	_body, _err := client.ListAutoScalingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListLdpsComputeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLdpsComputeGroupsResponse
func (client *Client) ListLdpsComputeGroupsWithOptions(request *ListLdpsComputeGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListLdpsComputeGroupsResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLdpsComputeGroups"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLdpsComputeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListLdpsComputeGroupsRequest
//
// @return ListLdpsComputeGroupsResponse
func (client *Client) ListLdpsComputeGroups(request *ListLdpsComputeGroupsRequest) (_result *ListLdpsComputeGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLdpsComputeGroupsResponse{}
	_body, _err := client.ListLdpsComputeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags associated with the specified Lindorm instance.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2020-06-15"),
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
// Queries the tags associated with the specified Lindorm instance.
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
// 单可用区迁移多可用区基础版
//
// @param request - MigrateSingleZoneToMultiZoneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MigrateSingleZoneToMultiZoneResponse
func (client *Client) MigrateSingleZoneToMultiZoneWithOptions(request *MigrateSingleZoneToMultiZoneRequest, runtime *dara.RuntimeOptions) (_result *MigrateSingleZoneToMultiZoneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ArbitraryVSwitchId) {
		query["ArbitraryVSwitchId"] = request.ArbitraryVSwitchId
	}

	if !dara.IsNil(request.ArbitraryZoneId) {
		query["ArbitraryZoneId"] = request.ArbitraryZoneId
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.StandbyVSwitchId) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !dara.IsNil(request.StandbyZoneId) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MigrateSingleZoneToMultiZone"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MigrateSingleZoneToMultiZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 单可用区迁移多可用区基础版
//
// @param request - MigrateSingleZoneToMultiZoneRequest
//
// @return MigrateSingleZoneToMultiZoneResponse
func (client *Client) MigrateSingleZoneToMultiZone(request *MigrateSingleZoneToMultiZoneRequest) (_result *MigrateSingleZoneToMultiZoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MigrateSingleZoneToMultiZoneResponse{}
	_body, _err := client.MigrateSingleZoneToMultiZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改弹性伸缩配置
//
// @param tmpReq - ModifyAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoScalingConfigResponse
func (client *Client) ModifyAutoScalingConfigWithOptions(tmpReq *ModifyAutoScalingConfigRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoScalingConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ModifyAutoScalingConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ScaleRuleList) {
		request.ScaleRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScaleRuleList, dara.String("ScaleRuleList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.ConfigName) {
		query["ConfigName"] = request.ConfigName
	}

	if !dara.IsNil(request.EffectiveTimeEnd) {
		query["EffectiveTimeEnd"] = request.EffectiveTimeEnd
	}

	if !dara.IsNil(request.EffectiveTimeStart) {
		query["EffectiveTimeStart"] = request.EffectiveTimeStart
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.Engine) {
		query["Engine"] = request.Engine
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodesMax) {
		query["NodesMax"] = request.NodesMax
	}

	if !dara.IsNil(request.NodesMin) {
		query["NodesMin"] = request.NodesMin
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

	if !dara.IsNil(request.ScaleRuleListShrink) {
		query["ScaleRuleList"] = request.ScaleRuleListShrink
	}

	if !dara.IsNil(request.ScaleType) {
		query["ScaleType"] = request.ScaleType
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SpecId) {
		query["SpecId"] = request.SpecId
	}

	if !dara.IsNil(request.StorageCapacityMax) {
		query["StorageCapacityMax"] = request.StorageCapacityMax
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAutoScalingConfig"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAutoScalingConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改弹性伸缩配置
//
// @param request - ModifyAutoScalingConfigRequest
//
// @return ModifyAutoScalingConfigResponse
func (client *Client) ModifyAutoScalingConfig(request *ModifyAutoScalingConfigRequest) (_result *ModifyAutoScalingConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAutoScalingConfigResponse{}
	_body, _err := client.ModifyAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改弹性伸缩规则
//
// @param request - ModifyAutoScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoScalingRuleResponse
func (client *Client) ModifyAutoScalingRuleWithOptions(request *ModifyAutoScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyAutoScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigId) {
		query["ConfigId"] = request.ConfigId
	}

	if !dara.IsNil(request.Enabled) {
		query["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ObservationWindow) {
		query["ObservationWindow"] = request.ObservationWindow
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
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

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.ScaleInStep) {
		query["ScaleInStep"] = request.ScaleInStep
	}

	if !dara.IsNil(request.ScaleOutStep) {
		query["ScaleOutStep"] = request.ScaleOutStep
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SilenceTime) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TargetMetric) {
		query["TargetMetric"] = request.TargetMetric
	}

	if !dara.IsNil(request.TargetNodes) {
		query["TargetNodes"] = request.TargetNodes
	}

	if !dara.IsNil(request.ThresholdLower) {
		query["ThresholdLower"] = request.ThresholdLower
	}

	if !dara.IsNil(request.ThresholdUpper) {
		query["ThresholdUpper"] = request.ThresholdUpper
	}

	if !dara.IsNil(request.TriggerCronExpr) {
		query["TriggerCronExpr"] = request.TriggerCronExpr
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyAutoScalingRule"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyAutoScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改弹性伸缩规则
//
// @param request - ModifyAutoScalingRuleRequest
//
// @return ModifyAutoScalingRuleResponse
func (client *Client) ModifyAutoScalingRule(request *ModifyAutoScalingRuleRequest) (_result *ModifyAutoScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyAutoScalingRuleResponse{}
	_body, _err := client.ModifyAutoScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the billing method of the specified Lindorm instance.
//
// Description:
//
// You can call this operation to change the billing method of an instance to subscription or pay-as-you-go.
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/en/pricing-calculator?spm=a2c63.p38356.0.0.2b024c2adcHeXL&_p_lc=1#/commodity/hitsdb_lindormpre_public_intl) of Lindorm. Published on only international site (alibabacloud.com).
//
// @param request - ModifyInstancePayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstancePayTypeResponse
func (client *Client) ModifyInstancePayTypeWithOptions(request *ModifyInstancePayTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstancePayTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstancePayType"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstancePayTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the billing method of the specified Lindorm instance.
//
// Description:
//
// You can call this operation to change the billing method of an instance to subscription or pay-as-you-go.
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/en/pricing-calculator?spm=a2c63.p38356.0.0.2b024c2adcHeXL&_p_lc=1#/commodity/hitsdb_lindormpre_public_intl) of Lindorm. Published on only international site (alibabacloud.com).
//
// @param request - ModifyInstancePayTypeRequest
//
// @return ModifyInstancePayTypeResponse
func (client *Client) ModifyInstancePayType(request *ModifyInstancePayTypeRequest) (_result *ModifyInstancePayTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstancePayTypeResponse{}
	_body, _err := client.ModifyInstancePayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 变配实例
//
// @param request - ModifyLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLindormV2InstanceResponse
func (client *Client) ModifyLindormV2InstanceWithOptions(request *ModifyLindormV2InstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyLindormV2InstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloudStorageSize) {
		query["CloudStorageSize"] = request.CloudStorageSize
	}

	if !dara.IsNil(request.CloudStorageType) {
		query["CloudStorageType"] = request.CloudStorageType
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeGroupList) {
		query["NodeGroupList"] = request.NodeGroupList
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.UpgradeType) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLindormV2Instance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLindormV2InstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 变配实例
//
// @param request - ModifyLindormV2InstanceRequest
//
// @return ModifyLindormV2InstanceResponse
func (client *Client) ModifyLindormV2Instance(request *ModifyLindormV2InstanceRequest) (_result *ModifyLindormV2InstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLindormV2InstanceResponse{}
	_body, _err := client.ModifyLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新架构修改安全组接口
//
// @param request - ModifyLindormV2InstanceSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLindormV2InstanceSecurityGroupsResponse
func (client *Client) ModifyLindormV2InstanceSecurityGroupsWithOptions(request *ModifyLindormV2InstanceSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *ModifyLindormV2InstanceSecurityGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityGroups) {
		query["SecurityGroups"] = request.SecurityGroups
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLindormV2InstanceSecurityGroups"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLindormV2InstanceSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新架构修改安全组接口
//
// @param request - ModifyLindormV2InstanceSecurityGroupsRequest
//
// @return ModifyLindormV2InstanceSecurityGroupsResponse
func (client *Client) ModifyLindormV2InstanceSecurityGroups(request *ModifyLindormV2InstanceSecurityGroupsRequest) (_result *ModifyLindormV2InstanceSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLindormV2InstanceSecurityGroupsResponse{}
	_body, _err := client.ModifyLindormV2InstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Lindorm新版实例白名单
//
// @param request - ModifyLindormV2WhiteIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLindormV2WhiteIpListResponse
func (client *Client) ModifyLindormV2WhiteIpListWithOptions(request *ModifyLindormV2WhiteIpListRequest, runtime *dara.RuntimeOptions) (_result *ModifyLindormV2WhiteIpListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteGroup) {
		query["DeleteGroup"] = request.DeleteGroup
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WhiteIpList) {
		query["WhiteIpList"] = request.WhiteIpList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyLindormV2WhiteIpList"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyLindormV2WhiteIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Lindorm新版实例白名单
//
// @param request - ModifyLindormV2WhiteIpListRequest
//
// @return ModifyLindormV2WhiteIpListResponse
func (client *Client) ModifyLindormV2WhiteIpList(request *ModifyLindormV2WhiteIpListRequest) (_result *ModifyLindormV2WhiteIpListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyLindormV2WhiteIpListResponse{}
	_body, _err := client.ModifyLindormV2WhiteIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通计算引擎
//
// @param request - OpenComputeEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenComputeEngineResponse
func (client *Client) OpenComputeEngineWithOptions(request *OpenComputeEngineRequest, runtime *dara.RuntimeOptions) (_result *OpenComputeEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CpuLimit) {
		query["CpuLimit"] = request.CpuLimit
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemoryLimit) {
		query["MemoryLimit"] = request.MemoryLimit
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenComputeEngine"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenComputeEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通计算引擎
//
// @param request - OpenComputeEngineRequest
//
// @return OpenComputeEngineResponse
func (client *Client) OpenComputeEngine(request *OpenComputeEngineRequest) (_result *OpenComputeEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenComputeEngineResponse{}
	_body, _err := client.OpenComputeEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通计算引擎前置校验
//
// @param request - OpenComputePreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenComputePreCheckResponse
func (client *Client) OpenComputePreCheckWithOptions(request *OpenComputePreCheckRequest, runtime *dara.RuntimeOptions) (_result *OpenComputePreCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CpuLimit) {
		query["CpuLimit"] = request.CpuLimit
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MemoryLimit) {
		query["MemoryLimit"] = request.MemoryLimit
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenComputePreCheck"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenComputePreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通计算引擎前置校验
//
// @param request - OpenComputePreCheckRequest
//
// @return OpenComputePreCheckResponse
func (client *Client) OpenComputePreCheck(request *OpenComputePreCheckRequest) (_result *OpenComputePreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenComputePreCheckResponse{}
	_body, _err := client.OpenComputePreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开通列存索引
//
// @param request - OpenLdpsColumnarIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenLdpsColumnarIndexResponse
func (client *Client) OpenLdpsColumnarIndexWithOptions(request *OpenLdpsColumnarIndexRequest, runtime *dara.RuntimeOptions) (_result *OpenLdpsColumnarIndexResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenLdpsColumnarIndex"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenLdpsColumnarIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开通列存索引
//
// @param request - OpenLdpsColumnarIndexRequest
//
// @return OpenLdpsColumnarIndexResponse
func (client *Client) OpenLdpsColumnarIndex(request *OpenLdpsColumnarIndexRequest) (_result *OpenLdpsColumnarIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenLdpsColumnarIndexResponse{}
	_body, _err := client.OpenLdpsColumnarIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a Lindorm instance.
//
// @param request - ReleaseLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseLindormInstanceResponse
func (client *Client) ReleaseLindormInstanceWithOptions(request *ReleaseLindormInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseLindormInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Immediately) {
		query["Immediately"] = request.Immediately
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseLindormInstance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a Lindorm instance.
//
// @param request - ReleaseLindormInstanceRequest
//
// @return ReleaseLindormInstanceResponse
func (client *Client) ReleaseLindormInstance(request *ReleaseLindormInstanceRequest) (_result *ReleaseLindormInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseLindormInstanceResponse{}
	_body, _err := client.ReleaseLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 释放实例
//
// @param request - ReleaseLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseLindormV2InstanceResponse
func (client *Client) ReleaseLindormV2InstanceWithOptions(request *ReleaseLindormV2InstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseLindormV2InstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Immediately) {
		query["Immediately"] = request.Immediately
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseLindormV2Instance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseLindormV2InstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 释放实例
//
// @param request - ReleaseLindormV2InstanceRequest
//
// @return ReleaseLindormV2InstanceResponse
func (client *Client) ReleaseLindormV2Instance(request *ReleaseLindormV2InstanceRequest) (_result *ReleaseLindormV2InstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseLindormV2InstanceResponse{}
	_body, _err := client.ReleaseLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a subscription Lindorm instance.
//
// Description:
//
// You can call this operation to renew a subscription Lindorm instance for 1 to 9 months or 1 to 3 years.
//
// Before you call this operation, make sure that you fully understand the billing methods and pricing of Lindorm.
//
// @param request - RenewLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewLindormInstanceResponse
func (client *Client) RenewLindormInstanceWithOptions(request *RenewLindormInstanceRequest, runtime *dara.RuntimeOptions) (_result *RenewLindormInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RenewLindormInstance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a subscription Lindorm instance.
//
// Description:
//
// You can call this operation to renew a subscription Lindorm instance for 1 to 9 months or 1 to 3 years.
//
// Before you call this operation, make sure that you fully understand the billing methods and pricing of Lindorm.
//
// @param request - RenewLindormInstanceRequest
//
// @return RenewLindormInstanceResponse
func (client *Client) RenewLindormInstance(request *RenewLindormInstanceRequest) (_result *RenewLindormInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RenewLindormInstanceResponse{}
	_body, _err := client.RenewLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RestartLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartLdpsComputeGroupResponse
func (client *Client) RestartLdpsComputeGroupWithOptions(request *RestartLdpsComputeGroupRequest, runtime *dara.RuntimeOptions) (_result *RestartLdpsComputeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartLdpsComputeGroup"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartLdpsComputeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RestartLdpsComputeGroupRequest
//
// @return RestartLdpsComputeGroupResponse
func (client *Client) RestartLdpsComputeGroup(request *RestartLdpsComputeGroupRequest) (_result *RestartLdpsComputeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartLdpsComputeGroupResponse{}
	_body, _err := client.RestartLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetDefaultOlapComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultOlapComputeGroupResponse
func (client *Client) SetDefaultOlapComputeGroupWithOptions(request *SetDefaultOlapComputeGroupRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultOlapComputeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultOlapComputeGroup"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultOlapComputeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SetDefaultOlapComputeGroupRequest
//
// @return SetDefaultOlapComputeGroupResponse
func (client *Client) SetDefaultOlapComputeGroup(request *SetDefaultOlapComputeGroupRequest) (_result *SetDefaultOlapComputeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDefaultOlapComputeGroupResponse{}
	_body, _err := client.SetDefaultOlapComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the MySQL compatibility feature for a Lindorm instance.
//
// Description:
//
// Prerequisites
//
//   - The LindormTable version of your instance is 2.6.0 or later.
//
//   - The LindormTable of your instance supports LindormSQL V3. The value of the EnableLsqlVersionV3 parameter in the response of the GetLindormInstance operation is true for Lindorm instances purchased after Oct 24, 2023, which indicates that LindormSQL is supported by these instances by default. If you want to enable LindormSQL for instances purchased before Oct 24, 2023, contact the on-duty technical support.
//
// You can enable the MySQL compatibility feature for a Lindorm instance only when the instance meets the preceding requirements.
//
// @param request - SwitchLSQLV3MySQLServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchLSQLV3MySQLServiceResponse
func (client *Client) SwitchLSQLV3MySQLServiceWithOptions(request *SwitchLSQLV3MySQLServiceRequest, runtime *dara.RuntimeOptions) (_result *SwitchLSQLV3MySQLServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionType) {
		query["ActionType"] = request.ActionType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchLSQLV3MySQLService"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchLSQLV3MySQLServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the MySQL compatibility feature for a Lindorm instance.
//
// Description:
//
// Prerequisites
//
//   - The LindormTable version of your instance is 2.6.0 or later.
//
//   - The LindormTable of your instance supports LindormSQL V3. The value of the EnableLsqlVersionV3 parameter in the response of the GetLindormInstance operation is true for Lindorm instances purchased after Oct 24, 2023, which indicates that LindormSQL is supported by these instances by default. If you want to enable LindormSQL for instances purchased before Oct 24, 2023, contact the on-duty technical support.
//
// You can enable the MySQL compatibility feature for a Lindorm instance only when the instance meets the preceding requirements.
//
// @param request - SwitchLSQLV3MySQLServiceRequest
//
// @return SwitchLSQLV3MySQLServiceResponse
func (client *Client) SwitchLSQLV3MySQLService(request *SwitchLSQLV3MySQLServiceRequest) (_result *SwitchLSQLV3MySQLServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SwitchLSQLV3MySQLServiceResponse{}
	_body, _err := client.SwitchLSQLV3MySQLServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates tags with a single or multiple Lindorm instances.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2020-06-15"),
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
// Associates tags with a single or multiple Lindorm instances.
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
// Removes tags from a Lindorm instance.
//
// Description:
//
// If a tag is not added to any Lindorm instance, it is deleted.
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2020-06-15"),
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
// Removes tags from a Lindorm instance.
//
// Description:
//
// If a tag is not added to any Lindorm instance, it is deleted.
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
// Configures an IP address whitelist for a Lindorm instance.
//
// @param request - UpdateInstanceIpWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceIpWhiteListResponse
func (client *Client) UpdateInstanceIpWhiteListWithOptions(request *UpdateInstanceIpWhiteListRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceIpWhiteListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Delete) {
		query["Delete"] = request.Delete
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityIpList) {
		query["SecurityIpList"] = request.SecurityIpList
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceIpWhiteList"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures an IP address whitelist for a Lindorm instance.
//
// @param request - UpdateInstanceIpWhiteListRequest
//
// @return UpdateInstanceIpWhiteListResponse
func (client *Client) UpdateInstanceIpWhiteList(request *UpdateInstanceIpWhiteListRequest) (_result *UpdateInstanceIpWhiteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceIpWhiteListResponse{}
	_body, _err := client.UpdateInstanceIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateInstanceSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceSecurityGroupsResponse
func (client *Client) UpdateInstanceSecurityGroupsWithOptions(request *UpdateInstanceSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceSecurityGroupsResponse, _err error) {
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

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityGroups) {
		query["SecurityGroups"] = request.SecurityGroups
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceSecurityGroups"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateInstanceSecurityGroupsRequest
//
// @return UpdateInstanceSecurityGroupsResponse
func (client *Client) UpdateInstanceSecurityGroups(request *UpdateInstanceSecurityGroupsRequest) (_result *UpdateInstanceSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceSecurityGroupsResponse{}
	_body, _err := client.UpdateInstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLdpsComputeGroupResponse
func (client *Client) UpdateLdpsComputeGroupWithOptions(request *UpdateLdpsComputeGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateLdpsComputeGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Properties) {
		query["Properties"] = request.Properties
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLdpsComputeGroup"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLdpsComputeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateLdpsComputeGroupRequest
//
// @return UpdateLdpsComputeGroupResponse
func (client *Client) UpdateLdpsComputeGroup(request *UpdateLdpsComputeGroupRequest) (_result *UpdateLdpsComputeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLdpsComputeGroupResponse{}
	_body, _err := client.UpdateLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例名称或删除保护
//
// @param request - UpdateLindormInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLindormInstanceAttributeResponse
func (client *Client) UpdateLindormInstanceAttributeWithOptions(request *UpdateLindormInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *UpdateLindormInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.InstanceAlias) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLindormInstanceAttribute"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLindormInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例名称或删除保护
//
// @param request - UpdateLindormInstanceAttributeRequest
//
// @return UpdateLindormInstanceAttributeResponse
func (client *Client) UpdateLindormInstanceAttribute(request *UpdateLindormInstanceAttributeRequest) (_result *UpdateLindormInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLindormInstanceAttributeResponse{}
	_body, _err := client.UpdateLindormInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新LindormV2Instance
//
// @param request - UpdateLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLindormV2InstanceResponse
func (client *Client) UpdateLindormV2InstanceWithOptions(request *UpdateLindormV2InstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateLindormV2InstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CapacityStorageSize) {
		query["CapacityStorageSize"] = request.CapacityStorageSize
	}

	if !dara.IsNil(request.CloudStorageSize) {
		query["CloudStorageSize"] = request.CloudStorageSize
	}

	if !dara.IsNil(request.CloudStorageType) {
		query["CloudStorageType"] = request.CloudStorageType
	}

	if !dara.IsNil(request.EnableCapacityStorage) {
		query["EnableCapacityStorage"] = request.EnableCapacityStorage
	}

	if !dara.IsNil(request.EngineList) {
		query["EngineList"] = request.EngineList
	}

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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLindormV2Instance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLindormV2InstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新LindormV2Instance
//
// @param request - UpdateLindormV2InstanceRequest
//
// @return UpdateLindormV2InstanceResponse
func (client *Client) UpdateLindormV2Instance(request *UpdateLindormV2InstanceRequest) (_result *UpdateLindormV2InstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLindormV2InstanceResponse{}
	_body, _err := client.UpdateLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateLindormV2InstanceParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLindormV2InstanceParameterResponse
func (client *Client) UpdateLindormV2InstanceParameterWithOptions(request *UpdateLindormV2InstanceParameterRequest, runtime *dara.RuntimeOptions) (_result *UpdateLindormV2InstanceParameterResponse, _err error) {
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

	if !dara.IsNil(request.ParameterKey) {
		query["ParameterKey"] = request.ParameterKey
	}

	if !dara.IsNil(request.ParameterValue) {
		query["ParameterValue"] = request.ParameterValue
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.UpdateType) {
		query["UpdateType"] = request.UpdateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLindormV2InstanceParameter"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLindormV2InstanceParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateLindormV2InstanceParameterRequest
//
// @return UpdateLindormV2InstanceParameterResponse
func (client *Client) UpdateLindormV2InstanceParameter(request *UpdateLindormV2InstanceParameterRequest) (_result *UpdateLindormV2InstanceParameterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLindormV2InstanceParameterResponse{}
	_body, _err := client.UpdateLindormV2InstanceParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改Lindorm新版实例白名单分组列表
//
// @param request - UpdateLindormV2WhiteIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLindormV2WhiteIpListResponse
func (client *Client) UpdateLindormV2WhiteIpListWithOptions(request *UpdateLindormV2WhiteIpListRequest, runtime *dara.RuntimeOptions) (_result *UpdateLindormV2WhiteIpListResponse, _err error) {
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.WhiteIpGroupList) {
		query["WhiteIpGroupList"] = request.WhiteIpGroupList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLindormV2WhiteIpList"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLindormV2WhiteIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改Lindorm新版实例白名单分组列表
//
// @param request - UpdateLindormV2WhiteIpListRequest
//
// @return UpdateLindormV2WhiteIpListResponse
func (client *Client) UpdateLindormV2WhiteIpList(request *UpdateLindormV2WhiteIpListRequest) (_result *UpdateLindormV2WhiteIpListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLindormV2WhiteIpListResponse{}
	_body, _err := client.UpdateLindormV2WhiteIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades, scales up, or enable cold storage for a Lindorm instance.
//
// Description:
//
// For more information about how to select the storage type and engine type when you create a Lindorm instance, see [Select engine typpes](https://help.aliyun.com/document_detail/181971.html) and [Select storage types](https://help.aliyun.com/document_detail/174643.html).
//
// @param request - UpgradeLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeLindormInstanceResponse
func (client *Client) UpgradeLindormInstanceWithOptions(request *UpgradeLindormInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpgradeLindormInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterStorage) {
		query["ClusterStorage"] = request.ClusterStorage
	}

	if !dara.IsNil(request.ColdStorage) {
		query["ColdStorage"] = request.ColdStorage
	}

	if !dara.IsNil(request.CoreSingleStorage) {
		query["CoreSingleStorage"] = request.CoreSingleStorage
	}

	if !dara.IsNil(request.FilestoreNum) {
		query["FilestoreNum"] = request.FilestoreNum
	}

	if !dara.IsNil(request.FilestoreSpec) {
		query["FilestoreSpec"] = request.FilestoreSpec
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LindormNum) {
		query["LindormNum"] = request.LindormNum
	}

	if !dara.IsNil(request.LindormSpec) {
		query["LindormSpec"] = request.LindormSpec
	}

	if !dara.IsNil(request.LogNum) {
		query["LogNum"] = request.LogNum
	}

	if !dara.IsNil(request.LogSingleStorage) {
		query["LogSingleStorage"] = request.LogSingleStorage
	}

	if !dara.IsNil(request.LogSpec) {
		query["LogSpec"] = request.LogSpec
	}

	if !dara.IsNil(request.LtsCoreNum) {
		query["LtsCoreNum"] = request.LtsCoreNum
	}

	if !dara.IsNil(request.LtsCoreSpec) {
		query["LtsCoreSpec"] = request.LtsCoreSpec
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

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.SolrNum) {
		query["SolrNum"] = request.SolrNum
	}

	if !dara.IsNil(request.SolrSpec) {
		query["SolrSpec"] = request.SolrSpec
	}

	if !dara.IsNil(request.StreamNum) {
		query["StreamNum"] = request.StreamNum
	}

	if !dara.IsNil(request.StreamSpec) {
		query["StreamSpec"] = request.StreamSpec
	}

	if !dara.IsNil(request.TsdbNum) {
		query["TsdbNum"] = request.TsdbNum
	}

	if !dara.IsNil(request.TsdbSpec) {
		query["TsdbSpec"] = request.TsdbSpec
	}

	if !dara.IsNil(request.UpgradeType) {
		query["UpgradeType"] = request.UpgradeType
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeLindormInstance"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeLindormInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades, scales up, or enable cold storage for a Lindorm instance.
//
// Description:
//
// For more information about how to select the storage type and engine type when you create a Lindorm instance, see [Select engine typpes](https://help.aliyun.com/document_detail/181971.html) and [Select storage types](https://help.aliyun.com/document_detail/174643.html).
//
// @param request - UpgradeLindormInstanceRequest
//
// @return UpgradeLindormInstanceResponse
func (client *Client) UpgradeLindormInstance(request *UpgradeLindormInstanceRequest) (_result *UpgradeLindormInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeLindormInstanceResponse{}
	_body, _err := client.UpgradeLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpgradeLindormV2StreamEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeLindormV2StreamEngineResponse
func (client *Client) UpgradeLindormV2StreamEngineWithOptions(request *UpgradeLindormV2StreamEngineRequest, runtime *dara.RuntimeOptions) (_result *UpgradeLindormV2StreamEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomConfig) {
		query["CustomConfig"] = request.CustomConfig
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Quantity) {
		query["Quantity"] = request.Quantity
	}

	if !dara.IsNil(request.ResourceGroupName) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SecurityToken) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !dara.IsNil(request.Spec) {
		query["Spec"] = request.Spec
	}

	if !dara.IsNil(request.SpecId) {
		query["SpecId"] = request.SpecId
	}

	if !dara.IsNil(request.UpgradeType) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeLindormV2StreamEngine"),
		Version:     dara.String("2020-06-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeLindormV2StreamEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpgradeLindormV2StreamEngineRequest
//
// @return UpgradeLindormV2StreamEngineResponse
func (client *Client) UpgradeLindormV2StreamEngine(request *UpgradeLindormV2StreamEngineRequest) (_result *UpgradeLindormV2StreamEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeLindormV2StreamEngineResponse{}
	_body, _err := client.UpgradeLindormV2StreamEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
