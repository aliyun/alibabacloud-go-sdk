// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// Specifies whether to enable auto-renewal for the instance. This parameter takes effect only if you set PaymentType to Subscription.
	//
	// 	- true: enable
	//
	// 	- false: disable
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
	// The commodity code.
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
	// If you do not configure this parameter, the instance ID is used as the instance name.
	//
	// example:
	//
	// rmq-cn-72u3048uxxx
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The network configurations.
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
	// The subscription duration of the instance. This parameter takes effect only if you set PaymentType to Subscription.
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
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzy6pist7uuna
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
	// >  After an instance is created, you can only upgrade the primary edition of the instance. The following editions are sorted in ascending order: Standard Edition, Professional Edition, Enterprise Platinum Edition. For example, you can upgrade an instance of Standard Edition to Professional Edition, but cannot downgrade an instance of Professional Edition to Standard Edition.
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
	// The sub-category edition of the instance. For information about the differences between sub-category edition instances, see [Instance selection](https://help.aliyun.com/document_detail/444722.html).
	//
	// Valid values:
	//
	// 	- cluster_ha: High-availability Cluster Edition
	//
	// 	- single_node: Standalone Edition
	//
	// 	- serverless: Serverless Edition
	//
	// If you set seriesCode to ultimate, you can set this parameter to only cluster_ha.
	//
	// >  After an instance is created, you cannot change the sub-category edition of the instance.
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
	// 	- payByBandwidth: pay-by-bandwidth. If Internet access is enabled for an instance, specify this value for the parameter.
	//
	// 	- payByTraffic: pay-by-traffic. If Internet access is enabled for an instance, specify this value for the parameter.
	//
	// 	- uninvolved: No billing method is involved. If Internet access is disabled for an instance, specify this value for the parameter.
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
	// By default, ApsaraMQ for RocketMQ allows you to access instances in VPCs. If you enable Internet access for an instance, you can access the instance over the Internet. After you enable this feature, you are charged for outbound Internet traffic. For more information, see [Internet access fee](https://help.aliyun.com/document_detail/427240.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// disable
	InternetSpec *string `json:"internetSpec,omitempty" xml:"internetSpec,omitempty"`
	// The whitelist that includes the IP addresses that are allowed to access the ApsaraMQ for RocketMQ broker over the Internet. This parameter can be configured only if you use a public endpoint to access the instance.
	//
	// 	- If you do not configure an IP address whitelist, all CIDR blocks are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
	//
	// 	- If configure an IP address whitlist, only the IP addresses in the whitelist are allowed to access the ApsaraMQ for RocketMQ broker over the Internet.
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
	// The ID of the vSwitch with which you want to associate the instance, If there are multiple vSwitchs, please concatenate them using the "|" character.
	//
	// >  After an ApsaraMQ for RocketMQ instance is created, you cannot change the vSwitch with which the instance is associated. If you want to change the vSwitch with which the instance is associated, you must release the instance and purchase a new instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6gwtbn6etadpv*******
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the VPC in which you want to deploy the instance.
	//
	// >  After an ApsaraMQ for RocketMQ instance is created, you cannot change the VPC in which the instance is deployed. If you want to change the VPC in which the instance is deployed, you must release the instance and create a new instance.
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

func (s *CreateInstanceRequestNetworkInfoVpcInfo) SetVpcId(v string) *CreateInstanceRequestNetworkInfoVpcInfo {
	s.VpcId = &v
	return s
}

type CreateInstanceRequestProductInfo struct {
	// Specifies whether to enable the elastic TPS feature for the instance.
	//
	// Valid values:
	//
	// 	- true: enable
	//
	// 	- false: disable
	//
	// After you enable the elastic TPS feature for an ApsaraMQ for RocketMQ instance, you can use a specific number of TPS that exceeds the specification limit. You are charged for using the elastic TPS feature. For more information, see [Computing fees](https://help.aliyun.com/document_detail/427237.html).
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
	// The computing specification that specifies the messaging transactions per second (TPS) of the instance. For information about computing specification limits, see [Instance specifications](https://help.aliyun.com/document_detail/444715.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rmq.s2.2xlarge
	MsgProcessSpec *string `json:"msgProcessSpec,omitempty" xml:"msgProcessSpec,omitempty"`
	// The proportion of message sending TPS to the messaging TPS on the instance.
	//
	// For example, you create an instance whose peak messaging TPS is specified as 1,000 and the proportion of message sending TPS is specified as 0.8. In this case, the peak message sending TPS is 800 and the peak message receiving TPS is 200 on the instance.
	//
	// Valid values: 0 to 1. Default value: 0.5.
	//
	// example:
	//
	// 0.5
	SendReceiveRatio *float32 `json:"sendReceiveRatio,omitempty" xml:"sendReceiveRatio,omitempty"`
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

type CreateTopicRequest struct {
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

type GetInstanceResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
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
	// 	- PayAsYouGo: pay-as-you-go
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
	// Valid values:
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
	// The metering method for Internet usage.
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
	// The information about endpoints.
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
	// - TCP_VPC: VPC endpoint
	//
	// - TCP_INTERNET:public endpoint
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
	VSwitchId *string                                                   `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
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
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	ZoneId    *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
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

type ListAvailableZonesResponseBody struct {
	// example:
	//
	// Topic.NotFound
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListAvailableZonesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// cn-qingdao-b
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
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
	// example:
	//
	// MissingPageNumber
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListConsumerConnectionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A3620115-6F1F-5CFB-AA3F-BBD4853B2EC4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Connections []*ListConsumerConnectionsResponseBodyDataConnections `json:"connections,omitempty" xml:"connections,omitempty" type:"Repeated"`
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
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
	// example:
	//
	// 172.17.135.197@17392#1936705963#551717232#9873695589062458
	ClientId *string `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// xx.xx.xx.xx
	EgressIp *string `json:"egressIp,omitempty" xml:"egressIp,omitempty"`
	// example:
	//
	// vos
	Hostname *string `json:"hostname,omitempty" xml:"hostname,omitempty"`
	// example:
	//
	// java
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
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
	// The returned error code.
	//
	// example:
	//
	// MissingInstanceId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data []*ListConsumerGroupSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// InstanceId
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
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5F4D9D5F-625B-59FF-BD4F-DA8284575DB4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
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
	Consistency *bool `json:"consistency,omitempty" xml:"consistency,omitempty"`
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
	// The type of the filter expression. Valid values: SQL, TAG, and UNSPECIFIED.
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

type ListInstancesRequest struct {
	// The filter condition that is used to query instances. If you do not configure this parameter, all instances are queried.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The number of the page to return.
	//
	// Valid values: 1 to 100000000.
	//
	// If the value that you specify for this parameter is less than 1, the system uses 1 as the value. If the value that you specify for this parameter is greater than 100000000, the system uses 100000000 as the value.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// Valid values: 10 to 200.
	//
	// If the value that you specify for this parameter is less than 10, the system uses 10 as the value. If the value that you specify for this parameter is greater than 200, the system uses 200 as the value.
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

func (s *ListInstancesRequest) SetTags(v string) *ListInstancesRequest {
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
	// The returned data.
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
	// The paginated data.
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
	// The product information of instance.
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
	// The time when the instance was started.
	//
	// example:
	//
	// 2022-08-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The instance status.
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
	// Whether to enable tracking capability. Non-serverless instances are enabled by default, and serverless instances are optional for users.
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

type ListRegionsResponseBody struct {
	// example:
	//
	// MissingPageNumber
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListRegionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// ConsumerGroupId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 0B962390-D84B-5D44-8C11-79DF40299D41
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId   *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	RegionName *string `json:"regionName,omitempty" xml:"regionName,omitempty"`
	// example:
	//
	// true
	SupportRocketmqV4 *bool `json:"supportRocketmqV4,omitempty" xml:"supportRocketmqV4,omitempty"`
	// example:
	//
	// true
	SupportRocketmqV5 *bool                              `json:"supportRocketmqV5,omitempty" xml:"supportRocketmqV5,omitempty"`
	Tags              []*ListRegionsResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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
	// example:
	//
	// xx
	TagCode *string `json:"tagCode,omitempty" xml:"tagCode,omitempty"`
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

type ListTopicSubscriptionsResponseBody struct {
	// example:
	//
	// MissingInstanceId
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListTopicSubscriptionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// Topic
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 92A9BE4E-B794-50C8-979C-0456E4D32943
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Consistency *string `json:"consistency,omitempty" xml:"consistency,omitempty"`
	// example:
	//
	// CID-TEST
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// example:
	//
	// *
	FilterExpression *string `json:"filterExpression,omitempty" xml:"filterExpression,omitempty"`
	// example:
	//
	// SQL
	FilterExpressionType *string `json:"filterExpressionType,omitempty" xml:"filterExpressionType,omitempty"`
	// example:
	//
	// BROADCASTING
	MessageModel *string `json:"messageModel,omitempty" xml:"messageModel,omitempty"`
	// example:
	//
	// ONLINE
	SubscriptionStatus *string `json:"subscriptionStatus,omitempty" xml:"subscriptionStatus,omitempty"`
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

type UpdateTopicRequest struct {
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
// Queries the subscriptions of a consumer group.
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
// Queries the subscriptions of a consumer group.
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
// Queries instances.
//
// Description:
//
// > API operations provided by Alibaba Cloud are used to manage and query resources of Alibaba Cloud services. We recommend that you integrate these API operations only in management systems. Do not use these API operations in the core system of messaging services. Otherwise, system risks may occur.
//
// @param request - ListInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
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
// Queries instances.
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
// 查询region列表
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
// 查询region列表
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
// 查询主题订阅关系列表
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
// 查询主题订阅关系列表
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
