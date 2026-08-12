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
		"ap-southeast-1": dara.String("starrocks.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5": dara.String("starrocks.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":     dara.String("starrocks.cn-beijing.aliyuncs.com"),
		"cn-hangzhou":    dara.String("starrocks.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":    dara.String("starrocks.cn-hongkong.aliyuncs.com"),
		"cn-qingdao":     dara.String("starrocks.cn-qingdao.aliyuncs.com"),
		"cn-shanghai":    dara.String("starrocks.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":    dara.String("starrocks.cn-shenzhen.aliyuncs.com"),
		"cn-wulanchabu":  dara.String("starrocks.cn-wulanchabu.aliyuncs.com"),
		"cn-zhangjiakou": dara.String("starrocks.cn-zhangjiakou.aliyuncs.com"),
		"us-west-1":      dara.String("starrocks.us-west-1.aliyuncs.com"),
		"us-east-1":      dara.String("starrocks.us-east-1.aliyuncs.com"),
		"eu-central-1":   dara.String("starrocks.eu-central-1.aliyuncs.com"),
	}
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
// Adds a backup policy.
//
// @param request - AddBackupPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBackupPolicyResponse
func (client *Client) AddBackupPolicyWithOptions(request *AddBackupPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireDays) {
		body["ExpireDays"] = request.ExpireDays
	}

	if !dara.IsNil(request.Hour) {
		body["Hour"] = request.Hour
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Minute) {
		body["Minute"] = request.Minute
	}

	if !dara.IsNil(request.RecurrenceType) {
		body["RecurrenceType"] = request.RecurrenceType
	}

	if !dara.IsNil(request.RecurrenceValues) {
		body["RecurrenceValues"] = request.RecurrenceValues
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		body["TimeoutSeconds"] = request.TimeoutSeconds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddBackupPolicy"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/backupRestore/policy/add"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a backup policy.
//
// @param request - AddBackupPolicyRequest
//
// @return AddBackupPolicyResponse
func (client *Client) AddBackupPolicy(request *AddBackupPolicyRequest) (_result *AddBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddBackupPolicyResponse{}
	_body, _err := client.AddBackupPolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a gateway.
//
// Description:
//
// The AddGateway operation requires software stack version 1.7.6 or later and at least three front-end (FE) nodes.
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
// Creates a gateway.
//
// Description:
//
// The AddGateway operation requires software stack version 1.7.6 or later and at least three front-end (FE) nodes.
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
// This operation changes the resource group of a Serverless StarRocks instance.
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
// This operation changes the resource group of a Serverless StarRocks instance.
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
// Check the inventory of resources.
//
// @param request - CheckInventoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckInventoryResponse
func (client *Client) CheckInventoryWithOptions(request *CheckInventoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckInventoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterInfo) {
		query["ClusterInfo"] = request.ClusterInfo
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckInventory"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/check/inventory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckInventoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check the inventory of resources.
//
// @param request - CheckInventoryRequest
//
// @return CheckInventoryResponse
func (client *Client) CheckInventory(request *CheckInventoryRequest) (_result *CheckInventoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckInventoryResponse{}
	_body, _err := client.CheckInventoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an Agent compute group.
//
// @param request - CreateAgentResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentResourceResponse
func (client *Client) CreateAgentResourceWithOptions(request *CreateAgentResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentResourceResponse, _err error) {
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

	if !dara.IsNil(request.Cu) {
		query["Cu"] = request.Cu
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.PricingCycle) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		query["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentResource"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/lifecycle/createAgentNodeGroup"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an Agent compute group.
//
// @param request - CreateAgentResourceRequest
//
// @return CreateAgentResourceResponse
func (client *Client) CreateAgentResource(request *CreateAgentResourceRequest) (_result *CreateAgentResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentResourceResponse{}
	_body, _err := client.CreateAgentResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing method and <props="china">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
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

	if !dara.IsNil(request.EnableAiFunction) {
		body["EnableAiFunction"] = request.EnableAiFunction
	}

	if !dara.IsNil(request.EnableMultiAz) {
		body["EnableMultiAz"] = request.EnableMultiAz
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
// Creates a Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing method and <props="china">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
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
// Create a scaling rule.
//
// @param request - CreateScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScalingRuleResponse
func (client *Client) CreateScalingRuleWithOptions(request *CreateScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScalingRule"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/scalingRule/createScalingRule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a scaling rule.
//
// @param request - CreateScalingRuleRequest
//
// @return CreateScalingRuleResponse
func (client *Client) CreateScalingRule(request *CreateScalingRuleRequest) (_result *CreateScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScalingRuleResponse{}
	_body, _err := client.CreateScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation creates the AliyunServiceRoleForEMRStarRocks role.
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
// This operation creates the AliyunServiceRoleForEMRStarRocks role.
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
// Deletes a data backup.
//
// Description:
//
// Deletes a data backup.
//
// @param request - DeleteBackupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackupWithOptions(request *DeleteBackupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBackupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupTaskId) {
		query["BackupTaskId"] = request.BackupTaskId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackup"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/backup/manage/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data backup.
//
// Description:
//
// Deletes a data backup.
//
// @param request - DeleteBackupRequest
//
// @return DeleteBackupResponse
func (client *Client) DeleteBackup(request *DeleteBackupRequest) (_result *DeleteBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBackupResponse{}
	_body, _err := client.DeleteBackupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a backup policy.
//
// Description:
//
// Deletes a backup policy.
//
// @param request - DeleteBackupPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBackupPolicyResponse
func (client *Client) DeleteBackupPolicyWithOptions(request *DeleteBackupPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBackupPolicy"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/backupRestore/policy/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a backup policy.
//
// Description:
//
// Deletes a backup policy.
//
// @param request - DeleteBackupPolicyRequest
//
// @return DeleteBackupPolicyResponse
func (client *Client) DeleteBackupPolicy(request *DeleteBackupPolicyRequest) (_result *DeleteBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBackupPolicyResponse{}
	_body, _err := client.DeleteBackupPolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Deleting a gateway
//
// Description:
//
// Deletes a gateway group. After the gateway group is deleted, its FE nodes are automatically assigned to the default gateway group.
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
// # Deleting a gateway
//
// Description:
//
// Deletes a gateway group. After the gateway group is deleted, its FE nodes are automatically assigned to the default gateway group.
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
// Deletes an internal network whitelist group.
//
// @param request - DeleteInnerIpWhitelistGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInnerIpWhitelistGroupResponse
func (client *Client) DeleteInnerIpWhitelistGroupWithOptions(request *DeleteInnerIpWhitelistGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInnerIpWhitelistGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InnerIpWhitelistGroupId) {
		body["InnerIpWhitelistGroupId"] = request.InnerIpWhitelistGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInnerIpWhitelistGroup"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/securityGroup/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInnerIpWhitelistGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an internal network whitelist group.
//
// @param request - DeleteInnerIpWhitelistGroupRequest
//
// @return DeleteInnerIpWhitelistGroupResponse
func (client *Client) DeleteInnerIpWhitelistGroup(request *DeleteInnerIpWhitelistGroupRequest) (_result *DeleteInnerIpWhitelistGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInnerIpWhitelistGroupResponse{}
	_body, _err := client.DeleteInnerIpWhitelistGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scaling rule.
//
// @param request - DeleteScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScalingRuleResponse
func (client *Client) DeleteScalingRuleWithOptions(request *DeleteScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.ScalingRuleId) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScalingRule"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/scalingRule/deleteScalingRule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scaling rule.
//
// @param request - DeleteScalingRuleRequest
//
// @return DeleteScalingRuleResponse
func (client *Client) DeleteScalingRule(request *DeleteScalingRuleRequest) (_result *DeleteScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteScalingRuleResponse{}
	_body, _err := client.DeleteScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available zones in a region.
//
// @param request - DescribeAvailableZonesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAvailableZonesResponse
func (client *Client) DescribeAvailableZonesWithOptions(request *DescribeAvailableZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeAvailableZonesResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAvailableZones"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/zone/describeZones"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAvailableZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available zones in a region.
//
// @param request - DescribeAvailableZonesRequest
//
// @return DescribeAvailableZonesResponse
func (client *Client) DescribeAvailableZones(request *DescribeAvailableZonesRequest) (_result *DescribeAvailableZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAvailableZonesResponse{}
	_body, _err := client.DescribeAvailableZonesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve details of backup policies.
//
// Description:
//
// Retrieve details of backup policies.
//
// @param request - DescribeBackupPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupPoliciesResponse
func (client *Client) DescribeBackupPoliciesWithOptions(request *DescribeBackupPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeBackupPoliciesResponse, _err error) {
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

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackupPolicies"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/backupRestore/policy/describe"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve details of backup policies.
//
// Description:
//
// Retrieve details of backup policies.
//
// @param request - DescribeBackupPoliciesRequest
//
// @return DescribeBackupPoliciesResponse
func (client *Client) DescribeBackupPolicies(request *DescribeBackupPoliciesRequest) (_result *DescribeBackupPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeBackupPoliciesResponse{}
	_body, _err := client.DescribeBackupPoliciesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtain backup details
//
// Description:
//
// This is an asynchronous interface. Instance restarts are not immediate. After a successful call, the instance first enters the restarting state. The instance is successfully restarted when its status returns to running.
//
// @param request - DescribeBackupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackupsWithOptions(request *DescribeBackupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeBackupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupTaskId) {
		query["BackupTaskId"] = request.BackupTaskId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.Statuses) {
		query["Statuses"] = request.Statuses
	}

	if !dara.IsNil(request.TimePeriodEndTime) {
		query["TimePeriodEndTime"] = request.TimePeriodEndTime
	}

	if !dara.IsNil(request.TimePeriodStartTime) {
		query["TimePeriodStartTime"] = request.TimePeriodStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeBackups"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/backup/manage/describe"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeBackupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain backup details
//
// Description:
//
// This is an asynchronous interface. Instance restarts are not immediate. After a successful call, the instance first enters the restarting state. The instance is successfully restarted when its status returns to running.
//
// @param request - DescribeBackupsRequest
//
// @return DescribeBackupsResponse
func (client *Client) DescribeBackups(request *DescribeBackupsRequest) (_result *DescribeBackupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeBackupsResponse{}
	_body, _err := client.DescribeBackupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Querying instance configuration history
//
// Description:
//
// This is an asynchronous API, meaning the instance restart is not immediate. After a successful call, the instance enters the `restarting` state, and the process completes when its state changes to `running`.
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
// # Querying instance configuration history
//
// Description:
//
// This is an asynchronous API, meaning the instance restart is not immediate. After a successful call, the instance enters the `restarting` state, and the process completes when its state changes to `running`.
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
// Retrieve the cluster event name.
//
// @param request - DescribeEventNamesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEventNamesResponse
func (client *Client) DescribeEventNamesWithOptions(request *DescribeEventNamesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeEventNamesResponse, _err error) {
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
		Action:      dara.String("DescribeEventNames"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/event/describeEventNames"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeEventNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the cluster event name.
//
// @param request - DescribeEventNamesRequest
//
// @return DescribeEventNamesResponse
func (client *Client) DescribeEventNames(request *DescribeEventNamesRequest) (_result *DescribeEventNamesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeEventNamesResponse{}
	_body, _err := client.DescribeEventNamesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Describes the whitelist groups for a private network.
//
// @param request - DescribeInnerIpWhitelistGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInnerIpWhitelistGroupsResponse
func (client *Client) DescribeInnerIpWhitelistGroupsWithOptions(request *DescribeInnerIpWhitelistGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInnerIpWhitelistGroupsResponse, _err error) {
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
		Action:      dara.String("DescribeInnerIpWhitelistGroups"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/securityGroup/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInnerIpWhitelistGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Describes the whitelist groups for a private network.
//
// @param request - DescribeInnerIpWhitelistGroupsRequest
//
// @return DescribeInnerIpWhitelistGroupsResponse
func (client *Client) DescribeInnerIpWhitelistGroups(request *DescribeInnerIpWhitelistGroupsRequest) (_result *DescribeInnerIpWhitelistGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInnerIpWhitelistGroupsResponse{}
	_body, _err := client.DescribeInnerIpWhitelistGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instance configurations.
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
// Queries instance configurations.
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
// Retrieve instance health diagnosis results.
//
// @param request - DescribeInstanceDiagnosisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceDiagnosisResultResponse
func (client *Client) DescribeInstanceDiagnosisResultWithOptions(request *DescribeInstanceDiagnosisResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInstanceDiagnosisResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Dimension) {
		query["Dimension"] = request.Dimension
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReportDate) {
		query["ReportDate"] = request.ReportDate
	}

	if !dara.IsNil(request.Statuses) {
		query["Statuses"] = request.Statuses
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInstanceDiagnosisResult"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/diagnosis/describe"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceDiagnosisResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve instance health diagnosis results.
//
// @param request - DescribeInstanceDiagnosisResultRequest
//
// @return DescribeInstanceDiagnosisResultResponse
func (client *Client) DescribeInstanceDiagnosisResult(request *DescribeInstanceDiagnosisResultRequest) (_result *DescribeInstanceDiagnosisResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceDiagnosisResultResponse{}
	_body, _err := client.DescribeInstanceDiagnosisResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the meta token for a StarRocks instance.
//
// @param request - DescribeInstanceMetaTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceMetaTokenResponse
func (client *Client) DescribeInstanceMetaTokenWithOptions(request *DescribeInstanceMetaTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeInstanceMetaTokenResponse, _err error) {
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
		Action:      dara.String("DescribeInstanceMetaToken"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/migration/getMetaToken"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeInstanceMetaTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the meta token for a StarRocks instance.
//
// @param request - DescribeInstanceMetaTokenRequest
//
// @return DescribeInstanceMetaTokenResponse
func (client *Client) DescribeInstanceMetaToken(request *DescribeInstanceMetaTokenRequest) (_result *DescribeInstanceMetaTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInstanceMetaTokenResponse{}
	_body, _err := client.DescribeInstanceMetaTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This operation queries Serverless StarRocks instances. You can filter the instances by criteria such as instance name and tags.
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
// This operation queries Serverless StarRocks instances. You can filter the instances by criteria such as instance name and tags.
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
// Retrieves node group information.
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
// Retrieves node group information.
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
// You can view the list of available regions.
//
// Description:
//
// This operation is asynchronous. After you call this operation, the instance enters a restarting state but does not restart immediately. The restart is complete when the instance status changes to running.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/region/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// You can view the list of available regions.
//
// Description:
//
// This operation is asynchronous. After you call this operation, the instance enters a restarting state but does not restart immediately. The restart is complete when the instance status changes to running.
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get resource constraint configurations.
//
// @param request - DescribeResourceConstraintsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceConstraintsResponse
func (client *Client) DescribeResourceConstraintsWithOptions(request *DescribeResourceConstraintsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeResourceConstraintsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Architecture) {
		query["Architecture"] = request.Architecture
	}

	if !dara.IsNil(request.PackageType) {
		query["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.RunMode) {
		query["RunMode"] = request.RunMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeResourceConstraints"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/describeResourceConstraints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeResourceConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get resource constraint configurations.
//
// @param request - DescribeResourceConstraintsRequest
//
// @return DescribeResourceConstraintsResponse
func (client *Client) DescribeResourceConstraints(request *DescribeResourceConstraintsRequest) (_result *DescribeResourceConstraintsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourceConstraintsResponse{}
	_body, _err := client.DescribeResourceConstraintsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtain the system time zone of a StarRocks instance
//
// @param request - DescribeSystemTimezoneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSystemTimezoneResponse
func (client *Client) DescribeSystemTimezoneWithOptions(request *DescribeSystemTimezoneRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeSystemTimezoneResponse, _err error) {
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
		Action:      dara.String("DescribeSystemTimezone"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/timezone/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSystemTimezoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtain the system time zone of a StarRocks instance
//
// @param request - DescribeSystemTimezoneRequest
//
// @return DescribeSystemTimezoneResponse
func (client *Client) DescribeSystemTimezone(request *DescribeSystemTimezoneRequest) (_result *DescribeSystemTimezoneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSystemTimezoneResponse{}
	_body, _err := client.DescribeSystemTimezoneWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of time-triggered scaling rules.
//
// @param request - DescribeTimeTriggerScalingRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTimeTriggerScalingRulesResponse
func (client *Client) DescribeTimeTriggerScalingRulesWithOptions(request *DescribeTimeTriggerScalingRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeTimeTriggerScalingRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTimeTriggerScalingRules"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/scalingRule/describeTimeTriggerScalingRules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeTimeTriggerScalingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of time-triggered scaling rules.
//
// @param request - DescribeTimeTriggerScalingRulesRequest
//
// @return DescribeTimeTriggerScalingRulesResponse
func (client *Client) DescribeTimeTriggerScalingRules(request *DescribeTimeTriggerScalingRulesRequest) (_result *DescribeTimeTriggerScalingRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTimeTriggerScalingRulesResponse{}
	_body, _err := client.DescribeTimeTriggerScalingRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables SSL for a StarRocks connection.
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
// Disables SSL for a StarRocks connection.
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
// Enables the internal Server Load Balancer (SLB) for the default gateway of an EMR Serverless StarRocks instance.
//
// Description:
//
// The instance must be in the Running state when you call this operation.
//
// @param request - EnableInternalSlbRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableInternalSlbResponse
func (client *Client) EnableInternalSlbWithOptions(request *EnableInternalSlbRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableInternalSlbResponse, _err error) {
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
		Action:      dara.String("EnableInternalSlb"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/gateway/enableInternalSlb"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableInternalSlbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the internal Server Load Balancer (SLB) for the default gateway of an EMR Serverless StarRocks instance.
//
// Description:
//
// The instance must be in the Running state when you call this operation.
//
// @param request - EnableInternalSlbRequest
//
// @return EnableInternalSlbResponse
func (client *Client) EnableInternalSlb(request *EnableInternalSlbRequest) (_result *EnableInternalSlbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableInternalSlbResponse{}
	_body, _err := client.EnableInternalSlbWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables Multi-AZ deployment.
//
// @param request - EnableMultiAzRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableMultiAzResponse
func (client *Client) EnableMultiAzWithOptions(request *EnableMultiAzRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EnableMultiAzResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Observers) {
		body["observers"] = request.Observers
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		body["promotionOptionNo"] = request.PromotionOptionNo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableMultiAz"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/lifecycle/enableMultiAz"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableMultiAzResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables Multi-AZ deployment.
//
// @param request - EnableMultiAzRequest
//
// @return EnableMultiAzResponse
func (client *Client) EnableMultiAz(request *EnableMultiAzRequest) (_result *EnableMultiAzResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableMultiAzResponse{}
	_body, _err := client.EnableMultiAzWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables SSL for a StarRocks connection.
//
// Description:
//
// You can call this operation only when the instance is in the Running state.
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
// Enables SSL for a StarRocks connection.
//
// Description:
//
// You can call this operation only when the instance is in the Running state.
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
// Retrieves the feature gates for a StarRocks cluster instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
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
// Retrieves the feature gates for a StarRocks cluster instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
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
// Retrieve the feature gate for a StarRocks compute group instance.
//
// @param request - GetNodeGroupFeatureGateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeGroupFeatureGateResponse
func (client *Client) GetNodeGroupFeatureGateWithOptions(request *GetNodeGroupFeatureGateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetNodeGroupFeatureGateResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeGroupFeatureGate"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/features/nodeGroupFeatureGate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeGroupFeatureGateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the feature gate for a StarRocks compute group instance.
//
// @param request - GetNodeGroupFeatureGateRequest
//
// @return GetNodeGroupFeatureGateResponse
func (client *Client) GetNodeGroupFeatureGate(request *GetNodeGroupFeatureGateRequest) (_result *GetNodeGroupFeatureGateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetNodeGroupFeatureGateResponse{}
	_body, _err := client.GetNodeGroupFeatureGateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enable read/write splitting. The Leader FE node handles write requests, and other nodes handle read requests.
//
// Description:
//
// The instance must be in the running state when you call this operation.
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
// Enable read/write splitting. The Leader FE node handles write requests, and other nodes handle read requests.
//
// Description:
//
// The instance must be in the running state when you call this operation.
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
// Lists gateways.
//
// Description:
//
// Lists the gateways of a cluster.
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
// Lists gateways.
//
// Description:
//
// Lists the gateways of a cluster.
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
// Retrieves the details of an operation.
//
// @param request - ListOperationActivityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationActivityResponse
func (client *Client) ListOperationActivityWithOptions(request *ListOperationActivityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOperationActivityResponse, _err error) {
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

	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationActivity"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/operation/listOperationActivity"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an operation.
//
// @param request - ListOperationActivityRequest
//
// @return ListOperationActivityResponse
func (client *Client) ListOperationActivity(request *ListOperationActivityRequest) (_result *ListOperationActivityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOperationActivityResponse{}
	_body, _err := client.ListOperationActivityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the operation history of a cluster.
//
// @param request - ListOperationHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationHistoryResponse
func (client *Client) ListOperationHistoryWithOptions(request *ListOperationHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOperationHistoryResponse, _err error) {
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

	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	if !dara.IsNil(request.OperationStatus) {
		query["OperationStatus"] = request.OperationStatus
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
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
		Action:      dara.String("ListOperationHistory"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/operation/listOperationHistory"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the operation history of a cluster.
//
// @param request - ListOperationHistoryRequest
//
// @return ListOperationHistoryResponse
func (client *Client) ListOperationHistory(request *ListOperationHistoryRequest) (_result *ListOperationHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOperationHistoryResponse{}
	_body, _err := client.ListOperationHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the ssl certificate details for a cluster.
//
// Description:
//
// Retrieves the ssl certificate details for a cluster.
//
// @param request - ListSSLDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSSLDetailsResponse
func (client *Client) ListSSLDetailsWithOptions(request *ListSSLDetailsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSSLDetailsResponse, _err error) {
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
		Action:      dara.String("ListSSLDetails"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/listSSLDetails"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSSLDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the ssl certificate details for a cluster.
//
// Description:
//
// Retrieves the ssl certificate details for a cluster.
//
// @param request - ListSSLDetailsRequest
//
// @return ListSSLDetailsResponse
func (client *Client) ListSSLDetails(request *ListSSLDetailsRequest) (_result *ListSSLDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSSLDetailsResponse{}
	_body, _err := client.ListSSLDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Change the billing method for an instance
//
// Description:
//
// Before you call this operation, review the billing methods for Serverless StarRocks and the <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P).
//
// When you call this operation, note the following:
//
// - Only standard instances support changing the number of compute group Compute Units (CUs). Entry-level instances do not support this change.
//
// - Only instances of the standard compute group specification type support increasing the number of disks.
//
// - The instance must be in the Running state.
//
// After you change the CU count, billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new CU count.
//
// - Subscription: During the current billing cycle, you pay the difference between the old and new configurations. This amount is calculated based on the number of days remaining in the subscription period, starting from 00:00 of the following day.
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
// # Change the billing method for an instance
//
// Description:
//
// Before you call this operation, review the billing methods for Serverless StarRocks and the <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P).
//
// When you call this operation, note the following:
//
// - Only standard instances support changing the number of compute group Compute Units (CUs). Entry-level instances do not support this change.
//
// - Only instances of the standard compute group specification type support increasing the number of disks.
//
// - The instance must be in the Running state.
//
// After you change the CU count, billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new CU count.
//
// - Subscription: During the current billing cycle, you pay the difference between the old and new configurations. This amount is calculated based on the number of days remaining in the subscription period, starting from 00:00 of the following day.
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
// You can use this API to change the number of Compute Units (CUs) in a compute group of a Serverless StarRocks instance.
//
// Description:
//
// Before you call this API, make sure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// When you call this API, note the following:
//
// - This operation is supported only for Standard Edition instances. It is not supported for Starter Edition instances.
//
// - You can increase the number of disks only for instances with a standard compute group specification.
//
// - The instance must be in the Running state.
//
// After you change the number of CUs, the billing for the instance is updated as follows:
//
// - Pay-as-you-go: You are billed based on the new number of CUs.
//
// - Subscription: A supplementary fee is charged. This fee is calculated based on the price difference between the old and new configurations and the remaining subscription period. The remaining period starts at 00:00 on the following day and ends when the subscription expires.
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
// You can use this API to change the number of Compute Units (CUs) in a compute group of a Serverless StarRocks instance.
//
// Description:
//
// Before you call this API, make sure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// When you call this API, note the following:
//
// - This operation is supported only for Standard Edition instances. It is not supported for Starter Edition instances.
//
// - You can increase the number of disks only for instances with a standard compute group specification.
//
// - The instance must be in the Running state.
//
// After you change the number of CUs, the billing for the instance is updated as follows:
//
// - Pay-as-you-go: You are billed based on the new number of CUs.
//
// - Subscription: A supplementary fee is charged. This fee is calculated based on the price difference between the old and new configurations and the remaining subscription period. The remaining period starts at 00:00 on the following day and ends when the subscription expires.
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
// Increases the number of disks for the nodes in a compute group of a Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure you fully understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// Note the following:
//
// - The number of disks can be increased only for standard instances. This operation is not supported for entry-level instances.
//
// - The number of disks can be increased only for instances whose compute group is the Standard Edition (standard).
//
// - The instance must be in the Running state.
//
// After you increase the number of disks, billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new disk type.
//
// - Subscription: A supplemental fee is calculated. This fee is based on the price difference between the old and new configurations and the remaining days in the billing cycle. The remaining period starts from 00:00 on the following day and ends when the subscription expires.
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
// Increases the number of disks for the nodes in a compute group of a Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure you fully understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// Note the following:
//
// - The number of disks can be increased only for standard instances. This operation is not supported for entry-level instances.
//
// - The number of disks can be increased only for instances whose compute group is the Standard Edition (standard).
//
// - The instance must be in the Running state.
//
// After you increase the number of disks, billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new disk type.
//
// - Subscription: A supplemental fee is calculated. This fee is based on the price difference between the old and new configurations and the remaining days in the billing cycle. The remaining period starts from 00:00 on the following day and ends when the subscription expires.
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
// This operation modifies the disk performance level of the nodes in a compute group of a Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, understand the billing methods and <props="china">[pricing](https://www.aliyun.com/price/product?spm=openapi-amp.newDocPublishment.0.0.1837281f3hbi2d#/ecs/detail/vm)<props="intl">[pricing](https://www.alibabacloud.com/zh/product/ecs?_p_lc=1#pricing) of Serverless StarRocks.
//
// Note the following when you call this operation:
//
// - You can upgrade or downgrade disks only for standard instances. This operation is not supported for Starter Edition instances.
//
// - Disk upgrades and downgrades are supported only for instances that have a Standard Edition (standard) compute group.
//
// - The instance must be in the Running state.
//
// - You cannot downgrade the disk performance level (PL) to PL0.
//
// - The performance level of an Enhanced SSD (ESSD) is limited by its capacity. If you cannot upgrade the performance level, increase the disk capacity and try again. For more information, see <props="china">[ESSDs](https://help.aliyun.com/zh/ecs/user-guide/essds)<props="intl">[ESSDs](https://www.alibabacloud.com/help/en/ecs/user-guide/essds).
//
// After you change the disk configuration, the billing is adjusted as follows:
//
// - Pay-as-you-go: You are charged based on the new disk type.
//
// - Subscription: A supplemental fee is calculated based on the price difference between the new and old configurations for the remainder of the billing cycle. The remaining period starts at 00:00 on the next day and ends when the subscription expires.
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
// This operation modifies the disk performance level of the nodes in a compute group of a Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, understand the billing methods and <props="china">[pricing](https://www.aliyun.com/price/product?spm=openapi-amp.newDocPublishment.0.0.1837281f3hbi2d#/ecs/detail/vm)<props="intl">[pricing](https://www.alibabacloud.com/zh/product/ecs?_p_lc=1#pricing) of Serverless StarRocks.
//
// Note the following when you call this operation:
//
// - You can upgrade or downgrade disks only for standard instances. This operation is not supported for Starter Edition instances.
//
// - Disk upgrades and downgrades are supported only for instances that have a Standard Edition (standard) compute group.
//
// - The instance must be in the Running state.
//
// - You cannot downgrade the disk performance level (PL) to PL0.
//
// - The performance level of an Enhanced SSD (ESSD) is limited by its capacity. If you cannot upgrade the performance level, increase the disk capacity and try again. For more information, see <props="china">[ESSDs](https://help.aliyun.com/zh/ecs/user-guide/essds)<props="intl">[ESSDs](https://www.alibabacloud.com/help/en/ecs/user-guide/essds).
//
// After you change the disk configuration, the billing is adjusted as follows:
//
// - Pay-as-you-go: You are charged based on the new disk type.
//
// - Subscription: A supplemental fee is calculated based on the price difference between the new and old configurations for the remainder of the billing cycle. The remaining period starts at 00:00 on the next day and ends when the subscription expires.
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
// Scales out the disk of a compute group node for a Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, ensure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// Note the following:
//
// - Only standard instances support disk scale-out. Entry-level instances do not.
//
// - Only instances that have a compute group of the standard specification type support disk scale-out.
//
// - The instance must be in the Running state.
//
// After you scale out the disk, your billing changes as follows:
//
// - Pay-as-you-go: You are charged based on the new disk size.
//
// - Subscription: You must pay an upgrade fee. The fee is calculated based on the price difference between the old and new configurations and the remaining subscription period. The remaining subscription period is calculated starting from 00:00 on the next day.
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
// Scales out the disk of a compute group node for a Serverless StarRocks instance.
//
// Description:
//
// Before you call this operation, ensure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// Note the following:
//
// - Only standard instances support disk scale-out. Entry-level instances do not.
//
// - Only instances that have a compute group of the standard specification type support disk scale-out.
//
// - The instance must be in the Running state.
//
// After you scale out the disk, your billing changes as follows:
//
// - Pay-as-you-go: You are charged based on the new disk size.
//
// - Subscription: You must pay an upgrade fee. The fee is calculated based on the price difference between the old and new configurations and the remaining subscription period. The remaining subscription period is calculated starting from 00:00 on the next day.
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
// # Change the disk type for nodes in a compute group
//
// Description:
//
// Before you call this operation, review the Serverless StarRocks billing model and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P).
//
// - Only standard instances support disk scaling. Starter instances do not support disk scaling.
//
// - Only instances with a compute group specification type of Standard Edition support disk scaling.
//
// - The instance must be in the Running state.
//
// After disk scaling, billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new disk size.
//
// - Subscription: During the billing cycle, the additional fee is calculated based on the price difference between the old and new configurations and the remaining days, from 00:00 the next day to the end of the validity period.
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
// # Change the disk type for nodes in a compute group
//
// Description:
//
// Before you call this operation, review the Serverless StarRocks billing model and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P).
//
// - Only standard instances support disk scaling. Starter instances do not support disk scaling.
//
// - Only instances with a compute group specification type of Standard Edition support disk scaling.
//
// - The instance must be in the Running state.
//
// After disk scaling, billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new disk size.
//
// - Subscription: During the billing cycle, the additional fee is calculated based on the price difference between the old and new configurations and the remaining days, from 00:00 the next day to the end of the validity period.
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
// Modifies the /etc/hosts file.
//
// @param request - ModifyHostAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyHostAliasResponse
func (client *Client) ModifyHostAliasWithOptions(request *ModifyHostAliasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyHostAliasResponse, _err error) {
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
	if !dara.IsNil(request.HostAliases) {
		body["hostAliases"] = request.HostAliases
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyHostAlias"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/network/modifyHostAlias"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyHostAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the /etc/hosts file.
//
// @param request - ModifyHostAliasRequest
//
// @return ModifyHostAliasResponse
func (client *Client) ModifyHostAlias(request *ModifyHostAliasRequest) (_result *ModifyHostAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyHostAliasResponse{}
	_body, _err := client.ModifyHostAliasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This API modifies the configuration of a Serverless StarRocks instance.
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
// This API modifies the configuration of a Serverless StarRocks instance.
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
// This operation pre-checks modifications to the instance configuration of a Serverless StarRocks instance.
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
// This operation pre-checks modifications to the instance configuration of a Serverless StarRocks instance.
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
// Modifies the maintenance window for a StarRocks instance.
//
// @param request - ModifyMaintainableTimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMaintainableTimeResponse
func (client *Client) ModifyMaintainableTimeWithOptions(request *ModifyMaintainableTimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyMaintainableTimeResponse, _err error) {
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

	if !dara.IsNil(request.MaintainableTimePeriod) {
		query["MaintainableTimePeriod"] = request.MaintainableTimePeriod
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyMaintainableTime"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/modifyMaintainableTime"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMaintainableTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the maintenance window for a StarRocks instance.
//
// @param request - ModifyMaintainableTimeRequest
//
// @return ModifyMaintainableTimeResponse
func (client *Client) ModifyMaintainableTime(request *ModifyMaintainableTimeRequest) (_result *ModifyMaintainableTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyMaintainableTimeResponse{}
	_body, _err := client.ModifyMaintainableTimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the number of nodes in a compute group of a Serverless StarRocks instance.
//
// Description:
//
// Before using this operation, make sure that you fully understand the billing method and <props="china">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// When calling this operation, note the following:
//
// - Only Standard Standard instances support modifying the number of compute group nodes. Basic Standard instances do not support this operation.
//
// - The instance must be in the Running state.
//
// - The number of FE nodes cannot be an even number, and FE nodes do not support scale-in.
//
// After modifying the node count, billing changes are as follows:
//
// - Pay-as-you-go: Billed based on the new node count.
//
// - Subscription: Within the billing cycle, the additional fee is calculated based on the price difference between the old and new configurations and the remaining days (from 00:00 of the next day to the end of the validity period).
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
// Modifies the number of nodes in a compute group of a Serverless StarRocks instance.
//
// Description:
//
// Before using this operation, make sure that you fully understand the billing method and <props="china">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// When calling this operation, note the following:
//
// - Only Standard Standard instances support modifying the number of compute group nodes. Basic Standard instances do not support this operation.
//
// - The instance must be in the Running state.
//
// - The number of FE nodes cannot be an even number, and FE nodes do not support scale-in.
//
// After modifying the node count, billing changes are as follows:
//
// - Pay-as-you-go: Billed based on the new node count.
//
// - Subscription: Within the billing cycle, the additional fee is calculated based on the price difference between the old and new configurations and the remaining days (from 00:00 of the next day to the end of the validity period).
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
// Modifies an Auto Scaling rule.
//
// @param request - ModifyScalingRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScalingRuleResponse
func (client *Client) ModifyScalingRuleWithOptions(request *ModifyScalingRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyScalingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewTriggerType) {
		query["NewTriggerType"] = request.NewTriggerType
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.OldTriggerType) {
		query["OldTriggerType"] = request.OldTriggerType
	}

	if !dara.IsNil(request.Rule) {
		query["Rule"] = request.Rule
	}

	if !dara.IsNil(request.ScalingRuleId) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScalingRule"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/scalingRule/modifyScalingRule"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an Auto Scaling rule.
//
// @param request - ModifyScalingRuleRequest
//
// @return ModifyScalingRuleResponse
func (client *Client) ModifyScalingRule(request *ModifyScalingRuleRequest) (_result *ModifyScalingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.ModifyScalingRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modify the node specifications type of the compute group.
//
// @param request - ModifySpecTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySpecTypeResponse
func (client *Client) ModifySpecTypeWithOptions(request *ModifySpecTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifySpecTypeResponse, _err error) {
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

	if !dara.IsNil(request.TargetSpecType) {
		query["TargetSpecType"] = request.TargetSpecType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySpecType"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifySpecType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySpecTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modify the node specifications type of the compute group.
//
// @param request - ModifySpecTypeRequest
//
// @return ModifySpecTypeResponse
func (client *Client) ModifySpecType(request *ModifySpecTypeRequest) (_result *ModifySpecTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifySpecTypeResponse{}
	_body, _err := client.ModifySpecTypeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs a precheck to modify the node specification type for a compute group.
//
// @param request - ModifySpecTypePreCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySpecTypePreCheckResponse
func (client *Client) ModifySpecTypePreCheckWithOptions(request *ModifySpecTypePreCheckRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifySpecTypePreCheckResponse, _err error) {
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

	if !dara.IsNil(request.TargetSpecType) {
		query["TargetSpecType"] = request.TargetSpecType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySpecTypePreCheck"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/resourceChange/modifySpecTypePreCheck"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySpecTypePreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a precheck to modify the node specification type for a compute group.
//
// @param request - ModifySpecTypePreCheckRequest
//
// @return ModifySpecTypePreCheckResponse
func (client *Client) ModifySpecTypePreCheck(request *ModifySpecTypePreCheckRequest) (_result *ModifySpecTypePreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifySpecTypePreCheckResponse{}
	_body, _err := client.ModifySpecTypePreCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the password for a StarRocks user.
//
// @param request - ModifyUserPasswordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserPasswordResponse
func (client *Client) ModifyUserPasswordWithOptions(request *ModifyUserPasswordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyUserPasswordResponse, _err error) {
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

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyUserPassword"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/password/modify"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the password for a StarRocks user.
//
// @param request - ModifyUserPasswordRequest
//
// @return ModifyUserPasswordResponse
func (client *Client) ModifyUserPassword(request *ModifyUserPasswordRequest) (_result *ModifyUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyUserPasswordResponse{}
	_body, _err := client.ModifyUserPasswordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price for enabling multi-zone deployment.
//
// @param request - QueryEnableMultiAzPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryEnableMultiAzPriceResponse
func (client *Client) QueryEnableMultiAzPriceWithOptions(request *QueryEnableMultiAzPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryEnableMultiAzPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Observers) {
		body["observers"] = request.Observers
	}

	if !dara.IsNil(request.PromotionOptionNo) {
		body["promotionOptionNo"] = request.PromotionOptionNo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryEnableMultiAzPrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/priceInquiry/enableMultiAz"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryEnableMultiAzPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price for enabling multi-zone deployment.
//
// @param request - QueryEnableMultiAzPriceRequest
//
// @return QueryEnableMultiAzPriceResponse
func (client *Client) QueryEnableMultiAzPrice(request *QueryEnableMultiAzPriceRequest) (_result *QueryEnableMultiAzPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryEnableMultiAzPriceResponse{}
	_body, _err := client.QueryEnableMultiAzPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the latest minor version for the current major version.
//
// @param request - QueryMinorVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMinorVersionResponse
func (client *Client) QueryMinorVersionWithOptions(request *QueryMinorVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryMinorVersionResponse, _err error) {
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

	if !dara.IsNil(request.Version) {
		query["Version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMinorVersion"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/queryAppDefineVersion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMinorVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest minor version for the current major version.
//
// @param request - QueryMinorVersionRequest
//
// @return QueryMinorVersionResponse
func (client *Client) QueryMinorVersion(request *QueryMinorVersionRequest) (_result *QueryMinorVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryMinorVersionResponse{}
	_body, _err := client.QueryMinorVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price for changing the billing method of a StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// Note the following when you call this operation:
//
// - You can change the number of compute units (CUs) in a compute group only for Standard Edition instances. This feature is not supported for Starter Edition instances.
//
// - You can increase the disk size only for instances that have a standard compute group specification.
//
// - The instance must be in the Running state.
//
// After you change the number of CUs, the billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new number of CUs.
//
// - Subscription: The system calculates the supplementary fee based on the price difference between the old and new configurations and the remaining days in the billing cycle. The calculation starts from 00:00 on the following day.
//
// @param request - QueryModifyChargeTypePriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyChargeTypePriceResponse
func (client *Client) QueryModifyChargeTypePriceWithOptions(request *QueryModifyChargeTypePriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryModifyChargeTypePriceResponse, _err error) {
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
		Action:      dara.String("QueryModifyChargeTypePrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/buy/query_modify_charge_type_price"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifyChargeTypePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price for changing the billing method of a StarRocks instance.
//
// Description:
//
// Before you call this operation, make sure that you fully understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// Note the following when you call this operation:
//
// - You can change the number of compute units (CUs) in a compute group only for Standard Edition instances. This feature is not supported for Starter Edition instances.
//
// - You can increase the disk size only for instances that have a standard compute group specification.
//
// - The instance must be in the Running state.
//
// After you change the number of CUs, the billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new number of CUs.
//
// - Subscription: The system calculates the supplementary fee based on the price difference between the old and new configurations and the remaining days in the billing cycle. The calculation starts from 00:00 on the following day.
//
// @param request - QueryModifyChargeTypePriceRequest
//
// @return QueryModifyChargeTypePriceResponse
func (client *Client) QueryModifyChargeTypePrice(request *QueryModifyChargeTypePriceRequest) (_result *QueryModifyChargeTypePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryModifyChargeTypePriceResponse{}
	_body, _err := client.QueryModifyChargeTypePriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price for modifying the CU of compute group nodes.
//
// Description:
//
// Before you use this API, understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) for Serverless StarRocks.
//
// When you call this API, note the following:
//
// - Only standard instances allow you to modify the number of CUs for compute groups. Entry-level instances do not.
//
// - You can increase the disk size only for instances with a Standard Edition (standard) compute group.
//
// - The instance must be in the Running state.
//
// After you modify the number of CUs, billing changes as follows:
//
// - Pay-as-you-go: Billing is based on the new number of CUs.
//
// - Subscription: The amount to pay is calculated based on the price difference between the old and new configurations for the remainder of the subscription period. This period starts at 00:00 on the following day.
//
// @param request - QueryModifyCuPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyCuPriceResponse
func (client *Client) QueryModifyCuPriceWithOptions(request *QueryModifyCuPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryModifyCuPriceResponse, _err error) {
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
		Action:      dara.String("QueryModifyCuPrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/priceInquiry/modifyCu"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifyCuPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price for modifying the CU of compute group nodes.
//
// Description:
//
// Before you use this API, understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) for Serverless StarRocks.
//
// When you call this API, note the following:
//
// - Only standard instances allow you to modify the number of CUs for compute groups. Entry-level instances do not.
//
// - You can increase the disk size only for instances with a Standard Edition (standard) compute group.
//
// - The instance must be in the Running state.
//
// After you modify the number of CUs, billing changes as follows:
//
// - Pay-as-you-go: Billing is based on the new number of CUs.
//
// - Subscription: The amount to pay is calculated based on the price difference between the old and new configurations for the remainder of the subscription period. This period starts at 00:00 on the following day.
//
// @param request - QueryModifyCuPriceRequest
//
// @return QueryModifyCuPriceResponse
func (client *Client) QueryModifyCuPrice(request *QueryModifyCuPriceRequest) (_result *QueryModifyCuPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryModifyCuPriceResponse{}
	_body, _err := client.QueryModifyCuPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve the price for changing the number of disks in a compute group node
//
// Description:
//
// Before you call this operation, review the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) for Serverless StarRocks.
//
// - You can increase the disk count only for Standard Edition instances. You cannot increase the disk count for Starter Edition instances.
//
// - You can increase the disk count only for instances with a compute group specification type of Standard Edition (standard).
//
// - The instance must be in the Running state.
//
// After you increase the disk count, billing changes as follows:
//
// - Pay-as-you-go: Billing is based on the new disk type.
//
// - Subscription: During the billing cycle, the additional fee is calculated based on the price difference between the new and original configurations and the remaining days—from 00:00 the next day to the end of the validity period.
//
// @param request - QueryModifyDiskNumberPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyDiskNumberPriceResponse
func (client *Client) QueryModifyDiskNumberPriceWithOptions(request *QueryModifyDiskNumberPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryModifyDiskNumberPriceResponse, _err error) {
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
		Action:      dara.String("QueryModifyDiskNumberPrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/priceInquiry/modifyDiskNumber"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifyDiskNumberPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve the price for changing the number of disks in a compute group node
//
// Description:
//
// Before you call this operation, review the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) for Serverless StarRocks.
//
// - You can increase the disk count only for Standard Edition instances. You cannot increase the disk count for Starter Edition instances.
//
// - You can increase the disk count only for instances with a compute group specification type of Standard Edition (standard).
//
// - The instance must be in the Running state.
//
// After you increase the disk count, billing changes as follows:
//
// - Pay-as-you-go: Billing is based on the new disk type.
//
// - Subscription: During the billing cycle, the additional fee is calculated based on the price difference between the new and original configurations and the remaining days—from 00:00 the next day to the end of the validity period.
//
// @param request - QueryModifyDiskNumberPriceRequest
//
// @return QueryModifyDiskNumberPriceResponse
func (client *Client) QueryModifyDiskNumberPrice(request *QueryModifyDiskNumberPriceRequest) (_result *QueryModifyDiskNumberPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryModifyDiskNumberPriceResponse{}
	_body, _err := client.QueryModifyDiskNumberPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price of changing the disk performance level for a compute group node.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and <props="intl">[pricing](https://www.alibabacloud.com/product/ecs?_p_lc=1#pricing) of Serverless StarRocks.
//
// Note the following when you call this operation:
//
// - Only standard instances support disk upgrade or downgrade. Entry-level instances do not.
//
// - Only instances with a standard compute group specification support disk upgrade or downgrade.
//
// - The instance must be in the running state.
//
// - You cannot downgrade the disk to performance level (PL) 0.
//
// - The performance level of an Enhanced SSD (ESSD) is limited by its capacity. If you cannot upgrade the performance level, scale out the disk and try again. For more information, see <props="intl">[ESSDs](https://www.alibabacloud.com/help/en/ecs/user-guide/essds).
//
// After you upgrade or downgrade a disk, the billing changes as follows:
//
// - Pay-as-you-go: You are charged based on the new disk type.
//
// - Subscription: A supplementary fee is calculated based on the price difference between the new and old configurations and the remaining days in the subscription period, starting from 00:00 on the following day.
//
// @param request - QueryModifyDiskPerformanceLevelPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyDiskPerformanceLevelPriceResponse
func (client *Client) QueryModifyDiskPerformanceLevelPriceWithOptions(request *QueryModifyDiskPerformanceLevelPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryModifyDiskPerformanceLevelPriceResponse, _err error) {
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
		Action:      dara.String("QueryModifyDiskPerformanceLevelPrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/priceInquiry/modifyDiskPerformanceLevel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifyDiskPerformanceLevelPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price of changing the disk performance level for a compute group node.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and <props="intl">[pricing](https://www.alibabacloud.com/product/ecs?_p_lc=1#pricing) of Serverless StarRocks.
//
// Note the following when you call this operation:
//
// - Only standard instances support disk upgrade or downgrade. Entry-level instances do not.
//
// - Only instances with a standard compute group specification support disk upgrade or downgrade.
//
// - The instance must be in the running state.
//
// - You cannot downgrade the disk to performance level (PL) 0.
//
// - The performance level of an Enhanced SSD (ESSD) is limited by its capacity. If you cannot upgrade the performance level, scale out the disk and try again. For more information, see <props="intl">[ESSDs](https://www.alibabacloud.com/help/en/ecs/user-guide/essds).
//
// After you upgrade or downgrade a disk, the billing changes as follows:
//
// - Pay-as-you-go: You are charged based on the new disk type.
//
// - Subscription: A supplementary fee is calculated based on the price difference between the new and old configurations and the remaining days in the subscription period, starting from 00:00 on the following day.
//
// @param request - QueryModifyDiskPerformanceLevelPriceRequest
//
// @return QueryModifyDiskPerformanceLevelPriceResponse
func (client *Client) QueryModifyDiskPerformanceLevelPrice(request *QueryModifyDiskPerformanceLevelPriceRequest) (_result *QueryModifyDiskPerformanceLevelPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryModifyDiskPerformanceLevelPriceResponse{}
	_body, _err := client.QueryModifyDiskPerformanceLevelPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price for modifying the disk size of a single node in a compute group.
//
// Description:
//
// Before you call this operation, make sure you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks. When you call this operation, note the following:
//
// - Disk scale-out is supported only for standard instances. It is not supported for entry-level instances.
//
// - Disk scale-out is supported only for instances that use the standard compute group specification.
//
// - The instance must be in the Running state.
//
// After you scale out the disk, the billing changes are as follows:
//
// - Pay-as-you-go: You are billed for the new disk size.
//
// - Subscription: A supplementary fee is calculated based on the price difference between the old and new configurations and the remaining days in your subscription. The remaining days are calculated from 00:00 on the next day until the subscription expires.
//
// @param request - QueryModifyDiskSizePriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyDiskSizePriceResponse
func (client *Client) QueryModifyDiskSizePriceWithOptions(request *QueryModifyDiskSizePriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryModifyDiskSizePriceResponse, _err error) {
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
		Action:      dara.String("QueryModifyDiskSizePrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/priceInquiry/modifyDiskSize"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifyDiskSizePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price for modifying the disk size of a single node in a compute group.
//
// Description:
//
// Before you call this operation, make sure you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks. When you call this operation, note the following:
//
// - Disk scale-out is supported only for standard instances. It is not supported for entry-level instances.
//
// - Disk scale-out is supported only for instances that use the standard compute group specification.
//
// - The instance must be in the Running state.
//
// After you scale out the disk, the billing changes are as follows:
//
// - Pay-as-you-go: You are billed for the new disk size.
//
// - Subscription: A supplementary fee is calculated based on the price difference between the old and new configurations and the remaining days in your subscription. The remaining days are calculated from 00:00 on the next day until the subscription expires.
//
// @param request - QueryModifyDiskSizePriceRequest
//
// @return QueryModifyDiskSizePriceResponse
func (client *Client) QueryModifyDiskSizePrice(request *QueryModifyDiskSizePriceRequest) (_result *QueryModifyDiskSizePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryModifyDiskSizePriceResponse{}
	_body, _err := client.QueryModifyDiskSizePriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price for modifying the disk type of nodes in a compute group.
//
// Description:
//
// Before you call this API, make sure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of the Serverless StarRocks product. Take note of the following:
//
// - Only standard instances support disk scale-out. Entry-level instances do not.
//
// - Only instances that have a compute group of the Standard Edition (standard) specification type support disk scale-out.
//
// - Instances must be in the running (Running) state.
//
// After a disk scale-out, the billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new disk size.
//
// - Subscription: A supplemental fee is calculated based on the price difference between the new and old configurations for the remaining duration of the subscription. The remaining duration is calculated from 00:00 on the next day until the subscription expires.
//
// @param request - QueryModifyDiskTypePriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyDiskTypePriceResponse
func (client *Client) QueryModifyDiskTypePriceWithOptions(request *QueryModifyDiskTypePriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryModifyDiskTypePriceResponse, _err error) {
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
		Action:      dara.String("QueryModifyDiskTypePrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/priceInquiry/modifyDiskType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifyDiskTypePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price for modifying the disk type of nodes in a compute group.
//
// Description:
//
// Before you call this API, make sure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of the Serverless StarRocks product. Take note of the following:
//
// - Only standard instances support disk scale-out. Entry-level instances do not.
//
// - Only instances that have a compute group of the Standard Edition (standard) specification type support disk scale-out.
//
// - Instances must be in the running (Running) state.
//
// After a disk scale-out, the billing changes as follows:
//
// - Pay-as-you-go: You are billed based on the new disk size.
//
// - Subscription: A supplemental fee is calculated based on the price difference between the new and old configurations for the remaining duration of the subscription. The remaining duration is calculated from 00:00 on the next day until the subscription expires.
//
// @param request - QueryModifyDiskTypePriceRequest
//
// @return QueryModifyDiskTypePriceResponse
func (client *Client) QueryModifyDiskTypePrice(request *QueryModifyDiskTypePriceRequest) (_result *QueryModifyDiskTypePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryModifyDiskTypePriceResponse{}
	_body, _err := client.QueryModifyDiskTypePriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price for changing the number of nodes in a compute group.
//
// Description:
//
// Before you call this API, make sure you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// When you call this API, note the following:
//
// - You can change the number of nodes in a compute group only for standard instances. This operation is not supported for entry-level instances.
//
// - The instance must be in the Running state.
//
// - The number of frontend (FE) nodes cannot be an even number. You cannot scale in FE nodes.
//
// After you change the number of nodes, billing is affected as follows:
//
// - Pay-as-you-go: You are charged based on the new number of nodes.
//
// - Subscription: A supplemental fee is calculated based on the price difference between the old and new configurations and the number of remaining days in the subscription period. The remaining period is calculated from 00:00 of the next day until the subscription expires.
//
// @param request - QueryModifyNodeNumberPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifyNodeNumberPriceResponse
func (client *Client) QueryModifyNodeNumberPriceWithOptions(request *QueryModifyNodeNumberPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryModifyNodeNumberPriceResponse, _err error) {
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
		Action:      dara.String("QueryModifyNodeNumberPrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/priceInquiry/modifyNodeNumber"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifyNodeNumberPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price for changing the number of nodes in a compute group.
//
// Description:
//
// Before you call this API, make sure you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// When you call this API, note the following:
//
// - You can change the number of nodes in a compute group only for standard instances. This operation is not supported for entry-level instances.
//
// - The instance must be in the Running state.
//
// - The number of frontend (FE) nodes cannot be an even number. You cannot scale in FE nodes.
//
// After you change the number of nodes, billing is affected as follows:
//
// - Pay-as-you-go: You are charged based on the new number of nodes.
//
// - Subscription: A supplemental fee is calculated based on the price difference between the old and new configurations and the number of remaining days in the subscription period. The remaining period is calculated from 00:00 of the next day until the subscription expires.
//
// @param request - QueryModifyNodeNumberPriceRequest
//
// @return QueryModifyNodeNumberPriceResponse
func (client *Client) QueryModifyNodeNumberPrice(request *QueryModifyNodeNumberPriceRequest) (_result *QueryModifyNodeNumberPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryModifyNodeNumberPriceResponse{}
	_body, _err := client.QueryModifyNodeNumberPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the price to change the specification type of a compute group.
//
// Description:
//
// Before you call this operation, ensure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks. Note the following:
//
// - Only standard instances support disk scale-out. Basic instances do not.
//
// - Only instances with the Standard (standard) compute group specification type support disk scale-out.
//
// - The instance must be in the Running state.
//
// After a disk scale-out, the billing changes are as follows:
//
// - Pay-as-you-go: You are billed based on the new disk size.
//
// - Subscription: A supplementary fee is calculated based on the price difference between the old and new configurations and the number of remaining days in the subscription period (from 00:00 on the next day to the end of the validity period).
//
// @param request - QueryModifySpecTypePriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryModifySpecTypePriceResponse
func (client *Client) QueryModifySpecTypePriceWithOptions(request *QueryModifySpecTypePriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryModifySpecTypePriceResponse, _err error) {
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

	if !dara.IsNil(request.TargetSpecType) {
		query["TargetSpecType"] = request.TargetSpecType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryModifySpecTypePrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/priceInquiry/modifySpecType"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryModifySpecTypePriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the price to change the specification type of a compute group.
//
// Description:
//
// Before you call this operation, ensure that you understand the billing methods and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks. Note the following:
//
// - Only standard instances support disk scale-out. Basic instances do not.
//
// - Only instances with the Standard (standard) compute group specification type support disk scale-out.
//
// - The instance must be in the Running state.
//
// After a disk scale-out, the billing changes are as follows:
//
// - Pay-as-you-go: You are billed based on the new disk size.
//
// - Subscription: A supplementary fee is calculated based on the price difference between the old and new configurations and the number of remaining days in the subscription period (from 00:00 on the next day to the end of the validity period).
//
// @param request - QueryModifySpecTypePriceRequest
//
// @return QueryModifySpecTypePriceResponse
func (client *Client) QueryModifySpecTypePrice(request *QueryModifySpecTypePriceRequest) (_result *QueryModifySpecTypePriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryModifySpecTypePriceResponse{}
	_body, _err := client.QueryModifySpecTypePriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// An API for querying the price of new StarRocks purchases.
//
// Description:
//
// Before you use this API, review the billing model and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) for Serverless StarRocks.
//
// @param request - QueryPriceV1Request
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPriceV1Response
func (client *Client) QueryPriceV1WithOptions(request *QueryPriceV1Request, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryPriceV1Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentNodeGroup) {
		body["AgentNodeGroup"] = request.AgentNodeGroup
	}

	if !dara.IsNil(request.BackendNodeGroups) {
		body["BackendNodeGroups"] = request.BackendNodeGroups
	}

	if !dara.IsNil(request.Duration) {
		body["Duration"] = request.Duration
	}

	if !dara.IsNil(request.FrontendNodeGroups) {
		body["FrontendNodeGroups"] = request.FrontendNodeGroups
	}

	if !dara.IsNil(request.ObserverNodeGroups) {
		body["ObserverNodeGroups"] = request.ObserverNodeGroups
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

	if !dara.IsNil(request.PromotionOptionNo) {
		body["PromotionOptionNo"] = request.PromotionOptionNo
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RunMode) {
		body["RunMode"] = request.RunMode
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPriceV1"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/price/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPriceV1Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// An API for querying the price of new StarRocks purchases.
//
// Description:
//
// Before you use this API, review the billing model and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G)<props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) for Serverless StarRocks.
//
// @param request - QueryPriceV1Request
//
// @return QueryPriceV1Response
func (client *Client) QueryPriceV1(request *QueryPriceV1Request) (_result *QueryPriceV1Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPriceV1Response{}
	_body, _err := client.QueryPriceV1WithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Pricing information for unsubscribing from StarRocks subscription instances
//
// Description:
//
//	Danger:
//
// After you release an instance, Alibaba Cloud reclaims all physical resources used by the instance. All associated data is permanently lost and cannot be recovered.
//
// @param request - QueryRefundPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRefundPriceResponse
func (client *Client) QueryRefundPriceWithOptions(request *QueryRefundPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryRefundPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingInstanceIds) {
		query["billingInstanceIds"] = request.BillingInstanceIds
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRefundPrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/buy/queryRefundPrice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRefundPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Pricing information for unsubscribing from StarRocks subscription instances
//
// Description:
//
//	Danger:
//
// After you release an instance, Alibaba Cloud reclaims all physical resources used by the instance. All associated data is permanently lost and cannot be recovered.
//
// @param request - QueryRefundPriceRequest
//
// @return QueryRefundPriceResponse
func (client *Client) QueryRefundPrice(request *QueryRefundPriceRequest) (_result *QueryRefundPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRefundPriceResponse{}
	_body, _err := client.QueryRefundPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the renewal price for a StarRocks billing instance.
//
// @param request - QueryRenewPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRenewPriceResponse
func (client *Client) QueryRenewPriceWithOptions(request *QueryRenewPriceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryRenewPriceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("QueryRenewPrice"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/price/renew"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRenewPriceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the renewal price for a StarRocks billing instance.
//
// @param request - QueryRenewPriceRequest
//
// @return QueryRenewPriceResponse
func (client *Client) QueryRenewPrice(request *QueryRenewPriceRequest) (_result *QueryRenewPriceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryRenewPriceResponse{}
	_body, _err := client.QueryRenewPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the unpaid orders of a compute group or cluster.
//
// Description:
//
// Queries the unpaid orders of a compute group or cluster.
//
// @param request - QueryUnpaidOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUnpaidOrderResponse
func (client *Client) QueryUnpaidOrderWithOptions(request *QueryUnpaidOrderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryUnpaidOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BillingInstanceId) {
		query["BillingInstanceId"] = request.BillingInstanceId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUnpaidOrder"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/order/queryUnpaidOrder"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUnpaidOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the unpaid orders of a compute group or cluster.
//
// Description:
//
// Queries the unpaid orders of a compute group or cluster.
//
// @param request - QueryUnpaidOrderRequest
//
// @return QueryUnpaidOrderResponse
func (client *Client) QueryUnpaidOrder(request *QueryUnpaidOrderRequest) (_result *QueryUnpaidOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryUnpaidOrderResponse{}
	_body, _err := client.QueryUnpaidOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available upgrade versions for an EMR Serverless StarRocks instance. StarRocks instances have two version layers: a major version and a minor version. You can view both versions in the Version Information section on the Instance Details page in the EMR console. Use the Minor parameter to specify whether to query upgradable minor versions or major versions.
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
// Queries the available upgrade versions for an EMR Serverless StarRocks instance. StarRocks instances have two version layers: a major version and a minor version. You can view both versions in the Version Information section on the Instance Details page in the EMR console. Use the Minor parameter to specify whether to query upgradable minor versions or major versions.
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
// Reboot ECS instances in a cluster.
//
// Description:
//
// This operation is asynchronous. The instance does not reboot immediately. After a successful call, the instance enters the Rebooting state. The reboot is complete when the instance status changes to Running.
//
// @param request - RebootECSRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootECSResponse
func (client *Client) RebootECSWithOptions(request *RebootECSRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RebootECSResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventId) {
		query["EventId"] = request.EventId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RebootTime) {
		query["RebootTime"] = request.RebootTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootECS"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/event/rebootEcs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RebootECSResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reboot ECS instances in a cluster.
//
// Description:
//
// This operation is asynchronous. The instance does not reboot immediately. After a successful call, the instance enters the Rebooting state. The reboot is complete when the instance status changes to Running.
//
// @param request - RebootECSRequest
//
// @return RebootECSResponse
func (client *Client) RebootECS(request *RebootECSRequest) (_result *RebootECSResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RebootECSResponse{}
	_body, _err := client.RebootECSWithOptions(request, headers, runtime)
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
// Renews an instance.
//
// Description:
//
// Before you call this API, make sure that you fully understand the billing model and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// When you call this API, note that only subscription instances can be renewed.
//
// @param request - RenewInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("RenewInstance"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/order/renew_instance"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews an instance.
//
// Description:
//
// Before you call this API, make sure that you fully understand the billing model and <props="china">[billable items](https://help.aliyun.com/zh/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c4g.11186623.help-menu-28066.d_0_1_1.4db82b05p3Gg0G) <props="intl">[billable items](https://www.alibabacloud.com/help/en/emr/emr-serverless-starrocks/product-overview/billable-items?spm=a2c63.p38356.help-menu-28066.d_0_1_0.3aaf4b0b69jN1P) of Serverless StarRocks.
//
// When you call this API, note that only subscription instances can be renewed.
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, headers, runtime)
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
// Restarts a specified node group.
//
// Description:
//
// This operation is asynchronous. The instance does not restart immediately. After you call the operation successfully, the instance enters the restarting state. When the instance status changes to running, the restart is complete.
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
// Restarts a specified node group.
//
// Description:
//
// This operation is asynchronous. The instance does not restart immediately. After you call the operation successfully, the instance enters the restarting state. When the instance status changes to running, the restart is complete.
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
// Restart nodes in a cluster.
//
// Description:
//
// This API is asynchronous. The instance restart is not immediate. After a successful call, the instance enters the restarting state. The restart is complete when the returned status is running.
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
// Restart nodes in a cluster.
//
// Description:
//
// This API is asynchronous. The instance restart is not immediate. After a successful call, the instance enters the restarting state. The restart is complete when the returned status is running.
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
// # Restore an instance from a backup
//
// Description:
//
// This is an asynchronous API. The instance does not restart immediately. After you call this API successfully, the instance enters the restarting state. When the instance status changes to running, the restore is complete.
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
// # Restore an instance from a backup
//
// Description:
//
// This is an asynchronous API. The instance does not restart immediately. After you call this API successfully, the instance enters the restarting state. When the instance status changes to running, the restore is complete.
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
// Resumes an instance automatically stopped by openlake.
//
// Description:
//
//	Danger:
//
// When you release an instance, its physical resources are reclaimed. All data on the instance is lost and cannot be recovered.
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
// Resumes an instance automatically stopped by openlake.
//
// Description:
//
//	Danger:
//
// When you release an instance, its physical resources are reclaimed. All data on the instance is lost and cannot be recovered.
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
// Rolls back a configuration modification that is currently in progress.
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
// Rolls back a configuration modification that is currently in progress.
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
// Switches the active and standby zones.
//
// Description:
//
// Switches the active and standby zones.
//
// @param request - SwitchActiveStandbyZonesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchActiveStandbyZonesResponse
func (client *Client) SwitchActiveStandbyZonesWithOptions(request *SwitchActiveStandbyZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SwitchActiveStandbyZonesResponse, _err error) {
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

	if !dara.IsNil(request.TargetZoneId) {
		query["TargetZoneId"] = request.TargetZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SwitchActiveStandbyZones"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/recovery/switchZones"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SwitchActiveStandbyZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Switches the active and standby zones.
//
// Description:
//
// Switches the active and standby zones.
//
// @param request - SwitchActiveStandbyZonesRequest
//
// @return SwitchActiveStandbyZonesResponse
func (client *Client) SwitchActiveStandbyZones(request *SwitchActiveStandbyZonesRequest) (_result *SwitchActiveStandbyZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SwitchActiveStandbyZonesResponse{}
	_body, _err := client.SwitchActiveStandbyZonesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches tags to specified resources.
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
// Attaches tags to specified resources.
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
// Enables or disables automatic minor version upgrades for a StarRocks instance.
//
// @param request - ToggleAutoMinorVersionUpgradeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ToggleAutoMinorVersionUpgradeResponse
func (client *Client) ToggleAutoMinorVersionUpgradeWithOptions(request *ToggleAutoMinorVersionUpgradeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ToggleAutoMinorVersionUpgradeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoUpgrade) {
		query["AutoUpgrade"] = request.AutoUpgrade
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ToggleAutoMinorVersionUpgrade"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/starrocks/toggleAutoMinorVersionUpgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ToggleAutoMinorVersionUpgradeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables automatic minor version upgrades for a StarRocks instance.
//
// @param request - ToggleAutoMinorVersionUpgradeRequest
//
// @return ToggleAutoMinorVersionUpgradeResponse
func (client *Client) ToggleAutoMinorVersionUpgrade(request *ToggleAutoMinorVersionUpgradeRequest) (_result *ToggleAutoMinorVersionUpgradeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ToggleAutoMinorVersionUpgradeResponse{}
	_body, _err := client.ToggleAutoMinorVersionUpgradeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Switches an Internet-facing SLB instance on or off.
//
// Description:
//
//	Danger:
//
// After you release an instance, Alibaba Cloud reclaims all physical resources used by the instance. All associated data is permanently lost and cannot be recovered.
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
// Switches an Internet-facing SLB instance on or off.
//
// Description:
//
//	Danger:
//
// After you release an instance, Alibaba Cloud reclaims all physical resources used by the instance. All associated data is permanently lost and cannot be recovered.
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
// Detaches tags from multiple resources.
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
// Detaches tags from multiple resources.
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
// Updates the description of a backup job.
//
// @param request - UpdateBackupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupResponse
func (client *Client) UpdateBackupWithOptions(request *UpdateBackupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBackupResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.BackupTaskId) {
		body["backupTaskId"] = request.BackupTaskId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBackup"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/backup/manage/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBackupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a backup job.
//
// @param request - UpdateBackupRequest
//
// @return UpdateBackupResponse
func (client *Client) UpdateBackup(request *UpdateBackupRequest) (_result *UpdateBackupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBackupResponse{}
	_body, _err := client.UpdateBackupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a backup policy.
//
// @param request - UpdateBackupPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBackupPolicyResponse
func (client *Client) UpdateBackupPolicyWithOptions(request *UpdateBackupPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateBackupPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireDays) {
		body["ExpireDays"] = request.ExpireDays
	}

	if !dara.IsNil(request.Hour) {
		body["Hour"] = request.Hour
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Minute) {
		body["Minute"] = request.Minute
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RecurrenceValues) {
		body["RecurrenceValues"] = request.RecurrenceValues
	}

	if !dara.IsNil(request.TimeoutSeconds) {
		body["TimeoutSeconds"] = request.TimeoutSeconds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBackupPolicy"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/backupRestore/policy/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBackupPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a backup policy.
//
// @param request - UpdateBackupPolicyRequest
//
// @return UpdateBackupPolicyResponse
func (client *Client) UpdateBackupPolicy(request *UpdateBackupPolicyRequest) (_result *UpdateBackupPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBackupPolicyResponse{}
	_body, _err := client.UpdateBackupPolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the number of gateway nodes.
//
// Description:
//
// Updates the number of gateway nodes.
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
// Updates the number of gateway nodes.
//
// Description:
//
// Updates the number of gateway nodes.
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
// Updates the CIDR blocks in a whitelist group.
//
// @param request - UpdateInnerIpWhitelistGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInnerIpWhitelistGroupResponse
func (client *Client) UpdateInnerIpWhitelistGroupWithOptions(request *UpdateInnerIpWhitelistGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInnerIpWhitelistGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CidrIpList) {
		body["CidrIpList"] = request.CidrIpList
	}

	if !dara.IsNil(request.InnerIpWhitelistGroupId) {
		body["InnerIpWhitelistGroupId"] = request.InnerIpWhitelistGroupId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInnerIpWhitelistGroup"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/securityGroup/update"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInnerIpWhitelistGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the CIDR blocks in a whitelist group.
//
// @param request - UpdateInnerIpWhitelistGroupRequest
//
// @return UpdateInnerIpWhitelistGroupResponse
func (client *Client) UpdateInnerIpWhitelistGroup(request *UpdateInnerIpWhitelistGroupRequest) (_result *UpdateInnerIpWhitelistGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInnerIpWhitelistGroupResponse{}
	_body, _err := client.UpdateInnerIpWhitelistGroupWithOptions(request, headers, runtime)
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
// Updates the description of a compute group.
//
// @param request - UpdateNodeGroupDescriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeGroupDescriptionResponse
func (client *Client) UpdateNodeGroupDescriptionWithOptions(request *UpdateNodeGroupDescriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateNodeGroupDescriptionResponse, _err error) {
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

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.XAcsRamAuthContext) {
		query["X-Acs-Ram-Auth-Context"] = request.XAcsRamAuthContext
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNodeGroupDescription"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/nodegroup/updateDescription"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodeGroupDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the description of a compute group.
//
// @param request - UpdateNodeGroupDescriptionRequest
//
// @return UpdateNodeGroupDescriptionResponse
func (client *Client) UpdateNodeGroupDescription(request *UpdateNodeGroupDescriptionRequest) (_result *UpdateNodeGroupDescriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateNodeGroupDescriptionResponse{}
	_body, _err := client.UpdateNodeGroupDescriptionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enable or disable the Internet-facing SLB for the FE or BE component.
//
// @param request - UpdatePublicNetworkStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePublicNetworkStatusResponse
func (client *Client) UpdatePublicNetworkStatusWithOptions(request *UpdatePublicNetworkStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePublicNetworkStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComponentType) {
		query["ComponentType"] = request.ComponentType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeGroupId) {
		query["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.PublicNetworkEnabled) {
		query["PublicNetworkEnabled"] = request.PublicNetworkEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePublicNetworkStatus"),
		Version:     dara.String("2022-10-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/webapi/network/updatePublicNetworkStatus"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePublicNetworkStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enable or disable the Internet-facing SLB for the FE or BE component.
//
// @param request - UpdatePublicNetworkStatusRequest
//
// @return UpdatePublicNetworkStatusResponse
func (client *Client) UpdatePublicNetworkStatus(request *UpdatePublicNetworkStatusRequest) (_result *UpdatePublicNetworkStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePublicNetworkStatusResponse{}
	_body, _err := client.UpdatePublicNetworkStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the version of a Serverless StarRocks instance. Serverless StarRocks has two levels of version definition: the major version displayed in the "Version" field on the cluster details page, and the minor version displayed in the "Minor Version" field on the cluster details page. This operation can be used to upgrade either the minor version or the major version. You can call the QueryUpgradableVersions operation to query the versions to which a cluster can be upgraded.
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
// Upgrades the version of a Serverless StarRocks instance. Serverless StarRocks has two levels of version definition: the major version displayed in the "Version" field on the cluster details page, and the minor version displayed in the "Minor Version" field on the cluster details page. This operation can be used to upgrade either the minor version or the major version. You can call the QueryUpgradableVersions operation to query the versions to which a cluster can be upgraded.
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
