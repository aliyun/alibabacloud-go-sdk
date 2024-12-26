// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataTopicLagMapValue struct {
	// Ready message count
	//
	// example:
	//
	// 1
	ReadyCount *int64 `json:"readyCount,omitempty" xml:"readyCount,omitempty"`
	// The number of messages being consumed.
	//
	// example:
	//
	// 1
	InflightCount *int64 `json:"inflightCount,omitempty" xml:"inflightCount,omitempty"`
	// Delivery delay time, in seconds
	//
	// example:
	//
	// 12
	DeliveryDuration *int64 `json:"deliveryDuration,omitempty" xml:"deliveryDuration,omitempty"`
}

func (s DataTopicLagMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataTopicLagMapValue) GoString() string {
	return s.String()
}

func (s *DataTopicLagMapValue) SetReadyCount(v int64) *DataTopicLagMapValue {
	s.ReadyCount = &v
	return s
}

func (s *DataTopicLagMapValue) SetInflightCount(v int64) *DataTopicLagMapValue {
	s.InflightCount = &v
	return s
}

func (s *DataTopicLagMapValue) SetDeliveryDuration(v int64) *DataTopicLagMapValue {
	s.DeliveryDuration = &v
	return s
}

type AddDisasterRecoveryItemRequest struct {
	Topics []*AddDisasterRecoveryItemRequestTopics `json:"topics,omitempty" xml:"topics,omitempty" type:"Repeated"`
}

func (s AddDisasterRecoveryItemRequest) String() string {
	return tea.Prettify(s)
}

func (s AddDisasterRecoveryItemRequest) GoString() string {
	return s.String()
}

func (s *AddDisasterRecoveryItemRequest) SetTopics(v []*AddDisasterRecoveryItemRequestTopics) *AddDisasterRecoveryItemRequest {
	s.Topics = v
	return s
}

type AddDisasterRecoveryItemRequestTopics struct {
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	InstanceId      *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// ALIYUN_ROCKETMQ
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// regionId
	//
	// example:
	//
	// cn-hangzhou
	RegionId  *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s AddDisasterRecoveryItemRequestTopics) String() string {
	return tea.Prettify(s)
}

func (s AddDisasterRecoveryItemRequestTopics) GoString() string {
	return s.String()
}

func (s *AddDisasterRecoveryItemRequestTopics) SetConsumerGroupId(v string) *AddDisasterRecoveryItemRequestTopics {
	s.ConsumerGroupId = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetInstanceId(v string) *AddDisasterRecoveryItemRequestTopics {
	s.InstanceId = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetInstanceType(v string) *AddDisasterRecoveryItemRequestTopics {
	s.InstanceType = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetRegionId(v string) *AddDisasterRecoveryItemRequestTopics {
	s.RegionId = &v
	return s
}

func (s *AddDisasterRecoveryItemRequestTopics) SetTopicName(v string) *AddDisasterRecoveryItemRequestTopics {
	s.TopicName = &v
	return s
}

type AddDisasterRecoveryItemResponseBody struct {
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	Code               *string `json:"code,omitempty" xml:"code,omitempty"`
	Data               *int64  `json:"data,omitempty" xml:"data,omitempty"`
	DynamicCode        *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	DynamicMessage     *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	HttpStatusCode     *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message            *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId          *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success            *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddDisasterRecoveryItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddDisasterRecoveryItemResponseBody) GoString() string {
	return s.String()
}

func (s *AddDisasterRecoveryItemResponseBody) SetAccessDeniedDetail(v string) *AddDisasterRecoveryItemResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetCode(v string) *AddDisasterRecoveryItemResponseBody {
	s.Code = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetData(v int64) *AddDisasterRecoveryItemResponseBody {
	s.Data = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetDynamicCode(v string) *AddDisasterRecoveryItemResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetDynamicMessage(v string) *AddDisasterRecoveryItemResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetHttpStatusCode(v int32) *AddDisasterRecoveryItemResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetMessage(v string) *AddDisasterRecoveryItemResponseBody {
	s.Message = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetRequestId(v string) *AddDisasterRecoveryItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDisasterRecoveryItemResponseBody) SetSuccess(v bool) *AddDisasterRecoveryItemResponseBody {
	s.Success = &v
	return s
}

type AddDisasterRecoveryItemResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDisasterRecoveryItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDisasterRecoveryItemResponse) String() string {
	return tea.Prettify(s)
}

func (s AddDisasterRecoveryItemResponse) GoString() string {
	return s.String()
}

func (s *AddDisasterRecoveryItemResponse) SetHeaders(v map[string]*string) *AddDisasterRecoveryItemResponse {
	s.Headers = v
	return s
}

func (s *AddDisasterRecoveryItemResponse) SetStatusCode(v int32) *AddDisasterRecoveryItemResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDisasterRecoveryItemResponse) SetBody(v *AddDisasterRecoveryItemResponseBody) *AddDisasterRecoveryItemResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the resource group to which the instance is changed.
	//
	// You can call the [ListResourceGroups](https://www.alibabacloud.com/help/resource-management/latest/listresourcegroups) operation to query existing resource groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The ID of the resource. Set this parameter to the ID of the ApsaraMQ for RocketMQ instance whose resource group you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// c2c5d1274a8d4317a13bc5b0d4******
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of resource.
	//
	// Set this parameter to **instance**. The value of this parameter cannot be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetCode(v string) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetData(v bool) *ChangeResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetDynamicCode(v string) *ChangeResourceGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetDynamicMessage(v string) *ChangeResourceGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetHttpStatusCode(v int32) *ChangeResourceGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetSuccess(v bool) *ChangeResourceGroupResponseBody {
	s.Success = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CreateConsumerGroupRequest struct {
	// The consumption retry policy that you want to configure for the consumer group. For more information, see [Consumption retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// This parameter is required.
	ConsumeRetryPolicy *CreateConsumerGroupRequestConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	// The message delivery order of the consumer group.
	//
	// Valid values:
	//
	// 	- Concurrently: concurrent delivery
	//
	// 	- Orderly: ordered delivery
	//
	// This parameter is required.
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	MaxReceiveTps     *int64  `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The remarks on the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) SetConsumeRetryPolicy(v *CreateConsumerGroupRequestConsumeRetryPolicy) *CreateConsumerGroupRequest {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *CreateConsumerGroupRequest) SetDeliveryOrderType(v string) *CreateConsumerGroupRequest {
	s.DeliveryOrderType = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetMaxReceiveTps(v int64) *CreateConsumerGroupRequest {
	s.MaxReceiveTps = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRemark(v string) *CreateConsumerGroupRequest {
	s.Remark = &v
	return s
}

type CreateConsumerGroupRequestConsumeRetryPolicy struct {
	// The dead-letter topic.
	//
	// If a consumer still fails to consume a message after the message is retried for a specified number of times, the message is delivered to a dead-letter topic for subsequent business recovery or troubleshooting. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 16
	MaxRetryTimes *int32 `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	// The retry policy. For more information, see [Message retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// Valid values:
	//
	// 	- FixedRetryPolicy: Failed messages are retried at a fixed interval.
	//
	// 	- DefaultRetryPolicy: Failed messages are retried at incremental intervals as the number of retries increases.
	//
	// This parameter is required.
	//
	// example:
	//
	// DefaultRetryPolicy
	RetryPolicy *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
}

func (s CreateConsumerGroupRequestConsumeRetryPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequestConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetDeadLetterTargetTopic(v string) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.DeadLetterTargetTopic = &v
	return s
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetMaxRetryTimes(v int32) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *CreateConsumerGroupRequestConsumeRetryPolicy) SetRetryPolicy(v string) *CreateConsumerGroupRequestConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

type CreateConsumerGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// consumerGroupId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter consumerGroupId is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) SetCode(v string) *CreateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetData(v bool) *CreateConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetDynamicCode(v string) *CreateConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetDynamicMessage(v string) *CreateConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetHttpStatusCode(v int32) *CreateConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetMessage(v string) *CreateConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetRequestId(v string) *CreateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetSuccess(v bool) *CreateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

type CreateConsumerGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponse) SetHeaders(v map[string]*string) *CreateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerGroupResponse) SetStatusCode(v int32) *CreateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConsumerGroupResponse) SetBody(v *CreateConsumerGroupResponseBody) *CreateConsumerGroupResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	// Specifies whether to enable auto-renewal for the instance. This parameter takes effect only if you set paymentType to Subscription. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// The auto-renewal cycle of the instance. This parameter takes effect only if you set autoRenew to true. Unit: months.
	//
	// Valid values:
	//
	// 	- Monthly renewal: 1, 2, 3, 6, and 12
	//
	// example:
	//
	// 3
	AutoRenewPeriod *int32 `json:"autoRenewPeriod,omitempty" xml:"autoRenewPeriod,omitempty"`
	// The commodity code. Valid values:
	//
	// 	- ons_rmqpost_public_intl: pay-as-you-go
	//
	// 	- ons_rmqsub_public_intl: subscription
	//
	// example:
	//
	// ons_ rmqpost_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The name of the instance that you want to create.
	//
	// If you leave this parameter empty, the instance ID is used as the instance name.
	//
	// example:
	//
	// rmq-cn-72u3048uxxx
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The information about the network.
	//
	// This parameter is required.
	NetworkInfo *CreateInstanceRequestNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// The billing method of the instance. ApsaraMQ for RocketMQ supports the subscription and pay-as-you-go billing methods.
	//
	// Valid values:
	//
	// 	- PayAsYouGo: This billing method allows you to use resources before you pay for the resources.
	//
	// 	- Subscription: This billing method allows you to use resources after you pay for the resources.
	//
	// For more information, see [Billing methods](https://help.aliyun.com/document_detail/427234.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The subscription duration of the instance. This parameter takes effect only if you set paymentType to Subscription.
	//
	// Valid values:
	//
	// 	- Monthly subscription: 1, 2, 3, 4, 5, and 6
	//
	// 	- Yearly subscription: 1, 2, and 3
	//
	// example:
	//
	// 3
	Period *int32 `json:"period,omitempty" xml:"period,omitempty"`
	// The unit of the subscription duration.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"periodUnit,omitempty" xml:"periodUnit,omitempty"`
	// The information about the instance specifications.
	ProductInfo *CreateInstanceRequestProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The instance description.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzy6pist7uuna
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary edition of the instance. For information about the differences among primary edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Enterprise Platinum Edition
	//
	// 	- professional: Professional Edition
	//
	// >  After you create an instance, you can only upgrade the primary edition of the instance. The following editions are sorted in ascending order: Standard Edition, Professional Edition, Enterprise Platinum Edition. For example, you can upgrade an instance from Standard Edition to Professional Edition, but you cannot downgrade an instance from Professional Edition to Standard Edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// standard
	SeriesCode *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	// The code of the service to which the instance belongs. The service code of ApsaraMQ for RocketMQ is rmq.
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// The sub-category edition of the instance. For information about the differences among sub-category edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- cluster_ha: Cluster High-availability Edition
	//
	// 	- single_node: Standalone Edition
	//
	// 	- serverless: serverless
	//
	// If you set seriesCode to ultimate, you can set this parameter only to cluster_ha.
	//
	// >  After you create an instance, you cannot change the sub-category edition of the instance.
	//
	// Valid values:
	//
	// 	- serverless: serverless
	//
	// 	- cluster_ha: Cluster High-availability Edition
	//
	// 	- single_node: Standalone Edition
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster_ha
	SubSeriesCode *string `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value of this parameter, but you must ensure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// c2c5d1274a8d4317a13bc5b0d4******
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int32) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetCommodityCode(v string) *CreateInstanceRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetNetworkInfo(v *CreateInstanceRequestNetworkInfo) *CreateInstanceRequest {
	s.NetworkInfo = v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodUnit(v string) *CreateInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetProductInfo(v *CreateInstanceRequestProductInfo) *CreateInstanceRequest {
	s.ProductInfo = v
	return s
}

func (s *CreateInstanceRequest) SetRemark(v string) *CreateInstanceRequest {
	s.Remark = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetSeriesCode(v string) *CreateInstanceRequest {
	s.SeriesCode = &v
	return s
}

func (s *CreateInstanceRequest) SetServiceCode(v string) *CreateInstanceRequest {
	s.ServiceCode = &v
	return s
}

func (s *CreateInstanceRequest) SetSubSeriesCode(v string) *CreateInstanceRequest {
	s.SubSeriesCode = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

type CreateInstanceRequestNetworkInfo struct {
	// The Internet-related configurations.
	//
	// This parameter is required.
	InternetInfo *CreateInstanceRequestNetworkInfoInternetInfo `json:"internetInfo,omitempty" xml:"internetInfo,omitempty" type:"Struct"`
	// The virtual private cloud (VPC)-related configurations.
	//
	// This parameter is required.
	VpcInfo *CreateInstanceRequestNetworkInfoVpcInfo `json:"vpcInfo,omitempty" xml:"vpcInfo,omitempty" type:"Struct"`
}

func (s CreateInstanceRequestNetworkInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestNetworkInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestNetworkInfo) SetInternetInfo(v *CreateInstanceRequestNetworkInfoInternetInfo) *CreateInstanceRequestNetworkInfo {
	s.InternetInfo = v
	return s
}

func (s *CreateInstanceRequestNetworkInfo) SetVpcInfo(v *CreateInstanceRequestNetworkInfoVpcInfo) *CreateInstanceRequestNetworkInfo {
	s.VpcInfo = v
	return s
}

type CreateInstanceRequestNetworkInfoInternetInfo struct {
	// The Internet bandwidth. Unit: MB/s.
	//
	// This parameter is required only if you set flowOutType to payByBandwidth.
	//
	// Valid values: 1 to 1000.
	//
	// example:
	//
	// 100
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The billing method of Internet usage.
	//
	// Valid values:
	//
	// 	- payByBandwidth: pay-by-bandwidth. This value is valid only if you enable Internet access.
	//
	// 	- payByTraffic: pay-by-traffic. This value is valid only if you enable Internet access.
	//
	// 	- uninvolved: No billing method is involved. This value is valid only if you disable Internet access.
	//
	// This parameter is required.
	//
	// example:
	//
	// uninvolved
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Specifies whether to enable the Internet access feature.
	//
	// Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// By default, ApsaraMQ for RocketMQ allows you to access instances in VPCs. If you enable Internet access for an instance, you can access the instance over the Internet. After you enable this feature, you are charged for outbound Internet traffic. For more information, see [Internet access fees](https://help.aliyun.com/document_detail/427240.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// disable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// Deprecated
	//
	// The whitelist that includes the IP addresses that are allowed to access the ApsaraMQ for RocketMQ broker over the Internet. This parameter can be configured only if you use the public endpoint to access the instance.
	//
	// 	- If you do not configure an IP address whitelist, all CIDR blocks are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If you configure an IP address whitelist, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	IpWhitelist []*string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
}

func (s CreateInstanceRequestNetworkInfoInternetInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestNetworkInfoInternetInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) SetFlowOutBandwidth(v int32) *CreateInstanceRequestNetworkInfoInternetInfo {
	s.FlowOutBandwidth = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) SetFlowOutType(v string) *CreateInstanceRequestNetworkInfoInternetInfo {
	s.FlowOutType = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) SetInternetSpec(v string) *CreateInstanceRequestNetworkInfoInternetInfo {
	s.InternetSpec = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoInternetInfo) SetIpWhitelist(v []*string) *CreateInstanceRequestNetworkInfoInternetInfo {
	s.IpWhitelist = v
	return s
}

type CreateInstanceRequestNetworkInfoVpcInfo struct {
	// The ID of the security group to which the instance belongs.
	//
	// example:
	//
	// sg-bp17hpmgz96tvnsdy6so
	SecurityGroupIds *string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty"`
	// Deprecated
	//
	// The ID of the vSwitch with which the instance is associated. If you want to specify multiple vSwitches, separate the vSwitches with vertical bars (|).
	//
	// >  After you create an ApsaraMQ for RocketMQ instance, you cannot change the vSwitch with which the instance is associated. If you want to change the vSwitch with which the instance is associated, you must release the instance and purchase a new instance.
	//
	// >  We recommend that you configure vSwitches instead of this parameter.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpv*******
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The vSwitches.
	//
	// >  After you create an ApsaraMQ for RocketMQ instance, you cannot change the vSwitch with which the instance is associated. If you want to change the vSwitch with which the instance is associated, you must release the instance and purchase a new instance.
	//
	// >  This parameter is required. We recommend that you configure this parameter instead of vSwitchId.
	VSwitches []*CreateInstanceRequestNetworkInfoVpcInfoVSwitches `json:"vSwitches,omitempty" xml:"vSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC with which the instance to be created is associated.
	//
	// >  After you create an ApsaraMQ for RocketMQ instance, you cannot change the VPC with which the instance is associated. If you want to change the VPC with which the instance is associated, you must release the instance and create a new instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-wz9qt50xhtj9krb******
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateInstanceRequestNetworkInfoVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestNetworkInfoVpcInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetSecurityGroupIds(v string) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.SecurityGroupIds = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetVSwitchId(v string) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetVSwitches(v []*CreateInstanceRequestNetworkInfoVpcInfoVSwitches) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.VSwitches = v
	return s
}

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetVpcId(v string) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.VpcId = &v
	return s
}

type CreateInstanceRequestNetworkInfoVpcInfoVSwitches struct {
	// The ID of the vSwitch with which the instance is associated.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpv*******
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
}

func (s CreateInstanceRequestNetworkInfoVpcInfoVSwitches) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestNetworkInfoVpcInfoVSwitches) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestNetworkInfoVpcInfoVSwitches) SetVSwitchId(v string) *CreateInstanceRequestNetworkInfoVpcInfoVSwitches {
	s.VSwitchId = &v
	return s
}

type CreateInstanceRequestProductInfo struct {
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// After you enable the elastic TPS feature for an ApsaraMQ for RocketMQ instance, you can use a specific amount of TPS that exceeds the specification limit. You are charged for the elastic TPS feature. For more information, see [Computing fees](https://help.aliyun.com/document_detail/427237.html).
	//
	// >  The elastic TPS feature is supported only by instances of specific editions. For more information, see [Instance editions](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- provisioned
	//
	// 	- ondemand
	//
	// example:
	//
	// provisioned
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// This parameter is no longer used. You do not need to configure this parameter.
	//
	// example:
	//
	// xxxx
	IntranetSpec *string `json:"intranetSpec,omitempty" xml:"intranetSpec,omitempty"`
	// The retention period of messages. Unit: hours.
	//
	// For information about the valid values of this parameter, see the "Limits on resource quotas" section of the [Limits](https://help.aliyun.com/document_detail/440347.html) topic.
	//
	// ApsaraMQ for RocketMQ supports serverless scaling of message storage. You are charged storage fees based on your actual storage usage. You can change the retention period of messages to manage storage capacity. For more information, see [Storage fees](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The computing specification that specifies the messaging transactions per second (TPS) of the instance. For more information, see [Instance editions](https://help.aliyun.com/document_detail/444715.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The ratio of the message sending TPS to the messaging TPS of the instance.
	//
	// For example, if the maximum messaging TPS of an instance is 1,000 and the ratio of the message sending TPS to the messaging TPS of the instance is 0.8, the maximum message sending TPS of the instance is 800 and the maximum message receiving TPS is 200.
	//
	// Valid values: 0 to 1. Default value: 0.5.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Indicates whether storage encryption is enabled.
	//
	// example:
	//
	// false
	StorageEncryption *bool `json:"storageEncryption,omitempty" xml:"storageEncryption,omitempty"`
	// The storage encryption key.
	//
	// example:
	//
	// xxx
	StorageSecretKey *string `json:"storageSecretKey,omitempty" xml:"storageSecretKey,omitempty"`
}

func (s CreateInstanceRequestProductInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequestProductInfo) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestProductInfo) SetAutoScaling(v bool) *CreateInstanceRequestProductInfo {
	s.AutoScaling = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetChargeType(v string) *CreateInstanceRequestProductInfo {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetIntranetSpec(v string) *CreateInstanceRequestProductInfo {
	s.IntranetSpec = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetMessageRetentionTime(v int32) *CreateInstanceRequestProductInfo {
	s.MessageRetentionTime = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetMsgProcessSpec(v string) *CreateInstanceRequestProductInfo {
	s.MsgProcessSpec = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetSendReceiveRatio(v float32) *CreateInstanceRequestProductInfo {
	s.SendReceiveRatio = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetStorageEncryption(v bool) *CreateInstanceRequestProductInfo {
	s.StorageEncryption = &v
	return s
}

func (s *CreateInstanceRequestProductInfo) SetStorageSecretKey(v string) *CreateInstanceRequestProductInfo {
	s.StorageSecretKey = &v
	return s
}

type CreateInstanceResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the created instance.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetCode(v string) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetData(v string) *CreateInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceResponseBody) SetDynamicCode(v string) *CreateInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetDynamicMessage(v string) *CreateInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInstanceResponseBody) SetHttpStatusCode(v int32) *CreateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceResponseBody) SetMessage(v string) *CreateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetSuccess(v bool) *CreateInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateInstanceAccountRequest struct {
	// The password of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountRequest) SetPassword(v string) *CreateInstanceAccountRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceAccountRequest) SetUsername(v string) *CreateInstanceAccountRequest {
	s.Username = &v
	return s
}

type CreateInstanceAccountResponseBody struct {
	// No permission details
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 3AE0999C-8DBA-5CEE-8D9A-BE8D4A90DF8D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponseBody) SetAccessDeniedDetail(v string) *CreateInstanceAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetCode(v string) *CreateInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetData(v bool) *CreateInstanceAccountResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetDynamicCode(v string) *CreateInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetDynamicMessage(v string) *CreateInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetHttpStatusCode(v int32) *CreateInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetMessage(v string) *CreateInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetRequestId(v string) *CreateInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetSuccess(v bool) *CreateInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponse) SetHeaders(v map[string]*string) *CreateInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceAccountResponse) SetStatusCode(v int32) *CreateInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceAccountResponse) SetBody(v *CreateInstanceAccountResponseBody) *CreateInstanceAccountResponse {
	s.Body = v
	return s
}

type CreateInstanceAclRequest struct {
	// The type of operations that can be performed on the resource.
	//
	// The following types of operations are supported based on the resource type:
	//
	// 	- Topic: Pub, Sub, and Pub|Sub
	//
	// 	- Consumer group: Sub
	//
	// Valid values:
	//
	// 	- SUB: subscribe
	//
	// 	- Pub|Sub: publish and subscribe
	//
	// 	- Pub: publish
	//
	// This parameter is required.
	//
	// example:
	//
	// Pub
	Actions *string `json:"actions,omitempty" xml:"actions,omitempty"`
	// The decision result of the authorization.
	//
	// Valid values:
	//
	// 	- Deny
	//
	// 	- Allow
	//
	// This parameter is required.
	//
	// example:
	//
	// Allow
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// The IP address whitelists.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
	// The name of the resource on which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The type of the resource on which you want to grant permissions.
	//
	// Valid values:
	//
	// 	- Group
	//
	// 	- Topic
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CreateInstanceAclRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceAclRequest) SetActions(v string) *CreateInstanceAclRequest {
	s.Actions = &v
	return s
}

func (s *CreateInstanceAclRequest) SetDecision(v string) *CreateInstanceAclRequest {
	s.Decision = &v
	return s
}

func (s *CreateInstanceAclRequest) SetIpWhitelists(v []*string) *CreateInstanceAclRequest {
	s.IpWhitelists = v
	return s
}

func (s *CreateInstanceAclRequest) SetResourceName(v string) *CreateInstanceAclRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateInstanceAclRequest) SetResourceType(v string) *CreateInstanceAclRequest {
	s.ResourceType = &v
	return s
}

type CreateInstanceAclResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7E8AE3A-219B-52EE-BE32-4036F5F88833
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInstanceAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceAclResponseBody) SetAccessDeniedDetail(v string) *CreateInstanceAclResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetCode(v string) *CreateInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetData(v bool) *CreateInstanceAclResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetDynamicCode(v string) *CreateInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetDynamicMessage(v string) *CreateInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetHttpStatusCode(v int32) *CreateInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetMessage(v string) *CreateInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetRequestId(v string) *CreateInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceAclResponseBody) SetSuccess(v bool) *CreateInstanceAclResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceAclResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceAclResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceAclResponse) SetHeaders(v map[string]*string) *CreateInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceAclResponse) SetStatusCode(v int32) *CreateInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceAclResponse) SetBody(v *CreateInstanceAclResponseBody) *CreateInstanceAclResponse {
	s.Body = v
	return s
}

type CreateInstanceIpWhitelistRequest struct {
	// The IP address whitelists.
	//
	// This parameter is required.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
}

func (s CreateInstanceIpWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceIpWhitelistRequest) SetIpWhitelists(v []*string) *CreateInstanceIpWhitelistRequest {
	s.IpWhitelists = v
	return s
}

type CreateInstanceIpWhitelistResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A07B41BD-6DD3-5349-9E76-00303DF04BBE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateInstanceIpWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceIpWhitelistResponseBody) SetAccessDeniedDetail(v string) *CreateInstanceIpWhitelistResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetCode(v string) *CreateInstanceIpWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetData(v bool) *CreateInstanceIpWhitelistResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetDynamicCode(v string) *CreateInstanceIpWhitelistResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetDynamicMessage(v string) *CreateInstanceIpWhitelistResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetHttpStatusCode(v int32) *CreateInstanceIpWhitelistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetMessage(v string) *CreateInstanceIpWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetRequestId(v string) *CreateInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponseBody) SetSuccess(v bool) *CreateInstanceIpWhitelistResponseBody {
	s.Success = &v
	return s
}

type CreateInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceIpWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *CreateInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceIpWhitelistResponse) SetStatusCode(v int32) *CreateInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceIpWhitelistResponse) SetBody(v *CreateInstanceIpWhitelistResponseBody) *CreateInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

type CreateTopicRequest struct {
	MaxSendTps *int64 `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// The type of messages in the topic that you want to create.
	//
	// Valid values:
	//
	// 	- TRANSACTION: transactional messages
	//
	// 	- FIFO: ordered messages
	//
	// 	- DELAY: scheduled messages or delayed Message
	//
	// 	- NORMAL: normal messages
	//
	// > The type of messages in the topic must be the same as the type of messages that you want to send. For example, if you create a topic whose message type is ordered messages, the topic can be used to send and receive only ordered messages.
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The description of the topic that you want to create.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) SetMaxSendTps(v int64) *CreateTopicRequest {
	s.MaxSendTps = &v
	return s
}

func (s *CreateTopicRequest) SetMessageType(v string) *CreateTopicRequest {
	s.MessageType = &v
	return s
}

func (s *CreateTopicRequest) SetRemark(v string) *CreateTopicRequest {
	s.Remark = &v
	return s
}

type CreateTopicResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// Topic.Existed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// topicName
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The topic already exists.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) SetCode(v string) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetData(v bool) *CreateTopicResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTopicResponseBody) SetDynamicCode(v string) *CreateTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *CreateTopicResponseBody) SetDynamicMessage(v string) *CreateTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *CreateTopicResponseBody) SetHttpStatusCode(v int32) *CreateTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTopicResponseBody) SetMessage(v string) *CreateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

type CreateTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponse) GoString() string {
	return s.String()
}

func (s *CreateTopicResponse) SetHeaders(v map[string]*string) *CreateTopicResponse {
	s.Headers = v
	return s
}

func (s *CreateTopicResponse) SetStatusCode(v int32) *CreateTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTopicResponse) SetBody(v *CreateTopicResponseBody) *CreateTopicResponse {
	s.Body = v
	return s
}

type DeleteConsumerGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// consumerGroupId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter consumerGroupId is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// C7F94090-3358-506A-97DC-34BC803C****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) SetCode(v string) *DeleteConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetData(v bool) *DeleteConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetDynamicCode(v string) *DeleteConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetDynamicMessage(v string) *DeleteConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetHttpStatusCode(v int32) *DeleteConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetMessage(v string) *DeleteConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetRequestId(v string) *DeleteConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetSuccess(v bool) *DeleteConsumerGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteConsumerGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponse) SetHeaders(v map[string]*string) *DeleteConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerGroupResponse) SetStatusCode(v int32) *DeleteConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConsumerGroupResponse) SetBody(v *DeleteConsumerGroupResponseBody) *DeleteConsumerGroupResponse {
	s.Body = v
	return s
}

type DeleteConsumerGroupSubscriptionRequest struct {
	// The filter expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// *
	FilterExpression *string `json:"filterExpression,omitempty" xml:"filterExpression,omitempty"`
	// The type of the filter expression. Valid values:
	//
	// 	- SQL: filters messages by using SQL expressions.
	//
	// 	- TAG: filters messages by using tags.
	//
	// Valid values:
	//
	// 	- TAG: filters messages by using SQL expressions.
	//
	// 	- SQL: filters messages by using SQL expressions.
	//
	// This parameter is required.
	//
	// example:
	//
	// TAG
	FilterType *string `json:"filterType,omitempty" xml:"filterType,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s DeleteConsumerGroupSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupSubscriptionRequest) SetFilterExpression(v string) *DeleteConsumerGroupSubscriptionRequest {
	s.FilterExpression = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionRequest) SetFilterType(v string) *DeleteConsumerGroupSubscriptionRequest {
	s.FilterType = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionRequest) SetTopicName(v string) *DeleteConsumerGroupSubscriptionRequest {
	s.TopicName = &v
	return s
}

type DeleteConsumerGroupSubscriptionResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied due to the reason that the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 723CDA5C-E25C-5EAF-9601-08C286DF8A4D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteConsumerGroupSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetAccessDeniedDetail(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetCode(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetData(v bool) *DeleteConsumerGroupSubscriptionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetDynamicCode(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetDynamicMessage(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetHttpStatusCode(v int32) *DeleteConsumerGroupSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetMessage(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetRequestId(v string) *DeleteConsumerGroupSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponseBody) SetSuccess(v bool) *DeleteConsumerGroupSubscriptionResponseBody {
	s.Success = &v
	return s
}

type DeleteConsumerGroupSubscriptionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConsumerGroupSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConsumerGroupSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupSubscriptionResponse) SetHeaders(v map[string]*string) *DeleteConsumerGroupSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponse) SetStatusCode(v int32) *DeleteConsumerGroupSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponse) SetBody(v *DeleteConsumerGroupSubscriptionResponseBody) *DeleteConsumerGroupSubscriptionResponse {
	s.Body = v
	return s
}

type DeleteInstanceResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 048242AA-BADA-5F29-B2CD-ED9FA344467F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetCode(v string) *DeleteInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetData(v bool) *DeleteInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetDynamicCode(v string) *DeleteInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetDynamicMessage(v string) *DeleteInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetMessage(v string) *DeleteInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetStatusCode(v int32) *DeleteInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
	s.Body = v
	return s
}

type DeleteInstanceAccountResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 157DF7D4-53FB-58C6-BEBC-A9400E7EF68A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAccountResponseBody) SetAccessDeniedDetail(v string) *DeleteInstanceAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetCode(v string) *DeleteInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetData(v bool) *DeleteInstanceAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetDynamicCode(v string) *DeleteInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetDynamicMessage(v string) *DeleteInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetMessage(v string) *DeleteInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetRequestId(v string) *DeleteInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceAccountResponseBody) SetSuccess(v bool) *DeleteInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAccountResponse) SetHeaders(v map[string]*string) *DeleteInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceAccountResponse) SetStatusCode(v int32) *DeleteInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceAccountResponse) SetBody(v *DeleteInstanceAccountResponseBody) *DeleteInstanceAccountResponse {
	s.Body = v
	return s
}

type DeleteInstanceAclRequest struct {
	// The name of the resource on which the permissions are granted.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The type of the resource on which the permissions are granted.
	//
	// Valid values:
	//
	// 	- Group
	//
	// 	- Topic
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s DeleteInstanceAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAclRequest) SetResourceName(v string) *DeleteInstanceAclRequest {
	s.ResourceName = &v
	return s
}

func (s *DeleteInstanceAclRequest) SetResourceType(v string) *DeleteInstanceAclRequest {
	s.ResourceType = &v
	return s
}

type DeleteInstanceAclResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied due to the reason that the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7358418D-83BD-507A-8079-611C63E05674
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteInstanceAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAclResponseBody) SetAccessDeniedDetail(v string) *DeleteInstanceAclResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetCode(v string) *DeleteInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetData(v bool) *DeleteInstanceAclResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetDynamicCode(v string) *DeleteInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetDynamicMessage(v string) *DeleteInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetMessage(v string) *DeleteInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetRequestId(v string) *DeleteInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceAclResponseBody) SetSuccess(v bool) *DeleteInstanceAclResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceAclResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceAclResponse) SetHeaders(v map[string]*string) *DeleteInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceAclResponse) SetStatusCode(v int32) *DeleteInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceAclResponse) SetBody(v *DeleteInstanceAclResponseBody) *DeleteInstanceAclResponse {
	s.Body = v
	return s
}

type DeleteInstanceIpWhitelistRequest struct {
	// The IP address whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	IpWhitelist *string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty"`
}

func (s DeleteInstanceIpWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceIpWhitelistRequest) SetIpWhitelist(v string) *DeleteInstanceIpWhitelistRequest {
	s.IpWhitelist = &v
	return s
}

type DeleteInstanceIpWhitelistResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied due to the reason that the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16425867-C948-5A0C-9A24-5259727BE727
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteInstanceIpWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetAccessDeniedDetail(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetCode(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetData(v bool) *DeleteInstanceIpWhitelistResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetDynamicCode(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetDynamicMessage(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceIpWhitelistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetMessage(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetRequestId(v string) *DeleteInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponseBody) SetSuccess(v bool) *DeleteInstanceIpWhitelistResponseBody {
	s.Success = &v
	return s
}

type DeleteInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceIpWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *DeleteInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceIpWhitelistResponse) SetStatusCode(v int32) *DeleteInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceIpWhitelistResponse) SetBody(v *DeleteInstanceIpWhitelistResponseBody) *DeleteInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

type DeleteTopicResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// topicName
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The topic cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) SetCode(v string) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetData(v bool) *DeleteTopicResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTopicResponseBody) SetDynamicCode(v string) *DeleteTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DeleteTopicResponseBody) SetDynamicMessage(v string) *DeleteTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteTopicResponseBody) SetHttpStatusCode(v int32) *DeleteTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTopicResponseBody) SetMessage(v string) *DeleteTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTopicResponseBody) SetRequestId(v string) *DeleteTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

type DeleteTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponse) SetHeaders(v map[string]*string) *DeleteTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteTopicResponse) SetStatusCode(v int32) *DeleteTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTopicResponse) SetBody(v *DeleteTopicResponseBody) *DeleteTopicResponse {
	s.Body = v
	return s
}

type GetConsumerGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	Data *GetConsumerGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// consumerGroupId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter consumerGroupId is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// C7F94090-3358-506A-97DC-34BC803C****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBody) SetCode(v string) *GetConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetData(v *GetConsumerGroupResponseBodyData) *GetConsumerGroupResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerGroupResponseBody) SetDynamicCode(v string) *GetConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetDynamicMessage(v string) *GetConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetHttpStatusCode(v int32) *GetConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetMessage(v string) *GetConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetRequestId(v string) *GetConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerGroupResponseBody) SetSuccess(v bool) *GetConsumerGroupResponseBody {
	s.Success = &v
	return s
}

type GetConsumerGroupResponseBodyData struct {
	// The consumption retry policy that you want to configure for the consumer group. For more information, see [Consumption retry](https://help.aliyun.com/document_detail/440356.html).
	ConsumeRetryPolicy *GetConsumerGroupResponseBodyDataConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	// The ID of the consumer group.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The time when the consumer group was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The message delivery order of the consumer group.
	//
	// Valid values:
	//
	// 	- Concurrently
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     concurrent delivery
	//
	//     <!-- -->
	//
	// 	- Orderly
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     ordered delivery
	//
	//     <!-- -->
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The remarks on the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The state of the consumer group.
	//
	// Valid values:
	//
	// 	- RUNNING
	//
	//     <!-- -->
	//
	//     : The consumer group is
	//
	//     <!-- -->
	//
	//     running
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- CREATING
	//
	//     <!-- -->
	//
	//     : The consumer group is
	//
	//     <!-- -->
	//
	//     being created
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the consumer group was last updated.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetConsumerGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBodyData) SetConsumeRetryPolicy(v *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) *GetConsumerGroupResponseBodyData {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetConsumerGroupId(v string) *GetConsumerGroupResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetCreateTime(v string) *GetConsumerGroupResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetDeliveryOrderType(v string) *GetConsumerGroupResponseBodyData {
	s.DeliveryOrderType = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetInstanceId(v string) *GetConsumerGroupResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetRegionId(v string) *GetConsumerGroupResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetRemark(v string) *GetConsumerGroupResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetStatus(v string) *GetConsumerGroupResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetConsumerGroupResponseBodyData) SetUpdateTime(v string) *GetConsumerGroupResponseBodyData {
	s.UpdateTime = &v
	return s
}

type GetConsumerGroupResponseBodyDataConsumeRetryPolicy struct {
	// The dead-letter topic.
	//
	// If a consumer still fails to consume a message after the message is retried for a specified number of times, the message is delivered to a dead-letter topic for subsequent business recovery or troubleshooting. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 16
	MaxRetryTimes *int32 `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	// The retry policy.
	//
	// Valid values:
	//
	// 	- FixedRetryPolicy
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Failed messages are retried at a fixed interval
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- DefaultRetryPolicy
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     Failed messages are retried at incremental intervals as the number of retries increases
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// DefaultRetryPolicy
	RetryPolicy *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
}

func (s GetConsumerGroupResponseBodyDataConsumeRetryPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupResponseBodyDataConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetDeadLetterTargetTopic(v string) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.DeadLetterTargetTopic = &v
	return s
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetMaxRetryTimes(v int32) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *GetConsumerGroupResponseBodyDataConsumeRetryPolicy) SetRetryPolicy(v string) *GetConsumerGroupResponseBodyDataConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

type GetConsumerGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupResponse) SetHeaders(v map[string]*string) *GetConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerGroupResponse) SetStatusCode(v int32) *GetConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerGroupResponse) SetBody(v *GetConsumerGroupResponseBody) *GetConsumerGroupResponse {
	s.Body = v
	return s
}

type GetConsumerGroupLagResponseBody struct {
	// Error code
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetConsumerGroupLagResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5764C40-FB8C-53AE-B95D-96AB3D0E9375
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerGroupLagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupLagResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagResponseBody) SetCode(v string) *GetConsumerGroupLagResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetData(v *GetConsumerGroupLagResponseBodyData) *GetConsumerGroupLagResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetDynamicCode(v string) *GetConsumerGroupLagResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetDynamicMessage(v string) *GetConsumerGroupLagResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetHttpStatusCode(v int32) *GetConsumerGroupLagResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetMessage(v string) *GetConsumerGroupLagResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetRequestId(v string) *GetConsumerGroupLagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerGroupLagResponseBody) SetSuccess(v bool) *GetConsumerGroupLagResponseBody {
	s.Success = &v
	return s
}

type GetConsumerGroupLagResponseBodyData struct {
	// Consumer Group ID
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// Instance ID
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Backlog for each topic
	TopicLagMap map[string]*DataTopicLagMapValue `json:"topicLagMap,omitempty" xml:"topicLagMap,omitempty"`
	// Total lag count
	TotalLag *GetConsumerGroupLagResponseBodyDataTotalLag `json:"totalLag,omitempty" xml:"totalLag,omitempty" type:"Struct"`
}

func (s GetConsumerGroupLagResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupLagResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagResponseBodyData) SetConsumerGroupId(v string) *GetConsumerGroupLagResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) SetInstanceId(v string) *GetConsumerGroupLagResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) SetRegionId(v string) *GetConsumerGroupLagResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) SetTopicLagMap(v map[string]*DataTopicLagMapValue) *GetConsumerGroupLagResponseBodyData {
	s.TopicLagMap = v
	return s
}

func (s *GetConsumerGroupLagResponseBodyData) SetTotalLag(v *GetConsumerGroupLagResponseBodyDataTotalLag) *GetConsumerGroupLagResponseBodyData {
	s.TotalLag = v
	return s
}

type GetConsumerGroupLagResponseBodyDataTotalLag struct {
	// Delivery delay time, in seconds
	//
	// example:
	//
	// 12
	DeliveryDuration *int64 `json:"deliveryDuration,omitempty" xml:"deliveryDuration,omitempty"`
	// The number of messages being consumed.
	//
	// example:
	//
	// 1
	InflightCount        *int64 `json:"inflightCount,omitempty" xml:"inflightCount,omitempty"`
	LastConsumeTimestamp *int64 `json:"lastConsumeTimestamp,omitempty" xml:"lastConsumeTimestamp,omitempty"`
	// Ready message count
	//
	// example:
	//
	// 1
	ReadyCount *int64 `json:"readyCount,omitempty" xml:"readyCount,omitempty"`
}

func (s GetConsumerGroupLagResponseBodyDataTotalLag) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupLagResponseBodyDataTotalLag) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) SetDeliveryDuration(v int64) *GetConsumerGroupLagResponseBodyDataTotalLag {
	s.DeliveryDuration = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) SetInflightCount(v int64) *GetConsumerGroupLagResponseBodyDataTotalLag {
	s.InflightCount = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) SetLastConsumeTimestamp(v int64) *GetConsumerGroupLagResponseBodyDataTotalLag {
	s.LastConsumeTimestamp = &v
	return s
}

func (s *GetConsumerGroupLagResponseBodyDataTotalLag) SetReadyCount(v int64) *GetConsumerGroupLagResponseBodyDataTotalLag {
	s.ReadyCount = &v
	return s
}

type GetConsumerGroupLagResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerGroupLagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerGroupLagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupLagResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupLagResponse) SetHeaders(v map[string]*string) *GetConsumerGroupLagResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerGroupLagResponse) SetStatusCode(v int32) *GetConsumerGroupLagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerGroupLagResponse) SetBody(v *GetConsumerGroupLagResponseBody) *GetConsumerGroupLagResponse {
	s.Body = v
	return s
}

type GetConsumerGroupSubscriptionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data []*GetConsumerGroupSubscriptionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 157DF7D4-53FB-58C6-BEBC-A9400E7EF68A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerGroupSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetCode(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetData(v []*GetConsumerGroupSubscriptionResponseBodyData) *GetConsumerGroupSubscriptionResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetDynamicCode(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetDynamicMessage(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetHttpStatusCode(v int32) *GetConsumerGroupSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetMessage(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetRequestId(v string) *GetConsumerGroupSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBody) SetSuccess(v bool) *GetConsumerGroupSubscriptionResponseBody {
	s.Success = &v
	return s
}

type GetConsumerGroupSubscriptionResponseBodyData struct {
	// The connection details.
	ConnectionDTO *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO `json:"connectionDTO,omitempty" xml:"connectionDTO,omitempty" type:"Struct"`
	// The subscription details.
	SubscriptionDTO *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO `json:"subscriptionDTO,omitempty" xml:"subscriptionDTO,omitempty" type:"Struct"`
}

func (s GetConsumerGroupSubscriptionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponseBodyData) SetConnectionDTO(v *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) *GetConsumerGroupSubscriptionResponseBodyData {
	s.ConnectionDTO = v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyData) SetSubscriptionDTO(v *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) *GetConsumerGroupSubscriptionResponseBodyData {
	s.SubscriptionDTO = v
	return s
}

type GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO struct {
	// The client ID.
	//
	// example:
	//
	// 192.168.50.191@19908#-2093249153#1534215565#40385215750900
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The public IP address of the host.
	//
	// example:
	//
	// xx.xx.xx.xx
	EgressIp *string `json:"egressIp,omitempty" xml:"egressIp,omitempty"`
	// The host name.
	//
	// example:
	//
	// nginx
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The language used by the client.
	//
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The consumption mode of the consumer group. Valid values:
	//
	// 	- BROADCASTING: broadcasting consumption
	//
	// 	- CLUSTERING: clustering consumption
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The client version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetClientId(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.ClientId = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetEgressIp(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.EgressIp = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetHostname(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.Hostname = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetLanguage(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.Language = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetMessageModel(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.MessageModel = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO) SetVersion(v string) *GetConsumerGroupSubscriptionResponseBodyDataConnectionDTO {
	s.Version = &v
	return s
}

type GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO struct {
	// The consumer group ID.
	//
	// example:
	//
	// GID_inspector_group
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The filter expression.
	//
	// example:
	//
	// *
	FilterExpression *string `json:"filterExpression,omitempty" xml:"filterExpression,omitempty"`
	// The type of the filter expression. Valid values:
	//
	// 	- SQL: filters messages by using SQL expressions.
	//
	// 	- TAG: filters messages by using tags.
	//
	// example:
	//
	// UNSPECIFIED
	FilterExpressionType *string `json:"filterExpressionType,omitempty" xml:"filterExpressionType,omitempty"`
	// The consumption mode of the consumer group. Valid values:
	//
	// 	- BROADCASTING: broadcasting consumption
	//
	// 	- CLUSTERING: clustering consumption
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The subscription status. Valid values:
	//
	// 	- ONLINE: The consumer group is online. If the consumer group contains multiple consumers, this value is returned if at least one of the consumers is online.
	//
	// 	- OFFLINE: The consumer group is offline. If the consumer group contains multiple consumers, this value is returned only if all consumers are offline.
	//
	// example:
	//
	// ONLINE
	SubscriptionStatus *string `json:"subscriptionStatus,omitempty" xml:"subscriptionStatus,omitempty"`
	// The topic to which the consumer group subscribes.
	//
	// example:
	//
	// Topic_normal_inspector
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetConsumerGroupId(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetFilterExpression(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.FilterExpression = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetFilterExpressionType(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.FilterExpressionType = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetMessageModel(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.MessageModel = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetSubscriptionStatus(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.SubscriptionStatus = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO) SetTopicName(v string) *GetConsumerGroupSubscriptionResponseBodyDataSubscriptionDTO {
	s.TopicName = &v
	return s
}

type GetConsumerGroupSubscriptionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerGroupSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerGroupSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponse) SetHeaders(v map[string]*string) *GetConsumerGroupSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerGroupSubscriptionResponse) SetStatusCode(v int32) *GetConsumerGroupSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponse) SetBody(v *GetConsumerGroupSubscriptionResponseBody) *GetConsumerGroupSubscriptionResponse {
	s.Body = v
	return s
}

type GetConsumerStackRequest struct {
	// The client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.26.76.48@Lqd7dImlp9KJ5V84
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
}

func (s GetConsumerStackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStackRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerStackRequest) SetClientId(v string) *GetConsumerStackRequest {
	s.ClientId = &v
	return s
}

type GetConsumerStackResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetConsumerStackResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30F2CBC7-F69D-5D78-9661-0254C9E1FBFA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetConsumerStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerStackResponseBody) SetCode(v string) *GetConsumerStackResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetData(v *GetConsumerStackResponseBodyData) *GetConsumerStackResponseBody {
	s.Data = v
	return s
}

func (s *GetConsumerStackResponseBody) SetDynamicCode(v string) *GetConsumerStackResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetDynamicMessage(v string) *GetConsumerStackResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetHttpStatusCode(v int32) *GetConsumerStackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetMessage(v string) *GetConsumerStackResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetRequestId(v string) *GetConsumerStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerStackResponseBody) SetSuccess(v bool) *GetConsumerStackResponseBody {
	s.Success = &v
	return s
}

type GetConsumerStackResponseBodyData struct {
	// The ID of the consumer group.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Stack Information.
	Stacks []*GetConsumerStackResponseBodyDataStacks `json:"stacks,omitempty" xml:"stacks,omitempty" type:"Repeated"`
}

func (s GetConsumerStackResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStackResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetConsumerStackResponseBodyData) SetConsumerGroupId(v string) *GetConsumerStackResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetConsumerStackResponseBodyData) SetInstanceId(v string) *GetConsumerStackResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerStackResponseBodyData) SetRegionId(v string) *GetConsumerStackResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetConsumerStackResponseBodyData) SetStacks(v []*GetConsumerStackResponseBodyDataStacks) *GetConsumerStackResponseBodyData {
	s.Stacks = v
	return s
}

type GetConsumerStackResponseBodyDataStacks struct {
	// Thread id.
	//
	// example:
	//
	// 123
	Thread *string `json:"thread,omitempty" xml:"thread,omitempty"`
	// Stack Information.
	Tracks []*string `json:"tracks,omitempty" xml:"tracks,omitempty" type:"Repeated"`
}

func (s GetConsumerStackResponseBodyDataStacks) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStackResponseBodyDataStacks) GoString() string {
	return s.String()
}

func (s *GetConsumerStackResponseBodyDataStacks) SetThread(v string) *GetConsumerStackResponseBodyDataStacks {
	s.Thread = &v
	return s
}

func (s *GetConsumerStackResponseBodyDataStacks) SetTracks(v []*string) *GetConsumerStackResponseBodyDataStacks {
	s.Tracks = v
	return s
}

type GetConsumerStackResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerStackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerStackResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerStackResponse) SetHeaders(v map[string]*string) *GetConsumerStackResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerStackResponse) SetStatusCode(v int32) *GetConsumerStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerStackResponse) SetBody(v *GetConsumerStackResponseBody) *GetConsumerStackResponse {
	s.Body = v
	return s
}

type GetInstanceResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 05AB7FBD-F1D3-5D87-BF78-BD782249****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetDynamicCode(v string) *GetInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetDynamicMessage(v string) *GetInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetInstanceResponseBody) SetHttpStatusCode(v int32) *GetInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetSuccess(v bool) *GetInstanceResponseBody {
	s.Success = &v
	return s
}

type GetInstanceResponseBodyData struct {
	// The account information.
	AccountInfo *GetInstanceResponseBodyDataAccountInfo `json:"accountInfo,omitempty" xml:"accountInfo,omitempty" type:"Struct"`
	// The information about access control.
	AclInfo *GetInstanceResponseBodyDataAclInfo `json:"aclInfo,omitempty" xml:"aclInfo,omitempty" type:"Struct"`
	// The business ID (BID) of the commodity.
	//
	// example:
	//
	// 26842
	Bid *string `json:"bid,omitempty" xml:"bid,omitempty"`
	// The commodity code of the instance. The commodity code of a ApsaraMQ for RocketMQ 5.0 instance has a similar format as ons_rmqsub_public_cn.
	//
	// example:
	//
	// ons_rmqsub_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2022-09-01 00:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// Deprecated
	//
	// The extended configurations. We recommend you configure productInfo, internetInfo, or aclInfo instead of this parameter.
	ExtConfig *GetInstanceResponseBodyDataExtConfig `json:"extConfig,omitempty" xml:"extConfig,omitempty" type:"Struct"`
	// The number of groups.
	//
	// example:
	//
	// 10
	GroupCount *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The instance quotas.
	InstanceQuotas []*GetInstanceResponseBodyDataInstanceQuotas `json:"instanceQuotas,omitempty" xml:"instanceQuotas,omitempty" type:"Repeated"`
	// The network information.
	NetworkInfo *GetInstanceResponseBodyDataNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// The billing method of the instance.
	//
	// Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Subscription
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The extended configurations of the instance.
	ProductInfo *GetInstanceResponseBodyDataProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The time when the instance was released.
	//
	// example:
	//
	// 2022-09-07 00:00:00
	ReleaseTime *string `json:"releaseTime,omitempty" xml:"releaseTime,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// This is remark for instance.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3tmjruyribi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary edition of the instance. For information about the differences between primary edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Enterprise Platinum Edition
	//
	// 	- professional: Professional Edition
	//
	// example:
	//
	// standard
	SeriesCode *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	// The code of the service to which the instance belongs. The service code of ApsaraMQ for RocketMQ is rmq.
	//
	// example:
	//
	// rmq
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// The instance software information.
	Software *GetInstanceResponseBodyDataSoftware `json:"software,omitempty" xml:"software,omitempty" type:"Struct"`
	// The time when the instance was started.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// 	- RELEASED
	//
	// 	- RUNNING
	//
	// 	- STOPPED
	//
	// 	- CHANGING
	//
	// 	- CREATING
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The sub-category edition of the instance. For information about the differences between sub-category edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- cluster_ha: Cluster High-availability Edition
	//
	// 	- single_node: Standalone Edition
	//
	// example:
	//
	// cluster_ha
	SubSeriesCode *string `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	// The resource tags.
	Tags []*GetInstanceResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The number of topics.
	//
	// example:
	//
	// 10
	TopicCount *int64 `json:"topicCount,omitempty" xml:"topicCount,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 2022-08-02 00:00:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the user who owns the instance.
	//
	// example:
	//
	// 111111111111111
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) SetAccountInfo(v *GetInstanceResponseBodyDataAccountInfo) *GetInstanceResponseBodyData {
	s.AccountInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetAclInfo(v *GetInstanceResponseBodyDataAclInfo) *GetInstanceResponseBodyData {
	s.AclInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetBid(v string) *GetInstanceResponseBodyData {
	s.Bid = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCommodityCode(v string) *GetInstanceResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetCreateTime(v string) *GetInstanceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetExpireTime(v string) *GetInstanceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetExtConfig(v *GetInstanceResponseBodyDataExtConfig) *GetInstanceResponseBodyData {
	s.ExtConfig = v
	return s
}

func (s *GetInstanceResponseBodyData) SetGroupCount(v int64) *GetInstanceResponseBodyData {
	s.GroupCount = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceId(v string) *GetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceName(v string) *GetInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceQuotas(v []*GetInstanceResponseBodyDataInstanceQuotas) *GetInstanceResponseBodyData {
	s.InstanceQuotas = v
	return s
}

func (s *GetInstanceResponseBodyData) SetNetworkInfo(v *GetInstanceResponseBodyDataNetworkInfo) *GetInstanceResponseBodyData {
	s.NetworkInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetPaymentType(v string) *GetInstanceResponseBodyData {
	s.PaymentType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetProductInfo(v *GetInstanceResponseBodyDataProductInfo) *GetInstanceResponseBodyData {
	s.ProductInfo = v
	return s
}

func (s *GetInstanceResponseBodyData) SetRegionId(v string) *GetInstanceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetReleaseTime(v string) *GetInstanceResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetRemark(v string) *GetInstanceResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetResourceGroupId(v string) *GetInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSeriesCode(v string) *GetInstanceResponseBodyData {
	s.SeriesCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetServiceCode(v string) *GetInstanceResponseBodyData {
	s.ServiceCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSoftware(v *GetInstanceResponseBodyDataSoftware) *GetInstanceResponseBodyData {
	s.Software = v
	return s
}

func (s *GetInstanceResponseBodyData) SetStartTime(v string) *GetInstanceResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSubSeriesCode(v string) *GetInstanceResponseBodyData {
	s.SubSeriesCode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetTags(v []*GetInstanceResponseBodyDataTags) *GetInstanceResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBodyData) SetTopicCount(v int64) *GetInstanceResponseBodyData {
	s.TopicCount = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUpdateTime(v string) *GetInstanceResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetUserId(v string) *GetInstanceResponseBodyData {
	s.UserId = &v
	return s
}

type GetInstanceResponseBodyDataAccountInfo struct {
	// The username of the instance. If you access a ApsaraMQ for RocketMQ instance over the Internet, you must configure the username and password of the instance in the SDK code for authentication.
	//
	// example:
	//
	// 6W0xz2uPfiwp****
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetInstanceResponseBodyDataAccountInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataAccountInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataAccountInfo) SetUsername(v string) *GetInstanceResponseBodyDataAccountInfo {
	s.Username = &v
	return s
}

type GetInstanceResponseBodyDataAclInfo struct {
	// Deprecated
	//
	// The authentication type of the instance. This parameter is no longer in use. We recommend that you configure aclTypes.
	//
	// Valid values:
	//
	// - default: intelligent identity authentication
	//
	// - apache_acl:access control list (ACL) identity authentication**
	//
	// example:
	//
	// default
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// The authentication types of the instance.
	AclTypes []*string `json:"aclTypes,omitempty" xml:"aclTypes,omitempty" type:"Repeated"`
	// Indicates whether the authentication-free in VPCs feature is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	DefaultVpcAuthFree *bool `json:"defaultVpcAuthFree,omitempty" xml:"defaultVpcAuthFree,omitempty"`
}

func (s GetInstanceResponseBodyDataAclInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataAclInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataAclInfo) SetAclType(v string) *GetInstanceResponseBodyDataAclInfo {
	s.AclType = &v
	return s
}

func (s *GetInstanceResponseBodyDataAclInfo) SetAclTypes(v []*string) *GetInstanceResponseBodyDataAclInfo {
	s.AclTypes = v
	return s
}

func (s *GetInstanceResponseBodyDataAclInfo) SetDefaultVpcAuthFree(v bool) *GetInstanceResponseBodyDataAclInfo {
	s.DefaultVpcAuthFree = &v
	return s
}

type GetInstanceResponseBodyDataExtConfig struct {
	// The authentication type of the instance.
	//
	// Valid value:
	//
	// 	- default: intelligent authentication
	//
	// example:
	//
	// default
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// This parameter is valid only when the supportAutoScaling parameter is set to enable.
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The Internet bandwidth. Unit: MB/s.
	//
	// example:
	//
	// 10
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The metering method of Internet usage.
	//
	// Valid values:
	//
	// 	- PayByTraffic: pay-by-traffic
	//
	// 	- paybybandwidth: pay-by-bandwidth
	//
	// 	- uninvolved: N/A
	//
	// example:
	//
	// payByBandwidth
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Indicates whether Internet access is enabled.
	//
	// Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// By default, you can access ApsaraMQ for RocketMQ instances in virtual private clouds (VPCs). If you enable the Internet access feature, you are charged for Internet outbound bandwidth. For more information, see [Internet access fees](https://help.aliyun.com/document_detail/427240.html).
	//
	// example:
	//
	// enable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// The retention period of messages. Unit: hours.
	//
	// For information about the valid values of this parameter, see the "Limits on resource quotas" section in [Usage limits](https://help.aliyun.com/document_detail/440347.html).
	//
	// The storage of messages in ApsaraMQ for RocketMQ is serverless and scalable. You are charged for message storage based on your actual usage. You can change the retention period of messages to adjust storage capacity. For more information, see [Storage fee](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The computing specification that is used to send and receive messages. For information about the upper limit of TPS, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The ratio between sent messages and received messages in the instance.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Specifies whether the elastic TPS feature is supported by the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// After you enable the elastic TPS feature for a ApsaraMQ for RocketMQ instance, you can use a specific amount of TPS that exceeds the specification limit. You are charged for the elastic TPS feature. For more information, see [Computing fee](https://help.aliyun.com/document_detail/427237.html).
	//
	// > The elastic TPS feature is supported only by specific instance editions. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	SupportAutoScaling *bool `json:"supportAutoScaling,omitempty" xml:"supportAutoScaling,omitempty"`
}

func (s GetInstanceResponseBodyDataExtConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataExtConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataExtConfig) SetAclType(v string) *GetInstanceResponseBodyDataExtConfig {
	s.AclType = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetAutoScaling(v bool) *GetInstanceResponseBodyDataExtConfig {
	s.AutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetFlowOutBandwidth(v int32) *GetInstanceResponseBodyDataExtConfig {
	s.FlowOutBandwidth = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetFlowOutType(v string) *GetInstanceResponseBodyDataExtConfig {
	s.FlowOutType = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetInternetSpec(v string) *GetInstanceResponseBodyDataExtConfig {
	s.InternetSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetMessageRetentionTime(v int32) *GetInstanceResponseBodyDataExtConfig {
	s.MessageRetentionTime = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetMsgProcessSpec(v string) *GetInstanceResponseBodyDataExtConfig {
	s.MsgProcessSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetSendReceiveRatio(v float32) *GetInstanceResponseBodyDataExtConfig {
	s.SendReceiveRatio = &v
	return s
}

func (s *GetInstanceResponseBodyDataExtConfig) SetSupportAutoScaling(v bool) *GetInstanceResponseBodyDataExtConfig {
	s.SupportAutoScaling = &v
	return s
}

type GetInstanceResponseBodyDataInstanceQuotas struct {
	// The number of topics that are free of charge on the instance.
	//
	// example:
	//
	// 20
	FreeCount *float64 `json:"freeCount,omitempty" xml:"freeCount,omitempty"`
	// The quota name.
	//
	// Valid value:
	//
	// 	- TOPIC_COUNT: the number of topics that can be created on the instance
	//
	// example:
	//
	// TOPIC_COUNT
	QuotaName *string `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	// The total number of topics on the instance.
	//
	// example:
	//
	// 100
	TotalCount *float64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The number of used topics on the instance.
	//
	// example:
	//
	// 10
	UsedCount *float64 `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
}

func (s GetInstanceResponseBodyDataInstanceQuotas) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataInstanceQuotas) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetFreeCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.FreeCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetQuotaName(v string) *GetInstanceResponseBodyDataInstanceQuotas {
	s.QuotaName = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetTotalCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceResponseBodyDataInstanceQuotas) SetUsedCount(v float64) *GetInstanceResponseBodyDataInstanceQuotas {
	s.UsedCount = &v
	return s
}

type GetInstanceResponseBodyDataNetworkInfo struct {
	// The endpoints.
	Endpoints []*GetInstanceResponseBodyDataNetworkInfoEndpoints `json:"endpoints,omitempty" xml:"endpoints,omitempty" type:"Repeated"`
	// The information about the Internet.
	InternetInfo *GetInstanceResponseBodyDataNetworkInfoInternetInfo `json:"internetInfo,omitempty" xml:"internetInfo,omitempty" type:"Struct"`
	// The virtual private cloud (VPC) information.
	VpcInfo *GetInstanceResponseBodyDataNetworkInfoVpcInfo `json:"vpcInfo,omitempty" xml:"vpcInfo,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBodyDataNetworkInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfo) SetEndpoints(v []*GetInstanceResponseBodyDataNetworkInfoEndpoints) *GetInstanceResponseBodyDataNetworkInfo {
	s.Endpoints = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfo) SetInternetInfo(v *GetInstanceResponseBodyDataNetworkInfoInternetInfo) *GetInstanceResponseBodyDataNetworkInfo {
	s.InternetInfo = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfo) SetVpcInfo(v *GetInstanceResponseBodyDataNetworkInfoVpcInfo) *GetInstanceResponseBodyDataNetworkInfo {
	s.VpcInfo = v
	return s
}

type GetInstanceResponseBodyDataNetworkInfoEndpoints struct {
	// The type of the endpoint that is used to access the instance.
	//
	// Valid values:
	//
	// 	- TCP_VPC: VPC endpoint
	//
	// 	- TCP_INTERNET: public endpoint
	//
	// example:
	//
	// TCP_INTERNET
	EndpointType *string `json:"endpointType,omitempty" xml:"endpointType,omitempty"`
	// The endpoint that is used to access the instance.
	//
	// example:
	//
	// rmq-cn-c4d2tbk****-vpc.cn-hangzhou.rmq.aliyuncs.com:8080
	EndpointUrl *string `json:"endpointUrl,omitempty" xml:"endpointUrl,omitempty"`
	// The whitelist that includes the IP addresses that are allowed to access the ApsaraMQ for RocketMQ broker over the Internet. This parameter can be configured only if you use the public endpoint to access the instance.
	//
	// 	- If you do not configure an IP address whitelist, all CIDR blocks are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If you configure an IP address whitelist, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// We recommend that you configure internetInfo.ipWhitelist instead of this parameter.
	//
	// example:
	//
	// 192.168.x.x/24
	IpWhitelist []*string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyDataNetworkInfoEndpoints) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoEndpoints) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetEndpointType(v string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.EndpointType = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetEndpointUrl(v string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.EndpointUrl = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoEndpoints) SetIpWhitelist(v []*string) *GetInstanceResponseBodyDataNetworkInfoEndpoints {
	s.IpWhitelist = v
	return s
}

type GetInstanceResponseBodyDataNetworkInfoInternetInfo struct {
	// The Internet bandwidth. Unit: MB/s.
	//
	// example:
	//
	// 1
	FlowOutBandwidth *int32 `json:"flowOutBandwidth,omitempty" xml:"flowOutBandwidth,omitempty"`
	// The metering method for Internet usage.
	//
	// Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth. If the Internet access feature is enabled, specify this value for the parameter.
	//
	// 	- uninvolved: N/A. If the Internet access feature is not enabled, specify this value for the parameter.
	//
	// example:
	//
	// payByBandwidth
	FlowOutType *string `json:"flowOutType,omitempty" xml:"flowOutType,omitempty"`
	// Specifies whether to enable the Internet access feature.
	//
	// Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// By default, ApsaraMQ for RocketMQ instances are accessed in virtual private clouds (VPCs). If you enable the Internet access feature, you are charged for Internet outbound bandwidth. For more information, see [Internet access fee](https://help.aliyun.com/document_detail/427240.html).
	//
	// example:
	//
	// enable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// The whitelist that includes the IP addresses that are allowed to access the ApsaraMQ for RocketMQ broker.
	//
	// 	- If this parameter is not configured, all IP addresses are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If this parameter is configured, only the IP addresses that are included in the whitelist can access the ApsaraMQ for RocketMQ broker over the Internet.
	IpWhitelist []*string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyDataNetworkInfoInternetInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoInternetInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) SetFlowOutBandwidth(v int32) *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	s.FlowOutBandwidth = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) SetFlowOutType(v string) *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	s.FlowOutType = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) SetInternetSpec(v string) *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	s.InternetSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoInternetInfo) SetIpWhitelist(v []*string) *GetInstanceResponseBodyDataNetworkInfoInternetInfo {
	s.IpWhitelist = v
	return s
}

type GetInstanceResponseBodyDataNetworkInfoVpcInfo struct {
	// The security group ID.
	//
	// example:
	//
	// sg-hp35r2hc3a3sv8q2sb16
	SecurityGroupIds *string `json:"securityGroupIds,omitempty" xml:"securityGroupIds,omitempty"`
	// Deprecated
	//
	// The ID of the vSwitch with which the instance is associated.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpvz7****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The vSwitches.
	VSwitches []*GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches `json:"vSwitches,omitempty" xml:"vSwitches,omitempty" type:"Repeated"`
	// The ID of the VPC with which the instance is associated.
	//
	// example:
	//
	// vpc-uf6of9452b2pba82c****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetSecurityGroupIds(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.SecurityGroupIds = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetVSwitchId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetVSwitches(v []*GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.VSwitches = v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfo) SetVpcId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfo {
	s.VpcId = &v
	return s
}

type GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches struct {
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpvz7****
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) SetVSwitchId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches) SetZoneId(v string) *GetInstanceResponseBodyDataNetworkInfoVpcInfoVSwitches {
	s.ZoneId = &v
	return s
}

type GetInstanceResponseBodyDataProductInfo struct {
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// This parameter is valid only when the supportAutoScaling parameter is set to enable.
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The retention period of messages. Unit: hours.
	//
	// For information about the valid values of this parameter, see the "Limits on resource quotas" section in [Usage limits](https://help.aliyun.com/document_detail/440347.html).
	//
	// The storage of messages in ApsaraMQ for RocketMQ is serverless and scalable. You are charged for message storage based on your actual usage. You can change the retention period of messages to adjust storage capacity. For more information, see [Storage fee](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The computing specification that is used to send and receive messages. For information about the upper limit of TPS, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The ratio between sent messages and received messages in the instance.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Indicates whether storage encryption is enabled.
	//
	// example:
	//
	// false
	StorageEncryption *bool `json:"storageEncryption,omitempty" xml:"storageEncryption,omitempty"`
	// The storage encryption key.
	//
	// example:
	//
	// xxxxx
	StorageSecretKey *string `json:"storageSecretKey,omitempty" xml:"storageSecretKey,omitempty"`
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// After you enable the elastic TPS feature for a ApsaraMQ for RocketMQ instance, you can use a specific amount of TPS that exceeds the specification limit. You are charged for the elastic TPS feature. For more information, see [Computing fee](https://help.aliyun.com/document_detail/427237.html).
	//
	// > The elastic TPS feature is supported by only specific instance editions. For more information, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	SupportAutoScaling *bool `json:"supportAutoScaling,omitempty" xml:"supportAutoScaling,omitempty"`
	// Indicates whether the message trace feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is not in use. By default, the message trace feature is enabled for ApsaraMQ for RocketMQ instances, regardless of whether this parameter is configured.
	//
	// example:
	//
	// true
	TraceOn *bool `json:"traceOn,omitempty" xml:"traceOn,omitempty"`
}

func (s GetInstanceResponseBodyDataProductInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataProductInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataProductInfo) SetAutoScaling(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.AutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetMessageRetentionTime(v int32) *GetInstanceResponseBodyDataProductInfo {
	s.MessageRetentionTime = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetMsgProcessSpec(v string) *GetInstanceResponseBodyDataProductInfo {
	s.MsgProcessSpec = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetSendReceiveRatio(v float32) *GetInstanceResponseBodyDataProductInfo {
	s.SendReceiveRatio = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetStorageEncryption(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.StorageEncryption = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetStorageSecretKey(v string) *GetInstanceResponseBodyDataProductInfo {
	s.StorageSecretKey = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetSupportAutoScaling(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.SupportAutoScaling = &v
	return s
}

func (s *GetInstanceResponseBodyDataProductInfo) SetTraceOn(v bool) *GetInstanceResponseBodyDataProductInfo {
	s.TraceOn = &v
	return s
}

type GetInstanceResponseBodyDataSoftware struct {
	// The period of upgrade time.
	//
	// example:
	//
	// 02:00-06:00
	MaintainTime *string `json:"maintainTime,omitempty" xml:"maintainTime,omitempty"`
	// The version of software.
	//
	// example:
	//
	// 5.0-rmq-20230619-1
	SoftwareVersion *string `json:"softwareVersion,omitempty" xml:"softwareVersion,omitempty"`
	// The upgrade method.
	//
	// Valid values:
	//
	// - Auto: automatic upgrade
	//
	// - Manual: manual upgrade
	//
	// example:
	//
	// auto
	UpgradeMethod *string `json:"upgradeMethod,omitempty" xml:"upgradeMethod,omitempty"`
}

func (s GetInstanceResponseBodyDataSoftware) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataSoftware) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataSoftware) SetMaintainTime(v string) *GetInstanceResponseBodyDataSoftware {
	s.MaintainTime = &v
	return s
}

func (s *GetInstanceResponseBodyDataSoftware) SetSoftwareVersion(v string) *GetInstanceResponseBodyDataSoftware {
	s.SoftwareVersion = &v
	return s
}

func (s *GetInstanceResponseBodyDataSoftware) SetUpgradeMethod(v string) *GetInstanceResponseBodyDataSoftware {
	s.UpgradeMethod = &v
	return s
}

type GetInstanceResponseBodyDataTags struct {
	// The tag key of the resource.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetInstanceResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataTags) SetKey(v string) *GetInstanceResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyDataTags) SetValue(v string) *GetInstanceResponseBodyDataTags {
	s.Value = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetInstanceAccountRequest struct {
	// The username of the account.
	//
	// If you do not configure this parameter, the default username of the instance is used.
	//
	// example:
	//
	// test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceAccountRequest) SetUsername(v string) *GetInstanceAccountRequest {
	s.Username = &v
	return s
}

type GetInstanceAccountResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetInstanceAccountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, each request\\"s ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// B5C59E80-FCFC-5796-ABE4-D39EAAE578E4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceAccountResponseBody) SetCode(v string) *GetInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetData(v *GetInstanceAccountResponseBodyData) *GetInstanceAccountResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceAccountResponseBody) SetDynamicCode(v string) *GetInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetDynamicMessage(v string) *GetInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetHttpStatusCode(v int32) *GetInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetMessage(v string) *GetInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetRequestId(v string) *GetInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceAccountResponseBody) SetSuccess(v bool) *GetInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type GetInstanceAccountResponseBodyData struct {
	// The password of the account.
	//
	// example:
	//
	// *************
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// xxx
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetInstanceAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceAccountResponseBodyData) SetPassword(v string) *GetInstanceAccountResponseBodyData {
	s.Password = &v
	return s
}

func (s *GetInstanceAccountResponseBodyData) SetUsername(v string) *GetInstanceAccountResponseBodyData {
	s.Username = &v
	return s
}

type GetInstanceAccountResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceAccountResponse) SetHeaders(v map[string]*string) *GetInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceAccountResponse) SetStatusCode(v int32) *GetInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceAccountResponse) SetBody(v *GetInstanceAccountResponseBody) *GetInstanceAccountResponse {
	s.Body = v
	return s
}

type GetMessageDetailResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetMessageDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FAEBD71F-E839-52F9-BD7B-8F1290525841
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMessageDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMessageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageDetailResponseBody) SetCode(v string) *GetMessageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetData(v *GetMessageDetailResponseBodyData) *GetMessageDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetMessageDetailResponseBody) SetDynamicCode(v string) *GetMessageDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetDynamicMessage(v string) *GetMessageDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetHttpStatusCode(v int32) *GetMessageDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetMessage(v string) *GetMessageDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetRequestId(v string) *GetMessageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetSuccess(v bool) *GetMessageDetailResponseBody {
	s.Success = &v
	return s
}

type GetMessageDetailResponseBodyData struct {
	// The message body.
	//
	// example:
	//
	// {}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The size of the message body.
	//
	// example:
	//
	// 123
	BodySize *int32 `json:"bodySize,omitempty" xml:"bodySize,omitempty"`
	// The client on which the message was produced.
	//
	// example:
	//
	// xxx.xx.xxx.xx
	BornHost *string `json:"bornHost,omitempty" xml:"bornHost,omitempty"`
	// The time when the message was generated.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	BornTime *string `json:"bornTime,omitempty" xml:"bornTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The sharding key. This parameter is returned only for ordered messages.
	//
	// example:
	//
	// xx
	MessageGroup *string `json:"messageGroup,omitempty" xml:"messageGroup,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 01BE87E485F0C7808C04543CAF00000001
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The message keys.
	MessageKeys []*string `json:"messageKeys,omitempty" xml:"messageKeys,omitempty" type:"Repeated"`
	// The tags.
	//
	// example:
	//
	// xx
	MessageTag *string `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
	// The message type.
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The broker on which the message was stored.
	//
	// example:
	//
	// xxx.xx.xxx.xx
	StoreHost *string `json:"storeHost,omitempty" xml:"storeHost,omitempty"`
	// The time when the message was stored.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	StoreTime *string `json:"storeTime,omitempty" xml:"storeTime,omitempty"`
	// The default system attributes.
	SystemProperties map[string]*string `json:"systemProperties,omitempty" xml:"systemProperties,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// The user attributes.
	UserProperties map[string]*string `json:"userProperties,omitempty" xml:"userProperties,omitempty"`
}

func (s GetMessageDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMessageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMessageDetailResponseBodyData) SetBody(v string) *GetMessageDetailResponseBodyData {
	s.Body = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetBodySize(v int32) *GetMessageDetailResponseBodyData {
	s.BodySize = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetBornHost(v string) *GetMessageDetailResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetBornTime(v string) *GetMessageDetailResponseBodyData {
	s.BornTime = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetInstanceId(v string) *GetMessageDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageGroup(v string) *GetMessageDetailResponseBodyData {
	s.MessageGroup = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageId(v string) *GetMessageDetailResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageKeys(v []*string) *GetMessageDetailResponseBodyData {
	s.MessageKeys = v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageTag(v string) *GetMessageDetailResponseBodyData {
	s.MessageTag = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageType(v string) *GetMessageDetailResponseBodyData {
	s.MessageType = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetRegionId(v string) *GetMessageDetailResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetStoreHost(v string) *GetMessageDetailResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetStoreTime(v string) *GetMessageDetailResponseBodyData {
	s.StoreTime = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetSystemProperties(v map[string]*string) *GetMessageDetailResponseBodyData {
	s.SystemProperties = v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetTopicName(v string) *GetMessageDetailResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetUserProperties(v map[string]*string) *GetMessageDetailResponseBodyData {
	s.UserProperties = v
	return s
}

type GetMessageDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMessageDetailResponse) GoString() string {
	return s.String()
}

func (s *GetMessageDetailResponse) SetHeaders(v map[string]*string) *GetMessageDetailResponse {
	s.Headers = v
	return s
}

func (s *GetMessageDetailResponse) SetStatusCode(v int32) *GetMessageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageDetailResponse) SetBody(v *GetMessageDetailResponseBody) *GetMessageDetailResponse {
	s.Body = v
	return s
}

type GetTopicResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	Data *GetTopicResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// topicName
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The topic cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBody) SetCode(v string) *GetTopicResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicResponseBody) SetData(v *GetTopicResponseBodyData) *GetTopicResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicResponseBody) SetDynamicCode(v string) *GetTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetTopicResponseBody) SetDynamicMessage(v string) *GetTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetTopicResponseBody) SetHttpStatusCode(v int32) *GetTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicResponseBody) SetMessage(v string) *GetTopicResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicResponseBody) SetRequestId(v string) *GetTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicResponseBody) SetSuccess(v bool) *GetTopicResponseBody {
	s.Success = &v
	return s
}

type GetTopicResponseBodyData struct {
	// The time when the topic was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MaxSendTps *int64  `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// The message type of the topic.
	//
	// Valid values:
	//
	// 	- TRANSACTION: transactional message
	//
	// 	- FIFO: ordered message
	//
	// 	- DELAY: scheduled or delayed message
	//
	// 	- NORMAL: normal message
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The remarks on the topic.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The state of the topic.
	//
	// Valid values:
	//
	// 	- RUNNING: The topic is running.
	//
	// 	- CREATING: The topic is being created.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// The time when the topic was last updated.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBodyData) SetCreateTime(v string) *GetTopicResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetInstanceId(v string) *GetTopicResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetMaxSendTps(v int64) *GetTopicResponseBodyData {
	s.MaxSendTps = &v
	return s
}

func (s *GetTopicResponseBodyData) SetMessageType(v string) *GetTopicResponseBodyData {
	s.MessageType = &v
	return s
}

func (s *GetTopicResponseBodyData) SetRegionId(v string) *GetTopicResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetRemark(v string) *GetTopicResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetTopicResponseBodyData) SetStatus(v string) *GetTopicResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTopicResponseBodyData) SetTopicName(v string) *GetTopicResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetTopicResponseBodyData) SetUpdateTime(v string) *GetTopicResponseBodyData {
	s.UpdateTime = &v
	return s
}

type GetTopicResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicResponse) GoString() string {
	return s.String()
}

func (s *GetTopicResponse) SetHeaders(v map[string]*string) *GetTopicResponse {
	s.Headers = v
	return s
}

func (s *GetTopicResponse) SetStatusCode(v int32) *GetTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicResponse) SetBody(v *GetTopicResponseBody) *GetTopicResponse {
	s.Body = v
	return s
}

type GetTraceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetTraceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 7779A8FC-1BCD-5A1D-A603-C4A9BD8ADC49
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTraceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) SetCode(v string) *GetTraceResponseBody {
	s.Code = &v
	return s
}

func (s *GetTraceResponseBody) SetData(v *GetTraceResponseBodyData) *GetTraceResponseBody {
	s.Data = v
	return s
}

func (s *GetTraceResponseBody) SetDynamicCode(v string) *GetTraceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetTraceResponseBody) SetDynamicMessage(v string) *GetTraceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetTraceResponseBody) SetHttpStatusCode(v int32) *GetTraceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTraceResponseBody) SetMessage(v string) *GetTraceResponseBody {
	s.Message = &v
	return s
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponseBody) SetSuccess(v bool) *GetTraceResponseBody {
	s.Success = &v
	return s
}

type GetTraceResponseBodyData struct {
	// Broker trace info.
	BrokerInfo *GetTraceResponseBodyDataBrokerInfo `json:"brokerInfo,omitempty" xml:"brokerInfo,omitempty" type:"Struct"`
	// Consumer trace info.
	ConsumerInfos []*GetTraceResponseBodyDataConsumerInfos `json:"consumerInfos,omitempty" xml:"consumerInfos,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The message information.
	MessageInfo *GetTraceResponseBodyDataMessageInfo `json:"messageInfo,omitempty" xml:"messageInfo,omitempty" type:"Struct"`
	// Producer trace info.
	ProducerInfo *GetTraceResponseBodyDataProducerInfo `json:"producerInfo,omitempty" xml:"producerInfo,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetTraceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyData) SetBrokerInfo(v *GetTraceResponseBodyDataBrokerInfo) *GetTraceResponseBodyData {
	s.BrokerInfo = v
	return s
}

func (s *GetTraceResponseBodyData) SetConsumerInfos(v []*GetTraceResponseBodyDataConsumerInfos) *GetTraceResponseBodyData {
	s.ConsumerInfos = v
	return s
}

func (s *GetTraceResponseBodyData) SetInstanceId(v string) *GetTraceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTraceResponseBodyData) SetMessageInfo(v *GetTraceResponseBodyDataMessageInfo) *GetTraceResponseBodyData {
	s.MessageInfo = v
	return s
}

func (s *GetTraceResponseBodyData) SetProducerInfo(v *GetTraceResponseBodyDataProducerInfo) *GetTraceResponseBodyData {
	s.ProducerInfo = v
	return s
}

func (s *GetTraceResponseBodyData) SetRegionId(v string) *GetTraceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetTraceResponseBodyData) SetTopicName(v string) *GetTraceResponseBodyData {
	s.TopicName = &v
	return s
}

type GetTraceResponseBodyDataBrokerInfo struct {
	// Delay status.
	//
	// example:
	//
	// SUCCESS
	DelayStatus *string `json:"delayStatus,omitempty" xml:"delayStatus,omitempty"`
	// Operation list.
	Operations []*GetTraceResponseBodyDataBrokerInfoOperations `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
	// Preset delivery time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	PresetDelayTime *string `json:"presetDelayTime,omitempty" xml:"presetDelayTime,omitempty"`
}

func (s GetTraceResponseBodyDataBrokerInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataBrokerInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataBrokerInfo) SetDelayStatus(v string) *GetTraceResponseBodyDataBrokerInfo {
	s.DelayStatus = &v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfo) SetOperations(v []*GetTraceResponseBodyDataBrokerInfoOperations) *GetTraceResponseBodyDataBrokerInfo {
	s.Operations = v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfo) SetPresetDelayTime(v string) *GetTraceResponseBodyDataBrokerInfo {
	s.PresetDelayTime = &v
	return s
}

type GetTraceResponseBodyDataBrokerInfoOperations struct {
	// Operation time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// Operation type.
	//
	// example:
	//
	// ADD
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
}

func (s GetTraceResponseBodyDataBrokerInfoOperations) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataBrokerInfoOperations) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataBrokerInfoOperations) SetOperateTime(v string) *GetTraceResponseBodyDataBrokerInfoOperations {
	s.OperateTime = &v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfoOperations) SetOperateType(v string) *GetTraceResponseBodyDataBrokerInfoOperations {
	s.OperateType = &v
	return s
}

type GetTraceResponseBodyDataConsumerInfos struct {
	// Consume status.
	//
	// example:
	//
	// SUCCESS
	ConsumeStatus *string `json:"consumeStatus,omitempty" xml:"consumeStatus,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// GID_inspector_group
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// Dead letter info.
	DeadLetterInfo *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo `json:"deadLetterInfo,omitempty" xml:"deadLetterInfo,omitempty" type:"Struct"`
	// Whether it is a dead letter message.
	//
	// example:
	//
	// true
	DeadMessage *bool `json:"deadMessage,omitempty" xml:"deadMessage,omitempty"`
	// Consumer record list.
	Records []*GetTraceResponseBodyDataConsumerInfosRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodyDataConsumerInfos) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataConsumerInfos) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetConsumeStatus(v string) *GetTraceResponseBodyDataConsumerInfos {
	s.ConsumeStatus = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetConsumerGroupId(v string) *GetTraceResponseBodyDataConsumerInfos {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetDeadLetterInfo(v *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) *GetTraceResponseBodyDataConsumerInfos {
	s.DeadLetterInfo = v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetDeadMessage(v bool) *GetTraceResponseBodyDataConsumerInfos {
	s.DeadMessage = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetRecords(v []*GetTraceResponseBodyDataConsumerInfosRecords) *GetTraceResponseBodyDataConsumerInfos {
	s.Records = v
	return s
}

type GetTraceResponseBodyDataConsumerInfosDeadLetterInfo struct {
	// MessageId.
	//
	// example:
	//
	// 7F000001001F7A4F0F29463F0376047D
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// Arrival time in the dead letter queue.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	ToDlqTime *string `json:"toDlqTime,omitempty" xml:"toDlqTime,omitempty"`
	// The topic name.
	//
	// example:
	//
	// Register_Sync
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) SetMessageId(v string) *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo {
	s.MessageId = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) SetToDlqTime(v string) *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo {
	s.ToDlqTime = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) SetTopicName(v string) *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo {
	s.TopicName = &v
	return s
}

type GetTraceResponseBodyDataConsumerInfosRecords struct {
	// Client host.
	//
	// example:
	//
	// xx.xx.xx.xx
	ClientHost *string `json:"clientHost,omitempty" xml:"clientHost,omitempty"`
	// Consume status.
	//
	// example:
	//
	// SUCCESS
	ConsumeStatus *string `json:"consumeStatus,omitempty" xml:"consumeStatus,omitempty"`
	// Whether to consume fifo.
	//
	// example:
	//
	// true
	FifoEnable *bool `json:"fifoEnable,omitempty" xml:"fifoEnable,omitempty"`
	// Operation list.
	Operations []*GetTraceResponseBodyDataConsumerInfosRecordsOperations `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
	// POP_CK
	//
	// example:
	//
	// 123
	PopCk *string `json:"popCk,omitempty" xml:"popCk,omitempty"`
	// Consumer name.
	//
	// example:
	//
	// test
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetTraceResponseBodyDataConsumerInfosRecords) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataConsumerInfosRecords) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetClientHost(v string) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.ClientHost = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetConsumeStatus(v string) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.ConsumeStatus = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetFifoEnable(v bool) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.FifoEnable = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetOperations(v []*GetTraceResponseBodyDataConsumerInfosRecordsOperations) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.Operations = v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetPopCk(v string) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.PopCk = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetUserName(v string) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.UserName = &v
	return s
}

type GetTraceResponseBodyDataConsumerInfosRecordsOperations struct {
	// Whether it is a dead letter message.
	//
	// example:
	//
	// true
	DeadMessage *bool `json:"deadMessage,omitempty" xml:"deadMessage,omitempty"`
	// Invisible time, milliseconds.
	//
	// example:
	//
	// 100
	InvisibleTime *int64 `json:"invisibleTime,omitempty" xml:"invisibleTime,omitempty"`
	// Operation time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// Operation type.
	//
	// example:
	//
	// ADD
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
}

func (s GetTraceResponseBodyDataConsumerInfosRecordsOperations) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataConsumerInfosRecordsOperations) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) SetDeadMessage(v bool) *GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	s.DeadMessage = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) SetInvisibleTime(v int64) *GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	s.InvisibleTime = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) SetOperateTime(v string) *GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	s.OperateTime = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) SetOperateType(v string) *GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	s.OperateType = &v
	return s
}

type GetTraceResponseBodyDataMessageInfo struct {
	// Message body.
	//
	// example:
	//
	// {}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Message born host.
	//
	// example:
	//
	// x.x.x.x
	BornHost *string `json:"bornHost,omitempty" xml:"bornHost,omitempty"`
	// Message born time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	BornTime *string `json:"bornTime,omitempty" xml:"bornTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-u0t2ygjq505
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Message grpup.
	//
	// example:
	//
	// xx
	MessageGroup *string `json:"messageGroup,omitempty" xml:"messageGroup,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 0A79275A00207A4F0F2916C92F9A0B94
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// Message keys.
	MessageKeys []*string `json:"messageKeys,omitempty" xml:"messageKeys,omitempty" type:"Repeated"`
	// Message tag.
	//
	// example:
	//
	// xx
	MessageTag *string `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
	// Message type.
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Message store host.
	//
	// example:
	//
	// x.x.x.x
	StoreHost *string `json:"storeHost,omitempty" xml:"storeHost,omitempty"`
	// Message store time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	StoreTime *string `json:"storeTime,omitempty" xml:"storeTime,omitempty"`
	// The topic name.
	//
	// example:
	//
	// Topic_normal_inspector
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// Message transaction id.
	//
	// example:
	//
	// xx
	TransactionId *string `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// Message user properties.
	UserProperties map[string]*string `json:"userProperties,omitempty" xml:"userProperties,omitempty"`
}

func (s GetTraceResponseBodyDataMessageInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataMessageInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataMessageInfo) SetBody(v string) *GetTraceResponseBodyDataMessageInfo {
	s.Body = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetBornHost(v string) *GetTraceResponseBodyDataMessageInfo {
	s.BornHost = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetBornTime(v string) *GetTraceResponseBodyDataMessageInfo {
	s.BornTime = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetInstanceId(v string) *GetTraceResponseBodyDataMessageInfo {
	s.InstanceId = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageGroup(v string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageGroup = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageId(v string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageId = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageKeys(v []*string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageKeys = v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageTag(v string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageTag = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageType(v string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageType = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetRegionId(v string) *GetTraceResponseBodyDataMessageInfo {
	s.RegionId = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetStoreHost(v string) *GetTraceResponseBodyDataMessageInfo {
	s.StoreHost = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetStoreTime(v string) *GetTraceResponseBodyDataMessageInfo {
	s.StoreTime = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetTopicName(v string) *GetTraceResponseBodyDataMessageInfo {
	s.TopicName = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetTransactionId(v string) *GetTraceResponseBodyDataMessageInfo {
	s.TransactionId = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetUserProperties(v map[string]*string) *GetTraceResponseBodyDataMessageInfo {
	s.UserProperties = v
	return s
}

type GetTraceResponseBodyDataProducerInfo struct {
	// Producer record list.
	Records []*GetTraceResponseBodyDataProducerInfoRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodyDataProducerInfo) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataProducerInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataProducerInfo) SetRecords(v []*GetTraceResponseBodyDataProducerInfoRecords) *GetTraceResponseBodyDataProducerInfo {
	s.Records = v
	return s
}

type GetTraceResponseBodyDataProducerInfoRecords struct {
	// Arrive time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	ArriveTime *string `json:"arriveTime,omitempty" xml:"arriveTime,omitempty"`
	// Client host.
	//
	// example:
	//
	// xx.xx.xx.xx
	ClientHost *string `json:"clientHost,omitempty" xml:"clientHost,omitempty"`
	// Dead-letter queue message ID.
	//
	// example:
	//
	// 0A79275A00207A4F0F2916C92F9A0B94
	DlqOriginMessageId *string `json:"dlqOriginMessageId,omitempty" xml:"dlqOriginMessageId,omitempty"`
	// Dead-letter queue topic.
	//
	// example:
	//
	// test_topic
	DlqOriginTopic *string `json:"dlqOriginTopic,omitempty" xml:"dlqOriginTopic,omitempty"`
	// Message source.
	//
	// example:
	//
	// CONSOLE
	MessageSource *string `json:"messageSource,omitempty" xml:"messageSource,omitempty"`
	// Producer duration.
	//
	// example:
	//
	// 100
	ProduceDuration *int64 `json:"produceDuration,omitempty" xml:"produceDuration,omitempty"`
	// Producer status.
	//
	// example:
	//
	// SUCCESS
	ProduceStatus *string `json:"produceStatus,omitempty" xml:"produceStatus,omitempty"`
	// Producer time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	ProduceTime *string `json:"produceTime,omitempty" xml:"produceTime,omitempty"`
	// Producer name.
	//
	// example:
	//
	// xxx
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetTraceResponseBodyDataProducerInfoRecords) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponseBodyDataProducerInfoRecords) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetArriveTime(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ArriveTime = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetClientHost(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ClientHost = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetDlqOriginMessageId(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.DlqOriginMessageId = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetDlqOriginTopic(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.DlqOriginTopic = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetMessageSource(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.MessageSource = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetProduceDuration(v int64) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ProduceDuration = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetProduceStatus(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ProduceStatus = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetProduceTime(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ProduceTime = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetUserName(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.UserName = &v
	return s
}

type GetTraceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTraceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTraceResponse) GoString() string {
	return s.String()
}

func (s *GetTraceResponse) SetHeaders(v map[string]*string) *GetTraceResponse {
	s.Headers = v
	return s
}

func (s *GetTraceResponse) SetStatusCode(v int32) *GetTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTraceResponse) SetBody(v *GetTraceResponseBody) *GetTraceResponse {
	s.Body = v
	return s
}

type ListAvailableZonesResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	Data []*ListAvailableZonesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListAvailableZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableZonesResponseBody) SetCode(v string) *ListAvailableZonesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetData(v []*ListAvailableZonesResponseBodyData) *ListAvailableZonesResponseBody {
	s.Data = v
	return s
}

func (s *ListAvailableZonesResponseBody) SetDynamicCode(v string) *ListAvailableZonesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetDynamicMessage(v string) *ListAvailableZonesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetHttpStatusCode(v int32) *ListAvailableZonesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetMessage(v string) *ListAvailableZonesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetRequestId(v string) *ListAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetSuccess(v bool) *ListAvailableZonesResponseBody {
	s.Success = &v
	return s
}

type ListAvailableZonesResponseBodyData struct {
	// The time when the zone was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the zone was last updated.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the current zone.
	//
	// example:
	//
	// cn-qingdao-b
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
	// The name of the current zone.
	//
	// example:
	//
	// ha-cn-t9b30w902vm_qrs
	ZoneName *string `json:"zoneName,omitempty" xml:"zoneName,omitempty"`
}

func (s ListAvailableZonesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvailableZonesResponseBodyData) SetCreateTime(v string) *ListAvailableZonesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAvailableZonesResponseBodyData) SetUpdateTime(v string) *ListAvailableZonesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListAvailableZonesResponseBodyData) SetZoneId(v string) *ListAvailableZonesResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *ListAvailableZonesResponseBodyData) SetZoneName(v string) *ListAvailableZonesResponseBodyData {
	s.ZoneName = &v
	return s
}

type ListAvailableZonesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableZonesResponse) SetHeaders(v map[string]*string) *ListAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableZonesResponse) SetStatusCode(v int32) *ListAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableZonesResponse) SetBody(v *ListAvailableZonesResponseBody) *ListAvailableZonesResponse {
	s.Body = v
	return s
}

type ListConsumerConnectionsResponseBody struct {
	// The returned error code.
	//
	// example:
	//
	// MissingPageNumber
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListConsumerConnectionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A3620115-6F1F-5CFB-AA3F-BBD4853B2EC4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListConsumerConnectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsResponseBody) SetCode(v string) *ListConsumerConnectionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetData(v *ListConsumerConnectionsResponseBodyData) *ListConsumerConnectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetDynamicCode(v string) *ListConsumerConnectionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetDynamicMessage(v string) *ListConsumerConnectionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetHttpStatusCode(v int32) *ListConsumerConnectionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetMessage(v string) *ListConsumerConnectionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetRequestId(v string) *ListConsumerConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBody) SetSuccess(v bool) *ListConsumerConnectionsResponseBody {
	s.Success = &v
	return s
}

type ListConsumerConnectionsResponseBodyData struct {
	// The client connection list
	Connections []*ListConsumerConnectionsResponseBodyDataConnections `json:"connections,omitempty" xml:"connections,omitempty" type:"Repeated"`
	// The consumer group ID.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListConsumerConnectionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerConnectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsResponseBodyData) SetConnections(v []*ListConsumerConnectionsResponseBodyDataConnections) *ListConsumerConnectionsResponseBodyData {
	s.Connections = v
	return s
}

func (s *ListConsumerConnectionsResponseBodyData) SetConsumerGroupId(v string) *ListConsumerConnectionsResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyData) SetInstanceId(v string) *ListConsumerConnectionsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyData) SetRegionId(v string) *ListConsumerConnectionsResponseBodyData {
	s.RegionId = &v
	return s
}

type ListConsumerConnectionsResponseBodyDataConnections struct {
	// The ID of the client.
	//
	// example:
	//
	// 172.17.135.197@17392#1936705963#551717232#9873695589062458
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// Host IP/Public IP
	//
	// example:
	//
	// xx.xx.xx.xx
	EgressIp *string `json:"egressIp,omitempty" xml:"egressIp,omitempty"`
	// The `hostname` of the cloud-native box.
	//
	// example:
	//
	// vos
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// The language of the client.
	//
	// example:
	//
	// java
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// Consumption Mode
	//
	// - BROADCASTING
	//
	// - CLUSTERING
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The version of the client.
	//
	// example:
	//
	// 1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListConsumerConnectionsResponseBodyDataConnections) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerConnectionsResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetClientId(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.ClientId = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetEgressIp(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.EgressIp = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetHostname(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.Hostname = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetLanguage(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.Language = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetMessageModel(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.MessageModel = &v
	return s
}

func (s *ListConsumerConnectionsResponseBodyDataConnections) SetVersion(v string) *ListConsumerConnectionsResponseBodyDataConnections {
	s.Version = &v
	return s
}

type ListConsumerConnectionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerConnectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerConnectionsResponse) SetHeaders(v map[string]*string) *ListConsumerConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerConnectionsResponse) SetStatusCode(v int32) *ListConsumerConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerConnectionsResponse) SetBody(v *ListConsumerConnectionsResponseBody) *ListConsumerConnectionsResponse {
	s.Body = v
	return s
}

type ListConsumerGroupSubscriptionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*ListConsumerGroupSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5F4D9D5F-625B-59FF-BD4F-DA8284575DB4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListConsumerGroupSubscriptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetCode(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetData(v []*ListConsumerGroupSubscriptionsResponseBodyData) *ListConsumerGroupSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetDynamicCode(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetDynamicMessage(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetHttpStatusCode(v int32) *ListConsumerGroupSubscriptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetMessage(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetRequestId(v string) *ListConsumerGroupSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBody) SetSuccess(v bool) *ListConsumerGroupSubscriptionsResponseBody {
	s.Success = &v
	return s
}

type ListConsumerGroupSubscriptionsResponseBodyData struct {
	// Indicates whether message consumption is consistent. Valid values:
	//
	// 	- false: Unconsumed messages exist in the consumer group.
	//
	// 	- true: No unconsumed message exists in the consumer group.
	//
	// example:
	//
	// true
	Consistency *bool `json:"consistency,omitempty" xml:"consistency,omitempty"`
	// The ID of the consumer group.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The filter expression.
	//
	// example:
	//
	// *
	FilterExpression *string `json:"filterExpression,omitempty" xml:"filterExpression,omitempty"`
	// The type of the filter expression. Valid values:
	//
	// 	- SQL: filters messages by using SQL expressions.
	//
	// 	- TAG: filters messages by using tags.
	//
	// example:
	//
	// SQL
	FilterExpressionType *string `json:"filterExpressionType,omitempty" xml:"filterExpressionType,omitempty"`
	// The consumption mode of the consumer group. Valid values:
	//
	// 	- BROADCASTING: broadcasting consumption
	//
	// 	- CLUSTERING: clustering consumption
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The subscription status. Valid values:
	//
	// 	- ONLINE: The consumer group is online. If the consumer group contains multiple consumers, this value is returned as long as one of the consumers is online.
	//
	// 	- OFFLINE: The consumer group is offline. If the consumer group contains multiple consumers, this value is returned only if all consumers are offline.
	//
	// example:
	//
	// ONLINE
	SubscriptionStatus *string `json:"subscriptionStatus,omitempty" xml:"subscriptionStatus,omitempty"`
	// Indicates whether the topic is created.
	//
	// example:
	//
	// true
	TopicCreated *bool `json:"topicCreated,omitempty" xml:"topicCreated,omitempty"`
	// The topic to which the consumer group subscribes.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListConsumerGroupSubscriptionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetConsistency(v bool) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.Consistency = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetConsumerGroupId(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetFilterExpression(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.FilterExpression = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetFilterExpressionType(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.FilterExpressionType = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetMessageModel(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetSubscriptionStatus(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.SubscriptionStatus = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetTopicCreated(v bool) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.TopicCreated = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponseBodyData) SetTopicName(v string) *ListConsumerGroupSubscriptionsResponseBodyData {
	s.TopicName = &v
	return s
}

type ListConsumerGroupSubscriptionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerGroupSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerGroupSubscriptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupSubscriptionsResponse) SetHeaders(v map[string]*string) *ListConsumerGroupSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponse) SetStatusCode(v int32) *ListConsumerGroupSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponse) SetBody(v *ListConsumerGroupSubscriptionsResponseBody) *ListConsumerGroupSubscriptionsResponse {
	s.Body = v
	return s
}

type ListConsumerGroupsRequest struct {
	// The condition that you want to use to filter consumer groups in the instance. If you leave this parameter empty, all consumer groups in the instance are queried.
	//
	// example:
	//
	// CID-TEST
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListConsumerGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsRequest) SetFilter(v string) *ListConsumerGroupsRequest {
	s.Filter = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetPageNumber(v int32) *ListConsumerGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupsRequest) SetPageSize(v int32) *ListConsumerGroupsRequest {
	s.PageSize = &v
	return s
}

type ListConsumerGroupsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	Data *ListConsumerGroupsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter InstanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// 5503A460-98ED-5543-92CF-4853DE28****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListConsumerGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBody) SetCode(v string) *ListConsumerGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetData(v *ListConsumerGroupsResponseBodyData) *ListConsumerGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetDynamicCode(v string) *ListConsumerGroupsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetDynamicMessage(v string) *ListConsumerGroupsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetHttpStatusCode(v int32) *ListConsumerGroupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetMessage(v string) *ListConsumerGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetRequestId(v string) *ListConsumerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerGroupsResponseBody) SetSuccess(v bool) *ListConsumerGroupsResponseBody {
	s.Success = &v
	return s
}

type ListConsumerGroupsResponseBodyData struct {
	// The paginated data.
	List []*ListConsumerGroupsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListConsumerGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBodyData) SetList(v []*ListConsumerGroupsResponseBodyDataList) *ListConsumerGroupsResponseBodyData {
	s.List = v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetPageNumber(v int64) *ListConsumerGroupsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetPageSize(v int64) *ListConsumerGroupsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyData) SetTotalCount(v int64) *ListConsumerGroupsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListConsumerGroupsResponseBodyDataList struct {
	// The ID of the consumer group.
	//
	// example:
	//
	// GID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The time when the consumer group was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId    *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MaxReceiveTps *int64  `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The remarks on the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The state of the consumer group.
	//
	// Valid values:
	//
	// 	- RUNNING
	//
	//     <!-- -->
	//
	//     : The consumer group is
	//
	//     <!-- -->
	//
	//     running
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- CREATING
	//
	//     <!-- -->
	//
	//     : The consumer group is
	//
	//     <!-- -->
	//
	//     being created
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the consumer group was last updated.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListConsumerGroupsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponseBodyDataList) SetConsumerGroupId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetCreateTime(v string) *ListConsumerGroupsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetInstanceId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetMaxReceiveTps(v int64) *ListConsumerGroupsResponseBodyDataList {
	s.MaxReceiveTps = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetRegionId(v string) *ListConsumerGroupsResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetRemark(v string) *ListConsumerGroupsResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetStatus(v string) *ListConsumerGroupsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListConsumerGroupsResponseBodyDataList) SetUpdateTime(v string) *ListConsumerGroupsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

type ListConsumerGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConsumerGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupsResponse) SetHeaders(v map[string]*string) *ListConsumerGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupsResponse) SetStatusCode(v int32) *ListConsumerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupsResponse) SetBody(v *ListConsumerGroupsResponseBody) *ListConsumerGroupsResponse {
	s.Body = v
	return s
}

type ListInstanceAccountRequest struct {
	// The status of the account.
	//
	// Valid values:
	//
	// 	- DISABLE
	//
	// 	- ENABLE
	//
	// example:
	//
	// ENABLE
	AccountStatus *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	// The account type.
	//
	//   - CUSTOMER
	//
	//   - DEFAULT
	//
	// example:
	//
	// CUSTOMER
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// The page number. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountRequest) SetAccountStatus(v string) *ListInstanceAccountRequest {
	s.AccountStatus = &v
	return s
}

func (s *ListInstanceAccountRequest) SetAccountType(v string) *ListInstanceAccountRequest {
	s.AccountType = &v
	return s
}

func (s *ListInstanceAccountRequest) SetPageNumber(v int32) *ListInstanceAccountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceAccountRequest) SetPageSize(v int32) *ListInstanceAccountRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceAccountRequest) SetUsername(v string) *ListInstanceAccountRequest {
	s.Username = &v
	return s
}

type ListInstanceAccountResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *ListInstanceAccountResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// C115601B-8736-5BBF-AC99-7FEAE1245A80
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountResponseBody) SetAccessDeniedDetail(v string) *ListInstanceAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetCode(v string) *ListInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetData(v *ListInstanceAccountResponseBodyData) *ListInstanceAccountResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceAccountResponseBody) SetDynamicCode(v string) *ListInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetDynamicMessage(v string) *ListInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetHttpStatusCode(v int32) *ListInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetMessage(v string) *ListInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetRequestId(v string) *ListInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceAccountResponseBody) SetSuccess(v bool) *ListInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type ListInstanceAccountResponseBodyData struct {
	// The pagination information.
	List []*ListInstanceAccountResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 24
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstanceAccountResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountResponseBodyData) SetList(v []*ListInstanceAccountResponseBodyDataList) *ListInstanceAccountResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstanceAccountResponseBodyData) SetPageNumber(v int64) *ListInstanceAccountResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceAccountResponseBodyData) SetPageSize(v int64) *ListInstanceAccountResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstanceAccountResponseBodyData) SetTotalCount(v int64) *ListInstanceAccountResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListInstanceAccountResponseBodyDataList struct {
	// The status of the account.
	//
	// Valid values:
	//
	//   - DISABLE
	//
	//   - ENABLE
	//
	// example:
	//
	// ENABLE
	AccountStatus *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	// The account type.
	//
	//   - CUSTOMER
	//
	//   - DEFAULT
	//
	// example:
	//
	// CUSTOMER
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListInstanceAccountResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAccountResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountResponseBodyDataList) SetAccountStatus(v string) *ListInstanceAccountResponseBodyDataList {
	s.AccountStatus = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) SetAccountType(v string) *ListInstanceAccountResponseBodyDataList {
	s.AccountType = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) SetInstanceId(v string) *ListInstanceAccountResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) SetRegionId(v string) *ListInstanceAccountResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListInstanceAccountResponseBodyDataList) SetUsername(v string) *ListInstanceAccountResponseBodyDataList {
	s.Username = &v
	return s
}

type ListInstanceAccountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceAccountResponse) SetHeaders(v map[string]*string) *ListInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceAccountResponse) SetStatusCode(v int32) *ListInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceAccountResponse) SetBody(v *ListInstanceAccountResponseBody) *ListInstanceAccountResponse {
	s.Body = v
	return s
}

type ListInstanceAclRequest struct {
	// The condition that you specify to filter the ACLs. If you do not specify this parameter, all ACLs are queried.
	//
	// example:
	//
	// CID-TEST
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListInstanceAclRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceAclRequest) SetFilter(v string) *ListInstanceAclRequest {
	s.Filter = &v
	return s
}

func (s *ListInstanceAclRequest) SetPageNumber(v int32) *ListInstanceAclRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceAclRequest) SetPageSize(v int32) *ListInstanceAclRequest {
	s.PageSize = &v
	return s
}

type ListInstanceAclResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied due to the reason that the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListInstanceAclResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DA4D2F89-E2C8-5F04-936B-60D55B055FA7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstanceAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceAclResponseBody) SetAccessDeniedDetail(v string) *ListInstanceAclResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetCode(v string) *ListInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetData(v *ListInstanceAclResponseBodyData) *ListInstanceAclResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceAclResponseBody) SetDynamicCode(v string) *ListInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetDynamicMessage(v string) *ListInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetHttpStatusCode(v int32) *ListInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetMessage(v string) *ListInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetRequestId(v string) *ListInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceAclResponseBody) SetSuccess(v bool) *ListInstanceAclResponseBody {
	s.Success = &v
	return s
}

type ListInstanceAclResponseBodyData struct {
	// The pagination information.
	List []*ListInstanceAclResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 24
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstanceAclResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAclResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceAclResponseBodyData) SetList(v []*ListInstanceAclResponseBodyDataList) *ListInstanceAclResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstanceAclResponseBodyData) SetPageNumber(v int64) *ListInstanceAclResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceAclResponseBodyData) SetPageSize(v int64) *ListInstanceAclResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstanceAclResponseBodyData) SetTotalCount(v int64) *ListInstanceAclResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListInstanceAclResponseBodyDataList struct {
	// The ACL type.
	//
	// Valid value:
	//
	// 	- APACHE: open source ACL.
	//
	// example:
	//
	// APACHE
	AclType *string `json:"aclType,omitempty" xml:"aclType,omitempty"`
	// The types of the operations that are allowed by the ACL.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The decision result.
	//
	// Valid values:
	//
	// 	- Deny: Access is denied.
	//
	// 	- Allow: Access is allowed.
	//
	// example:
	//
	// Allow
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The IP address whitelists.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// test
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- Group
	//
	// 	- Topic
	//
	// example:
	//
	// Topic
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The username.
	//
	// example:
	//
	// test
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s ListInstanceAclResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAclResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstanceAclResponseBodyDataList) SetAclType(v string) *ListInstanceAclResponseBodyDataList {
	s.AclType = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetActions(v []*string) *ListInstanceAclResponseBodyDataList {
	s.Actions = v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetDecision(v string) *ListInstanceAclResponseBodyDataList {
	s.Decision = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetInstanceId(v string) *ListInstanceAclResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetIpWhitelists(v []*string) *ListInstanceAclResponseBodyDataList {
	s.IpWhitelists = v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetRegionId(v string) *ListInstanceAclResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetResourceName(v string) *ListInstanceAclResponseBodyDataList {
	s.ResourceName = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetResourceType(v string) *ListInstanceAclResponseBodyDataList {
	s.ResourceType = &v
	return s
}

func (s *ListInstanceAclResponseBodyDataList) SetUsername(v string) *ListInstanceAclResponseBodyDataList {
	s.Username = &v
	return s
}

type ListInstanceAclResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceAclResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceAclResponse) SetHeaders(v map[string]*string) *ListInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceAclResponse) SetStatusCode(v int32) *ListInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceAclResponse) SetBody(v *ListInstanceAclResponseBody) *ListInstanceAclResponse {
	s.Body = v
	return s
}

type ListInstanceIpWhitelistRequest struct {
	// IP whitelist.
	//
	// example:
	//
	// 0.0.0.0/0
	IpWhitelist *string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty"`
	// The page number. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListInstanceIpWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceIpWhitelistRequest) SetIpWhitelist(v string) *ListInstanceIpWhitelistRequest {
	s.IpWhitelist = &v
	return s
}

func (s *ListInstanceIpWhitelistRequest) SetPageNumber(v int32) *ListInstanceIpWhitelistRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceIpWhitelistRequest) SetPageSize(v int32) *ListInstanceIpWhitelistRequest {
	s.PageSize = &v
	return s
}

type ListInstanceIpWhitelistResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *ListInstanceIpWhitelistResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 7358418D-83BD-507A-8079-611C63E05674
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstanceIpWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIpWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceIpWhitelistResponseBody) SetAccessDeniedDetail(v string) *ListInstanceIpWhitelistResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetCode(v string) *ListInstanceIpWhitelistResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetData(v *ListInstanceIpWhitelistResponseBodyData) *ListInstanceIpWhitelistResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetDynamicCode(v string) *ListInstanceIpWhitelistResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetDynamicMessage(v string) *ListInstanceIpWhitelistResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetHttpStatusCode(v int32) *ListInstanceIpWhitelistResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetMessage(v string) *ListInstanceIpWhitelistResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetRequestId(v string) *ListInstanceIpWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBody) SetSuccess(v bool) *ListInstanceIpWhitelistResponseBody {
	s.Success = &v
	return s
}

type ListInstanceIpWhitelistResponseBodyData struct {
	// The pagination information.
	List []*string `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstanceIpWhitelistResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIpWhitelistResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceIpWhitelistResponseBodyData) SetList(v []*string) *ListInstanceIpWhitelistResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstanceIpWhitelistResponseBodyData) SetPageNumber(v int64) *ListInstanceIpWhitelistResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBodyData) SetPageSize(v int64) *ListInstanceIpWhitelistResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstanceIpWhitelistResponseBodyData) SetTotalCount(v int64) *ListInstanceIpWhitelistResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListInstanceIpWhitelistResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceIpWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceIpWhitelistResponse) SetHeaders(v map[string]*string) *ListInstanceIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceIpWhitelistResponse) SetStatusCode(v int32) *ListInstanceIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceIpWhitelistResponse) SetBody(v *ListInstanceIpWhitelistResponseBody) *ListInstanceIpWhitelistResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// The filter condition that is used to query instances. If you do not configure this parameter, all instances are queried.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The page number.
	//
	// Valid values: 1 to 100000000.
	//
	// If you set this parameter to a value smaller than 1, the system uses 1 as the value. If you set this parameter to a value greater than 100000000, the system uses 100000000 as the value.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// Value values: 10 to 200.
	//
	// If you set this parameter to a value smaller than 10, the system uses 10 as the value. If you set this parameter to a value greater than 200, the system uses 200 as the value.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmx7caj******
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary edition of the instance.
	//
	// Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Enterprise Platinum Edition
	//
	// 	- professional: Professional Edition
	SeriesCodes []*string `json:"seriesCodes,omitempty" xml:"seriesCodes,omitempty" type:"Repeated"`
	// The storage encryption key.
	//
	// example:
	//
	// xxxxx
	StorageSecretKey *string `json:"storageSecretKey,omitempty" xml:"storageSecretKey,omitempty"`
	// The tags that are used to filter instances.
	//
	// example:
	//
	// [{"key": "rmq-test", "value": "test"}]
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetFilter(v string) *ListInstancesRequest {
	s.Filter = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetSeriesCodes(v []*string) *ListInstancesRequest {
	s.SeriesCodes = v
	return s
}

func (s *ListInstancesRequest) SetStorageSecretKey(v string) *ListInstancesRequest {
	s.StorageSecretKey = &v
	return s
}

func (s *ListInstancesRequest) SetTags(v string) *ListInstancesRequest {
	s.Tags = &v
	return s
}

type ListInstancesShrinkRequest struct {
	// The filter condition that is used to query instances. If you do not configure this parameter, all instances are queried.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The page number.
	//
	// Valid values: 1 to 100000000.
	//
	// If you set this parameter to a value smaller than 1, the system uses 1 as the value. If you set this parameter to a value greater than 100000000, the system uses 100000000 as the value.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// Value values: 10 to 200.
	//
	// If you set this parameter to a value smaller than 10, the system uses 10 as the value. If you set this parameter to a value greater than 200, the system uses 200 as the value.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmx7caj******
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary edition of the instance.
	//
	// Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Enterprise Platinum Edition
	//
	// 	- professional: Professional Edition
	SeriesCodesShrink *string `json:"seriesCodes,omitempty" xml:"seriesCodes,omitempty"`
	// The storage encryption key.
	//
	// example:
	//
	// xxxxx
	StorageSecretKey *string `json:"storageSecretKey,omitempty" xml:"storageSecretKey,omitempty"`
	// The tags that are used to filter instances.
	//
	// example:
	//
	// [{"key": "rmq-test", "value": "test"}]
	Tags *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesShrinkRequest) SetFilter(v string) *ListInstancesShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageNumber(v int32) *ListInstancesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetPageSize(v int32) *ListInstancesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetResourceGroupId(v string) *ListInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetSeriesCodesShrink(v string) *ListInstancesShrinkRequest {
	s.SeriesCodesShrink = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetStorageSecretKey(v string) *ListInstancesShrinkRequest {
	s.StorageSecretKey = &v
	return s
}

func (s *ListInstancesShrinkRequest) SetTags(v string) *ListInstancesShrinkRequest {
	s.Tags = &v
	return s
}

type ListInstancesResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// MissingPageNumber
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *ListInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// PageNumber
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// pageNumber
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter pageNumber is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 84445A20-2B50-5306-A3C0-AF99FC1833C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetDynamicCode(v string) *ListInstancesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetDynamicMessage(v string) *ListInstancesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

type ListInstancesResponseBodyData struct {
	// The pagination information.
	List []*ListInstancesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) SetList(v []*ListInstancesResponseBodyDataList) *ListInstancesResponseBodyData {
	s.List = v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageNumber(v int64) *ListInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetPageSize(v int64) *ListInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetTotalCount(v int64) *ListInstancesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyDataList struct {
	// The commodity code of the instance. The commodity code of ApsaraMQ for RocketMQ 5.0 instances has a similar format to ons_rmqsub_public_cn.
	//
	// example:
	//
	// ons_rmqsub_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the version of the instance was updated.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2022-09-01 00:00:00
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The number of consumer groups that are created on the instance.
	//
	// example:
	//
	// 10
	GroupCount *int64 `json:"groupCount,omitempty" xml:"groupCount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// test instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The billing method of the instance.
	//
	// Valid values:
	//
	// 	- PayAsYouGo
	//
	// 	- Subscription
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	// The product information.
	ProductInfo *ListInstancesResponseBodyDataListProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The time when the instance was released.
	//
	// example:
	//
	// 2022-09-07 00:00:00
	ReleaseTime *string `json:"releaseTime,omitempty" xml:"releaseTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// This is remark for instance.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmx7caj******
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The primary edition of the instance.
	//
	// Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Enterprise Platinum Edition
	//
	// 	- professional: Professional Edition
	//
	// example:
	//
	// standard
	SeriesCode *string `json:"seriesCode,omitempty" xml:"seriesCode,omitempty"`
	// The code of the service to which the instance belongs. The service code of ApsaraMQ for RocketMQ is rmq.
	//
	// example:
	//
	// rmq
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the instance.
	//
	// Valid values:
	//
	// 	- RELEASED
	//
	// 	- RUNNING
	//
	// 	- STOPPED
	//
	// 	- CHANGING
	//
	// 	- CREATING
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The sub-category edition of the instance.
	//
	// Valid values:
	//
	// 	- cluster_ha: Cluster High-availability Edition
	//
	// 	- single_node: Standalone Edition
	//
	// example:
	//
	// cluster_ha
	SubSeriesCode *string `json:"subSeriesCode,omitempty" xml:"subSeriesCode,omitempty"`
	// The resource tags.
	Tags []*ListInstancesResponseBodyDataListTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The number of topics that are created on the instance.
	//
	// example:
	//
	// 20
	TopicCount *int64 `json:"topicCount,omitempty" xml:"topicCount,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 2022-08-02 00:00:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the user who owns the instance.
	//
	// example:
	//
	// 6W0xz2uPfiwp****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListInstancesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataList) SetCommodityCode(v string) *ListInstancesResponseBodyDataList {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetCreateTime(v string) *ListInstancesResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetExpireTime(v string) *ListInstancesResponseBodyDataList {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetGroupCount(v int64) *ListInstancesResponseBodyDataList {
	s.GroupCount = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetInstanceId(v string) *ListInstancesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetInstanceName(v string) *ListInstancesResponseBodyDataList {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetPaymentType(v string) *ListInstancesResponseBodyDataList {
	s.PaymentType = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetProductInfo(v *ListInstancesResponseBodyDataListProductInfo) *ListInstancesResponseBodyDataList {
	s.ProductInfo = v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetRegionId(v string) *ListInstancesResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetReleaseTime(v string) *ListInstancesResponseBodyDataList {
	s.ReleaseTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetRemark(v string) *ListInstancesResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetResourceGroupId(v string) *ListInstancesResponseBodyDataList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetSeriesCode(v string) *ListInstancesResponseBodyDataList {
	s.SeriesCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetServiceCode(v string) *ListInstancesResponseBodyDataList {
	s.ServiceCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetStartTime(v string) *ListInstancesResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetStatus(v string) *ListInstancesResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetSubSeriesCode(v string) *ListInstancesResponseBodyDataList {
	s.SubSeriesCode = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetTags(v []*ListInstancesResponseBodyDataListTags) *ListInstancesResponseBodyDataList {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetTopicCount(v int64) *ListInstancesResponseBodyDataList {
	s.TopicCount = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetUpdateTime(v string) *ListInstancesResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataList) SetUserId(v string) *ListInstancesResponseBodyDataList {
	s.UserId = &v
	return s
}

type ListInstancesResponseBodyDataListProductInfo struct {
	// Indicates whether the message trace feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is not in use. By default, the message trace feature is enabled for ApsaraMQ for RocketMQ instances, regardless of whether this parameter is configured.
	//
	// example:
	//
	// true
	TraceOn *bool `json:"traceOn,omitempty" xml:"traceOn,omitempty"`
}

func (s ListInstancesResponseBodyDataListProductInfo) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataListProductInfo) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataListProductInfo) SetTraceOn(v bool) *ListInstancesResponseBodyDataListProductInfo {
	s.TraceOn = &v
	return s
}

type ListInstancesResponseBodyDataListTags struct {
	// The tag key of the resource.
	//
	// example:
	//
	// key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesResponseBodyDataListTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyDataListTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataListTags) SetKey(v string) *ListInstancesResponseBodyDataListTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyDataListTags) SetValue(v string) *ListInstancesResponseBodyDataListTags {
	s.Value = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListMessagesRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 2024-09-09 09:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Message Id.
	//
	// example:
	//
	// 7F00000100207A4F0F294A938F7807AE
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// Message key.
	//
	// example:
	//
	// XSCBillResult
	MessageKey *string `json:"messageKey,omitempty" xml:"messageKey,omitempty"`
	// The page number. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The scroll ID of the request.
	//
	// You do not need to configure this parameter for the first page. This parameter is included in the pagination request based on the result returned for the first page.
	//
	// example:
	//
	// B13D0B07-F24B-4790-88D8-D47A38063D00
	ScrollId *string `json:"scrollId,omitempty" xml:"scrollId,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2024-09-09 08:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListMessagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListMessagesRequest) SetEndTime(v string) *ListMessagesRequest {
	s.EndTime = &v
	return s
}

func (s *ListMessagesRequest) SetMessageId(v string) *ListMessagesRequest {
	s.MessageId = &v
	return s
}

func (s *ListMessagesRequest) SetMessageKey(v string) *ListMessagesRequest {
	s.MessageKey = &v
	return s
}

func (s *ListMessagesRequest) SetPageNumber(v int32) *ListMessagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMessagesRequest) SetPageSize(v int32) *ListMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMessagesRequest) SetScrollId(v string) *ListMessagesRequest {
	s.ScrollId = &v
	return s
}

func (s *ListMessagesRequest) SetStartTime(v string) *ListMessagesRequest {
	s.StartTime = &v
	return s
}

type ListMessagesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *ListMessagesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A3531B6A-5A88-52BD-B3C4-A024C3D0AA2E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMessagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBody) SetCode(v string) *ListMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListMessagesResponseBody) SetData(v *ListMessagesResponseBodyData) *ListMessagesResponseBody {
	s.Data = v
	return s
}

func (s *ListMessagesResponseBody) SetDynamicCode(v string) *ListMessagesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListMessagesResponseBody) SetDynamicMessage(v string) *ListMessagesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListMessagesResponseBody) SetHttpStatusCode(v int32) *ListMessagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMessagesResponseBody) SetMessage(v string) *ListMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListMessagesResponseBody) SetRequestId(v string) *ListMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessagesResponseBody) SetSuccess(v bool) *ListMessagesResponseBody {
	s.Success = &v
	return s
}

type ListMessagesResponseBodyData struct {
	// The pagination information.
	List []*ListMessagesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The scroll ID of the request.
	//
	// The ID is automatically generated by the system. The result can be paginated only if this parameter is included in the pagination request.
	//
	// example:
	//
	// B13D0B07-F24B-4790-88D8-D47A38063D00
	ScrollId *string `json:"scrollId,omitempty" xml:"scrollId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListMessagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyData) SetList(v []*ListMessagesResponseBodyDataList) *ListMessagesResponseBodyData {
	s.List = v
	return s
}

func (s *ListMessagesResponseBodyData) SetPageNumber(v int64) *ListMessagesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetPageSize(v int64) *ListMessagesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetScrollId(v string) *ListMessagesResponseBodyData {
	s.ScrollId = &v
	return s
}

func (s *ListMessagesResponseBodyData) SetTotalCount(v int64) *ListMessagesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListMessagesResponseBodyDataList struct {
	// Message body.
	//
	// example:
	//
	// {}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Message body size.
	//
	// example:
	//
	// 100
	BodySize *int32 `json:"bodySize,omitempty" xml:"bodySize,omitempty"`
	// The client on which messages are produced.
	//
	// example:
	//
	// xx.xx.xx.xx
	BornHost *string `json:"bornHost,omitempty" xml:"bornHost,omitempty"`
	// Message born time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	BornTime *string `json:"bornTime,omitempty" xml:"bornTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The message group. This parameter is returned only for ordered messages.
	//
	// example:
	//
	// xx
	MessageGroup *string `json:"messageGroup,omitempty" xml:"messageGroup,omitempty"`
	// Message Id.
	//
	// example:
	//
	// 7F000001000114B4340C5ABF94500079
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// Message keys.
	MessageKeys []*string `json:"messageKeys,omitempty" xml:"messageKeys,omitempty" type:"Repeated"`
	// The message tag.
	//
	// example:
	//
	// xx
	MessageTag *string `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
	// Message type.
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The broker on which messages are stored.
	//
	// example:
	//
	// xx.xx.xx.xx
	StoreHost *string `json:"storeHost,omitempty" xml:"storeHost,omitempty"`
	// Message store time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	StoreTime *string `json:"storeTime,omitempty" xml:"storeTime,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// Message user properties.
	UserProperties map[string]*string `json:"userProperties,omitempty" xml:"userProperties,omitempty"`
}

func (s ListMessagesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListMessagesResponseBodyDataList) SetBody(v string) *ListMessagesResponseBodyDataList {
	s.Body = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetBodySize(v int32) *ListMessagesResponseBodyDataList {
	s.BodySize = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetBornHost(v string) *ListMessagesResponseBodyDataList {
	s.BornHost = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetBornTime(v string) *ListMessagesResponseBodyDataList {
	s.BornTime = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetInstanceId(v string) *ListMessagesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetMessageGroup(v string) *ListMessagesResponseBodyDataList {
	s.MessageGroup = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetMessageId(v string) *ListMessagesResponseBodyDataList {
	s.MessageId = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetMessageKeys(v []*string) *ListMessagesResponseBodyDataList {
	s.MessageKeys = v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetMessageTag(v string) *ListMessagesResponseBodyDataList {
	s.MessageTag = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetMessageType(v string) *ListMessagesResponseBodyDataList {
	s.MessageType = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetRegionId(v string) *ListMessagesResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetStoreHost(v string) *ListMessagesResponseBodyDataList {
	s.StoreHost = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetStoreTime(v string) *ListMessagesResponseBodyDataList {
	s.StoreTime = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetTopicName(v string) *ListMessagesResponseBodyDataList {
	s.TopicName = &v
	return s
}

func (s *ListMessagesResponseBodyDataList) SetUserProperties(v map[string]*string) *ListMessagesResponseBodyDataList {
	s.UserProperties = v
	return s
}

type ListMessagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMessagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMessagesResponse) GoString() string {
	return s.String()
}

func (s *ListMessagesResponse) SetHeaders(v map[string]*string) *ListMessagesResponse {
	s.Headers = v
	return s
}

func (s *ListMessagesResponse) SetStatusCode(v int32) *ListMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMessagesResponse) SetBody(v *ListMessagesResponseBody) *ListMessagesResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingPageNumber
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*ListRegionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B962390-D84B-5D44-8C11-79DF40299D41
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetCode(v string) *ListRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionsResponseBody) SetData(v []*ListRegionsResponseBodyData) *ListRegionsResponseBody {
	s.Data = v
	return s
}

func (s *ListRegionsResponseBody) SetDynamicCode(v string) *ListRegionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListRegionsResponseBody) SetDynamicMessage(v string) *ListRegionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListRegionsResponseBody) SetHttpStatusCode(v int32) *ListRegionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRegionsResponseBody) SetMessage(v string) *ListRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetSuccess(v bool) *ListRegionsResponseBody {
	s.Success = &v
	return s
}

type ListRegionsResponseBodyData struct {
	// The time when the ApsaraMQ for RocketMQ instance was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// hangzhou
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// Indicates whether ApsaraMQ for RocketMQ V4 is activated.
	//
	// example:
	//
	// true
	SupportRocketmqV4 *bool `json:"supportRocketmqV4,omitempty" xml:"supportRocketmqV4,omitempty"`
	// Indicates whether ApsaraMQ for RocketMQ V5 is activated.
	//
	// example:
	//
	// true
	SupportRocketmqV5 *bool `json:"supportRocketmqV5,omitempty" xml:"supportRocketmqV5,omitempty"`
	// The region tags.
	Tags []*ListRegionsResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the ApsaraMQ for RocketMQ instance was last modified.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListRegionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyData) SetCreateTime(v string) *ListRegionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetRegionId(v string) *ListRegionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetRegionName(v string) *ListRegionsResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetSupportRocketmqV4(v bool) *ListRegionsResponseBodyData {
	s.SupportRocketmqV4 = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetSupportRocketmqV5(v bool) *ListRegionsResponseBodyData {
	s.SupportRocketmqV5 = &v
	return s
}

func (s *ListRegionsResponseBodyData) SetTags(v []*ListRegionsResponseBodyDataTags) *ListRegionsResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListRegionsResponseBodyData) SetUpdateTime(v string) *ListRegionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListRegionsResponseBodyDataTags struct {
	// The tag code.
	//
	// example:
	//
	// xx
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
	// The tag value.
	//
	// example:
	//
	// xx
	TagValue interface{} `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListRegionsResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyDataTags) SetTagCode(v string) *ListRegionsResponseBodyDataTags {
	s.TagCode = &v
	return s
}

func (s *ListRegionsResponseBodyDataTags) SetTagValue(v interface{}) *ListRegionsResponseBodyDataTags {
	s.TagValue = v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The position from which the next query starts.
	//
	// example:
	//
	// d09e2b63e1b12d905b7080ff70
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmx7caj******
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// List of resource IDs, in JSON format.
	//
	// example:
	//
	// ["rmq-cn-pe334n08h08"]
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// Resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// List of tags, in JSON format.
	//
	// example:
	//
	// [{"key": "rmq-test", "value": "test"}]
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceGroupId(v string) *ListTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v string) *ListTagResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v string) *ListTagResourcesRequest {
	s.Tag = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// Error code
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Return result
	Data *ListTagResourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Dynamic error code
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// Dynamic error message
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// The topic already exists.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID
	//
	// example:
	//
	// F00C6A70-C782-5DD6-9D11-0CFC710100C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetCode(v string) *ListTagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetData(v *ListTagResourcesResponseBodyData) *ListTagResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListTagResourcesResponseBody) SetDynamicCode(v string) *ListTagResourcesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetDynamicMessage(v string) *ListTagResourcesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetHttpStatusCode(v int32) *ListTagResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetMessage(v string) *ListTagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

type ListTagResourcesResponseBodyData struct {
	// The position from which the next query starts.
	//
	// example:
	//
	// d09e2b63e1b12d905b7080ff70
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Request ID.
	//
	// example:
	//
	// F00C6A70-C782-5DD6-9D11-0CFC710100C7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Resource tag relationships.
	TagResources []*ListTagResourcesResponseBodyDataTagResources `json:"tagResources,omitempty" xml:"tagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyData) SetNextToken(v string) *ListTagResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetRequestId(v string) *ListTagResourcesResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBodyData) SetTagResources(v []*ListTagResourcesResponseBodyDataTagResources) *ListTagResourcesResponseBodyData {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyDataTagResources struct {
	// UID of the resource owner.
	//
	// example:
	//
	// 1876441048322426
	AliUid *int64 `json:"aliUid,omitempty" xml:"aliUid,omitempty"`
	// Tag category.
	//
	// example:
	//
	// custom
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// Resource ID.
	//
	// example:
	//
	// rmq-cn-pe334n08h08
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// Visibility scope.
	//
	// example:
	//
	// public
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// Tag key.
	//
	// example:
	//
	// key
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// Tag value.
	//
	// example:
	//
	// value
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyDataTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyDataTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetAliUid(v int64) *ListTagResourcesResponseBodyDataTagResources {
	s.AliUid = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetCategory(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.Category = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetScope(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.Scope = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyDataTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyDataTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTopicSubscriptionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*ListTopicSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// Topic
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 92A9BE4E-B794-50C8-979C-0456E4D32943
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTopicSubscriptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTopicSubscriptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicSubscriptionsResponseBody) SetCode(v string) *ListTopicSubscriptionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetData(v []*ListTopicSubscriptionsResponseBodyData) *ListTopicSubscriptionsResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetDynamicCode(v string) *ListTopicSubscriptionsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetDynamicMessage(v string) *ListTopicSubscriptionsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetHttpStatusCode(v int32) *ListTopicSubscriptionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetMessage(v string) *ListTopicSubscriptionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetRequestId(v string) *ListTopicSubscriptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBody) SetSuccess(v bool) *ListTopicSubscriptionsResponseBody {
	s.Success = &v
	return s
}

type ListTopicSubscriptionsResponseBodyData struct {
	// Indicates whether message consumption is consistent. Valid values:
	//
	// 	- false: Unconsumed messages exist in the consumer group.
	//
	// 	- true: No unconsumed message exists in the consumer group.
	//
	// example:
	//
	// true
	Consistency *string `json:"consistency,omitempty" xml:"consistency,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The filter expression.
	//
	// example:
	//
	// *
	FilterExpression *string `json:"filterExpression,omitempty" xml:"filterExpression,omitempty"`
	// The type of the filter expression. Valid values: SQL, TAG, and UNSPECIFIED. The value SQL indicates that messages are filtered by using SQL expressions. The value TAG indicates that messages are filtered by using tags. The value UNSPECIFIED indicates that no filter expression type is specified.
	//
	// example:
	//
	// SQL
	FilterExpressionType *string `json:"filterExpressionType,omitempty" xml:"filterExpressionType,omitempty"`
	// The consumption mode. Valid values: BROADCASTING and CLUSTERING.
	//
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// The subscription status. Valid values: ONLINE and OFFLINE.
	//
	// example:
	//
	// ONLINE
	SubscriptionStatus *string `json:"subscriptionStatus,omitempty" xml:"subscriptionStatus,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListTopicSubscriptionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTopicSubscriptionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTopicSubscriptionsResponseBodyData) SetConsistency(v string) *ListTopicSubscriptionsResponseBodyData {
	s.Consistency = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetConsumerGroupId(v string) *ListTopicSubscriptionsResponseBodyData {
	s.ConsumerGroupId = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetFilterExpression(v string) *ListTopicSubscriptionsResponseBodyData {
	s.FilterExpression = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetFilterExpressionType(v string) *ListTopicSubscriptionsResponseBodyData {
	s.FilterExpressionType = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetMessageModel(v string) *ListTopicSubscriptionsResponseBodyData {
	s.MessageModel = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetSubscriptionStatus(v string) *ListTopicSubscriptionsResponseBodyData {
	s.SubscriptionStatus = &v
	return s
}

func (s *ListTopicSubscriptionsResponseBodyData) SetTopicName(v string) *ListTopicSubscriptionsResponseBodyData {
	s.TopicName = &v
	return s
}

type ListTopicSubscriptionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicSubscriptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTopicSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListTopicSubscriptionsResponse) SetHeaders(v map[string]*string) *ListTopicSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListTopicSubscriptionsResponse) SetStatusCode(v int32) *ListTopicSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicSubscriptionsResponse) SetBody(v *ListTopicSubscriptionsResponseBody) *ListTopicSubscriptionsResponse {
	s.Body = v
	return s
}

type ListTopicsRequest struct {
	// The condition that you want to use to filter topics in the instance. If you leave this parameter empty, all topics in the instance are queried.
	//
	// example:
	//
	// topic_test
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The message types of the topics.
	MessageTypes []*string `json:"messageTypes,omitempty" xml:"messageTypes,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 3
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListTopicsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListTopicsRequest) SetFilter(v string) *ListTopicsRequest {
	s.Filter = &v
	return s
}

func (s *ListTopicsRequest) SetMessageTypes(v []*string) *ListTopicsRequest {
	s.MessageTypes = v
	return s
}

func (s *ListTopicsRequest) SetPageNumber(v int32) *ListTopicsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsRequest) SetPageSize(v int32) *ListTopicsRequest {
	s.PageSize = &v
	return s
}

type ListTopicsShrinkRequest struct {
	// The condition that you want to use to filter topics in the instance. If you leave this parameter empty, all topics in the instance are queried.
	//
	// example:
	//
	// topic_test
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The message types of the topics.
	MessageTypesShrink *string `json:"messageTypes,omitempty" xml:"messageTypes,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 3
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListTopicsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTopicsShrinkRequest) SetFilter(v string) *ListTopicsShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListTopicsShrinkRequest) SetMessageTypesShrink(v string) *ListTopicsShrinkRequest {
	s.MessageTypesShrink = &v
	return s
}

func (s *ListTopicsShrinkRequest) SetPageNumber(v int32) *ListTopicsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsShrinkRequest) SetPageSize(v int32) *ListTopicsShrinkRequest {
	s.PageSize = &v
	return s
}

type ListTopicsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	Data *ListTopicsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// topicName
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The topic cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTopicsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBody) SetCode(v string) *ListTopicsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicsResponseBody) SetData(v *ListTopicsResponseBodyData) *ListTopicsResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicsResponseBody) SetDynamicCode(v string) *ListTopicsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTopicsResponseBody) SetDynamicMessage(v string) *ListTopicsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTopicsResponseBody) SetHttpStatusCode(v int32) *ListTopicsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicsResponseBody) SetMessage(v string) *ListTopicsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicsResponseBody) SetRequestId(v string) *ListTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicsResponseBody) SetSuccess(v bool) *ListTopicsResponseBody {
	s.Success = &v
	return s
}

type ListTopicsResponseBodyData struct {
	// The paginated data.
	List []*ListTopicsResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 3
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTopicsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyData) SetList(v []*ListTopicsResponseBodyDataList) *ListTopicsResponseBodyData {
	s.List = v
	return s
}

func (s *ListTopicsResponseBodyData) SetPageNumber(v int64) *ListTopicsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsResponseBodyData) SetPageSize(v int64) *ListTopicsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTopicsResponseBodyData) SetTotalCount(v int64) *ListTopicsResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListTopicsResponseBodyDataList struct {
	// The time when the topic was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	MaxSendTps *int64  `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// The message type of the topic.
	//
	// Valid values:
	//
	// 	- TRANSACTION
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     transactional message
	//
	//     <!-- -->
	//
	// 	- FIFO
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     ordered message
	//
	//     <!-- -->
	//
	// 	- DELAY
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     scheduled or delayed message
	//
	//     <!-- -->
	//
	// 	- NORMAL
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     normal message
	//
	//     <!-- -->
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The remarks on the topic.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// The state of the topic.
	//
	// Valid values:
	//
	// 	- RUNNING
	//
	//     <!-- -->
	//
	//     : The topic is
	//
	//     <!-- -->
	//
	//     running
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- CREATING
	//
	//     <!-- -->
	//
	//     : The topic is
	//
	//     <!-- -->
	//
	//     being created
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// The time when the topic was last updated.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListTopicsResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyDataList) SetCreateTime(v string) *ListTopicsResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetInstanceId(v string) *ListTopicsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetMaxSendTps(v int64) *ListTopicsResponseBodyDataList {
	s.MaxSendTps = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetMessageType(v string) *ListTopicsResponseBodyDataList {
	s.MessageType = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetRegionId(v string) *ListTopicsResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetRemark(v string) *ListTopicsResponseBodyDataList {
	s.Remark = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetStatus(v string) *ListTopicsResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetTopicName(v string) *ListTopicsResponseBodyDataList {
	s.TopicName = &v
	return s
}

func (s *ListTopicsResponseBodyDataList) SetUpdateTime(v string) *ListTopicsResponseBodyDataList {
	s.UpdateTime = &v
	return s
}

type ListTopicsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTopicsResponse) GoString() string {
	return s.String()
}

func (s *ListTopicsResponse) SetHeaders(v map[string]*string) *ListTopicsResponse {
	s.Headers = v
	return s
}

func (s *ListTopicsResponse) SetStatusCode(v int32) *ListTopicsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicsResponse) SetBody(v *ListTopicsResponseBody) *ListTopicsResponse {
	s.Body = v
	return s
}

type ListTracesRequest struct {
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-19 10:10:09
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The message ID.
	//
	// This parameter is required if you set queryType to MESSAGE_ID.
	//
	// example:
	//
	// 0100163E0EC1F1965C04C7906700000000
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The message key.
	//
	// This parameter is required if you set queryType to MESSAGE_ID.
	//
	// example:
	//
	// order_ceating
	MessageKey *string `json:"messageKey,omitempty" xml:"messageKey,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The query type.
	//
	// Valid values:
	//
	// 	- MESSAGE_ID
	//
	// 	- MESSAGE_KEY
	//
	// 	- TOPIC
	//
	// This parameter is required.
	//
	// example:
	//
	// MESSAGE_ID
	QueryType *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-10 10:42:11
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListTracesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTracesRequest) GoString() string {
	return s.String()
}

func (s *ListTracesRequest) SetEndTime(v string) *ListTracesRequest {
	s.EndTime = &v
	return s
}

func (s *ListTracesRequest) SetMessageId(v string) *ListTracesRequest {
	s.MessageId = &v
	return s
}

func (s *ListTracesRequest) SetMessageKey(v string) *ListTracesRequest {
	s.MessageKey = &v
	return s
}

func (s *ListTracesRequest) SetPageNumber(v int32) *ListTracesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTracesRequest) SetPageSize(v int32) *ListTracesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTracesRequest) SetQueryType(v string) *ListTracesRequest {
	s.QueryType = &v
	return s
}

func (s *ListTracesRequest) SetStartTime(v string) *ListTracesRequest {
	s.StartTime = &v
	return s
}

type ListTracesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	Data *ListTracesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EDFF77E1-1ED1-5389-B6A8-651D9433BBE5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListTracesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTracesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTracesResponseBody) SetCode(v string) *ListTracesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTracesResponseBody) SetData(v *ListTracesResponseBodyData) *ListTracesResponseBody {
	s.Data = v
	return s
}

func (s *ListTracesResponseBody) SetDynamicCode(v string) *ListTracesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListTracesResponseBody) SetDynamicMessage(v string) *ListTracesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListTracesResponseBody) SetHttpStatusCode(v int32) *ListTracesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTracesResponseBody) SetMessage(v string) *ListTracesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTracesResponseBody) SetRequestId(v string) *ListTracesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTracesResponseBody) SetSuccess(v bool) *ListTracesResponseBody {
	s.Success = &v
	return s
}

type ListTracesResponseBodyData struct {
	// Trace list.
	List []*ListTracesResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListTracesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTracesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTracesResponseBodyData) SetList(v []*ListTracesResponseBodyDataList) *ListTracesResponseBodyData {
	s.List = v
	return s
}

func (s *ListTracesResponseBodyData) SetPageNumber(v int64) *ListTracesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTracesResponseBodyData) SetPageSize(v int64) *ListTracesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTracesResponseBodyData) SetTotalCount(v int64) *ListTracesResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListTracesResponseBodyDataList struct {
	// Message born time.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	BornTime *string `json:"bornTime,omitempty" xml:"bornTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Message id.
	//
	// example:
	//
	// 7F00000100207A4F0F294A938F7807AE
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// Message keys.
	MessageKeys []*string `json:"messageKeys,omitempty" xml:"messageKeys,omitempty" type:"Repeated"`
	// Message tag.
	//
	// example:
	//
	// xx
	MessageTag *string `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s ListTracesResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListTracesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTracesResponseBodyDataList) SetBornTime(v string) *ListTracesResponseBodyDataList {
	s.BornTime = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetInstanceId(v string) *ListTracesResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetMessageId(v string) *ListTracesResponseBodyDataList {
	s.MessageId = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetMessageKeys(v []*string) *ListTracesResponseBodyDataList {
	s.MessageKeys = v
	return s
}

func (s *ListTracesResponseBodyDataList) SetMessageTag(v string) *ListTracesResponseBodyDataList {
	s.MessageTag = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetRegionId(v string) *ListTracesResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *ListTracesResponseBodyDataList) SetTopicName(v string) *ListTracesResponseBodyDataList {
	s.TopicName = &v
	return s
}

type ListTracesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTracesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTracesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTracesResponse) GoString() string {
	return s.String()
}

func (s *ListTracesResponse) SetHeaders(v map[string]*string) *ListTracesResponse {
	s.Headers = v
	return s
}

func (s *ListTracesResponse) SetStatusCode(v int32) *ListTracesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTracesResponse) SetBody(v *ListTracesResponseBody) *ListTracesResponse {
	s.Body = v
	return s
}

type ResetConsumeOffsetRequest struct {
	// The time when the consumer offset is reset.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	ResetTime *string `json:"resetTime,omitempty" xml:"resetTime,omitempty"`
	// The method that is used to reset the consumer offset. Valid values: LATEST_OFFSET and SPECIFIED_TIME.
	//
	// example:
	//
	// LATEST_OFFSET
	ResetType *string `json:"resetType,omitempty" xml:"resetType,omitempty"`
}

func (s ResetConsumeOffsetRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetConsumeOffsetRequest) GoString() string {
	return s.String()
}

func (s *ResetConsumeOffsetRequest) SetResetTime(v string) *ResetConsumeOffsetRequest {
	s.ResetTime = &v
	return s
}

func (s *ResetConsumeOffsetRequest) SetResetType(v string) *ResetConsumeOffsetRequest {
	s.ResetType = &v
	return s
}

type ResetConsumeOffsetResponseBody struct {
	// The returned error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The returned dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F9A95891-EAD4-5A2B-8A30-676CD18921AF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResetConsumeOffsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetConsumeOffsetResponseBody) GoString() string {
	return s.String()
}

func (s *ResetConsumeOffsetResponseBody) SetCode(v string) *ResetConsumeOffsetResponseBody {
	s.Code = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetDynamicCode(v string) *ResetConsumeOffsetResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetDynamicMessage(v string) *ResetConsumeOffsetResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetHttpStatusCode(v int32) *ResetConsumeOffsetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetMessage(v string) *ResetConsumeOffsetResponseBody {
	s.Message = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetRequestId(v string) *ResetConsumeOffsetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetConsumeOffsetResponseBody) SetSuccess(v bool) *ResetConsumeOffsetResponseBody {
	s.Success = &v
	return s
}

type ResetConsumeOffsetResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetConsumeOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetConsumeOffsetResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetConsumeOffsetResponse) GoString() string {
	return s.String()
}

func (s *ResetConsumeOffsetResponse) SetHeaders(v map[string]*string) *ResetConsumeOffsetResponse {
	s.Headers = v
	return s
}

func (s *ResetConsumeOffsetResponse) SetStatusCode(v int32) *ResetConsumeOffsetResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetConsumeOffsetResponse) SetBody(v *ResetConsumeOffsetResponseBody) *ResetConsumeOffsetResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource IDs, in the JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq-cn-pe3355cs707
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of resource.
	//
	// Set this parameter to **instance**. The value of this parameter cannot be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// tag, in JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"key": "rmq-test", "value": "test"}]
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v string) *TagResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v string) *TagResourcesRequest {
	s.Tag = &v
	return s
}

type TagResourcesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error code.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B962390-D84B-5D44-8C11-79DF40299D41
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetCode(v string) *TagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *TagResourcesResponseBody) SetData(v bool) *TagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *TagResourcesResponseBody) SetDynamicCode(v string) *TagResourcesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetDynamicMessage(v string) *TagResourcesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *TagResourcesResponseBody) SetHttpStatusCode(v int32) *TagResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetMessage(v string) *TagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Whether to delete all tags.
	//
	// example:
	//
	// true
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The resource IDs, in the JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq-cn-pe3355cs707
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of resource.
	//
	// Set this parameter to **instance**. The value of this parameter cannot be changed.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The keys of tags.
	//
	// example:
	//
	// ["key1", "key2"]
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v string) *UntagResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v string) *UntagResourcesRequest {
	s.TagKey = &v
	return s
}

type UntagResourcesResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter deliveryOrderType is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// A07B41BD-6DD3-5349-9E76-00303DF04BBE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetCode(v string) *UntagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *UntagResourcesResponseBody) SetData(v bool) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *UntagResourcesResponseBody) SetDynamicCode(v string) *UntagResourcesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetDynamicMessage(v string) *UntagResourcesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UntagResourcesResponseBody) SetHttpStatusCode(v int32) *UntagResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetMessage(v string) *UntagResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateConsumerGroupRequest struct {
	// The new consumption retry policy that you want to configure for the consumer group. For more information, see [Consumption retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// This parameter is required.
	ConsumeRetryPolicy *UpdateConsumerGroupRequestConsumeRetryPolicy `json:"consumeRetryPolicy,omitempty" xml:"consumeRetryPolicy,omitempty" type:"Struct"`
	// The new message delivery order of the consumer group.
	//
	// Valid values:
	//
	// 	- Concurrently: concurrent delivery
	//
	// 	- Orderly: ordered delivery
	//
	// This parameter is required.
	//
	// example:
	//
	// Concurrently
	DeliveryOrderType *string `json:"deliveryOrderType,omitempty" xml:"deliveryOrderType,omitempty"`
	MaxReceiveTps     *int64  `json:"maxReceiveTps,omitempty" xml:"maxReceiveTps,omitempty"`
	// The new remarks on the consumer group.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequest) SetConsumeRetryPolicy(v *UpdateConsumerGroupRequestConsumeRetryPolicy) *UpdateConsumerGroupRequest {
	s.ConsumeRetryPolicy = v
	return s
}

func (s *UpdateConsumerGroupRequest) SetDeliveryOrderType(v string) *UpdateConsumerGroupRequest {
	s.DeliveryOrderType = &v
	return s
}

func (s *UpdateConsumerGroupRequest) SetMaxReceiveTps(v int64) *UpdateConsumerGroupRequest {
	s.MaxReceiveTps = &v
	return s
}

func (s *UpdateConsumerGroupRequest) SetRemark(v string) *UpdateConsumerGroupRequest {
	s.Remark = &v
	return s
}

type UpdateConsumerGroupRequestConsumeRetryPolicy struct {
	// The dead-letter topic.
	//
	// If a consumer still fails to consume a message after the message is retried for a specified number of times, the message is delivered to a dead-letter topic for subsequent business recovery or troubleshooting. For more information, see [Consumption retry and dead-letter messages](https://help.aliyun.com/document_detail/440356.html).
	//
	// example:
	//
	// DLQ_mqtest
	DeadLetterTargetTopic *string `json:"deadLetterTargetTopic,omitempty" xml:"deadLetterTargetTopic,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 16
	MaxRetryTimes *int32 `json:"maxRetryTimes,omitempty" xml:"maxRetryTimes,omitempty"`
	// The retry policy. For more information, see [Message retry](https://help.aliyun.com/document_detail/440356.html).
	//
	// Valid values:
	//
	// 	- FixedRetryPolicy: Failed messages are retried at a fixed interval.
	//
	// 	- DefaultRetryPolicy: Failed messages are retried at incremental intervals as the number of retries increases.
	//
	// This parameter is required.
	//
	// example:
	//
	// DefaultRetryPolicy
	RetryPolicy *string `json:"retryPolicy,omitempty" xml:"retryPolicy,omitempty"`
}

func (s UpdateConsumerGroupRequestConsumeRetryPolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupRequestConsumeRetryPolicy) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetDeadLetterTargetTopic(v string) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.DeadLetterTargetTopic = &v
	return s
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetMaxRetryTimes(v int32) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.MaxRetryTimes = &v
	return s
}

func (s *UpdateConsumerGroupRequestConsumeRetryPolicy) SetRetryPolicy(v string) *UpdateConsumerGroupRequestConsumeRetryPolicy {
	s.RetryPolicy = &v
	return s
}

type UpdateConsumerGroupResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidDeliveryOrderType
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// DeliveryOrderType
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// deliveryOrderType
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter deliveryOrderType is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// C7F94090-3358-506A-97DC-34BC803C****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponseBody) SetCode(v string) *UpdateConsumerGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetData(v bool) *UpdateConsumerGroupResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetDynamicCode(v string) *UpdateConsumerGroupResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetDynamicMessage(v string) *UpdateConsumerGroupResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetHttpStatusCode(v int32) *UpdateConsumerGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetMessage(v string) *UpdateConsumerGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetRequestId(v string) *UpdateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConsumerGroupResponseBody) SetSuccess(v bool) *UpdateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateConsumerGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponse) SetHeaders(v map[string]*string) *UpdateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerGroupResponse) SetStatusCode(v int32) *UpdateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConsumerGroupResponse) SetBody(v *UpdateConsumerGroupResponseBody) *UpdateConsumerGroupResponse {
	s.Body = v
	return s
}

type UpdateInstanceRequest struct {
	// The access control list for the instance.
	AclInfo *UpdateInstanceRequestAclInfo `json:"aclInfo,omitempty" xml:"aclInfo,omitempty" type:"Struct"`
	// The updated name of the instance.
	//
	// example:
	//
	// test_instance
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The updated network information about the instance.
	NetworkInfo *UpdateInstanceRequestNetworkInfo `json:"networkInfo,omitempty" xml:"networkInfo,omitempty" type:"Struct"`
	// Additional configurations of the instance.
	ProductInfo *UpdateInstanceRequestProductInfo `json:"productInfo,omitempty" xml:"productInfo,omitempty" type:"Struct"`
	// The updated description of the instance.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) SetAclInfo(v *UpdateInstanceRequestAclInfo) *UpdateInstanceRequest {
	s.AclInfo = v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceName(v string) *UpdateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceRequest) SetNetworkInfo(v *UpdateInstanceRequestNetworkInfo) *UpdateInstanceRequest {
	s.NetworkInfo = v
	return s
}

func (s *UpdateInstanceRequest) SetProductInfo(v *UpdateInstanceRequestProductInfo) *UpdateInstanceRequest {
	s.ProductInfo = v
	return s
}

func (s *UpdateInstanceRequest) SetRemark(v string) *UpdateInstanceRequest {
	s.Remark = &v
	return s
}

type UpdateInstanceRequestAclInfo struct {
	// The authentication type of the instance.
	AclTypes []*string `json:"aclTypes,omitempty" xml:"aclTypes,omitempty" type:"Repeated"`
	// Indicates whether the authentication-free in VPCs feature is enabled.
	//
	// Indicates whether the authentication-free in VPCs feature is enabled.
	//
	// Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	DefaultVpcAuthFree *bool `json:"defaultVpcAuthFree,omitempty" xml:"defaultVpcAuthFree,omitempty"`
}

func (s UpdateInstanceRequestAclInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestAclInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestAclInfo) SetAclTypes(v []*string) *UpdateInstanceRequestAclInfo {
	s.AclTypes = v
	return s
}

func (s *UpdateInstanceRequestAclInfo) SetDefaultVpcAuthFree(v bool) *UpdateInstanceRequestAclInfo {
	s.DefaultVpcAuthFree = &v
	return s
}

type UpdateInstanceRequestNetworkInfo struct {
	// The information about the Internet over which the instance is accessed. This parameter takes effect only if the Internet access feature is enabled for the instance.
	InternetInfo *UpdateInstanceRequestNetworkInfoInternetInfo `json:"internetInfo,omitempty" xml:"internetInfo,omitempty" type:"Struct"`
}

func (s UpdateInstanceRequestNetworkInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestNetworkInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestNetworkInfo) SetInternetInfo(v *UpdateInstanceRequestNetworkInfoInternetInfo) *UpdateInstanceRequestNetworkInfo {
	s.InternetInfo = v
	return s
}

type UpdateInstanceRequestNetworkInfoInternetInfo struct {
	// The whitelist that includes the IP addresses that are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If you do not configure an IP address whitelist, all CIDR blocks are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If you configure an IP address whitelist, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	IpWhitelist []*string `json:"ipWhitelist,omitempty" xml:"ipWhitelist,omitempty" type:"Repeated"`
}

func (s UpdateInstanceRequestNetworkInfoInternetInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestNetworkInfoInternetInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestNetworkInfoInternetInfo) SetIpWhitelist(v []*string) *UpdateInstanceRequestNetworkInfoInternetInfo {
	s.IpWhitelist = v
	return s
}

type UpdateInstanceRequestProductInfo struct {
	// Specifies whether to enable the elastic transactions per second (TPS) feature for the instance.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// After you enable the elastic TPS feature for an ApsaraMQ for RocketMQ instance, you can use a specific number of TPS that exceeds the specification limit. You are charged for using the elastic TPS feature. For more information, see [Computing fees](https://help.aliyun.com/document_detail/427237.html).
	//
	// >  The elastic TPS feature is supported only by specific instance editions. For more information, see [Instance editions](https://help.aliyun.com/document_detail/444715.html).
	//
	// example:
	//
	// true
	AutoScaling *bool `json:"autoScaling,omitempty" xml:"autoScaling,omitempty"`
	// The retention period of messages. Unit: hours.
	//
	// For information about the valid values of this parameter, see the "Limits on resource quotas" section of the [Limits](https://help.aliyun.com/document_detail/440347.html) topic.
	//
	// ApsaraMQ for RocketMQ supports serverless scaling of message storage. You are charged storage fees based on your actual storage usage. You can change the retention period of messages to manage storage capacity. For more information, see [Storage fees](https://help.aliyun.com/document_detail/427238.html).
	//
	// example:
	//
	// 72
	MessageRetentionTime *int32 `json:"messageRetentionTime,omitempty" xml:"messageRetentionTime,omitempty"`
	// The ratio of the number of messages that you can send to the number of messages that you can receive on the instance.
	//
	// Value values: 0.25 to 1.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
	// Specifies whether to enable the message trace feature.
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is not in use. By default, the message trace feature is enabled for ApsaraMQ for RocketMQ instances, regardless of whether this parameter is configured.
	//
	// example:
	//
	// true
	TraceOn *bool `json:"traceOn,omitempty" xml:"traceOn,omitempty"`
}

func (s UpdateInstanceRequestProductInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceRequestProductInfo) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequestProductInfo) SetAutoScaling(v bool) *UpdateInstanceRequestProductInfo {
	s.AutoScaling = &v
	return s
}

func (s *UpdateInstanceRequestProductInfo) SetMessageRetentionTime(v int32) *UpdateInstanceRequestProductInfo {
	s.MessageRetentionTime = &v
	return s
}

func (s *UpdateInstanceRequestProductInfo) SetSendReceiveRatio(v float32) *UpdateInstanceRequestProductInfo {
	s.SendReceiveRatio = &v
	return s
}

func (s *UpdateInstanceRequestProductInfo) SetTraceOn(v bool) *UpdateInstanceRequestProductInfo {
	s.TraceOn = &v
	return s
}

type UpdateInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// AA87DE09-DA44-52F4-9515-78B1B607****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponseBody) SetCode(v string) *UpdateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetData(v bool) *UpdateInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetDynamicCode(v string) *UpdateInstanceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetDynamicMessage(v string) *UpdateInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetMessage(v string) *UpdateInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetRequestId(v string) *UpdateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceResponseBody) SetSuccess(v bool) *UpdateInstanceResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResponse) SetHeaders(v map[string]*string) *UpdateInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceResponse) SetStatusCode(v int32) *UpdateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceResponse) SetBody(v *UpdateInstanceResponseBody) *UpdateInstanceResponse {
	s.Body = v
	return s
}

type UpdateInstanceAccountRequest struct {
	// The status of the account.
	//
	// Valid values:
	//
	// 	- DISABLE
	//
	// 	- ENABLE
	//
	// example:
	//
	// ENABLE
	AccountStatus *string `json:"accountStatus,omitempty" xml:"accountStatus,omitempty"`
	// The password of the account.
	//
	// example:
	//
	// test
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s UpdateInstanceAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAccountRequest) SetAccountStatus(v string) *UpdateInstanceAccountRequest {
	s.AccountStatus = &v
	return s
}

func (s *UpdateInstanceAccountRequest) SetPassword(v string) *UpdateInstanceAccountRequest {
	s.Password = &v
	return s
}

type UpdateInstanceAccountResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// Instance.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInstanceAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAccountResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceAccountResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetCode(v string) *UpdateInstanceAccountResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetData(v bool) *UpdateInstanceAccountResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetDynamicCode(v string) *UpdateInstanceAccountResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetDynamicMessage(v string) *UpdateInstanceAccountResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceAccountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetMessage(v string) *UpdateInstanceAccountResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetRequestId(v string) *UpdateInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceAccountResponseBody) SetSuccess(v bool) *UpdateInstanceAccountResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAccountResponse) SetHeaders(v map[string]*string) *UpdateInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAccountResponse) SetStatusCode(v int32) *UpdateInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAccountResponse) SetBody(v *UpdateInstanceAccountResponseBody) *UpdateInstanceAccountResponse {
	s.Body = v
	return s
}

type UpdateInstanceAclRequest struct {
	// The following items describe the types of permissions that can be granted based on the resource type:
	//
	// 	- Topic: Pub, Sub, and Pub|Sub
	//
	// 	- Consumer group: Sub
	//
	// Valid values:
	//
	// 	- SUB: subscribe
	//
	// 	- Pub|Sub: publish and subscribe
	//
	// 	- Pub: publish
	//
	// example:
	//
	// Pub
	Actions *string `json:"actions,omitempty" xml:"actions,omitempty"`
	// The decision result of the authorization.
	//
	// Valid values:
	//
	// 	- Deny
	//
	// 	- Allow
	//
	// example:
	//
	// Allow
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// The IP address whitelists.
	IpWhitelists []*string `json:"ipWhitelists,omitempty" xml:"ipWhitelists,omitempty" type:"Repeated"`
	// The name of the resource on which you want to grant permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The type of the resource on which you want to grant permissions.
	//
	// Valid values:
	//
	// 	- Group
	//
	// 	- Topic
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s UpdateInstanceAclRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAclRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAclRequest) SetActions(v string) *UpdateInstanceAclRequest {
	s.Actions = &v
	return s
}

func (s *UpdateInstanceAclRequest) SetDecision(v string) *UpdateInstanceAclRequest {
	s.Decision = &v
	return s
}

func (s *UpdateInstanceAclRequest) SetIpWhitelists(v []*string) *UpdateInstanceAclRequest {
	s.IpWhitelists = v
	return s
}

func (s *UpdateInstanceAclRequest) SetResourceName(v string) *UpdateInstanceAclRequest {
	s.ResourceName = &v
	return s
}

func (s *UpdateInstanceAclRequest) SetResourceType(v string) *UpdateInstanceAclRequest {
	s.ResourceType = &v
	return s
}

type UpdateInstanceAclResponseBody struct {
	// The details about the access denial. This parameter is returned only if the access is denied because the Resource Access Management (RAM) user does not have the required permissions.
	//
	// example:
	//
	// xxx
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C115601B-8736-5BBF-AC99-7FEAE1245A80
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateInstanceAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAclResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAclResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceAclResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetCode(v string) *UpdateInstanceAclResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetData(v bool) *UpdateInstanceAclResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetDynamicCode(v string) *UpdateInstanceAclResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetDynamicMessage(v string) *UpdateInstanceAclResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceAclResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetMessage(v string) *UpdateInstanceAclResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetRequestId(v string) *UpdateInstanceAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceAclResponseBody) SetSuccess(v bool) *UpdateInstanceAclResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceAclResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceAclResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAclResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAclResponse) SetHeaders(v map[string]*string) *UpdateInstanceAclResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAclResponse) SetStatusCode(v int32) *UpdateInstanceAclResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAclResponse) SetBody(v *UpdateInstanceAclResponseBody) *UpdateInstanceAclResponse {
	s.Body = v
	return s
}

type UpdateTopicRequest struct {
	MaxSendTps *int64 `json:"maxSendTps,omitempty" xml:"maxSendTps,omitempty"`
	// The new remarks on the topic.
	//
	// example:
	//
	// This is the remark for test.
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s UpdateTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicRequest) GoString() string {
	return s.String()
}

func (s *UpdateTopicRequest) SetMaxSendTps(v int64) *UpdateTopicRequest {
	s.MaxSendTps = &v
	return s
}

func (s *UpdateTopicRequest) SetRemark(v string) *UpdateTopicRequest {
	s.Remark = &v
	return s
}

type UpdateTopicResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// TopicName
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// topicName
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The topic cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. The system generates a unique ID for each request. You can troubleshoot issues based on the request ID.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTopicResponseBody) SetCode(v string) *UpdateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTopicResponseBody) SetData(v bool) *UpdateTopicResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTopicResponseBody) SetDynamicCode(v string) *UpdateTopicResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *UpdateTopicResponseBody) SetDynamicMessage(v string) *UpdateTopicResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpdateTopicResponseBody) SetHttpStatusCode(v int32) *UpdateTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTopicResponseBody) SetMessage(v string) *UpdateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTopicResponseBody) SetRequestId(v string) *UpdateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTopicResponseBody) SetSuccess(v bool) *UpdateTopicResponseBody {
	s.Success = &v
	return s
}

type UpdateTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicResponse) GoString() string {
	return s.String()
}

func (s *UpdateTopicResponse) SetHeaders(v map[string]*string) *UpdateTopicResponse {
	s.Headers = v
	return s
}

func (s *UpdateTopicResponse) SetStatusCode(v int32) *UpdateTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTopicResponse) SetBody(v *UpdateTopicResponseBody) *UpdateTopicResponse {
	s.Body = v
	return s
}

type VerifyConsumeMessageRequest struct {
	// The client ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// zeekr-settlement-server-dc555456f-v2lcg@1@1@qfvorazqns
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// The ID of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST_FINANCE_STOCK_OUT_GROUP
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
}

func (s VerifyConsumeMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyConsumeMessageRequest) GoString() string {
	return s.String()
}

func (s *VerifyConsumeMessageRequest) SetClientId(v string) *VerifyConsumeMessageRequest {
	s.ClientId = &v
	return s
}

func (s *VerifyConsumeMessageRequest) SetConsumerGroupId(v string) *VerifyConsumeMessageRequest {
	s.ConsumerGroupId = &v
	return s
}

type VerifyConsumeMessageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5304143F-AB0E-5AB4-A227-7C5489216FD5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifyConsumeMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyConsumeMessageResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyConsumeMessageResponseBody) SetCode(v string) *VerifyConsumeMessageResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetData(v bool) *VerifyConsumeMessageResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetDynamicCode(v string) *VerifyConsumeMessageResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetDynamicMessage(v string) *VerifyConsumeMessageResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetHttpStatusCode(v int32) *VerifyConsumeMessageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetMessage(v string) *VerifyConsumeMessageResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetRequestId(v string) *VerifyConsumeMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyConsumeMessageResponseBody) SetSuccess(v bool) *VerifyConsumeMessageResponseBody {
	s.Success = &v
	return s
}

type VerifyConsumeMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyConsumeMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyConsumeMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyConsumeMessageResponse) GoString() string {
	return s.String()
}

func (s *VerifyConsumeMessageResponse) SetHeaders(v map[string]*string) *VerifyConsumeMessageResponse {
	s.Headers = v
	return s
}

func (s *VerifyConsumeMessageResponse) SetStatusCode(v int32) *VerifyConsumeMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyConsumeMessageResponse) SetBody(v *VerifyConsumeMessageResponseBody) *VerifyConsumeMessageResponse {
	s.Body = v
	return s
}

type VerifySendMessageRequest struct {
	// The message body.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The message key.
	//
	// example:
	//
	// xx
	MessageKey *string `json:"messageKey,omitempty" xml:"messageKey,omitempty"`
	// The message tag.
	//
	// example:
	//
	// xx
	MessageTag *string `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
}

func (s VerifySendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifySendMessageRequest) GoString() string {
	return s.String()
}

func (s *VerifySendMessageRequest) SetMessage(v string) *VerifySendMessageRequest {
	s.Message = &v
	return s
}

func (s *VerifySendMessageRequest) SetMessageKey(v string) *VerifySendMessageRequest {
	s.MessageKey = &v
	return s
}

func (s *VerifySendMessageRequest) SetMessageTag(v string) *VerifySendMessageRequest {
	s.MessageTag = &v
	return s
}

type VerifySendMessageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 0A64228900207A4F0F2931A4E0D40BE5
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// consumerGroupId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3BD2C19B-66DE-59C7-B2F6-FD1BE21DC8C1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifySendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifySendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *VerifySendMessageResponseBody) SetCode(v string) *VerifySendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetData(v string) *VerifySendMessageResponseBody {
	s.Data = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetDynamicCode(v string) *VerifySendMessageResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetDynamicMessage(v string) *VerifySendMessageResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetHttpStatusCode(v int32) *VerifySendMessageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetMessage(v string) *VerifySendMessageResponseBody {
	s.Message = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetRequestId(v string) *VerifySendMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifySendMessageResponseBody) SetSuccess(v bool) *VerifySendMessageResponseBody {
	s.Success = &v
	return s
}

type VerifySendMessageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifySendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifySendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifySendMessageResponse) GoString() string {
	return s.String()
}

func (s *VerifySendMessageResponse) SetHeaders(v map[string]*string) *VerifySendMessageResponse {
	s.Headers = v
	return s
}

func (s *VerifySendMessageResponse) SetStatusCode(v int32) *VerifySendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySendMessageResponse) SetBody(v *VerifySendMessageResponseBody) *VerifySendMessageResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("rocketmq"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加容灾计划条目
//
// @param request - AddDisasterRecoveryItemRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDisasterRecoveryItemResponse
func (client *Client) AddDisasterRecoveryItemWithOptions(planId *string, request *AddDisasterRecoveryItemRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddDisasterRecoveryItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Topics)) {
		body["topics"] = request.Topics
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddDisasterRecoveryItem"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/disaster_recovery/" + tea.StringValue(openapiutil.GetEncodeParam(planId)) + "/items"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddDisasterRecoveryItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加容灾计划条目
//
// @param request - AddDisasterRecoveryItemRequest
//
// @return AddDisasterRecoveryItemResponse
func (client *Client) AddDisasterRecoveryItem(planId *string, request *AddDisasterRecoveryItemRequest) (_result *AddDisasterRecoveryItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/resourceGroup/change"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - CreateConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, request *CreateConsumerGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumeRetryPolicy)) {
		body["consumeRetryPolicy"] = request.ConsumeRetryPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryOrderType)) {
		body["deliveryOrderType"] = request.DeliveryOrderType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxReceiveTps)) {
		body["maxReceiveTps"] = request.MaxReceiveTps
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConsumerGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - CreateConsumerGroupRequest
//
// @return CreateConsumerGroupResponse
func (client *Client) CreateConsumerGroup(instanceId *string, consumerGroupId *string, request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["autoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		body["autoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["commodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInfo)) {
		body["networkInfo"] = request.NetworkInfo
	}

	if !tea.BoolValue(util.IsUnset(request.PaymentType)) {
		body["paymentType"] = request.PaymentType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["periodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ProductInfo)) {
		body["productInfo"] = request.ProductInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SeriesCode)) {
		body["seriesCode"] = request.SeriesCode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["serviceCode"] = request.ServiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.SubSeriesCode)) {
		body["subSeriesCode"] = request.SubSeriesCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateInstanceAccountWithOptions(instanceId *string, request *CreateInstanceAccountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceAccount"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/accounts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateInstanceAclWithOptions(instanceId *string, username *string, request *CreateInstanceAclRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Actions)) {
		body["actions"] = request.Actions
	}

	if !tea.BoolValue(util.IsUnset(request.Decision)) {
		body["decision"] = request.Decision
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhitelists)) {
		body["ipWhitelists"] = request.IpWhitelists
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		body["resourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceAcl"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/acl/account/" + tea.StringValue(openapiutil.GetEncodeParam(username))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) CreateInstanceIpWhitelistWithOptions(instanceId *string, request *CreateInstanceIpWhitelistRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInstanceIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpWhitelists)) {
		body["ipWhitelists"] = request.IpWhitelists
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceIpWhitelist"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/ip/whitelist"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Creates a topic.
//
// @param request - CreateTopicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTopicResponse
func (client *Client) CreateTopicWithOptions(instanceId *string, topicName *string, request *CreateTopicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxSendTps)) {
		body["maxSendTps"] = request.MaxSendTps
	}

	if !tea.BoolValue(util.IsUnset(request.MessageType)) {
		body["messageType"] = request.MessageType
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTopic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a topic.
//
// @param request - CreateTopicRequest
//
// @return CreateTopicResponse
func (client *Client) CreateTopic(instanceId *string, topicName *string, request *CreateTopicRequest) (_result *CreateTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) DeleteConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConsumerGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) DeleteConsumerGroupSubscriptionWithOptions(instanceId *string, consumerGroupId *string, request *DeleteConsumerGroupSubscriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteConsumerGroupSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterExpression)) {
		query["filterExpression"] = request.FilterExpression
	}

	if !tea.BoolValue(util.IsUnset(request.FilterType)) {
		query["filterType"] = request.FilterType
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["topicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConsumerGroupSubscription"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId)) + "/subscriptions"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConsumerGroupSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Deletes a ApsaraMQ for RocketMQ instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// 	- After an instance is deleted, the instance cannot be restored. Exercise caution when you call this operation.
//
// 	- This operation is used to delete a pay-as-you-go instance. A subscription instance is automatically released after it expires. You do not need to manually delete a subscription instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
// 	- After an instance is deleted, the instance cannot be restored. Exercise caution when you call this operation.
//
// 	- This operation is used to delete a pay-as-you-go instance. A subscription instance is automatically released after it expires. You do not need to manually delete a subscription instance.
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(instanceId *string) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Delete access control ACL user
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceAccountResponse
func (client *Client) DeleteInstanceAccountWithOptions(instanceId *string, username *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceAccountResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceAccount"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/accounts/" + tea.StringValue(openapiutil.GetEncodeParam(username))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete access control ACL user
//
// @return DeleteInstanceAccountResponse
func (client *Client) DeleteInstanceAccount(instanceId *string, username *string) (_result *DeleteInstanceAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) DeleteInstanceAclWithOptions(instanceId *string, username *string, request *DeleteInstanceAclRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["resourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceAcl"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/acl/account/" + tea.StringValue(openapiutil.GetEncodeParam(username))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// @param request - DeleteInstanceIpWhitelistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceIpWhitelistResponse
func (client *Client) DeleteInstanceIpWhitelistWithOptions(instanceId *string, request *DeleteInstanceIpWhitelistRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteInstanceIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpWhitelist)) {
		query["ipWhitelist"] = request.IpWhitelist
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceIpWhitelist"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/ip/whitelist"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) DeleteTopicWithOptions(instanceId *string, topicName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTopic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetConsumerGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsumerGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Query Consumer Group Backlog Information
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerGroupLagResponse
func (client *Client) GetConsumerGroupLagWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetConsumerGroupLagResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsumerGroupLag"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId)) + "/lag"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConsumerGroupLagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query Consumer Group Backlog Information
//
// @return GetConsumerGroupLagResponse
func (client *Client) GetConsumerGroupLag(instanceId *string, consumerGroupId *string) (_result *GetConsumerGroupLagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConsumerGroupLagResponse{}
	_body, _err := client.GetConsumerGroupLagWithOptions(instanceId, consumerGroupId, headers, runtime)
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
func (client *Client) GetConsumerGroupSubscriptionWithOptions(instanceId *string, consumerGroupId *string, topicName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetConsumerGroupSubscriptionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsumerGroupSubscription"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId)) + "/subscriptions/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConsumerGroupSubscriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscriptions of a consumer group.
//
// @return GetConsumerGroupSubscriptionResponse
func (client *Client) GetConsumerGroupSubscription(instanceId *string, consumerGroupId *string, topicName *string) (_result *GetConsumerGroupSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetConsumerStackWithOptions(instanceId *string, consumerGroupId *string, request *GetConsumerStackRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetConsumerStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["clientId"] = request.ClientId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsumerStack"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId)) + "/stack"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConsumerStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetInstanceWithOptions(instanceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) GetInstanceAccountWithOptions(instanceId *string, request *GetInstanceAccountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceAccount"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/account"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Obtains the details of a specific message.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMessageDetailResponse
func (client *Client) GetMessageDetailWithOptions(instanceId *string, topicName *string, messageId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetMessageDetailResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetMessageDetail"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName)) + "/messages/" + tea.StringValue(openapiutil.GetEncodeParam(messageId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMessageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a specific message.
//
// @return GetMessageDetailResponse
func (client *Client) GetMessageDetail(instanceId *string, topicName *string, messageId *string) (_result *GetMessageDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Queries the details of a specified topic.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicResponse
func (client *Client) GetTopicWithOptions(instanceId *string, topicName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTopicResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified topic.
//
// @return GetTopicResponse
func (client *Client) GetTopic(instanceId *string, topicName *string) (_result *GetTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTraceResponse
func (client *Client) GetTraceWithOptions(instanceId *string, topicName *string, messageId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTraceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrace"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName)) + "/traces/" + tea.StringValue(openapiutil.GetEncodeParam(messageId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTraceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trace of a specific message in a specific topic.
//
// @return GetTraceResponse
func (client *Client) GetTrace(instanceId *string, topicName *string, messageId *string) (_result *GetTraceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTraceResponse{}
	_body, _err := client.GetTraceWithOptions(instanceId, topicName, messageId, headers, runtime)
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
func (client *Client) ListAvailableZonesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAvailableZonesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListAvailableZones"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/zones"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAvailableZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询支持的可用区
//
// @return ListAvailableZonesResponse
func (client *Client) ListAvailableZones() (_result *ListAvailableZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerConnectionsResponse
func (client *Client) ListConsumerConnectionsWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConsumerConnectionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListConsumerConnections"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId)) + "/connections"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConsumerConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询消费者客户端连接信息
//
// @return ListConsumerConnectionsResponse
func (client *Client) ListConsumerConnections(instanceId *string, consumerGroupId *string) (_result *ListConsumerConnectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerConnectionsResponse{}
	_body, _err := client.ListConsumerConnectionsWithOptions(instanceId, consumerGroupId, headers, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupSubscriptionsResponse
func (client *Client) ListConsumerGroupSubscriptionsWithOptions(instanceId *string, consumerGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConsumerGroupSubscriptionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListConsumerGroupSubscriptions"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId)) + "/subscriptions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConsumerGroupSubscriptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscriptions of a specific consumer group.
//
// @return ListConsumerGroupSubscriptionsResponse
func (client *Client) ListConsumerGroupSubscriptions(instanceId *string, consumerGroupId *string) (_result *ListConsumerGroupSubscriptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConsumerGroupSubscriptionsResponse{}
	_body, _err := client.ListConsumerGroupSubscriptionsWithOptions(instanceId, consumerGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the consumer groups in a specified instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - ListConsumerGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConsumerGroupsResponse
func (client *Client) ListConsumerGroupsWithOptions(instanceId *string, request *ListConsumerGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListConsumerGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConsumerGroups"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConsumerGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the consumer groups in a specified instance.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - ListConsumerGroupsRequest
//
// @return ListConsumerGroupsResponse
func (client *Client) ListConsumerGroups(instanceId *string, request *ListConsumerGroupsRequest) (_result *ListConsumerGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Queries the accounts that are used to access a specific instance.
//
// @param request - ListInstanceAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAccountResponse
func (client *Client) ListInstanceAccountWithOptions(instanceId *string, request *ListInstanceAccountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountStatus)) {
		query["accountStatus"] = request.AccountStatus
	}

	if !tea.BoolValue(util.IsUnset(request.AccountType)) {
		query["accountType"] = request.AccountType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceAccount"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/accounts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListInstanceAclWithOptions(instanceId *string, request *ListInstanceAclRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceAcl"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/acl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListInstanceIpWhitelistWithOptions(instanceId *string, request *ListInstanceIpWhitelistRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceIpWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpWhitelist)) {
		query["ipWhitelist"] = request.IpWhitelist
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceIpWhitelist"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/ip/whitelist"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceIpWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListInstancesWithOptions(tmpReq *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SeriesCodes)) {
		request.SeriesCodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SeriesCodes, tea.String("seriesCodes"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SeriesCodesShrink)) {
		query["seriesCodes"] = request.SeriesCodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StorageSecretKey)) {
		query["storageSecretKey"] = request.StorageSecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListMessagesWithOptions(instanceId *string, topicName *string, request *ListMessagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMessagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageKey)) {
		query["messageKey"] = request.MessageKey
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ScrollId)) {
		query["scrollId"] = request.ScrollId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMessages"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName)) + "/messages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Queries regions in which ApsaraMQ for RocketMQ is available.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/regions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regions in which ApsaraMQ for RocketMQ is available.
//
// @return ListRegionsResponse
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Query visible resource tag relationships
//
// @param request - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/resourceTag/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query visible resource tag relationships
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListTopicSubscriptionsWithOptions(instanceId *string, topicName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTopicSubscriptionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListTopicSubscriptions"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName)) + "/subscriptions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTopicSubscriptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscriptions of a specific topic.
//
// @return ListTopicSubscriptionsResponse
func (client *Client) ListTopicSubscriptions(instanceId *string, topicName *string) (_result *ListTopicSubscriptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Queries the topics in a specified instance.
//
// @param tmpReq - ListTopicsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicsResponse
func (client *Client) ListTopicsWithOptions(instanceId *string, tmpReq *ListTopicsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTopicsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTopicsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MessageTypes)) {
		request.MessageTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MessageTypes, tea.String("messageTypes"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTypesShrink)) {
		query["messageTypes"] = request.MessageTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTopics"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTopicsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topics in a specified instance.
//
// @param request - ListTopicsRequest
//
// @return ListTopicsResponse
func (client *Client) ListTopics(instanceId *string, request *ListTopicsRequest) (_result *ListTopicsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ListTracesWithOptions(instanceId *string, topicName *string, request *ListTracesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTracesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MessageId)) {
		query["messageId"] = request.MessageId
	}

	if !tea.BoolValue(util.IsUnset(request.MessageKey)) {
		query["messageKey"] = request.MessageKey
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryType)) {
		query["queryType"] = request.QueryType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTraces"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName)) + "/traces"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTracesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) ResetConsumeOffsetWithOptions(instanceId *string, consumerGroupId *string, topicName *string, request *ResetConsumeOffsetRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ResetConsumeOffsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResetTime)) {
		body["resetTime"] = request.ResetTime
	}

	if !tea.BoolValue(util.IsUnset(request.ResetType)) {
		body["resetType"] = request.ResetType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetConsumeOffset"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId)) + "/consumeOffsets/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetConsumeOffsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Creates resource tags.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/resourceTag/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["all"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["regionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["tagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/resourceTag/delete"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Updates the basic information about and the consumption retry policy of a consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - UpdateConsumerGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConsumerGroupResponse
func (client *Client) UpdateConsumerGroupWithOptions(instanceId *string, consumerGroupId *string, request *UpdateConsumerGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumeRetryPolicy)) {
		body["consumeRetryPolicy"] = request.ConsumeRetryPolicy
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryOrderType)) {
		body["deliveryOrderType"] = request.DeliveryOrderType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxReceiveTps)) {
		body["maxReceiveTps"] = request.MaxReceiveTps
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConsumerGroup"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/consumerGroups/" + tea.StringValue(openapiutil.GetEncodeParam(consumerGroupId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConsumerGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about and the consumption retry policy of a consumer group.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - UpdateConsumerGroupRequest
//
// @return UpdateConsumerGroupResponse
func (client *Client) UpdateConsumerGroup(instanceId *string, consumerGroupId *string, request *UpdateConsumerGroupRequest) (_result *UpdateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpdateInstanceWithOptions(instanceId *string, request *UpdateInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclInfo)) {
		body["aclInfo"] = request.AclInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkInfo)) {
		body["networkInfo"] = request.NetworkInfo
	}

	if !tea.BoolValue(util.IsUnset(request.ProductInfo)) {
		body["productInfo"] = request.ProductInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstance"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpdateInstanceAccountWithOptions(instanceId *string, username *string, request *UpdateInstanceAccountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountStatus)) {
		query["accountStatus"] = request.AccountStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceAccount"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/accounts/" + tea.StringValue(openapiutil.GetEncodeParam(username))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) UpdateInstanceAclWithOptions(instanceId *string, username *string, request *UpdateInstanceAclRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInstanceAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Actions)) {
		body["actions"] = request.Actions
	}

	if !tea.BoolValue(util.IsUnset(request.Decision)) {
		body["decision"] = request.Decision
	}

	if !tea.BoolValue(util.IsUnset(request.IpWhitelists)) {
		body["ipWhitelists"] = request.IpWhitelists
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		body["resourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceAcl"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/acl/account/" + tea.StringValue(openapiutil.GetEncodeParam(username))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
// Updates the basic information about a topic.
//
// @param request - UpdateTopicRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTopicResponse
func (client *Client) UpdateTopicWithOptions(instanceId *string, topicName *string, request *UpdateTopicRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxSendTps)) {
		body["maxSendTps"] = request.MaxSendTps
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		body["remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTopic"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName))),
		Method:      tea.String("PATCH"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the basic information about a topic.
//
// @param request - UpdateTopicRequest
//
// @return UpdateTopicResponse
func (client *Client) UpdateTopic(instanceId *string, topicName *string, request *UpdateTopicRequest) (_result *UpdateTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
func (client *Client) VerifyConsumeMessageWithOptions(instanceId *string, topicName *string, messageId *string, request *VerifyConsumeMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VerifyConsumeMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["clientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ConsumerGroupId)) {
		query["consumerGroupId"] = request.ConsumerGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyConsumeMessage"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName)) + "/messages/" + tea.StringValue(openapiutil.GetEncodeParam(messageId)) + "/action/verifyConsume"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyConsumeMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
func (client *Client) VerifySendMessageWithOptions(instanceId *string, topicName *string, request *VerifySendMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VerifySendMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Message)) {
		body["message"] = request.Message
	}

	if !tea.BoolValue(util.IsUnset(request.MessageKey)) {
		body["messageKey"] = request.MessageKey
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTag)) {
		body["messageTag"] = request.MessageTag
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifySendMessage"),
		Version:     tea.String("2022-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/topics/" + tea.StringValue(openapiutil.GetEncodeParam(topicName)) + "/messages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifySendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifySendMessageResponse{}
	_body, _err := client.VerifySendMessageWithOptions(instanceId, topicName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
