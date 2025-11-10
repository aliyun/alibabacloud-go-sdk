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
	client.Endpoint, _err = client.GetEndpoint(dara.String("rocketmq"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Add Disaster Recovery Plan Entry
//
// @param request - AddDisasterRecoveryItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDisasterRecoveryItemResponse
func (client *Client) AddDisasterRecoveryItemWithOptions(planId *string, request *AddDisasterRecoveryItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddDisasterRecoveryItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Topics) {
		body["topics"] = request.Topics
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDisasterRecoveryItem"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDisasterRecoveryItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add Disaster Recovery Plan Entry
//
// @param request - AddDisasterRecoveryItemRequest
//
// @return AddDisasterRecoveryItemResponse
func (client *Client) AddDisasterRecoveryItem(planId *string, request *AddDisasterRecoveryItemRequest) (_result *AddDisasterRecoveryItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddDisasterRecoveryItemResponse{}
	_body, _err := client.AddDisasterRecoveryItemWithOptions(planId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a ApsaraMQ for RocketMQ instance belongs.
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
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resourceGroup/change"),
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
// Changes the resource group to which a ApsaraMQ for RocketMQ instance belongs.
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
// Creates a consumer group.
//
// Description:
//
// The ID of the consumer group. The ID is globally unique and is used to identify a consumer group.
//
// The following limits are imposed on the ID:
//
//   - Character limit: The ID can contain letters, digits, underscores (_), hyphens (-), and percent signs (%).
//
//   - Length limit: The ID must be 1 to 60 characters in length.
//
// For more information about strings that are reserved for the system, see [Limits on parameters](https://help.aliyun.com/document_detail/440347.html).
//
// @param request - CreateConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, request *CreateConsumerGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConsumeRetryPolicy) {
		body["consumeRetryPolicy"] = request.ConsumeRetryPolicy
	}

	if !dara.IsNil(request.DeliveryOrderType) {
		body["deliveryOrderType"] = request.DeliveryOrderType
	}

	if !dara.IsNil(request.MaxReceiveTps) {
		body["maxReceiveTps"] = request.MaxReceiveTps
	}

	if !dara.IsNil(request.MessageModel) {
		body["messageModel"] = request.MessageModel
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.TopicName) {
		body["topicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerGroup"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId))),
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
// Creates a consumer group.
//
// Description:
//
// The ID of the consumer group. The ID is globally unique and is used to identify a consumer group.
//
// The following limits are imposed on the ID:
//
//   - Character limit: The ID can contain letters, digits, underscores (_), hyphens (-), and percent signs (%).
//
//   - Length limit: The ID must be 1 to 60 characters in length.
//
// For more information about strings that are reserved for the system, see [Limits on parameters](https://help.aliyun.com/document_detail/440347.html).
//
// @param request - CreateConsumerGroupRequest
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroup(instanceId *string, consumerGroupId *string, request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Disaster Recovery Plan
//
// @param request - CreateDisasterRecoveryPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDisasterRecoveryPlanResponse
func (client *Client) CreateDisasterRecoveryPlanWithOptions(request *CreateDisasterRecoveryPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDisasterRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoSyncCheckpoint) {
		body["autoSyncCheckpoint"] = request.AutoSyncCheckpoint
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	if !dara.IsNil(request.PlanDesc) {
		body["planDesc"] = request.PlanDesc
	}

	if !dara.IsNil(request.PlanName) {
		body["planName"] = request.PlanName
	}

	if !dara.IsNil(request.PlanType) {
		body["planType"] = request.PlanType
	}

	if !dara.IsNil(request.SyncCheckpointEnabled) {
		body["syncCheckpointEnabled"] = request.SyncCheckpointEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDisasterRecoveryPlan"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDisasterRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Disaster Recovery Plan
//
// @param request - CreateDisasterRecoveryPlanRequest
//
// @return CreateDisasterRecoveryPlanResponse
func (client *Client) CreateDisasterRecoveryPlan(request *CreateDisasterRecoveryPlanRequest) (_result *CreateDisasterRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDisasterRecoveryPlanResponse{}
	_body, _err := client.CreateDisasterRecoveryPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an ApsaraMQ for RocketMQ 5.x instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - CreateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
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
	if !dara.IsNil(request.AutoRenew) {
		body["autoRenew"] = request.AutoRenew
	}

	if !dara.IsNil(request.AutoRenewPeriod) {
		body["autoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !dara.IsNil(request.CommodityCode) {
		body["commodityCode"] = request.CommodityCode
	}

	if !dara.IsNil(request.InstanceName) {
		body["instanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.NetworkInfo) {
		body["networkInfo"] = request.NetworkInfo
	}

	if !dara.IsNil(request.PaymentType) {
		body["paymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.Period) {
		body["period"] = request.Period
	}

	if !dara.IsNil(request.PeriodUnit) {
		body["periodUnit"] = request.PeriodUnit
	}

	if !dara.IsNil(request.ProductInfo) {
		body["productInfo"] = request.ProductInfo
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SeriesCode) {
		body["seriesCode"] = request.SeriesCode
	}

	if !dara.IsNil(request.ServiceCode) {
		body["serviceCode"] = request.ServiceCode
	}

	if !dara.IsNil(request.SubSeriesCode) {
		body["subSeriesCode"] = request.SubSeriesCode
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
		Action:      dara.String("CreateInstance"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ApsaraMQ for RocketMQ 5.x instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - CreateInstanceRequest
//
// @return CreateInstanceResponse
func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an account that is used to access an instance.
//
// @param request - CreateInstanceAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceAccountResponse
func (client *Client) CreateInstanceAccountWithOptions(instanceId *string, request *CreateInstanceAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Password) {
		body["password"] = request.Password
	}

	if !dara.IsNil(request.Username) {
		body["username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceAccount"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/accounts"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an account that is used to access an instance.
//
// @param request - CreateInstanceAccountRequest
//
// @return CreateInstanceAccountResponse
func (client *Client) CreateInstanceAccount(instanceId *string, request *CreateInstanceAccountRequest) (_result *CreateInstanceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.CreateInstanceAccountWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) in a specific instance.
//
// @param request - CreateInstanceAclRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceAclResponse
func (client *Client) CreateInstanceAclWithOptions(instanceId *string, username *string, request *CreateInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Actions) {
		body["actions"] = request.Actions
	}

	if !dara.IsNil(request.Decision) {
		body["decision"] = request.Decision
	}

	if !dara.IsNil(request.IpWhitelists) {
		body["ipWhitelists"] = request.IpWhitelists
	}

	if !dara.IsNil(request.ResourceName) {
		body["resourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceAcl"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/acl/account/" + dara.PercentEncode(dara.StringValue(username))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL) in a specific instance.
//
// @param request - CreateInstanceAclRequest
//
// @return CreateInstanceAclResponse
func (client *Client) CreateInstanceAcl(instanceId *string, username *string, request *CreateInstanceAclRequest) (_result *CreateInstanceAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceAclResponse{}
	_body, _err := client.CreateInstanceAclWithOptions(instanceId, username, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an IP address whitelist.
//
// @param request - CreateInstanceIpWhitelistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceIpWhitelistResponse
func (client *Client) CreateInstanceIpWhitelistWithOptions(instanceId *string, request *CreateInstanceIpWhitelistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IpWhitelists) {
		body["ipWhitelists"] = request.IpWhitelists
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateInstanceIpWhitelist"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/ip/whitelist"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IP address whitelist.
//
// @param request - CreateInstanceIpWhitelistRequest
//
// @return CreateInstanceIpWhitelistResponse
func (client *Client) CreateInstanceIpWhitelist(instanceId *string, request *CreateInstanceIpWhitelistRequest) (_result *CreateInstanceIpWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInstanceIpWhitelistResponse{}
	_body, _err := client.CreateInstanceIpWhitelistWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Topic
//
// @param request - CreateTopicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTopicResponse
func (client *Client) CreateTopicWithOptions(instanceId *string, topicName *string, request *CreateTopicRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LiteTopicExpiration) {
		body["liteTopicExpiration"] = request.LiteTopicExpiration
	}

	if !dara.IsNil(request.MaxSendTps) {
		body["maxSendTps"] = request.MaxSendTps
	}

	if !dara.IsNil(request.MessageType) {
		body["messageType"] = request.MessageType
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTopic"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Topic
//
// @param request - CreateTopicRequest
//
// @return CreateTopicResponse
func (client *Client) CreateTopic(instanceId *string, topicName *string, request *CreateTopicRequest) (_result *CreateTopicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTopicResponse{}
	_body, _err := client.CreateTopicWithOptions(instanceId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// After you delete a consumer group, the consumer client associated with the consumer group cannot consume messages. Exercise caution when you call this operation.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerGroup"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId))),
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
// Deletes a specified consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// After you delete a consumer group, the consumer client associated with the consumer group cannot consume messages. Exercise caution when you call this operation.
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroup(instanceId *string, consumerGroupId *string) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(instanceId, consumerGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the subscriptions of a consumer group.
//
// @param request - DeleteConsumerGroupSubscriptionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupSubscriptionResponse
func (client *Client) DeleteConsumerGroupSubscriptionWithOptions(instanceId *string, consumerGroupId *string, request *DeleteConsumerGroupSubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupSubscriptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterExpression) {
		query["filterExpression"] = request.FilterExpression
	}

	if !dara.IsNil(request.FilterType) {
		query["filterType"] = request.FilterType
	}

	if !dara.IsNil(request.TopicName) {
		query["topicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerGroupSubscription"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/subscriptions"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConsumerGroupSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the subscriptions of a consumer group.
//
// @param request - DeleteConsumerGroupSubscriptionRequest
//
// @return DeleteConsumerGroupSubscriptionResponse
func (client *Client) DeleteConsumerGroupSubscription(instanceId *string, consumerGroupId *string, request *DeleteConsumerGroupSubscriptionRequest) (_result *DeleteConsumerGroupSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConsumerGroupSubscriptionResponse{}
	_body, _err := client.DeleteConsumerGroupSubscriptionWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除容灾计划条目
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDisasterRecoveryItemResponse
func (client *Client) DeleteDisasterRecoveryItemWithOptions(planId *string, itemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDisasterRecoveryItemResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDisasterRecoveryItem"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items/" + dara.PercentEncode(dara.StringValue(itemId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDisasterRecoveryItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除容灾计划条目
//
// @return DeleteDisasterRecoveryItemResponse
func (client *Client) DeleteDisasterRecoveryItem(planId *string, itemId *string) (_result *DeleteDisasterRecoveryItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDisasterRecoveryItemResponse{}
	_body, _err := client.DeleteDisasterRecoveryItemWithOptions(planId, itemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a global message backup plan.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDisasterRecoveryPlanResponse
func (client *Client) DeleteDisasterRecoveryPlanWithOptions(planId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDisasterRecoveryPlanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDisasterRecoveryPlan"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDisasterRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a global message backup plan.
//
// @return DeleteDisasterRecoveryPlanResponse
func (client *Client) DeleteDisasterRecoveryPlan(planId *string) (_result *DeleteDisasterRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDisasterRecoveryPlanResponse{}
	_body, _err := client.DeleteDisasterRecoveryPlanWithOptions(planId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - After an instance is deleted, the instance cannot be restored. Exercise caution when you call this operation.
//
//   - This operation is used to delete a pay-as-you-go instance. A subscription instance is automatically released after it expires. You do not need to manually delete a subscription instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
//   - After an instance is deleted, the instance cannot be restored. Exercise caution when you call this operation.
//
//   - This operation is used to delete a pay-as-you-go instance. A subscription instance is automatically released after it expires. You do not need to manually delete a subscription instance.
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(instanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete access control ACL user
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceAccountResponse
func (client *Client) DeleteInstanceAccountWithOptions(instanceId *string, username *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceAccountResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceAccount"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/accounts/" + dara.PercentEncode(dara.StringValue(username))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete access control ACL user
//
// @return DeleteInstanceAccountResponse
func (client *Client) DeleteInstanceAccount(instanceId *string, username *string) (_result *DeleteInstanceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceAccountResponse{}
	_body, _err := client.DeleteInstanceAccountWithOptions(instanceId, username, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the permissions of a specific account of an instance.
//
// @param request - DeleteInstanceAclRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceAclResponse
func (client *Client) DeleteInstanceAclWithOptions(instanceId *string, username *string, request *DeleteInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceName) {
		query["resourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceAcl"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/acl/account/" + dara.PercentEncode(dara.StringValue(username))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the permissions of a specific account of an instance.
//
// @param request - DeleteInstanceAclRequest
//
// @return DeleteInstanceAclResponse
func (client *Client) DeleteInstanceAcl(instanceId *string, username *string, request *DeleteInstanceAclRequest) (_result *DeleteInstanceAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceAclResponse{}
	_body, _err := client.DeleteInstanceAclWithOptions(instanceId, username, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific IP address whitelist from an instance.
//
// @param tmpReq - DeleteInstanceIpWhitelistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceIpWhitelistResponse
func (client *Client) DeleteInstanceIpWhitelistWithOptions(instanceId *string, tmpReq *DeleteInstanceIpWhitelistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteInstanceIpWhitelistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IpWhitelists) {
		request.IpWhitelistsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpWhitelists, dara.String("ipWhitelists"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IpWhitelist) {
		query["ipWhitelist"] = request.IpWhitelist
	}

	if !dara.IsNil(request.IpWhitelistsShrink) {
		query["ipWhitelists"] = request.IpWhitelistsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstanceIpWhitelist"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/ip/whitelist"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific IP address whitelist from an instance.
//
// @param request - DeleteInstanceIpWhitelistRequest
//
// @return DeleteInstanceIpWhitelistResponse
func (client *Client) DeleteInstanceIpWhitelist(instanceId *string, request *DeleteInstanceIpWhitelistRequest) (_result *DeleteInstanceIpWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteInstanceIpWhitelistResponse{}
	_body, _err := client.DeleteInstanceIpWhitelistWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified topic.
//
// Description:
//
// If you delete the topic, the publishing and subscription relationships that are established based on the topic are cleared. Exercise caution when you call this operation.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopicWithOptions(instanceId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTopic"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified topic.
//
// Description:
//
// If you delete the topic, the publishing and subscription relationships that are established based on the topic are cleared. Exercise caution when you call this operation.
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopic(instanceId *string, topicName *string) (_result *DeleteTopicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTopicResponse{}
	_body, _err := client.DeleteTopicWithOptions(instanceId, topicName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行迁移操作
//
// @param request - ExecuteMigrationOperationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteMigrationOperationResponse
func (client *Client) ExecuteMigrationOperationWithOptions(migrationId *string, stageType *string, operationId *string, request *ExecuteMigrationOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteMigrationOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OperationParam) {
		body["operationParam"] = request.OperationParam
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteMigrationOperation"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/migrations/" + dara.PercentEncode(dara.StringValue(migrationId)) + "/stages/" + dara.PercentEncode(dara.StringValue(stageType)) + "/operations/" + dara.PercentEncode(dara.StringValue(operationId)) + "/execute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteMigrationOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行迁移操作
//
// @param request - ExecuteMigrationOperationRequest
//
// @return ExecuteMigrationOperationResponse
func (client *Client) ExecuteMigrationOperation(migrationId *string, stageType *string, operationId *string, request *ExecuteMigrationOperationRequest) (_result *ExecuteMigrationOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteMigrationOperationResponse{}
	_body, _err := client.ExecuteMigrationOperationWithOptions(migrationId, stageType, operationId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 完成当前迁移阶段
//
// @param request - FinishMigrationStageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishMigrationStageResponse
func (client *Client) FinishMigrationStageWithOptions(migrationId *string, stageType *string, request *FinishMigrationStageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FinishMigrationStageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FinishMigrationStage"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/migrations/" + dara.PercentEncode(dara.StringValue(migrationId)) + "/stages/" + dara.PercentEncode(dara.StringValue(stageType)) + "/finish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FinishMigrationStageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 完成当前迁移阶段
//
// @param request - FinishMigrationStageRequest
//
// @return FinishMigrationStageResponse
func (client *Client) FinishMigrationStage(migrationId *string, stageType *string, request *FinishMigrationStageRequest) (_result *FinishMigrationStageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishMigrationStageResponse{}
	_body, _err := client.FinishMigrationStageWithOptions(migrationId, stageType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询topic可重置时间范围
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumeTimespanResponse
func (client *Client) GetConsumeTimespanWithOptions(instanceId *string, consumerGroupId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumeTimespanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumeTimespan"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/consumeTimespan/" + dara.PercentEncode(dara.StringValue(topicName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumeTimespanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询topic可重置时间范围
//
// @return GetConsumeTimespanResponse
func (client *Client) GetConsumeTimespan(instanceId *string, consumerGroupId *string, topicName *string) (_result *GetConsumeTimespanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumeTimespanResponse{}
	_body, _err := client.GetConsumeTimespanWithOptions(instanceId, consumerGroupId, topicName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerGroupResponse
func (client *Client) GetConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerGroupResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerGroup"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId))),
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
// Queries the details of a specified consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @return GetConsumerGroupResponse
func (client *Client) GetConsumerGroup(instanceId *string, consumerGroupId *string) (_result *GetConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerGroupResponse{}
	_body, _err := client.GetConsumerGroupWithOptions(instanceId, consumerGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Consumer Group Backlog Information
//
// @param request - GetConsumerGroupLagRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerGroupLagResponse
func (client *Client) GetConsumerGroupLagWithOptions(instanceId *string, consumerGroupId *string, request *GetConsumerGroupLagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerGroupLagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LiteTopicName) {
		query["liteTopicName"] = request.LiteTopicName
	}

	if !dara.IsNil(request.TopicName) {
		query["topicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerGroupLag"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/lag"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerGroupLagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Consumer Group Backlog Information
//
// @param request - GetConsumerGroupLagRequest
//
// @return GetConsumerGroupLagResponse
func (client *Client) GetConsumerGroupLag(instanceId *string, consumerGroupId *string, request *GetConsumerGroupLagRequest) (_result *GetConsumerGroupLagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerGroupLagResponse{}
	_body, _err := client.GetConsumerGroupLagWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the subscriptions of a consumer group.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerGroupSubscriptionResponse
func (client *Client) GetConsumerGroupSubscriptionWithOptions(instanceId *string, consumerGroupId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerGroupSubscriptionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerGroupSubscription"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/subscriptions/" + dara.PercentEncode(dara.StringValue(topicName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerGroupSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscriptions of a consumer group.
//
// @return GetConsumerGroupSubscriptionResponse
func (client *Client) GetConsumerGroupSubscription(instanceId *string, consumerGroupId *string, topicName *string) (_result *GetConsumerGroupSubscriptionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerGroupSubscriptionResponse{}
	_body, _err := client.GetConsumerGroupSubscriptionWithOptions(instanceId, consumerGroupId, topicName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the stack information about a consumer.
//
// @param request - GetConsumerStackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerStackResponse
func (client *Client) GetConsumerStackWithOptions(instanceId *string, consumerGroupId *string, request *GetConsumerStackRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerStackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerStack"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/stack"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the stack information about a consumer.
//
// @param request - GetConsumerStackRequest
//
// @return GetConsumerStackResponse
func (client *Client) GetConsumerStack(instanceId *string, consumerGroupId *string, request *GetConsumerStackRequest) (_result *GetConsumerStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerStackResponse{}
	_body, _err := client.GetConsumerStackWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询容灾计划条目详情
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDisasterRecoveryItemResponse
func (client *Client) GetDisasterRecoveryItemWithOptions(planId *string, itemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDisasterRecoveryItemResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDisasterRecoveryItem"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items/" + dara.PercentEncode(dara.StringValue(itemId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDisasterRecoveryItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询容灾计划条目详情
//
// @return GetDisasterRecoveryItemResponse
func (client *Client) GetDisasterRecoveryItem(planId *string, itemId *string) (_result *GetDisasterRecoveryItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDisasterRecoveryItemResponse{}
	_body, _err := client.GetDisasterRecoveryItemWithOptions(planId, itemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a Global Replicator task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDisasterRecoveryPlanResponse
func (client *Client) GetDisasterRecoveryPlanWithOptions(planId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDisasterRecoveryPlanResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDisasterRecoveryPlan"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDisasterRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a Global Replicator task.
//
// @return GetDisasterRecoveryPlanResponse
func (client *Client) GetDisasterRecoveryPlan(planId *string) (_result *GetDisasterRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDisasterRecoveryPlanResponse{}
	_body, _err := client.GetDisasterRecoveryPlanWithOptions(planId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detailed information about an instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detailed information about an instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @return GetInstanceResponse
func (client *Client) GetInstance(instanceId *string) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the account used to access a specific instance.
//
// @param request - GetInstanceAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceAccountResponse
func (client *Client) GetInstanceAccountWithOptions(instanceId *string, request *GetInstanceAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Username) {
		query["username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceAccount"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/account"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the account used to access a specific instance.
//
// @param request - GetInstanceAccountRequest
//
// @return GetInstanceAccountResponse
func (client *Client) GetInstanceAccount(instanceId *string, request *GetInstanceAccountRequest) (_result *GetInstanceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceAccountResponse{}
	_body, _err := client.GetInstanceAccountWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the access control list (ACL) of an instance.
//
// @param request - GetInstanceAclRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceAclResponse
func (client *Client) GetInstanceAclWithOptions(instanceId *string, username *string, request *GetInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceName) {
		query["resourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceAcl"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/acl/account/" + dara.PercentEncode(dara.StringValue(username))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the access control list (ACL) of an instance.
//
// @param request - GetInstanceAclRequest
//
// @return GetInstanceAclResponse
func (client *Client) GetInstanceAcl(instanceId *string, username *string, request *GetInstanceAclRequest) (_result *GetInstanceAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceAclResponse{}
	_body, _err := client.GetInstanceAclWithOptions(instanceId, username, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the IP address whitelist of an instance.
//
// @param tmpReq - GetInstanceIpWhitelistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceIpWhitelistResponse
func (client *Client) GetInstanceIpWhitelistWithOptions(instanceId *string, tmpReq *GetInstanceIpWhitelistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetInstanceIpWhitelistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IpWhitelists) {
		request.IpWhitelistsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpWhitelists, dara.String("ipWhitelists"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IpWhitelistsShrink) {
		query["ipWhitelists"] = request.IpWhitelistsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceIpWhitelist"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/ip/whitelists"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the IP address whitelist of an instance.
//
// @param request - GetInstanceIpWhitelistRequest
//
// @return GetInstanceIpWhitelistResponse
func (client *Client) GetInstanceIpWhitelist(instanceId *string, request *GetInstanceIpWhitelistRequest) (_result *GetInstanceIpWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceIpWhitelistResponse{}
	_body, _err := client.GetInstanceIpWhitelistWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a specific message.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageDetailResponse
func (client *Client) GetMessageDetailWithOptions(instanceId *string, topicName *string, messageId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMessageDetailResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMessageDetail"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName)) + "/messages/" + dara.PercentEncode(dara.StringValue(messageId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMessageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a specific message.
//
// @return GetMessageDetailResponse
func (client *Client) GetMessageDetail(instanceId *string, topicName *string, messageId *string) (_result *GetMessageDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMessageDetailResponse{}
	_body, _err := client.GetMessageDetailWithOptions(instanceId, topicName, messageId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Topic Details
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicResponse
func (client *Client) GetTopicWithOptions(instanceId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTopicResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopic"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Topic Details
//
// @return GetTopicResponse
func (client *Client) GetTopic(instanceId *string, topicName *string) (_result *GetTopicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTopicResponse{}
	_body, _err := client.GetTopicWithOptions(instanceId, topicName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trace of a specific message in a specific topic.
//
// @param request - GetTraceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceResponse
func (client *Client) GetTraceWithOptions(instanceId *string, topicName *string, messageId *string, request *GetTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrace"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName)) + "/traces/" + dara.PercentEncode(dara.StringValue(messageId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trace of a specific message in a specific topic.
//
// @param request - GetTraceRequest
//
// @return GetTraceResponse
func (client *Client) GetTrace(instanceId *string, topicName *string, messageId *string, request *GetTraceRequest) (_result *GetTraceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTraceResponse{}
	_body, _err := client.GetTraceWithOptions(instanceId, topicName, messageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询支持的可用区
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableZonesResponse
func (client *Client) ListAvailableZonesWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAvailableZonesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvailableZones"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/zones"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvailableZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支持的可用区
//
// @return ListAvailableZonesResponse
func (client *Client) ListAvailableZones() (_result *ListAvailableZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAvailableZonesResponse{}
	_body, _err := client.ListAvailableZonesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询消费者客户端连接信息
//
// @param request - ListConsumerConnectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerConnectionsResponse
func (client *Client) ListConsumerConnectionsWithOptions(instanceId *string, consumerGroupId *string, request *ListConsumerConnectionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerConnectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LiteTopicName) {
		query["liteTopicName"] = request.LiteTopicName
	}

	if !dara.IsNil(request.TopicName) {
		query["topicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumerConnections"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/connections"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumerConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者客户端连接信息
//
// @param request - ListConsumerConnectionsRequest
//
// @return ListConsumerConnectionsResponse
func (client *Client) ListConsumerConnections(instanceId *string, consumerGroupId *string, request *ListConsumerConnectionsRequest) (_result *ListConsumerConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerConnectionsResponse{}
	_body, _err := client.ListConsumerConnectionsWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the subscriptions of a specific consumer group.
//
// @param request - ListConsumerGroupSubscriptionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupSubscriptionsResponse
func (client *Client) ListConsumerGroupSubscriptionsWithOptions(instanceId *string, consumerGroupId *string, request *ListConsumerGroupSubscriptionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerGroupSubscriptionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TopicName) {
		query["topicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConsumerGroupSubscriptions"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/subscriptions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConsumerGroupSubscriptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscriptions of a specific consumer group.
//
// @param request - ListConsumerGroupSubscriptionsRequest
//
// @return ListConsumerGroupSubscriptionsResponse
func (client *Client) ListConsumerGroupSubscriptions(instanceId *string, consumerGroupId *string, request *ListConsumerGroupSubscriptionsRequest) (_result *ListConsumerGroupSubscriptionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerGroupSubscriptionsResponse{}
	_body, _err := client.ListConsumerGroupSubscriptionsWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Consumer Groups
//
// Description:
//
//	Notice: The OpenAPI provided by Alibaba Cloud is a management API used for managing and querying related resources of Alibaba Cloud services. It is recommended to integrate it only in the management chain. Do not rely on OpenAPI implementation in the core data chain for message sending and receiving, as this may lead to risks in the chain.
//
// @param request - ListConsumerGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupsResponse
func (client *Client) ListConsumerGroupsWithOptions(instanceId *string, request *ListConsumerGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
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
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups"),
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
// # List Consumer Groups
//
// Description:
//
//	Notice: The OpenAPI provided by Alibaba Cloud is a management API used for managing and querying related resources of Alibaba Cloud services. It is recommended to integrate it only in the management chain. Do not rely on OpenAPI implementation in the core data chain for message sending and receiving, as this may lead to risks in the chain.
//
// @param request - ListConsumerGroupsRequest
//
// @return ListConsumerGroupsResponse
func (client *Client) ListConsumerGroups(instanceId *string, request *ListConsumerGroupsRequest) (_result *ListConsumerGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerGroupsResponse{}
	_body, _err := client.ListConsumerGroupsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query disaster recovery plan consumption progress information
//
// @param request - ListDisasterRecoveryCheckpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisasterRecoveryCheckpointsResponse
func (client *Client) ListDisasterRecoveryCheckpointsWithOptions(planId *string, itemId *string, request *ListDisasterRecoveryCheckpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDisasterRecoveryCheckpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
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
		Action:      dara.String("ListDisasterRecoveryCheckpoints"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items/" + dara.PercentEncode(dara.StringValue(itemId)) + "/checkpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDisasterRecoveryCheckpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query disaster recovery plan consumption progress information
//
// @param request - ListDisasterRecoveryCheckpointsRequest
//
// @return ListDisasterRecoveryCheckpointsResponse
func (client *Client) ListDisasterRecoveryCheckpoints(planId *string, itemId *string, request *ListDisasterRecoveryCheckpointsRequest) (_result *ListDisasterRecoveryCheckpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDisasterRecoveryCheckpointsResponse{}
	_body, _err := client.ListDisasterRecoveryCheckpointsWithOptions(planId, itemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Global Replicator tasks of an instance.
//
// @param request - ListDisasterRecoveryItemsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisasterRecoveryItemsResponse
func (client *Client) ListDisasterRecoveryItemsWithOptions(planId *string, request *ListDisasterRecoveryItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDisasterRecoveryItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TopicName) {
		query["topicName"] = request.TopicName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDisasterRecoveryItems"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDisasterRecoveryItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Global Replicator tasks of an instance.
//
// @param request - ListDisasterRecoveryItemsRequest
//
// @return ListDisasterRecoveryItemsResponse
func (client *Client) ListDisasterRecoveryItems(planId *string, request *ListDisasterRecoveryItemsRequest) (_result *ListDisasterRecoveryItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDisasterRecoveryItemsResponse{}
	_body, _err := client.ListDisasterRecoveryItemsWithOptions(planId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Global Replicator tasks.
//
// @param request - ListDisasterRecoveryPlansRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisasterRecoveryPlansResponse
func (client *Client) ListDisasterRecoveryPlansWithOptions(request *ListDisasterRecoveryPlansRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDisasterRecoveryPlansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
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
		Action:      dara.String("ListDisasterRecoveryPlans"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDisasterRecoveryPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Global Replicator tasks.
//
// @param request - ListDisasterRecoveryPlansRequest
//
// @return ListDisasterRecoveryPlansResponse
func (client *Client) ListDisasterRecoveryPlans(request *ListDisasterRecoveryPlansRequest) (_result *ListDisasterRecoveryPlansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDisasterRecoveryPlansResponse{}
	_body, _err := client.ListDisasterRecoveryPlansWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the accounts that are used to access a specific instance.
//
// @param request - ListInstanceAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAccountResponse
func (client *Client) ListInstanceAccountWithOptions(instanceId *string, request *ListInstanceAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountStatus) {
		query["accountStatus"] = request.AccountStatus
	}

	if !dara.IsNil(request.AccountType) {
		query["accountType"] = request.AccountType
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Username) {
		query["username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceAccount"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/accounts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accounts that are used to access a specific instance.
//
// @param request - ListInstanceAccountRequest
//
// @return ListInstanceAccountResponse
func (client *Client) ListInstanceAccount(instanceId *string, request *ListInstanceAccountRequest) (_result *ListInstanceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceAccountResponse{}
	_body, _err := client.ListInstanceAccountWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the access control lists (ACLs) of an instance.
//
// @param request - ListInstanceAclRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAclResponse
func (client *Client) ListInstanceAclWithOptions(instanceId *string, request *ListInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
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
		Action:      dara.String("ListInstanceAcl"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/acl"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the access control lists (ACLs) of an instance.
//
// @param request - ListInstanceAclRequest
//
// @return ListInstanceAclResponse
func (client *Client) ListInstanceAcl(instanceId *string, request *ListInstanceAclRequest) (_result *ListInstanceAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceAclResponse{}
	_body, _err := client.ListInstanceAclWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists of an instance.
//
// @param request - ListInstanceIpWhitelistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceIpWhitelistResponse
func (client *Client) ListInstanceIpWhitelistWithOptions(instanceId *string, request *ListInstanceIpWhitelistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceIpWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpWhitelist) {
		query["ipWhitelist"] = request.IpWhitelist
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
		Action:      dara.String("ListInstanceIpWhitelist"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/ip/whitelist"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelists of an instance.
//
// @param request - ListInstanceIpWhitelistRequest
//
// @return ListInstanceIpWhitelistResponse
func (client *Client) ListInstanceIpWhitelist(instanceId *string, request *ListInstanceIpWhitelistRequest) (_result *ListInstanceIpWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceIpWhitelistResponse{}
	_body, _err := client.ListInstanceIpWhitelistWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all ApsaraMQ for RocketMQ instances in a specific region.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param tmpReq - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SeriesCodes) {
		request.SeriesCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SeriesCodes, dara.String("seriesCodes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
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

	if !dara.IsNil(request.SeriesCodesShrink) {
		query["seriesCodes"] = request.SeriesCodesShrink
	}

	if !dara.IsNil(request.StorageSecretKey) {
		query["storageSecretKey"] = request.StorageSecretKey
	}

	if !dara.IsNil(request.Tags) {
		query["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all ApsaraMQ for RocketMQ instances in a specific region.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of messages.
//
// @param request - ListMessagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessagesResponse
func (client *Client) ListMessagesWithOptions(instanceId *string, topicName *string, request *ListMessagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.LiteTopicName) {
		query["liteTopicName"] = request.LiteTopicName
	}

	if !dara.IsNil(request.MessageId) {
		query["messageId"] = request.MessageId
	}

	if !dara.IsNil(request.MessageKey) {
		query["messageKey"] = request.MessageKey
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScrollId) {
		query["scrollId"] = request.ScrollId
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMessages"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName)) + "/messages"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of messages.
//
// @param request - ListMessagesRequest
//
// @return ListMessagesResponse
func (client *Client) ListMessages(instanceId *string, topicName *string, request *ListMessagesRequest) (_result *ListMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMessagesResponse{}
	_body, _err := client.ListMessagesWithOptions(instanceId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Monitoring Items List
//
// @param request - ListMetricMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetricMetaResponse
func (client *Client) ListMetricMetaWithOptions(request *ListMetricMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMetricMetaResponse, _err error) {
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
		Action:      dara.String("ListMetricMeta"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/monitor/metrics/meta"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetricMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Monitoring Items List
//
// @param request - ListMetricMetaRequest
//
// @return ListMetricMetaResponse
func (client *Client) ListMetricMeta(request *ListMetricMetaRequest) (_result *ListMetricMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMetricMetaResponse{}
	_body, _err := client.ListMetricMetaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询迁移操作列表
//
// @param request - ListMigrationOperationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMigrationOperationsResponse
func (client *Client) ListMigrationOperationsWithOptions(migrationId *string, stageType *string, request *ListMigrationOperationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMigrationOperationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OperationType) {
		query["operationType"] = request.OperationType
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
		Action:      dara.String("ListMigrationOperations"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/migrations/" + dara.PercentEncode(dara.StringValue(migrationId)) + "/stages/" + dara.PercentEncode(dara.StringValue(stageType)) + "/operations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMigrationOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询迁移操作列表
//
// @param request - ListMigrationOperationsRequest
//
// @return ListMigrationOperationsResponse
func (client *Client) ListMigrationOperations(migrationId *string, stageType *string, request *ListMigrationOperationsRequest) (_result *ListMigrationOperationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMigrationOperationsResponse{}
	_body, _err := client.ListMigrationOperationsWithOptions(migrationId, stageType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询迁移列表
//
// @param request - ListMigrationsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMigrationsResponse
func (client *Client) ListMigrationsWithOptions(request *ListMigrationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMigrationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.MigrationType) {
		query["migrationType"] = request.MigrationType
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
		Action:      dara.String("ListMigrations"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/migrations"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMigrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询迁移列表
//
// @param request - ListMigrationsRequest
//
// @return ListMigrationsResponse
func (client *Client) ListMigrations(request *ListMigrationsRequest) (_result *ListMigrationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMigrationsResponse{}
	_body, _err := client.ListMigrationsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries regions in which ApsaraMQ for RocketMQ is available.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regions in which ApsaraMQ for RocketMQ is available.
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query visible resource tag relationships
//
// @param request - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resourceTag/list"),
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
// # Query visible resource tag relationships
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
// Queries the subscriptions of a specific topic.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicSubscriptionsResponse
func (client *Client) ListTopicSubscriptionsWithOptions(instanceId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTopicSubscriptionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTopicSubscriptions"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName)) + "/subscriptions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTopicSubscriptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscriptions of a specific topic.
//
// @return ListTopicSubscriptionsResponse
func (client *Client) ListTopicSubscriptions(instanceId *string, topicName *string) (_result *ListTopicSubscriptionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTopicSubscriptionsResponse{}
	_body, _err := client.ListTopicSubscriptionsWithOptions(instanceId, topicName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Topic List
//
// @param tmpReq - ListTopicsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicsResponse
func (client *Client) ListTopicsWithOptions(instanceId *string, tmpReq *ListTopicsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTopicsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTopicsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MessageTypes) {
		request.MessageTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MessageTypes, dara.String("messageTypes"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.MessageTypesShrink) {
		query["messageTypes"] = request.MessageTypesShrink
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
		Action:      dara.String("ListTopics"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTopicsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Topic List
//
// @param request - ListTopicsRequest
//
// @return ListTopicsResponse
func (client *Client) ListTopics(instanceId *string, request *ListTopicsRequest) (_result *ListTopicsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTopicsResponse{}
	_body, _err := client.ListTopicsWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the message traces of a specific topic.
//
// @param request - ListTracesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTracesResponse
func (client *Client) ListTracesWithOptions(instanceId *string, topicName *string, request *ListTracesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTracesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.LiteTopicName) {
		query["liteTopicName"] = request.LiteTopicName
	}

	if !dara.IsNil(request.MessageId) {
		query["messageId"] = request.MessageId
	}

	if !dara.IsNil(request.MessageKey) {
		query["messageKey"] = request.MessageKey
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["queryType"] = request.QueryType
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTraces"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName)) + "/traces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTracesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the message traces of a specific topic.
//
// @param request - ListTracesRequest
//
// @return ListTracesResponse
func (client *Client) ListTraces(instanceId *string, topicName *string, request *ListTracesRequest) (_result *ListTracesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTracesResponse{}
	_body, _err := client.ListTracesWithOptions(instanceId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the consumer offset of a consumer group.
//
// @param request - ResetConsumeOffsetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetConsumeOffsetResponse
func (client *Client) ResetConsumeOffsetWithOptions(instanceId *string, consumerGroupId *string, topicName *string, request *ResetConsumeOffsetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetConsumeOffsetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ResetTime) {
		body["resetTime"] = request.ResetTime
	}

	if !dara.IsNil(request.ResetType) {
		body["resetType"] = request.ResetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetConsumeOffset"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId)) + "/consumeOffsets/" + dara.PercentEncode(dara.StringValue(topicName))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetConsumeOffsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the consumer offset of a consumer group.
//
// @param request - ResetConsumeOffsetRequest
//
// @return ResetConsumeOffsetResponse
func (client *Client) ResetConsumeOffset(instanceId *string, consumerGroupId *string, topicName *string, request *ResetConsumeOffsetRequest) (_result *ResetConsumeOffsetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetConsumeOffsetResponse{}
	_body, _err := client.ResetConsumeOffsetWithOptions(instanceId, consumerGroupId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Enable Disaster Recovery Plan Entry
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDisasterRecoveryItemResponse
func (client *Client) StartDisasterRecoveryItemWithOptions(planId *string, itemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartDisasterRecoveryItemResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDisasterRecoveryItem"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items/" + dara.PercentEncode(dara.StringValue(itemId)) + "/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDisasterRecoveryItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Enable Disaster Recovery Plan Entry
//
// @return StartDisasterRecoveryItemResponse
func (client *Client) StartDisasterRecoveryItem(planId *string, itemId *string) (_result *StartDisasterRecoveryItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartDisasterRecoveryItemResponse{}
	_body, _err := client.StartDisasterRecoveryItemWithOptions(planId, itemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Deactivate Disaster Recovery Plan Entry
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDisasterRecoveryItemResponse
func (client *Client) StopDisasterRecoveryItemWithOptions(planId *string, itemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopDisasterRecoveryItemResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDisasterRecoveryItem"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items/" + dara.PercentEncode(dara.StringValue(itemId)) + "/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDisasterRecoveryItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Deactivate Disaster Recovery Plan Entry
//
// @return StopDisasterRecoveryItemResponse
func (client *Client) StopDisasterRecoveryItem(planId *string, itemId *string) (_result *StopDisasterRecoveryItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopDisasterRecoveryItemResponse{}
	_body, _err := client.StopDisasterRecoveryItemWithOptions(planId, itemId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Synchronize Disaster Recovery Plan Consumption Progress
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDisasterRecoveryCheckpointResponse
func (client *Client) SyncDisasterRecoveryCheckpointWithOptions(planId *string, itemId *string, checkpointId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncDisasterRecoveryCheckpointResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDisasterRecoveryCheckpoint"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items/" + dara.PercentEncode(dara.StringValue(itemId)) + "/checkpoints/" + dara.PercentEncode(dara.StringValue(checkpointId))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDisasterRecoveryCheckpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Synchronize Disaster Recovery Plan Consumption Progress
//
// @return SyncDisasterRecoveryCheckpointResponse
func (client *Client) SyncDisasterRecoveryCheckpoint(planId *string, itemId *string, checkpointId *string) (_result *SyncDisasterRecoveryCheckpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SyncDisasterRecoveryCheckpointResponse{}
	_body, _err := client.SyncDisasterRecoveryCheckpointWithOptions(planId, itemId, checkpointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates resource tags.
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
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resourceTag/create"),
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
// Creates resource tags.
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
// Removes tags from resources.
//
// @param request - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["all"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["tagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/resourceTag/delete"),
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
// # Update ConsumerGroup
//
// Description:
//
//	Notice: The OpenAPI provided by Alibaba Cloud is a management API used for managing and querying related resources of Alibaba Cloud services. It is recommended to integrate it only in the management chain. It is strictly prohibited to rely on OpenAPI implementation in the core data chain of message sending and receiving, otherwise it may lead to risks in the chain.
//
// @param request - UpdateConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerGroupResponse
func (client *Client) UpdateConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, request *UpdateConsumerGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConsumerGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConsumeRetryPolicy) {
		body["consumeRetryPolicy"] = request.ConsumeRetryPolicy
	}

	if !dara.IsNil(request.DeliveryOrderType) {
		body["deliveryOrderType"] = request.DeliveryOrderType
	}

	if !dara.IsNil(request.MaxReceiveTps) {
		body["maxReceiveTps"] = request.MaxReceiveTps
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConsumerGroup"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/consumerGroups/" + dara.PercentEncode(dara.StringValue(consumerGroupId))),
		Method:      dara.String("PATCH"),
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
// # Update ConsumerGroup
//
// Description:
//
//	Notice: The OpenAPI provided by Alibaba Cloud is a management API used for managing and querying related resources of Alibaba Cloud services. It is recommended to integrate it only in the management chain. It is strictly prohibited to rely on OpenAPI implementation in the core data chain of message sending and receiving, otherwise it may lead to risks in the chain.
//
// @param request - UpdateConsumerGroupRequest
//
// @return UpdateConsumerGroupResponse
func (client *Client) UpdateConsumerGroup(instanceId *string, consumerGroupId *string, request *UpdateConsumerGroupRequest) (_result *UpdateConsumerGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.UpdateConsumerGroupWithOptions(instanceId, consumerGroupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a topic mapping in a global message backup plan.
//
// @param request - UpdateDisasterRecoveryItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDisasterRecoveryItemResponse
func (client *Client) UpdateDisasterRecoveryItemWithOptions(planId *string, itemId *string, request *UpdateDisasterRecoveryItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDisasterRecoveryItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Topics) {
		body["topics"] = request.Topics
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDisasterRecoveryItem"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId)) + "/items/" + dara.PercentEncode(dara.StringValue(itemId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDisasterRecoveryItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a topic mapping in a global message backup plan.
//
// @param request - UpdateDisasterRecoveryItemRequest
//
// @return UpdateDisasterRecoveryItemResponse
func (client *Client) UpdateDisasterRecoveryItem(planId *string, itemId *string, request *UpdateDisasterRecoveryItemRequest) (_result *UpdateDisasterRecoveryItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDisasterRecoveryItemResponse{}
	_body, _err := client.UpdateDisasterRecoveryItemWithOptions(planId, itemId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a global message backup plan.
//
// @param request - UpdateDisasterRecoveryPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDisasterRecoveryPlanResponse
func (client *Client) UpdateDisasterRecoveryPlanWithOptions(planId *string, request *UpdateDisasterRecoveryPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDisasterRecoveryPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoSyncCheckpoint) {
		body["autoSyncCheckpoint"] = request.AutoSyncCheckpoint
	}

	if !dara.IsNil(request.Instances) {
		body["instances"] = request.Instances
	}

	if !dara.IsNil(request.PlanDesc) {
		body["planDesc"] = request.PlanDesc
	}

	if !dara.IsNil(request.PlanName) {
		body["planName"] = request.PlanName
	}

	if !dara.IsNil(request.PlanType) {
		body["planType"] = request.PlanType
	}

	if !dara.IsNil(request.SyncCheckpointEnabled) {
		body["syncCheckpointEnabled"] = request.SyncCheckpointEnabled
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDisasterRecoveryPlan"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/disaster_recovery/" + dara.PercentEncode(dara.StringValue(planId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDisasterRecoveryPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a global message backup plan.
//
// @param request - UpdateDisasterRecoveryPlanRequest
//
// @return UpdateDisasterRecoveryPlanResponse
func (client *Client) UpdateDisasterRecoveryPlan(planId *string, request *UpdateDisasterRecoveryPlanRequest) (_result *UpdateDisasterRecoveryPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDisasterRecoveryPlanResponse{}
	_body, _err := client.UpdateDisasterRecoveryPlanWithOptions(planId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the basic information and specifications of an ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - UpdateInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(instanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AclInfo) {
		body["aclInfo"] = request.AclInfo
	}

	if !dara.IsNil(request.InstanceName) {
		body["instanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.NetworkInfo) {
		body["networkInfo"] = request.NetworkInfo
	}

	if !dara.IsNil(request.ProductInfo) {
		body["productInfo"] = request.ProductInfo
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information and specifications of an ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(instanceId *string, request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a specific account in a specific instance.
//
// @param request - UpdateInstanceAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceAccountResponse
func (client *Client) UpdateInstanceAccountWithOptions(instanceId *string, username *string, request *UpdateInstanceAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountStatus) {
		query["accountStatus"] = request.AccountStatus
	}

	if !dara.IsNil(request.Password) {
		query["password"] = request.Password
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceAccount"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/accounts/" + dara.PercentEncode(dara.StringValue(username))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a specific account in a specific instance.
//
// @param request - UpdateInstanceAccountRequest
//
// @return UpdateInstanceAccountResponse
func (client *Client) UpdateInstanceAccount(instanceId *string, username *string, request *UpdateInstanceAccountRequest) (_result *UpdateInstanceAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceAccountResponse{}
	_body, _err := client.UpdateInstanceAccountWithOptions(instanceId, username, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the permissions on the resources of a specific instance for a specific user.
//
// @param request - UpdateInstanceAclRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceAclResponse
func (client *Client) UpdateInstanceAclWithOptions(instanceId *string, username *string, request *UpdateInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceAclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Actions) {
		body["actions"] = request.Actions
	}

	if !dara.IsNil(request.Decision) {
		body["decision"] = request.Decision
	}

	if !dara.IsNil(request.IpWhitelists) {
		body["ipWhitelists"] = request.IpWhitelists
	}

	if !dara.IsNil(request.ResourceName) {
		body["resourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceAcl"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/acl/account/" + dara.PercentEncode(dara.StringValue(username))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the permissions on the resources of a specific instance for a specific user.
//
// @param request - UpdateInstanceAclRequest
//
// @return UpdateInstanceAclResponse
func (client *Client) UpdateInstanceAcl(instanceId *string, username *string, request *UpdateInstanceAclRequest) (_result *UpdateInstanceAclResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInstanceAclResponse{}
	_body, _err := client.UpdateInstanceAclWithOptions(instanceId, username, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Topic
//
// @param request - UpdateTopicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTopicResponse
func (client *Client) UpdateTopicWithOptions(instanceId *string, topicName *string, request *UpdateTopicRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTopicResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LiteTopicExpiration) {
		body["liteTopicExpiration"] = request.LiteTopicExpiration
	}

	if !dara.IsNil(request.MaxSendTps) {
		body["maxSendTps"] = request.MaxSendTps
	}

	if !dara.IsNil(request.Remark) {
		body["remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTopic"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName))),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Topic
//
// @param request - UpdateTopicRequest
//
// @return UpdateTopicResponse
func (client *Client) UpdateTopic(instanceId *string, topicName *string, request *UpdateTopicRequest) (_result *UpdateTopicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTopicResponse{}
	_body, _err := client.UpdateTopicWithOptions(instanceId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the consumption status of a message in a specific topic on a specific instance.
//
// @param request - VerifyConsumeMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyConsumeMessageResponse
func (client *Client) VerifyConsumeMessageWithOptions(instanceId *string, topicName *string, messageId *string, request *VerifyConsumeMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyConsumeMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		query["clientId"] = request.ClientId
	}

	if !dara.IsNil(request.ConsumerGroupId) {
		query["consumerGroupId"] = request.ConsumerGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyConsumeMessage"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName)) + "/messages/" + dara.PercentEncode(dara.StringValue(messageId)) + "/action/verifyConsume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyConsumeMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the consumption status of a message in a specific topic on a specific instance.
//
// @param request - VerifyConsumeMessageRequest
//
// @return VerifyConsumeMessageResponse
func (client *Client) VerifyConsumeMessage(instanceId *string, topicName *string, messageId *string, request *VerifyConsumeMessageRequest) (_result *VerifyConsumeMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyConsumeMessageResponse{}
	_body, _err := client.VerifyConsumeMessageWithOptions(instanceId, topicName, messageId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Verifies the message sending feature of a specific topic on a specific instance.
//
// @param request - VerifySendMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySendMessageResponse
func (client *Client) VerifySendMessageWithOptions(instanceId *string, topicName *string, request *VerifySendMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifySendMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LiteTopicName) {
		body["liteTopicName"] = request.LiteTopicName
	}

	if !dara.IsNil(request.Message) {
		body["message"] = request.Message
	}

	if !dara.IsNil(request.MessageKey) {
		body["messageKey"] = request.MessageKey
	}

	if !dara.IsNil(request.MessageTag) {
		body["messageTag"] = request.MessageTag
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifySendMessage"),
		Version:     dara.String("2022-08-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/topics/" + dara.PercentEncode(dara.StringValue(topicName)) + "/messages"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifySendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies the message sending feature of a specific topic on a specific instance.
//
// @param request - VerifySendMessageRequest
//
// @return VerifySendMessageResponse
func (client *Client) VerifySendMessage(instanceId *string, topicName *string, request *VerifySendMessageRequest) (_result *VerifySendMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifySendMessageResponse{}
	_body, _err := client.VerifySendMessageWithOptions(instanceId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
