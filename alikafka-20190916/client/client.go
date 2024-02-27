// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeResourceGroupRequest struct {
	// The ID of the resource group to which you want to transfer the cloud resource.
	//
	// >  You can use resource groups to manage resources owned by your Alibaba Cloud account. Resource groups simplify the resource and permission management of your Alibaba Cloud account. For more information, see [What is resource management?](~~94475~~)
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID of the resource.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource to which you want to attach a tag. Only the ID of a Message Queue for Apache Kafka instance is supported.
	//
	// For example, if the ID of the instance is alikafka_post-cn-v0h1fgs2xxxx, the resource ID is alikafka_post-cn-v0h1fgs2xxxx.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the new resource group. You can view the available resource groups in the Resource Management console.
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetCode(v int32) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetNewResourceGroupId(v string) *ChangeResourceGroupResponseBody {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetSuccess(v int64) *ChangeResourceGroupResponseBody {
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

type ConvertPostPayOrderRequest struct {
	// The subscription duration. Unit: months. Valid values:
	//
	// *   **1~12**
	// *   **24**
	// *   **36**
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConvertPostPayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s ConvertPostPayOrderRequest) GoString() string {
	return s.String()
}

func (s *ConvertPostPayOrderRequest) SetDuration(v int32) *ConvertPostPayOrderRequest {
	s.Duration = &v
	return s
}

func (s *ConvertPostPayOrderRequest) SetInstanceId(v string) *ConvertPostPayOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertPostPayOrderRequest) SetRegionId(v string) *ConvertPostPayOrderRequest {
	s.RegionId = &v
	return s
}

type ConvertPostPayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConvertPostPayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConvertPostPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertPostPayOrderResponseBody) SetCode(v int32) *ConvertPostPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) SetMessage(v string) *ConvertPostPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) SetOrderId(v string) *ConvertPostPayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) SetRequestId(v string) *ConvertPostPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertPostPayOrderResponseBody) SetSuccess(v bool) *ConvertPostPayOrderResponseBody {
	s.Success = &v
	return s
}

type ConvertPostPayOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConvertPostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConvertPostPayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s ConvertPostPayOrderResponse) GoString() string {
	return s.String()
}

func (s *ConvertPostPayOrderResponse) SetHeaders(v map[string]*string) *ConvertPostPayOrderResponse {
	s.Headers = v
	return s
}

func (s *ConvertPostPayOrderResponse) SetStatusCode(v int32) *ConvertPostPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ConvertPostPayOrderResponse) SetBody(v *ConvertPostPayOrderResponseBody) *ConvertPostPayOrderResponse {
	s.Body = v
	return s
}

type CreateAclRequest struct {
	// The type of operation allowed by the ACL. Valid values:
	//
	// *   **Write**: data writes.
	// *   **Read**: data reads.
	// *   **Describe**: reads of **transaction IDs**.
	// *   **IdempotentWrite**: idempotent data writes to **clusters**.
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The name or ID of the resource.
	//
	// *   The value can be the name of a topic, consumer group, or cluster, or the ID of a transaction.
	// *   You can use an asterisk (\*) to represent the names or IDs of all relevant resources.
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The matching mode. Valid values:
	//
	// *   **LITERAL**: exact match
	// *   **PREFIXED**: prefix match
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// *   **Topic**: topic
	// *   **Group**: consumer group
	// *   **Cluster**: cluster
	// *   **TransactionalId**: transaction
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The username.
	//
	// You can use an asterisk (\*) to represent all usernames.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateAclRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAclRequest) GoString() string {
	return s.String()
}

func (s *CreateAclRequest) SetAclOperationType(v string) *CreateAclRequest {
	s.AclOperationType = &v
	return s
}

func (s *CreateAclRequest) SetAclResourceName(v string) *CreateAclRequest {
	s.AclResourceName = &v
	return s
}

func (s *CreateAclRequest) SetAclResourcePatternType(v string) *CreateAclRequest {
	s.AclResourcePatternType = &v
	return s
}

func (s *CreateAclRequest) SetAclResourceType(v string) *CreateAclRequest {
	s.AclResourceType = &v
	return s
}

func (s *CreateAclRequest) SetInstanceId(v string) *CreateAclRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAclRequest) SetRegionId(v string) *CreateAclRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAclRequest) SetUsername(v string) *CreateAclRequest {
	s.Username = &v
	return s
}

type CreateAclResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclResponseBody) SetCode(v int32) *CreateAclResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAclResponseBody) SetMessage(v string) *CreateAclResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAclResponseBody) SetRequestId(v string) *CreateAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAclResponseBody) SetSuccess(v bool) *CreateAclResponseBody {
	s.Success = &v
	return s
}

type CreateAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAclResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAclResponse) GoString() string {
	return s.String()
}

func (s *CreateAclResponse) SetHeaders(v map[string]*string) *CreateAclResponse {
	s.Headers = v
	return s
}

func (s *CreateAclResponse) SetStatusCode(v int32) *CreateAclResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAclResponse) SetBody(v *CreateAclResponseBody) *CreateAclResponse {
	s.Body = v
	return s
}

type CreateConsumerGroupRequest struct {
	// The name of the consumer group.
	//
	// *   The value can contain only letters, digits, hyphens (-), and underscores (\_), and the value must contain at least one letter or digit.
	// *   The value must be 3 to 128 characters in length. If the value that you specify contains more than 128 characters, the system automatically truncates the value to 128 characters.
	// *   After a consumer group is created, you cannot change the name of the consumer group.
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the consumer group.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The tags.
	Tag []*CreateConsumerGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) SetConsumerId(v string) *CreateConsumerGroupRequest {
	s.ConsumerId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetInstanceId(v string) *CreateConsumerGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRegionId(v string) *CreateConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetRemark(v string) *CreateConsumerGroupRequest {
	s.Remark = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetTag(v []*CreateConsumerGroupRequestTag) *CreateConsumerGroupRequest {
	s.Tag = v
	return s
}

type CreateConsumerGroupRequestTag struct {
	// The tag key.
	//
	// *   You must specify this parameter.
	// *   The tag key can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// *   You can leave this parameter empty.
	// *   The tag value can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateConsumerGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequestTag) SetKey(v string) *CreateConsumerGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateConsumerGroupRequestTag) SetValue(v string) *CreateConsumerGroupRequestTag {
	s.Value = &v
	return s
}

type CreateConsumerGroupResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) SetCode(v int32) *CreateConsumerGroupResponseBody {
	s.Code = &v
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

type CreatePostPayOrderRequest struct {
	// The deployment mode of the instance. Valid values:
	//
	// *   **4**: deploys the instance that allows access from the Internet and a VPC.
	// *   **5**: deploys the instance that allows access only from a VPC.
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size.
	//
	// For more information about the valid values, see [Billing](~~84737~~).
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// *   **0**: ultra disk
	// *   **1**: standard SSD
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic for the instance.
	//
	// *   This parameter is required if the **DeployType** parameter is set to **4**.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic for the instance. We recommend that you do not configure this parameter.
	//
	// *   You must specify at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you specify only the IoMaxSpec parameter.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// *   You must specify at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you specify only the IoMaxSpec parameter.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// If this parameter is left empty, the default resource group is used. You can view the resource group ID on the Resource Group page in the Resource Management console.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The edition of the instance. Valid values:
	//
	// *   **normal**: Standard Edition (High Write)
	// *   **professional**: Professional Edition (High Write)
	// *   **professionalForHighRead**: Professional Edition (High Read)
	//
	// For more information about these instance editions, see [Billing](~~84737~~).
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePostPayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   The default value of the TopicQuota parameter varies based on the value of the IoMaxSpec parameter. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s CreatePostPayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderRequest) SetDeployType(v int32) *CreatePostPayOrderRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetDiskSize(v int32) *CreatePostPayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetDiskType(v string) *CreatePostPayOrderRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetEipMax(v int32) *CreatePostPayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetIoMax(v int32) *CreatePostPayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetIoMaxSpec(v string) *CreatePostPayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetPartitionNum(v int32) *CreatePostPayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetRegionId(v string) *CreatePostPayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetResourceGroupId(v string) *CreatePostPayOrderRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetSpecType(v string) *CreatePostPayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetTag(v []*CreatePostPayOrderRequestTag) *CreatePostPayOrderRequest {
	s.Tag = v
	return s
}

func (s *CreatePostPayOrderRequest) SetTopicQuota(v int32) *CreatePostPayOrderRequest {
	s.TopicQuota = &v
	return s
}

type CreatePostPayOrderRequestTag struct {
	// The key of tag N.
	//
	// *   Valid values of N: 1 to 20.
	// *   If this parameter is left empty, the keys of all tags are matched.
	// *   The tag key must be up to 128 characters in length. It cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// *   Valid values of N: 1 to 20.
	// *   If you do not specify a tag key, you cannot specify a tag value. If this parameter is not configured, all tag values are matched.
	// *   The tag value must be 1 to 128 characters in length. It cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePostPayOrderRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayOrderRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderRequestTag) SetKey(v string) *CreatePostPayOrderRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayOrderRequestTag) SetValue(v string) *CreatePostPayOrderRequestTag {
	s.Value = &v
	return s
}

type CreatePostPayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePostPayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderResponseBody) SetCode(v int32) *CreatePostPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) SetMessage(v string) *CreatePostPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) SetOrderId(v string) *CreatePostPayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) SetRequestId(v string) *CreatePostPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePostPayOrderResponseBody) SetSuccess(v bool) *CreatePostPayOrderResponseBody {
	s.Success = &v
	return s
}

type CreatePostPayOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePostPayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderResponse) SetHeaders(v map[string]*string) *CreatePostPayOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePostPayOrderResponse) SetStatusCode(v int32) *CreatePostPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePostPayOrderResponse) SetBody(v *CreatePostPayOrderResponseBody) *CreatePostPayOrderResponse {
	s.Body = v
	return s
}

type CreatePrePayOrderRequest struct {
	// The deployment mode of the instance. Valid values:
	//
	// *   **4**: deploys the instance that allows access from the Internet and a VPC.
	// *   **5**: deploys the instance that allows access only from a VPC.
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size. Unit: GB.
	//
	// For more information about the valid values, see [Billing](~~84737~~).
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// *   **0**: ultra disk
	// *   **1**: standard SSD
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic for the instance.
	//
	// *   This parameter is required if the **DeployType** parameter is set to **4**.
	// *   For more information about the valid values, see [Pay-as-you-go](~~72142~~).
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic for the instance. We recommend that you do not configure this parameter.
	//
	// *   You must configure at least one of the **IoMax** and **IoMaxSpec** parameters. If both parameters are configured, the value of the **IoMaxSpec** parameter takes effect. We recommend that you configure only the **IoMaxSpec** parameter.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// *   You must configure at least one of the **IoMax** and **IoMaxSpec** parameters. If both parameters are configured, the value of the **IoMaxSpec** parameter takes effect. We recommend that you configure only the **IoMaxSpec** parameter.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// If this parameter is left empty, the default resource group is used. You can view the resource group ID on the Resource Group page in the Resource Management console.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The edition of the instance. Valid values:
	//
	// *   **normal**: Standard Edition (High Write)
	// *   **professional**: Professional Edition (High Write)
	// *   **professionalForHighRead**: Professional Edition (High Read)
	//
	// For more information, see [Billing](~~84737~~).
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePrePayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   The default value of the TopicQuota parameter varies based on the value of the IoMaxSpec parameter. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s CreatePrePayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderRequest) SetDeployType(v int32) *CreatePrePayOrderRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetDiskSize(v int32) *CreatePrePayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetDiskType(v string) *CreatePrePayOrderRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetEipMax(v int32) *CreatePrePayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetIoMax(v int32) *CreatePrePayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetIoMaxSpec(v string) *CreatePrePayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetPartitionNum(v int32) *CreatePrePayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetRegionId(v string) *CreatePrePayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetResourceGroupId(v string) *CreatePrePayOrderRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetSpecType(v string) *CreatePrePayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetTag(v []*CreatePrePayOrderRequestTag) *CreatePrePayOrderRequest {
	s.Tag = v
	return s
}

func (s *CreatePrePayOrderRequest) SetTopicQuota(v int32) *CreatePrePayOrderRequest {
	s.TopicQuota = &v
	return s
}

type CreatePrePayOrderRequestTag struct {
	// The key of tag N.
	//
	// *   Valid values of N: 1 to 20.
	// *   If this parameter is left empty, the keys of all tags are matched.
	// *   The tag key can be up to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// *   Valid values of N: 1 to 20.
	// *   This parameter can be left empty.
	// *   The tag value can be 1 to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrePayOrderRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayOrderRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderRequestTag) SetKey(v string) *CreatePrePayOrderRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrePayOrderRequestTag) SetValue(v string) *CreatePrePayOrderRequestTag {
	s.Value = &v
	return s
}

type CreatePrePayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePrePayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderResponseBody) SetCode(v int32) *CreatePrePayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) SetMessage(v string) *CreatePrePayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) SetOrderId(v string) *CreatePrePayOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) SetRequestId(v string) *CreatePrePayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrePayOrderResponseBody) SetSuccess(v bool) *CreatePrePayOrderResponseBody {
	s.Success = &v
	return s
}

type CreatePrePayOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrePayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrePayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayOrderResponse) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderResponse) SetHeaders(v map[string]*string) *CreatePrePayOrderResponse {
	s.Headers = v
	return s
}

func (s *CreatePrePayOrderResponse) SetStatusCode(v int32) *CreatePrePayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrePayOrderResponse) SetBody(v *CreatePrePayOrderResponseBody) *CreatePrePayOrderResponse {
	s.Body = v
	return s
}

type CreateSaslUserRequest struct {
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password of the SASL user.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SASL mechanism. Valid values:
	//
	// *   **plain**: a simple mechanism that uses usernames and passwords to verify user identities. Message Queue for Apache Kafka provides an optimized PLAIN mechanism that allows you to dynamically create SASL users for an instance without the need to restart the instance.
	// *   **scram**: a mechanism that uses usernames and passwords to verify user identities. This mechanism provides better security protection than the PLAIN mechanism. Message Queue for Apache Kafka uses SCRAM-SHA-256.
	//
	// Default value: **plain**.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the SASL user.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateSaslUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSaslUserRequest) GoString() string {
	return s.String()
}

func (s *CreateSaslUserRequest) SetInstanceId(v string) *CreateSaslUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSaslUserRequest) SetPassword(v string) *CreateSaslUserRequest {
	s.Password = &v
	return s
}

func (s *CreateSaslUserRequest) SetRegionId(v string) *CreateSaslUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSaslUserRequest) SetType(v string) *CreateSaslUserRequest {
	s.Type = &v
	return s
}

func (s *CreateSaslUserRequest) SetUsername(v string) *CreateSaslUserRequest {
	s.Username = &v
	return s
}

type CreateSaslUserResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSaslUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSaslUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSaslUserResponseBody) SetCode(v int32) *CreateSaslUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSaslUserResponseBody) SetMessage(v string) *CreateSaslUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSaslUserResponseBody) SetRequestId(v string) *CreateSaslUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSaslUserResponseBody) SetSuccess(v bool) *CreateSaslUserResponseBody {
	s.Success = &v
	return s
}

type CreateSaslUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSaslUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSaslUserResponse) GoString() string {
	return s.String()
}

func (s *CreateSaslUserResponse) SetHeaders(v map[string]*string) *CreateSaslUserResponse {
	s.Headers = v
	return s
}

func (s *CreateSaslUserResponse) SetStatusCode(v int32) *CreateSaslUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSaslUserResponse) SetBody(v *CreateSaslUserResponseBody) *CreateSaslUserResponse {
	s.Body = v
	return s
}

type CreateTopicRequest struct {
	// The log cleanup policy that is used for the topic. This parameter is available only when LocalTopic is set to true. Valid values:
	//
	// *   false: The topic uses the default log cleanup policy.
	// *   true: The topic uses the log compaction policy.
	CompactTopic *bool `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	// The additional configurations.
	//
	// *   The value of this parameter must be in JSON format.
	// *   The key must be **replications**. The value indicates the number of replicas for the topic. The value must be an integer that ranges from 1 to 3.
	// *   This parameter is available only when **LocalTopic** is set to **true**, or the instance is of the **Open Source Edition (Local Disk)**.****
	//
	// > If you specify this parameter, **ReplicationFactor** does not take effect.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of storage that the topic uses. Valid values:
	//
	// *   false: The topic uses cloud storage.
	// *   true: The topic uses local storage.
	LocalTopic *bool `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	// The minimum number of in-sync replicas (ISRs).
	//
	// *   This parameter is available only when **LocalTopic** is set to **true**, or the instance is of the **Open Source Edition (Local Disk)**.****
	// *   The value of this parameter must be smaller than the value of ReplicationFactor.
	// *   Valid values: 1 to 3.
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" xml:"MinInsyncReplicas,omitempty"`
	// The number of partitions in the topic.
	//
	// *   Valid values: 1 to 360.
	// *   The system recommends the number of partitions based on the specification of the instance. You can view the recommended number in the Message Queue for Apache Kafka console. We recommend that you specify the number that is recommended by the system as the value of this parameter to reduce the risk of data skew.
	PartitionNum *string `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance in which you want to create a topic.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the topic.
	//
	// *   The description can contain only letters, digits, hyphens (-), and underscores (\_).
	// *   The description must be 3 to 64 characters in length.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The number of replicas for the topic.
	//
	// *   This parameter is available only when **LocalTopic** is set to **true**, or the instance is of the **Open Source Edition (Local Disk)**.****
	// *   Valid values: 1 to 3.
	//
	// > If you set this parameter to **1**, data loss may occur. Exercise caution when you configure this parameter.
	ReplicationFactor *int64 `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The tags that you want to add to the topic.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The topic name.
	//
	// *   The name can contain only letters, digits, hyphens (-), and underscores (\_).
	// *   The name must be 3 to 64 characters in length. If the name that you specify contains more than 64 characters, the system automatically truncates the name.
	// *   After a topic is created, you cannot change the name of the topic.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) SetCompactTopic(v bool) *CreateTopicRequest {
	s.CompactTopic = &v
	return s
}

func (s *CreateTopicRequest) SetConfig(v string) *CreateTopicRequest {
	s.Config = &v
	return s
}

func (s *CreateTopicRequest) SetInstanceId(v string) *CreateTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTopicRequest) SetLocalTopic(v bool) *CreateTopicRequest {
	s.LocalTopic = &v
	return s
}

func (s *CreateTopicRequest) SetMinInsyncReplicas(v int64) *CreateTopicRequest {
	s.MinInsyncReplicas = &v
	return s
}

func (s *CreateTopicRequest) SetPartitionNum(v string) *CreateTopicRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreateTopicRequest) SetRegionId(v string) *CreateTopicRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTopicRequest) SetRemark(v string) *CreateTopicRequest {
	s.Remark = &v
	return s
}

func (s *CreateTopicRequest) SetReplicationFactor(v int64) *CreateTopicRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *CreateTopicRequest) SetTag(v []*CreateTopicRequestTag) *CreateTopicRequest {
	s.Tag = v
	return s
}

func (s *CreateTopicRequest) SetTopic(v string) *CreateTopicRequest {
	s.Topic = &v
	return s
}

type CreateTopicRequestTag struct {
	// The tag key.
	//
	// *   If you do not specify this parameter, the keys of all tags are matched.
	// *   The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// *   You can leave this parameter empty.
	// *   The tag value must be 1 to 128 characters in length and cannot contain http:// or https://. The tag value cannot start with aliyun or acs:.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTopicRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTopicRequestTag) SetKey(v string) *CreateTopicRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTopicRequestTag) SetValue(v string) *CreateTopicRequestTag {
	s.Value = &v
	return s
}

type CreateTopicResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the call is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) SetCode(v int32) *CreateTopicResponseBody {
	s.Code = &v
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

type DeleteAclRequest struct {
	// The type of operation allowed by the ACL. Valid values:
	//
	// *   **Write**
	// *   **Read**
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The name of the resource.
	//
	// *   The value can be the name of a topic or consumer group.
	// *   You can use an asterisk (\*) to indicate the names of all topics or consumer groups.
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The mode that is used to match resources. Valid values:
	//
	// *   **LITERAL:** full match
	// *   **PREFIXED**: prefix match
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The type of the resource.
	//
	// *   **Topic**
	// *   **Group**
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DeleteAclRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclRequest) GoString() string {
	return s.String()
}

func (s *DeleteAclRequest) SetAclOperationType(v string) *DeleteAclRequest {
	s.AclOperationType = &v
	return s
}

func (s *DeleteAclRequest) SetAclResourceName(v string) *DeleteAclRequest {
	s.AclResourceName = &v
	return s
}

func (s *DeleteAclRequest) SetAclResourcePatternType(v string) *DeleteAclRequest {
	s.AclResourcePatternType = &v
	return s
}

func (s *DeleteAclRequest) SetAclResourceType(v string) *DeleteAclRequest {
	s.AclResourceType = &v
	return s
}

func (s *DeleteAclRequest) SetInstanceId(v string) *DeleteAclRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAclRequest) SetRegionId(v string) *DeleteAclRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAclRequest) SetUsername(v string) *DeleteAclRequest {
	s.Username = &v
	return s
}

type DeleteAclResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclResponseBody) SetCode(v int32) *DeleteAclResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAclResponseBody) SetMessage(v string) *DeleteAclResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAclResponseBody) SetRequestId(v string) *DeleteAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAclResponseBody) SetSuccess(v bool) *DeleteAclResponseBody {
	s.Success = &v
	return s
}

type DeleteAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAclResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAclResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclResponse) SetHeaders(v map[string]*string) *DeleteAclResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclResponse) SetStatusCode(v int32) *DeleteAclResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAclResponse) SetBody(v *DeleteAclResponseBody) *DeleteAclResponse {
	s.Body = v
	return s
}

type DeleteConsumerGroupRequest struct {
	// The name of the consumer group.
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupRequest) SetConsumerId(v string) *DeleteConsumerGroupRequest {
	s.ConsumerId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetInstanceId(v string) *DeleteConsumerGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetRegionId(v string) *DeleteConsumerGroupRequest {
	s.RegionId = &v
	return s
}

type DeleteConsumerGroupResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) SetCode(v int32) *DeleteConsumerGroupResponseBody {
	s.Code = &v
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

type DeleteInstanceRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetRegionId(v string) *DeleteInstanceRequest {
	s.RegionId = &v
	return s
}

type DeleteInstanceResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetCode(v int32) *DeleteInstanceResponseBody {
	s.Code = &v
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

type DeleteSaslUserRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The SASL mechanism. Valid values:
	//
	// *   **plain**: a simple mechanism that uses usernames and passwords to verify user identities. Message Queue for Apache Kafka provides an optimized PLAIN mechanism that allows you to dynamically create SASL users for an instance without the need to restart the instance.
	// *   **scram**: a mechanism that uses usernames and passwords to verify user identities. This mechanism provides better security protection than the PLAIN mechanism. Message Queue for Apache Kafka uses SCRAM-SHA-256.
	//
	// Default value: **plain**.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the user.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DeleteSaslUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSaslUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteSaslUserRequest) SetInstanceId(v string) *DeleteSaslUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSaslUserRequest) SetRegionId(v string) *DeleteSaslUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSaslUserRequest) SetType(v string) *DeleteSaslUserRequest {
	s.Type = &v
	return s
}

func (s *DeleteSaslUserRequest) SetUsername(v string) *DeleteSaslUserRequest {
	s.Username = &v
	return s
}

type DeleteSaslUserResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSaslUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSaslUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSaslUserResponseBody) SetCode(v int32) *DeleteSaslUserResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteSaslUserResponseBody) SetMessage(v string) *DeleteSaslUserResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSaslUserResponseBody) SetRequestId(v string) *DeleteSaslUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSaslUserResponseBody) SetSuccess(v bool) *DeleteSaslUserResponseBody {
	s.Success = &v
	return s
}

type DeleteSaslUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSaslUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSaslUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteSaslUserResponse) SetHeaders(v map[string]*string) *DeleteSaslUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteSaslUserResponse) SetStatusCode(v int32) *DeleteSaslUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSaslUserResponse) SetBody(v *DeleteSaslUserResponseBody) *DeleteSaslUserResponse {
	s.Body = v
	return s
}

type DeleteTopicRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DeleteTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicRequest) SetInstanceId(v string) *DeleteTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTopicRequest) SetRegionId(v string) *DeleteTopicRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTopicRequest) SetTopic(v string) *DeleteTopicRequest {
	s.Topic = &v
	return s
}

type DeleteTopicResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) SetCode(v int32) *DeleteTopicResponseBody {
	s.Code = &v
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

type DescribeAclsRequest struct {
	// The name or ID of the resource.
	//
	// *   The value can be the name of a topic or a consumer group.
	// *   You can use an asterisk (\*) to represent the names of all topics or consumer groups.
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The match mode. Valid values:
	//
	// *   LITERAL: full-name match
	// *   PREFIXED: prefix match
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// *   **Topic**
	// *   **Group**
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeAclsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclsRequest) SetAclResourceName(v string) *DescribeAclsRequest {
	s.AclResourceName = &v
	return s
}

func (s *DescribeAclsRequest) SetAclResourcePatternType(v string) *DescribeAclsRequest {
	s.AclResourcePatternType = &v
	return s
}

func (s *DescribeAclsRequest) SetAclResourceType(v string) *DescribeAclsRequest {
	s.AclResourceType = &v
	return s
}

func (s *DescribeAclsRequest) SetInstanceId(v string) *DescribeAclsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAclsRequest) SetRegionId(v string) *DescribeAclsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAclsRequest) SetUsername(v string) *DescribeAclsRequest {
	s.Username = &v
	return s
}

type DescribeAclsResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ACLs.
	KafkaAclList *DescribeAclsResponseBodyKafkaAclList `json:"KafkaAclList,omitempty" xml:"KafkaAclList,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAclsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclsResponseBody) SetCode(v int32) *DescribeAclsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAclsResponseBody) SetKafkaAclList(v *DescribeAclsResponseBodyKafkaAclList) *DescribeAclsResponseBody {
	s.KafkaAclList = v
	return s
}

func (s *DescribeAclsResponseBody) SetMessage(v string) *DescribeAclsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAclsResponseBody) SetRequestId(v string) *DescribeAclsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclsResponseBody) SetSuccess(v bool) *DescribeAclsResponseBody {
	s.Success = &v
	return s
}

type DescribeAclsResponseBodyKafkaAclList struct {
	KafkaAclVO []*DescribeAclsResponseBodyKafkaAclListKafkaAclVO `json:"KafkaAclVO,omitempty" xml:"KafkaAclVO,omitempty" type:"Repeated"`
}

func (s DescribeAclsResponseBodyKafkaAclList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclsResponseBodyKafkaAclList) GoString() string {
	return s.String()
}

func (s *DescribeAclsResponseBodyKafkaAclList) SetKafkaAclVO(v []*DescribeAclsResponseBodyKafkaAclListKafkaAclVO) *DescribeAclsResponseBodyKafkaAclList {
	s.KafkaAclVO = v
	return s
}

type DescribeAclsResponseBodyKafkaAclListKafkaAclVO struct {
	// The type of the operation. Valid values:
	//
	// *   **Write**
	// *   **Read**
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The name of the resource.
	//
	// *   The value can be the name of a topic or a consumer group.
	// *   An asterisk (\*) represents the names of all topics or consumer groups.
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The match mode. Valid values:
	//
	// *   **LITERAL**: full-name match
	// *   **PREFIXED**: prefix match
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The type of the resources to which you want to attach tags. Valid values:
	//
	// *   **Topic**
	// *   **Group**
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The host.
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The name of the user.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeAclsResponseBodyKafkaAclListKafkaAclVO) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclsResponseBodyKafkaAclListKafkaAclVO) GoString() string {
	return s.String()
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclOperationType(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclOperationType = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclResourceName(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclResourceName = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclResourcePatternType(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclResourcePatternType = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclResourceType(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclResourceType = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetHost(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.Host = &v
	return s
}

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetUsername(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.Username = &v
	return s
}

type DescribeAclsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclsResponse) SetHeaders(v map[string]*string) *DescribeAclsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclsResponse) SetStatusCode(v int32) *DescribeAclsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclsResponse) SetBody(v *DescribeAclsResponseBody) *DescribeAclsResponse {
	s.Body = v
	return s
}

type DescribeSaslUsersRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSaslUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSaslUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersRequest) SetInstanceId(v string) *DescribeSaslUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSaslUsersRequest) SetRegionId(v string) *DescribeSaslUsersRequest {
	s.RegionId = &v
	return s
}

type DescribeSaslUsersResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SASL users.
	SaslUserList *DescribeSaslUsersResponseBodySaslUserList `json:"SaslUserList,omitempty" xml:"SaslUserList,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSaslUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSaslUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponseBody) SetCode(v int32) *DescribeSaslUsersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSaslUsersResponseBody) SetMessage(v string) *DescribeSaslUsersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSaslUsersResponseBody) SetRequestId(v string) *DescribeSaslUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSaslUsersResponseBody) SetSaslUserList(v *DescribeSaslUsersResponseBodySaslUserList) *DescribeSaslUsersResponseBody {
	s.SaslUserList = v
	return s
}

func (s *DescribeSaslUsersResponseBody) SetSuccess(v bool) *DescribeSaslUsersResponseBody {
	s.Success = &v
	return s
}

type DescribeSaslUsersResponseBodySaslUserList struct {
	SaslUserVO []*DescribeSaslUsersResponseBodySaslUserListSaslUserVO `json:"SaslUserVO,omitempty" xml:"SaslUserVO,omitempty" type:"Repeated"`
}

func (s DescribeSaslUsersResponseBodySaslUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSaslUsersResponseBodySaslUserList) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponseBodySaslUserList) SetSaslUserVO(v []*DescribeSaslUsersResponseBodySaslUserListSaslUserVO) *DescribeSaslUsersResponseBodySaslUserList {
	s.SaslUserVO = v
	return s
}

type DescribeSaslUsersResponseBodySaslUserListSaslUserVO struct {
	// The password that is used to access the Elasticsearch cluster.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The request type. Valid values:
	//
	// *   **plain**: a simple mechanism that uses usernames and passwords to verify user identities. Message Queue for Apache Kafka provides an optimized PLAIN mechanism that allows you to dynamically create SASL users for an instance without the need to restart the instance.
	// *   **scram**: a mechanism that uses usernames and passwords to verify user identities. This mechanism provides better security protection than the PLAIN mechanism. Message Queue for Apache Kafka uses SCRAM-SHA-256.
	//
	// Default value: **plain**.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the user.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeSaslUsersResponseBodySaslUserListSaslUserVO) String() string {
	return tea.Prettify(s)
}

func (s DescribeSaslUsersResponseBodySaslUserListSaslUserVO) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) SetPassword(v string) *DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	s.Password = &v
	return s
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) SetType(v string) *DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	s.Type = &v
	return s
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) SetUsername(v string) *DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	s.Username = &v
	return s
}

type DescribeSaslUsersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSaslUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSaslUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSaslUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponse) SetHeaders(v map[string]*string) *DescribeSaslUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSaslUsersResponse) SetStatusCode(v int32) *DescribeSaslUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSaslUsersResponse) SetBody(v *DescribeSaslUsersResponseBody) *DescribeSaslUsersResponse {
	s.Body = v
	return s
}

type GetAllInstanceIdListRequest struct {
	// The region ID of the instance. This parameter is reserved.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAllInstanceIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllInstanceIdListRequest) GoString() string {
	return s.String()
}

func (s *GetAllInstanceIdListRequest) SetRegionId(v string) *GetAllInstanceIdListRequest {
	s.RegionId = &v
	return s
}

type GetAllInstanceIdListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of instances that are managed by the Alibaba Cloud account in all the regions.
	InstanceIds map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllInstanceIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllInstanceIdListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllInstanceIdListResponseBody) SetCode(v int32) *GetAllInstanceIdListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllInstanceIdListResponseBody) SetInstanceIds(v map[string]interface{}) *GetAllInstanceIdListResponseBody {
	s.InstanceIds = v
	return s
}

func (s *GetAllInstanceIdListResponseBody) SetMessage(v string) *GetAllInstanceIdListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllInstanceIdListResponseBody) SetRequestId(v string) *GetAllInstanceIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllInstanceIdListResponseBody) SetSuccess(v bool) *GetAllInstanceIdListResponseBody {
	s.Success = &v
	return s
}

type GetAllInstanceIdListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllInstanceIdListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllInstanceIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllInstanceIdListResponse) GoString() string {
	return s.String()
}

func (s *GetAllInstanceIdListResponse) SetHeaders(v map[string]*string) *GetAllInstanceIdListResponse {
	s.Headers = v
	return s
}

func (s *GetAllInstanceIdListResponse) SetStatusCode(v int32) *GetAllInstanceIdListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllInstanceIdListResponse) SetBody(v *GetAllInstanceIdListResponseBody) *GetAllInstanceIdListResponse {
	s.Body = v
	return s
}

type GetAllowedIpListRequest struct {
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAllowedIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllowedIpListRequest) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListRequest) SetInstanceId(v string) *GetAllowedIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAllowedIpListRequest) SetRegionId(v string) *GetAllowedIpListRequest {
	s.RegionId = &v
	return s
}

type GetAllowedIpListResponseBody struct {
	// The IP address whitelist.
	AllowedList *GetAllowedIpListResponseBodyAllowedList `json:"AllowedList,omitempty" xml:"AllowedList,omitempty" type:"Struct"`
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllowedIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllowedIpListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponseBody) SetAllowedList(v *GetAllowedIpListResponseBodyAllowedList) *GetAllowedIpListResponseBody {
	s.AllowedList = v
	return s
}

func (s *GetAllowedIpListResponseBody) SetCode(v int32) *GetAllowedIpListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllowedIpListResponseBody) SetMessage(v string) *GetAllowedIpListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllowedIpListResponseBody) SetRequestId(v string) *GetAllowedIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllowedIpListResponseBody) SetSuccess(v bool) *GetAllowedIpListResponseBody {
	s.Success = &v
	return s
}

type GetAllowedIpListResponseBodyAllowedList struct {
	// The deployment mode of the instance. Valid values:
	//
	// *   **4**: allows access from the Internet and a virtual private cloud (VPC).
	// *   **5**: allows access from a VPC.
	//
	// >  Only integrators need to concern themselves with the value of this parameter.
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The whitelist for access from the Internet.
	InternetList []*GetAllowedIpListResponseBodyAllowedListInternetList `json:"InternetList,omitempty" xml:"InternetList,omitempty" type:"Repeated"`
	// The whitelist for access from a virtual private cloud (VPC).
	VpcList []*GetAllowedIpListResponseBodyAllowedListVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
}

func (s GetAllowedIpListResponseBodyAllowedList) String() string {
	return tea.Prettify(s)
}

func (s GetAllowedIpListResponseBodyAllowedList) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponseBodyAllowedList) SetDeployType(v int32) *GetAllowedIpListResponseBodyAllowedList {
	s.DeployType = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedList) SetInternetList(v []*GetAllowedIpListResponseBodyAllowedListInternetList) *GetAllowedIpListResponseBodyAllowedList {
	s.InternetList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedList) SetVpcList(v []*GetAllowedIpListResponseBodyAllowedListVpcList) *GetAllowedIpListResponseBodyAllowedList {
	s.VpcList = v
	return s
}

type GetAllowedIpListResponseBodyAllowedListInternetList struct {
	// The group to which the IP address whitelist belongs.
	AllowedIpGroup map[string]*string `json:"AllowedIpGroup,omitempty" xml:"AllowedIpGroup,omitempty"`
	// The information about the IP address whitelist.
	AllowedIpList []*string `json:"AllowedIpList,omitempty" xml:"AllowedIpList,omitempty" type:"Repeated"`
	// The port range. Valid value:
	//
	// **9093/9093**.
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
}

func (s GetAllowedIpListResponseBodyAllowedListInternetList) String() string {
	return tea.Prettify(s)
}

func (s GetAllowedIpListResponseBodyAllowedListInternetList) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetAllowedIpGroup(v map[string]*string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.AllowedIpGroup = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetAllowedIpList(v []*string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.AllowedIpList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetPortRange(v string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.PortRange = &v
	return s
}

type GetAllowedIpListResponseBodyAllowedListVpcList struct {
	// The group to which the IP address whitelist belongs.
	AllowedIpGroup map[string]*string `json:"AllowedIpGroup,omitempty" xml:"AllowedIpGroup,omitempty"`
	// The information about the IP address whitelist.
	AllowedIpList []*string `json:"AllowedIpList,omitempty" xml:"AllowedIpList,omitempty" type:"Repeated"`
	// The port range. Valid value:
	//
	// **9092/9092**.
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
}

func (s GetAllowedIpListResponseBodyAllowedListVpcList) String() string {
	return tea.Prettify(s)
}

func (s GetAllowedIpListResponseBodyAllowedListVpcList) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetAllowedIpGroup(v map[string]*string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.AllowedIpGroup = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetAllowedIpList(v []*string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.AllowedIpList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetPortRange(v string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.PortRange = &v
	return s
}

type GetAllowedIpListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAllowedIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAllowedIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllowedIpListResponse) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponse) SetHeaders(v map[string]*string) *GetAllowedIpListResponse {
	s.Headers = v
	return s
}

func (s *GetAllowedIpListResponse) SetStatusCode(v int32) *GetAllowedIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllowedIpListResponse) SetBody(v *GetAllowedIpListResponseBody) *GetAllowedIpListResponse {
	s.Body = v
	return s
}

type GetConsumerListRequest struct {
	// The name of the consumer group. If you do not configure this parameter, all consumer groups are queried.
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance to which the consumer group belongs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsumerListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerListRequest) SetConsumerId(v string) *GetConsumerListRequest {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerListRequest) SetInstanceId(v string) *GetConsumerListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerListRequest) SetRegionId(v string) *GetConsumerListRequest {
	s.RegionId = &v
	return s
}

type GetConsumerListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The consumer groups.
	ConsumerList *GetConsumerListResponseBodyConsumerList `json:"ConsumerList,omitempty" xml:"ConsumerList,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConsumerListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBody) SetCode(v int32) *GetConsumerListResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerListResponseBody) SetConsumerList(v *GetConsumerListResponseBodyConsumerList) *GetConsumerListResponseBody {
	s.ConsumerList = v
	return s
}

func (s *GetConsumerListResponseBody) SetMessage(v string) *GetConsumerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerListResponseBody) SetRequestId(v string) *GetConsumerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerListResponseBody) SetSuccess(v bool) *GetConsumerListResponseBody {
	s.Success = &v
	return s
}

type GetConsumerListResponseBodyConsumerList struct {
	ConsumerVO []*GetConsumerListResponseBodyConsumerListConsumerVO `json:"ConsumerVO,omitempty" xml:"ConsumerVO,omitempty" type:"Repeated"`
}

func (s GetConsumerListResponseBodyConsumerList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerList) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerList) SetConsumerVO(v []*GetConsumerListResponseBodyConsumerListConsumerVO) *GetConsumerListResponseBodyConsumerList {
	s.ConsumerVO = v
	return s
}

type GetConsumerListResponseBodyConsumerListConsumerVO struct {
	// The consumer group that is automatically created by the system.
	AutomaticallyCreatedGroup *bool `json:"AutomaticallyCreatedGroup,omitempty" xml:"AutomaticallyCreatedGroup,omitempty"`
	// The ID of the consumer group.
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the consumer group.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The tags.
	Tags *GetConsumerListResponseBodyConsumerListConsumerVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s GetConsumerListResponseBodyConsumerListConsumerVO) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerListConsumerVO) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetAutomaticallyCreatedGroup(v bool) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.AutomaticallyCreatedGroup = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetConsumerId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetInstanceId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetRegionId(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.RegionId = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetRemark(v string) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.Remark = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetTags(v *GetConsumerListResponseBodyConsumerListConsumerVOTags) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.Tags = v
	return s
}

type GetConsumerListResponseBodyConsumerListConsumerVOTags struct {
	TagVO []*GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO `json:"TagVO,omitempty" xml:"TagVO,omitempty" type:"Repeated"`
}

func (s GetConsumerListResponseBodyConsumerListConsumerVOTags) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerListConsumerVOTags) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTags) SetTagVO(v []*GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) *GetConsumerListResponseBodyConsumerListConsumerVOTags {
	s.TagVO = v
	return s
}

type GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) SetKey(v string) *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO {
	s.Key = &v
	return s
}

func (s *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO) SetValue(v string) *GetConsumerListResponseBodyConsumerListConsumerVOTagsTagVO {
	s.Value = &v
	return s
}

type GetConsumerListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponse) SetHeaders(v map[string]*string) *GetConsumerListResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerListResponse) SetStatusCode(v int32) *GetConsumerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerListResponse) SetBody(v *GetConsumerListResponseBody) *GetConsumerListResponse {
	s.Body = v
	return s
}

type GetConsumerProgressRequest struct {
	// The name of the consumer group.
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetConsumerProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressRequest) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressRequest) SetConsumerId(v string) *GetConsumerProgressRequest {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerProgressRequest) SetInstanceId(v string) *GetConsumerProgressRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerProgressRequest) SetRegionId(v string) *GetConsumerProgressRequest {
	s.RegionId = &v
	return s
}

type GetConsumerProgressResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The consumption status of the consumer group.
	ConsumerProgress *GetConsumerProgressResponseBodyConsumerProgress `json:"ConsumerProgress,omitempty" xml:"ConsumerProgress,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetConsumerProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBody) SetCode(v int32) *GetConsumerProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetConsumerProgress(v *GetConsumerProgressResponseBodyConsumerProgress) *GetConsumerProgressResponseBody {
	s.ConsumerProgress = v
	return s
}

func (s *GetConsumerProgressResponseBody) SetMessage(v string) *GetConsumerProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetRequestId(v string) *GetConsumerProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerProgressResponseBody) SetSuccess(v bool) *GetConsumerProgressResponseBody {
	s.Success = &v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgress struct {
	// The time when the last message consumed by the consumer group was generated.
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The consumption progress of each topic to which the consumer group is subscribed.
	TopicList *GetConsumerProgressResponseBodyConsumerProgressTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
	// The number of messages that were not consumed in all topics. This is also known as the number of accumulated messages in all topics.
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgress) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgress) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgress {
	s.LastTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetTopicList(v *GetConsumerProgressResponseBodyConsumerProgressTopicList) *GetConsumerProgressResponseBodyConsumerProgress {
	s.TopicList = v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetTotalDiff(v int64) *GetConsumerProgressResponseBodyConsumerProgress {
	s.TotalDiff = &v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressTopicList struct {
	TopicList []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Repeated"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicList) SetTopicList(v []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) *GetConsumerProgressResponseBodyConsumerProgressTopicList {
	s.TopicList = v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList struct {
	// The time when the last consumed message in the topic was generated.
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The information about offsets in the topic.
	OffsetList *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList `json:"OffsetList,omitempty" xml:"OffsetList,omitempty" type:"Struct"`
	// The name of the topic.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The number of messages that were not consumed in the topic. This is also known as the number of accumulated messages in the topic.
	TotalDiff *int64 `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.LastTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetOffsetList(v *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.OffsetList = v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetTopic(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.Topic = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList) SetTotalDiff(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicList {
	s.TotalDiff = &v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList struct {
	OffsetList []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList `json:"OffsetList,omitempty" xml:"OffsetList,omitempty" type:"Repeated"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList) SetOffsetList(v []*GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList {
	s.OffsetList = v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList struct {
	// The latest offset in the partition of the topic.
	BrokerOffset *int64 `json:"BrokerOffset,omitempty" xml:"BrokerOffset,omitempty"`
	// The consumer offset in the partition of the topic.
	ConsumerOffset *int64 `json:"ConsumerOffset,omitempty" xml:"ConsumerOffset,omitempty"`
	// The time when the last consumed message in the partition was generated.
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The ID of the partition.
	Partition *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetBrokerOffset(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.BrokerOffset = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetConsumerOffset(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.ConsumerOffset = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetLastTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.LastTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetPartition(v int32) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.Partition = &v
	return s
}

type GetConsumerProgressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponse) SetHeaders(v map[string]*string) *GetConsumerProgressResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerProgressResponse) SetStatusCode(v int32) *GetConsumerProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerProgressResponse) SetBody(v *GetConsumerProgressResponseBody) *GetConsumerProgressResponse {
	s.Body = v
	return s
}

type GetInstanceListRequest struct {
	// The IDs of instances.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The ID of the order. You can obtain the order ID on the [Orders](https://usercenter2-intl.aliyun.com/order/list?pageIndex=1\&pageSize=20\&spm=5176.12818093.top-nav.ditem-ord.36f016d0OQFmJa) page in Alibaba Cloud User Center.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the region where the instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can obtain this ID on the Resource Group page in the Resource Management console.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*GetInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceListRequest) SetInstanceId(v []*string) *GetInstanceListRequest {
	s.InstanceId = v
	return s
}

func (s *GetInstanceListRequest) SetOrderId(v string) *GetInstanceListRequest {
	s.OrderId = &v
	return s
}

func (s *GetInstanceListRequest) SetRegionId(v string) *GetInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceListRequest) SetResourceGroupId(v string) *GetInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceListRequest) SetTag(v []*GetInstanceListRequestTag) *GetInstanceListRequest {
	s.Tag = v
	return s
}

type GetInstanceListRequestTag struct {
	// The key of the resource tag.
	//
	// *   If this parameter is left empty, all tag keys are matched.
	// *   The tag key can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the resource tag.
	//
	// *   This parameter must be left empty if the Key parameter is left empty. If this parameter is left empty, the values of all tags are matched.
	// *   The tag value can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *GetInstanceListRequestTag) SetKey(v string) *GetInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *GetInstanceListRequestTag) SetValue(v string) *GetInstanceListRequestTag {
	s.Value = &v
	return s
}

type GetInstanceListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the call is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the instance.
	InstanceList *GetInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the region.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBody) SetCode(v int32) *GetInstanceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceListResponseBody) SetInstanceList(v *GetInstanceListResponseBodyInstanceList) *GetInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *GetInstanceListResponseBody) SetMessage(v string) *GetInstanceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceListResponseBody) SetRequestId(v string) *GetInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceListResponseBody) SetSuccess(v bool) *GetInstanceListResponseBody {
	s.Success = &v
	return s
}

type GetInstanceListResponseBodyInstanceList struct {
	InstanceVO []*GetInstanceListResponseBodyInstanceListInstanceVO `json:"InstanceVO,omitempty" xml:"InstanceVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceList) SetInstanceVO(v []*GetInstanceListResponseBodyInstanceListInstanceVO) *GetInstanceListResponseBodyInstanceList {
	s.InstanceVO = v
	return s
}

type GetInstanceListResponseBodyInstanceListInstanceVO struct {
	// The configurations of the deployed ApsaraMQ for Kafka instance.
	AllConfig *string `json:"AllConfig,omitempty" xml:"AllConfig,omitempty"`
	// The time when the instance was created. Unit: milliseconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the network in which the instance is deployed. Valid values:
	//
	// *   **4**: the Internet and VPCs
	// *   **5**: VPCs
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size of the instance. Unit: GB
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// *   **0**: ultra disk
	// *   **1**: standard SSD
	DiskType *int32 `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The default endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// *   Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	// *   Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	DomainEndpoint *string `json:"DomainEndpoint,omitempty" xml:"DomainEndpoint,omitempty"`
	// The peak Internet traffic allowed for the instance.
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The default endpoint of the instance in IP address mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// *   Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	// *   Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The time when the instance expires. Unit: milliseconds.
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The peak traffic allowed for the instance.
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification.
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The ID of the key that is used for disk encryption in the region where the instance resides.
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The retention period of messages on the instance. Unit: hours.
	MsgRetain *int32 `json:"MsgRetain,omitempty" xml:"MsgRetain,omitempty"`
	// The instance name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// *   **0**: the subscription billing method
	// *   **1**: the pay-as-you-go billing method
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The ID of the region where the instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Simple Authentication and Security Layer (SASL) endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// *   Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	// *   Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	SaslDomainEndpoint *string `json:"SaslDomainEndpoint,omitempty" xml:"SaslDomainEndpoint,omitempty"`
	// The security group to which the instance belongs.
	//
	// *   If the instance is deployed by using the ApsaraMQ for Kafka console or calling the [StartInstance](~~157786~~) operation without a security group configured, no value is returned.
	// *   If the instance is deployed by calling the [StartInstance](~~157786~~) operation with a security group configured, the return value is the configured security group.
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The status of the instance. Valid values:
	//
	// *   **0**: pending
	// *   **1**: preparing hardware resources
	// *   **2**: initializing
	// *   **3**: starting
	// *   **5**: running
	// *   **6**: migrating
	// *   **7**: ready for upgrade
	// *   **8**: upgrading
	// *   **9**: ready for changes
	// *   **10**: released
	// *   **11**: changing
	// *   **15**: expired
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The instance edition. Valid values:
	//
	// *   **professional**: Professional Edition (High Write)
	// *   **professionalForHighRead**: Professional Edition (High Read)
	// *   **normal**: Standard Edition
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The SSL endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// *   Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	// *   Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	SslDomainEndpoint *string `json:"SslDomainEndpoint,omitempty" xml:"SslDomainEndpoint,omitempty"`
	// The Secure Sockets Layer (SSL) endpoint of the instance in IP address mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// *   Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	// *   Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	SslEndPoint *string `json:"SslEndPoint,omitempty" xml:"SslEndPoint,omitempty"`
	// The zone ID.
	StandardZoneId *string `json:"StandardZoneId,omitempty" xml:"StandardZoneId,omitempty"`
	// The tags.
	Tags *GetInstanceListResponseBodyInstanceListInstanceVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The maximum number of topics that can be created on the instance.
	TopicNumLimit *int32 `json:"TopicNumLimit,omitempty" xml:"TopicNumLimit,omitempty"`
	// The upgrade information of the instance.
	UpgradeServiceDetailInfo *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo `json:"UpgradeServiceDetailInfo,omitempty" xml:"UpgradeServiceDetailInfo,omitempty" type:"Struct"`
	// The number of used consumer groups.
	UsedGroupCount *int32 `json:"UsedGroupCount,omitempty" xml:"UsedGroupCount,omitempty"`
	// The number of used partitions.
	UsedPartitionCount *int32 `json:"UsedPartitionCount,omitempty" xml:"UsedPartitionCount,omitempty"`
	// The number of used topics.
	UsedTopicCount *int32 `json:"UsedTopicCount,omitempty" xml:"UsedTopicCount,omitempty"`
	// The vSwitch ID of the instance.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the instance is deployed.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVO) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAllConfig(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AllConfig = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetCreateTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDeployType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DeployType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDiskSize(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DiskSize = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDiskType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DiskType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetEipMax(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.EipMax = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.EndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetExpiredTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ExpiredTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetInstanceId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMax(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMax = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxSpec(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxSpec = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetKmsKeyId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.KmsKeyId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetMsgRetain(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.MsgRetain = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetName(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Name = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetPaidType(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.PaidType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetRegionId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.RegionId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetResourceGroupId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSaslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SaslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSecurityGroup(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SecurityGroup = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetServiceStatus(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ServiceStatus = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSpecType(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SpecType = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetStandardZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.StandardZoneId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetTags(v *GetInstanceListResponseBodyInstanceListInstanceVOTags) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Tags = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetTopicNumLimit(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.TopicNumLimit = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUpgradeServiceDetailInfo(v *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UpgradeServiceDetailInfo = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedGroupCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedGroupCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedPartitionCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedPartitionCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetUsedTopicCount(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.UsedTopicCount = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVSwitchId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VSwitchId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ZoneId = &v
	return s
}

type GetInstanceListResponseBodyInstanceListInstanceVOTags struct {
	TagVO []*GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO `json:"TagVO,omitempty" xml:"TagVO,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTags) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTags) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTags) SetTagVO(v []*GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) *GetInstanceListResponseBodyInstanceListInstanceVOTags {
	s.TagVO = v
	return s
}

type GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) SetKey(v string) *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO {
	s.Key = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO) SetValue(v string) *GetInstanceListResponseBodyInstanceListInstanceVOTagsTagVO {
	s.Value = &v
	return s
}

type GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo struct {
	// The open source Apache Kafka version that corresponds to the instance.
	Current2OpenSourceVersion *string `json:"Current2OpenSourceVersion,omitempty" xml:"Current2OpenSourceVersion,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo) SetCurrent2OpenSourceVersion(v string) *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo {
	s.Current2OpenSourceVersion = &v
	return s
}

type GetInstanceListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponse) SetHeaders(v map[string]*string) *GetInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceListResponse) SetStatusCode(v int32) *GetInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceListResponse) SetBody(v *GetInstanceListResponseBody) *GetInstanceListResponse {
	s.Body = v
	return s
}

type GetQuotaTipRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetQuotaTipRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaTipRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaTipRequest) SetInstanceId(v string) *GetQuotaTipRequest {
	s.InstanceId = &v
	return s
}

func (s *GetQuotaTipRequest) SetRegionId(v string) *GetQuotaTipRequest {
	s.RegionId = &v
	return s
}

type GetQuotaTipResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional message. This message is typically used to describe API call failures for troubleshooting.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The quota.
	QuotaData *GetQuotaTipResponseBodyQuotaData `json:"QuotaData,omitempty" xml:"QuotaData,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQuotaTipResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaTipResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaTipResponseBody) SetCode(v int32) *GetQuotaTipResponseBody {
	s.Code = &v
	return s
}

func (s *GetQuotaTipResponseBody) SetMessage(v string) *GetQuotaTipResponseBody {
	s.Message = &v
	return s
}

func (s *GetQuotaTipResponseBody) SetQuotaData(v *GetQuotaTipResponseBodyQuotaData) *GetQuotaTipResponseBody {
	s.QuotaData = v
	return s
}

func (s *GetQuotaTipResponseBody) SetRequestId(v string) *GetQuotaTipResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaTipResponseBody) SetSuccess(v bool) *GetQuotaTipResponseBody {
	s.Success = &v
	return s
}

type GetQuotaTipResponseBodyQuotaData struct {
	// The number of available groups.
	GroupLeft *int32 `json:"GroupLeft,omitempty" xml:"GroupLeft,omitempty"`
	// The number of used groups.
	GroupUsed *int32 `json:"GroupUsed,omitempty" xml:"GroupUsed,omitempty"`
	// The method that you use to purchase partitions. Valid values:
	//
	// *   0: indicates that the instance is purchased based on topics.
	// *   1: indicates that the instance is purchased based on partitions.
	IsPartitionBuy *int32 `json:"IsPartitionBuy,omitempty" xml:"IsPartitionBuy,omitempty"`
	// The number of available partitions.
	PartitionLeft *int32 `json:"PartitionLeft,omitempty" xml:"PartitionLeft,omitempty"`
	// The number of purchased partitions.
	PartitionNumOfBuy *int32 `json:"PartitionNumOfBuy,omitempty" xml:"PartitionNumOfBuy,omitempty"`
	// The quota of partitions.
	PartitionQuota *int32 `json:"PartitionQuota,omitempty" xml:"PartitionQuota,omitempty"`
	// The number of used partitions.
	PartitionUsed *int32 `json:"PartitionUsed,omitempty" xml:"PartitionUsed,omitempty"`
	// The number of available topics.
	TopicLeft *int32 `json:"TopicLeft,omitempty" xml:"TopicLeft,omitempty"`
	// The number of purchased topics.
	TopicNumOfBuy *int32 `json:"TopicNumOfBuy,omitempty" xml:"TopicNumOfBuy,omitempty"`
	// The quota of topics.
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
	// The number of used topics.
	TopicUsed *int32 `json:"TopicUsed,omitempty" xml:"TopicUsed,omitempty"`
}

func (s GetQuotaTipResponseBodyQuotaData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaTipResponseBodyQuotaData) GoString() string {
	return s.String()
}

func (s *GetQuotaTipResponseBodyQuotaData) SetGroupLeft(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.GroupLeft = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetGroupUsed(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.GroupUsed = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetIsPartitionBuy(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.IsPartitionBuy = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetPartitionLeft(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.PartitionLeft = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetPartitionNumOfBuy(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.PartitionNumOfBuy = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetPartitionQuota(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.PartitionQuota = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetPartitionUsed(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.PartitionUsed = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetTopicLeft(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.TopicLeft = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetTopicNumOfBuy(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.TopicNumOfBuy = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetTopicQuota(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.TopicQuota = &v
	return s
}

func (s *GetQuotaTipResponseBodyQuotaData) SetTopicUsed(v int32) *GetQuotaTipResponseBodyQuotaData {
	s.TopicUsed = &v
	return s
}

type GetQuotaTipResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQuotaTipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQuotaTipResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaTipResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaTipResponse) SetHeaders(v map[string]*string) *GetQuotaTipResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaTipResponse) SetStatusCode(v int32) *GetQuotaTipResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaTipResponse) SetBody(v *GetQuotaTipResponseBody) *GetQuotaTipResponse {
	s.Body = v
	return s
}

type GetTopicListRequest struct {
	// The page number of the page to return. Default value: 1.
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Default value: 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance to which the topics that you want to query belong.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic that you want to query.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListRequest) GoString() string {
	return s.String()
}

func (s *GetTopicListRequest) SetCurrentPage(v string) *GetTopicListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetTopicListRequest) SetInstanceId(v string) *GetTopicListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopicListRequest) SetPageSize(v string) *GetTopicListRequest {
	s.PageSize = &v
	return s
}

func (s *GetTopicListRequest) SetRegionId(v string) *GetTopicListRequest {
	s.RegionId = &v
	return s
}

func (s *GetTopicListRequest) SetTopic(v string) *GetTopicListRequest {
	s.Topic = &v
	return s
}

type GetTopicListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number of the returned page.
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of entries returned on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The information about the topic.
	TopicList *GetTopicListResponseBodyTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
	// The number of topics.
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetTopicListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBody) SetCode(v int32) *GetTopicListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicListResponseBody) SetCurrentPage(v int32) *GetTopicListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetTopicListResponseBody) SetMessage(v string) *GetTopicListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicListResponseBody) SetPageSize(v int32) *GetTopicListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTopicListResponseBody) SetRequestId(v string) *GetTopicListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicListResponseBody) SetSuccess(v bool) *GetTopicListResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicListResponseBody) SetTopicList(v *GetTopicListResponseBodyTopicList) *GetTopicListResponseBody {
	s.TopicList = v
	return s
}

func (s *GetTopicListResponseBody) SetTotal(v int32) *GetTopicListResponseBody {
	s.Total = &v
	return s
}

type GetTopicListResponseBodyTopicList struct {
	TopicVO []*GetTopicListResponseBodyTopicListTopicVO `json:"TopicVO,omitempty" xml:"TopicVO,omitempty" type:"Repeated"`
}

func (s GetTopicListResponseBodyTopicList) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBodyTopicList) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicList) SetTopicVO(v []*GetTopicListResponseBodyTopicListTopicVO) *GetTopicListResponseBodyTopicList {
	s.TopicVO = v
	return s
}

type GetTopicListResponseBodyTopicListTopicVO struct {
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// The log cleanup policy that is used for the topic. This parameter is returned when the **LocalTopic** parameter is set to **true**. Valid values:
	//
	// *   false: The topic uses the default log cleanup policy.
	// *   true: The topic uses the log compaction policy.
	CompactTopic *bool `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	// The timestamp that indicates when the topic was created. Unit: milliseconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of storage used by the topic. Valid values:
	//
	// *   false: The topic uses cloud storage.
	// *   true: The topic uses local storage.
	LocalTopic *bool `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	// The number of partitions in the topic.
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance to which the topics that you want to query belong.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the topic. Valid values:
	//
	// *   The description contains only letters, digits, hyphens (-), and underscores (\_).
	// *   The description is 3 to 64 characters in length.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the topic. Valid values:
	//
	// **0:** indicates that the topic is running.
	//
	// If the topic is deleted, this parameter is not returned.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the topic. Valid values:
	//
	// **Running**
	//
	// If the topic is deleted, this parameter is not returned.
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The tags.
	Tags *GetTopicListResponseBodyTopicListTopicVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The name of the topic. Valid values:
	//
	// *   The name contains only letters, digits, hyphens (-), and underscores (\_).
	// *   The name is 3 to 64 characters in length. If the name that you specified contains more than 64 characters, the returned name is automatically truncated.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicListResponseBodyTopicListTopicVO) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBodyTopicListTopicVO) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetAutoCreate(v bool) *GetTopicListResponseBodyTopicListTopicVO {
	s.AutoCreate = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetCompactTopic(v bool) *GetTopicListResponseBodyTopicListTopicVO {
	s.CompactTopic = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetCreateTime(v int64) *GetTopicListResponseBodyTopicListTopicVO {
	s.CreateTime = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetInstanceId(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.InstanceId = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetLocalTopic(v bool) *GetTopicListResponseBodyTopicListTopicVO {
	s.LocalTopic = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetPartitionNum(v int32) *GetTopicListResponseBodyTopicListTopicVO {
	s.PartitionNum = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetRegionId(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.RegionId = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetRemark(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.Remark = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetStatus(v int32) *GetTopicListResponseBodyTopicListTopicVO {
	s.Status = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetStatusName(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.StatusName = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetTags(v *GetTopicListResponseBodyTopicListTopicVOTags) *GetTopicListResponseBodyTopicListTopicVO {
	s.Tags = v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVO) SetTopic(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.Topic = &v
	return s
}

type GetTopicListResponseBodyTopicListTopicVOTags struct {
	TagVO []*GetTopicListResponseBodyTopicListTopicVOTagsTagVO `json:"TagVO,omitempty" xml:"TagVO,omitempty" type:"Repeated"`
}

func (s GetTopicListResponseBodyTopicListTopicVOTags) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBodyTopicListTopicVOTags) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicListTopicVOTags) SetTagVO(v []*GetTopicListResponseBodyTopicListTopicVOTagsTagVO) *GetTopicListResponseBodyTopicListTopicVOTags {
	s.TagVO = v
	return s
}

type GetTopicListResponseBodyTopicListTopicVOTagsTagVO struct {
	// The tag key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTopicListResponseBodyTopicListTopicVOTagsTagVO) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBodyTopicListTopicVOTagsTagVO) GoString() string {
	return s.String()
}

func (s *GetTopicListResponseBodyTopicListTopicVOTagsTagVO) SetKey(v string) *GetTopicListResponseBodyTopicListTopicVOTagsTagVO {
	s.Key = &v
	return s
}

func (s *GetTopicListResponseBodyTopicListTopicVOTagsTagVO) SetValue(v string) *GetTopicListResponseBodyTopicListTopicVOTagsTagVO {
	s.Value = &v
	return s
}

type GetTopicListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponse) GoString() string {
	return s.String()
}

func (s *GetTopicListResponse) SetHeaders(v map[string]*string) *GetTopicListResponse {
	s.Headers = v
	return s
}

func (s *GetTopicListResponse) SetStatusCode(v int32) *GetTopicListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicListResponse) SetBody(v *GetTopicListResponseBody) *GetTopicListResponse {
	s.Body = v
	return s
}

type GetTopicStatusRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTopicStatusRequest) SetInstanceId(v string) *GetTopicStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopicStatusRequest) SetRegionId(v string) *GetTopicStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetTopicStatusRequest) SetTopic(v string) *GetTopicStatusRequest {
	s.Topic = &v
	return s
}

type GetTopicStatusResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The status information about messages in the topic.
	TopicStatus *GetTopicStatusResponseBodyTopicStatus `json:"TopicStatus,omitempty" xml:"TopicStatus,omitempty" type:"Struct"`
}

func (s GetTopicStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBody) SetCode(v int32) *GetTopicStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetMessage(v string) *GetTopicStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetRequestId(v string) *GetTopicStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetSuccess(v bool) *GetTopicStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicStatusResponseBody) SetTopicStatus(v *GetTopicStatusResponseBodyTopicStatus) *GetTopicStatusResponseBody {
	s.TopicStatus = v
	return s
}

type GetTopicStatusResponseBodyTopicStatus struct {
	// The time when the last consumed message was generated.
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// The information about offsets in the topic.
	OffsetTable *GetTopicStatusResponseBodyTopicStatusOffsetTable `json:"OffsetTable,omitempty" xml:"OffsetTable,omitempty" type:"Struct"`
	// The number of messages in the topic.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTopicStatusResponseBodyTopicStatus) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatus) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetLastTimeStamp(v int64) *GetTopicStatusResponseBodyTopicStatus {
	s.LastTimeStamp = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetOffsetTable(v *GetTopicStatusResponseBodyTopicStatusOffsetTable) *GetTopicStatusResponseBodyTopicStatus {
	s.OffsetTable = v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatus) SetTotalCount(v int64) *GetTopicStatusResponseBodyTopicStatus {
	s.TotalCount = &v
	return s
}

type GetTopicStatusResponseBodyTopicStatusOffsetTable struct {
	OffsetTable []*GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable `json:"OffsetTable,omitempty" xml:"OffsetTable,omitempty" type:"Repeated"`
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTable) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTable) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTable) SetOffsetTable(v []*GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) *GetTopicStatusResponseBodyTopicStatusOffsetTable {
	s.OffsetTable = v
	return s
}

type GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable struct {
	// The last time when the partition was modified.
	LastUpdateTimestamp *int64 `json:"LastUpdateTimestamp,omitempty" xml:"LastUpdateTimestamp,omitempty"`
	// The latest offset in the partition of the topic.
	MaxOffset *int64 `json:"MaxOffset,omitempty" xml:"MaxOffset,omitempty"`
	// The earliest offset in the partition of the topic.
	MinOffset *int64 `json:"MinOffset,omitempty" xml:"MinOffset,omitempty"`
	// The ID of the partition.
	Partition *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The name of the topic.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetLastUpdateTimestamp(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.LastUpdateTimestamp = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetMaxOffset(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.MaxOffset = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetMinOffset(v int64) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.MinOffset = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetPartition(v int32) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.Partition = &v
	return s
}

func (s *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable) SetTopic(v string) *GetTopicStatusResponseBodyTopicStatusOffsetTableOffsetTable {
	s.Topic = &v
	return s
}

type GetTopicStatusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTopicStatusResponse) SetHeaders(v map[string]*string) *GetTopicStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTopicStatusResponse) SetStatusCode(v int32) *GetTopicStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicStatusResponse) SetBody(v *GetTopicStatusResponseBody) *GetTopicStatusResponse {
	s.Body = v
	return s
}

type GetTopicSubscribeStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicSubscribeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSubscribeStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTopicSubscribeStatusRequest) SetInstanceId(v string) *GetTopicSubscribeStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTopicSubscribeStatusRequest) SetRegionId(v string) *GetTopicSubscribeStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetTopicSubscribeStatusRequest) SetTopic(v string) *GetTopicSubscribeStatusRequest {
	s.Topic = &v
	return s
}

type GetTopicSubscribeStatusResponseBody struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *bool                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TopicSubscribeStatus *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus `json:"TopicSubscribeStatus,omitempty" xml:"TopicSubscribeStatus,omitempty" type:"Struct"`
}

func (s GetTopicSubscribeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSubscribeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicSubscribeStatusResponseBody) SetCode(v int32) *GetTopicSubscribeStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) SetMessage(v string) *GetTopicSubscribeStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) SetRequestId(v string) *GetTopicSubscribeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) SetSuccess(v bool) *GetTopicSubscribeStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicSubscribeStatusResponseBody) SetTopicSubscribeStatus(v *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) *GetTopicSubscribeStatusResponseBody {
	s.TopicSubscribeStatus = v
	return s
}

type GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus struct {
	ConsumerGroups []*string `json:"ConsumerGroups,omitempty" xml:"ConsumerGroups,omitempty" type:"Repeated"`
	Topic          *string   `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) GoString() string {
	return s.String()
}

func (s *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) SetConsumerGroups(v []*string) *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus {
	s.ConsumerGroups = v
	return s
}

func (s *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus) SetTopic(v string) *GetTopicSubscribeStatusResponseBodyTopicSubscribeStatus {
	s.Topic = &v
	return s
}

type GetTopicSubscribeStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicSubscribeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicSubscribeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicSubscribeStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTopicSubscribeStatusResponse) SetHeaders(v map[string]*string) *GetTopicSubscribeStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTopicSubscribeStatusResponse) SetStatusCode(v int32) *GetTopicSubscribeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicSubscribeStatusResponse) SetBody(v *GetTopicSubscribeStatusResponseBody) *GetTopicSubscribeStatusResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The token that determines the start point of the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the resource is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID. The following items describe the formats of resource IDs:
	//
	// *   Instance ID: instanceId
	// *   Topic ID: Kafka_alikafka_instanceId_topic
	// *   Group ID: Kafka_alikafka_instanceId_consumerGroup
	//
	// For example, you create an instance whose ID is alikafka_post-cn-v0h1fgs2xxxx, a topic whose name is test-topic, and a group whose ID is test-consumer-group. In this case, the resource IDs are alikafka_post-cn-v0h1fgs2xxxx, Kafka_alikafka_post-cn-v0h1fgs2xxxx_test-topic, and Kafka_alikafka_post-cn-v0h1fgs2xxxx_test-consumer-group.
	//
	// >  You must specify one of the **ResourceId** and **Tag** parameters to query the tags that are attached to a resource. Otherwise, the call fails.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource whose tags you want to query. The value is an enumerated value. Valid values:
	//
	// *   **INSTANCE**
	// *   **TOPIC**
	// *   **CONSUMERGROUP**
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The tag key.
	//
	// *   If you leave this parameter empty, the keys of all tags are matched.
	// *   The tag key must be 1 to 128 characters in length and cannot start with acs: or aliyun. The tag key cannot contain http:// or https://.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// *   If you do not specify the tag key, you cannot specify the tag value. If you leave this parameter empty, the values of all tags are matched.
	// *   The tag value must be 1 to 128 characters in length and cannot start with acs: or aliyun. The tag value cannot contain http:// or https://.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The token that determines the start point of the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the resource and tags, such as the resource ID, the resource type, tag keys, and tag values.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the resource. A resource ID complies with the following rules:
	//
	// *   The resource ID of an instance is the value of the instanceId parameter.
	// *   The resource ID of a topic is the value of the Kafka_alikafka_instanceId_topic parameter.
	// *   The resource ID of a consumer group is the value of the Kafka_alikafka_instanceId_consumerGroup parameter.
	//
	// For example, the resources whose tags you want to query include the alikafka_post-cn-v0h1fgs2xxxx instance, the test-topic topic, and the test-consumer-group consumer group. In this case, their resource IDs are alikafka_post-cn-v0h1fgs2xxxx, Kafka_alikafka_post-cn-v0h1fgs2xxxx_test-topic, and Kafka_alikafka_post-cn-v0h1fgs2xxxx_test-consumer-group.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. The value is an enumerated value. Valid values:
	//
	// *   **Instance**
	// *   **Topic**
	// *   **Consumergroup**
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
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

type ModifyInstanceNameRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name. Valid values:
	//
	// *   The name can contain only letters, digits, hyphens (-), and underscores (\_).
	// *   The name must be 3 to 64 characters in length. A name that contains more than 64 characters is automatically truncated.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameRequest) SetInstanceId(v string) *ModifyInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetInstanceName(v string) *ModifyInstanceNameRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetRegionId(v string) *ModifyInstanceNameRequest {
	s.RegionId = &v
	return s
}

type ModifyInstanceNameResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponseBody) SetCode(v int32) *ModifyInstanceNameResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) SetMessage(v string) *ModifyInstanceNameResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) SetRequestId(v string) *ModifyInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) SetSuccess(v bool) *ModifyInstanceNameResponseBody {
	s.Success = &v
	return s
}

type ModifyInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponse) SetHeaders(v map[string]*string) *ModifyInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNameResponse) SetStatusCode(v int32) *ModifyInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceNameResponse) SetBody(v *ModifyInstanceNameResponseBody) *ModifyInstanceNameResponse {
	s.Body = v
	return s
}

type ModifyPartitionNumRequest struct {
	// The number of partitions that you want to add to the topic.
	//
	// *   The value must be an integer that is greater than 0.
	// *   To reduce the risk of data skew, we recommend that you set the value to a multiple of 6.
	// *   The number of total partitions ranges from 1 to 360.
	AddPartitionNum *int32 `json:"AddPartitionNum,omitempty" xml:"AddPartitionNum,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ModifyPartitionNumRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPartitionNumRequest) GoString() string {
	return s.String()
}

func (s *ModifyPartitionNumRequest) SetAddPartitionNum(v int32) *ModifyPartitionNumRequest {
	s.AddPartitionNum = &v
	return s
}

func (s *ModifyPartitionNumRequest) SetInstanceId(v string) *ModifyPartitionNumRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPartitionNumRequest) SetRegionId(v string) *ModifyPartitionNumRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPartitionNumRequest) SetTopic(v string) *ModifyPartitionNumRequest {
	s.Topic = &v
	return s
}

type ModifyPartitionNumResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyPartitionNumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPartitionNumResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPartitionNumResponseBody) SetCode(v int32) *ModifyPartitionNumResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyPartitionNumResponseBody) SetMessage(v string) *ModifyPartitionNumResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyPartitionNumResponseBody) SetRequestId(v string) *ModifyPartitionNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPartitionNumResponseBody) SetSuccess(v bool) *ModifyPartitionNumResponseBody {
	s.Success = &v
	return s
}

type ModifyPartitionNumResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPartitionNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPartitionNumResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPartitionNumResponse) GoString() string {
	return s.String()
}

func (s *ModifyPartitionNumResponse) SetHeaders(v map[string]*string) *ModifyPartitionNumResponse {
	s.Headers = v
	return s
}

func (s *ModifyPartitionNumResponse) SetStatusCode(v int32) *ModifyPartitionNumResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPartitionNumResponse) SetBody(v *ModifyPartitionNumResponseBody) *ModifyPartitionNumResponse {
	s.Body = v
	return s
}

type ModifyTopicRemarkRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the topic.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the topic.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ModifyTopicRemarkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTopicRemarkRequest) GoString() string {
	return s.String()
}

func (s *ModifyTopicRemarkRequest) SetInstanceId(v string) *ModifyTopicRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTopicRemarkRequest) SetRegionId(v string) *ModifyTopicRemarkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyTopicRemarkRequest) SetRemark(v string) *ModifyTopicRemarkRequest {
	s.Remark = &v
	return s
}

func (s *ModifyTopicRemarkRequest) SetTopic(v string) *ModifyTopicRemarkRequest {
	s.Topic = &v
	return s
}

type ModifyTopicRemarkResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyTopicRemarkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTopicRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTopicRemarkResponseBody) SetCode(v int32) *ModifyTopicRemarkResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyTopicRemarkResponseBody) SetMessage(v string) *ModifyTopicRemarkResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyTopicRemarkResponseBody) SetRequestId(v string) *ModifyTopicRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTopicRemarkResponseBody) SetSuccess(v bool) *ModifyTopicRemarkResponseBody {
	s.Success = &v
	return s
}

type ModifyTopicRemarkResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTopicRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTopicRemarkResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTopicRemarkResponse) GoString() string {
	return s.String()
}

func (s *ModifyTopicRemarkResponse) SetHeaders(v map[string]*string) *ModifyTopicRemarkResponse {
	s.Headers = v
	return s
}

func (s *ModifyTopicRemarkResponse) SetStatusCode(v int32) *ModifyTopicRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTopicRemarkResponse) SetBody(v *ModifyTopicRemarkResponseBody) *ModifyTopicRemarkResponse {
	s.Body = v
	return s
}

type QueryMessageRequest struct {
	BeginTime  *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Offset     *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Partition  *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	QueryType  *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageRequest) GoString() string {
	return s.String()
}

func (s *QueryMessageRequest) SetBeginTime(v int64) *QueryMessageRequest {
	s.BeginTime = &v
	return s
}

func (s *QueryMessageRequest) SetInstanceId(v string) *QueryMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMessageRequest) SetOffset(v string) *QueryMessageRequest {
	s.Offset = &v
	return s
}

func (s *QueryMessageRequest) SetPartition(v string) *QueryMessageRequest {
	s.Partition = &v
	return s
}

func (s *QueryMessageRequest) SetQueryType(v string) *QueryMessageRequest {
	s.QueryType = &v
	return s
}

func (s *QueryMessageRequest) SetRegionId(v string) *QueryMessageRequest {
	s.RegionId = &v
	return s
}

func (s *QueryMessageRequest) SetTopic(v string) *QueryMessageRequest {
	s.Topic = &v
	return s
}

type QueryMessageResponseBody struct {
	Code        *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	MessageList []*QueryMessageResponseBodyMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBody) SetCode(v int32) *QueryMessageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessage(v string) *QueryMessageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMessageResponseBody) SetMessageList(v []*QueryMessageResponseBodyMessageList) *QueryMessageResponseBody {
	s.MessageList = v
	return s
}

func (s *QueryMessageResponseBody) SetRequestId(v string) *QueryMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMessageResponseBody) SetSuccess(v bool) *QueryMessageResponseBody {
	s.Success = &v
	return s
}

type QueryMessageResponseBodyMessageList struct {
	Checksum            *int64  `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	Key                 *string `json:"Key,omitempty" xml:"Key,omitempty"`
	KeyTruncated        *bool   `json:"KeyTruncated,omitempty" xml:"KeyTruncated,omitempty"`
	Offset              *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	Partition           *int64  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	SerializedKeySize   *int32  `json:"SerializedKeySize,omitempty" xml:"SerializedKeySize,omitempty"`
	SerializedValueSize *int32  `json:"SerializedValueSize,omitempty" xml:"SerializedValueSize,omitempty"`
	Timestamp           *int64  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TimestampType       *string `json:"TimestampType,omitempty" xml:"TimestampType,omitempty"`
	Topic               *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TruncatedKeySize    *int32  `json:"TruncatedKeySize,omitempty" xml:"TruncatedKeySize,omitempty"`
	TruncatedValueSize  *int32  `json:"TruncatedValueSize,omitempty" xml:"TruncatedValueSize,omitempty"`
	Value               *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ValueTruncated      *bool   `json:"ValueTruncated,omitempty" xml:"ValueTruncated,omitempty"`
}

func (s QueryMessageResponseBodyMessageList) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageResponseBodyMessageList) GoString() string {
	return s.String()
}

func (s *QueryMessageResponseBodyMessageList) SetChecksum(v int64) *QueryMessageResponseBodyMessageList {
	s.Checksum = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetKey(v string) *QueryMessageResponseBodyMessageList {
	s.Key = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetKeyTruncated(v bool) *QueryMessageResponseBodyMessageList {
	s.KeyTruncated = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetOffset(v int64) *QueryMessageResponseBodyMessageList {
	s.Offset = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetPartition(v int64) *QueryMessageResponseBodyMessageList {
	s.Partition = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetSerializedKeySize(v int32) *QueryMessageResponseBodyMessageList {
	s.SerializedKeySize = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetSerializedValueSize(v int32) *QueryMessageResponseBodyMessageList {
	s.SerializedValueSize = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTimestamp(v int64) *QueryMessageResponseBodyMessageList {
	s.Timestamp = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTimestampType(v string) *QueryMessageResponseBodyMessageList {
	s.TimestampType = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTopic(v string) *QueryMessageResponseBodyMessageList {
	s.Topic = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTruncatedKeySize(v int32) *QueryMessageResponseBodyMessageList {
	s.TruncatedKeySize = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetTruncatedValueSize(v int32) *QueryMessageResponseBodyMessageList {
	s.TruncatedValueSize = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetValue(v string) *QueryMessageResponseBodyMessageList {
	s.Value = &v
	return s
}

func (s *QueryMessageResponseBodyMessageList) SetValueTruncated(v bool) *QueryMessageResponseBodyMessageList {
	s.ValueTruncated = &v
	return s
}

type QueryMessageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMessageResponse) GoString() string {
	return s.String()
}

func (s *QueryMessageResponse) SetHeaders(v map[string]*string) *QueryMessageResponse {
	s.Headers = v
	return s
}

func (s *QueryMessageResponse) SetStatusCode(v int32) *QueryMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMessageResponse) SetBody(v *QueryMessageResponseBody) *QueryMessageResponse {
	s.Body = v
	return s
}

type ReleaseInstanceRequest struct {
	// Specifies whether to immediately release the physical resources of the instance. Valid values:
	//
	// *   **true**: The physical resources of the instance are immediately released.
	// *   **false**: The physical resources of the instance are retained for a period of time before they are released.
	ForceDeleteInstance *bool `json:"ForceDeleteInstance,omitempty" xml:"ForceDeleteInstance,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) SetForceDeleteInstance(v bool) *ReleaseInstanceRequest {
	s.ForceDeleteInstance = &v
	return s
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseInstanceRequest) SetRegionId(v string) *ReleaseInstanceRequest {
	s.RegionId = &v
	return s
}

type ReleaseInstanceResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) SetCode(v int32) *ReleaseInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetMessage(v string) *ReleaseInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetSuccess(v bool) *ReleaseInstanceResponseBody {
	s.Success = &v
	return s
}

type ReleaseInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponse) SetHeaders(v map[string]*string) *ReleaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceResponse) SetStatusCode(v int32) *ReleaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceResponse) SetBody(v *ReleaseInstanceResponseBody) *ReleaseInstanceResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	// The initial configurations of ApsaraMQ for Kafka. The value must be a valid JSON string.
	//
	// If you do not specify this parameter, it is left empty.
	//
	// The following parameters can be configured for **Config**:
	//
	// *   **enable.vpc_sasl_ssl**: specifies whether to enable VPC transmission encryption. Valid values:
	//
	//     *   **true**: enables VPC transmission encryption. If you enable VPC transmission encryption, you must also enable access control list (ACL).
	//     *   **false**: disables VPC transmission encryption. By default, VPC transmission encryption is disabled.
	//
	// *   **enable.acl**: specifies whether to enable ACL. Valid values:
	//
	//     *   **true**: enables ACL.
	//     *   **false**: disables ACL. By default, ACL is disabled.
	//
	// *   **kafka.log.retention.hours**: the maximum message retention period when the disk capacity is sufficient. Unit: hours. Valid values: 24 to 480. Default value: **72**. When the disk usage reaches 85%, the disk capacity is considered insufficient and the system deletes messages in the order in which they are stored to ensure service availability.
	//
	// *   **kafka.message.max.bytes**: the maximum size of messages that ApsaraMQ for Kafka can send and receive. Unit: bytes. Valid values: 1048576 to 10485760. Default value: **1048576**. Before you change the value of this parameter, make sure that the new value matches the corresponding configurations on the producers and consumers.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether cross-zone deployment is required. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: true.
	CrossZone *bool `json:"CrossZone,omitempty" xml:"CrossZone,omitempty"`
	// The deployment mode of the instance. Valid values:
	//
	// *   **vpc**: deploys the instance that allows access only from a VPC.
	// *   **eip**: deploys the instance that allows access from the Internet and a VPC.
	//
	// The deployment mode of the instance must match the type of the instance. If the instance allows access only from a VPC, set the value to **vpc**. If the instance allows access from the Internet and a VPC, set the value to **eip**.
	DeployModule *string `json:"DeployModule,omitempty" xml:"DeployModule,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the instance supports elastic IP addresses (EIPs). Valid values:
	//
	// *   **true**: supports EIPs and allows access from the Internet and a VPC.
	// *   **false**: does not support EIPs and allows access only from a VPC.
	//
	// The value of this parameter must match the type of the instance. For example, if the instance allows access only from a VPC, set this parameter to **false**.
	IsEipInner *bool `json:"IsEipInner,omitempty" xml:"IsEipInner,omitempty"`
	// Specifies whether to forcibly deploy the instance in the selected zones.
	IsForceSelectedZones *bool `json:"IsForceSelectedZones,omitempty" xml:"IsForceSelectedZones,omitempty"`
	// Specifies whether to set a new username and password. Valid values:
	//
	// *   **true**: sets a new username and password.
	// *   **false**: does not set a new username or password.
	//
	// This parameter is available only if you deploy an instance that allows access from the Internet and a VPC.
	IsSetUserAndPassword *bool `json:"IsSetUserAndPassword,omitempty" xml:"IsSetUserAndPassword,omitempty"`
	// The ID of the key that is used for disk encryption in the region where the instance is deployed. You can obtain the ID of the key in the [Key Management Service (KMS) console](https://kms.console.aliyun.com/?spm=a2c4g.11186623.2.5.336745b8hfiU21) or create a key. For more information, see [Manage CMKs](~~181610~~).
	//
	// If this parameter is configured, disk encryption is enabled for the instance. You cannot disable disk encryption after disk encryption is enabled. When you call this operation, the system checks whether the AliyunServiceRoleForAlikafkaInstanceEncryption service-linked role is created. If the role is not created, the system automatically creates the role. For more information, see [Service-linked roles](~~190460~~).
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The name of the instance.
	//
	// >  If you specify a value for this parameter, make sure that the specified value is unique in the region of the instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert contact.
	Notifier *string `json:"Notifier,omitempty" xml:"Notifier,omitempty"`
	// The password that corresponds to the username.
	//
	// This parameter is available only if you deploy an instance that allows access from the Internet and a VPC.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group of the instance.
	//
	// If you do not specify this parameter, ApsaraMQ for Kafka automatically configures a security group for your instance. If you specify this parameter, you must create a security group in advance. For more information, see [Create a security group](~~25468~~).
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The two-dimensional arrays that consist of the candidate set for primary zones and the candidate set for secondary zones.
	//
	// *   If you set CrossZone to true and specify Zone H and Zone F as the candidate set for primary zones and Zone K as the candidate set for secondary zones, set this parameter to `[[\"zoneh\",\"zonef\"],[\"zonek\"]]`.
	//
	//     **
	//
	//     **Note** If you specify multiple zones as the primary or secondary zones, the system deploys the instance in one of the zones without prioritizing them. For example, if you set this parameter to `[[\"zoneh\",\"zonef\"],[\"zonek\"]]`, the primary zone in which the instance is deployed can be Zone H or Zone F, and the secondary zone is Zone K.
	//
	// *   If you set CrossZone to false and want to deploy the instance in Zone K, set this parameter to `[[\"zonek\"],[]]`. In this case, the value of this parameter must still be two-dimensional arrays, but the array that specifies the candidate for secondary zones is left empty.
	SelectedZones *string `json:"SelectedZones,omitempty" xml:"SelectedZones,omitempty"`
	// The version of ApsaraMQ for Kafka. Valid values: 0.10.2 and 2.2.0.
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The mobile phone number of the alert contact.
	UserPhoneNum *string `json:"UserPhoneNum,omitempty" xml:"UserPhoneNum,omitempty"`
	// The username that is used to access the instance.
	//
	// This parameter is available only if you deploy an instance that allows access from the Internet and a VPC.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The ID of the vSwitch to which you want to connect the instance.
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which you want to deploy the instance.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone in which you want to deploy the instance.
	//
	// *   The zone ID of the instance must be the same as that of the vSwitch.
	// *   The value must be in the format of zoneX or Region ID-X. For example, you can set this parameter to zonea or cn-hangzhou-k.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetConfig(v string) *StartInstanceRequest {
	s.Config = &v
	return s
}

func (s *StartInstanceRequest) SetCrossZone(v bool) *StartInstanceRequest {
	s.CrossZone = &v
	return s
}

func (s *StartInstanceRequest) SetDeployModule(v string) *StartInstanceRequest {
	s.DeployModule = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetIsEipInner(v bool) *StartInstanceRequest {
	s.IsEipInner = &v
	return s
}

func (s *StartInstanceRequest) SetIsForceSelectedZones(v bool) *StartInstanceRequest {
	s.IsForceSelectedZones = &v
	return s
}

func (s *StartInstanceRequest) SetIsSetUserAndPassword(v bool) *StartInstanceRequest {
	s.IsSetUserAndPassword = &v
	return s
}

func (s *StartInstanceRequest) SetKMSKeyId(v string) *StartInstanceRequest {
	s.KMSKeyId = &v
	return s
}

func (s *StartInstanceRequest) SetName(v string) *StartInstanceRequest {
	s.Name = &v
	return s
}

func (s *StartInstanceRequest) SetNotifier(v string) *StartInstanceRequest {
	s.Notifier = &v
	return s
}

func (s *StartInstanceRequest) SetPassword(v string) *StartInstanceRequest {
	s.Password = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstanceRequest) SetSecurityGroup(v string) *StartInstanceRequest {
	s.SecurityGroup = &v
	return s
}

func (s *StartInstanceRequest) SetSelectedZones(v string) *StartInstanceRequest {
	s.SelectedZones = &v
	return s
}

func (s *StartInstanceRequest) SetServiceVersion(v string) *StartInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *StartInstanceRequest) SetUserPhoneNum(v string) *StartInstanceRequest {
	s.UserPhoneNum = &v
	return s
}

func (s *StartInstanceRequest) SetUsername(v string) *StartInstanceRequest {
	s.Username = &v
	return s
}

func (s *StartInstanceRequest) SetVSwitchId(v string) *StartInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *StartInstanceRequest) SetVpcId(v string) *StartInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *StartInstanceRequest) SetZoneId(v string) *StartInstanceRequest {
	s.ZoneId = &v
	return s
}

type StartInstanceResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetCode(v int32) *StartInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartInstanceResponseBody) SetMessage(v string) *StartInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) SetSuccess(v bool) *StartInstanceResponseBody {
	s.Success = &v
	return s
}

type StartInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of the Message Queue for Apache RocketMQ instance which contains the resource to which you want to attach tags.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the resource is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resources. The value is an enumerated value. Valid values:
	//
	// *   **INSTANCE**
	// *   **TOPIC**
	// *   **CONSUMERGROUP**
	//
	// >  The value of this parameter is not case-sensitive.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags that you want to add.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetInstanceId(v string) *TagResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The tag key.
	//
	// *   You must specify this parameter.
	// *   The tag key must be 1 to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// *   You can leave this parameter empty.
	// *   The tag value must be 1 to 128 characters in length and cannot start with acs: or aliyun. The tag key cannot contain http:// or https://.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
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
	// Specifies whether to detach all tags from the resource. This parameter only takes effect when the TagKey.N parameter is not configured. Default value: **false**.
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The ID of the region in which the resource is deployed.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources from which you want to detach tags.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resources. Valid values:
	//
	// *   **INSTANCE**
	// *   **TOPIC**
	// *   **CONSUMERGROUP**
	//
	// >  The value of this parameter is not case-sensitive.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
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

type UpdateAllowedIpRequest struct {
	// The IP addresses that you want to manage. You can specify a CIDR block. Example: **192.168.0.0/16**.
	//
	// *   If the **UpdateType** parameter is set to **add**, specify one or more IP addresses for this parameter. Separate multiple IP addresses with commas (,).
	// *   If the **UpdateType** parameter is set to **delete**, specify only one IP address.
	// *   Exercise caution when you delete IP addresses.
	AllowedListIp *string `json:"AllowedListIp,omitempty" xml:"AllowedListIp,omitempty"`
	// The type of the whitelist. Valid values:
	//
	// *   **vpc**: a whitelist for access from a VPC.
	// *   **internet**: a whitelist for access from the Internet.
	AllowedListType *string `json:"AllowedListType,omitempty" xml:"AllowedListType,omitempty"`
	// The description of the whitelist.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The port range. Valid values:
	//
	// *   **9092/9092**: the port range for access from a virtual private cloud (VPC).
	// *   **9093/9093**: the port range for access from the Internet.
	//
	// The value of this parameter must match the value of the **AllowdedListType** parameter.
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the region where the instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of configuration change. Valid values:
	//
	// *   **add**
	// *   **delete**
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
}

func (s UpdateAllowedIpRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAllowedIpRequest) GoString() string {
	return s.String()
}

func (s *UpdateAllowedIpRequest) SetAllowedListIp(v string) *UpdateAllowedIpRequest {
	s.AllowedListIp = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetAllowedListType(v string) *UpdateAllowedIpRequest {
	s.AllowedListType = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetDescription(v string) *UpdateAllowedIpRequest {
	s.Description = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetInstanceId(v string) *UpdateAllowedIpRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetPortRange(v string) *UpdateAllowedIpRequest {
	s.PortRange = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetRegionId(v string) *UpdateAllowedIpRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetUpdateType(v string) *UpdateAllowedIpRequest {
	s.UpdateType = &v
	return s
}

type UpdateAllowedIpResponseBody struct {
	// The HTTP status code that is returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAllowedIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAllowedIpResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAllowedIpResponseBody) SetCode(v int32) *UpdateAllowedIpResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAllowedIpResponseBody) SetMessage(v string) *UpdateAllowedIpResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAllowedIpResponseBody) SetRequestId(v string) *UpdateAllowedIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAllowedIpResponseBody) SetSuccess(v bool) *UpdateAllowedIpResponseBody {
	s.Success = &v
	return s
}

type UpdateAllowedIpResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAllowedIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAllowedIpResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAllowedIpResponse) GoString() string {
	return s.String()
}

func (s *UpdateAllowedIpResponse) SetHeaders(v map[string]*string) *UpdateAllowedIpResponse {
	s.Headers = v
	return s
}

func (s *UpdateAllowedIpResponse) SetStatusCode(v int32) *UpdateAllowedIpResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAllowedIpResponse) SetBody(v *UpdateAllowedIpResponseBody) *UpdateAllowedIpResponse {
	s.Body = v
	return s
}

type UpdateConsumerOffsetRequest struct {
	// The name of the consumer group.
	//
	// *   The name can contain letters, digits, hyphens (-), and underscores (\_).
	// *   The name must be **3 to 64** characters in length. If a name contains more than **64** characters, the name is automatically truncated.
	// *   The name of a consumer group cannot be changed after the consumer group is created.
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// If you set resetType to offset, you can use this parameter to reset the consumer offset of each partition of a specific topic in the consumer group.
	Offsets []*UpdateConsumerOffsetRequestOffsets `json:"Offsets,omitempty" xml:"Offsets,omitempty" type:"Repeated"`
	// The region ID of the instance to which the consumer group belongs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method that is used to reset the consumer offsets of the subscribed topics of a consumer group. Valid values:
	//
	// *   **timestamp** (default)
	// *   **offset**
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The point in time when message consumption starts. The value of this parameter is a UNIX timestamp. Unit: milliseconds. The value of this parameter must be **less than 0** or **within the retention period of the consumer offset**. This parameter takes effect only if you set resetType to timestamp.
	//
	// **If you want to reset the consumer offset to the latest offset, specify a value that is less than 0. Recommended value: -1.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The topic name.
	//
	// *   The name can contain letters, digits, underscores (\_), and hyphens (-).
	// *   The name must be **3 to 64** characters in length. If a name contains more than **64** characters, the name is automatically truncated.
	// *   The name of a topic cannot be changed after the topic is created.
	//
	// **If you want to reset the consumer offsets of all topics to which the consumer subscribes, specify an empty string.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateConsumerOffsetRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerOffsetRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetRequest) SetConsumerId(v string) *UpdateConsumerOffsetRequest {
	s.ConsumerId = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetInstanceId(v string) *UpdateConsumerOffsetRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetOffsets(v []*UpdateConsumerOffsetRequestOffsets) *UpdateConsumerOffsetRequest {
	s.Offsets = v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetRegionId(v string) *UpdateConsumerOffsetRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetResetType(v string) *UpdateConsumerOffsetRequest {
	s.ResetType = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetTime(v string) *UpdateConsumerOffsetRequest {
	s.Time = &v
	return s
}

func (s *UpdateConsumerOffsetRequest) SetTopic(v string) *UpdateConsumerOffsetRequest {
	s.Topic = &v
	return s
}

type UpdateConsumerOffsetRequestOffsets struct {
	// The consumer offset of the partition.
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The partition ID.
	Partition *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
}

func (s UpdateConsumerOffsetRequestOffsets) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerOffsetRequestOffsets) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetRequestOffsets) SetOffset(v int64) *UpdateConsumerOffsetRequestOffsets {
	s.Offset = &v
	return s
}

func (s *UpdateConsumerOffsetRequestOffsets) SetPartition(v int32) *UpdateConsumerOffsetRequestOffsets {
	s.Partition = &v
	return s
}

type UpdateConsumerOffsetShrinkRequest struct {
	// The name of the consumer group.
	//
	// *   The name can contain letters, digits, hyphens (-), and underscores (\_).
	// *   The name must be **3 to 64** characters in length. If a name contains more than **64** characters, the name is automatically truncated.
	// *   The name of a consumer group cannot be changed after the consumer group is created.
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// If you set resetType to offset, you can use this parameter to reset the consumer offset of each partition of a specific topic in the consumer group.
	OffsetsShrink *string `json:"Offsets,omitempty" xml:"Offsets,omitempty"`
	// The region ID of the instance to which the consumer group belongs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method that is used to reset the consumer offsets of the subscribed topics of a consumer group. Valid values:
	//
	// *   **timestamp** (default)
	// *   **offset**
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The point in time when message consumption starts. The value of this parameter is a UNIX timestamp. Unit: milliseconds. The value of this parameter must be **less than 0** or **within the retention period of the consumer offset**. This parameter takes effect only if you set resetType to timestamp.
	//
	// **If you want to reset the consumer offset to the latest offset, specify a value that is less than 0. Recommended value: -1.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The topic name.
	//
	// *   The name can contain letters, digits, underscores (\_), and hyphens (-).
	// *   The name must be **3 to 64** characters in length. If a name contains more than **64** characters, the name is automatically truncated.
	// *   The name of a topic cannot be changed after the topic is created.
	//
	// **If you want to reset the consumer offsets of all topics to which the consumer subscribes, specify an empty string.
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s UpdateConsumerOffsetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerOffsetShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetShrinkRequest) SetConsumerId(v string) *UpdateConsumerOffsetShrinkRequest {
	s.ConsumerId = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetInstanceId(v string) *UpdateConsumerOffsetShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetOffsetsShrink(v string) *UpdateConsumerOffsetShrinkRequest {
	s.OffsetsShrink = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetRegionId(v string) *UpdateConsumerOffsetShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetResetType(v string) *UpdateConsumerOffsetShrinkRequest {
	s.ResetType = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetTime(v string) *UpdateConsumerOffsetShrinkRequest {
	s.Time = &v
	return s
}

func (s *UpdateConsumerOffsetShrinkRequest) SetTopic(v string) *UpdateConsumerOffsetShrinkRequest {
	s.Topic = &v
	return s
}

type UpdateConsumerOffsetResponseBody struct {
	// The HTTP status code that is returned. The status code **200** indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateConsumerOffsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerOffsetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetResponseBody) SetCode(v int32) *UpdateConsumerOffsetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateConsumerOffsetResponseBody) SetMessage(v string) *UpdateConsumerOffsetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateConsumerOffsetResponseBody) SetRequestId(v string) *UpdateConsumerOffsetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConsumerOffsetResponseBody) SetSuccess(v bool) *UpdateConsumerOffsetResponseBody {
	s.Success = &v
	return s
}

type UpdateConsumerOffsetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConsumerOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConsumerOffsetResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConsumerOffsetResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetResponse) SetHeaders(v map[string]*string) *UpdateConsumerOffsetResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerOffsetResponse) SetStatusCode(v int32) *UpdateConsumerOffsetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConsumerOffsetResponse) SetBody(v *UpdateConsumerOffsetResponseBody) *UpdateConsumerOffsetResponse {
	s.Body = v
	return s
}

type UpdateInstanceConfigRequest struct {
	// The configurations that you want to update for the ApsaraMQ for Kafka instance. The value must be a valid JSON string.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateInstanceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConfigRequest) SetConfig(v string) *UpdateInstanceConfigRequest {
	s.Config = &v
	return s
}

func (s *UpdateInstanceConfigRequest) SetInstanceId(v string) *UpdateInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceConfigRequest) SetRegionId(v string) *UpdateInstanceConfigRequest {
	s.RegionId = &v
	return s
}

type UpdateInstanceConfigResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConfigResponseBody) SetCode(v int32) *UpdateInstanceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceConfigResponseBody) SetMessage(v string) *UpdateInstanceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceConfigResponseBody) SetRequestId(v string) *UpdateInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceConfigResponseBody) SetSuccess(v bool) *UpdateInstanceConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateInstanceConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceConfigResponse) SetHeaders(v map[string]*string) *UpdateInstanceConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceConfigResponse) SetStatusCode(v int32) *UpdateInstanceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceConfigResponse) SetBody(v *UpdateInstanceConfigResponseBody) *UpdateInstanceConfigResponse {
	s.Body = v
	return s
}

type UpdateTopicConfigRequest struct {
	Config     *string `json:"Config,omitempty" xml:"Config,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	Value      *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTopicConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateTopicConfigRequest) SetConfig(v string) *UpdateTopicConfigRequest {
	s.Config = &v
	return s
}

func (s *UpdateTopicConfigRequest) SetInstanceId(v string) *UpdateTopicConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTopicConfigRequest) SetRegionId(v string) *UpdateTopicConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTopicConfigRequest) SetTopic(v string) *UpdateTopicConfigRequest {
	s.Topic = &v
	return s
}

func (s *UpdateTopicConfigRequest) SetValue(v string) *UpdateTopicConfigRequest {
	s.Value = &v
	return s
}

type UpdateTopicConfigResponseBody struct {
	Code    *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTopicConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTopicConfigResponseBody) SetCode(v int64) *UpdateTopicConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) SetData(v string) *UpdateTopicConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) SetMessage(v string) *UpdateTopicConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) SetRequestId(v string) *UpdateTopicConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTopicConfigResponseBody) SetSuccess(v bool) *UpdateTopicConfigResponseBody {
	s.Success = &v
	return s
}

type UpdateTopicConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTopicConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTopicConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTopicConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateTopicConfigResponse) SetHeaders(v map[string]*string) *UpdateTopicConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateTopicConfigResponse) SetStatusCode(v int32) *UpdateTopicConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTopicConfigResponse) SetBody(v *UpdateTopicConfigResponseBody) *UpdateTopicConfigResponse {
	s.Body = v
	return s
}

type UpgradeInstanceVersionRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The major version to be upgraded to. Valid values:
	//
	// *   **0.10.2**
	// *   **2.2.0**
	//
	// If you set this parameter to the current major version, the system upgrades the instance to the latest minor version.
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s UpgradeInstanceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionRequest) SetInstanceId(v string) *UpgradeInstanceVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRegionId(v string) *UpgradeInstanceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetTargetVersion(v string) *UpgradeInstanceVersionRequest {
	s.TargetVersion = &v
	return s
}

type UpgradeInstanceVersionResponseBody struct {
	// The HTTP status code that is returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeInstanceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponseBody) SetCode(v int32) *UpgradeInstanceVersionResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetMessage(v string) *UpgradeInstanceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetRequestId(v string) *UpgradeInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetSuccess(v bool) *UpgradeInstanceVersionResponseBody {
	s.Success = &v
	return s
}

type UpgradeInstanceVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeInstanceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponse) SetHeaders(v map[string]*string) *UpgradeInstanceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetStatusCode(v int32) *UpgradeInstanceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeInstanceVersionResponse) SetBody(v *UpgradeInstanceVersionResponseBody) *UpgradeInstanceVersionResponse {
	s.Body = v
	return s
}

type UpgradePostPayOrderRequest struct {
	// The disk size. Unit: GB.
	//
	// *   The disk size that you specify must be greater than or equal to the current disk size of the instance.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The Internet traffic for the instance.
	//
	// *   The Internet traffic volume that you specify must be greater than or equal to the current Internet traffic volume of the instance.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	// > - If the **EipModel** parameter is set to **true**, set the **EipMax** parameter to a value that is greater than 0.
	// > - If the **EipModel** parameter is set to **false**, set the **EipMax** parameter to **0**.
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// *   true: enables Internet access.
	// *   false: disables Internet access.
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic for the instance. We recommend that you do not configure this parameter.
	//
	// *   The maximum traffic that you specify must be greater than or equal to the current maximum traffic of the instance.
	// *   You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you specify only the IoMaxSpec parameter.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// *   The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	// *   You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you specify only the IoMaxSpec parameter.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	RegionId         *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerlessConfig *UpgradePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The edition of the instance. Valid values:
	//
	// *   **normal**: Standard Edition (High Write)
	// *   **professional**: Professional Edition (High Write)
	// *   **professionalForHighRead**: Professional Edition (High Read)
	//
	// You cannot downgrade an instance from the Professional Edition to the Standard Edition. For more information about these instance editions, see [Billing](~~84737~~).
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   The default value of the TopicQuota parameter varies based on the value of the IoMaxSpec parameter. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s UpgradePostPayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradePostPayOrderRequest) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderRequest) SetDiskSize(v int32) *UpgradePostPayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetEipMax(v int32) *UpgradePostPayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetEipModel(v bool) *UpgradePostPayOrderRequest {
	s.EipModel = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetInstanceId(v string) *UpgradePostPayOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetIoMax(v int32) *UpgradePostPayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetIoMaxSpec(v string) *UpgradePostPayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetPartitionNum(v int32) *UpgradePostPayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetRegionId(v string) *UpgradePostPayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetServerlessConfig(v *UpgradePostPayOrderRequestServerlessConfig) *UpgradePostPayOrderRequest {
	s.ServerlessConfig = v
	return s
}

func (s *UpgradePostPayOrderRequest) SetSpecType(v string) *UpgradePostPayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetTopicQuota(v int32) *UpgradePostPayOrderRequest {
	s.TopicQuota = &v
	return s
}

type UpgradePostPayOrderRequestServerlessConfig struct {
	ReservedPublishCapacity   *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	ReservedSubscribeCapacity *int64 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
}

func (s UpgradePostPayOrderRequestServerlessConfig) String() string {
	return tea.Prettify(s)
}

func (s UpgradePostPayOrderRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderRequestServerlessConfig) SetReservedPublishCapacity(v int64) *UpgradePostPayOrderRequestServerlessConfig {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *UpgradePostPayOrderRequestServerlessConfig) SetReservedSubscribeCapacity(v int64) *UpgradePostPayOrderRequestServerlessConfig {
	s.ReservedSubscribeCapacity = &v
	return s
}

type UpgradePostPayOrderShrinkRequest struct {
	// The disk size. Unit: GB.
	//
	// *   The disk size that you specify must be greater than or equal to the current disk size of the instance.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The Internet traffic for the instance.
	//
	// *   The Internet traffic volume that you specify must be greater than or equal to the current Internet traffic volume of the instance.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	// > - If the **EipModel** parameter is set to **true**, set the **EipMax** parameter to a value that is greater than 0.
	// > - If the **EipModel** parameter is set to **false**, set the **EipMax** parameter to **0**.
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// *   true: enables Internet access.
	// *   false: disables Internet access.
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic for the instance. We recommend that you do not configure this parameter.
	//
	// *   The maximum traffic that you specify must be greater than or equal to the current maximum traffic of the instance.
	// *   You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you specify only the IoMaxSpec parameter.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// *   The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	// *   You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you specify only the IoMaxSpec parameter.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// The edition of the instance. Valid values:
	//
	// *   **normal**: Standard Edition (High Write)
	// *   **professional**: Professional Edition (High Write)
	// *   **professionalForHighRead**: Professional Edition (High Read)
	//
	// You cannot downgrade an instance from the Professional Edition to the Standard Edition. For more information about these instance editions, see [Billing](~~84737~~).
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   The default value of the TopicQuota parameter varies based on the value of the IoMaxSpec parameter. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	// *   For more information about the valid values, see [Billing](~~84737~~).
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s UpgradePostPayOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradePostPayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderShrinkRequest) SetDiskSize(v int32) *UpgradePostPayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetEipMax(v int32) *UpgradePostPayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetEipModel(v bool) *UpgradePostPayOrderShrinkRequest {
	s.EipModel = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetInstanceId(v string) *UpgradePostPayOrderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetIoMax(v int32) *UpgradePostPayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetIoMaxSpec(v string) *UpgradePostPayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetPartitionNum(v int32) *UpgradePostPayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetRegionId(v string) *UpgradePostPayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetServerlessConfigShrink(v string) *UpgradePostPayOrderShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetSpecType(v string) *UpgradePostPayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetTopicQuota(v int32) *UpgradePostPayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

type UpgradePostPayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradePostPayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradePostPayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderResponseBody) SetCode(v int32) *UpgradePostPayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradePostPayOrderResponseBody) SetMessage(v string) *UpgradePostPayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradePostPayOrderResponseBody) SetRequestId(v string) *UpgradePostPayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradePostPayOrderResponseBody) SetSuccess(v bool) *UpgradePostPayOrderResponseBody {
	s.Success = &v
	return s
}

type UpgradePostPayOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradePostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradePostPayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradePostPayOrderResponse) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderResponse) SetHeaders(v map[string]*string) *UpgradePostPayOrderResponse {
	s.Headers = v
	return s
}

func (s *UpgradePostPayOrderResponse) SetStatusCode(v int32) *UpgradePostPayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradePostPayOrderResponse) SetBody(v *UpgradePostPayOrderResponseBody) *UpgradePostPayOrderResponse {
	s.Body = v
	return s
}

type UpgradePrePayOrderRequest struct {
	// The size of the disk.
	//
	// *   The disk size that you specify must be greater than or equal to the current disk size of the instance.
	// *   For more information about the valid values, see [Billing overview](~~84737~~).
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The Internet traffic for the instance.
	//
	// *   The Internet traffic volume that you specify must be greater than or equal to the current Internet traffic volume of the instance.
	// *   For more information about the valid values, see [Billing overview](~~84737~~).
	// > - If the **EipModel** parameter is set to **true**, set the **EipMax** parameter to a value that is greater than 0.
	// > - If the **EipModel** parameter is set to **false**, set the **EipMax** parameter to **0**.
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// *   true: enables Internet access.
	// *   false: disables Internet access.
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic for the instance. We recommend that you do not configure this parameter.
	//
	// *   The maximum traffic volume that you specify must be greater than or equal to the current maximum traffic volume of the instance.
	// *   You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you configure only the IoMaxSpec parameter.
	// *   For more information about the valid values, see [Billing overview](~~84737~~).
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// *   The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	// *   You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you configure only the IoMaxSpec parameter.
	// *   For more information about the valid values, see [Billing overview](~~84737~~).
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   For more information about the valid values, see [Billing overview](~~84737~~).
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The edition of the instance. Valid values:
	//
	// *   **normal**: Standard Edition (High Write)
	// *   **professional**: Professional Edition (High Write)
	// *   **professionalForHighRead**: Professional Edition (High Read)
	//
	// You cannot downgrade an instance from the Professional Edition to the Standard Edition. For more information about these instance editions, see [Billing overview](~~84737~~).
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// *   You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	// *   If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	// *   The default value of the TopicQuota parameter varies based on the value of the IoMaxSpec parameter. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	// *   For more information about the valid values, see [Billing overview](~~84737~~).
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s UpgradePrePayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradePrePayOrderRequest) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderRequest) SetDiskSize(v int32) *UpgradePrePayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetEipMax(v int32) *UpgradePrePayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetEipModel(v bool) *UpgradePrePayOrderRequest {
	s.EipModel = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetInstanceId(v string) *UpgradePrePayOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetIoMax(v int32) *UpgradePrePayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetIoMaxSpec(v string) *UpgradePrePayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetPartitionNum(v int32) *UpgradePrePayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetRegionId(v string) *UpgradePrePayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetSpecType(v string) *UpgradePrePayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetTopicQuota(v int32) *UpgradePrePayOrderRequest {
	s.TopicQuota = &v
	return s
}

type UpgradePrePayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradePrePayOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradePrePayOrderResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderResponseBody) SetCode(v int32) *UpgradePrePayOrderResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradePrePayOrderResponseBody) SetMessage(v string) *UpgradePrePayOrderResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradePrePayOrderResponseBody) SetRequestId(v string) *UpgradePrePayOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradePrePayOrderResponseBody) SetSuccess(v bool) *UpgradePrePayOrderResponseBody {
	s.Success = &v
	return s
}

type UpgradePrePayOrderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradePrePayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradePrePayOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradePrePayOrderResponse) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderResponse) SetHeaders(v map[string]*string) *UpgradePrePayOrderResponse {
	s.Headers = v
	return s
}

func (s *UpgradePrePayOrderResponse) SetStatusCode(v int32) *UpgradePrePayOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradePrePayOrderResponse) SetBody(v *UpgradePrePayOrderResponseBody) *UpgradePrePayOrderResponse {
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
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("alikafka"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConvertPostPayOrderWithOptions(request *ConvertPostPayOrderRequest, runtime *util.RuntimeOptions) (_result *ConvertPostPayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConvertPostPayOrder"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConvertPostPayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConvertPostPayOrder(request *ConvertPostPayOrderRequest) (_result *ConvertPostPayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConvertPostPayOrderResponse{}
	_body, _err := client.ConvertPostPayOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAclWithOptions(request *CreateAclRequest, runtime *util.RuntimeOptions) (_result *CreateAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclOperationType)) {
		query["AclOperationType"] = request.AclOperationType
	}

	if !tea.BoolValue(util.IsUnset(request.AclResourceName)) {
		query["AclResourceName"] = request.AclResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.AclResourcePatternType)) {
		query["AclResourcePatternType"] = request.AclResourcePatternType
	}

	if !tea.BoolValue(util.IsUnset(request.AclResourceType)) {
		query["AclResourceType"] = request.AclResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAcl"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAcl(request *CreateAclRequest) (_result *CreateAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAclResponse{}
	_body, _err := client.CreateAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConsumerGroupWithOptions(request *CreateConsumerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerId)) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConsumerGroup"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing method and pricing of pay-as-you-go Message Queue for Apache Kafka instances. For more information, see [Billing](~~84737~~).
 *
 * @param request CreatePostPayOrderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreatePostPayOrderResponse
 */
func (client *Client) CreatePostPayOrderWithOptions(request *CreatePostPayOrderRequest, runtime *util.RuntimeOptions) (_result *CreatePostPayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		query["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.EipMax)) {
		query["EipMax"] = request.EipMax
	}

	if !tea.BoolValue(util.IsUnset(request.IoMax)) {
		query["IoMax"] = request.IoMax
	}

	if !tea.BoolValue(util.IsUnset(request.IoMaxSpec)) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNum)) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecType)) {
		query["SpecType"] = request.SpecType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TopicQuota)) {
		query["TopicQuota"] = request.TopicQuota
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePostPayOrder"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePostPayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing method and pricing of pay-as-you-go Message Queue for Apache Kafka instances. For more information, see [Billing](~~84737~~).
 *
 * @param request CreatePostPayOrderRequest
 * @return CreatePostPayOrderResponse
 */
func (client *Client) CreatePostPayOrder(request *CreatePostPayOrderRequest) (_result *CreatePostPayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePostPayOrderResponse{}
	_body, _err := client.CreatePostPayOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Before you call this operation, make sure that you understand the billing methods and pricing of subscription ApsaraMQ for Kafka instances. For more information, see [Billing](~~84737~~).
 * *   If you create an ApsaraMQ for Kafka instance by calling this operation, the subscription duration is one month and the auto-renewal feature is enabled by default. The auto-renewal cycle is also one month. If you want to change the auto-renewal cycle or disable the auto-renewal feature, you can go to the [Renewal](https://renew.console.aliyun.com/#/ecs) page in the Alibaba Cloud Management Console.
 *
 * @param request CreatePrePayOrderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreatePrePayOrderResponse
 */
func (client *Client) CreatePrePayOrderWithOptions(request *CreatePrePayOrderRequest, runtime *util.RuntimeOptions) (_result *CreatePrePayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		query["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.EipMax)) {
		query["EipMax"] = request.EipMax
	}

	if !tea.BoolValue(util.IsUnset(request.IoMax)) {
		query["IoMax"] = request.IoMax
	}

	if !tea.BoolValue(util.IsUnset(request.IoMaxSpec)) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNum)) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecType)) {
		query["SpecType"] = request.SpecType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TopicQuota)) {
		query["TopicQuota"] = request.TopicQuota
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePrePayOrder"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePrePayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Before you call this operation, make sure that you understand the billing methods and pricing of subscription ApsaraMQ for Kafka instances. For more information, see [Billing](~~84737~~).
 * *   If you create an ApsaraMQ for Kafka instance by calling this operation, the subscription duration is one month and the auto-renewal feature is enabled by default. The auto-renewal cycle is also one month. If you want to change the auto-renewal cycle or disable the auto-renewal feature, you can go to the [Renewal](https://renew.console.aliyun.com/#/ecs) page in the Alibaba Cloud Management Console.
 *
 * @param request CreatePrePayOrderRequest
 * @return CreatePrePayOrderResponse
 */
func (client *Client) CreatePrePayOrder(request *CreatePrePayOrderRequest) (_result *CreatePrePayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePrePayOrderResponse{}
	_body, _err := client.CreatePrePayOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSaslUserWithOptions(request *CreateSaslUserRequest, runtime *util.RuntimeOptions) (_result *CreateSaslUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSaslUser"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSaslUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSaslUser(request *CreateSaslUserRequest) (_result *CreateSaslUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSaslUserResponse{}
	_body, _err := client.CreateSaslUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Each Alibaba Cloud account can call this operation up to once per second.
 * *   The maximum number of topics that you can create in an instance is determined by the specification of the instance.
 *
 * @param request CreateTopicRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateTopicResponse
 */
func (client *Client) CreateTopicWithOptions(request *CreateTopicRequest, runtime *util.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompactTopic)) {
		query["CompactTopic"] = request.CompactTopic
	}

	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LocalTopic)) {
		query["LocalTopic"] = request.LocalTopic
	}

	if !tea.BoolValue(util.IsUnset(request.MinInsyncReplicas)) {
		query["MinInsyncReplicas"] = request.MinInsyncReplicas
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNum)) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.ReplicationFactor)) {
		query["ReplicationFactor"] = request.ReplicationFactor
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTopic"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * *   Each Alibaba Cloud account can call this operation up to once per second.
 * *   The maximum number of topics that you can create in an instance is determined by the specification of the instance.
 *
 * @param request CreateTopicRequest
 * @return CreateTopicResponse
 */
func (client *Client) CreateTopic(request *CreateTopicRequest) (_result *CreateTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTopicResponse{}
	_body, _err := client.CreateTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAclWithOptions(request *DeleteAclRequest, runtime *util.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclOperationType)) {
		query["AclOperationType"] = request.AclOperationType
	}

	if !tea.BoolValue(util.IsUnset(request.AclResourceName)) {
		query["AclResourceName"] = request.AclResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.AclResourcePatternType)) {
		query["AclResourcePatternType"] = request.AclResourcePatternType
	}

	if !tea.BoolValue(util.IsUnset(request.AclResourceType)) {
		query["AclResourceType"] = request.AclResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAcl"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAcl(request *DeleteAclRequest) (_result *DeleteAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAclResponse{}
	_body, _err := client.DeleteAclWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConsumerGroupWithOptions(request *DeleteConsumerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerId)) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConsumerGroup"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSaslUserWithOptions(request *DeleteSaslUserRequest, runtime *util.RuntimeOptions) (_result *DeleteSaslUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSaslUser"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSaslUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSaslUser(request *DeleteSaslUserRequest) (_result *DeleteSaslUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSaslUserResponse{}
	_body, _err := client.DeleteSaslUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTopicWithOptions(request *DeleteTopicRequest, runtime *util.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTopic"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) DeleteTopic(request *DeleteTopicRequest) (_result *DeleteTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTopicResponse{}
	_body, _err := client.DeleteTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAclsWithOptions(request *DescribeAclsRequest, runtime *util.RuntimeOptions) (_result *DescribeAclsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclResourceName)) {
		query["AclResourceName"] = request.AclResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.AclResourcePatternType)) {
		query["AclResourcePatternType"] = request.AclResourcePatternType
	}

	if !tea.BoolValue(util.IsUnset(request.AclResourceType)) {
		query["AclResourceType"] = request.AclResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAcls"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAclsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAcls(request *DescribeAclsRequest) (_result *DescribeAclsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAclsResponse{}
	_body, _err := client.DescribeAclsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSaslUsersWithOptions(request *DescribeSaslUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeSaslUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSaslUsers"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSaslUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSaslUsers(request *DescribeSaslUsersRequest) (_result *DescribeSaslUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSaslUsersResponse{}
	_body, _err := client.DescribeSaslUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllInstanceIdListWithOptions(request *GetAllInstanceIdListRequest, runtime *util.RuntimeOptions) (_result *GetAllInstanceIdListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAllInstanceIdList"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllInstanceIdListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllInstanceIdList(request *GetAllInstanceIdListRequest) (_result *GetAllInstanceIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllInstanceIdListResponse{}
	_body, _err := client.GetAllInstanceIdListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllowedIpListWithOptions(request *GetAllowedIpListRequest, runtime *util.RuntimeOptions) (_result *GetAllowedIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAllowedIpList"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllowedIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllowedIpList(request *GetAllowedIpListRequest) (_result *GetAllowedIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllowedIpListResponse{}
	_body, _err := client.GetAllowedIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConsumerListWithOptions(request *GetConsumerListRequest, runtime *util.RuntimeOptions) (_result *GetConsumerListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerId)) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsumerList"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConsumerListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConsumerList(request *GetConsumerListRequest) (_result *GetConsumerListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConsumerListResponse{}
	_body, _err := client.GetConsumerListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConsumerProgressWithOptions(request *GetConsumerProgressRequest, runtime *util.RuntimeOptions) (_result *GetConsumerProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerId)) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConsumerProgress"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConsumerProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConsumerProgress(request *GetConsumerProgressRequest) (_result *GetConsumerProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConsumerProgressResponse{}
	_body, _err := client.GetConsumerProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceListWithOptions(request *GetInstanceListRequest, runtime *util.RuntimeOptions) (_result *GetInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceList"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceList(request *GetInstanceListRequest) (_result *GetInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceListResponse{}
	_body, _err := client.GetInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuotaTipWithOptions(request *GetQuotaTipRequest, runtime *util.RuntimeOptions) (_result *GetQuotaTipResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaTip"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaTipResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuotaTip(request *GetQuotaTipRequest) (_result *GetQuotaTipResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQuotaTipResponse{}
	_body, _err := client.GetQuotaTipWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTopicListWithOptions(request *GetTopicListRequest, runtime *util.RuntimeOptions) (_result *GetTopicListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopicList"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTopicList(request *GetTopicListRequest) (_result *GetTopicListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicListResponse{}
	_body, _err := client.GetTopicListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTopicStatusWithOptions(request *GetTopicStatusRequest, runtime *util.RuntimeOptions) (_result *GetTopicStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopicStatus"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTopicStatus(request *GetTopicStatusRequest) (_result *GetTopicStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicStatusResponse{}
	_body, _err := client.GetTopicStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTopicSubscribeStatusWithOptions(request *GetTopicSubscribeStatusRequest, runtime *util.RuntimeOptions) (_result *GetTopicSubscribeStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopicSubscribeStatus"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicSubscribeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTopicSubscribeStatus(request *GetTopicSubscribeStatusRequest) (_result *GetTopicSubscribeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicSubscribeStatusResponse{}
	_body, _err := client.GetTopicSubscribeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceNameWithOptions(request *ModifyInstanceNameRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceName"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (_result *ModifyInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.ModifyInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPartitionNumWithOptions(request *ModifyPartitionNumRequest, runtime *util.RuntimeOptions) (_result *ModifyPartitionNumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddPartitionNum)) {
		query["AddPartitionNum"] = request.AddPartitionNum
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPartitionNum"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPartitionNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPartitionNum(request *ModifyPartitionNumRequest) (_result *ModifyPartitionNumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPartitionNumResponse{}
	_body, _err := client.ModifyPartitionNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTopicRemarkWithOptions(request *ModifyTopicRemarkRequest, runtime *util.RuntimeOptions) (_result *ModifyTopicRemarkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTopicRemark"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTopicRemarkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTopicRemark(request *ModifyTopicRemarkRequest) (_result *ModifyTopicRemarkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTopicRemarkResponse{}
	_body, _err := client.ModifyTopicRemarkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryMessageWithOptions(request *QueryMessageRequest, runtime *util.RuntimeOptions) (_result *QueryMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryMessage"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryMessage(request *QueryMessageRequest) (_result *QueryMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryMessageResponse{}
	_body, _err := client.QueryMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You cannot call this operation to release a subscription Message Queue for Apache Kafka instance.
 *
 * @param request ReleaseInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReleaseInstanceResponse
 */
func (client *Client) ReleaseInstanceWithOptions(request *ReleaseInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceDeleteInstance)) {
		query["ForceDeleteInstance"] = request.ForceDeleteInstance
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstance"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You cannot call this operation to release a subscription Message Queue for Apache Kafka instance.
 *
 * @param request ReleaseInstanceRequest
 * @return ReleaseInstanceResponse
 */
func (client *Client) ReleaseInstance(request *ReleaseInstanceRequest) (_result *ReleaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstanceResponse{}
	_body, _err := client.ReleaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  You can call this operation up to twice per second.
 *
 * @param request StartInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartInstanceResponse
 */
func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.CrossZone)) {
		query["CrossZone"] = request.CrossZone
	}

	if !tea.BoolValue(util.IsUnset(request.DeployModule)) {
		query["DeployModule"] = request.DeployModule
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsEipInner)) {
		query["IsEipInner"] = request.IsEipInner
	}

	if !tea.BoolValue(util.IsUnset(request.IsForceSelectedZones)) {
		query["IsForceSelectedZones"] = request.IsForceSelectedZones
	}

	if !tea.BoolValue(util.IsUnset(request.IsSetUserAndPassword)) {
		query["IsSetUserAndPassword"] = request.IsSetUserAndPassword
	}

	if !tea.BoolValue(util.IsUnset(request.KMSKeyId)) {
		query["KMSKeyId"] = request.KMSKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Notifier)) {
		query["Notifier"] = request.Notifier
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroup)) {
		query["SecurityGroup"] = request.SecurityGroup
	}

	if !tea.BoolValue(util.IsUnset(request.SelectedZones)) {
		query["SelectedZones"] = request.SelectedZones
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.UserPhoneNum)) {
		query["UserPhoneNum"] = request.UserPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstance"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  You can call this operation up to twice per second.
 *
 * @param request StartInstanceRequest
 * @return StartInstanceResponse
 */
func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAllowedIpWithOptions(request *UpdateAllowedIpRequest, runtime *util.RuntimeOptions) (_result *UpdateAllowedIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowedListIp)) {
		query["AllowedListIp"] = request.AllowedListIp
	}

	if !tea.BoolValue(util.IsUnset(request.AllowedListType)) {
		query["AllowedListType"] = request.AllowedListType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PortRange)) {
		query["PortRange"] = request.PortRange
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateType)) {
		query["UpdateType"] = request.UpdateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAllowedIp"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAllowedIpResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAllowedIp(request *UpdateAllowedIpRequest) (_result *UpdateAllowedIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAllowedIpResponse{}
	_body, _err := client.UpdateAllowedIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to reset the consumer offsets of the subscribed topics of a consumer group. You can specify a timestamp or an offset to reset a consumer offset. You can implement the following features by configuring a combination of different parameters:
 * *   Reset the consumer offsets of one or all subscribed topics of a consumer group to the latest offset. This way, you can consume messages in the topics from the latest offset.
 * *   Reset the consumer offsets of one or all subscribed topics of a consumer group to a specific point in time. This way, you can consume messages in the topics from the specified point in time.
 * *   Reset the consumer offset of one subscribed topic of a consumer group to a specific offset in a specific partition. This way, you can consume messages from the specified offset in the specified partition.
 *
 * @param tmpReq UpdateConsumerOffsetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateConsumerOffsetResponse
 */
func (client *Client) UpdateConsumerOffsetWithOptions(tmpReq *UpdateConsumerOffsetRequest, runtime *util.RuntimeOptions) (_result *UpdateConsumerOffsetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateConsumerOffsetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Offsets)) {
		request.OffsetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Offsets, tea.String("Offsets"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerId)) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OffsetsShrink)) {
		query["Offsets"] = request.OffsetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResetType)) {
		query["ResetType"] = request.ResetType
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		query["Time"] = request.Time
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConsumerOffset"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConsumerOffsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to reset the consumer offsets of the subscribed topics of a consumer group. You can specify a timestamp or an offset to reset a consumer offset. You can implement the following features by configuring a combination of different parameters:
 * *   Reset the consumer offsets of one or all subscribed topics of a consumer group to the latest offset. This way, you can consume messages in the topics from the latest offset.
 * *   Reset the consumer offsets of one or all subscribed topics of a consumer group to a specific point in time. This way, you can consume messages in the topics from the specified point in time.
 * *   Reset the consumer offset of one subscribed topic of a consumer group to a specific offset in a specific partition. This way, you can consume messages from the specified offset in the specified partition.
 *
 * @param request UpdateConsumerOffsetRequest
 * @return UpdateConsumerOffsetResponse
 */
func (client *Client) UpdateConsumerOffset(request *UpdateConsumerOffsetRequest) (_result *UpdateConsumerOffsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConsumerOffsetResponse{}
	_body, _err := client.UpdateConsumerOffsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## **Permissions**
 * If a RAM user wants to call the **UpdateInstanceConfig** operation, the RAM user must be granted the required permissions. For more information about how to grant permissions, see [RAM policies](~~185815~~).
 * |API|Action|Resource|
 * |---|---|---|
 * |UpdateInstanceConfig|alikafka: UpdateInstance|acs:alikafka:*:*:{instanceId}|
 *
 * @param request UpdateInstanceConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateInstanceConfigResponse
 */
func (client *Client) UpdateInstanceConfigWithOptions(request *UpdateInstanceConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceConfig"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## **Permissions**
 * If a RAM user wants to call the **UpdateInstanceConfig** operation, the RAM user must be granted the required permissions. For more information about how to grant permissions, see [RAM policies](~~185815~~).
 * |API|Action|Resource|
 * |---|---|---|
 * |UpdateInstanceConfig|alikafka: UpdateInstance|acs:alikafka:*:*:{instanceId}|
 *
 * @param request UpdateInstanceConfigRequest
 * @return UpdateInstanceConfigResponse
 */
func (client *Client) UpdateInstanceConfig(request *UpdateInstanceConfigRequest) (_result *UpdateInstanceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceConfigResponse{}
	_body, _err := client.UpdateInstanceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTopicConfigWithOptions(request *UpdateTopicConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateTopicConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTopicConfig"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTopicConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTopicConfig(request *UpdateTopicConfigRequest) (_result *UpdateTopicConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTopicConfigResponse{}
	_body, _err := client.UpdateTopicConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## **Permissions**
 * A RAM user must be granted the required permissions before the RAM user calls the **UpgradeInstanceVersion** operation. For information about how to grant permissions, see [RAM policies](~~185815~~).
 * |API|Action|Resource|
 * |---|---|---|
 * |UpgradeInstanceVersion|UpdateInstance|acs:alikafka:*:*:{instanceId}|
 * ## **QPS limits**
 * You can send a maximum of two queries per second (QPS).
 *
 * @param request UpgradeInstanceVersionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpgradeInstanceVersionResponse
 */
func (client *Client) UpgradeInstanceVersionWithOptions(request *UpgradeInstanceVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetVersion)) {
		query["TargetVersion"] = request.TargetVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeInstanceVersion"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## **Permissions**
 * A RAM user must be granted the required permissions before the RAM user calls the **UpgradeInstanceVersion** operation. For information about how to grant permissions, see [RAM policies](~~185815~~).
 * |API|Action|Resource|
 * |---|---|---|
 * |UpgradeInstanceVersion|UpdateInstance|acs:alikafka:*:*:{instanceId}|
 * ## **QPS limits**
 * You can send a maximum of two queries per second (QPS).
 *
 * @param request UpgradeInstanceVersionRequest
 * @return UpgradeInstanceVersionResponse
 */
func (client *Client) UpgradeInstanceVersion(request *UpgradeInstanceVersionRequest) (_result *UpgradeInstanceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceVersionResponse{}
	_body, _err := client.UpgradeInstanceVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing method and pricing of pay-as-you-go Message Queue for Apache Kafka instances. For more information, see [Billing](~~84737~~).
 *
 * @param tmpReq UpgradePostPayOrderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpgradePostPayOrderResponse
 */
func (client *Client) UpgradePostPayOrderWithOptions(tmpReq *UpgradePostPayOrderRequest, runtime *util.RuntimeOptions) (_result *UpgradePostPayOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpgradePostPayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ServerlessConfig)) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, tea.String("ServerlessConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		query["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.EipMax)) {
		query["EipMax"] = request.EipMax
	}

	if !tea.BoolValue(util.IsUnset(request.EipModel)) {
		query["EipModel"] = request.EipModel
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IoMax)) {
		query["IoMax"] = request.IoMax
	}

	if !tea.BoolValue(util.IsUnset(request.IoMaxSpec)) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNum)) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerlessConfigShrink)) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SpecType)) {
		query["SpecType"] = request.SpecType
	}

	if !tea.BoolValue(util.IsUnset(request.TopicQuota)) {
		query["TopicQuota"] = request.TopicQuota
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradePostPayOrder"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradePostPayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing method and pricing of pay-as-you-go Message Queue for Apache Kafka instances. For more information, see [Billing](~~84737~~).
 *
 * @param request UpgradePostPayOrderRequest
 * @return UpgradePostPayOrderResponse
 */
func (client *Client) UpgradePostPayOrder(request *UpgradePostPayOrderRequest) (_result *UpgradePostPayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradePostPayOrderResponse{}
	_body, _err := client.UpgradePostPayOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing method and pricing of subscription Message Queue for Apache Kafka instances. For more information, see [Billing overview](~~84737~~).
 *
 * @param request UpgradePrePayOrderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpgradePrePayOrderResponse
 */
func (client *Client) UpgradePrePayOrderWithOptions(request *UpgradePrePayOrderRequest, runtime *util.RuntimeOptions) (_result *UpgradePrePayOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		query["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.EipMax)) {
		query["EipMax"] = request.EipMax
	}

	if !tea.BoolValue(util.IsUnset(request.EipModel)) {
		query["EipModel"] = request.EipModel
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IoMax)) {
		query["IoMax"] = request.IoMax
	}

	if !tea.BoolValue(util.IsUnset(request.IoMaxSpec)) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNum)) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecType)) {
		query["SpecType"] = request.SpecType
	}

	if !tea.BoolValue(util.IsUnset(request.TopicQuota)) {
		query["TopicQuota"] = request.TopicQuota
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradePrePayOrder"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradePrePayOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call this operation, make sure that you understand the billing method and pricing of subscription Message Queue for Apache Kafka instances. For more information, see [Billing overview](~~84737~~).
 *
 * @param request UpgradePrePayOrderRequest
 * @return UpgradePrePayOrderResponse
 */
func (client *Client) UpgradePrePayOrder(request *UpgradePrePayOrderRequest) (_result *UpgradePrePayOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradePrePayOrderResponse{}
	_body, _err := client.UpgradePrePayOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
