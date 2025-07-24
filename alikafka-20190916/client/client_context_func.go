// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Changes the resource group of an ApsaraMQ for Kafka instance.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithContext(ctx context.Context, request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Changes the billing method of a Message Queue for Apache Kafka instance from pay-as-you-go to subscription.
//
// @param request - ConvertPostPayOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertPostPayOrderResponse
func (client *Client) ConvertPostPayOrderWithContext(ctx context.Context, request *ConvertPostPayOrderRequest, runtime *dara.RuntimeOptions) (_result *ConvertPostPayOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PaidType) {
		query["PaidType"] = request.PaidType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ConvertPostPayOrder"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertPostPayOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an access control list (ACL).
//
// @param request - CreateAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAclResponse
func (client *Client) CreateAclWithContext(ctx context.Context, request *CreateAclRequest, runtime *dara.RuntimeOptions) (_result *CreateAclResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclOperationType) {
		query["AclOperationType"] = request.AclOperationType
	}

	if !dara.IsNil(request.AclOperationTypes) {
		query["AclOperationTypes"] = request.AclOperationTypes
	}

	if !dara.IsNil(request.AclPermissionType) {
		query["AclPermissionType"] = request.AclPermissionType
	}

	if !dara.IsNil(request.AclResourceName) {
		query["AclResourceName"] = request.AclResourceName
	}

	if !dara.IsNil(request.AclResourcePatternType) {
		query["AclResourcePatternType"] = request.AclResourcePatternType
	}

	if !dara.IsNil(request.AclResourceType) {
		query["AclResourceType"] = request.AclResourceType
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAcl"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAclResponse{}
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
// @param request - CreateConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroupWithContext(ctx context.Context, request *CreateConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConsumerGroup"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 创建后付费实例。
//
// @param tmpReq - CreatePostPayInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePostPayInstanceResponse
func (client *Client) CreatePostPayInstanceWithContext(ctx context.Context, tmpReq *CreatePostPayInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreatePostPayInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreatePostPayInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServerlessConfig) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, dara.String("ServerlessConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeployType) {
		query["DeployType"] = request.DeployType
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.EipMax) {
		query["EipMax"] = request.EipMax
	}

	if !dara.IsNil(request.IoMaxSpec) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !dara.IsNil(request.PaidType) {
		query["PaidType"] = request.PaidType
	}

	if !dara.IsNil(request.PartitionNum) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServerlessConfigShrink) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePostPayInstance"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePostPayInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a pay-as-you-go ApsaraMQ for Kafka instance. Pay-as-you-go instances allow you to pay after you use the resources. You are charged for pay-as-you-go instances based on the actual resource usage. You can use pay-as-you-go instances in test scenarios or scenarios in which the peak traffic is uncertain.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of pay-as-you-go Message Queue for Apache Kafka instances. For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
//
// @param tmpReq - CreatePostPayOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePostPayOrderResponse
func (client *Client) CreatePostPayOrderWithContext(ctx context.Context, tmpReq *CreatePostPayOrderRequest, runtime *dara.RuntimeOptions) (_result *CreatePostPayOrderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreatePostPayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServerlessConfig) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, dara.String("ServerlessConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DeployType) {
		query["DeployType"] = request.DeployType
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.EipMax) {
		query["EipMax"] = request.EipMax
	}

	if !dara.IsNil(request.IoMax) {
		query["IoMax"] = request.IoMax
	}

	if !dara.IsNil(request.IoMaxSpec) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !dara.IsNil(request.PaidType) {
		query["PaidType"] = request.PaidType
	}

	if !dara.IsNil(request.PartitionNum) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ServerlessConfigShrink) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TopicQuota) {
		query["TopicQuota"] = request.TopicQuota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePostPayOrder"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePostPayOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预付费实例
//
// @param tmpReq - CreatePrePayInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrePayInstanceResponse
func (client *Client) CreatePrePayInstanceWithContext(ctx context.Context, tmpReq *CreatePrePayInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreatePrePayInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreatePrePayInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfluentConfig) {
		request.ConfluentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfluentConfig, dara.String("ConfluentConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfluentConfigShrink) {
		query["ConfluentConfig"] = request.ConfluentConfigShrink
	}

	if !dara.IsNil(request.DeployType) {
		query["DeployType"] = request.DeployType
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EipMax) {
		query["EipMax"] = request.EipMax
	}

	if !dara.IsNil(request.IoMaxSpec) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !dara.IsNil(request.PaidType) {
		query["PaidType"] = request.PaidType
	}

	if !dara.IsNil(request.PartitionNum) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrePayInstance"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrePayInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a subscription ApsaraMQ for Kafka instance. You can use subscription instances only after you pay for them. Subscription instances are suitable for long-term and stable business scenarios.
//
// Description:
//
//	  Before you call this operation, make sure that you understand the billing methods and pricing of subscription ApsaraMQ for Kafka instances. For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
//
//		- If you create an ApsaraMQ for Kafka instance by calling this operation, the subscription duration is one month and the auto-renewal feature is enabled by default. The auto-renewal cycle is also one month. If you want to change the auto-renewal cycle or disable the auto-renewal feature, you can go to the [Renewal](https://renew.console.aliyun.com/#/ecs) page in the Alibaba Cloud Management Console.
//
// @param tmpReq - CreatePrePayOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrePayOrderResponse
func (client *Client) CreatePrePayOrderWithContext(ctx context.Context, tmpReq *CreatePrePayOrderRequest, runtime *dara.RuntimeOptions) (_result *CreatePrePayOrderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreatePrePayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfluentConfig) {
		request.ConfluentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfluentConfig, dara.String("ConfluentConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfluentConfigShrink) {
		query["ConfluentConfig"] = request.ConfluentConfigShrink
	}

	if !dara.IsNil(request.DeployType) {
		query["DeployType"] = request.DeployType
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EipMax) {
		query["EipMax"] = request.EipMax
	}

	if !dara.IsNil(request.IoMax) {
		query["IoMax"] = request.IoMax
	}

	if !dara.IsNil(request.IoMaxSpec) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !dara.IsNil(request.PaidType) {
		query["PaidType"] = request.PaidType
	}

	if !dara.IsNil(request.PartitionNum) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TopicQuota) {
		query["TopicQuota"] = request.TopicQuota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrePayOrder"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrePayOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Simple Authentication and Security Layer (SASL) user.
//
// @param request - CreateSaslUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSaslUserResponse
func (client *Client) CreateSaslUserWithContext(ctx context.Context, request *CreateSaslUserRequest, runtime *dara.RuntimeOptions) (_result *CreateSaslUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mechanism) {
		query["Mechanism"] = request.Mechanism
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSaslUser"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSaslUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled scaling rule for a serverless ApsaraMQ for Kafka V3 instance.
//
// Description:
//
// ###### [](#-v3-serverless-)This operation is supported only by serverless ApsaraMQ for Kafka V3 instances.
//
// @param tmpReq - CreateScheduledScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledScalingRuleResponse
func (client *Client) CreateScheduledScalingRuleWithContext(ctx context.Context, tmpReq *CreateScheduledScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateScheduledScalingRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateScheduledScalingRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WeeklyTypes) {
		request.WeeklyTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WeeklyTypes, dara.String("WeeklyTypes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DurationMinutes) {
		query["DurationMinutes"] = request.DurationMinutes
	}

	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.FirstScheduledTime) {
		query["FirstScheduledTime"] = request.FirstScheduledTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RepeatType) {
		query["RepeatType"] = request.RepeatType
	}

	if !dara.IsNil(request.ReservedPubFlow) {
		query["ReservedPubFlow"] = request.ReservedPubFlow
	}

	if !dara.IsNil(request.ReservedSubFlow) {
		query["ReservedSubFlow"] = request.ReservedSubFlow
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.ScheduleType) {
		query["ScheduleType"] = request.ScheduleType
	}

	if !dara.IsNil(request.TimeZone) {
		query["TimeZone"] = request.TimeZone
	}

	if !dara.IsNil(request.WeeklyTypesShrink) {
		query["WeeklyTypes"] = request.WeeklyTypesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScheduledScalingRule"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScheduledScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a topic.
//
// Description:
//
//	  Each Alibaba Cloud account can call this operation up to once per second.
//
//		- The maximum number of topics that you can create in an instance is determined by the specification of the instance.
//
// @param request - CreateTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTopicResponse
func (client *Client) CreateTopicWithContext(ctx context.Context, request *CreateTopicRequest, runtime *dara.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompactTopic) {
		query["CompactTopic"] = request.CompactTopic
	}

	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.LocalTopic) {
		query["LocalTopic"] = request.LocalTopic
	}

	if !dara.IsNil(request.MinInsyncReplicas) {
		query["MinInsyncReplicas"] = request.MinInsyncReplicas
	}

	if !dara.IsNil(request.PartitionNum) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.ReplicationFactor) {
		query["ReplicationFactor"] = request.ReplicationFactor
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTopic"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Deletes an access control list (ACL).
//
// @param request - DeleteAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAclResponse
func (client *Client) DeleteAclWithContext(ctx context.Context, request *DeleteAclRequest, runtime *dara.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclOperationType) {
		query["AclOperationType"] = request.AclOperationType
	}

	if !dara.IsNil(request.AclOperationTypes) {
		query["AclOperationTypes"] = request.AclOperationTypes
	}

	if !dara.IsNil(request.AclPermissionType) {
		query["AclPermissionType"] = request.AclPermissionType
	}

	if !dara.IsNil(request.AclResourceName) {
		query["AclResourceName"] = request.AclResourceName
	}

	if !dara.IsNil(request.AclResourcePatternType) {
		query["AclResourcePatternType"] = request.AclResourcePatternType
	}

	if !dara.IsNil(request.AclResourceType) {
		query["AclResourceType"] = request.AclResourceType
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAcl"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a consumer group from a specified Message Queue for Apache Kafka instance.
//
// @param request - DeleteConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
func (client *Client) DeleteConsumerGroupWithContext(ctx context.Context, request *DeleteConsumerGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConsumerGroup"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Deletes an instance. You can delete subscription and pay-as-you-go instances after you release them.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithContext(ctx context.Context, request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Deletes a Simple Authentication and Security Layer (SASL) user.
//
// @param request - DeleteSaslUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSaslUserResponse
func (client *Client) DeleteSaslUserWithContext(ctx context.Context, request *DeleteSaslUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteSaslUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Mechanism) {
		query["Mechanism"] = request.Mechanism
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSaslUser"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSaslUserResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the scheduled scaling policy of a serverless ApsaraMQ for Kafka instance after you deploy the instance.
//
// Description:
//
// ###### [](#-serverless-)This operation is available only for serverless ApsaraMQ for Kafka instances.
//
// @param request - DeleteScheduledScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledScalingRuleResponse
func (client *Client) DeleteScheduledScalingRuleWithContext(ctx context.Context, request *DeleteScheduledScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteScheduledScalingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScheduledScalingRule"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScheduledScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a topic.
//
// @param request - DeleteTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopicWithContext(ctx context.Context, request *DeleteTopicRequest, runtime *dara.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTopic"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// 查询acl资源名
//
// @param request - DescribeAclResourceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclResourceNameResponse
func (client *Client) DescribeAclResourceNameWithContext(ctx context.Context, request *DescribeAclResourceNameRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclResourceNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclResourcePatternType) {
		query["AclResourcePatternType"] = request.AclResourcePatternType
	}

	if !dara.IsNil(request.AclResourceType) {
		query["AclResourceType"] = request.AclResourceType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAclResourceName"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclResourceNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries access control lists (ACLs).
//
// @param request - DescribeAclsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclsResponse
func (client *Client) DescribeAclsWithContext(ctx context.Context, request *DescribeAclsRequest, runtime *dara.RuntimeOptions) (_result *DescribeAclsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AclOperationType) {
		query["AclOperationType"] = request.AclOperationType
	}

	if !dara.IsNil(request.AclPermissionType) {
		query["AclPermissionType"] = request.AclPermissionType
	}

	if !dara.IsNil(request.AclResourceName) {
		query["AclResourceName"] = request.AclResourceName
	}

	if !dara.IsNil(request.AclResourcePatternType) {
		query["AclResourcePatternType"] = request.AclResourcePatternType
	}

	if !dara.IsNil(request.AclResourceType) {
		query["AclResourceType"] = request.AclResourceType
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAcls"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAclsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Simple Authentication and Security Layer (SASL) users.
//
// @param request - DescribeSaslUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSaslUsersResponse
func (client *Client) DescribeSaslUsersWithContext(ctx context.Context, request *DescribeSaslUsersRequest, runtime *dara.RuntimeOptions) (_result *DescribeSaslUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSaslUsers"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSaslUsersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables and disables the flexible group creation feature.
//
// @param request - EnableAutoGroupCreationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAutoGroupCreationResponse
func (client *Client) EnableAutoGroupCreationWithContext(ctx context.Context, request *EnableAutoGroupCreationRequest, runtime *dara.RuntimeOptions) (_result *EnableAutoGroupCreationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAutoGroupCreation"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAutoGroupCreationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the automatic topic creation feature, or changes the number of partitions in topics that are automatically created.
//
// @param request - EnableAutoTopicCreationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAutoTopicCreationResponse
func (client *Client) EnableAutoTopicCreationWithContext(ctx context.Context, request *EnableAutoTopicCreationRequest, runtime *dara.RuntimeOptions) (_result *EnableAutoTopicCreationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Operate) {
		query["Operate"] = request.Operate
	}

	if !dara.IsNil(request.PartitionNum) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UpdatePartition) {
		query["UpdatePartition"] = request.UpdatePartition
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableAutoTopicCreation"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableAutoTopicCreationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IDs of all instances in the current account.
//
// @param request - GetAllInstanceIdListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllInstanceIdListResponse
func (client *Client) GetAllInstanceIdListWithContext(ctx context.Context, request *GetAllInstanceIdListRequest, runtime *dara.RuntimeOptions) (_result *GetAllInstanceIdListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAllInstanceIdList"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAllInstanceIdListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP address whitelist.
//
// @param request - GetAllowedIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllowedIpListResponse
func (client *Client) GetAllowedIpListWithContext(ctx context.Context, request *GetAllowedIpListRequest, runtime *dara.RuntimeOptions) (_result *GetAllowedIpListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAllowedIpList"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAllowedIpListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scheduled scaling policy of a serverless ApsaraMQ for Kafka instance after you deploy the instance.
//
// Description:
//
// ###### [](#-serverless-)**This operation is available only for serverless ApsaraMQ for Kafka instances.
//
// @param request - GetAutoScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoScalingConfigurationResponse
func (client *Client) GetAutoScalingConfigurationWithContext(ctx context.Context, request *GetAutoScalingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetAutoScalingConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAutoScalingConfiguration"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAutoScalingConfigurationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more consumer groups in a specified Message Queue for Apache Kafka instance.
//
// @param request - GetConsumerListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerListResponse
func (client *Client) GetConsumerListWithContext(ctx context.Context, request *GetConsumerListRequest, runtime *dara.RuntimeOptions) (_result *GetConsumerListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerList"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the consumer progress of a consumer group.
//
// @param request - GetConsumerProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerProgressResponse
func (client *Client) GetConsumerProgressWithContext(ctx context.Context, request *GetConsumerProgressRequest, runtime *dara.RuntimeOptions) (_result *GetConsumerProgressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.HideLastTimestamp) {
		query["HideLastTimestamp"] = request.HideLastTimestamp
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConsumerProgress"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConsumerProgressResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about instances in a specified region.
//
// @param request - GetInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceListResponse
func (client *Client) GetInstanceListWithContext(ctx context.Context, request *GetInstanceListRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Series) {
		query["Series"] = request.Series
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceList"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IP addresses of the clients that are connected to an ApsaraMQ for Kafka instance.
//
// Description:
//
//	  The IP information is obtained from the sampled logs generated for the requests that the client sends to the broker by calling the API operations of ApsaraMQ for Kafka.
//
//		- Statistics refers to the number of connections on different ports of an IP address within a specific period of time.
//
//		- If the broker is not of the latest minor version, the sampled logs may not be accurate. This may cause inaccurate IP information. Therefore, we recommend that you update your broker to the latest version at the earliest opportunity.
//
// @param request - GetKafkaClientIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKafkaClientIpResponse
func (client *Client) GetKafkaClientIpWithContext(ctx context.Context, request *GetKafkaClientIpRequest, runtime *dara.RuntimeOptions) (_result *GetKafkaClientIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Group) {
		query["Group"] = request.Group
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKafkaClientIp"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKafkaClientIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the used quota of topics and partitions.
//
// @param request - GetQuotaTipRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQuotaTipResponse
func (client *Client) GetQuotaTipWithContext(ctx context.Context, request *GetQuotaTipRequest, runtime *dara.RuntimeOptions) (_result *GetQuotaTipResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQuotaTip"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQuotaTipResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例风险列表
//
// @param request - GetRiskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRiskListResponse
func (client *Client) GetRiskListWithContext(ctx context.Context, request *GetRiskListRequest, runtime *dara.RuntimeOptions) (_result *GetRiskListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartIndex) {
		query["StartIndex"] = request.StartIndex
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRiskList"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRiskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a topic.
//
// @param request - GetTopicListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicListResponse
func (client *Client) GetTopicListWithContext(ctx context.Context, request *GetTopicListRequest, runtime *dara.RuntimeOptions) (_result *GetTopicListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopicList"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the messaging status of a topic.
//
// @param request - GetTopicStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicStatusResponse
func (client *Client) GetTopicStatusWithContext(ctx context.Context, request *GetTopicStatusRequest, runtime *dara.RuntimeOptions) (_result *GetTopicStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopicStatus"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the groups that subscribe to a topic.
//
// @param request - GetTopicSubscribeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicSubscribeStatusResponse
func (client *Client) GetTopicSubscribeStatusWithContext(ctx context.Context, request *GetTopicSubscribeStatusRequest, runtime *dara.RuntimeOptions) (_result *GetTopicSubscribeStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopicSubscribeStatus"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicSubscribeStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Rebalance详情
//
// @param request - ListRebalanceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRebalanceInfoResponse
func (client *Client) ListRebalanceInfoWithContext(ctx context.Context, request *ListRebalanceInfoRequest, runtime *dara.RuntimeOptions) (_result *ListRebalanceInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRebalanceInfo"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRebalanceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are attached to a specified resource.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithContext(ctx context.Context, request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Changes the name of an ApsaraMQ for Kafka instance. After you deploy an instance, you can call this operation to change the name of the instance.
//
// @param request - ModifyInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceNameResponse
func (client *Client) ModifyInstanceNameWithContext(ctx context.Context, request *ModifyInstanceNameRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstanceName"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the number of partitions in a topic.
//
// @param request - ModifyPartitionNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPartitionNumResponse
func (client *Client) ModifyPartitionNumWithContext(ctx context.Context, request *ModifyPartitionNumRequest, runtime *dara.RuntimeOptions) (_result *ModifyPartitionNumResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AddPartitionNum) {
		query["AddPartitionNum"] = request.AddPartitionNum
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyPartitionNum"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPartitionNumResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the scheduled scaling policy of a serverless ApsaraMQ for Kafka instance after you deploy the instance.
//
// Description:
//
// ###### [](#-serverless-)This operation is available only for serverless ApsaraMQ for Kafka instances.
//
// @param request - ModifyScheduledScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduledScalingRuleResponse
func (client *Client) ModifyScheduledScalingRuleWithContext(ctx context.Context, request *ModifyScheduledScalingRuleRequest, runtime *dara.RuntimeOptions) (_result *ModifyScheduledScalingRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Enable) {
		query["Enable"] = request.Enable
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyScheduledScalingRule"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyScheduledScalingRuleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a topic.
//
// @param request - ModifyTopicRemarkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTopicRemarkResponse
func (client *Client) ModifyTopicRemarkWithContext(ctx context.Context, request *ModifyTopicRemarkRequest, runtime *dara.RuntimeOptions) (_result *ModifyTopicRemarkResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTopicRemark"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTopicRemarkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries messages stored in a topic. You can query messages by creation time or offset.
//
// @param request - QueryMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageResponse
func (client *Client) QueryMessageWithContext(ctx context.Context, request *QueryMessageRequest, runtime *dara.RuntimeOptions) (_result *QueryMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMessage"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases a pay-as-you-go instance.
//
// Description:
//
// You cannot call this operation to release a subscription Message Queue for Apache Kafka instance.
//
// @param request - ReleaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstanceResponse
func (client *Client) ReleaseInstanceWithContext(ctx context.Context, request *ReleaseInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceDeleteInstance) {
		query["ForceDeleteInstance"] = request.ForceDeleteInstance
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseInstance"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an ApsaraMQ for Kafka instance.
//
// Description:
//
// You can call this operation only if your instance is in the Stopped state.
//
// @param request - ReopenInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReopenInstanceResponse
func (client *Client) ReopenInstanceWithContext(ctx context.Context, request *ReopenInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReopenInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReopenInstance"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReopenInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys an ApsaraMQ for Kafka instance. You must purchase and deploy an ApsaraMQ for Kafka instance before you can use the instance to send and receive messages.
//
// Description:
//
// >  You can call this operation up to twice per second.
//
// @param request - StartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithContext(ctx context.Context, request *StartInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.CrossZone) {
		query["CrossZone"] = request.CrossZone
	}

	if !dara.IsNil(request.DeployModule) {
		query["DeployModule"] = request.DeployModule
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IsEipInner) {
		query["IsEipInner"] = request.IsEipInner
	}

	if !dara.IsNil(request.IsForceSelectedZones) {
		query["IsForceSelectedZones"] = request.IsForceSelectedZones
	}

	if !dara.IsNil(request.IsSetUserAndPassword) {
		query["IsSetUserAndPassword"] = request.IsSetUserAndPassword
	}

	if !dara.IsNil(request.KMSKeyId) {
		query["KMSKeyId"] = request.KMSKeyId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Notifier) {
		query["Notifier"] = request.Notifier
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SecurityGroup) {
		query["SecurityGroup"] = request.SecurityGroup
	}

	if !dara.IsNil(request.SelectedZones) {
		query["SelectedZones"] = request.SelectedZones
	}

	if !dara.IsNil(request.ServiceVersion) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !dara.IsNil(request.UserPhoneNum) {
		query["UserPhoneNum"] = request.UserPhoneNum
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VSwitchIds) {
		query["VSwitchIds"] = request.VSwitchIds
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
		Action:      dara.String("StartInstance"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an ApsaraMQ for Kafka instance.
//
// Description:
//
// You cannot stop a subscription ApsaraMQ for Kafka instance. If you want to stop a subscription ApsaraMQ for Kafka instance, submit a ticket.
//
// @param request - StopInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithContext(ctx context.Context, request *StopInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInstance"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a tag to a resource.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithContext(ctx context.Context, request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
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

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Detaches tags from a specified resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithContext(ctx context.Context, request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
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
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
// Updates the IP address whitelist of an ApsaraMQ for Kafka instance. Only IP addresses and ports that are configured in the IP address whitelist of an instance can access the instance.
//
// @param request - UpdateAllowedIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAllowedIpResponse
func (client *Client) UpdateAllowedIpWithContext(ctx context.Context, request *UpdateAllowedIpRequest, runtime *dara.RuntimeOptions) (_result *UpdateAllowedIpResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowedListIp) {
		query["AllowedListIp"] = request.AllowedListIp
	}

	if !dara.IsNil(request.AllowedListType) {
		query["AllowedListType"] = request.AllowedListType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PortRange) {
		query["PortRange"] = request.PortRange
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.UpdateType) {
		query["UpdateType"] = request.UpdateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAllowedIp"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAllowedIpResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the consumer offsets of the subscribed topics of a consumer group.
//
// Description:
//
// You can call this operation to reset the consumer offset of a specific consumer group. You can use the timestamp or offset parameter to reset the consumer offset of a consumer group. You can implement the following features by configuring a combination of different parameters:
//
//   - Reset the consumer offsets of one or all subscribed topics of a consumer group to the latest offset. This way, you can consume messages in the topics from the latest offset.
//
//   - Reset the consumer offsets of one or all subscribed topics of a consumer group to a specific point in time. This way, you can consume messages in the topics from the specified point in time.
//
//   - Reset the consumer offset of one subscribed topic of a consumer group to a specific offset in a specific partition. This way, you can consume messages from the specified offset in the specified partition.
//
// @param tmpReq - UpdateConsumerOffsetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerOffsetResponse
func (client *Client) UpdateConsumerOffsetWithContext(ctx context.Context, tmpReq *UpdateConsumerOffsetRequest, runtime *dara.RuntimeOptions) (_result *UpdateConsumerOffsetResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateConsumerOffsetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Offsets) {
		request.OffsetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Offsets, dara.String("Offsets"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConsumerId) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.OffsetsShrink) {
		query["Offsets"] = request.OffsetsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResetType) {
		query["ResetType"] = request.ResetType
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConsumerOffset"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConsumerOffsetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of an ApsaraMQ for Kafka instance. ApsaraMQ for Kafka allows you to modify the configurations of an instance, including the access control list (ACL) feature, the Secure Sockets Layer (SSL) feature, the message retention period, and the maximum message size.
//
// Description:
//
// ## **Permissions**
//
// If a RAM user wants to call the **UpdateInstanceConfig*	- operation, the RAM user must be granted the required permissions. For more information about how to grant permissions, see [RAM policies](https://help.aliyun.com/document_detail/185815.html).
//
// |API|Action|Resource|
//
// |---|---|---|
//
// |UpdateInstanceConfig|alikafka: UpdateInstance|acs:alikafka:*:*:{instanceId}|
//
// @param request - UpdateInstanceConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceConfigResponse
func (client *Client) UpdateInstanceConfigWithContext(ctx context.Context, request *UpdateInstanceConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstanceConfig"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configurations of a topic. After you create a topic, you can modify the message retention period and maximum message size of the topic.
//
// @param request - UpdateTopicConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTopicConfigResponse
func (client *Client) UpdateTopicConfigWithContext(ctx context.Context, request *UpdateTopicConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateTopicConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Topic) {
		query["Topic"] = request.Topic
	}

	if !dara.IsNil(request.Value) {
		query["Value"] = request.Value
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTopicConfig"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTopicConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version of an instance.
//
// Description:
//
// ## **Permissions**
//
// A RAM user must be granted the required permissions before the RAM user calls the **UpgradeInstanceVersion*	- operation. For information about how to grant permissions, see [RAM policies](https://help.aliyun.com/document_detail/185815.html).
//
// |API|Action|Resource|
//
// |---|---|---|
//
// |UpgradeInstanceVersion|UpdateInstance|acs:alikafka:*:*:{instanceId}|
//
// ## **QPS limits**
//
// You can send a maximum of two queries per second (QPS).
//
// @param request - UpgradeInstanceVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeInstanceVersionResponse
func (client *Client) UpgradeInstanceVersionWithContext(ctx context.Context, request *UpgradeInstanceVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeInstanceVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetVersion) {
		query["TargetVersion"] = request.TargetVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeInstanceVersion"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a pay-as-you-go ApsaraMQ for Kafka instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of pay-as-you-go Message Queue for Apache Kafka instances. For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
//
// @param tmpReq - UpgradePostPayOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradePostPayOrderResponse
func (client *Client) UpgradePostPayOrderWithContext(ctx context.Context, tmpReq *UpgradePostPayOrderRequest, runtime *dara.RuntimeOptions) (_result *UpgradePostPayOrderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpgradePostPayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServerlessConfig) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, dara.String("ServerlessConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.EipMax) {
		query["EipMax"] = request.EipMax
	}

	if !dara.IsNil(request.EipModel) {
		query["EipModel"] = request.EipModel
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IoMax) {
		query["IoMax"] = request.IoMax
	}

	if !dara.IsNil(request.IoMaxSpec) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !dara.IsNil(request.PartitionNum) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServerlessConfigShrink) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	if !dara.IsNil(request.TopicQuota) {
		query["TopicQuota"] = request.TopicQuota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradePostPayOrder"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradePostPayOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades a Message Queue for Apache Kafka instance that uses the subscription billing method.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of subscription Message Queue for Apache Kafka instances. For more information, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
//
// @param tmpReq - UpgradePrePayOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradePrePayOrderResponse
func (client *Client) UpgradePrePayOrderWithContext(ctx context.Context, tmpReq *UpgradePrePayOrderRequest, runtime *dara.RuntimeOptions) (_result *UpgradePrePayOrderResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpgradePrePayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfluentConfig) {
		request.ConfluentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfluentConfig, dara.String("ConfluentConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfluentConfigShrink) {
		query["ConfluentConfig"] = request.ConfluentConfigShrink
	}

	if !dara.IsNil(request.DiskSize) {
		query["DiskSize"] = request.DiskSize
	}

	if !dara.IsNil(request.EipMax) {
		query["EipMax"] = request.EipMax
	}

	if !dara.IsNil(request.EipModel) {
		query["EipModel"] = request.EipModel
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IoMax) {
		query["IoMax"] = request.IoMax
	}

	if !dara.IsNil(request.IoMaxSpec) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !dara.IsNil(request.PaidType) {
		query["PaidType"] = request.PaidType
	}

	if !dara.IsNil(request.PartitionNum) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.SpecType) {
		query["SpecType"] = request.SpecType
	}

	if !dara.IsNil(request.TopicQuota) {
		query["TopicQuota"] = request.TopicQuota
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradePrePayOrder"),
		Version:     dara.String("2019-09-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradePrePayOrderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
