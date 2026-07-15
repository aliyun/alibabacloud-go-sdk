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
		"us-west-1":      dara.String("apig.us-west-1.aliyuncs.com"),
		"us-east-1":      dara.String("apig.us-east-1.aliyuncs.com"),
		"me-east-1":      dara.String("apig.me-east-1.aliyuncs.com"),
		"me-central-1":   dara.String("apig.me-central-1.aliyuncs.com"),
		"eu-west-1":      dara.String("apig.eu-west-1.aliyuncs.com"),
		"eu-central-1":   dara.String("apig.eu-central-1.aliyuncs.com"),
		"cn-zhangjiakou": dara.String("apig.cn-zhangjiakou.aliyuncs.com"),
		"cn-wulanchabu":  dara.String("apig.cn-wulanchabu.aliyuncs.com"),
		"cn-shenzhen":    dara.String("apig.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai":    dara.String("apig.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":     dara.String("apig.cn-qingdao.aliyuncs.com"),
		"cn-hongkong":    dara.String("apig.cn-hongkong.aliyuncs.com"),
		"cn-heyuan":      dara.String("apig.cn-heyuan.aliyuncs.com"),
		"cn-hangzhou":    dara.String("apig.cn-hangzhou.aliyuncs.com"),
		"cn-guangzhou":   dara.String("apig.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":     dara.String("apig.cn-chengdu.aliyuncs.com"),
		"cn-beijing":     dara.String("apig.cn-beijing.aliyuncs.com"),
		"ap-southeast-7": dara.String("apig.ap-southeast-7.aliyuncs.com"),
		"ap-southeast-6": dara.String("apig.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-5": dara.String("apig.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-3": dara.String("apig.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-2": dara.String("apig.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-1": dara.String("apig.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2": dara.String("apig.ap-northeast-2.aliyuncs.com"),
		"ap-northeast-1": dara.String("apig.ap-northeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("apig"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a gateway quota throttling rule.
//
// Description:
//
// Creates a consumer-based quota rule for an AI gateway. This operation takes effect only on AI gateways of version 2.1.19 or later.
//
// >
//
// >  Recommended call sequence:
//
// > - Step 1: Perform a dry run to check for rule conflicts.
//
// > - - Set dryRun to true.
//
// > - - The response contains a conflict preview with a conflictHash value.
//
// > - Step 2: Submit the request after confirmation.
//
// > - - No conflicts: Set dryRun to false and overwrite to false.
//
// > - - Conflicts exist and you confirm the overwrite: Set dryRun to false, overwrite to true, and conflictHash to the value returned in the previous step.
//
// @param request - AddGatewayQuotaRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewayQuotaRuleResponse
func (client *Client) AddGatewayQuotaRuleWithOptions(gatewayId *string, request *AddGatewayQuotaRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddGatewayQuotaRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConflictHash) {
		body["conflictHash"] = request.ConflictHash
	}

	if !dara.IsNil(request.ConsumerGroupIds) {
		body["consumerGroupIds"] = request.ConsumerGroupIds
	}

	if !dara.IsNil(request.ConsumerIds) {
		body["consumerIds"] = request.ConsumerIds
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Overwrite) {
		body["overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.PeriodMultiplier) {
		body["periodMultiplier"] = request.PeriodMultiplier
	}

	if !dara.IsNil(request.PeriodType) {
		body["periodType"] = request.PeriodType
	}

	if !dara.IsNil(request.QuotaDimension) {
		body["quotaDimension"] = request.QuotaDimension
	}

	if !dara.IsNil(request.QuotaLimit) {
		body["quotaLimit"] = request.QuotaLimit
	}

	if !dara.IsNil(request.RuleName) {
		body["ruleName"] = request.RuleName
	}

	if !dara.IsNil(request.Timezone) {
		body["timezone"] = request.Timezone
	}

	if !dara.IsNil(request.WindowAlignment) {
		body["windowAlignment"] = request.WindowAlignment
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewayQuotaRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/quota-rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewayQuotaRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a gateway quota throttling rule.
//
// Description:
//
// Creates a consumer-based quota rule for an AI gateway. This operation takes effect only on AI gateways of version 2.1.19 or later.
//
// >
//
// >  Recommended call sequence:
//
// > - Step 1: Perform a dry run to check for rule conflicts.
//
// > - - Set dryRun to true.
//
// > - - The response contains a conflict preview with a conflictHash value.
//
// > - Step 2: Submit the request after confirmation.
//
// > - - No conflicts: Set dryRun to false and overwrite to false.
//
// > - - Conflicts exist and you confirm the overwrite: Set dryRun to false, overwrite to true, and conflictHash to the value returned in the previous step.
//
// @param request - AddGatewayQuotaRuleRequest
//
// @return AddGatewayQuotaRuleResponse
func (client *Client) AddGatewayQuotaRule(gatewayId *string, request *AddGatewayQuotaRuleRequest) (_result *AddGatewayQuotaRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddGatewayQuotaRuleResponse{}
	_body, _err := client.AddGatewayQuotaRuleWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Authorizes a security group that allows a gateway to access services.
//
// @param request - AddGatewaySecurityGroupRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGatewaySecurityGroupRuleResponse
func (client *Client) AddGatewaySecurityGroupRuleWithOptions(gatewayId *string, request *AddGatewaySecurityGroupRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddGatewaySecurityGroupRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.PortRanges) {
		body["portRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["securityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddGatewaySecurityGroupRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/security-group-rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddGatewaySecurityGroupRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Authorizes a security group that allows a gateway to access services.
//
// @param request - AddGatewaySecurityGroupRuleRequest
//
// @return AddGatewaySecurityGroupRuleResponse
func (client *Client) AddGatewaySecurityGroupRule(gatewayId *string, request *AddGatewaySecurityGroupRuleRequest) (_result *AddGatewaySecurityGroupRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddGatewaySecurityGroupRuleResponse{}
	_body, _err := client.AddGatewaySecurityGroupRuleWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量添加消费者组成员
//
// @param request - BatchAddConsumerGroupConsumersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddConsumerGroupConsumersResponse
func (client *Client) BatchAddConsumerGroupConsumersWithOptions(consumerGroupId *string, request *BatchAddConsumerGroupConsumersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchAddConsumerGroupConsumersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerIds) {
		body["consumerIds"] = request.ConsumerIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAddConsumerGroupConsumers"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumer-groups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/consumers/batch-add"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAddConsumerGroupConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量添加消费者组成员
//
// @param request - BatchAddConsumerGroupConsumersRequest
//
// @return BatchAddConsumerGroupConsumersResponse
func (client *Client) BatchAddConsumerGroupConsumers(consumerGroupId *string, request *BatchAddConsumerGroupConsumersRequest) (_result *BatchAddConsumerGroupConsumersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchAddConsumerGroupConsumersResponse{}
	_body, _err := client.BatchAddConsumerGroupConsumersWithOptions(consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes consumer authorization rules in batches.
//
// @param request - BatchDeleteConsumerAuthorizationRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteConsumerAuthorizationRuleResponse
func (client *Client) BatchDeleteConsumerAuthorizationRuleWithOptions(request *BatchDeleteConsumerAuthorizationRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchDeleteConsumerAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerAuthorizationRuleIds) {
		query["consumerAuthorizationRuleIds"] = request.ConsumerAuthorizationRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/authorization-rules"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes consumer authorization rules in batches.
//
// @param request - BatchDeleteConsumerAuthorizationRuleRequest
//
// @return BatchDeleteConsumerAuthorizationRuleResponse
func (client *Client) BatchDeleteConsumerAuthorizationRule(request *BatchDeleteConsumerAuthorizationRuleRequest) (_result *BatchDeleteConsumerAuthorizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchDeleteConsumerAuthorizationRuleResponse{}
	_body, _err := client.BatchDeleteConsumerAuthorizationRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量移除消费者组成员
//
// @param request - BatchRemoveConsumerGroupConsumersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchRemoveConsumerGroupConsumersResponse
func (client *Client) BatchRemoveConsumerGroupConsumersWithOptions(consumerGroupId *string, request *BatchRemoveConsumerGroupConsumersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchRemoveConsumerGroupConsumersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerIds) {
		body["consumerIds"] = request.ConsumerIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchRemoveConsumerGroupConsumers"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumer-groups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/consumers/batch-remove"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchRemoveConsumerGroupConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量移除消费者组成员
//
// @param request - BatchRemoveConsumerGroupConsumersRequest
//
// @return BatchRemoveConsumerGroupConsumersResponse
func (client *Client) BatchRemoveConsumerGroupConsumers(consumerGroupId *string, request *BatchRemoveConsumerGroupConsumersRequest) (_result *BatchRemoveConsumerGroupConsumersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchRemoveConsumerGroupConsumersResponse{}
	_body, _err := client.BatchRemoveConsumerGroupConsumersWithOptions(consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a resource to a different resource group.
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
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/move-resource-group"),
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
// Moves a resource to a different resource group.
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
// 创建AI模型卡片
//
// Description:
//
// 在指定AI网关实例的已有模型供应商下创建模型卡片。目标网关必须存在、属于当前账号且类型为AI网关，modelProvider必须引用该网关中已存在的模型供应商。
//
// 同一AI网关实例、同一模型供应商下的modelName必须唯一；单个网关实例最多可创建1000张模型卡片。credit当前仅支持fixed类型，费用单位为Credits/百万Token；未传时type默认为fixed，各项费用默认为0。availablePaths中的每一项必须同时包含path和type。
//
// @param request - CreateAiModelCardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAiModelCardResponse
func (client *Client) CreateAiModelCardWithOptions(request *CreateAiModelCardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAiModelCardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AvailablePaths) {
		body["availablePaths"] = request.AvailablePaths
	}

	if !dara.IsNil(request.Credit) {
		body["credit"] = request.Credit
	}

	if !dara.IsNil(request.Features) {
		body["features"] = request.Features
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Meta) {
		body["meta"] = request.Meta
	}

	if !dara.IsNil(request.ModelName) {
		body["modelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelProvider) {
		body["modelProvider"] = request.ModelProvider
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAiModelCard"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-cards"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAiModelCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AI模型卡片
//
// Description:
//
// 在指定AI网关实例的已有模型供应商下创建模型卡片。目标网关必须存在、属于当前账号且类型为AI网关，modelProvider必须引用该网关中已存在的模型供应商。
//
// 同一AI网关实例、同一模型供应商下的modelName必须唯一；单个网关实例最多可创建1000张模型卡片。credit当前仅支持fixed类型，费用单位为Credits/百万Token；未传时type默认为fixed，各项费用默认为0。availablePaths中的每一项必须同时包含path和type。
//
// @param request - CreateAiModelCardRequest
//
// @return CreateAiModelCardResponse
func (client *Client) CreateAiModelCard(request *CreateAiModelCardRequest) (_result *CreateAiModelCardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAiModelCardResponse{}
	_body, _err := client.CreateAiModelCardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建AI模型供应商
//
// @param request - CreateAiModelProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAiModelProviderResponse
func (client *Client) CreateAiModelProviderWithOptions(request *CreateAiModelProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAiModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Provider) {
		body["provider"] = request.Provider
	}

	if !dara.IsNil(request.ServiceIds) {
		body["serviceIds"] = request.ServiceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAiModelProvider"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-providers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAiModelProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AI模型供应商
//
// @param request - CreateAiModelProviderRequest
//
// @return CreateAiModelProviderResponse
func (client *Client) CreateAiModelProvider(request *CreateAiModelProviderRequest) (_result *CreateAiModelProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAiModelProviderResponse{}
	_body, _err := client.CreateAiModelProviderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and attaches a policy.
//
// @param request - CreateAndAttachPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndAttachPolicyResponse
func (client *Client) CreateAndAttachPolicyWithOptions(request *CreateAndAttachPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAndAttachPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceIds) {
		body["attachResourceIds"] = request.AttachResourceIds
	}

	if !dara.IsNil(request.AttachResourceType) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.ClassName) {
		body["className"] = request.ClassName
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAndAttachPolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndAttachPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and attaches a policy.
//
// @param request - CreateAndAttachPolicyRequest
//
// @return CreateAndAttachPolicyResponse
func (client *Client) CreateAndAttachPolicy(request *CreateAndAttachPolicyRequest) (_result *CreateAndAttachPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAndAttachPolicyResponse{}
	_body, _err := client.CreateAndAttachPolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a consumer.
//
// @param request - CreateConsumerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerResponse
func (client *Client) CreateConsumerWithOptions(request *CreateConsumerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AkSkIdentityConfigs) {
		body["akSkIdentityConfigs"] = request.AkSkIdentityConfigs
	}

	if !dara.IsNil(request.ApikeyIdentityConfig) {
		body["apikeyIdentityConfig"] = request.ApikeyIdentityConfig
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.GatewayType) {
		body["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.JwtIdentityConfig) {
		body["jwtIdentityConfig"] = request.JwtIdentityConfig
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a consumer.
//
// @param request - CreateConsumerRequest
//
// @return CreateConsumerResponse
func (client *Client) CreateConsumer(request *CreateConsumerRequest) (_result *CreateConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsumerResponse{}
	_body, _err := client.CreateConsumerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a consumer authorization rule.
//
// @param request - CreateConsumerAuthorizationRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerAuthorizationRuleResponse
func (client *Client) CreateConsumerAuthorizationRuleWithOptions(consumerId *string, request *CreateConsumerAuthorizationRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceInfos) {
		body["authorizationResourceInfos"] = request.AuthorizationResourceInfos
	}

	if !dara.IsNil(request.ExpireMode) {
		body["expireMode"] = request.ExpireMode
	}

	if !dara.IsNil(request.ExpireTimestamp) {
		body["expireTimestamp"] = request.ExpireTimestamp
	}

	if !dara.IsNil(request.ParentResourceType) {
		body["parentResourceType"] = request.ParentResourceType
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a consumer authorization rule.
//
// @param request - CreateConsumerAuthorizationRuleRequest
//
// @return CreateConsumerAuthorizationRuleResponse
func (client *Client) CreateConsumerAuthorizationRule(consumerId *string, request *CreateConsumerAuthorizationRuleRequest) (_result *CreateConsumerAuthorizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsumerAuthorizationRuleResponse{}
	_body, _err := client.CreateConsumerAuthorizationRuleWithOptions(consumerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates consumer authorization rules.
//
// @param request - CreateConsumerAuthorizationRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerAuthorizationRulesResponse
func (client *Client) CreateConsumerAuthorizationRulesWithOptions(request *CreateConsumerAuthorizationRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerAuthorizationRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationRules) {
		body["authorizationRules"] = request.AuthorizationRules
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerAuthorizationRules"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/authorization-rules"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerAuthorizationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates consumer authorization rules.
//
// @param request - CreateConsumerAuthorizationRulesRequest
//
// @return CreateConsumerAuthorizationRulesResponse
func (client *Client) CreateConsumerAuthorizationRules(request *CreateConsumerAuthorizationRulesRequest) (_result *CreateConsumerAuthorizationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsumerAuthorizationRulesResponse{}
	_body, _err := client.CreateConsumerAuthorizationRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建消费者组
//
// @param request - CreateConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroupWithOptions(request *CreateConsumerGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerGroupId) {
		body["consumerGroupId"] = request.ConsumerGroupId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GatewayType) {
		body["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerGroup"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumer-groups"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建消费者组
//
// @param request - CreateConsumerGroupRequest
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a domain name.
//
// @param request - CreateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CaCertIdentifier) {
		body["caCertIdentifier"] = request.CaCertIdentifier
	}

	if !dara.IsNil(request.CertIdentifier) {
		body["certIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.ClientCACert) {
		body["clientCACert"] = request.ClientCACert
	}

	if !dara.IsNil(request.DomainScope) {
		body["domainScope"] = request.DomainScope
	}

	if !dara.IsNil(request.ForceHttps) {
		body["forceHttps"] = request.ForceHttps
	}

	if !dara.IsNil(request.GatewayType) {
		body["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.Http2Option) {
		body["http2Option"] = request.Http2Option
	}

	if !dara.IsNil(request.MTLSEnabled) {
		body["mTLSEnabled"] = request.MTLSEnabled
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TlsCipherSuitesConfig) {
		body["tlsCipherSuitesConfig"] = request.TlsCipherSuitesConfig
	}

	if !dara.IsNil(request.TlsMax) {
		body["tlsMax"] = request.TlsMax
	}

	if !dara.IsNil(request.TlsMin) {
		body["tlsMin"] = request.TlsMin
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a domain name.
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateEnvironment is deprecated
//
// Summary:
//
// Creates an environment.
//
// @param request - CreateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnvironmentResponse
func (client *Client) CreateEnvironmentWithOptions(request *CreateEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEnvironmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["alias"] = request.Alias
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnvironment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateEnvironment is deprecated
//
// Summary:
//
// Creates an environment.
//
// @param request - CreateEnvironmentRequest
//
// @return CreateEnvironmentResponse
// Deprecated
func (client *Client) CreateEnvironment(request *CreateEnvironmentRequest) (_result *CreateEnvironmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEnvironmentResponse{}
	_body, _err := client.CreateEnvironmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a cloud-native gateway.
//
// @param request - CreateGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGatewayResponse
func (client *Client) CreateGatewayWithOptions(request *CreateGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		body["chargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.GatewayEdition) {
		body["gatewayEdition"] = request.GatewayEdition
	}

	if !dara.IsNil(request.GatewayType) {
		body["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.LogConfig) {
		body["logConfig"] = request.LogConfig
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.NetworkAccessConfig) {
		body["networkAccessConfig"] = request.NetworkAccessConfig
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Spec) {
		body["spec"] = request.Spec
	}

	if !dara.IsNil(request.Tag) {
		body["tag"] = request.Tag
	}

	if !dara.IsNil(request.VpcId) {
		body["vpcId"] = request.VpcId
	}

	if !dara.IsNil(request.ZoneConfig) {
		body["zoneConfig"] = request.ZoneConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a cloud-native gateway.
//
// @param request - CreateGatewayRequest
//
// @return CreateGatewayResponse
func (client *Client) CreateGateway(request *CreateGatewayRequest) (_result *CreateGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGatewayResponse{}
	_body, _err := client.CreateGatewayWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an HTTP API.
//
// @param request - CreateHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiResponse
func (client *Client) CreateHttpApiWithOptions(request *CreateHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHttpApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProtocols) {
		body["agentProtocols"] = request.AgentProtocols
	}

	if !dara.IsNil(request.AiProtocols) {
		body["aiProtocols"] = request.AiProtocols
	}

	if !dara.IsNil(request.AuthConfig) {
		body["authConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.BasePath) {
		body["basePath"] = request.BasePath
	}

	if !dara.IsNil(request.BelongGatewayId) {
		body["belongGatewayId"] = request.BelongGatewayId
	}

	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.EnableAuth) {
		body["enableAuth"] = request.EnableAuth
	}

	if !dara.IsNil(request.FirstByteTimeout) {
		body["firstByteTimeout"] = request.FirstByteTimeout
	}

	if !dara.IsNil(request.IngressConfig) {
		body["ingressConfig"] = request.IngressConfig
	}

	if !dara.IsNil(request.ModelCategory) {
		body["modelCategory"] = request.ModelCategory
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Protocols) {
		body["protocols"] = request.Protocols
	}

	if !dara.IsNil(request.RemoveBasePathOnForward) {
		body["removeBasePathOnForward"] = request.RemoveBasePathOnForward
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.VersionConfig) {
		body["versionConfig"] = request.VersionConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an HTTP API.
//
// @param request - CreateHttpApiRequest
//
// @return CreateHttpApiResponse
func (client *Client) CreateHttpApi(request *CreateHttpApiRequest) (_result *CreateHttpApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHttpApiResponse{}
	_body, _err := client.CreateHttpApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates operations for an HTTP API.
//
// @param request - CreateHttpApiOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiOperationResponse
func (client *Client) CreateHttpApiOperationWithOptions(httpApiId *string, request *CreateHttpApiOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHttpApiOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Operations) {
		body["operations"] = request.Operations
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpApiOperation"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpApiOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates operations for an HTTP API.
//
// @param request - CreateHttpApiOperationRequest
//
// @return CreateHttpApiOperationResponse
func (client *Client) CreateHttpApiOperation(httpApiId *string, request *CreateHttpApiOperationRequest) (_result *CreateHttpApiOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHttpApiOperationResponse{}
	_body, _err := client.CreateHttpApiOperationWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a route for an HTTP API.
//
// @param request - CreateHttpApiRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHttpApiRouteResponse
func (client *Client) CreateHttpApiRouteWithOptions(httpApiId *string, request *CreateHttpApiRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateHttpApiRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BackendConfig) {
		body["backendConfig"] = request.BackendConfig
	}

	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DomainIds) {
		body["domainIds"] = request.DomainIds
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Match) {
		body["match"] = request.Match
	}

	if !dara.IsNil(request.McpRouteConfig) {
		body["mcpRouteConfig"] = request.McpRouteConfig
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.PolicyConfigs) {
		body["policyConfigs"] = request.PolicyConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHttpApiRoute"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHttpApiRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a route for an HTTP API.
//
// @param request - CreateHttpApiRouteRequest
//
// @return CreateHttpApiRouteResponse
func (client *Client) CreateHttpApiRoute(httpApiId *string, request *CreateHttpApiRouteRequest) (_result *CreateHttpApiRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHttpApiRouteResponse{}
	_body, _err := client.CreateHttpApiRouteWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Model Context Protocol (MCP) server.
//
// @param request - CreateMcpServerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMcpServerResponse
func (client *Client) CreateMcpServerWithOptions(request *CreateMcpServerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssembledSources) {
		body["assembledSources"] = request.AssembledSources
	}

	if !dara.IsNil(request.BackendConfig) {
		body["backendConfig"] = request.BackendConfig
	}

	if !dara.IsNil(request.CreateFromType) {
		body["createFromType"] = request.CreateFromType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DomainIds) {
		body["domainIds"] = request.DomainIds
	}

	if !dara.IsNil(request.ExposedUriPath) {
		body["exposedUriPath"] = request.ExposedUriPath
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GrayMcpServerConfigs) {
		body["grayMcpServerConfigs"] = request.GrayMcpServerConfigs
	}

	if !dara.IsNil(request.Match) {
		body["match"] = request.Match
	}

	if !dara.IsNil(request.McpServerConfig) {
		body["mcpServerConfig"] = request.McpServerConfig
	}

	if !dara.IsNil(request.McpStatisticsEnable) {
		body["mcpStatisticsEnable"] = request.McpStatisticsEnable
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMcpServer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/mcp-servers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Model Context Protocol (MCP) server.
//
// @param request - CreateMcpServerRequest
//
// @return CreateMcpServerResponse
func (client *Client) CreateMcpServer(request *CreateMcpServerRequest) (_result *CreateMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMcpServerResponse{}
	_body, _err := client.CreateMcpServerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Mounts a plug-in.
//
// @param request - CreatePluginAttachmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePluginAttachmentResponse
func (client *Client) CreatePluginAttachmentWithOptions(request *CreatePluginAttachmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePluginAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceIds) {
		body["attachResourceIds"] = request.AttachResourceIds
	}

	if !dara.IsNil(request.AttachResourceType) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PluginConfig) {
		body["pluginConfig"] = request.PluginConfig
	}

	if !dara.IsNil(request.PluginId) {
		body["pluginId"] = request.PluginId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePluginAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePluginAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Mounts a plug-in.
//
// @param request - CreatePluginAttachmentRequest
//
// @return CreatePluginAttachmentResponse
func (client *Client) CreatePluginAttachment(request *CreatePluginAttachmentRequest) (_result *CreatePluginAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePluginAttachmentResponse{}
	_body, _err := client.CreatePluginAttachmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom plugin class.
//
// @param request - CreatePluginClassRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePluginClassResponse
func (client *Client) CreatePluginClassWithOptions(request *CreatePluginClassRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePluginClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["alias"] = request.Alias
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ExecutePriority) {
		body["executePriority"] = request.ExecutePriority
	}

	if !dara.IsNil(request.ExecuteStage) {
		body["executeStage"] = request.ExecuteStage
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SupportedMinGatewayVersion) {
		body["supportedMinGatewayVersion"] = request.SupportedMinGatewayVersion
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	if !dara.IsNil(request.VersionDescription) {
		body["versionDescription"] = request.VersionDescription
	}

	if !dara.IsNil(request.WasmLanguage) {
		body["wasmLanguage"] = request.WasmLanguage
	}

	if !dara.IsNil(request.WasmUrl) {
		body["wasmUrl"] = request.WasmUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePluginClass"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-classes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePluginClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom plugin class.
//
// @param request - CreatePluginClassRequest
//
// @return CreatePluginClassResponse
func (client *Client) CreatePluginClass(request *CreatePluginClassRequest) (_result *CreatePluginClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePluginClassResponse{}
	_body, _err := client.CreatePluginClassWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a policy.
//
// @param request - CreatePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicyWithOptions(request *CreatePolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		body["className"] = request.ClassName
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/policies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a policy.
//
// @param request - CreatePolicyRequest
//
// @return CreatePolicyResponse
func (client *Client) CreatePolicy(request *CreatePolicyRequest) (_result *CreatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePolicyResponse{}
	_body, _err := client.CreatePolicyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a policy attachment to a resource.
//
// @param request - CreatePolicyAttachmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePolicyAttachmentResponse
func (client *Client) CreatePolicyAttachmentWithOptions(request *CreatePolicyAttachmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePolicyAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		body["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PolicyId) {
		body["policyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePolicyAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policy-attachments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePolicyAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a policy attachment to a resource.
//
// @param request - CreatePolicyAttachmentRequest
//
// @return CreatePolicyAttachmentResponse
func (client *Client) CreatePolicyAttachment(request *CreatePolicyAttachmentRequest) (_result *CreatePolicyAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePolicyAttachmentResponse{}
	_body, _err := client.CreatePolicyAttachmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a key.
//
// @param request - CreateSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretResponse
func (client *Client) CreateSecretWithOptions(request *CreateSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.GatewayType) {
		body["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.KmsConfig) {
		body["kmsConfig"] = request.KmsConfig
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.SecretData) {
		body["secretData"] = request.SecretData
	}

	if !dara.IsNil(request.SecretSource) {
		body["secretSource"] = request.SecretSource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSecret"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/secrets"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a key.
//
// @param request - CreateSecretRequest
//
// @return CreateSecretResponse
func (client *Client) CreateSecret(request *CreateSecretRequest) (_result *CreateSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSecretResponse{}
	_body, _err := client.CreateSecretWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates services.
//
// Description:
//
// This operation supports creating multiple services.
//
// @param request - CreateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceResponse
func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServiceConfigs) {
		body["serviceConfigs"] = request.ServiceConfigs
	}

	if !dara.IsNil(request.SourceType) {
		body["sourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateService"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates services.
//
// Description:
//
// This operation supports creating multiple services.
//
// @param request - CreateServiceRequest
//
// @return CreateServiceResponse
func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service version.
//
// @param request - CreateServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceVersionResponse
func (client *Client) CreateServiceVersionWithOptions(serviceId *string, request *CreateServiceVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateServiceVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["labels"] = request.Labels
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateServiceVersion"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service version.
//
// @param request - CreateServiceVersionRequest
//
// @return CreateServiceVersionResponse
func (client *Client) CreateServiceVersion(serviceId *string, request *CreateServiceVersionRequest) (_result *CreateServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceVersionResponse{}
	_body, _err := client.CreateServiceVersionWithOptions(serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a source.
//
// @param request - CreateSourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSourceResponse
func (client *Client) CreateSourceWithOptions(request *CreateSourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.K8sSourceConfig) {
		body["k8sSourceConfig"] = request.K8sSourceConfig
	}

	if !dara.IsNil(request.NacosSourceConfig) {
		body["nacosSourceConfig"] = request.NacosSourceConfig
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSource"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/sources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a source.
//
// @param request - CreateSourceRequest
//
// @return CreateSourceResponse
func (client *Client) CreateSource(request *CreateSourceRequest) (_result *CreateSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSourceResponse{}
	_body, _err := client.CreateSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除AI模型卡片
//
// @param request - DeleteAiModelCardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAiModelCardResponse
func (client *Client) DeleteAiModelCardWithOptions(modelCardId *string, request *DeleteAiModelCardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAiModelCardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAiModelCard"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-cards/" + dara.PercentEncode(dara.StringValue(modelCardId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAiModelCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AI模型卡片
//
// @param request - DeleteAiModelCardRequest
//
// @return DeleteAiModelCardResponse
func (client *Client) DeleteAiModelCard(modelCardId *string, request *DeleteAiModelCardRequest) (_result *DeleteAiModelCardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAiModelCardResponse{}
	_body, _err := client.DeleteAiModelCardWithOptions(modelCardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除AI模型供应商
//
// @param request - DeleteAiModelProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAiModelProviderResponse
func (client *Client) DeleteAiModelProviderWithOptions(modelProviderId *string, request *DeleteAiModelProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAiModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAiModelProvider"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-providers/" + dara.PercentEncode(dara.StringValue(modelProviderId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAiModelProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AI模型供应商
//
// @param request - DeleteAiModelProviderRequest
//
// @return DeleteAiModelProviderResponse
func (client *Client) DeleteAiModelProvider(modelProviderId *string, request *DeleteAiModelProviderRequest) (_result *DeleteAiModelProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAiModelProviderResponse{}
	_body, _err := client.DeleteAiModelProviderWithOptions(modelProviderId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an API consumer.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerResponse
func (client *Client) DeleteConsumerWithOptions(consumerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API consumer.
//
// @return DeleteConsumerResponse
func (client *Client) DeleteConsumer(consumerId *string) (_result *DeleteConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConsumerResponse{}
	_body, _err := client.DeleteConsumerWithOptions(consumerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an API consumer authorization rule.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerAuthorizationRuleResponse
func (client *Client) DeleteConsumerAuthorizationRuleWithOptions(consumerAuthorizationRuleId *string, consumerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerAuthorizationRuleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules/" + dara.PercentEncode(dara.StringValue(consumerAuthorizationRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API consumer authorization rule.
//
// @return DeleteConsumerAuthorizationRuleResponse
func (client *Client) DeleteConsumerAuthorizationRule(consumerAuthorizationRuleId *string, consumerId *string) (_result *DeleteConsumerAuthorizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConsumerAuthorizationRuleResponse{}
	_body, _err := client.DeleteConsumerAuthorizationRuleWithOptions(consumerAuthorizationRuleId, consumerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除消费者组
//
// @param request - DeleteConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroupWithOptions(consumerGroupId *string, request *DeleteConsumerGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerGroup"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumer-groups/" + dara.PercentEncode(dara.StringValue(consumerGroupId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除消费者组
//
// @param request - DeleteConsumerGroupRequest
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroup(consumerGroupId *string, request *DeleteConsumerGroupRequest) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a domain name.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(domainId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains/" + dara.PercentEncode(dara.StringValue(domainId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a domain name.
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(domainId *string) (_result *DeleteDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(domainId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteEnvironment is deprecated
//
// Summary:
//
// Deletes an environment.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnvironmentResponse
func (client *Client) DeleteEnvironmentWithOptions(environmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEnvironmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnvironment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments/" + dara.PercentEncode(dara.StringValue(environmentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteEnvironment is deprecated
//
// Summary:
//
// Deletes an environment.
//
// @return DeleteEnvironmentResponse
// Deprecated
func (client *Client) DeleteEnvironment(environmentId *string) (_result *DeleteEnvironmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEnvironmentResponse{}
	_body, _err := client.DeleteEnvironmentWithOptions(environmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGatewayWithOptions(gatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewayResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId))),
		Method:      dara.String("DELETE"),
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
// Deletes a gateway.
//
// @return DeleteGatewayResponse
func (client *Client) DeleteGateway(gatewayId *string) (_result *DeleteGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayResponse{}
	_body, _err := client.DeleteGatewayWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a quota throttling rule from a gateway.
//
// Description:
//
// Deletes a consumer-based quota rule from an AI gateway. This operation takes effect only for AI gateways of version 2.1.19 or later.
//
// @param request - DeleteGatewayQuotaRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewayQuotaRuleResponse
func (client *Client) DeleteGatewayQuotaRuleWithOptions(gatewayId *string, ruleId *string, request *DeleteGatewayQuotaRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewayQuotaRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewayQuotaRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/quota-rules/" + dara.PercentEncode(dara.StringValue(ruleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewayQuotaRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a quota throttling rule from a gateway.
//
// Description:
//
// Deletes a consumer-based quota rule from an AI gateway. This operation takes effect only for AI gateways of version 2.1.19 or later.
//
// @param request - DeleteGatewayQuotaRuleRequest
//
// @return DeleteGatewayQuotaRuleResponse
func (client *Client) DeleteGatewayQuotaRule(gatewayId *string, ruleId *string, request *DeleteGatewayQuotaRuleRequest) (_result *DeleteGatewayQuotaRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewayQuotaRuleResponse{}
	_body, _err := client.DeleteGatewayQuotaRuleWithOptions(gatewayId, ruleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a security group rule from a gateway.
//
// @param request - DeleteGatewaySecurityGroupRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGatewaySecurityGroupRuleResponse
func (client *Client) DeleteGatewaySecurityGroupRuleWithOptions(gatewayId *string, securityGroupRuleId *string, request *DeleteGatewaySecurityGroupRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteGatewaySecurityGroupRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CascadingDelete) {
		query["cascadingDelete"] = request.CascadingDelete
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteGatewaySecurityGroupRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/security-group-rules/" + dara.PercentEncode(dara.StringValue(securityGroupRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteGatewaySecurityGroupRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a security group rule from a gateway.
//
// @param request - DeleteGatewaySecurityGroupRuleRequest
//
// @return DeleteGatewaySecurityGroupRuleResponse
func (client *Client) DeleteGatewaySecurityGroupRule(gatewayId *string, securityGroupRuleId *string, request *DeleteGatewaySecurityGroupRuleRequest) (_result *DeleteGatewaySecurityGroupRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGatewaySecurityGroupRuleResponse{}
	_body, _err := client.DeleteGatewaySecurityGroupRuleWithOptions(gatewayId, securityGroupRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified HTTP API.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiResponse
func (client *Client) DeleteHttpApiWithOptions(httpApiId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHttpApiResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified HTTP API.
//
// @return DeleteHttpApiResponse
func (client *Client) DeleteHttpApi(httpApiId *string) (_result *DeleteHttpApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHttpApiResponse{}
	_body, _err := client.DeleteHttpApiWithOptions(httpApiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified operation.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiOperationResponse
func (client *Client) DeleteHttpApiOperationWithOptions(httpApiId *string, operationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHttpApiOperationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpApiOperation"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations/" + dara.PercentEncode(dara.StringValue(operationId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpApiOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified operation.
//
// @return DeleteHttpApiOperationResponse
func (client *Client) DeleteHttpApiOperation(httpApiId *string, operationId *string) (_result *DeleteHttpApiOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHttpApiOperationResponse{}
	_body, _err := client.DeleteHttpApiOperationWithOptions(httpApiId, operationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a route of an HTTP API.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHttpApiRouteResponse
func (client *Client) DeleteHttpApiRouteWithOptions(httpApiId *string, routeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteHttpApiRouteResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHttpApiRoute"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes/" + dara.PercentEncode(dara.StringValue(routeId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHttpApiRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a route of an HTTP API.
//
// @return DeleteHttpApiRouteResponse
func (client *Client) DeleteHttpApiRoute(httpApiId *string, routeId *string) (_result *DeleteHttpApiRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteHttpApiRouteResponse{}
	_body, _err := client.DeleteHttpApiRouteWithOptions(httpApiId, routeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an MCP server.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMcpServerResponse
func (client *Client) DeleteMcpServerWithOptions(mcpServerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMcpServerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMcpServer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an MCP server.
//
// @return DeleteMcpServerResponse
func (client *Client) DeleteMcpServer(mcpServerId *string) (_result *DeleteMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMcpServerResponse{}
	_body, _err := client.DeleteMcpServerWithOptions(mcpServerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a plugin mount.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePluginAttachmentResponse
func (client *Client) DeletePluginAttachmentWithOptions(pluginAttachmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePluginAttachmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePluginAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments/" + dara.PercentEncode(dara.StringValue(pluginAttachmentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePluginAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a plugin mount.
//
// @return DeletePluginAttachmentResponse
func (client *Client) DeletePluginAttachment(pluginAttachmentId *string) (_result *DeletePluginAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePluginAttachmentResponse{}
	_body, _err := client.DeletePluginAttachmentWithOptions(pluginAttachmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # DeletePluginClass
//
// @param request - DeletePluginClassRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePluginClassResponse
func (client *Client) DeletePluginClassWithOptions(pluginClassId *string, request *DeletePluginClassRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePluginClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePluginClass"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-classes/" + dara.PercentEncode(dara.StringValue(pluginClassId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePluginClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeletePluginClass
//
// @param request - DeletePluginClassRequest
//
// @return DeletePluginClassResponse
func (client *Client) DeletePluginClass(pluginClassId *string, request *DeletePluginClassRequest) (_result *DeletePluginClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePluginClassResponse{}
	_body, _err := client.DeletePluginClassWithOptions(pluginClassId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a policy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicyWithOptions(policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy.
//
// @return DeletePolicyResponse
func (client *Client) DeletePolicy(policyId *string) (_result *DeletePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePolicyResponse{}
	_body, _err := client.DeletePolicyWithOptions(policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a policy attachment.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePolicyAttachmentResponse
func (client *Client) DeletePolicyAttachmentWithOptions(policyAttachmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePolicyAttachmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePolicyAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policy-attachments/" + dara.PercentEncode(dara.StringValue(policyAttachmentId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePolicyAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a policy attachment.
//
// @return DeletePolicyAttachmentResponse
func (client *Client) DeletePolicyAttachment(policyAttachmentId *string) (_result *DeletePolicyAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePolicyAttachmentResponse{}
	_body, _err := client.DeletePolicyAttachmentWithOptions(policyAttachmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a key value.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretResponse
func (client *Client) DeleteSecretWithOptions(secretId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSecret"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/secrets/" + dara.PercentEncode(dara.StringValue(secretId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a key value.
//
// Description:
//
// The operation supports creating multiple services.
//
// @return DeleteSecretResponse
func (client *Client) DeleteSecret(secretId *string) (_result *DeleteSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSecretResponse{}
	_body, _err := client.DeleteSecretWithOptions(secretId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceResponse
func (client *Client) DeleteServiceWithOptions(serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteService"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service.
//
// @return DeleteServiceResponse
func (client *Client) DeleteService(serviceId *string) (_result *DeleteServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(serviceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a service version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceVersionResponse
func (client *Client) DeleteServiceVersionWithOptions(serviceId *string, name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteServiceVersionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteServiceVersion"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId)) + "/versions/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service version.
//
// @return DeleteServiceVersionResponse
func (client *Client) DeleteServiceVersion(serviceId *string, name *string) (_result *DeleteServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceVersionResponse{}
	_body, _err := client.DeleteServiceVersionWithOptions(serviceId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a service source.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSourceResponse
func (client *Client) DeleteSourceWithOptions(sourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSource"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/sources/" + dara.PercentEncode(dara.StringValue(sourceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service source.
//
// @return DeleteSourceResponse
func (client *Client) DeleteSource(sourceId *string) (_result *DeleteSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSourceResponse{}
	_body, _err := client.DeleteSourceWithOptions(sourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes an HTTP API. This includes REST APIs and routes within HTTP APIs.
//
// @param request - DeployHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployHttpApiResponse
func (client *Client) DeployHttpApiWithOptions(httpApiId *string, request *DeployHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployHttpApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HttpApiConfig) {
		body["httpApiConfig"] = request.HttpApiConfig
	}

	if !dara.IsNil(request.RestApiConfig) {
		body["restApiConfig"] = request.RestApiConfig
	}

	if !dara.IsNil(request.RouteId) {
		body["routeId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/deploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes an HTTP API. This includes REST APIs and routes within HTTP APIs.
//
// @param request - DeployHttpApiRequest
//
// @return DeployHttpApiResponse
func (client *Client) DeployHttpApi(httpApiId *string, request *DeployHttpApiRequest) (_result *DeployHttpApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployHttpApiResponse{}
	_body, _err := client.DeployHttpApiWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes an MCP server.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployMcpServerResponse
func (client *Client) DeployMcpServerWithOptions(mcpServerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeployMcpServerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployMcpServer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId)) + "/deploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes an MCP server.
//
// @return DeployMcpServerResponse
func (client *Client) DeployMcpServer(mcpServerId *string) (_result *DeployMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeployMcpServerResponse{}
	_body, _err := client.DeployMcpServerWithOptions(mcpServerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions where the cloud-native API gateway is available for the current account.
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/regions"),
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
// Queries the regions where the cloud-native API gateway is available for the current account.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Exports a specified HTTP API.
//
// @param request - ExportHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportHttpApiResponse
func (client *Client) ExportHttpApiWithOptions(httpApiId *string, request *ExportHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExportHttpApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtensionConfig) {
		body["extensionConfig"] = request.ExtensionConfig
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.OperationIds) {
		body["operationIds"] = request.OperationIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/export"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports a specified HTTP API.
//
// @param request - ExportHttpApiRequest
//
// @return ExportHttpApiResponse
func (client *Client) ExportHttpApi(httpApiId *string, request *ExportHttpApiRequest) (_result *ExportHttpApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExportHttpApiResponse{}
	_body, _err := client.ExportHttpApiWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询AI模型卡片详情
//
// @param request - GetAiModelCardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiModelCardResponse
func (client *Client) GetAiModelCardWithOptions(modelCardId *string, request *GetAiModelCardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAiModelCardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAiModelCard"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-cards/" + dara.PercentEncode(dara.StringValue(modelCardId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAiModelCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AI模型卡片详情
//
// @param request - GetAiModelCardRequest
//
// @return GetAiModelCardResponse
func (client *Client) GetAiModelCard(modelCardId *string, request *GetAiModelCardRequest) (_result *GetAiModelCardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAiModelCardResponse{}
	_body, _err := client.GetAiModelCardWithOptions(modelCardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询AI模型供应商详情
//
// @param request - GetAiModelProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAiModelProviderResponse
func (client *Client) GetAiModelProviderWithOptions(modelProviderId *string, request *GetAiModelProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAiModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAiModelProvider"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-providers/" + dara.PercentEncode(dara.StringValue(modelProviderId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAiModelProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AI模型供应商详情
//
// @param request - GetAiModelProviderRequest
//
// @return GetAiModelProviderResponse
func (client *Client) GetAiModelProvider(modelProviderId *string, request *GetAiModelProviderRequest) (_result *GetAiModelProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAiModelProviderResponse{}
	_body, _err := client.GetAiModelProviderWithOptions(modelProviderId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves an API consumer.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerResponse
func (client *Client) GetConsumerWithOptions(consumerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an API consumer.
//
// @return GetConsumerResponse
func (client *Client) GetConsumer(consumerId *string) (_result *GetConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerResponse{}
	_body, _err := client.GetConsumerWithOptions(consumerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a consumer authorization rule.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerAuthorizationRuleResponse
func (client *Client) GetConsumerAuthorizationRuleWithOptions(consumerAuthorizationRuleId *string, consumerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerAuthorizationRuleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules/" + dara.PercentEncode(dara.StringValue(consumerAuthorizationRuleId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a consumer authorization rule.
//
// @return GetConsumerAuthorizationRuleResponse
func (client *Client) GetConsumerAuthorizationRule(consumerAuthorizationRuleId *string, consumerId *string) (_result *GetConsumerAuthorizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerAuthorizationRuleResponse{}
	_body, _err := client.GetConsumerAuthorizationRuleWithOptions(consumerAuthorizationRuleId, consumerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询消费者组
//
// @param request - GetConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerGroupResponse
func (client *Client) GetConsumerGroupWithOptions(consumerGroupId *string, request *GetConsumerGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerGroup"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumer-groups/" + dara.PercentEncode(dara.StringValue(consumerGroupId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者组
//
// @param request - GetConsumerGroupRequest
//
// @return GetConsumerGroupResponse
func (client *Client) GetConsumerGroup(consumerGroupId *string, request *GetConsumerGroupRequest) (_result *GetConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerGroupResponse{}
	_body, _err := client.GetConsumerGroupWithOptions(consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the monitoring log dashboard.
//
// @param tmpReq - GetDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDashboardResponse
func (client *Client) GetDashboardWithOptions(gatewayId *string, tmpReq *GetDashboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDashboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDashboardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ApiId) {
		query["apiId"] = request.ApiId
	}

	if !dara.IsNil(request.FilterShrink) {
		query["filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PluginClassId) {
		query["pluginClassId"] = request.PluginClassId
	}

	if !dara.IsNil(request.PluginId) {
		query["pluginId"] = request.PluginId
	}

	if !dara.IsNil(request.RouteId) {
		query["routeId"] = request.RouteId
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.UpstreamCluster) {
		query["upstreamCluster"] = request.UpstreamCluster
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDashboard"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/dashboards"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDashboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the monitoring log dashboard.
//
// @param request - GetDashboardRequest
//
// @return GetDashboardResponse
func (client *Client) GetDashboard(gatewayId *string, request *GetDashboardRequest) (_result *GetDashboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDashboardResponse{}
	_body, _err := client.GetDashboardWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a domain name.
//
// @param request - GetDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithOptions(domainId *string, request *GetDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WithStatistics) {
		query["withStatistics"] = request.WithStatistics
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains/" + dara.PercentEncode(dara.StringValue(domainId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a domain name.
//
// @param request - GetDomainRequest
//
// @return GetDomainResponse
func (client *Client) GetDomain(domainId *string, request *GetDomainRequest) (_result *GetDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(domainId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetEnvironment is deprecated
//
// Summary:
//
// Queries an environment.
//
// @param request - GetEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnvironmentResponse
func (client *Client) GetEnvironmentWithOptions(environmentId *string, request *GetEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEnvironmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WithStatistics) {
		query["withStatistics"] = request.WithStatistics
	}

	if !dara.IsNil(request.WithVpcInfo) {
		query["withVpcInfo"] = request.WithVpcInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnvironment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments/" + dara.PercentEncode(dara.StringValue(environmentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetEnvironment is deprecated
//
// Summary:
//
// Queries an environment.
//
// @param request - GetEnvironmentRequest
//
// @return GetEnvironmentResponse
// Deprecated
func (client *Client) GetEnvironment(environmentId *string, request *GetEnvironmentRequest) (_result *GetEnvironmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEnvironmentResponse{}
	_body, _err := client.GetEnvironmentWithOptions(environmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves basic information about a gateway, including the associated VPC, vSwitch, and gateway ingress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayResponse
func (client *Client) GetGatewayWithOptions(gatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGatewayResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves basic information about a gateway, including the associated VPC, vSwitch, and gateway ingress.
//
// @return GetGatewayResponse
func (client *Client) GetGateway(gatewayId *string) (_result *GetGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGatewayResponse{}
	_body, _err := client.GetGatewayWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a gateway quota rate limiting rule.
//
// Description:
//
// Queries the details of a consumer quota rule on an AI gateway.
//
// @param request - GetGatewayQuotaRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayQuotaRuleResponse
func (client *Client) GetGatewayQuotaRuleWithOptions(gatewayId *string, ruleId *string, request *GetGatewayQuotaRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGatewayQuotaRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerPageNumber) {
		query["consumerPageNumber"] = request.ConsumerPageNumber
	}

	if !dara.IsNil(request.ConsumerPageSize) {
		query["consumerPageSize"] = request.ConsumerPageSize
	}

	if !dara.IsNil(request.WithConsumers) {
		query["withConsumers"] = request.WithConsumers
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayQuotaRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/quota-rules/" + dara.PercentEncode(dara.StringValue(ruleId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayQuotaRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a gateway quota rate limiting rule.
//
// Description:
//
// Queries the details of a consumer quota rule on an AI gateway.
//
// @param request - GetGatewayQuotaRuleRequest
//
// @return GetGatewayQuotaRuleResponse
func (client *Client) GetGatewayQuotaRule(gatewayId *string, ruleId *string, request *GetGatewayQuotaRuleRequest) (_result *GetGatewayQuotaRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGatewayQuotaRuleResponse{}
	_body, _err := client.GetGatewayQuotaRuleWithOptions(gatewayId, ruleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage details of a subject under a gateway quota throttling rule, including used quota, total quota, whether the limit is exceeded, usage details, and consumption records.
//
// Description:
//
// Retrieves the usage details of a specific consumer under a quota rule. This operation takes effect only for AI gateways with a version later than 2.1.19.
//
// @param request - GetGatewayQuotaRuleSubjectUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGatewayQuotaRuleSubjectUsageResponse
func (client *Client) GetGatewayQuotaRuleSubjectUsageWithOptions(gatewayId *string, ruleId *string, subjectId *string, request *GetGatewayQuotaRuleSubjectUsageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGatewayQuotaRuleSubjectUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGatewayQuotaRuleSubjectUsage"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/quota-rules/" + dara.PercentEncode(dara.StringValue(ruleId)) + "/subjects/" + dara.PercentEncode(dara.StringValue(subjectId)) + "/usage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGatewayQuotaRuleSubjectUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage details of a subject under a gateway quota throttling rule, including used quota, total quota, whether the limit is exceeded, usage details, and consumption records.
//
// Description:
//
// Retrieves the usage details of a specific consumer under a quota rule. This operation takes effect only for AI gateways with a version later than 2.1.19.
//
// @param request - GetGatewayQuotaRuleSubjectUsageRequest
//
// @return GetGatewayQuotaRuleSubjectUsageResponse
func (client *Client) GetGatewayQuotaRuleSubjectUsage(gatewayId *string, ruleId *string, subjectId *string, request *GetGatewayQuotaRuleSubjectUsageRequest) (_result *GetGatewayQuotaRuleSubjectUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGatewayQuotaRuleSubjectUsageResponse{}
	_body, _err := client.GetGatewayQuotaRuleSubjectUsageWithOptions(gatewayId, ruleId, subjectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves HTTP API information.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiResponse
func (client *Client) GetHttpApiWithOptions(httpApiId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHttpApiResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves HTTP API information.
//
// @return GetHttpApiResponse
func (client *Client) GetHttpApi(httpApiId *string) (_result *GetHttpApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHttpApiResponse{}
	_body, _err := client.GetHttpApiWithOptions(httpApiId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves operation information.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiOperationResponse
func (client *Client) GetHttpApiOperationWithOptions(httpApiId *string, operationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHttpApiOperationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHttpApiOperation"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations/" + dara.PercentEncode(dara.StringValue(operationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpApiOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves operation information.
//
// @return GetHttpApiOperationResponse
func (client *Client) GetHttpApiOperation(httpApiId *string, operationId *string) (_result *GetHttpApiOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHttpApiOperationResponse{}
	_body, _err := client.GetHttpApiOperationWithOptions(httpApiId, operationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the route details of an HTTP API.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHttpApiRouteResponse
func (client *Client) GetHttpApiRouteWithOptions(httpApiId *string, routeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHttpApiRouteResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHttpApiRoute"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes/" + dara.PercentEncode(dara.StringValue(routeId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHttpApiRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the route details of an HTTP API.
//
// @return GetHttpApiRouteResponse
func (client *Client) GetHttpApiRoute(httpApiId *string, routeId *string) (_result *GetHttpApiRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetHttpApiRouteResponse{}
	_body, _err := client.GetHttpApiRouteWithOptions(httpApiId, routeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves an MCP server.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMcpServerResponse
func (client *Client) GetMcpServerWithOptions(mcpServerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMcpServerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMcpServer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves an MCP server.
//
// Description:
//
// The operation supports creating multiple services.
//
// @return GetMcpServerResponse
func (client *Client) GetMcpServer(mcpServerId *string) (_result *GetMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMcpServerResponse{}
	_body, _err := client.GetMcpServerWithOptions(mcpServerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a plugin mount.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPluginAttachmentResponse
func (client *Client) GetPluginAttachmentWithOptions(pluginAttachmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPluginAttachmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPluginAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments/" + dara.PercentEncode(dara.StringValue(pluginAttachmentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPluginAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a plugin mount.
//
// @return GetPluginAttachmentResponse
func (client *Client) GetPluginAttachment(pluginAttachmentId *string) (_result *GetPluginAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPluginAttachmentResponse{}
	_body, _err := client.GetPluginAttachmentWithOptions(pluginAttachmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a custom plug-in class.
//
// @param request - GetPluginClassRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPluginClassResponse
func (client *Client) GetPluginClassWithOptions(pluginClassId *string, request *GetPluginClassRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPluginClassResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPluginClass"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-classes/" + dara.PercentEncode(dara.StringValue(pluginClassId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPluginClassResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a custom plug-in class.
//
// @param request - GetPluginClassRequest
//
// @return GetPluginClassResponse
func (client *Client) GetPluginClass(pluginClassId *string, request *GetPluginClassRequest) (_result *GetPluginClassResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPluginClassResponse{}
	_body, _err := client.GetPluginClassWithOptions(pluginClassId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a policy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyResponse
func (client *Client) GetPolicyWithOptions(policyId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a policy.
//
// @return GetPolicyResponse
func (client *Client) GetPolicy(policyId *string) (_result *GetPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPolicyResponse{}
	_body, _err := client.GetPolicyWithOptions(policyId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resource attachment of a policy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPolicyAttachmentResponse
func (client *Client) GetPolicyAttachmentWithOptions(policyAttachmentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPolicyAttachmentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPolicyAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policy-attachments/" + dara.PercentEncode(dara.StringValue(policyAttachmentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPolicyAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resource attachment of a policy.
//
// @return GetPolicyAttachmentResponse
func (client *Client) GetPolicyAttachment(policyAttachmentId *string) (_result *GetPolicyAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPolicyAttachmentResponse{}
	_body, _err := client.GetPolicyAttachmentWithOptions(policyAttachmentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves resource overview information.
//
// @param request - GetResourceOverviewRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceOverviewResponse
func (client *Client) GetResourceOverviewWithOptions(request *GetResourceOverviewRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetResourceOverviewResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceOverview"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/overview/resources"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceOverviewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves resource overview information.
//
// @param request - GetResourceOverviewRequest
//
// @return GetResourceOverviewResponse
func (client *Client) GetResourceOverview(request *GetResourceOverviewRequest) (_result *GetResourceOverviewResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetResourceOverviewResponse{}
	_body, _err := client.GetResourceOverviewWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a key.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretResponse
func (client *Client) GetSecretWithOptions(secretId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSecretResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecret"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/secrets/" + dara.PercentEncode(dara.StringValue(secretId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a key.
//
// Description:
//
// The operation supports creating multiple services.
//
// @return GetSecretResponse
func (client *Client) GetSecret(secretId *string) (_result *GetSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSecretResponse{}
	_body, _err := client.GetSecretWithOptions(secretId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the value of a key.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretValueResponse
func (client *Client) GetSecretValueWithOptions(name *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSecretValueResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSecretValue"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/secrets/name/" + dara.PercentEncode(dara.StringValue(name)) + "/value"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSecretValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the value of a key.
//
// Description:
//
// The operation supports creating multiple services.
//
// @return GetSecretValueResponse
func (client *Client) GetSecretValue(name *string) (_result *GetSecretValueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSecretValueResponse{}
	_body, _err := client.GetSecretValueWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceResponse
func (client *Client) GetServiceWithOptions(serviceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetService"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a service.
//
// @return GetServiceResponse
func (client *Client) GetService(serviceId *string) (_result *GetServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceResponse{}
	_body, _err := client.GetServiceWithOptions(serviceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a service source.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSourceResponse
func (client *Client) GetSourceWithOptions(sourceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSourceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSource"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/sources/" + dara.PercentEncode(dara.StringValue(sourceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a service source.
//
// @return GetSourceResponse
func (client *Client) GetSource(sourceId *string) (_result *GetSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSourceResponse{}
	_body, _err := client.GetSourceWithOptions(sourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the Tracing Analysis configuration.
//
// @param request - GetTraceConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceConfigResponse
func (client *Client) GetTraceConfigWithOptions(gatewayId *string, request *GetTraceConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTraceConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTraceConfig"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/trace"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTraceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the Tracing Analysis configuration.
//
// @param request - GetTraceConfigRequest
//
// @return GetTraceConfigResponse
func (client *Client) GetTraceConfig(gatewayId *string, request *GetTraceConfigRequest) (_result *GetTraceConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTraceConfigResponse{}
	_body, _err := client.GetTraceConfigWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports an HTTP API. This operation supports importing OpenAPI 2.0 and OpenAPI 3.0.x definition files as REST-type APIs.
//
// @param request - ImportHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportHttpApiResponse
func (client *Client) ImportHttpApiWithOptions(request *ImportHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ImportHttpApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.McpRouteId) {
		body["mcpRouteId"] = request.McpRouteId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SpecContentBase64) {
		body["specContentBase64"] = request.SpecContentBase64
	}

	if !dara.IsNil(request.SpecFileUrl) {
		body["specFileUrl"] = request.SpecFileUrl
	}

	if !dara.IsNil(request.SpecOssConfig) {
		body["specOssConfig"] = request.SpecOssConfig
	}

	if !dara.IsNil(request.Strategy) {
		body["strategy"] = request.Strategy
	}

	if !dara.IsNil(request.TargetHttpApiId) {
		body["targetHttpApiId"] = request.TargetHttpApiId
	}

	if !dara.IsNil(request.VersionConfig) {
		body["versionConfig"] = request.VersionConfig
	}

	if !dara.IsNil(request.WithGatewayExtension) {
		body["withGatewayExtension"] = request.WithGatewayExtension
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/import"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports an HTTP API. This operation supports importing OpenAPI 2.0 and OpenAPI 3.0.x definition files as REST-type APIs.
//
// @param request - ImportHttpApiRequest
//
// @return ImportHttpApiResponse
func (client *Client) ImportHttpApi(request *ImportHttpApiRequest) (_result *ImportHttpApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImportHttpApiResponse{}
	_body, _err := client.ImportHttpApiWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs a plug-in.
//
// @param request - InstallPluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallPluginResponse
func (client *Client) InstallPluginWithOptions(request *InstallPluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InstallPluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GatewayIds) {
		body["gatewayIds"] = request.GatewayIds
	}

	if !dara.IsNil(request.PluginClassId) {
		body["pluginClassId"] = request.PluginClassId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InstallPlugin"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugins/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InstallPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs a plug-in.
//
// @param request - InstallPluginRequest
//
// @return InstallPluginResponse
func (client *Client) InstallPlugin(request *InstallPluginRequest) (_result *InstallPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InstallPluginResponse{}
	_body, _err := client.InstallPluginWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询AI模型卡片列表
//
// @param request - ListAiModelCardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAiModelCardsResponse
func (client *Client) ListAiModelCardsWithOptions(request *ListAiModelCardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAiModelCardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAiModelCards"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-cards"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAiModelCardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AI模型卡片列表
//
// @param request - ListAiModelCardsRequest
//
// @return ListAiModelCardsResponse
func (client *Client) ListAiModelCards(request *ListAiModelCardsRequest) (_result *ListAiModelCardsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAiModelCardsResponse{}
	_body, _err := client.ListAiModelCardsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询AI模型供应商列表
//
// @param request - ListAiModelProvidersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAiModelProvidersResponse
func (client *Client) ListAiModelProvidersWithOptions(request *ListAiModelProvidersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAiModelProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Provider) {
		query["provider"] = request.Provider
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAiModelProviders"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-providers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAiModelProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AI模型供应商列表
//
// @param request - ListAiModelProvidersRequest
//
// @return ListAiModelProvidersResponse
func (client *Client) ListAiModelProviders(request *ListAiModelProvidersRequest) (_result *ListAiModelProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAiModelProvidersResponse{}
	_body, _err := client.ListAiModelProvidersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of consumer authorization rules.
//
// @param request - ListConsumerAuthorizationRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerAuthorizationRulesResponse
func (client *Client) ListConsumerAuthorizationRulesWithOptions(consumerId *string, request *ListConsumerAuthorizationRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerAuthorizationRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiNameLike) {
		query["apiNameLike"] = request.ApiNameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumerAuthorizationRules"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumerAuthorizationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of consumer authorization rules.
//
// @param request - ListConsumerAuthorizationRulesRequest
//
// @return ListConsumerAuthorizationRulesResponse
func (client *Client) ListConsumerAuthorizationRules(consumerId *string, request *ListConsumerAuthorizationRulesRequest) (_result *ListConsumerAuthorizationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerAuthorizationRulesResponse{}
	_body, _err := client.ListConsumerAuthorizationRulesWithOptions(consumerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询消费者组成员列表
//
// @param request - ListConsumerGroupConsumersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupConsumersResponse
func (client *Client) ListConsumerGroupConsumersWithOptions(consumerGroupId *string, request *ListConsumerGroupConsumersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerGroupConsumersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumerGroupConsumers"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumer-groups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/consumers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumerGroupConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者组成员列表
//
// @param request - ListConsumerGroupConsumersRequest
//
// @return ListConsumerGroupConsumersResponse
func (client *Client) ListConsumerGroupConsumers(consumerGroupId *string, request *ListConsumerGroupConsumersRequest) (_result *ListConsumerGroupConsumersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerGroupConsumersResponse{}
	_body, _err := client.ListConsumerGroupConsumersWithOptions(consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询消费者组列表
//
// @param request - ListConsumerGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupsResponse
func (client *Client) ListConsumerGroupsWithOptions(request *ListConsumerGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumerGroups"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumer-groups"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者组列表
//
// @param request - ListConsumerGroupsRequest
//
// @return ListConsumerGroupsResponse
func (client *Client) ListConsumerGroups(request *ListConsumerGroupsRequest) (_result *ListConsumerGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerGroupsResponse{}
	_body, _err := client.ListConsumerGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of quota rules configured for a specific consumer.
//
// @param request - ListConsumerQuotaRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerQuotaRulesResponse
func (client *Client) ListConsumerQuotaRulesWithOptions(consumerId *string, request *ListConsumerQuotaRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerQuotaRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumerQuotaRules"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/quota-rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumerQuotaRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of quota rules configured for a specific consumer.
//
// @param request - ListConsumerQuotaRulesRequest
//
// @return ListConsumerQuotaRulesResponse
func (client *Client) ListConsumerQuotaRules(consumerId *string, request *ListConsumerQuotaRulesRequest) (_result *ListConsumerQuotaRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerQuotaRulesResponse{}
	_body, _err := client.ListConsumerQuotaRulesWithOptions(consumerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of consumers.
//
// @param request - ListConsumersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumersResponse
func (client *Client) ListConsumersWithOptions(request *ListConsumersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumers"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of consumers.
//
// @param request - ListConsumersRequest
//
// @return ListConsumersResponse
func (client *Client) ListConsumers(request *ListConsumersRequest) (_result *ListConsumersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumersResponse{}
	_body, _err := client.ListConsumersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of domain names.
//
// @param request - ListDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainScope) {
		query["domainScope"] = request.DomainScope
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomains"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of domain names.
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListEnvironments is deprecated
//
// Summary:
//
// Queries the list of environments.
//
// @param request - ListEnvironmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnvironmentsResponse
func (client *Client) ListEnvironmentsWithOptions(request *ListEnvironmentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEnvironmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasLike) {
		query["aliasLike"] = request.AliasLike
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayNameLike) {
		query["gatewayNameLike"] = request.GatewayNameLike
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnvironments"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListEnvironments is deprecated
//
// Summary:
//
// Queries the list of environments.
//
// @param request - ListEnvironmentsRequest
//
// @return ListEnvironmentsResponse
// Deprecated
func (client *Client) ListEnvironments(request *ListEnvironmentsRequest) (_result *ListEnvironmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEnvironmentsResponse{}
	_body, _err := client.ListEnvironmentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the external service information of a gateway.
//
// Description:
//
// This operation supports creating multiple services.
//
// @param request - ListExternalServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExternalServicesResponse
func (client *Client) ListExternalServicesWithOptions(gatewayId *string, request *ListExternalServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExternalServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImportableOnly) {
		query["importableOnly"] = request.ImportableOnly
	}

	if !dara.IsNil(request.Limit) {
		query["limit"] = request.Limit
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PaiWorkspaceId) {
		query["paiWorkspaceId"] = request.PaiWorkspaceId
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExternalServices"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/external-services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExternalServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the external service information of a gateway.
//
// Description:
//
// This operation supports creating multiple services.
//
// @param request - ListExternalServicesRequest
//
// @return ListExternalServicesResponse
func (client *Client) ListExternalServices(gatewayId *string, request *ListExternalServicesRequest) (_result *ListExternalServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExternalServicesResponse{}
	_body, _err := client.ListExternalServicesWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of gateway attribute parameter settings.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayFeaturesResponse
func (client *Client) ListGatewayFeaturesWithOptions(gatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewayFeaturesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayFeatures"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/gateway-features"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayFeaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of gateway attribute parameter settings.
//
// @return ListGatewayFeaturesResponse
func (client *Client) ListGatewayFeatures(gatewayId *string) (_result *ListGatewayFeaturesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewayFeaturesResponse{}
	_body, _err := client.ListGatewayFeaturesWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of consumer quota rules bound to a gateway.
//
// Description:
//
// Queries the list of consumer quota rules bound to a gateway.
//
// @param request - ListGatewayQuotaRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewayQuotaRulesResponse
func (client *Client) ListGatewayQuotaRulesWithOptions(gatewayId *string, request *ListGatewayQuotaRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewayQuotaRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGatewayQuotaRules"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/quota-rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewayQuotaRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of consumer quota rules bound to a gateway.
//
// Description:
//
// Queries the list of consumer quota rules bound to a gateway.
//
// @param request - ListGatewayQuotaRulesRequest
//
// @return ListGatewayQuotaRulesResponse
func (client *Client) ListGatewayQuotaRules(gatewayId *string, request *ListGatewayQuotaRulesRequest) (_result *ListGatewayQuotaRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewayQuotaRulesResponse{}
	_body, _err := client.ListGatewayQuotaRulesWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of gateways.
//
// @param tmpReq - ListGatewaysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGatewaysResponse
func (client *Client) ListGatewaysWithOptions(tmpReq *ListGatewaysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListGatewaysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListGatewaysShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.TagShrink) {
		query["tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListGateways"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListGatewaysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of gateways.
//
// @param request - ListGatewaysRequest
//
// @return ListGatewaysResponse
func (client *Client) ListGateways(request *ListGatewaysRequest) (_result *ListGatewaysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGatewaysResponse{}
	_body, _err := client.ListGatewaysWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of API operations.
//
// @param request - ListHttpApiOperationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApiOperationsResponse
func (client *Client) ListHttpApiOperationsWithOptions(httpApiId *string, request *ListHttpApiOperationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHttpApiOperationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerAuthorizationRuleId) {
		query["consumerAuthorizationRuleId"] = request.ConsumerAuthorizationRuleId
	}

	if !dara.IsNil(request.EnableAuth) {
		query["enableAuth"] = request.EnableAuth
	}

	if !dara.IsNil(request.ForDeploy) {
		query["forDeploy"] = request.ForDeploy
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Method) {
		query["method"] = request.Method
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PathLike) {
		query["pathLike"] = request.PathLike
	}

	if !dara.IsNil(request.WithConsumerInEnvironmentId) {
		query["withConsumerInEnvironmentId"] = request.WithConsumerInEnvironmentId
	}

	if !dara.IsNil(request.WithConsumerInfoById) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !dara.IsNil(request.WithPluginAttachmentByPluginId) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHttpApiOperations"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpApiOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of API operations.
//
// @param request - ListHttpApiOperationsRequest
//
// @return ListHttpApiOperationsResponse
func (client *Client) ListHttpApiOperations(httpApiId *string, request *ListHttpApiOperationsRequest) (_result *ListHttpApiOperationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHttpApiOperationsResponse{}
	_body, _err := client.ListHttpApiOperationsWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the route list of an HTTP API.
//
// @param request - ListHttpApiRoutesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApiRoutesResponse
func (client *Client) ListHttpApiRoutesWithOptions(httpApiId *string, request *ListHttpApiRoutesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHttpApiRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackendServiceName) {
		query["backendServiceName"] = request.BackendServiceName
	}

	if !dara.IsNil(request.ConsumerAuthorizationRuleId) {
		query["consumerAuthorizationRuleId"] = request.ConsumerAuthorizationRuleId
	}

	if !dara.IsNil(request.DeployStatuses) {
		query["deployStatuses"] = request.DeployStatuses
	}

	if !dara.IsNil(request.DomainId) {
		query["domainId"] = request.DomainId
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.ForDeploy) {
		query["forDeploy"] = request.ForDeploy
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PathLike) {
		query["pathLike"] = request.PathLike
	}

	if !dara.IsNil(request.WithAuthPolicyInfo) {
		query["withAuthPolicyInfo"] = request.WithAuthPolicyInfo
	}

	if !dara.IsNil(request.WithConsumerInfoById) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !dara.IsNil(request.WithPluginAttachmentByPluginId) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHttpApiRoutes"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpApiRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the route list of an HTTP API.
//
// @param request - ListHttpApiRoutesRequest
//
// @return ListHttpApiRoutesResponse
func (client *Client) ListHttpApiRoutes(httpApiId *string, request *ListHttpApiRoutesRequest) (_result *ListHttpApiRoutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHttpApiRoutesResponse{}
	_body, _err := client.ListHttpApiRoutesWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of HTTP APIs.
//
// @param request - ListHttpApisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHttpApisResponse
func (client *Client) ListHttpApisWithOptions(request *ListHttpApisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListHttpApisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Types) {
		query["types"] = request.Types
	}

	if !dara.IsNil(request.WithAPIsPublishedToEnvironment) {
		query["withAPIsPublishedToEnvironment"] = request.WithAPIsPublishedToEnvironment
	}

	if !dara.IsNil(request.WithAuthPolicyInEnvironmentId) {
		query["withAuthPolicyInEnvironmentId"] = request.WithAuthPolicyInEnvironmentId
	}

	if !dara.IsNil(request.WithAuthPolicyList) {
		query["withAuthPolicyList"] = request.WithAuthPolicyList
	}

	if !dara.IsNil(request.WithConsumerInfoById) {
		query["withConsumerInfoById"] = request.WithConsumerInfoById
	}

	if !dara.IsNil(request.WithEnvironmentInfo) {
		query["withEnvironmentInfo"] = request.WithEnvironmentInfo
	}

	if !dara.IsNil(request.WithEnvironmentInfoById) {
		query["withEnvironmentInfoById"] = request.WithEnvironmentInfoById
	}

	if !dara.IsNil(request.WithIngressInfo) {
		query["withIngressInfo"] = request.WithIngressInfo
	}

	if !dara.IsNil(request.WithPluginAttachmentByPluginId) {
		query["withPluginAttachmentByPluginId"] = request.WithPluginAttachmentByPluginId
	}

	if !dara.IsNil(request.WithPolicyConfigs) {
		query["withPolicyConfigs"] = request.WithPolicyConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHttpApis"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHttpApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of HTTP APIs.
//
// @param request - ListHttpApisRequest
//
// @return ListHttpApisResponse
func (client *Client) ListHttpApis(request *ListHttpApisRequest) (_result *ListHttpApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListHttpApisResponse{}
	_body, _err := client.ListHttpApisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of MCP servers.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param request - ListMcpServersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMcpServersResponse
func (client *Client) ListMcpServersWithOptions(request *ListMcpServersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMcpServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateFromTypes) {
		query["createFromTypes"] = request.CreateFromTypes
	}

	if !dara.IsNil(request.DeployStatuses) {
		query["deployStatuses"] = request.DeployStatuses
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMcpServers"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/mcp-servers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMcpServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of MCP servers.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param request - ListMcpServersRequest
//
// @return ListMcpServersResponse
func (client *Client) ListMcpServers(request *ListMcpServersRequest) (_result *ListMcpServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMcpServersResponse{}
	_body, _err := client.ListMcpServersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of plug-in mounts.
//
// @param request - ListPluginAttachmentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginAttachmentsResponse
func (client *Client) ListPluginAttachmentsWithOptions(request *ListPluginAttachmentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginAttachmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		query["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.AttachResourceTypes) {
		query["attachResourceTypes"] = request.AttachResourceTypes
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginId) {
		query["pluginId"] = request.PluginId
	}

	if !dara.IsNil(request.WithParentResource) {
		query["withParentResource"] = request.WithParentResource
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPluginAttachments"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPluginAttachmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of plug-in mounts.
//
// @param request - ListPluginAttachmentsRequest
//
// @return ListPluginAttachmentsResponse
func (client *Client) ListPluginAttachments(request *ListPluginAttachmentsRequest) (_result *ListPluginAttachmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPluginAttachmentsResponse{}
	_body, _err := client.ListPluginAttachmentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves plug-ins.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param request - ListPluginClassesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginClassesResponse
func (client *Client) ListPluginClassesWithOptions(request *ListPluginClassesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginClassesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliasLike) {
		query["aliasLike"] = request.AliasLike
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.ExcludeBuiltinAiProxy) {
		query["excludeBuiltinAiProxy"] = request.ExcludeBuiltinAiProxy
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.Installed) {
		query["installed"] = request.Installed
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPluginClasses"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-classes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPluginClassesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves plug-ins.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param request - ListPluginClassesRequest
//
// @return ListPluginClassesResponse
func (client *Client) ListPluginClasses(request *ListPluginClassesRequest) (_result *ListPluginClassesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPluginClassesResponse{}
	_body, _err := client.ListPluginClassesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of plugins.
//
// @param request - ListPluginsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPluginsResponse
func (client *Client) ListPluginsWithOptions(request *ListPluginsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPluginsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		query["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.IncludeBuiltinAiGateway) {
		query["includeBuiltinAiGateway"] = request.IncludeBuiltinAiGateway
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginClassId) {
		query["pluginClassId"] = request.PluginClassId
	}

	if !dara.IsNil(request.PluginClassName) {
		query["pluginClassName"] = request.PluginClassName
	}

	if !dara.IsNil(request.WithAttachmentInfo) {
		query["withAttachmentInfo"] = request.WithAttachmentInfo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPlugins"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugins"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of plugins.
//
// @param request - ListPluginsRequest
//
// @return ListPluginsResponse
func (client *Client) ListPlugins(request *ListPluginsRequest) (_result *ListPluginsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPluginsResponse{}
	_body, _err := client.ListPluginsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of policies.
//
// @param request - ListPoliciesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPoliciesResponse
func (client *Client) ListPoliciesWithOptions(request *ListPoliciesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		query["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.WithAttachments) {
		query["withAttachments"] = request.WithAttachments
	}

	if !dara.IsNil(request.WithSystemPolicy) {
		query["withSystemPolicy"] = request.WithSystemPolicy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicies"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of policies.
//
// @param request - ListPoliciesRequest
//
// @return ListPoliciesResponse
func (client *Client) ListPolicies(request *ListPoliciesRequest) (_result *ListPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPoliciesResponse{}
	_body, _err := client.ListPoliciesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves policy templates.
//
// @param request - ListPolicyClassesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicyClassesResponse
func (client *Client) ListPolicyClassesWithOptions(request *ListPolicyClassesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPolicyClassesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceId) {
		query["attachResourceId"] = request.AttachResourceId
	}

	if !dara.IsNil(request.AttachResourceType) {
		query["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.Direction) {
		query["direction"] = request.Direction
	}

	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicyClasses"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policy-classes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicyClassesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves policy templates.
//
// @param request - ListPolicyClassesRequest
//
// @return ListPolicyClassesResponse
func (client *Client) ListPolicyClasses(request *ListPolicyClassesRequest) (_result *ListPolicyClassesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPolicyClassesResponse{}
	_body, _err := client.ListPolicyClassesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists secret references.
//
// Description:
//
// This operation supports creating multiple services.
//
// @param request - ListSecretReferencesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretReferencesResponse
func (client *Client) ListSecretReferencesWithOptions(secretId *string, request *ListSecretReferencesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSecretReferencesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecretReferences"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/secrets/" + dara.PercentEncode(dara.StringValue(secretId)) + "/references"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecretReferencesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists secret references.
//
// Description:
//
// This operation supports creating multiple services.
//
// @param request - ListSecretReferencesRequest
//
// @return ListSecretReferencesResponse
func (client *Client) ListSecretReferences(secretId *string, request *ListSecretReferencesRequest) (_result *ListSecretReferencesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSecretReferencesResponse{}
	_body, _err := client.ListSecretReferencesWithOptions(secretId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists keys.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param request - ListSecretsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretsResponse
func (client *Client) ListSecretsWithOptions(request *ListSecretsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSecretsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayType) {
		query["gatewayType"] = request.GatewayType
	}

	if !dara.IsNil(request.NameLike) {
		query["nameLike"] = request.NameLike
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSecrets"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/secrets"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSecretsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists keys.
//
// Description:
//
// The operation supports creating multiple services.
//
// @param request - ListSecretsRequest
//
// @return ListSecretsResponse
func (client *Client) ListSecrets(request *ListSecretsRequest) (_result *ListSecretsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSecretsResponse{}
	_body, _err := client.ListSecretsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of services.
//
// @param request - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(request *ListServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayId) {
		query["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SourceTypes) {
		query["sourceTypes"] = request.SourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServices"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of services.
//
// @param request - ListServicesRequest
//
// @return ListServicesResponse
func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of certificates.
//
// @param request - ListSslCertsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSslCertsResponse
func (client *Client) ListSslCertsWithOptions(request *ListSslCertsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSslCertsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CertNameLike) {
		query["certNameLike"] = request.CertNameLike
	}

	if !dara.IsNil(request.DomainName) {
		query["domainName"] = request.DomainName
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSslCerts"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ssl/certs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSslCertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of certificates.
//
// @param request - ListSslCertsRequest
//
// @return ListSslCertsResponse
func (client *Client) ListSslCerts(request *ListSslCertsRequest) (_result *ListSslCertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSslCertsResponse{}
	_body, _err := client.ListSslCertsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of resource labels.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceId) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, dara.String("ResourceId"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceIdShrink) {
		query["ResourceId"] = request.ResourceIdShrink
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/tags"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Retrieves the list of resource labels.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the zones available for a cloud-native API gateway in a specific region.
//
// @param request - ListZonesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListZonesResponse
func (client *Client) ListZonesWithOptions(request *ListZonesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListZonesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GatewayEdition) {
		query["gatewayEdition"] = request.GatewayEdition
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListZones"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/zones"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the zones available for a cloud-native API gateway in a specific region.
//
// @param request - ListZonesRequest
//
// @return ListZonesResponse
func (client *Client) ListZones(request *ListZonesRequest) (_result *ListZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListZonesResponse{}
	_body, _err := client.ListZonesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of consumer authorization rules.
//
// @param request - QueryConsumerAuthorizationRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryConsumerAuthorizationRulesResponse
func (client *Client) QueryConsumerAuthorizationRulesWithOptions(request *QueryConsumerAuthorizationRulesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryConsumerAuthorizationRulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiNameLike) {
		query["apiNameLike"] = request.ApiNameLike
	}

	if !dara.IsNil(request.ConsumerGroupId) {
		query["consumerGroupId"] = request.ConsumerGroupId
	}

	if !dara.IsNil(request.ConsumerId) {
		query["consumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.ConsumerNameLike) {
		query["consumerNameLike"] = request.ConsumerNameLike
	}

	if !dara.IsNil(request.EnvironmentId) {
		query["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GroupByApi) {
		query["groupByApi"] = request.GroupByApi
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentResourceId) {
		query["parentResourceId"] = request.ParentResourceId
	}

	if !dara.IsNil(request.PrincipalType) {
		query["principalType"] = request.PrincipalType
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ResourceTypes) {
		query["resourceTypes"] = request.ResourceTypes
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryConsumerAuthorizationRules"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/authorization-rules"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryConsumerAuthorizationRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of consumer authorization rules.
//
// @param request - QueryConsumerAuthorizationRulesRequest
//
// @return QueryConsumerAuthorizationRulesResponse
func (client *Client) QueryConsumerAuthorizationRules(request *QueryConsumerAuthorizationRulesRequest) (_result *QueryConsumerAuthorizationRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryConsumerAuthorizationRulesResponse{}
	_body, _err := client.QueryConsumerAuthorizationRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an API consumer authorization rule.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveConsumerAuthorizationRuleResponse
func (client *Client) RemoveConsumerAuthorizationRuleWithOptions(consumerAuthorizationRuleId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RemoveConsumerAuthorizationRuleResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/authorization-rules/" + dara.PercentEncode(dara.StringValue(consumerAuthorizationRuleId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API consumer authorization rule.
//
// @return RemoveConsumerAuthorizationRuleResponse
func (client *Client) RemoveConsumerAuthorizationRule(consumerAuthorizationRuleId *string) (_result *RemoveConsumerAuthorizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveConsumerAuthorizationRuleResponse{}
	_body, _err := client.RemoveConsumerAuthorizationRuleWithOptions(consumerAuthorizationRuleId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets a quota throttling rule on a gateway.
//
// Description:
//
// Resets a quota throttling rule on a gateway. This operation takes effect only on AI gateways running version 2.1.19 or later. Resetting a rule clears the historical usage of consumers associated with the rule.
//
// >
//
// >  Recommended call sequence:
//
// > - 1. Perform a dry run to check for rule conflicts.
//
// > - - Set dryRun to true.
//
// > - - The response contains a conflict preview with conflictHash.
//
// > - 2. Submit the request after confirmation.
//
// > - - No conflicts: Set dryRun to false and overwrite to false.
//
// > - - Conflicts exist and you confirm the overwrite: Set dryRun to false, overwrite to true, and conflictHash to the value returned in the previous step.
//
// @param request - ResetGatewayQuotaRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetGatewayQuotaRuleResponse
func (client *Client) ResetGatewayQuotaRuleWithOptions(gatewayId *string, ruleId *string, request *ResetGatewayQuotaRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetGatewayQuotaRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConflictHash) {
		body["conflictHash"] = request.ConflictHash
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Overwrite) {
		body["overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.PeriodMultiplier) {
		body["periodMultiplier"] = request.PeriodMultiplier
	}

	if !dara.IsNil(request.PeriodType) {
		body["periodType"] = request.PeriodType
	}

	if !dara.IsNil(request.QuotaLimit) {
		body["quotaLimit"] = request.QuotaLimit
	}

	if !dara.IsNil(request.Timezone) {
		body["timezone"] = request.Timezone
	}

	if !dara.IsNil(request.WindowAlignment) {
		body["windowAlignment"] = request.WindowAlignment
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetGatewayQuotaRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/quota-rules/" + dara.PercentEncode(dara.StringValue(ruleId)) + "/reset"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetGatewayQuotaRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets a quota throttling rule on a gateway.
//
// Description:
//
// Resets a quota throttling rule on a gateway. This operation takes effect only on AI gateways running version 2.1.19 or later. Resetting a rule clears the historical usage of consumers associated with the rule.
//
// >
//
// >  Recommended call sequence:
//
// > - 1. Perform a dry run to check for rule conflicts.
//
// > - - Set dryRun to true.
//
// > - - The response contains a conflict preview with conflictHash.
//
// > - 2. Submit the request after confirmation.
//
// > - - No conflicts: Set dryRun to false and overwrite to false.
//
// > - - Conflicts exist and you confirm the overwrite: Set dryRun to false, overwrite to true, and conflictHash to the value returned in the previous step.
//
// @param request - ResetGatewayQuotaRuleRequest
//
// @return ResetGatewayQuotaRuleResponse
func (client *Client) ResetGatewayQuotaRule(gatewayId *string, ruleId *string, request *ResetGatewayQuotaRuleRequest) (_result *ResetGatewayQuotaRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetGatewayQuotaRuleResponse{}
	_body, _err := client.ResetGatewayQuotaRuleWithOptions(gatewayId, ruleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a gateway.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartGatewayResponse
func (client *Client) RestartGatewayWithOptions(gatewayId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RestartGatewayResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/restart"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a gateway.
//
// @return RestartGatewayResponse
func (client *Client) RestartGateway(gatewayId *string) (_result *RestartGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestartGatewayResponse{}
	_body, _err := client.RestartGatewayWithOptions(gatewayId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Syncs an external MCP server.
//
// @param request - SyncMCPServersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncMCPServersResponse
func (client *Client) SyncMCPServersWithOptions(request *SyncMCPServersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncMCPServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DomainIds) {
		body["domainIds"] = request.DomainIds
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.NacosMcpServers) {
		body["nacosMcpServers"] = request.NacosMcpServers
	}

	if !dara.IsNil(request.Namespace) {
		body["namespace"] = request.Namespace
	}

	if !dara.IsNil(request.SourceId) {
		body["sourceId"] = request.SourceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncMCPServers"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/mcp-servers/sync-mcp-server"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncMCPServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Syncs an external MCP server.
//
// @param request - SyncMCPServersRequest
//
// @return SyncMCPServersResponse
func (client *Client) SyncMCPServers(request *SyncMCPServersRequest) (_result *SyncMCPServersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncMCPServersResponse{}
	_body, _err := client.SyncMCPServersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds labels to resources.
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
	if !dara.IsNil(request.ResourceId) {
		body["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		body["tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/tags"),
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
// Adds labels to resources.
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
// Cancels the publication of an MCP server.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnDeployMcpServerResponse
func (client *Client) UnDeployMcpServerWithOptions(mcpServerId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UnDeployMcpServerResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnDeployMcpServer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId)) + "/undeploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UnDeployMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the publication of an MCP server.
//
// @return UnDeployMcpServerResponse
func (client *Client) UnDeployMcpServer(mcpServerId *string) (_result *UnDeployMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnDeployMcpServerResponse{}
	_body, _err := client.UnDeployMcpServerWithOptions(mcpServerId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the deployment of an HTTP API.
//
// @param request - UndeployHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UndeployHttpApiResponse
func (client *Client) UndeployHttpApiWithOptions(httpApiId *string, request *UndeployHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UndeployHttpApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.OperationId) {
		body["operationId"] = request.OperationId
	}

	if !dara.IsNil(request.RouteId) {
		body["routeId"] = request.RouteId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UndeployHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/undeploy"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UndeployHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the deployment of an HTTP API.
//
// @param request - UndeployHttpApiRequest
//
// @return UndeployHttpApiResponse
func (client *Client) UndeployHttpApi(httpApiId *string, request *UndeployHttpApiRequest) (_result *UndeployHttpApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UndeployHttpApiResponse{}
	_body, _err := client.UndeployHttpApiWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uninstalls a plugin.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UninstallPluginResponse
func (client *Client) UninstallPluginWithOptions(pluginId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UninstallPluginResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("UninstallPlugin"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugins/" + dara.PercentEncode(dara.StringValue(pluginId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UninstallPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uninstalls a plugin.
//
// @return UninstallPluginResponse
func (client *Client) UninstallPlugin(pluginId *string) (_result *UninstallPluginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UninstallPluginResponse{}
	_body, _err := client.UninstallPluginWithOptions(pluginId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from resources.
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UntagResourcesShrinkRequest{}
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
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/tags"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新AI模型卡片
//
// @param request - UpdateAiModelCardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAiModelCardResponse
func (client *Client) UpdateAiModelCardWithOptions(modelCardId *string, request *UpdateAiModelCardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAiModelCardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AvailablePaths) {
		body["availablePaths"] = request.AvailablePaths
	}

	if !dara.IsNil(request.Credit) {
		body["credit"] = request.Credit
	}

	if !dara.IsNil(request.Features) {
		body["features"] = request.Features
	}

	if !dara.IsNil(request.Meta) {
		body["meta"] = request.Meta
	}

	if !dara.IsNil(request.ModelName) {
		body["modelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelProvider) {
		body["modelProvider"] = request.ModelProvider
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAiModelCard"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-cards/" + dara.PercentEncode(dara.StringValue(modelCardId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAiModelCardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AI模型卡片
//
// @param request - UpdateAiModelCardRequest
//
// @return UpdateAiModelCardResponse
func (client *Client) UpdateAiModelCard(modelCardId *string, request *UpdateAiModelCardRequest) (_result *UpdateAiModelCardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAiModelCardResponse{}
	_body, _err := client.UpdateAiModelCardWithOptions(modelCardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新AI模型供应商
//
// @param request - UpdateAiModelProviderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAiModelProviderResponse
func (client *Client) UpdateAiModelProviderWithOptions(modelProviderId *string, request *UpdateAiModelProviderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAiModelProviderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.ServiceIds) {
		body["serviceIds"] = request.ServiceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAiModelProvider"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ai-model-providers/" + dara.PercentEncode(dara.StringValue(modelProviderId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAiModelProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AI模型供应商
//
// @param request - UpdateAiModelProviderRequest
//
// @return UpdateAiModelProviderResponse
func (client *Client) UpdateAiModelProvider(modelProviderId *string, request *UpdateAiModelProviderRequest) (_result *UpdateAiModelProviderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAiModelProviderResponse{}
	_body, _err := client.UpdateAiModelProviderWithOptions(modelProviderId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates and mounts a policy.
//
// @param request - UpdateAndAttachPolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAndAttachPolicyResponse
func (client *Client) UpdateAndAttachPolicyWithOptions(policyId *string, request *UpdateAndAttachPolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAndAttachPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceIds) {
		body["attachResourceIds"] = request.AttachResourceIds
	}

	if !dara.IsNil(request.AttachResourceType) {
		body["attachResourceType"] = request.AttachResourceType
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.GatewayId) {
		body["gatewayId"] = request.GatewayId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAndAttachPolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAndAttachPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates and mounts a policy.
//
// @param request - UpdateAndAttachPolicyRequest
//
// @return UpdateAndAttachPolicyResponse
func (client *Client) UpdateAndAttachPolicy(policyId *string, request *UpdateAndAttachPolicyRequest) (_result *UpdateAndAttachPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAndAttachPolicyResponse{}
	_body, _err := client.UpdateAndAttachPolicyWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an API consumer.
//
// @param request - UpdateConsumerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerResponse
func (client *Client) UpdateConsumerWithOptions(consumerId *string, request *UpdateConsumerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConsumerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AkSkIdentityConfigs) {
		body["akSkIdentityConfigs"] = request.AkSkIdentityConfigs
	}

	if !dara.IsNil(request.ApikeyIdentityConfig) {
		body["apikeyIdentityConfig"] = request.ApikeyIdentityConfig
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.JwtIdentityConfig) {
		body["jwtIdentityConfig"] = request.JwtIdentityConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConsumer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an API consumer.
//
// @param request - UpdateConsumerRequest
//
// @return UpdateConsumerResponse
func (client *Client) UpdateConsumer(consumerId *string, request *UpdateConsumerRequest) (_result *UpdateConsumerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConsumerResponse{}
	_body, _err := client.UpdateConsumerWithOptions(consumerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a consumer authorization rule.
//
// @param request - UpdateConsumerAuthorizationRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerAuthorizationRuleResponse
func (client *Client) UpdateConsumerAuthorizationRuleWithOptions(consumerId *string, consumerAuthorizationRuleId *string, request *UpdateConsumerAuthorizationRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConsumerAuthorizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationResourceInfos) {
		body["authorizationResourceInfos"] = request.AuthorizationResourceInfos
	}

	if !dara.IsNil(request.ExpireMode) {
		body["expireMode"] = request.ExpireMode
	}

	if !dara.IsNil(request.ExpireTimestamp) {
		body["expireTimestamp"] = request.ExpireTimestamp
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConsumerAuthorizationRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumers/" + dara.PercentEncode(dara.StringValue(consumerId)) + "/authorization-rules/" + dara.PercentEncode(dara.StringValue(consumerAuthorizationRuleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConsumerAuthorizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a consumer authorization rule.
//
// @param request - UpdateConsumerAuthorizationRuleRequest
//
// @return UpdateConsumerAuthorizationRuleResponse
func (client *Client) UpdateConsumerAuthorizationRule(consumerId *string, consumerAuthorizationRuleId *string, request *UpdateConsumerAuthorizationRuleRequest) (_result *UpdateConsumerAuthorizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConsumerAuthorizationRuleResponse{}
	_body, _err := client.UpdateConsumerAuthorizationRuleWithOptions(consumerId, consumerAuthorizationRuleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新消费者组
//
// @param request - UpdateConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerGroupResponse
func (client *Client) UpdateConsumerGroupWithOptions(consumerGroupId *string, request *UpdateConsumerGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConsumerGroup"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/consumer-groups/" + dara.PercentEncode(dara.StringValue(consumerGroupId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新消费者组
//
// @param request - UpdateConsumerGroupRequest
//
// @return UpdateConsumerGroupResponse
func (client *Client) UpdateConsumerGroup(consumerGroupId *string, request *UpdateConsumerGroupRequest) (_result *UpdateConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.UpdateConsumerGroupWithOptions(consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a domain name.
//
// Description:
//
// Only sources of the **Container Service*	- type allow you to update the listener Ingress configuration.
//
// @param request - UpdateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomainWithOptions(domainId *string, request *UpdateDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDomainResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CaCertIdentifier) {
		body["caCertIdentifier"] = request.CaCertIdentifier
	}

	if !dara.IsNil(request.CertIdentifier) {
		body["certIdentifier"] = request.CertIdentifier
	}

	if !dara.IsNil(request.ClientCACert) {
		body["clientCACert"] = request.ClientCACert
	}

	if !dara.IsNil(request.DomainScope) {
		body["domainScope"] = request.DomainScope
	}

	if !dara.IsNil(request.ForceHttps) {
		body["forceHttps"] = request.ForceHttps
	}

	if !dara.IsNil(request.Http2Option) {
		body["http2Option"] = request.Http2Option
	}

	if !dara.IsNil(request.MTLSEnabled) {
		body["mTLSEnabled"] = request.MTLSEnabled
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.TlsCipherSuitesConfig) {
		body["tlsCipherSuitesConfig"] = request.TlsCipherSuitesConfig
	}

	if !dara.IsNil(request.TlsMax) {
		body["tlsMax"] = request.TlsMax
	}

	if !dara.IsNil(request.TlsMin) {
		body["tlsMin"] = request.TlsMin
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomain"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/domains/" + dara.PercentEncode(dara.StringValue(domainId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a domain name.
//
// Description:
//
// Only sources of the **Container Service*	- type allow you to update the listener Ingress configuration.
//
// @param request - UpdateDomainRequest
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomain(domainId *string, request *UpdateDomainRequest) (_result *UpdateDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDomainResponse{}
	_body, _err := client.UpdateDomainWithOptions(domainId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateEnvironment is deprecated
//
// Summary:
//
// Updates an environment.
//
// @param request - UpdateEnvironmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEnvironmentResponse
func (client *Client) UpdateEnvironmentWithOptions(environmentId *string, request *UpdateEnvironmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEnvironmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["alias"] = request.Alias
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEnvironment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/environments/" + dara.PercentEncode(dara.StringValue(environmentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateEnvironment is deprecated
//
// Summary:
//
// Updates an environment.
//
// @param request - UpdateEnvironmentRequest
//
// @return UpdateEnvironmentResponse
// Deprecated
func (client *Client) UpdateEnvironment(environmentId *string, request *UpdateEnvironmentRequest) (_result *UpdateEnvironmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEnvironmentResponse{}
	_body, _err := client.UpdateEnvironmentWithOptions(environmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the attribute parameters of a gateway.
//
// @param request - UpdateGatewayFeatureRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayFeatureResponse
func (client *Client) UpdateGatewayFeatureWithOptions(gatewayId *string, name *string, request *UpdateGatewayFeatureRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGatewayFeatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Value) {
		body["value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayFeature"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/gateway-features/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayFeatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the attribute parameters of a gateway.
//
// @param request - UpdateGatewayFeatureRequest
//
// @return UpdateGatewayFeatureResponse
func (client *Client) UpdateGatewayFeature(gatewayId *string, name *string, request *UpdateGatewayFeatureRequest) (_result *UpdateGatewayFeatureResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayFeatureResponse{}
	_body, _err := client.UpdateGatewayFeatureWithOptions(gatewayId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateGatewayName is deprecated
//
// Summary:
//
// Modifies the name of a gateway.
//
// @param request - UpdateGatewayNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayNameResponse
func (client *Client) UpdateGatewayNameWithOptions(gatewayId *string, request *UpdateGatewayNameRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGatewayNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayName"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/name"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateGatewayName is deprecated
//
// Summary:
//
// Modifies the name of a gateway.
//
// @param request - UpdateGatewayNameRequest
//
// @return UpdateGatewayNameResponse
// Deprecated
func (client *Client) UpdateGatewayName(gatewayId *string, request *UpdateGatewayNameRequest) (_result *UpdateGatewayNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayNameResponse{}
	_body, _err := client.UpdateGatewayNameWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits a quota throttling rule on a gateway.
//
// Description:
//
// Edits a quota rule on a gateway. This operation takes effect only on AI gateways with a version later than 2.1.19. Editing a rule preserves the historical usage of consumers on the rule.
//
// >  Recommended call logic:
//
// > - Step 1: Perform a dry run to check for rule conflicts.
//
// > - - Set dryRun to true.
//
// > - - The response contains a conflict preview with conflictHash.
//
// > - Step 2: Submit the request after confirmation.
//
// > - - No conflicts: Set dryRun to false and overwrite to false.
//
// > - - Conflicts exist and you confirm the overwrite: Set dryRun to false, overwrite to true, and conflictHash to the value returned in the previous step.
//
// @param request - UpdateGatewayQuotaRuleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayQuotaRuleResponse
func (client *Client) UpdateGatewayQuotaRuleWithOptions(gatewayId *string, ruleId *string, request *UpdateGatewayQuotaRuleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGatewayQuotaRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddIds) {
		body["addIds"] = request.AddIds
	}

	if !dara.IsNil(request.ConflictHash) {
		body["conflictHash"] = request.ConflictHash
	}

	if !dara.IsNil(request.ConsumerGroupIds) {
		body["consumerGroupIds"] = request.ConsumerGroupIds
	}

	if !dara.IsNil(request.DryRun) {
		body["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Overwrite) {
		body["overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.QuotaLimit) {
		body["quotaLimit"] = request.QuotaLimit
	}

	if !dara.IsNil(request.RemoveIds) {
		body["removeIds"] = request.RemoveIds
	}

	if !dara.IsNil(request.RuleName) {
		body["ruleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayQuotaRule"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/quota-rules/" + dara.PercentEncode(dara.StringValue(ruleId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayQuotaRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a quota throttling rule on a gateway.
//
// Description:
//
// Edits a quota rule on a gateway. This operation takes effect only on AI gateways with a version later than 2.1.19. Editing a rule preserves the historical usage of consumers on the rule.
//
// >  Recommended call logic:
//
// > - Step 1: Perform a dry run to check for rule conflicts.
//
// > - - Set dryRun to true.
//
// > - - The response contains a conflict preview with conflictHash.
//
// > - Step 2: Submit the request after confirmation.
//
// > - - No conflicts: Set dryRun to false and overwrite to false.
//
// > - - Conflicts exist and you confirm the overwrite: Set dryRun to false, overwrite to true, and conflictHash to the value returned in the previous step.
//
// @param request - UpdateGatewayQuotaRuleRequest
//
// @return UpdateGatewayQuotaRuleResponse
func (client *Client) UpdateGatewayQuotaRule(gatewayId *string, ruleId *string, request *UpdateGatewayQuotaRuleRequest) (_result *UpdateGatewayQuotaRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayQuotaRuleResponse{}
	_body, _err := client.UpdateGatewayQuotaRuleWithOptions(gatewayId, ruleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables a quota throttling rule for a gateway.
//
// Description:
//
// Enables or disables a quota rule on a gateway. This operation takes effect only for AI gateways with a version later than 2.1.19.
//
// @param request - UpdateGatewayQuotaRuleStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGatewayQuotaRuleStatusResponse
func (client *Client) UpdateGatewayQuotaRuleStatusWithOptions(gatewayId *string, ruleId *string, request *UpdateGatewayQuotaRuleStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateGatewayQuotaRuleStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClearHistory) {
		body["clearHistory"] = request.ClearHistory
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateGatewayQuotaRuleStatus"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/quota-rules/" + dara.PercentEncode(dara.StringValue(ruleId)) + "/status"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateGatewayQuotaRuleStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables a quota throttling rule for a gateway.
//
// Description:
//
// Enables or disables a quota rule on a gateway. This operation takes effect only for AI gateways with a version later than 2.1.19.
//
// @param request - UpdateGatewayQuotaRuleStatusRequest
//
// @return UpdateGatewayQuotaRuleStatusResponse
func (client *Client) UpdateGatewayQuotaRuleStatus(gatewayId *string, ruleId *string, request *UpdateGatewayQuotaRuleStatusRequest) (_result *UpdateGatewayQuotaRuleStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGatewayQuotaRuleStatusResponse{}
	_body, _err := client.UpdateGatewayQuotaRuleStatusWithOptions(gatewayId, ruleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an HTTP API.
//
// @param request - UpdateHttpApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiResponse
func (client *Client) UpdateHttpApiWithOptions(httpApiId *string, request *UpdateHttpApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHttpApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentProtocols) {
		body["agentProtocols"] = request.AgentProtocols
	}

	if !dara.IsNil(request.AiProtocols) {
		body["aiProtocols"] = request.AiProtocols
	}

	if !dara.IsNil(request.AuthConfig) {
		body["authConfig"] = request.AuthConfig
	}

	if !dara.IsNil(request.BasePath) {
		body["basePath"] = request.BasePath
	}

	if !dara.IsNil(request.DeployConfigs) {
		body["deployConfigs"] = request.DeployConfigs
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnableAuth) {
		body["enableAuth"] = request.EnableAuth
	}

	if !dara.IsNil(request.FirstByteTimeout) {
		body["firstByteTimeout"] = request.FirstByteTimeout
	}

	if !dara.IsNil(request.IngressConfig) {
		body["ingressConfig"] = request.IngressConfig
	}

	if !dara.IsNil(request.OnlyChangeConfig) {
		body["onlyChangeConfig"] = request.OnlyChangeConfig
	}

	if !dara.IsNil(request.Protocols) {
		body["protocols"] = request.Protocols
	}

	if !dara.IsNil(request.RemoveBasePathOnForward) {
		body["removeBasePathOnForward"] = request.RemoveBasePathOnForward
	}

	if !dara.IsNil(request.VersionConfig) {
		body["versionConfig"] = request.VersionConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpApi"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an HTTP API.
//
// @param request - UpdateHttpApiRequest
//
// @return UpdateHttpApiResponse
func (client *Client) UpdateHttpApi(httpApiId *string, request *UpdateHttpApiRequest) (_result *UpdateHttpApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHttpApiResponse{}
	_body, _err := client.UpdateHttpApiWithOptions(httpApiId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an API operation.
//
// @param request - UpdateHttpApiOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiOperationResponse
func (client *Client) UpdateHttpApiOperationWithOptions(httpApiId *string, operationId *string, request *UpdateHttpApiOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHttpApiOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Operation) {
		body["operation"] = request.Operation
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpApiOperation"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/operations/" + dara.PercentEncode(dara.StringValue(operationId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpApiOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an API operation.
//
// @param request - UpdateHttpApiOperationRequest
//
// @return UpdateHttpApiOperationResponse
func (client *Client) UpdateHttpApiOperation(httpApiId *string, operationId *string, request *UpdateHttpApiOperationRequest) (_result *UpdateHttpApiOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHttpApiOperationResponse{}
	_body, _err := client.UpdateHttpApiOperationWithOptions(httpApiId, operationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a route of an HttpApi.
//
// @param request - UpdateHttpApiRouteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHttpApiRouteResponse
func (client *Client) UpdateHttpApiRouteWithOptions(httpApiId *string, routeId *string, request *UpdateHttpApiRouteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateHttpApiRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BackendConfig) {
		body["backendConfig"] = request.BackendConfig
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DomainIds) {
		body["domainIds"] = request.DomainIds
	}

	if !dara.IsNil(request.EnvironmentId) {
		body["environmentId"] = request.EnvironmentId
	}

	if !dara.IsNil(request.Match) {
		body["match"] = request.Match
	}

	if !dara.IsNil(request.McpRouteConfig) {
		body["mcpRouteConfig"] = request.McpRouteConfig
	}

	if !dara.IsNil(request.PolicyConfigs) {
		body["policyConfigs"] = request.PolicyConfigs
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHttpApiRoute"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/http-apis/" + dara.PercentEncode(dara.StringValue(httpApiId)) + "/routes/" + dara.PercentEncode(dara.StringValue(routeId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHttpApiRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a route of an HttpApi.
//
// @param request - UpdateHttpApiRouteRequest
//
// @return UpdateHttpApiRouteResponse
func (client *Client) UpdateHttpApiRoute(httpApiId *string, routeId *string, request *UpdateHttpApiRouteRequest) (_result *UpdateHttpApiRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateHttpApiRouteResponse{}
	_body, _err := client.UpdateHttpApiRouteWithOptions(httpApiId, routeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an MCP server.
//
// Description:
//
// Only sources of the **Container Service*	- type can update the Ingress listener configuration.
//
// @param request - UpdateMcpServerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMcpServerResponse
func (client *Client) UpdateMcpServerWithOptions(mcpServerId *string, request *UpdateMcpServerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssembledSources) {
		body["assembledSources"] = request.AssembledSources
	}

	if !dara.IsNil(request.BackendConfig) {
		body["backendConfig"] = request.BackendConfig
	}

	if !dara.IsNil(request.CreateFromType) {
		body["createFromType"] = request.CreateFromType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DomainIds) {
		body["domainIds"] = request.DomainIds
	}

	if !dara.IsNil(request.ExposedUriPath) {
		body["exposedUriPath"] = request.ExposedUriPath
	}

	if !dara.IsNil(request.GrayMcpServerConfigs) {
		body["grayMcpServerConfigs"] = request.GrayMcpServerConfigs
	}

	if !dara.IsNil(request.Match) {
		body["match"] = request.Match
	}

	if !dara.IsNil(request.McpServerConfig) {
		body["mcpServerConfig"] = request.McpServerConfig
	}

	if !dara.IsNil(request.McpStatisticsEnable) {
		body["mcpStatisticsEnable"] = request.McpStatisticsEnable
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMcpServer"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/mcp-servers/" + dara.PercentEncode(dara.StringValue(mcpServerId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMcpServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an MCP server.
//
// Description:
//
// Only sources of the **Container Service*	- type can update the Ingress listener configuration.
//
// @param request - UpdateMcpServerRequest
//
// @return UpdateMcpServerResponse
func (client *Client) UpdateMcpServer(mcpServerId *string, request *UpdateMcpServerRequest) (_result *UpdateMcpServerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMcpServerResponse{}
	_body, _err := client.UpdateMcpServerWithOptions(mcpServerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a plugin mount.
//
// @param request - UpdatePluginAttachmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePluginAttachmentResponse
func (client *Client) UpdatePluginAttachmentWithOptions(pluginAttachmentId *string, request *UpdatePluginAttachmentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePluginAttachmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AttachResourceIds) {
		body["attachResourceIds"] = request.AttachResourceIds
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.PluginConfig) {
		body["pluginConfig"] = request.PluginConfig
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePluginAttachment"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/plugin-attachments/" + dara.PercentEncode(dara.StringValue(pluginAttachmentId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePluginAttachmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a plugin mount.
//
// @param request - UpdatePluginAttachmentRequest
//
// @return UpdatePluginAttachmentResponse
func (client *Client) UpdatePluginAttachment(pluginAttachmentId *string, request *UpdatePluginAttachmentRequest) (_result *UpdatePluginAttachmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePluginAttachmentResponse{}
	_body, _err := client.UpdatePluginAttachmentWithOptions(pluginAttachmentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a policy.
//
// @param request - UpdatePolicyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicyWithOptions(policyId *string, request *UpdatePolicyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePolicy"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/policies/" + dara.PercentEncode(dara.StringValue(policyId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a policy.
//
// @param request - UpdatePolicyRequest
//
// @return UpdatePolicyResponse
func (client *Client) UpdatePolicy(policyId *string, request *UpdatePolicyRequest) (_result *UpdatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePolicyResponse{}
	_body, _err := client.UpdatePolicyWithOptions(policyId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a key pair.
//
// Description:
//
// Only sources of the **container service*	- type allow you to update the configuration for listening to Ingress.
//
// @param request - UpdateSecretRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretResponse
func (client *Client) UpdateSecretWithOptions(secretId *string, request *UpdateSecretRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SecretData) {
		body["secretData"] = request.SecretData
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSecret"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/secrets/" + dara.PercentEncode(dara.StringValue(secretId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a key pair.
//
// Description:
//
// Only sources of the **container service*	- type allow you to update the configuration for listening to Ingress.
//
// @param request - UpdateSecretRequest
//
// @return UpdateSecretResponse
func (client *Client) UpdateSecret(secretId *string, request *UpdateSecretRequest) (_result *UpdateSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSecretResponse{}
	_body, _err := client.UpdateSecretWithOptions(secretId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a service. You can update the health check configuration, DNS domain name, and address configuration of fixed addresses for the service.
//
// @param request - UpdateServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceResponse
func (client *Client) UpdateServiceWithOptions(serviceId *string, request *UpdateServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Addresses) {
		body["addresses"] = request.Addresses
	}

	if !dara.IsNil(request.AgentServiceConfig) {
		body["agentServiceConfig"] = request.AgentServiceConfig
	}

	if !dara.IsNil(request.AiServiceConfig) {
		body["aiServiceConfig"] = request.AiServiceConfig
	}

	if !dara.IsNil(request.DnsServers) {
		body["dnsServers"] = request.DnsServers
	}

	if !dara.IsNil(request.HealthCheckConfig) {
		body["healthCheckConfig"] = request.HealthCheckConfig
	}

	if !dara.IsNil(request.HealthyPanicThreshold) {
		body["healthyPanicThreshold"] = request.HealthyPanicThreshold
	}

	if !dara.IsNil(request.ModelProviderId) {
		body["modelProviderId"] = request.ModelProviderId
	}

	if !dara.IsNil(request.OutlierDetectionConfig) {
		body["outlierDetectionConfig"] = request.OutlierDetectionConfig
	}

	if !dara.IsNil(request.Ports) {
		body["ports"] = request.Ports
	}

	if !dara.IsNil(request.Protocol) {
		body["protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateService"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a service. You can update the health check configuration, DNS domain name, and address configuration of fixed addresses for the service.
//
// @param request - UpdateServiceRequest
//
// @return UpdateServiceResponse
func (client *Client) UpdateService(serviceId *string, request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(serviceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a service version.
//
// @param request - UpdateServiceVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceVersionResponse
func (client *Client) UpdateServiceVersionWithOptions(serviceId *string, name *string, request *UpdateServiceVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateServiceVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Labels) {
		body["labels"] = request.Labels
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateServiceVersion"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/services/" + dara.PercentEncode(dara.StringValue(serviceId)) + "/versions/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a service version.
//
// @param request - UpdateServiceVersionRequest
//
// @return UpdateServiceVersionResponse
func (client *Client) UpdateServiceVersion(serviceId *string, name *string, request *UpdateServiceVersionRequest) (_result *UpdateServiceVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceVersionResponse{}
	_body, _err := client.UpdateServiceVersionWithOptions(serviceId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the gateway version.
//
// @param request - UpgradeGatewayRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeGatewayResponse
func (client *Client) UpgradeGatewayWithOptions(gatewayId *string, request *UpgradeGatewayRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpgradeGatewayResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeGateway"),
		Version:     dara.String("2024-03-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/gateways/" + dara.PercentEncode(dara.StringValue(gatewayId)) + "/upgrade"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeGatewayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the gateway version.
//
// @param request - UpgradeGatewayRequest
//
// @return UpgradeGatewayResponse
func (client *Client) UpgradeGateway(gatewayId *string, request *UpgradeGatewayRequest) (_result *UpgradeGatewayResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpgradeGatewayResponse{}
	_body, _err := client.UpgradeGatewayWithOptions(gatewayId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
