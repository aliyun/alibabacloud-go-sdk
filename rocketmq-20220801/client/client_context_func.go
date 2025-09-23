// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AddDisasterRecoveryItemWithContext(ctx context.Context, planId *string, request *AddDisasterRecoveryItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddDisasterRecoveryItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroupWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, request *CreateConsumerGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDisasterRecoveryPlanResponse
func (client *Client) CreateDisasterRecoveryPlanWithContext(ctx context.Context, request *CreateDisasterRecoveryPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDisasterRecoveryPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceResponse
func (client *Client) CreateInstanceWithContext(ctx context.Context, request *CreateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceAccountResponse
func (client *Client) CreateInstanceAccountWithContext(ctx context.Context, instanceId *string, request *CreateInstanceAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceAclResponse
func (client *Client) CreateInstanceAclWithContext(ctx context.Context, instanceId *string, username *string, request *CreateInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceAclResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceIpWhitelistResponse
func (client *Client) CreateInstanceIpWhitelistWithContext(ctx context.Context, instanceId *string, request *CreateInstanceIpWhitelistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateInstanceIpWhitelistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTopicResponse
func (client *Client) CreateTopicWithContext(ctx context.Context, instanceId *string, topicName *string, request *CreateTopicRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroupWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupSubscriptionResponse
func (client *Client) DeleteConsumerGroupSubscriptionWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, request *DeleteConsumerGroupSubscriptionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupSubscriptionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDisasterRecoveryItemResponse
func (client *Client) DeleteDisasterRecoveryItemWithContext(ctx context.Context, planId *string, itemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDisasterRecoveryItemResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDisasterRecoveryPlanResponse
func (client *Client) DeleteDisasterRecoveryPlanWithContext(ctx context.Context, planId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDisasterRecoveryPlanResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceAccountResponse
func (client *Client) DeleteInstanceAccountWithContext(ctx context.Context, instanceId *string, username *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceAccountResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceAclResponse
func (client *Client) DeleteInstanceAclWithContext(ctx context.Context, instanceId *string, username *string, request *DeleteInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceAclResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - DeleteInstanceIpWhitelistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceIpWhitelistResponse
func (client *Client) DeleteInstanceIpWhitelistWithContext(ctx context.Context, instanceId *string, tmpReq *DeleteInstanceIpWhitelistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteInstanceIpWhitelistResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopicWithContext(ctx context.Context, instanceId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteMigrationOperationResponse
func (client *Client) ExecuteMigrationOperationWithContext(ctx context.Context, migrationId *string, stageType *string, operationId *string, request *ExecuteMigrationOperationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteMigrationOperationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishMigrationStageResponse
func (client *Client) FinishMigrationStageWithContext(ctx context.Context, migrationId *string, stageType *string, request *FinishMigrationStageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FinishMigrationStageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerGroupResponse
func (client *Client) GetConsumerGroupWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerGroupResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerGroupLagResponse
func (client *Client) GetConsumerGroupLagWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, request *GetConsumerGroupLagRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerGroupLagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerGroupSubscriptionResponse
func (client *Client) GetConsumerGroupSubscriptionWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerGroupSubscriptionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerStackResponse
func (client *Client) GetConsumerStackWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, request *GetConsumerStackRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConsumerStackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDisasterRecoveryItemResponse
func (client *Client) GetDisasterRecoveryItemWithContext(ctx context.Context, planId *string, itemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDisasterRecoveryItemResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDisasterRecoveryPlanResponse
func (client *Client) GetDisasterRecoveryPlanWithContext(ctx context.Context, planId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDisasterRecoveryPlanResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithContext(ctx context.Context, instanceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceAccountResponse
func (client *Client) GetInstanceAccountWithContext(ctx context.Context, instanceId *string, request *GetInstanceAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceAclResponse
func (client *Client) GetInstanceAclWithContext(ctx context.Context, instanceId *string, username *string, request *GetInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceAclResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - GetInstanceIpWhitelistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceIpWhitelistResponse
func (client *Client) GetInstanceIpWhitelistWithContext(ctx context.Context, instanceId *string, tmpReq *GetInstanceIpWhitelistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInstanceIpWhitelistResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageDetailResponse
func (client *Client) GetMessageDetailWithContext(ctx context.Context, instanceId *string, topicName *string, messageId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMessageDetailResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicResponse
func (client *Client) GetTopicWithContext(ctx context.Context, instanceId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTopicResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceResponse
func (client *Client) GetTraceWithContext(ctx context.Context, instanceId *string, topicName *string, messageId *string, request *GetTraceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvailableZonesResponse
func (client *Client) ListAvailableZonesWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAvailableZonesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerConnectionsResponse
func (client *Client) ListConsumerConnectionsWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerConnectionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupSubscriptionsResponse
func (client *Client) ListConsumerGroupSubscriptionsWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, request *ListConsumerGroupSubscriptionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerGroupSubscriptionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupsResponse
func (client *Client) ListConsumerGroupsWithContext(ctx context.Context, instanceId *string, request *ListConsumerGroupsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConsumerGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisasterRecoveryCheckpointsResponse
func (client *Client) ListDisasterRecoveryCheckpointsWithContext(ctx context.Context, planId *string, itemId *string, request *ListDisasterRecoveryCheckpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDisasterRecoveryCheckpointsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisasterRecoveryItemsResponse
func (client *Client) ListDisasterRecoveryItemsWithContext(ctx context.Context, planId *string, request *ListDisasterRecoveryItemsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDisasterRecoveryItemsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisasterRecoveryPlansResponse
func (client *Client) ListDisasterRecoveryPlansWithContext(ctx context.Context, request *ListDisasterRecoveryPlansRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDisasterRecoveryPlansResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAccountResponse
func (client *Client) ListInstanceAccountWithContext(ctx context.Context, instanceId *string, request *ListInstanceAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAclResponse
func (client *Client) ListInstanceAclWithContext(ctx context.Context, instanceId *string, request *ListInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceAclResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceIpWhitelistResponse
func (client *Client) ListInstanceIpWhitelistWithContext(ctx context.Context, instanceId *string, request *ListInstanceIpWhitelistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstanceIpWhitelistResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithContext(ctx context.Context, tmpReq *ListInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMessagesResponse
func (client *Client) ListMessagesWithContext(ctx context.Context, instanceId *string, topicName *string, request *ListMessagesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetricMetaResponse
func (client *Client) ListMetricMetaWithContext(ctx context.Context, request *ListMetricMetaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMetricMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMigrationOperationsResponse
func (client *Client) ListMigrationOperationsWithContext(ctx context.Context, migrationId *string, stageType *string, request *ListMigrationOperationsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMigrationOperationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicSubscriptionsResponse
func (client *Client) ListTopicSubscriptionsWithContext(ctx context.Context, instanceId *string, topicName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTopicSubscriptionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param tmpReq - ListTopicsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicsResponse
func (client *Client) ListTopicsWithContext(ctx context.Context, instanceId *string, tmpReq *ListTopicsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTopicsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTracesResponse
func (client *Client) ListTracesWithContext(ctx context.Context, instanceId *string, topicName *string, request *ListTracesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTracesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetConsumeOffsetResponse
func (client *Client) ResetConsumeOffsetWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, topicName *string, request *ResetConsumeOffsetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetConsumeOffsetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDisasterRecoveryItemResponse
func (client *Client) StartDisasterRecoveryItemWithContext(ctx context.Context, planId *string, itemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartDisasterRecoveryItemResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDisasterRecoveryItemResponse
func (client *Client) StopDisasterRecoveryItemWithContext(ctx context.Context, planId *string, itemId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopDisasterRecoveryItemResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDisasterRecoveryCheckpointResponse
func (client *Client) SyncDisasterRecoveryCheckpointWithContext(ctx context.Context, planId *string, itemId *string, checkpointId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SyncDisasterRecoveryCheckpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerGroupResponse
func (client *Client) UpdateConsumerGroupWithContext(ctx context.Context, instanceId *string, consumerGroupId *string, request *UpdateConsumerGroupRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConsumerGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDisasterRecoveryItemResponse
func (client *Client) UpdateDisasterRecoveryItemWithContext(ctx context.Context, planId *string, itemId *string, request *UpdateDisasterRecoveryItemRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDisasterRecoveryItemResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDisasterRecoveryPlanResponse
func (client *Client) UpdateDisasterRecoveryPlanWithContext(ctx context.Context, planId *string, request *UpdateDisasterRecoveryPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDisasterRecoveryPlanResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithContext(ctx context.Context, instanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceAccountResponse
func (client *Client) UpdateInstanceAccountWithContext(ctx context.Context, instanceId *string, username *string, request *UpdateInstanceAccountRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceAccountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceAclResponse
func (client *Client) UpdateInstanceAclWithContext(ctx context.Context, instanceId *string, username *string, request *UpdateInstanceAclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateInstanceAclResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTopicResponse
func (client *Client) UpdateTopicWithContext(ctx context.Context, instanceId *string, topicName *string, request *UpdateTopicRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyConsumeMessageResponse
func (client *Client) VerifyConsumeMessageWithContext(ctx context.Context, instanceId *string, topicName *string, messageId *string, request *VerifyConsumeMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyConsumeMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifySendMessageResponse
func (client *Client) VerifySendMessageWithContext(ctx context.Context, instanceId *string, topicName *string, request *VerifySendMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifySendMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
