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
	// The ID of the resource group to which you want to transfer the cloud resource.
	//
	// >  You can use resource groups to manage resources owned by your Alibaba Cloud account. Resource groups simplify the resource and permission management of your Alibaba Cloud account. For more information, see [What is resource management?](https://help.aliyun.com/document_detail/94475.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-ac***********7q
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource to which you want to attach a tag. Only the ID of a Message Queue for Apache Kafka instance is supported.
	//
	// For example, if the ID of the instance is alikafka_post-cn-v0h1fgs2xxxx, the resource ID is alikafka_post-cn-v0h1fgs2xxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the new resource group. You can view the available resource groups in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C5CA600C-7D5A-45B5-B6DB-44FAC2C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// 	- **1~12**
	//
	// 	- **24**
	//
	// 	- **36**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PaidType   *int32  `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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

func (s *ConvertPostPayOrderRequest) SetPaidType(v int32) *ConvertPostPayOrderRequest {
	s.PaidType = &v
	return s
}

func (s *ConvertPostPayOrderRequest) SetRegionId(v string) *ConvertPostPayOrderRequest {
	s.RegionId = &v
	return s
}

type ConvertPostPayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20497346575****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// The type of the operation allowed by the access control list (ACL). Valid values:
	//
	// 	- **Write**
	//
	// 	- **Read**
	//
	// 	- **Describe**: reads of transactional IDs.
	//
	// 	- **IdempotentWrite**: idempotent data writes to clusters.
	//
	// 	- **IDEMPOTENT_WRITE**: idempotent data writes to clusters. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// 	- **DESCRIBE_CONFIGS**: configuration query. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// Read
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The types of operations allowed by the ACL. Separate multiple operation types with commas (,).
	//
	// Valid values:
	//
	// 	- **Write**
	//
	// 	- **Read**
	//
	// 	- **Describe**: reads of transactional IDs.
	//
	// 	- **IdempotentWrite**: idempotent data writes to clusters.
	//
	// 	- **IDEMPOTENT_WRITE**: idempotent data writes to clusters. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// 	- **DESCRIBE_CONFIGS**: configuration query. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// Write,Read
	AclOperationTypes *string `json:"AclOperationTypes,omitempty" xml:"AclOperationTypes,omitempty"`
	// The authorization method. Valid values:
	//
	// 	- **DENY**
	//
	// 	- **ALLOW**
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// The resource name.
	//
	// 	- The value can be a topic name, a group ID, a cluster name, or a transaction ID.
	//
	// 	- You can use an asterisk (\\*) to specify the names of all resources of the specified type.
	//
	// > You can use an asterisk (\\*) to query the resources on which permissions are granted only after you grant the user the required permissions on all resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// X****
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The matching mode. Valid values:
	//
	// 	- **LITERAL**: exact match
	//
	// 	- **PREFIXED**: prefix match
	//
	// This parameter is required.
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **Topic**
	//
	// 	- **Group**
	//
	// 	- **Cluster**
	//
	// 	- **TransactionalId**: transactional ID
	//
	// This parameter is required.
	//
	// example:
	//
	// Group
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The IP address of the source.
	//
	// >
	//
	// 	- You can specify a specific IP address or use the asterisk (\\*) wildcard character to specify all IP addresses. CIDR blocks are not supported.
	//
	// 	- This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// *
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The username.
	//
	// 	- You can use an asterisk (\\*) to specify all usernames.
	//
	// > You can use an asterisk (\\*) to query the authorized users only after you grant the required permissions to all users.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***
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

func (s *CreateAclRequest) SetAclOperationTypes(v string) *CreateAclRequest {
	s.AclOperationTypes = &v
	return s
}

func (s *CreateAclRequest) SetAclPermissionType(v string) *CreateAclRequest {
	s.AclPermissionType = &v
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

func (s *CreateAclRequest) SetHost(v string) *CreateAclRequest {
	s.Host = &v
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 56729737-C428-4E1B-AC68-7A8C2D5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// 	- The value can contain only letters, digits, hyphens (-), and underscores (_), and the value must contain at least one letter or digit.
	//
	// 	- The value must be 3 to 128 characters in length. If the value that you specify contains more than 128 characters, the system automatically truncates the value to 128 characters.
	//
	// 	- After a consumer group is created, you cannot change the name of the consumer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the consumer group.
	//
	// example:
	//
	// test
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
	// 	- You must specify this parameter.
	//
	// 	- The tag key can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// 	- You can leave this parameter empty.
	//
	// 	- The tag value can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E57A8862-DF68-4055-8E55-B80CB4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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

type CreatePostPayInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 1500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// example:
	//
	// 0
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// example:
	//
	// 100
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId  *string                                       `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerlessConfig *CreatePostPayInstanceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// example:
	//
	// professional
	SpecType *string                            `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Tag      []*CreatePostPayInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePostPayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequest) SetDeployType(v int32) *CreatePostPayInstanceRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetDiskSize(v int32) *CreatePostPayInstanceRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetDiskType(v string) *CreatePostPayInstanceRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetEipMax(v int32) *CreatePostPayInstanceRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetIoMaxSpec(v string) *CreatePostPayInstanceRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetPaidType(v int32) *CreatePostPayInstanceRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetPartitionNum(v int32) *CreatePostPayInstanceRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetRegionId(v string) *CreatePostPayInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetResourceGroupId(v string) *CreatePostPayInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetServerlessConfig(v *CreatePostPayInstanceRequestServerlessConfig) *CreatePostPayInstanceRequest {
	s.ServerlessConfig = v
	return s
}

func (s *CreatePostPayInstanceRequest) SetSpecType(v string) *CreatePostPayInstanceRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetTag(v []*CreatePostPayInstanceRequestTag) *CreatePostPayInstanceRequest {
	s.Tag = v
	return s
}

type CreatePostPayInstanceRequestServerlessConfig struct {
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// example:
	//
	// 50
	ReservedSubscribeCapacity *int64 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
}

func (s CreatePostPayInstanceRequestServerlessConfig) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayInstanceRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequestServerlessConfig) SetReservedPublishCapacity(v int64) *CreatePostPayInstanceRequestServerlessConfig {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *CreatePostPayInstanceRequestServerlessConfig) SetReservedSubscribeCapacity(v int64) *CreatePostPayInstanceRequestServerlessConfig {
	s.ReservedSubscribeCapacity = &v
	return s
}

type CreatePostPayInstanceRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePostPayInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequestTag) SetKey(v string) *CreatePostPayInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayInstanceRequestTag) SetValue(v string) *CreatePostPayInstanceRequestTag {
	s.Value = &v
	return s
}

type CreatePostPayInstanceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 1500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// example:
	//
	// 0
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// example:
	//
	// 100
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId        *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// example:
	//
	// professional
	SpecType *string                                  `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Tag      []*CreatePostPayInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePostPayInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceShrinkRequest) SetDeployType(v int32) *CreatePostPayInstanceShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetDiskSize(v int32) *CreatePostPayInstanceShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetDiskType(v string) *CreatePostPayInstanceShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetEipMax(v int32) *CreatePostPayInstanceShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetIoMaxSpec(v string) *CreatePostPayInstanceShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetPaidType(v int32) *CreatePostPayInstanceShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetPartitionNum(v int32) *CreatePostPayInstanceShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetRegionId(v string) *CreatePostPayInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetResourceGroupId(v string) *CreatePostPayInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetServerlessConfigShrink(v string) *CreatePostPayInstanceShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetSpecType(v string) *CreatePostPayInstanceShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequest) SetTag(v []*CreatePostPayInstanceShrinkRequestTag) *CreatePostPayInstanceShrinkRequest {
	s.Tag = v
	return s
}

type CreatePostPayInstanceShrinkRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePostPayInstanceShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceShrinkRequestTag) SetKey(v string) *CreatePostPayInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayInstanceShrinkRequestTag) SetValue(v string) *CreatePostPayInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

type CreatePostPayInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePostPayInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePostPayInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceResponseBody) SetCode(v int32) *CreatePostPayInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePostPayInstanceResponseBody) SetData(v *CreatePostPayInstanceResponseBodyData) *CreatePostPayInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreatePostPayInstanceResponseBody) SetMessage(v string) *CreatePostPayInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePostPayInstanceResponseBody) SetRequestId(v string) *CreatePostPayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePostPayInstanceResponseBody) SetSuccess(v bool) *CreatePostPayInstanceResponseBody {
	s.Success = &v
	return s
}

type CreatePostPayInstanceResponseBodyData struct {
	// example:
	//
	// alikafka_pre-cn-pe333xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 236972661580636
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreatePostPayInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceResponseBodyData) SetInstanceId(v string) *CreatePostPayInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreatePostPayInstanceResponseBodyData) SetOrderId(v int64) *CreatePostPayInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

type CreatePostPayInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePostPayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePostPayInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceResponse) SetHeaders(v map[string]*string) *CreatePostPayInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePostPayInstanceResponse) SetStatusCode(v int32) *CreatePostPayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePostPayInstanceResponse) SetBody(v *CreatePostPayInstanceResponseBody) *CreatePostPayInstanceResponse {
	s.Body = v
	return s
}

type CreatePostPayOrderRequest struct {
	// The deployment mode of the instance. Valid values:
	//
	// 	- **4**: deploys the instance that allows access from the Internet and a VPC.
	//
	// 	- **5**: deploys the instance that allows access only from a VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size.
	//
	// For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type of the instance. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic.
	//
	// 	- If you set **DeployType*	- to **4**, you must configure this parameter.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic in the instance. We recommend that you do not configure this parameter.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- 1: pay-as-you-go (reserved capacity).
	//
	// 	- 3: pay-as-you-go (reserved capacity) + pay-as-you-go (on-demand capacity)
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only ParittionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// If this parameter is left empty, the default resource group is used. You can view the resource group ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The parameters configured for the serverless ApsaraMQ for Kafka instance. These parameters are required only when you create a serverless instance.
	ServerlessConfig *CreatePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The instance edition.
	//
	// Valid values if you set PaidType to 1:
	//
	// 	- normal: Standard Edition (High Write)
	//
	// 	- professional: Professional Edition (High Write)
	//
	// 	- professionalForHighRead: Professional Edition (High Read)
	//
	// Valid values if you set PaidType to 3:
	//
	// 	- normal: Serverless Standard Edition
	//
	// For more information about the instance editions, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePostPayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only ParittionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
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

func (s *CreatePostPayOrderRequest) SetPaidType(v int32) *CreatePostPayOrderRequest {
	s.PaidType = &v
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

func (s *CreatePostPayOrderRequest) SetServerlessConfig(v *CreatePostPayOrderRequestServerlessConfig) *CreatePostPayOrderRequest {
	s.ServerlessConfig = v
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

type CreatePostPayOrderRequestServerlessConfig struct {
	// The reserved capacity for publishing messages. You can specify only an integer for this parameter. Minimum value: 60.
	//
	// >  The actual maximum reserved capacity for publishing messages varies based on available resources in the region. The actual range displayed on the buy page shall prevail.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved capacity for subscribing to messages. You can specify only an integer for this parameter. Minimum value: 20.
	//
	// >  The actual maximum reserved capacity for subscribing to messages varies based on available resources in the region. The actual range displayed on the buy page shall prevail.
	//
	// example:
	//
	// 50
	ReservedSubscribeCapacity *int64 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
}

func (s CreatePostPayOrderRequestServerlessConfig) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayOrderRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderRequestServerlessConfig) SetReservedPublishCapacity(v int64) *CreatePostPayOrderRequestServerlessConfig {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *CreatePostPayOrderRequestServerlessConfig) SetReservedSubscribeCapacity(v int64) *CreatePostPayOrderRequestServerlessConfig {
	s.ReservedSubscribeCapacity = &v
	return s
}

type CreatePostPayOrderRequestTag struct {
	// The key of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If this parameter is left empty, the keys of all tags are matched.
	//
	// 	- The tag key must be up to 128 characters in length. It cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If you do not specify a tag key, you cannot specify a tag value. If this parameter is not configured, all tag values are matched.
	//
	// 	- The tag value must be 1 to 128 characters in length. It cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// example:
	//
	// FinanceJoshua
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

type CreatePostPayOrderShrinkRequest struct {
	// The deployment mode of the instance. Valid values:
	//
	// 	- **4**: deploys the instance that allows access from the Internet and a VPC.
	//
	// 	- **5**: deploys the instance that allows access only from a VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size.
	//
	// For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type of the instance. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic.
	//
	// 	- If you set **DeployType*	- to **4**, you must configure this parameter.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic in the instance. We recommend that you do not configure this parameter.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- 1: pay-as-you-go (reserved capacity).
	//
	// 	- 3: pay-as-you-go (reserved capacity) + pay-as-you-go (on-demand capacity)
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only ParittionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// If this parameter is left empty, the default resource group is used. You can view the resource group ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The parameters configured for the serverless ApsaraMQ for Kafka instance. These parameters are required only when you create a serverless instance.
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// The instance edition.
	//
	// Valid values if you set PaidType to 1:
	//
	// 	- normal: Standard Edition (High Write)
	//
	// 	- professional: Professional Edition (High Write)
	//
	// 	- professionalForHighRead: Professional Edition (High Read)
	//
	// Valid values if you set PaidType to 3:
	//
	// 	- normal: Serverless Standard Edition
	//
	// For more information about the instance editions, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePostPayOrderShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only ParittionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s CreatePostPayOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderShrinkRequest) SetDeployType(v int32) *CreatePostPayOrderShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetDiskSize(v int32) *CreatePostPayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetDiskType(v string) *CreatePostPayOrderShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetEipMax(v int32) *CreatePostPayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetIoMax(v int32) *CreatePostPayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetIoMaxSpec(v string) *CreatePostPayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetPaidType(v int32) *CreatePostPayOrderShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetPartitionNum(v int32) *CreatePostPayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetRegionId(v string) *CreatePostPayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetResourceGroupId(v string) *CreatePostPayOrderShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetServerlessConfigShrink(v string) *CreatePostPayOrderShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetSpecType(v string) *CreatePostPayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetTag(v []*CreatePostPayOrderShrinkRequestTag) *CreatePostPayOrderShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetTopicQuota(v int32) *CreatePostPayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

type CreatePostPayOrderShrinkRequestTag struct {
	// The key of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If this parameter is left empty, the keys of all tags are matched.
	//
	// 	- The tag key must be up to 128 characters in length. It cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If you do not specify a tag key, you cannot specify a tag value. If this parameter is not configured, all tag values are matched.
	//
	// 	- The tag value must be 1 to 128 characters in length. It cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePostPayOrderShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePostPayOrderShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderShrinkRequestTag) SetKey(v string) *CreatePostPayOrderShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequestTag) SetValue(v string) *CreatePostPayOrderShrinkRequestTag {
	s.Value = &v
	return s
}

type CreatePostPayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20497346575****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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

type CreatePrePayInstanceRequest struct {
	ConfluentConfig *CreatePrePayInstanceRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// 1
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// example:
	//
	// 1000
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// professional
	SpecType *string                           `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Tag      []*CreatePrePayInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePrePayInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceRequest) SetConfluentConfig(v *CreatePrePayInstanceRequestConfluentConfig) *CreatePrePayInstanceRequest {
	s.ConfluentConfig = v
	return s
}

func (s *CreatePrePayInstanceRequest) SetDeployType(v int32) *CreatePrePayInstanceRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetDiskSize(v int32) *CreatePrePayInstanceRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetDiskType(v string) *CreatePrePayInstanceRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetDuration(v int32) *CreatePrePayInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetEipMax(v int32) *CreatePrePayInstanceRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetIoMaxSpec(v string) *CreatePrePayInstanceRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetPaidType(v int32) *CreatePrePayInstanceRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetPartitionNum(v int32) *CreatePrePayInstanceRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetRegionId(v string) *CreatePrePayInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetResourceGroupId(v string) *CreatePrePayInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetSpecType(v string) *CreatePrePayInstanceRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetTag(v []*CreatePrePayInstanceRequestTag) *CreatePrePayInstanceRequest {
	s.Tag = v
	return s
}

type CreatePrePayInstanceRequestConfluentConfig struct {
	// example:
	//
	// 4
	ConnectCU *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	// example:
	//
	// 2
	ConnectReplica *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	// example:
	//
	// 4
	ControlCenterCU *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	// example:
	//
	// 1
	ControlCenterReplica *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	// example:
	//
	// 300
	ControlCenterStorage *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	// example:
	//
	// 4
	KafkaCU *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	// example:
	//
	// 3
	KafkaReplica *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	// example:
	//
	// 4
	KafkaRestProxyCU *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	// example:
	//
	// 2
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	// example:
	//
	// 800
	KafkaStorage *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	// example:
	//
	// 4
	KsqlCU *int32 `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	// example:
	//
	// 2
	KsqlReplica *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	// example:
	//
	// 100
	KsqlStorage *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	// example:
	//
	// 1
	SchemaRegistryCU *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	// example:
	//
	// 2
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	// example:
	//
	// 2
	ZooKeeperCU *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	// example:
	//
	// 3
	ZooKeeperReplica *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	// example:
	//
	// 100
	ZooKeeperStorage *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s CreatePrePayInstanceRequestConfluentConfig) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayInstanceRequestConfluentConfig) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetConnectCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetConnectReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetControlCenterCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetControlCenterReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetControlCenterStorage(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaRestProxyCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaRestProxyReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaStorage(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKsqlCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKsqlReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKsqlStorage(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetSchemaRegistryCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetSchemaRegistryReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetZooKeeperCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetZooKeeperReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetZooKeeperStorage(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

type CreatePrePayInstanceRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrePayInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceRequestTag) SetKey(v string) *CreatePrePayInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrePayInstanceRequestTag) SetValue(v string) *CreatePrePayInstanceRequestTag {
	s.Value = &v
	return s
}

type CreatePrePayInstanceShrinkRequest struct {
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// 1
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// example:
	//
	// 1000
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// professional
	SpecType *string                                 `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Tag      []*CreatePrePayInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePrePayInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceShrinkRequest) SetConfluentConfigShrink(v string) *CreatePrePayInstanceShrinkRequest {
	s.ConfluentConfigShrink = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetDeployType(v int32) *CreatePrePayInstanceShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetDiskSize(v int32) *CreatePrePayInstanceShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetDiskType(v string) *CreatePrePayInstanceShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetDuration(v int32) *CreatePrePayInstanceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetEipMax(v int32) *CreatePrePayInstanceShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetIoMaxSpec(v string) *CreatePrePayInstanceShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetPaidType(v int32) *CreatePrePayInstanceShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetPartitionNum(v int32) *CreatePrePayInstanceShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetRegionId(v string) *CreatePrePayInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetResourceGroupId(v string) *CreatePrePayInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetSpecType(v string) *CreatePrePayInstanceShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetTag(v []*CreatePrePayInstanceShrinkRequestTag) *CreatePrePayInstanceShrinkRequest {
	s.Tag = v
	return s
}

type CreatePrePayInstanceShrinkRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrePayInstanceShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceShrinkRequestTag) SetKey(v string) *CreatePrePayInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequestTag) SetValue(v string) *CreatePrePayInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

type CreatePrePayInstanceResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePrePayInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E57A8862-DF68-4055-8E55-B80CB4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePrePayInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceResponseBody) SetCode(v int32) *CreatePrePayInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePrePayInstanceResponseBody) SetData(v *CreatePrePayInstanceResponseBodyData) *CreatePrePayInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreatePrePayInstanceResponseBody) SetMessage(v string) *CreatePrePayInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePrePayInstanceResponseBody) SetRequestId(v string) *CreatePrePayInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrePayInstanceResponseBody) SetSuccess(v bool) *CreatePrePayInstanceResponseBody {
	s.Success = &v
	return s
}

type CreatePrePayInstanceResponseBodyData struct {
	// example:
	//
	// alikafka_post-cn-xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 236972661xxxx
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreatePrePayInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceResponseBodyData) SetInstanceId(v string) *CreatePrePayInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreatePrePayInstanceResponseBodyData) SetOrderId(v int64) *CreatePrePayInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

type CreatePrePayInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrePayInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrePayInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceResponse) SetHeaders(v map[string]*string) *CreatePrePayInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePrePayInstanceResponse) SetStatusCode(v int32) *CreatePrePayInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrePayInstanceResponse) SetBody(v *CreatePrePayInstanceResponseBody) *CreatePrePayInstanceResponse {
	s.Body = v
	return s
}

type CreatePrePayOrderRequest struct {
	// The configurations of Confluent.
	//
	// >  When you create an ApsaraMQ for Confluent instance, you must configure this parameter.
	ConfluentConfig *CreatePrePayOrderRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	// The type of the network in which the instance is deployed. Valid values:
	//
	// 	- **4**: Internet and virtual private cloud (VPC)
	//
	// 	- **5**: VPC
	//
	// >  If you create an ApsaraMQ for Confluent instance, set the value to 5. After the instance is created, you can specify whether to enable each component.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size. Unit: GB
	//
	// For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The subscription duration. Unit: months. Default value: 1. Valid values:
	//
	// 	- **1 to 12**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The maximum Internet traffic in the instance.
	//
	// 	- If you set **DeployType*	- to **4**, you must configure this parameter.
	//
	// 	- For information about the valid values, see [Pay-as-you-go](https://help.aliyun.com/document_detail/72142.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic in the instance. We recommend that you do not configure this parameter.
	//
	// 	- You must set one of **IoMax*	- and **IoMaxSpec**. If both parameters are configured, the value of **IoMaxSpec*	- is used. We recommend that you configure only **IoMaxSpec**.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- You must configure one of **IoMax*	- and **IoMaxSpec**. If both parameters are configured, the value of **IoMaxSpec*	- is used. We recommend that you configure only **IoMaxSpec**.
	//
	// 	- For more information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **0**: the subscription billing method
	//
	// 	- **4**: the subscription billing method for ApsaraMQ for Confluent instances
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// If this parameter is left empty, the default resource group is used. You can view the resource group ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- **normal**: Standard Edition (High Write)
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePrePayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you use exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s CreatePrePayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderRequest) SetConfluentConfig(v *CreatePrePayOrderRequestConfluentConfig) *CreatePrePayOrderRequest {
	s.ConfluentConfig = v
	return s
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

func (s *CreatePrePayOrderRequest) SetDuration(v int32) *CreatePrePayOrderRequest {
	s.Duration = &v
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

func (s *CreatePrePayOrderRequest) SetPaidType(v int32) *CreatePrePayOrderRequest {
	s.PaidType = &v
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

type CreatePrePayOrderRequestConfluentConfig struct {
	// The number of CPU cores of Connect.
	//
	// example:
	//
	// 4
	ConnectCU *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	// The number of replicas of Connect.
	//
	// example:
	//
	// 2
	ConnectReplica *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	// The number of CPU cores of Control Center.
	//
	// example:
	//
	// 4
	ControlCenterCU *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	// The number of replicas of Control Center.
	//
	// example:
	//
	// 1
	ControlCenterReplica *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	// The disk capacity of Control Center. Unit: GB
	//
	// example:
	//
	// 300
	ControlCenterStorage *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	// The number of CPU cores of the Kafka broker.
	//
	// example:
	//
	// 4
	KafkaCU *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	// The number of replicas of the Kafka broker.
	//
	// example:
	//
	// 3
	KafkaReplica *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	// The number of CPU cores of Kafka Rest Proxy.
	//
	// example:
	//
	// 4
	KafkaRestProxyCU *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	// The number of replicas of Kafka Rest Proxy.
	//
	// example:
	//
	// 2
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	// The disk capacity of the Kafka broker. Unit: GB
	//
	// example:
	//
	// 800
	KafkaStorage *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	// The number of CPU cores of ksqIDB.
	//
	// example:
	//
	// 4
	KsqlCU *int32 `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	// The number of replicas of ksqlDB.
	//
	// example:
	//
	// 2
	KsqlReplica *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	// The disk capacity of ksqlDB. Unit: GB
	//
	// example:
	//
	// 100
	KsqlStorage *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	// The number of CPU cores of Schema Registry.
	//
	// example:
	//
	// 1
	SchemaRegistryCU *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	// The number of replicas of Schema Registry.
	//
	// example:
	//
	// 2
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	// The number of CPU cores of ZooKeeper.
	//
	// example:
	//
	// 2
	ZooKeeperCU *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	// The number of replicas of ZooKeeper.
	//
	// example:
	//
	// 3
	ZooKeeperReplica *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	// The disk capacity of ZooKeeper. Unit: GB
	//
	// example:
	//
	// 100
	ZooKeeperStorage *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s CreatePrePayOrderRequestConfluentConfig) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayOrderRequestConfluentConfig) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetConnectCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetConnectReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetControlCenterCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetControlCenterReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetControlCenterStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaRestProxyCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaRestProxyReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKsqlCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKsqlReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKsqlStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetSchemaRegistryCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetSchemaRegistryReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetZooKeeperCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetZooKeeperReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetZooKeeperStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

type CreatePrePayOrderRequestTag struct {
	// The key of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If this parameter is left empty, the keys of all tags are matched.
	//
	// 	- The tag key can be up to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- This parameter can be left empty.
	//
	// 	- The tag value can be 1 to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// example:
	//
	// FinanceJoshua
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

type CreatePrePayOrderShrinkRequest struct {
	// The configurations of Confluent.
	//
	// >  When you create an ApsaraMQ for Confluent instance, you must configure this parameter.
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	// The type of the network in which the instance is deployed. Valid values:
	//
	// 	- **4**: Internet and virtual private cloud (VPC)
	//
	// 	- **5**: VPC
	//
	// >  If you create an ApsaraMQ for Confluent instance, set the value to 5. After the instance is created, you can specify whether to enable each component.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size. Unit: GB
	//
	// For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The subscription duration. Unit: months. Default value: 1. Valid values:
	//
	// 	- **1 to 12**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The maximum Internet traffic in the instance.
	//
	// 	- If you set **DeployType*	- to **4**, you must configure this parameter.
	//
	// 	- For information about the valid values, see [Pay-as-you-go](https://help.aliyun.com/document_detail/72142.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic in the instance. We recommend that you do not configure this parameter.
	//
	// 	- You must set one of **IoMax*	- and **IoMaxSpec**. If both parameters are configured, the value of **IoMaxSpec*	- is used. We recommend that you configure only **IoMaxSpec**.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- You must configure one of **IoMax*	- and **IoMaxSpec**. If both parameters are configured, the value of **IoMaxSpec*	- is used. We recommend that you configure only **IoMaxSpec**.
	//
	// 	- For more information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **0**: the subscription billing method
	//
	// 	- **4**: the subscription billing method for ApsaraMQ for Confluent instances
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// If this parameter is left empty, the default resource group is used. You can view the resource group ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- **normal**: Standard Edition (High Write)
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePrePayOrderShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you use exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s CreatePrePayOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderShrinkRequest) SetConfluentConfigShrink(v string) *CreatePrePayOrderShrinkRequest {
	s.ConfluentConfigShrink = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetDeployType(v int32) *CreatePrePayOrderShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetDiskSize(v int32) *CreatePrePayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetDiskType(v string) *CreatePrePayOrderShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetDuration(v int32) *CreatePrePayOrderShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetEipMax(v int32) *CreatePrePayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetIoMax(v int32) *CreatePrePayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetIoMaxSpec(v string) *CreatePrePayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetPaidType(v int32) *CreatePrePayOrderShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetPartitionNum(v int32) *CreatePrePayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetRegionId(v string) *CreatePrePayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetResourceGroupId(v string) *CreatePrePayOrderShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetSpecType(v string) *CreatePrePayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetTag(v []*CreatePrePayOrderShrinkRequestTag) *CreatePrePayOrderShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetTopicQuota(v int32) *CreatePrePayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

type CreatePrePayOrderShrinkRequestTag struct {
	// The key of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If this parameter is left empty, the keys of all tags are matched.
	//
	// 	- The tag key can be up to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- This parameter can be left empty.
	//
	// 	- The tag value can be 1 to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrePayOrderShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePrePayOrderShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderShrinkRequestTag) SetKey(v string) *CreatePrePayOrderShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequestTag) SetValue(v string) *CreatePrePayOrderShrinkRequestTag {
	s.Value = &v
	return s
}

type CreatePrePayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 20497346575****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The encryption method. Valid values:
	//
	// 	- SCRAM-SHA-512 (default)
	//
	// 	- SCRAM-SHA-256
	//
	// >  This parameter is available only for ApsaraMQ for Kafka serverless instances.
	//
	// example:
	//
	// SCRAM-SHA-256
	Mechanism *string `json:"Mechanism,omitempty" xml:"Mechanism,omitempty"`
	// The password of the SASL user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the Simple Authentication and Security Layer (SASL) user. Valid values:
	//
	// 	- **plain**: a simple mechanism that uses usernames and passwords to verify user identities. ApsaraMQ for Kafka provides an improved PLAIN mechanism that allows you to dynamically add SASL users without the need to restart an instance.
	//
	// 	- **SCRAM**: a mechanism that uses usernames and passwords to verify user identities. Compared with the PLAIN mechanism, this mechanism provides better security protection. ApsaraMQ for Kafka uses the SCRAM-SHA-256 algorithm.
	//
	// 	- **LDAP**: This value is available only for the SASL users of ApsaraMQ for Confluent instances.
	//
	// Default value: **plain**.
	//
	// example:
	//
	// plain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the SASL user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***
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

func (s *CreateSaslUserRequest) SetMechanism(v string) *CreateSaslUserRequest {
	s.Mechanism = &v
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
	// The HTTP status code. The HTTP status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C5CA600C-7D5A-45B5-B6DB-44FAC2C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
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

type CreateScheduledScalingRuleRequest struct {
	// The duration of each scheduled scaling task. Unit: minutes.
	//
	// >  The value of this parameter must be greater than or equal to 15.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	DurationMinutes *int32 `json:"DurationMinutes,omitempty" xml:"DurationMinutes,omitempty"`
	// Specifies whether to enable the scheduled scaling rule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the scheduled scaling task is executed.
	//
	// If you set ScheduleType to at, make sure that the value of this parameter is at least 30 minutes later than the current point in time.
	//
	// 	Notice: To prevent the broker from repeatedly executing instance upgrade and downgrade tasks, make sure that the interval between two consecutive scheduled scaling tasks is at least 60 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1714467540000
	FirstScheduledTime *int64 `json:"FirstScheduledTime,omitempty" xml:"FirstScheduledTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_serverless-cn-vxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The frequency to execute the scheduled scaling task. This parameter is required only if you set ScheduleType to repeat. Valid values:
	//
	// 	- Daily: The scheduled scaling task is executed every day.
	//
	// 	- Weekly: The scheduled scaling task is executed every week.
	//
	// example:
	//
	// Weekly
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The reserved production capacity for scheduled scaling. Unit: MB/s.
	//
	// >  You must specify a higher value than the instance specification for at least one of ReservedPubFlow and ReservedSubFlow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	ReservedPubFlow *int32 `json:"ReservedPubFlow,omitempty" xml:"ReservedPubFlow,omitempty"`
	// The reserved consumption capacity for scheduled scaling. Unit: MB/s.
	//
	// >  You must specify a higher value than the instance specification for at least one of ReservedPubFlow and ReservedSubFlow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	ReservedSubFlow *int32 `json:"ReservedSubFlow,omitempty" xml:"ReservedSubFlow,omitempty"`
	// The name of the scheduled scaling rule.
	//
	// >  The name of the scheduled scaling rule cannot be the same as the names of other rules for the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the scheduled scaling task. Valid values:
	//
	// 	- at: The scheduled scaling task is executed only once.
	//
	// 	- repeat: The scheduled scaling task is repeatedly executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// at
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The time zone in Coordinated Universal Time (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// GMT+8
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The day on which the scheduled scaling task is executed every week. You can specify multiple days.
	WeeklyTypes []*string `json:"WeeklyTypes,omitempty" xml:"WeeklyTypes,omitempty" type:"Repeated"`
}

func (s CreateScheduledScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledScalingRuleRequest) SetDurationMinutes(v int32) *CreateScheduledScalingRuleRequest {
	s.DurationMinutes = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetEnable(v bool) *CreateScheduledScalingRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetFirstScheduledTime(v int64) *CreateScheduledScalingRuleRequest {
	s.FirstScheduledTime = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetInstanceId(v string) *CreateScheduledScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetRegionId(v string) *CreateScheduledScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetRepeatType(v string) *CreateScheduledScalingRuleRequest {
	s.RepeatType = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetReservedPubFlow(v int32) *CreateScheduledScalingRuleRequest {
	s.ReservedPubFlow = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetReservedSubFlow(v int32) *CreateScheduledScalingRuleRequest {
	s.ReservedSubFlow = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetRuleName(v string) *CreateScheduledScalingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetScheduleType(v string) *CreateScheduledScalingRuleRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetTimeZone(v string) *CreateScheduledScalingRuleRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetWeeklyTypes(v []*string) *CreateScheduledScalingRuleRequest {
	s.WeeklyTypes = v
	return s
}

type CreateScheduledScalingRuleShrinkRequest struct {
	// The duration of each scheduled scaling task. Unit: minutes.
	//
	// >  The value of this parameter must be greater than or equal to 15.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	DurationMinutes *int32 `json:"DurationMinutes,omitempty" xml:"DurationMinutes,omitempty"`
	// Specifies whether to enable the scheduled scaling rule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the scheduled scaling task is executed.
	//
	// If you set ScheduleType to at, make sure that the value of this parameter is at least 30 minutes later than the current point in time.
	//
	// 	Notice: To prevent the broker from repeatedly executing instance upgrade and downgrade tasks, make sure that the interval between two consecutive scheduled scaling tasks is at least 60 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1714467540000
	FirstScheduledTime *int64 `json:"FirstScheduledTime,omitempty" xml:"FirstScheduledTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_serverless-cn-vxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The frequency to execute the scheduled scaling task. This parameter is required only if you set ScheduleType to repeat. Valid values:
	//
	// 	- Daily: The scheduled scaling task is executed every day.
	//
	// 	- Weekly: The scheduled scaling task is executed every week.
	//
	// example:
	//
	// Weekly
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The reserved production capacity for scheduled scaling. Unit: MB/s.
	//
	// >  You must specify a higher value than the instance specification for at least one of ReservedPubFlow and ReservedSubFlow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	ReservedPubFlow *int32 `json:"ReservedPubFlow,omitempty" xml:"ReservedPubFlow,omitempty"`
	// The reserved consumption capacity for scheduled scaling. Unit: MB/s.
	//
	// >  You must specify a higher value than the instance specification for at least one of ReservedPubFlow and ReservedSubFlow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	ReservedSubFlow *int32 `json:"ReservedSubFlow,omitempty" xml:"ReservedSubFlow,omitempty"`
	// The name of the scheduled scaling rule.
	//
	// >  The name of the scheduled scaling rule cannot be the same as the names of other rules for the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the scheduled scaling task. Valid values:
	//
	// 	- at: The scheduled scaling task is executed only once.
	//
	// 	- repeat: The scheduled scaling task is repeatedly executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// at
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The time zone in Coordinated Universal Time (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// GMT+8
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The day on which the scheduled scaling task is executed every week. You can specify multiple days.
	WeeklyTypesShrink *string `json:"WeeklyTypes,omitempty" xml:"WeeklyTypes,omitempty"`
}

func (s CreateScheduledScalingRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledScalingRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetDurationMinutes(v int32) *CreateScheduledScalingRuleShrinkRequest {
	s.DurationMinutes = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetEnable(v bool) *CreateScheduledScalingRuleShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetFirstScheduledTime(v int64) *CreateScheduledScalingRuleShrinkRequest {
	s.FirstScheduledTime = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetInstanceId(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetRegionId(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetRepeatType(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.RepeatType = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetReservedPubFlow(v int32) *CreateScheduledScalingRuleShrinkRequest {
	s.ReservedPubFlow = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetReservedSubFlow(v int32) *CreateScheduledScalingRuleShrinkRequest {
	s.ReservedSubFlow = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetRuleName(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetScheduleType(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetTimeZone(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetWeeklyTypesShrink(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.WeeklyTypesShrink = &v
	return s
}

type CreateScheduledScalingRuleResponseBody struct {
	// The response code. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB6F1BEA-903B-4FD8-8809-46E7E9CE***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScheduledScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledScalingRuleResponseBody) SetCode(v int64) *CreateScheduledScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScheduledScalingRuleResponseBody) SetMessage(v string) *CreateScheduledScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScheduledScalingRuleResponseBody) SetRequestId(v string) *CreateScheduledScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledScalingRuleResponseBody) SetSuccess(v bool) *CreateScheduledScalingRuleResponseBody {
	s.Success = &v
	return s
}

type CreateScheduledScalingRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledScalingRuleResponse) SetHeaders(v map[string]*string) *CreateScheduledScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledScalingRuleResponse) SetStatusCode(v int32) *CreateScheduledScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledScalingRuleResponse) SetBody(v *CreateScheduledScalingRuleResponseBody) *CreateScheduledScalingRuleResponse {
	s.Body = v
	return s
}

type CreateTopicRequest struct {
	// The log cleanup policy that is used for the topic. This parameter is available only when LocalTopic is set to true. Valid values:
	//
	// 	- false: The topic uses the default log cleanup policy.
	//
	// 	- true: The topic uses the log compaction policy.
	//
	// example:
	//
	// false
	CompactTopic *bool `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	// The additional configuration.
	//
	// 	- The value must be in JSON format.
	//
	// 	- Set Key to **replications**. This value specifies the number of replicas of the topic. The value must be an integer that ranges from 1 to 3.
	//
	// 	- You can configure this parameter only if you set **LocalTopic*	- to **true*	- or specify **Open Source Edition (Local Disk)*	- as the instance edition.****
	//
	// >  If you specify replications in this parameter, **ReplicationFactor*	- does not take effect.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// {"replications": 3}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of storage that the topic uses. Valid values:
	//
	// 	- false: The topic uses cloud storage.
	//
	// 	- true: The topic uses local storage.
	//
	// example:
	//
	// false
	LocalTopic *bool `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	// The minimum number of in-sync replicas (ISRs).
	//
	// 	- This parameter is available only when **LocalTopic*	- is set to **true**, or the instance is of the **Open Source Edition (Local Disk)**.****
	//
	// 	- The value of this parameter must be smaller than the value of ReplicationFactor.
	//
	// 	- Valid values: 1 to 3.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" xml:"MinInsyncReplicas,omitempty"`
	// The number of partitions in the topic.
	//
	// 	- Valid values: 1 to 360.
	//
	// 	- In the ApsaraMQ for Kafka console, you can view the number of partitions that the system recommends based on the specifications of the instance. We recommend that you specify the number that is recommended by the system as the value of this parameter to reduce the risk of data skew.
	//
	// Default values:
	//
	// 	- ApsaraMQ for Kafka V2 instance: 12
	//
	// 	- ApsaraMQ for Kafka V3 instance: 3
	//
	// example:
	//
	// 12
	PartitionNum *string `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance in which you want to create a topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the topic.
	//
	// 	- The description can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The description must be 3 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_topic_test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The number of replicas for the topic.
	//
	// 	- This parameter is available only when **LocalTopic*	- is set to **true**, or the instance is of the **Open Source Edition (Local Disk)**.****
	//
	// 	- Valid values: 1 to 3.
	//
	// > If you set this parameter to **1**, data loss may occur. Exercise caution when you configure this parameter.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 3
	ReplicationFactor *int64 `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	// The tags that you want to add to the topic.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The topic name.
	//
	// 	- The name can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be 3 to 64 characters in length. If the name that you specify contains more than 64 characters, the system automatically truncates the name.
	//
	// 	- After a topic is created, you cannot change the name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_topic_test
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
	// 	- If you do not specify this parameter, the keys of all tags are matched.
	//
	// 	- The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// 	- You can leave this parameter empty.
	//
	// 	- The tag value must be 1 to 128 characters in length and cannot contain http:// or https://. The tag value cannot start with aliyun or acs:.
	//
	// example:
	//
	// FinanceJoshua
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C0F207C-77A6-43E5-991C-9D98510A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
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
	// The type of the operation allowed by the access control list (ACL). Valid values:
	//
	// 	- **Write**: data writes.
	//
	// 	- **Read**: data reads.
	//
	// 	- **Describe**: reads of transaction IDs.
	//
	// 	- **IdempotentWrite**: idempotent data writes to clusters.
	//
	// 	- **IDEMPOTENT_WRITE**: idempotent data writes to clusters. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// 	- **DESCRIBE_CONFIGS**: configuration query. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// Write
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The types of operations allowed by the ACL. Separate multiple operations with commas (,).
	//
	// Valid values:
	//
	// 	- **Write**: data writes.
	//
	// 	- **Read**: data reads.
	//
	// 	- **Describe**: reads of transaction IDs.
	//
	// 	- **IdempotentWrite**: idempotent data writes to clusters.
	//
	// 	- **IDEMPOTENT_WRITE**: idempotent data writes to clusters. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// 	- **DESCRIBE_CONFIGS**: configuration query. This value is available only for serverless ApsaraMQ for Kafka instances.
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// Write,Read
	AclOperationTypes *string `json:"AclOperationTypes,omitempty" xml:"AclOperationTypes,omitempty"`
	// The authorization method. Valid values:
	//
	// 	- Deny
	//
	// 	- ALLOW
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// The name of the resource.
	//
	// 	- The value can be the name of a topic or consumer group.
	//
	// 	- You can use an asterisk (\\*) to indicate the names of all topics or consumer groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The mode that is used to match resources. Valid values:
	//
	// 	- **LITERAL:*	- full match
	//
	// 	- **PREFIXED**: prefix match
	//
	// This parameter is required.
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **Topic**: topic
	//
	// 	- **Group**: consumer group
	//
	// 	- **Cluster**: cluster
	//
	// 	- **TransactionalId**: transactional ID
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The IP address of the source.
	//
	// >
	//
	// 	- You can specify only a specific IP address or use the asterisk (\\*) wildcard character to specify all IP addresses. CIDR blocks are not supported.
	//
	// 	- This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// *
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test12****
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

func (s *DeleteAclRequest) SetAclOperationTypes(v string) *DeleteAclRequest {
	s.AclOperationTypes = &v
	return s
}

func (s *DeleteAclRequest) SetAclPermissionType(v string) *DeleteAclRequest {
	s.AclPermissionType = &v
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

func (s *DeleteAclRequest) SetHost(v string) *DeleteAclRequest {
	s.Host = &v
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B0740227-AA9A-4E14-8E9F-36ED665****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	//
	// This parameter is required.
	//
	// example:
	//
	// CID-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The encryption method. Valid values:
	//
	// 	- SCRAM-SHA-512. This is the default value.
	//
	// 	- SCRAM-SHA-256
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// SCRAM-SHA-256
	Mechanism *string `json:"Mechanism,omitempty" xml:"Mechanism,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the Simple Authentication and Security Layer (SASL) user. Valid values:
	//
	// 	- **plain**: a simple mechanism that uses usernames and passwords to verify user identities. ApsaraMQ for Kafka provides an improved PLAIN mechanism that allows you to dynamically add SASL users without the need to restart an instance.
	//
	// 	- **SCRAM**: a mechanism that uses usernames and passwords to verify user identities. Compared with the PLAIN mechanism, this mechanism provides better security protection. ApsaraMQ for Kafka uses the SCRAM-SHA-256 algorithm.
	//
	// 	- **LDAP**: This value is available only for the SASL users of ApsaraMQ for Confluent instances.
	//
	// Default value: **plain**.
	//
	// example:
	//
	// scram
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***
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

func (s *DeleteSaslUserRequest) SetMechanism(v string) *DeleteSaslUserRequest {
	s.Mechanism = &v
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
	// The HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3CB89F5C-CD97-4C1D-BC7C-FEDEC2F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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

type DeleteScheduledScalingRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_serverless-cn-vxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the scheduled scaling rule.
	//
	// >  You can delete only rules that are disabled and rules that are scheduled only once and have been executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-name-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteScheduledScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledScalingRuleRequest) SetInstanceId(v string) *DeleteScheduledScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScheduledScalingRuleRequest) SetRegionId(v string) *DeleteScheduledScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScheduledScalingRuleRequest) SetRuleName(v string) *DeleteScheduledScalingRuleRequest {
	s.RuleName = &v
	return s
}

type DeleteScheduledScalingRuleResponseBody struct {
	// The responses code. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteScheduledScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledScalingRuleResponseBody) SetCode(v int64) *DeleteScheduledScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponseBody) SetMessage(v string) *DeleteScheduledScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponseBody) SetRequestId(v string) *DeleteScheduledScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponseBody) SetSuccess(v bool) *DeleteScheduledScalingRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteScheduledScalingRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteScheduledScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledScalingRuleResponse) SetStatusCode(v int32) *DeleteScheduledScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledScalingRuleResponse) SetBody(v *DeleteScheduledScalingRuleResponseBody) *DeleteScheduledScalingRuleResponse {
	s.Body = v
	return s
}

type DeleteTopicRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
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
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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

type DescribeAclResourceNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAclResourceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclResourceNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclResourceNameRequest) SetAclResourcePatternType(v string) *DescribeAclResourceNameRequest {
	s.AclResourcePatternType = &v
	return s
}

func (s *DescribeAclResourceNameRequest) SetAclResourceType(v string) *DescribeAclResourceNameRequest {
	s.AclResourceType = &v
	return s
}

func (s *DescribeAclResourceNameRequest) SetInstanceId(v string) *DescribeAclResourceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAclResourceNameRequest) SetRegionId(v string) *DescribeAclResourceNameRequest {
	s.RegionId = &v
	return s
}

type DescribeAclResourceNameResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeAclResourceNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAclResourceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclResourceNameResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAclResourceNameResponseBody) SetCode(v int32) *DescribeAclResourceNameResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAclResourceNameResponseBody) SetData(v *DescribeAclResourceNameResponseBodyData) *DescribeAclResourceNameResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAclResourceNameResponseBody) SetMessage(v string) *DescribeAclResourceNameResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAclResourceNameResponseBody) SetRequestId(v string) *DescribeAclResourceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAclResourceNameResponseBody) SetSuccess(v bool) *DescribeAclResourceNameResponseBody {
	s.Success = &v
	return s
}

type DescribeAclResourceNameResponseBodyData struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeAclResourceNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclResourceNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAclResourceNameResponseBodyData) SetData(v []*string) *DescribeAclResourceNameResponseBodyData {
	s.Data = v
	return s
}

type DescribeAclResourceNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclResourceNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclResourceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclResourceNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclResourceNameResponse) SetHeaders(v map[string]*string) *DescribeAclResourceNameResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclResourceNameResponse) SetStatusCode(v int32) *DescribeAclResourceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclResourceNameResponse) SetBody(v *DescribeAclResourceNameResponseBody) *DescribeAclResourceNameResponse {
	s.Body = v
	return s
}

type DescribeAclsRequest struct {
	// The types of operations allowed by the ACL. Separate multiple operation types with commas (,).
	//
	// - Valid values:
	//
	// - Write
	//
	// - Read
	//
	// - Describe: reads of transactional IDs.
	//
	// - IdempotentWrite: idempotent data writes to clusters.
	//
	// - IDEMPOTENT_WRITE: idempotent data writes to clusters. This value is available only for ApsaraMQ for Kafka V3 instances.
	//
	// - DESCRIBE_CONFIGS: queries of configurations. This value is available only for ApsaraMQ for Kafka V3 instances.
	//
	// > This parameter is available only for ApsaraMQ for Kafka V3 serverless instances.
	//
	// example:
	//
	// Write
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The authorization method. Valid values:
	//
	// - DENY
	//
	// - ALLOW
	//
	// > This parameter is available only for ApsaraMQ for Kafka V3 serverless instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// The resource name.
	//
	// 	- The value can be the name of a topic or consumer group.
	//
	// 	- You can use an asterisk (\\*) to specify the names of all topics or consumer groups.
	//
	// > You can query the resources on which permissions are granted only after you grant the user the required permissions on all resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The match mode. Valid values:
	//
	// 	- LITERAL: full-name match
	//
	// 	- PREFIXED: prefix match
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **Topic**
	//
	// 	- **Group**
	//
	// This parameter is required.
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The source IP address.
	//
	// >-  You can specify only a specific IP address or use the asterisk (*) wildcard character to specify all IP addresses. CIDR blocks are not supported.
	//
	// > - This parameter is available only for ApsaraMQ for Kafka V3 serverless instances.
	//
	// example:
	//
	// *
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The username.
	//
	// 	- You can use an asterisk (\\*) to specify all users.
	//
	// > You can use an asterisk (\\*) to query the authorized users only after you grant the required permissions to all users.
	//
	// This parameter is required.
	//
	// example:
	//
	// test12****
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeAclsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAclsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclsRequest) SetAclOperationType(v string) *DescribeAclsRequest {
	s.AclOperationType = &v
	return s
}

func (s *DescribeAclsRequest) SetAclPermissionType(v string) *DescribeAclsRequest {
	s.AclPermissionType = &v
	return s
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

func (s *DescribeAclsRequest) SetHost(v string) *DescribeAclsRequest {
	s.Host = &v
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The access control lists (ACLs).
	KafkaAclList *DescribeAclsResponseBodyKafkaAclList `json:"KafkaAclList,omitempty" xml:"KafkaAclList,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 46496E38-881E-4719-A2F3-F3DA6AE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// The types of operations allowed by the ACL. Separate multiple operation types with commas (,).
	//
	// - Valid values:
	//
	// - Write
	//
	// - Read
	//
	// - Describe: reads of transactional IDs.
	//
	// - IdempotentWrite: idempotent data writes to clusters.
	//
	// - IDEMPOTENT_WRITE: idempotent data writes to clusters. This value is available only for ApsaraMQ for Kafka V3 instances.
	//
	// - DESCRIBE_CONFIGS: queries of configurations. This value is available only for ApsaraMQ for Kafka V3 instances.
	//
	// > This parameter is available only for ApsaraMQ for Kafka V3 serverless instances.
	//
	// example:
	//
	// Write
	AclOperationType *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	// The authorization method. Valid values:
	//
	// - DENY
	//
	// - ALLOW
	//
	// > This parameter is available only for ApsaraMQ for Kafka V3 serverless instances.
	//
	// example:
	//
	// DENY
	AclPermissionType *string `json:"AclPermissionType,omitempty" xml:"AclPermissionType,omitempty"`
	// The resource name.
	//
	// 	- The value can be the name of a topic or consumer group.
	//
	// 	- You can use the asterisk (\\*) wildcard character to specify the names of all topics or consumer groups.
	//
	// example:
	//
	// demo
	AclResourceName *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	// The matching mode. Valid values:
	//
	// 	- **LITERAL:*	- full-name match
	//
	// 	- **PREFIXED**: prefix match
	//
	// example:
	//
	// LITERAL
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	// The resource type. Valid values:
	//
	// 	- **Topic**
	//
	// 	- **Group**
	//
	// example:
	//
	// Topic
	AclResourceType *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	// The host.
	//
	// example:
	//
	// ****
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The username.
	//
	// example:
	//
	// test12***
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

func (s *DescribeAclsResponseBodyKafkaAclListKafkaAclVO) SetAclPermissionType(v string) *DescribeAclsResponseBodyKafkaAclListKafkaAclVO {
	s.AclPermissionType = &v
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
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9E3B3592-5994-4F65-A61E-E62A77A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Simple Authentication and Security Layer (SASL) users.
	SaslUserList *DescribeSaslUsersResponseBodySaslUserList `json:"SaslUserList,omitempty" xml:"SaslUserList,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// The encryption method.
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// SCRAM-SHA-256
	Mechanism *string `json:"Mechanism,omitempty" xml:"Mechanism,omitempty"`
	// The password.
	//
	// example:
	//
	// ******
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The type of the SASL user. Valid values:
	//
	// 	- **plain**: a simple mechanism that uses usernames and passwords to verify user identities. ApsaraMQ for Kafka provides an improved PLAIN mechanism that allows you to dynamically add SASL users without the need to restart an instance.
	//
	// 	- **SCRAM**: a mechanism that uses usernames and passwords to verify user identities. Compared with the PLAIN mechanism, this mechanism provides better security protection. ApsaraMQ for Kafka uses the SCRAM-SHA-256 algorithm.
	//
	// 	- **LDAP**: This value is available only for the SASL users of ApsaraMQ for Confluent instances.
	//
	// Default value: **plain**.
	//
	// example:
	//
	// scram
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username.
	//
	// example:
	//
	// test12***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeSaslUsersResponseBodySaslUserListSaslUserVO) String() string {
	return tea.Prettify(s)
}

func (s DescribeSaslUsersResponseBodySaslUserListSaslUserVO) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersResponseBodySaslUserListSaslUserVO) SetMechanism(v string) *DescribeSaslUsersResponseBodySaslUserListSaslUserVO {
	s.Mechanism = &v
	return s
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

type EnableAutoGroupCreationRequest struct {
	// Specify whether to enable the flexible group creation feature. Valid values:
	//
	// 	- **true**: enables the flexible group creation feature.
	//
	// 	- **false**: disabled the flexible group creation feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The instance ID.
	//
	// You can call the [GetInstanceList](https://help.aliyun.com/document_detail/437663.html) operation to query instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EnableAutoGroupCreationRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoGroupCreationRequest) GoString() string {
	return s.String()
}

func (s *EnableAutoGroupCreationRequest) SetEnable(v bool) *EnableAutoGroupCreationRequest {
	s.Enable = &v
	return s
}

func (s *EnableAutoGroupCreationRequest) SetInstanceId(v string) *EnableAutoGroupCreationRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableAutoGroupCreationRequest) SetRegionId(v string) *EnableAutoGroupCreationRequest {
	s.RegionId = &v
	return s
}

type EnableAutoGroupCreationResponseBody struct {
	// The returned HTTP status code.
	//
	// If the value **200*	- is returned, the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A421CCD7-5BC5-4B32-8DD8-64668A8FCB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableAutoGroupCreationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoGroupCreationResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAutoGroupCreationResponseBody) SetCode(v int32) *EnableAutoGroupCreationResponseBody {
	s.Code = &v
	return s
}

func (s *EnableAutoGroupCreationResponseBody) SetMessage(v string) *EnableAutoGroupCreationResponseBody {
	s.Message = &v
	return s
}

func (s *EnableAutoGroupCreationResponseBody) SetRequestId(v string) *EnableAutoGroupCreationResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableAutoGroupCreationResponseBody) SetSuccess(v bool) *EnableAutoGroupCreationResponseBody {
	s.Success = &v
	return s
}

type EnableAutoGroupCreationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableAutoGroupCreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAutoGroupCreationResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoGroupCreationResponse) GoString() string {
	return s.String()
}

func (s *EnableAutoGroupCreationResponse) SetHeaders(v map[string]*string) *EnableAutoGroupCreationResponse {
	s.Headers = v
	return s
}

func (s *EnableAutoGroupCreationResponse) SetStatusCode(v int32) *EnableAutoGroupCreationResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAutoGroupCreationResponse) SetBody(v *EnableAutoGroupCreationResponseBody) *EnableAutoGroupCreationResponse {
	s.Body = v
	return s
}

type EnableAutoTopicCreationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation that you want to perform. Valid values:
	//
	// 	- enable: enables the automatic topic creation feature.
	//
	// 	- disable: disables the automatic topic creation feature.
	//
	// 	- updatePartition: changes the number of partitions in topics that are automatically created.
	//
	// example:
	//
	// enable
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// The changed number of partitions in topics that are automatically created.
	//
	// This parameter takes effect only if you set Operate to updatePartition.
	//
	// example:
	//
	// 12
	PartitionNum *int64 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpdatePartition *bool   `json:"UpdatePartition,omitempty" xml:"UpdatePartition,omitempty"`
}

func (s EnableAutoTopicCreationRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoTopicCreationRequest) GoString() string {
	return s.String()
}

func (s *EnableAutoTopicCreationRequest) SetInstanceId(v string) *EnableAutoTopicCreationRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableAutoTopicCreationRequest) SetOperate(v string) *EnableAutoTopicCreationRequest {
	s.Operate = &v
	return s
}

func (s *EnableAutoTopicCreationRequest) SetPartitionNum(v int64) *EnableAutoTopicCreationRequest {
	s.PartitionNum = &v
	return s
}

func (s *EnableAutoTopicCreationRequest) SetRegionId(v string) *EnableAutoTopicCreationRequest {
	s.RegionId = &v
	return s
}

func (s *EnableAutoTopicCreationRequest) SetUpdatePartition(v bool) *EnableAutoTopicCreationRequest {
	s.UpdatePartition = &v
	return s
}

type EnableAutoTopicCreationResponseBody struct {
	// The returned status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E3B3592-5994-4F65-A61E-E62A77A7***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableAutoTopicCreationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoTopicCreationResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAutoTopicCreationResponseBody) SetCode(v int32) *EnableAutoTopicCreationResponseBody {
	s.Code = &v
	return s
}

func (s *EnableAutoTopicCreationResponseBody) SetMessage(v string) *EnableAutoTopicCreationResponseBody {
	s.Message = &v
	return s
}

func (s *EnableAutoTopicCreationResponseBody) SetRequestId(v string) *EnableAutoTopicCreationResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableAutoTopicCreationResponseBody) SetSuccess(v bool) *EnableAutoTopicCreationResponseBody {
	s.Success = &v
	return s
}

type EnableAutoTopicCreationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableAutoTopicCreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableAutoTopicCreationResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoTopicCreationResponse) GoString() string {
	return s.String()
}

func (s *EnableAutoTopicCreationResponse) SetHeaders(v map[string]*string) *EnableAutoTopicCreationResponse {
	s.Headers = v
	return s
}

func (s *EnableAutoTopicCreationResponse) SetStatusCode(v int32) *EnableAutoTopicCreationResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAutoTopicCreationResponse) SetBody(v *EnableAutoTopicCreationResponseBody) *EnableAutoTopicCreationResponse {
	s.Body = v
	return s
}

type GetAllInstanceIdListRequest struct {
	// The region ID of the instance. This parameter is reserved.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of instances that are managed by the Alibaba Cloud account in all the regions.
	//
	// example:
	//
	// [{"cn-shenzhen": ["alikafka_post-cn-7pp2btvo****"],"us-west-1": ["alikafka_pre-cn-i7m2lxid****"],"cn-hangzhou": ["alikafka_pre-cn-i7m2hflj****","alikafka_pre-cn-zvp2hsje****","alikafka_pre-cn-zvp2kvc9****"]}]
	InstanceIds map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp91inkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A421CCD7-5BC5-4B32-8DD8-64668A8FCB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// 	- **4**: allows access from the Internet and a virtual private cloud (VPC).
	//
	// 	- **5**: allows access from a VPC.
	//
	// >  Only integrators need to concern themselves with the value of this parameter.
	//
	// example:
	//
	// 4
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
	AllowedIpList []*string          `json:"AllowedIpList,omitempty" xml:"AllowedIpList,omitempty" type:"Repeated"`
	BlackIPList   []*string          `json:"BlackIPList,omitempty" xml:"BlackIPList,omitempty" type:"Repeated"`
	BlackIPMap    map[string]*string `json:"BlackIPMap,omitempty" xml:"BlackIPMap,omitempty"`
	// The port range. Valid value:
	//
	// **9093/9093**.
	//
	// example:
	//
	// 9093/9093
	PortRange                      *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SecurityGroupId                *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	UserDefinedSharedSecurityGroup *bool   `json:"UserDefinedSharedSecurityGroup,omitempty" xml:"UserDefinedSharedSecurityGroup,omitempty"`
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

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetBlackIPList(v []*string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.BlackIPList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetBlackIPMap(v map[string]*string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.BlackIPMap = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetPortRange(v string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.PortRange = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetSecurityGroupId(v string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.SecurityGroupId = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetUserDefinedSharedSecurityGroup(v bool) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.UserDefinedSharedSecurityGroup = &v
	return s
}

type GetAllowedIpListResponseBodyAllowedListVpcList struct {
	// The group to which the IP address whitelist belongs.
	AllowedIpGroup map[string]*string `json:"AllowedIpGroup,omitempty" xml:"AllowedIpGroup,omitempty"`
	// The information about the IP address whitelist.
	AllowedIpList []*string          `json:"AllowedIpList,omitempty" xml:"AllowedIpList,omitempty" type:"Repeated"`
	BlackIPList   []*string          `json:"BlackIPList,omitempty" xml:"BlackIPList,omitempty" type:"Repeated"`
	BlackIPMap    map[string]*string `json:"BlackIPMap,omitempty" xml:"BlackIPMap,omitempty"`
	// The port range. Valid value:
	//
	// **9092/9092**.
	//
	// example:
	//
	// 9092/9092
	PortRange                      *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SecurityGroupId                *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	UserDefinedSharedSecurityGroup *bool   `json:"UserDefinedSharedSecurityGroup,omitempty" xml:"UserDefinedSharedSecurityGroup,omitempty"`
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

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetBlackIPList(v []*string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.BlackIPList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetBlackIPMap(v map[string]*string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.BlackIPMap = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetPortRange(v string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.PortRange = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetSecurityGroupId(v string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.SecurityGroupId = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetUserDefinedSharedSecurityGroup(v bool) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.UserDefinedSharedSecurityGroup = &v
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

type GetAutoScalingConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_serverless-cn-vxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAutoScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationRequest) SetInstanceId(v string) *GetAutoScalingConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingConfigurationRequest) SetRegionId(v string) *GetAutoScalingConfigurationRequest {
	s.RegionId = &v
	return s
}

type GetAutoScalingConfigurationResponseBody struct {
	// The response code. The value 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetAutoScalingConfigurationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B7A39AE5-0B36-4442-A304-E0885265***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBody) SetCode(v int64) *GetAutoScalingConfigurationResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) SetData(v *GetAutoScalingConfigurationResponseBodyData) *GetAutoScalingConfigurationResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) SetMessage(v string) *GetAutoScalingConfigurationResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) SetRequestId(v string) *GetAutoScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBody) SetSuccess(v bool) *GetAutoScalingConfigurationResponseBody {
	s.Success = &v
	return s
}

type GetAutoScalingConfigurationResponseBodyData struct {
	// The scheduled scaling rules.
	ScheduledScalingRules *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules `json:"ScheduledScalingRules,omitempty" xml:"ScheduledScalingRules,omitempty" type:"Struct"`
}

func (s GetAutoScalingConfigurationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBodyData) SetScheduledScalingRules(v *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) *GetAutoScalingConfigurationResponseBodyData {
	s.ScheduledScalingRules = v
	return s
}

type GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules struct {
	ScheduledScalingRules []*GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules `json:"ScheduledScalingRules,omitempty" xml:"ScheduledScalingRules,omitempty" type:"Repeated"`
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules) SetScheduledScalingRules(v []*GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRules {
	s.ScheduledScalingRules = v
	return s
}

type GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules struct {
	// The duration of a scheduled scaling task. Unit: minutes.
	//
	// example:
	//
	// 60
	DurationMinutes *int64 `json:"DurationMinutes,omitempty" xml:"DurationMinutes,omitempty"`
	// Indicates whether the scheduled scaling rule is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The estimated scale-in duration. Unit: seconds.
	//
	// example:
	//
	// 780
	EstimatedElasticScalingDownTimeSecs *int64 `json:"EstimatedElasticScalingDownTimeSecs,omitempty" xml:"EstimatedElasticScalingDownTimeSecs,omitempty"`
	// The estimated scale-out duration. Unit: seconds.
	//
	// example:
	//
	// 780
	EstimatedElasticScalingUpTimeSecs *int64 `json:"EstimatedElasticScalingUpTimeSecs,omitempty" xml:"EstimatedElasticScalingUpTimeSecs,omitempty"`
	// The timestamp that indicates the start time of the scheduled scaling task.
	//
	// example:
	//
	// 1714467540000
	FirstScheduledTime *int64 `json:"FirstScheduledTime,omitempty" xml:"FirstScheduledTime,omitempty"`
	// The frequency at which the scheduled scaling task is executed. This parameter is returned only if ScheduleType is set to repeat. Valid values:
	//
	// 	- Daily: The scheduled scaling task is executed every day.
	//
	// 	- Weekly: The scheduled scaling task is executed every week.
	//
	// example:
	//
	// Weekly
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The reserved production capacity for scheduled scaling. Unit: MB/s.
	//
	// example:
	//
	// 120
	ReservedPubFlow *int64 `json:"ReservedPubFlow,omitempty" xml:"ReservedPubFlow,omitempty"`
	// The reserved consumption capacity for scheduled scaling. Unit: MB/s.
	//
	// example:
	//
	// 120
	ReservedSubFlow *int64 `json:"ReservedSubFlow,omitempty" xml:"ReservedSubFlow,omitempty"`
	// The ID of the scheduled scaling rule.
	//
	// example:
	//
	// 64
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the scheduled scaling rule.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the scheduled scaling task. Valid values:
	//
	// 	- at: The scheduled scaling task is executed only once.
	//
	// 	- repeat: The scheduled scaling task is repeatedly executed.
	//
	// example:
	//
	// at
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The time zone in Coordinated Universal Time (UTC).
	//
	// example:
	//
	// GMT+8
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The day on which the scheduled scaling task is repeatedly executed. You can specify multiple days for this parameter.
	WeeklyTypes *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes `json:"WeeklyTypes,omitempty" xml:"WeeklyTypes,omitempty" type:"Struct"`
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetDurationMinutes(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.DurationMinutes = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetEnable(v bool) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.Enable = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetEstimatedElasticScalingDownTimeSecs(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.EstimatedElasticScalingDownTimeSecs = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetEstimatedElasticScalingUpTimeSecs(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.EstimatedElasticScalingUpTimeSecs = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetFirstScheduledTime(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.FirstScheduledTime = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetRepeatType(v string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.RepeatType = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetReservedPubFlow(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.ReservedPubFlow = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetReservedSubFlow(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.ReservedSubFlow = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetRuleId(v int64) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.RuleId = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetRuleName(v string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.RuleName = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetScheduleType(v string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.ScheduleType = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetTimeZone(v string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.TimeZone = &v
	return s
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules) SetWeeklyTypes(v *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRules {
	s.WeeklyTypes = v
	return s
}

type GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes struct {
	WeeklyTypes []*string `json:"WeeklyTypes,omitempty" xml:"WeeklyTypes,omitempty" type:"Repeated"`
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes) SetWeeklyTypes(v []*string) *GetAutoScalingConfigurationResponseBodyDataScheduledScalingRulesScheduledScalingRulesWeeklyTypes {
	s.WeeklyTypes = v
	return s
}

type GetAutoScalingConfigurationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigurationResponse) SetHeaders(v map[string]*string) *GetAutoScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScalingConfigurationResponse) SetStatusCode(v int32) *GetAutoScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScalingConfigurationResponse) SetBody(v *GetAutoScalingConfigurationResponseBody) *GetAutoScalingConfigurationResponse {
	s.Body = v
	return s
}

type GetConsumerListRequest struct {
	// The name of the consumer group. If you do not configure this parameter, all consumer groups are queried.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance to which the consumer group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h18sav****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to be returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance to which the consumer group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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

func (s *GetConsumerListRequest) SetCurrentPage(v int32) *GetConsumerListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetConsumerListRequest) SetInstanceId(v string) *GetConsumerListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetConsumerListRequest) SetPageSize(v int32) *GetConsumerListRequest {
	s.PageSize = &v
	return s
}

func (s *GetConsumerListRequest) SetRegionId(v string) *GetConsumerListRequest {
	s.RegionId = &v
	return s
}

type GetConsumerListResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The consumer groups.
	ConsumerList *GetConsumerListResponseBodyConsumerList `json:"ConsumerList,omitempty" xml:"ConsumerList,omitempty" type:"Struct"`
	// The number of the page to return. Pages start from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 808F042B-CB9A-4FBC-9009-00E7DDB6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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

func (s *GetConsumerListResponseBody) SetCurrentPage(v int32) *GetConsumerListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *GetConsumerListResponseBody) SetMessage(v string) *GetConsumerListResponseBody {
	s.Message = &v
	return s
}

func (s *GetConsumerListResponseBody) SetPageSize(v int32) *GetConsumerListResponseBody {
	s.PageSize = &v
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

func (s *GetConsumerListResponseBody) SetTotal(v int64) *GetConsumerListResponseBody {
	s.Total = &v
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
	// Indicates that the consumer group was automatically created by the system.
	//
	// example:
	//
	// false
	AutomaticallyCreatedGroup *bool `json:"AutomaticallyCreatedGroup,omitempty" xml:"AutomaticallyCreatedGroup,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The timestamp that indicates when the consumer group was created. Unit: milliseconds.
	//
	// example:
	//
	// 1729736584002
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// alikafka_post-cn-v0h18sav****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance description.
	//
	// example:
	//
	// test
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

func (s *GetConsumerListResponseBodyConsumerListConsumerVO) SetCreateTime(v int64) *GetConsumerListResponseBodyConsumerListConsumerVO {
	s.CreateTime = &v
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
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
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
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// Specifies whether to hide LastTimestamp. Default value: false. We recommend that you set this parameter to true.
	//
	// >
	//
	// 	- If you set this parameter to true, -1 is returned for LastTimestamp. If you set this parameter to false, a specific value is returned for LastTimestamp. This parameter is supported only by topics that use cloud storage on reserved instances.
	//
	// 	- A large amount of data is processed by this operation, which causes performance loss. We recommend that you set this parameter to true to accelerate processing.
	//
	// example:
	//
	// true
	HideLastTimestamp *bool `json:"HideLastTimestamp,omitempty" xml:"HideLastTimestamp,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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

func (s *GetConsumerProgressRequest) SetHideLastTimestamp(v bool) *GetConsumerProgressRequest {
	s.HideLastTimestamp = &v
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
	// The returned HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The consumer progress of the consumer group.
	ConsumerProgress *GetConsumerProgressResponseBodyConsumerProgress `json:"ConsumerProgress,omitempty" xml:"ConsumerProgress,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 252820E1-A2E6-45F2-B4C9-1056B8CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	//
	// example:
	//
	// 1566874931671
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The details of rebalances in the consumer group.
	RebalanceInfoList *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList `json:"RebalanceInfoList,omitempty" xml:"RebalanceInfoList,omitempty" type:"Struct"`
	// The consumer progress of each topic to which the consumer group subscribes.
	TopicList *GetConsumerProgressResponseBodyConsumerProgressTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
	// The total number of unconsumed messages in all topics to which the consumer group subscribes.
	//
	// example:
	//
	// 0
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

func (s *GetConsumerProgressResponseBodyConsumerProgress) SetRebalanceInfoList(v *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) *GetConsumerProgressResponseBodyConsumerProgress {
	s.RebalanceInfoList = v
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

type GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList struct {
	RebalanceInfoList []*GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList `json:"RebalanceInfoList,omitempty" xml:"RebalanceInfoList,omitempty" type:"Repeated"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList) SetRebalanceInfoList(v []*GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoList {
	s.RebalanceInfoList = v
	return s
}

type GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList struct {
	// The number of rebalances.
	//
	// example:
	//
	// 100
	Generation *int64 `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The group ID of the subscriber.
	//
	// example:
	//
	// falcon-uat
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the last rebalance occurred. Unit: milliseconds.
	//
	// example:
	//
	// 1709199270
	LastRebalanceTimestamp *int64 `json:"LastRebalanceTimestamp,omitempty" xml:"LastRebalanceTimestamp,omitempty"`
	// The cause of the rebalance.
	//
	// example:
	//
	// removing member consumer-1-cd14eb9c-379b-4b8e-9bbd-76f147f8536f on LeaveGroup
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Indicates whether new members are added to the consumer group in the rebalance.
	//
	// example:
	//
	// true
	RebalanceSuccess *bool `json:"RebalanceSuccess,omitempty" xml:"RebalanceSuccess,omitempty"`
	// The duration of the rebalance. Unit: milliseconds.
	//
	// example:
	//
	// 1
	RebalanceTimeConsuming *int64 `json:"RebalanceTimeConsuming,omitempty" xml:"RebalanceTimeConsuming,omitempty"`
}

func (s GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) GoString() string {
	return s.String()
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetGeneration(v int64) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.Generation = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetGroupId(v string) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.GroupId = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetLastRebalanceTimestamp(v int64) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.LastRebalanceTimestamp = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetReason(v string) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.Reason = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetRebalanceSuccess(v bool) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.RebalanceSuccess = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList) SetRebalanceTimeConsuming(v int64) *GetConsumerProgressResponseBodyConsumerProgressRebalanceInfoListRebalanceInfoList {
	s.RebalanceTimeConsuming = &v
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
	//
	// example:
	//
	// 1566874931649
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// The consumer offsets.
	OffsetList *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList `json:"OffsetList,omitempty" xml:"OffsetList,omitempty" type:"Struct"`
	// The topic name.
	//
	// example:
	//
	// kafka-test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The number of unconsumed messages in the topic to which the consumer group subscribes.
	//
	// example:
	//
	// 0
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
	//
	// example:
	//
	// 9
	BrokerOffset *int64 `json:"BrokerOffset,omitempty" xml:"BrokerOffset,omitempty"`
	// Client ID of the application.
	//
	// example:
	//
	// client-id-KafkaConsumerDemo
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 172.16.11.3
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// The consumer offset in the partition of the topic.
	//
	// example:
	//
	// 9
	ConsumerOffset *int64 `json:"ConsumerOffset,omitempty" xml:"ConsumerOffset,omitempty"`
	// The time when the last consumed message in the partition was generated.
	//
	// example:
	//
	// 1566874931649
	LastTimestamp *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	// Member ID.
	//
	// example:
	//
	// client-id-KafkaConsumerDemo-70b64883-a911-4882-8084-598b958848b4
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 0
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

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetClientId(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.ClientId = &v
	return s
}

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetClientIp(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.ClientIp = &v
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

func (s *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList) SetMemberId(v string) *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetListOffsetList {
	s.MemberId = &v
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
	//
	// example:
	//
	// alikafka_post-cn-mp91gnw0p***
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The ID of the order. You can obtain the order ID on the [Orders](https://usercenter2-intl.aliyun.com/order/list?pageIndex=1\\&pageSize=20\\&spm=5176.12818093.top-nav.ditem-ord.36f016d0OQFmJa) page in Alibaba Cloud User Center.
	//
	// example:
	//
	// 6072673****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can obtain this ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance version. You can use instance versions to filter different versions of instances. Valid values:
	//
	// 	- v2
	//
	// 	- v3
	//
	// 	- confluent
	//
	// example:
	//
	// v3
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
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

func (s *GetInstanceListRequest) SetSeries(v string) *GetInstanceListRequest {
	s.Series = &v
	return s
}

func (s *GetInstanceListRequest) SetTag(v []*GetInstanceListRequestTag) *GetInstanceListRequest {
	s.Tag = v
	return s
}

type GetInstanceListRequestTag struct {
	// The tag key.
	//
	// 	- If you leave this parameter empty, the keys of all tags are matched.
	//
	// 	- The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// 	- If you leave Key empty, you must also leave this parameter empty. If you leave this parameter empty, the values of all tags are matched.
	//
	// 	- The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The instances.
	InstanceList *GetInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// 4B6D821D-7F67-4CAA-9E13-A5A997C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
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
	//
	// example:
	//
	// {\\"enable.vpc_sasl_ssl\\":\\"false\\",\\"kafka.log.retention.hours\\":\\"66\\",\\"enable.acl\\":\\"false\\",\\"kafka.message.max.bytes\\":\\"6291456\\"}
	AllConfig *string `json:"AllConfig,omitempty" xml:"AllConfig,omitempty"`
	// Indicates whether the flexible group creation feature is enabled.
	//
	// example:
	//
	// true
	AutoCreateGroupEnable *bool `json:"AutoCreateGroupEnable,omitempty" xml:"AutoCreateGroupEnable,omitempty"`
	// Indicates whether the automatic topic creation feature is enabled.
	//
	// example:
	//
	// true
	AutoCreateTopicEnable *bool `json:"AutoCreateTopicEnable,omitempty" xml:"AutoCreateTopicEnable,omitempty"`
	// The ID of the secondary zone.
	//
	// example:
	//
	// cn-hangzhou-a
	BackupZoneId *string `json:"BackupZoneId,omitempty" xml:"BackupZoneId,omitempty"`
	// The parameters that are returned for the ApsaraMQ for Confluent instance.
	ConfluentConfig *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	// The time when the instance was created. Unit: milliseconds.
	//
	// example:
	//
	// 1577961819000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of partitions in a topic that is automatically created.
	//
	// example:
	//
	// 12
	DefaultPartitionNum *int32 `json:"DefaultPartitionNum,omitempty" xml:"DefaultPartitionNum,omitempty"`
	// The type of the network in which the instance is deployed. Valid values:
	//
	// 	- **4**: Internet and VPC
	//
	// 	- **5**: VPC
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size. Unit: GB
	//
	// example:
	//
	// 3600
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type of the instance. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// example:
	//
	// 1
	DiskType *int32 `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The default endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// alikafka-pre-cn-zv**********-1-vpc.alikafka.aliyuncs.com:9092,alikafka-pre-cn-zv**********-2-vpc.alikafka.aliyuncs.com:9092,alikafka-pre-cn-zv**********-3-vpc.alikafka.aliyuncs.com:9092
	DomainEndpoint *string `json:"DomainEndpoint,omitempty" xml:"DomainEndpoint,omitempty"`
	// The maximum Internet traffic in the instance.
	//
	// example:
	//
	// 20
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The default endpoint of the instance in IP address mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// 192.168.XX.XX:9092,192.168.XX.XX:9092,192.168.XX.XX:9092
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	// The time when the instance expires. Unit: milliseconds.
	//
	// example:
	//
	// 1893581018000
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic in the instance.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The maximum read traffic in the instance. Unit: Mbit/s.
	//
	// example:
	//
	// 1000
	IoMaxRead *int32 `json:"IoMaxRead,omitempty" xml:"IoMaxRead,omitempty"`
	// The traffic specification.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The maximum write traffic. Unit: Mbit/s.
	//
	// example:
	//
	// 1000
	IoMaxWrite *int32 `json:"IoMaxWrite,omitempty" xml:"IoMaxWrite,omitempty"`
	// The ID of the key that is used for disk encryption in the region where the instance is deployed.
	//
	// example:
	//
	// 0d24xxxx-da7b-4786-b981-9a164dxxxxxx
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The retention period of messages in the instance. Unit: hours.
	//
	// example:
	//
	// 72
	MsgRetain *int32 `json:"MsgRetain,omitempty" xml:"MsgRetain,omitempty"`
	// The instance name.
	//
	// example:
	//
	// alikafka_post-cn-mp91gnw0****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **0**: the subscription billing method
	//
	// 	- **1**: the pay-as-you-go billing method
	//
	// 	- **3**: the pay-as-you-go billing method for serverless ApsaraMQ for Kafka V3 instances
	//
	// 	- **4**: the pay-as-you-go billing method for ApsaraMQ for Confluent instances
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The traffic reserved for message publishing. Unit: MB/s.
	//
	// >  This parameter is returned only if the instance is a serverless ApsaraMQ for Kafka V3 instance.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int32 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The traffic reserved for message subscription. Unit: MB/s.
	//
	// >  This parameter is returned only if the instance is a serverless ApsaraMQ for Kafka V3 instance.
	//
	// example:
	//
	// 60
	ReservedSubscribeCapacity *int32 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Simple Authentication and Security Layer (SASL) endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// alikafka-pre-cn-zv**********-1-vpc.alikafka.aliyuncs.com:9094,alikafka-pre-cn-zv**********-2-vpc.alikafka.aliyuncs.com:9094,alikafka-pre-cn-zv**********-3-vpc.alikafka.aliyuncs.com:9094
	SaslDomainEndpoint *string `json:"SaslDomainEndpoint,omitempty" xml:"SaslDomainEndpoint,omitempty"`
	// The Simple Authentication and Security Layer (SASL) endpoint of the instance in IP address mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// 172.16.3.XX:9094,172.16.3.XX:9094,172.16.3.XX:9094
	SaslEndPoint *string `json:"SaslEndPoint,omitempty" xml:"SaslEndPoint,omitempty"`
	// The security group to which the instance belongs.
	//
	// 	- If the instance is deployed in the ApsaraMQ for Kafka console or by calling the [StartInstance](https://help.aliyun.com/document_detail/157786.html) operation without a security group configured, no value is returned.
	//
	// 	- If the instance is deployed by calling the [StartInstance](https://help.aliyun.com/document_detail/157786.html) operation with a security group configured, the returned value is the configured security group.
	//
	// example:
	//
	// sg-bp13wfx7kz9pkow****
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The instance version. Valid values: v2, v3, and confluent.
	//
	// example:
	//
	// v3
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// >  This parameter is out of date. We recommend that you refer to the ViewInstanceStatusCode parameter.
	//
	// The instance status. Valid values:
	//
	// 	- **0**: pending
	//
	// 	- **1**: preparing hardware resources
	//
	// 	- **2**: initializing
	//
	// 	- **3**: starting
	//
	// 	- **5**: running
	//
	// 	- **6**: migrating
	//
	// 	- **7**: ready for upgrade
	//
	// 	- **8**: upgrading
	//
	// 	- **9**: ready for change
	//
	// 	- **10**: released
	//
	// 	- **11**: changing
	//
	// 	- **15**: expired
	//
	// 	- **30**: scaling
	//
	// example:
	//
	// 5
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// 	- **normal**: Standard Edition
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The SSL endpoint of the instance in domain name mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// alikafka-pre-cn-zv**********-1.alikafka.aliyuncs.com:9093,alikafka-pre-cn-zv**********-2.alikafka.aliyuncs.com:9093,alikafka-pre-cn-zv**********-3.alikafka.aliyuncs.com:9093
	SslDomainEndpoint *string `json:"SslDomainEndpoint,omitempty" xml:"SslDomainEndpoint,omitempty"`
	// The Secure Sockets Layer (SSL) endpoint of the instance in IP address mode. ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// 192.0.XX.XX:9093,198.51.XX.XX:9093,203.0.XX.XX:9093
	SslEndPoint *string `json:"SslEndPoint,omitempty" xml:"SslEndPoint,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	StandardZoneId *string `json:"StandardZoneId,omitempty" xml:"StandardZoneId,omitempty"`
	// The tags.
	Tags *GetInstanceListResponseBodyInstanceListInstanceVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The maximum number of topics on the instance.
	//
	// example:
	//
	// 180
	TopicNumLimit *int32 `json:"TopicNumLimit,omitempty" xml:"TopicNumLimit,omitempty"`
	// The upgrade information about the instance.
	UpgradeServiceDetailInfo *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo `json:"UpgradeServiceDetailInfo,omitempty" xml:"UpgradeServiceDetailInfo,omitempty" type:"Struct"`
	// The number of used groups.
	//
	// example:
	//
	// 10
	UsedGroupCount *int32 `json:"UsedGroupCount,omitempty" xml:"UsedGroupCount,omitempty"`
	// The number of used partitions.
	//
	// example:
	//
	// 25
	UsedPartitionCount *int32 `json:"UsedPartitionCount,omitempty" xml:"UsedPartitionCount,omitempty"`
	// The number of used topics.
	//
	// example:
	//
	// 3
	UsedTopicCount *int32 `json:"UsedTopicCount,omitempty" xml:"UsedTopicCount,omitempty"`
	// The ID of the vSwitch to which the instance belongs.
	//
	// example:
	//
	// vsw-bp1fvuw0ljd7vzmo3****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch IDs.
	VSwitchIds *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	// The instance status. The valid values are consistent with the values displayed in the ApsaraMQ for Kafka console. This parameter is used in the new version of ApsaraMQ for Kafka.
	//
	// Valid values:
	//
	// 	- **0**: pending
	//
	// 	- **1**: deploying
	//
	// 	- **2**: running
	//
	// 	- **3**: stopped
	//
	// 	- **4**: expiring
	//
	// 	- **5**: expired
	//
	// 	- **6**: released
	//
	// 	- **7**: upgrading
	//
	// 	- **8**: migrating
	//
	// 	- **21**: stopping
	//
	// 	- **22**: starting
	//
	// 	- **23**: releasing
	//
	// 	- **30**: auto scaling
	//
	// 	- **101**: deployment failed
	//
	// 	- **102**: upgrade failed
	//
	// 	- **103**: migration failed
	//
	// example:
	//
	// 2
	ViewInstanceStatusCode *int32 `json:"ViewInstanceStatusCode,omitempty" xml:"ViewInstanceStatusCode,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-bp1ojac7bv448nifj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The SSL endpoint of the instance in domain name mode. You can use the endpoint to access the instance only in virtual private clouds (VPCs). ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// alikafka-post-cn-******-1-vpc.alikafka.aliyuncs.com:9095,alikafka-post-cn-******-2-vpc.alikafka.aliyuncs.com:9095,alikafka-post-cn-******-3-vpc.alikafka.aliyuncs.com:9095
	VpcSaslDomainEndpoint *string `json:"VpcSaslDomainEndpoint,omitempty" xml:"VpcSaslDomainEndpoint,omitempty"`
	// The Secure Sockets Layer (SSL) endpoint of the instance in IP address mode. You can use the endpoint to access the instance only in virtual private clouds (VPCs). ApsaraMQ for Kafka instances support endpoints in domain name mode and IP address mode.
	//
	// 	- Endpoints in domain name mode: An endpoint in this mode consists of the domain name of the instance and a port number. The format of an endpoint in this mode is `{Instance domain name}:{Port number}`.
	//
	// 	- Endpoints in IP address mode: An endpoint in this mode consists of the IP address of the broker and a port number. The format of an endpoint in this mode is `{Broker IP address}:{Port number}`.
	//
	// example:
	//
	// 172.16.3.XX:9095,172.16.3.XX:9095,172.16.3.XX:9095
	VpcSaslEndPoint *string `json:"VpcSaslEndPoint,omitempty" xml:"VpcSaslEndPoint,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// zonei
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

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAutoCreateGroupEnable(v bool) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AutoCreateGroupEnable = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetAutoCreateTopicEnable(v bool) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.AutoCreateTopicEnable = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetBackupZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.BackupZoneId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetConfluentConfig(v *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ConfluentConfig = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetCreateTime(v int64) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetDefaultPartitionNum(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.DefaultPartitionNum = &v
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

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxRead(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxRead = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxSpec(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxSpec = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetIoMaxWrite(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.IoMaxWrite = &v
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

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetReservedPublishCapacity(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetReservedSubscribeCapacity(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ReservedSubscribeCapacity = &v
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

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSaslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SaslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSecurityGroup(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.SecurityGroup = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetSeries(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.Series = &v
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

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVSwitchIds(v *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VSwitchIds = v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetViewInstanceStatusCode(v int32) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ViewInstanceStatusCode = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcId = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcSaslDomainEndpoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcSaslDomainEndpoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetVpcSaslEndPoint(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.VpcSaslEndPoint = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVO) SetZoneId(v string) *GetInstanceListResponseBodyInstanceListInstanceVO {
	s.ZoneId = &v
	return s
}

type GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig struct {
	// The number of CPU cores of Connect.
	//
	// example:
	//
	// 4
	ConnectCU *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	// The number of replicas of Connect.
	//
	// example:
	//
	// 2
	ConnectReplica *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	// The number of CPU cores of Control Center.
	//
	// example:
	//
	// 4
	ControlCenterCU *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	// The number of replicas of Control Center.
	//
	// example:
	//
	// 1
	ControlCenterReplica *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	// The disk capacity of Control Center. Unit: GB.
	//
	// example:
	//
	// 300
	ControlCenterStorage *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	// The number of CPU cores of the Kafka broker.
	//
	// example:
	//
	// 4
	KafkaCU *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	// The number of replicas of the Kafka broker.
	//
	// example:
	//
	// 3
	KafkaReplica *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	// The number of CPU cores of Kafka Rest Proxy.
	//
	// example:
	//
	// 4
	KafkaRestProxyCU *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	// The number of replicas of Kafka Rest Proxy.
	//
	// example:
	//
	// 2
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	// The disk capacity of the Kafka broker. Unit: GB.
	//
	// example:
	//
	// 800
	KafkaStorage *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	// The number of CPU cores of ksqlDB.
	//
	// example:
	//
	// 2
	KsqlCU *int32 `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	// The number of replicas of ksqlDB.
	//
	// example:
	//
	// 2
	KsqlReplica *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	// The disk capacity of ksqlDB. Unit: GB.
	//
	// example:
	//
	// 100
	KsqlStorage *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	// The number of CPU cores of Schema Registry.
	//
	// example:
	//
	// 4
	SchemaRegistryCU *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	// The number of replicas of Schema Registry.
	//
	// example:
	//
	// 2
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	// The number of CPU cores of ZooKeeper.
	//
	// example:
	//
	// 2
	ZooKeeperCU *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	// The number of replicas of ZooKeeper.
	//
	// example:
	//
	// 3
	ZooKeeperReplica *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	// The disk capacity of ZooKeeper. Unit: GB.
	//
	// example:
	//
	// 100
	ZooKeeperStorage *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetConnectCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetConnectReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetControlCenterStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaRestProxyCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaRestProxyReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKafkaStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetKsqlStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetSchemaRegistryCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetSchemaRegistryReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperCU(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperReplica(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig) SetZooKeeperStorage(v int32) *GetInstanceListResponseBodyInstanceListInstanceVOConfluentConfig {
	s.ZooKeeperStorage = &v
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
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
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
	//
	// example:
	//
	// 2.2.0
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

type GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds struct {
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) GoString() string {
	return s.String()
}

func (s *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds) SetVSwitchIds(v []*string) *GetInstanceListResponseBodyInstanceListInstanceVOVSwitchIds {
	s.VSwitchIds = v
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

type GetKafkaClientIpRequest struct {
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716343502000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the consumer group.
	//
	// >  This parameter is required only if you set Type to byGroup.
	//
	// example:
	//
	// group_name
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716343501000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The topic name.
	//
	// >
	//
	// 	- This parameter is required only if you set Type to byTopic.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The query method that you want to use to query the client IP addresses. Valid values:
	//
	// 	- byInstance: queries the IP addresses of the clients that are connected to the instance within a specific period of time.
	//
	// 	- byTopic: queries the IP addresses of the clients that are connected to a specific topic on the instance within a specific period of time.
	//
	// 	- byGroup: queries the IP addresses of the clients that are connected to a specific group on the instance within a specific period of time.
	//
	// This parameter is required.
	//
	// example:
	//
	// byInstance
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetKafkaClientIpRequest) String() string {
	return tea.Prettify(s)
}

func (s GetKafkaClientIpRequest) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpRequest) SetEndTime(v int64) *GetKafkaClientIpRequest {
	s.EndTime = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetGroup(v string) *GetKafkaClientIpRequest {
	s.Group = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetInstanceId(v string) *GetKafkaClientIpRequest {
	s.InstanceId = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetRegionId(v string) *GetKafkaClientIpRequest {
	s.RegionId = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetStartTime(v int64) *GetKafkaClientIpRequest {
	s.StartTime = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetTopic(v string) *GetKafkaClientIpRequest {
	s.Topic = &v
	return s
}

func (s *GetKafkaClientIpRequest) SetType(v string) *GetKafkaClientIpRequest {
	s.Type = &v
	return s
}

type GetKafkaClientIpResponseBody struct {
	// The returned status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetKafkaClientIpResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E57A8862-DF68-4055-8E55-B80CB4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKafkaClientIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKafkaClientIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBody) SetCode(v int64) *GetKafkaClientIpResponseBody {
	s.Code = &v
	return s
}

func (s *GetKafkaClientIpResponseBody) SetData(v *GetKafkaClientIpResponseBodyData) *GetKafkaClientIpResponseBody {
	s.Data = v
	return s
}

func (s *GetKafkaClientIpResponseBody) SetMessage(v string) *GetKafkaClientIpResponseBody {
	s.Message = &v
	return s
}

func (s *GetKafkaClientIpResponseBody) SetRequestId(v string) *GetKafkaClientIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKafkaClientIpResponseBody) SetSuccess(v bool) *GetKafkaClientIpResponseBody {
	s.Success = &v
	return s
}

type GetKafkaClientIpResponseBodyData struct {
	// The value true indicates that the broker is not of the latest minor version.
	//
	// >  If the broker is not of the latest minor version, the sampled logs may not be accurate. This may cause inaccurate IP information. Therefore, we recommend that you update your broker to the latest version at the earliest opportunity.
	//
	// example:
	//
	// true
	Alert *bool `json:"Alert,omitempty" xml:"Alert,omitempty"`
	// The data returned.
	Data *GetKafkaClientIpResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The end of the date range within which data is queried.
	//
	// example:
	//
	// 1716343502000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The time range within which the client IP addresses are queried.
	//
	// >  The valid value is 1 hour. If the beginning of the time range to query and the end of the time range to query exceeds 1 hour, only data within 1 hour is returned.
	//
	// example:
	//
	// 1
	SearchTimeRange *int32 `json:"SearchTimeRange,omitempty" xml:"SearchTimeRange,omitempty"`
	// The beginning of the date range within which data is queried.
	//
	// example:
	//
	// 1716343501000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The date range within which the client IP addresses are queried.
	//
	// >  The valid value is 7 days. If the beginning of the date range to query and the end of the date range to query exceeds 7 days, only data within 7 days is returned.
	//
	// example:
	//
	// 7
	TimeLimitDay *int32 `json:"TimeLimitDay,omitempty" xml:"TimeLimitDay,omitempty"`
}

func (s GetKafkaClientIpResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyData) SetAlert(v bool) *GetKafkaClientIpResponseBodyData {
	s.Alert = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetData(v *GetKafkaClientIpResponseBodyDataData) *GetKafkaClientIpResponseBodyData {
	s.Data = v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetEndDate(v int64) *GetKafkaClientIpResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetSearchTimeRange(v int32) *GetKafkaClientIpResponseBodyData {
	s.SearchTimeRange = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetStartDate(v int64) *GetKafkaClientIpResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyData) SetTimeLimitDay(v int32) *GetKafkaClientIpResponseBodyData {
	s.TimeLimitDay = &v
	return s
}

type GetKafkaClientIpResponseBodyDataData struct {
	Data []*GetKafkaClientIpResponseBodyDataDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetKafkaClientIpResponseBodyDataData) String() string {
	return tea.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyDataData) SetData(v []*GetKafkaClientIpResponseBodyDataDataData) *GetKafkaClientIpResponseBodyDataData {
	s.Data = v
	return s
}

type GetKafkaClientIpResponseBodyDataDataData struct {
	// The response parameters.
	Data *GetKafkaClientIpResponseBodyDataDataDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request name.
	//
	// >  The value of this parameter indicates the type of request that the client sends to the broker within the specified period of time.
	//
	// example:
	//
	// OFFSET_COMMIT
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetKafkaClientIpResponseBodyDataDataData) String() string {
	return tea.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyDataDataData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyDataDataData) SetData(v *GetKafkaClientIpResponseBodyDataDataDataData) *GetKafkaClientIpResponseBodyDataDataData {
	s.Data = v
	return s
}

func (s *GetKafkaClientIpResponseBodyDataDataData) SetName(v string) *GetKafkaClientIpResponseBodyDataDataData {
	s.Name = &v
	return s
}

type GetKafkaClientIpResponseBodyDataDataDataData struct {
	Data []*GetKafkaClientIpResponseBodyDataDataDataDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetKafkaClientIpResponseBodyDataDataDataData) String() string {
	return tea.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyDataDataDataData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyDataDataDataData) SetData(v []*GetKafkaClientIpResponseBodyDataDataDataDataData) *GetKafkaClientIpResponseBodyDataDataDataData {
	s.Data = v
	return s
}

type GetKafkaClientIpResponseBodyDataDataDataDataData struct {
	// The IP address of the client.
	//
	// example:
	//
	// 58.210.117.154
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The statistics.
	//
	// >  The value of this parameter indicates the number of connections on different ports of the IP address within the specified period of time.
	//
	// example:
	//
	// 3
	Num *int64 `json:"Num,omitempty" xml:"Num,omitempty"`
}

func (s GetKafkaClientIpResponseBodyDataDataDataDataData) String() string {
	return tea.Prettify(s)
}

func (s GetKafkaClientIpResponseBodyDataDataDataDataData) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponseBodyDataDataDataDataData) SetIp(v string) *GetKafkaClientIpResponseBodyDataDataDataDataData {
	s.Ip = &v
	return s
}

func (s *GetKafkaClientIpResponseBodyDataDataDataDataData) SetNum(v int64) *GetKafkaClientIpResponseBodyDataDataDataDataData {
	s.Num = &v
	return s
}

type GetKafkaClientIpResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKafkaClientIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKafkaClientIpResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKafkaClientIpResponse) GoString() string {
	return s.String()
}

func (s *GetKafkaClientIpResponse) SetHeaders(v map[string]*string) *GetKafkaClientIpResponse {
	s.Headers = v
	return s
}

func (s *GetKafkaClientIpResponse) SetStatusCode(v int32) *GetKafkaClientIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKafkaClientIpResponse) SetBody(v *GetKafkaClientIpResponseBody) *GetKafkaClientIpResponse {
	s.Body = v
	return s
}

type GetQuotaTipRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-i7m2wpm5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional message. This message is typically used to describe API call failures for troubleshooting.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The quota.
	QuotaData *GetQuotaTipResponseBodyQuotaData `json:"QuotaData,omitempty" xml:"QuotaData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0178A3A7-E87B-5E50-A16F-3E62F534****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
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
	//
	// example:
	//
	// 50
	GroupLeft *int32 `json:"GroupLeft,omitempty" xml:"GroupLeft,omitempty"`
	// The number of used groups.
	//
	// example:
	//
	// 50
	GroupUsed *int32 `json:"GroupUsed,omitempty" xml:"GroupUsed,omitempty"`
	// The method that you use to purchase partitions. Valid values:
	//
	// 	- 0: indicates that the instance is purchased based on topics.
	//
	// 	- 1: indicates that the instance is purchased based on partitions.
	//
	// example:
	//
	// 1
	IsPartitionBuy *int32 `json:"IsPartitionBuy,omitempty" xml:"IsPartitionBuy,omitempty"`
	// The number of available partitions.
	//
	// example:
	//
	// 1050
	PartitionLeft *int32 `json:"PartitionLeft,omitempty" xml:"PartitionLeft,omitempty"`
	// The number of purchased partitions.
	//
	// example:
	//
	// 100
	PartitionNumOfBuy *int32 `json:"PartitionNumOfBuy,omitempty" xml:"PartitionNumOfBuy,omitempty"`
	// The quota of partitions.
	//
	// example:
	//
	// 1100
	PartitionQuota *int32 `json:"PartitionQuota,omitempty" xml:"PartitionQuota,omitempty"`
	// The number of used partitions.
	//
	// example:
	//
	// 50
	PartitionUsed *int32 `json:"PartitionUsed,omitempty" xml:"PartitionUsed,omitempty"`
	// The number of available topics.
	//
	// example:
	//
	// 20
	TopicLeft *int32 `json:"TopicLeft,omitempty" xml:"TopicLeft,omitempty"`
	// The number of purchased topics.
	//
	// example:
	//
	// 50
	TopicNumOfBuy *int32 `json:"TopicNumOfBuy,omitempty" xml:"TopicNumOfBuy,omitempty"`
	// The quota of topics.
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
	// The number of used topics.
	//
	// example:
	//
	// 30
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
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-0pp1954n****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance to which the topics that you want to query belong.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic that you want to query.
	//
	// example:
	//
	// topic_name
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C0D3DC5B-5C37-47AD-9F22-1F559880****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The topics.
	TopicList *GetTopicListResponseBodyTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
	// The number of topics.
	//
	// example:
	//
	// 1
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
	// Indicates whether the topic was automatically created.
	//
	// example:
	//
	// false
	AutoCreate *bool `json:"AutoCreate,omitempty" xml:"AutoCreate,omitempty"`
	// The log cleanup policy for the topic. This parameter is returned only if **LocalTopic*	- is set to **true**. Valid values:
	//
	// 	- false: the default log cleanup policy.
	//
	// 	- true: the Apache Kafka log compaction policy.
	//
	// example:
	//
	// false
	CompactTopic *bool `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	// The timestamp that indicates when the topic was created. Unit: milliseconds.
	//
	// example:
	//
	// 1576563109000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// alikafka_pre-cn-0pp1954n****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The storage type that is used for the topic. Valid values:
	//
	// 	- false: cloud storage
	//
	// 	- true: local storage
	//
	// example:
	//
	// false
	LocalTopic *bool `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	// The number of partitions in the topic.
	//
	// example:
	//
	// 6
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The ID of the region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic description. Valid values:
	//
	// 	- The description can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The description must be 3 to 64 characters in length.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The topic status. Valid value:
	//
	// **0**: running.
	//
	// If the topic is deleted, this parameter is not returned.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The topic status. Valid value:
	//
	// **Running**.
	//
	// If the topic is deleted, this parameter is not returned.
	//
	// example:
	//
	// Running
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The tags.
	Tags *GetTopicListResponseBodyTopicListTopicVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The topic name. Valid values:
	//
	// 	- The name can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be 3 to 64 characters in length. If the name contains more than 64 characters, the system automatically truncates the name.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The topic configuration.
	//
	// example:
	//
	// {"replication-factor":3}
	TopicConfig *string `json:"TopicConfig,omitempty" xml:"TopicConfig,omitempty"`
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

func (s *GetTopicListResponseBodyTopicListTopicVO) SetTopicConfig(v string) *GetTopicListResponseBodyTopicListTopicVO {
	s.TopicConfig = &v
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
	//
	// example:
	//
	// Test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Test
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
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h15tjm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// normal_topic_9d034262835916103455551be06cc****
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E475C7E2-8C35-46EF-BE7D-5D2A9F5D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	//
	// example:
	//
	// 1566470063575
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// The information about offsets in the topic.
	OffsetTable *GetTopicStatusResponseBodyTopicStatusOffsetTable `json:"OffsetTable,omitempty" xml:"OffsetTable,omitempty" type:"Struct"`
	// The number of messages in the topic.
	//
	// example:
	//
	// 423
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
	//
	// example:
	//
	// 1566470063547
	LastUpdateTimestamp *int64 `json:"LastUpdateTimestamp,omitempty" xml:"LastUpdateTimestamp,omitempty"`
	// The latest offset in the partition of the topic.
	//
	// example:
	//
	// 76
	MaxOffset *int64 `json:"MaxOffset,omitempty" xml:"MaxOffset,omitempty"`
	// The earliest offset in the partition of the topic.
	//
	// example:
	//
	// 0
	MinOffset *int64 `json:"MinOffset,omitempty" xml:"MinOffset,omitempty"`
	// The ID of the partition.
	//
	// example:
	//
	// 0
	Partition *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// testkafka
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
	// The instance ID.
	//
	// You can call the [GetInstanceList](https://help.aliyun.com/document_detail/437663.html) operation to query the list of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	//
	// You can call the [GetTopicList](https://help.aliyun.com/document_detail/437677.html) operation to query the list of topics.
	//
	// This parameter is required.
	//
	// example:
	//
	// topic_name
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The subscription details.
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
	// The groups that subscribe to the topic.
	ConsumerGroups []*string `json:"ConsumerGroups,omitempty" xml:"ConsumerGroups,omitempty" type:"Repeated"`
	// The topic name.
	//
	// example:
	//
	// topic_api_1681624879908
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the resource is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource whose tags you want to query. The resource ID follows the following rules:
	//
	// 	- Instance ID: instanceId
	//
	// 	- Topic ID: Kafka_alikafka_instanceId_topic
	//
	// 	- Group ID: Kafka_alikafka_instanceId_consumerGroup
	//
	// For example, if the instance ID is alikafka_post-cn-v0h1fgs2xxxx, the topic name is test-topic, and the group name is test-consumer-group, the resource IDs are alikafka_post-cn-v0h1fgs2xxxx, Kafka_alikafka_post-cn-v0h1fgs2xxxx_test-topic, and Kafka_alikafka_post-cn-v0h1fgs2xxxx_test-consumer-group, respectively.
	//
	// >  You must configure one of **ResourceId*	- and **Tag*	- to query the tags that are bound to a resource. Otherwise, the request fails.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource whose tags you want to query. The value is an enumerated value. Valid values:
	//
	// 	- **INSTANCE**
	//
	// 	- **TOPIC**
	//
	// 	- **CONSUMERGROUP**
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
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
	// The key of the resource tag.
	//
	// 	- If you leave this parameter empty, the keys of all tags are matched.
	//
	// 	- The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the resource tag.
	//
	// 	- If you leave Key empty, you must also leave this parameter empty. If you leave this parameter empty, the values of all tags are matched.
	//
	// 	- The tag value can be up to 128 characters in length and cannot contain http:// or https://. The tag value cannot start with acs: or aliyun.
	//
	// example:
	//
	// FinanceJoshua
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
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DE65F6B7-7566-4802-9007-96F2494A****
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
	// 	- The resource ID of an instance is the value of the instanceId parameter.
	//
	// 	- The resource ID of a topic is the value of the Kafka_alikafka_instanceId_topic parameter.
	//
	// 	- The resource ID of a consumer group is the value of the Kafka_alikafka_instanceId_consumerGroup parameter.
	//
	// For example, the resources whose tags you want to query include the alikafka_post-cn-v0h1fgs2xxxx instance, the test-topic topic, and the test-consumer-group consumer group. In this case, their resource IDs are alikafka_post-cn-v0h1fgs2xxxx, Kafka_alikafka_post-cn-v0h1fgs2xxxx_test-topic, and Kafka_alikafka_post-cn-v0h1fgs2xxxx_test-consumer-group.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. The value is an enumerated value. Valid values:
	//
	// 	- **Instance**
	//
	// 	- **Topic**
	//
	// 	- **Consumergroup**
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// FinanceDept
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// FinanceJoshua
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
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name. Valid values:
	//
	// 	- The name can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be 3 to 64 characters in length. A name that contains more than 64 characters is automatically truncated.
	//
	// This parameter is required.
	//
	// example:
	//
	// dev-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
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
	// 	- The value must be an integer that is greater than 0.
	//
	// 	- To reduce the risk of data skew, we recommend that you set the value to a multiple of 6.
	//
	// 	- The number of total partitions ranges from 1 to 360.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	AddPartitionNum *int32 `json:"AddPartitionNum,omitempty" xml:"AddPartitionNum,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// TopicPartitionNum
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B7A39AE5-0B36-4442-A304-E088526****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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

type ModifyScheduledScalingRuleRequest struct {
	// Specifies whether to enable the scheduled scaling rule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If the scaling task is scheduled to execute only once and you want to enable the scheduled scaling rule, make sure that the value of this parameter is at least 30 minutes later than the current point in time.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_serverless-cn-vxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the scheduled scaling rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// contact-id
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifyScheduledScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledScalingRuleRequest) SetEnable(v bool) *ModifyScheduledScalingRuleRequest {
	s.Enable = &v
	return s
}

func (s *ModifyScheduledScalingRuleRequest) SetInstanceId(v string) *ModifyScheduledScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyScheduledScalingRuleRequest) SetRegionId(v string) *ModifyScheduledScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyScheduledScalingRuleRequest) SetRuleName(v string) *ModifyScheduledScalingRuleRequest {
	s.RuleName = &v
	return s
}

type ModifyScheduledScalingRuleResponseBody struct {
	// The response code.
	//
	// The value **200*	- indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DB6F1BEA-903B-4FD8-8809-46E7E9CE***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyScheduledScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledScalingRuleResponseBody) SetCode(v int64) *ModifyScheduledScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponseBody) SetMessage(v string) *ModifyScheduledScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponseBody) SetRequestId(v string) *ModifyScheduledScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponseBody) SetSuccess(v bool) *ModifyScheduledScalingRuleResponseBody {
	s.Success = &v
	return s
}

type ModifyScheduledScalingRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScheduledScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScheduledScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduledScalingRuleResponse) SetHeaders(v map[string]*string) *ModifyScheduledScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduledScalingRuleResponse) SetStatusCode(v int32) *ModifyScheduledScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScheduledScalingRuleResponse) SetBody(v *ModifyScheduledScalingRuleResponseBody) *ModifyScheduledScalingRuleResponse {
	s.Body = v
	return s
}

type ModifyTopicRemarkRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the topic.
	//
	// example:
	//
	// testremark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-0pp1l9z8****
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DB6F1BEA-903B-4FD8-8809-46E7E9CE***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1672410180000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The consumer offset of the partition.
	//
	// example:
	//
	// 100
	Offset *string `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 0
	Partition *string `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The query type. Valid values:
	//
	// 	- byOffset: queries messages by offset. If you select this value, you must configure Partition and Offset.
	//
	// 	- byTimestamp: queries messages by time. If you select this value, you must configure BeginTime.
	//
	// This parameter is required.
	//
	// example:
	//
	// byTimestamp
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The ID of the region where the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testkafka
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	// The returned HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The messages.
	MessageList []*QueryMessageResponseBodyMessageList `json:"MessageList,omitempty" xml:"MessageList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// The check value of the chaincode.
	//
	// example:
	//
	// 0
	Checksum *int64 `json:"Checksum,omitempty" xml:"Checksum,omitempty"`
	// The message key.
	//
	// example:
	//
	// this is key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Indicates whether the key is truncated.
	//
	// example:
	//
	// false
	KeyTruncated *bool `json:"KeyTruncated,omitempty" xml:"KeyTruncated,omitempty"`
	// The consumer offset of the partition.
	//
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 0
	Partition *int64 `json:"Partition,omitempty" xml:"Partition,omitempty"`
	// The size of the key after serialization. Unit: bytes.
	//
	// example:
	//
	// 11
	SerializedKeySize *int32 `json:"SerializedKeySize,omitempty" xml:"SerializedKeySize,omitempty"`
	// The size of the value after serialization. Unit: bytes.
	//
	// example:
	//
	// 20
	SerializedValueSize *int32 `json:"SerializedValueSize,omitempty" xml:"SerializedValueSize,omitempty"`
	// The time when the message was created. The value of this parameter is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1705482172640
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The time type.
	//
	// example:
	//
	// CreateTime
	TimestampType *string `json:"TimestampType,omitempty" xml:"TimestampType,omitempty"`
	// The topic name.
	//
	// example:
	//
	// dqc_test2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The truncated size of the message key. Unit: bytes.
	//
	// >  A maximum of 1 KB of content can be displayed for each message. Content that exceeds 1 KB is automatically truncated. For more information, see [Query messages](https://help.aliyun.com/document_detail/113172.html).
	//
	// example:
	//
	// 0
	TruncatedKeySize *int32 `json:"TruncatedKeySize,omitempty" xml:"TruncatedKeySize,omitempty"`
	// The truncated size of the message value. Unit: bytes.
	//
	// >  A maximum of 1 KB of content can be displayed for each message. Content that exceeds 1 KB is automatically truncated. For more information, see [Query messages](https://help.aliyun.com/document_detail/113172.html).
	//
	// example:
	//
	// 0
	TruncatedValueSize *int32 `json:"TruncatedValueSize,omitempty" xml:"TruncatedValueSize,omitempty"`
	// The message value.
	//
	// example:
	//
	// Welcome to Ali kafka
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// Indicates whether the value is truncated.
	//
	// example:
	//
	// false
	ValueTruncated *bool `json:"ValueTruncated,omitempty" xml:"ValueTruncated,omitempty"`
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
	// 	- **true**: The physical resources of the instance are immediately released.
	//
	// 	- **false**: The physical resources of the instance are retained for a period of time before they are released.
	//
	// example:
	//
	// false
	ForceDeleteInstance *bool `json:"ForceDeleteInstance,omitempty" xml:"ForceDeleteInstance,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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

type ReopenInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp91inkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReopenInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReopenInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReopenInstanceRequest) SetInstanceId(v string) *ReopenInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReopenInstanceRequest) SetRegionId(v string) *ReopenInstanceRequest {
	s.RegionId = &v
	return s
}

type ReopenInstanceResponseBody struct {
	// The returned HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 252820E1-A2E6-45F2-B4C9-1056B8CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReopenInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReopenInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReopenInstanceResponseBody) SetCode(v int32) *ReopenInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ReopenInstanceResponseBody) SetMessage(v string) *ReopenInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ReopenInstanceResponseBody) SetRequestId(v string) *ReopenInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReopenInstanceResponseBody) SetSuccess(v bool) *ReopenInstanceResponseBody {
	s.Success = &v
	return s
}

type ReopenInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReopenInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReopenInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReopenInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReopenInstanceResponse) SetHeaders(v map[string]*string) *ReopenInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReopenInstanceResponse) SetStatusCode(v int32) *ReopenInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReopenInstanceResponse) SetBody(v *ReopenInstanceResponseBody) *ReopenInstanceResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	// The initial configurations of the ApsaraMQ for Kafka instance. The values must be valid JSON strings. If you do not specify this parameter, it is left empty.
	//
	// > - You cannot configure this parameter when you deploy an ApsaraMQ for Confluent instance.
	//
	// > - You cannot configure enable.acl for instances whose versions are earlier than 2.2.0.
	//
	// The **Config*	- parameter supports the following parameters:
	//
	// 	- **enable.vpc_sasl_ssl**: specifies whether to enable VPC transmission encryption. Valid values:
	//
	//     	- **true**: enables VPC transmission encryption. If you enable VPC transmission encryption, you must also enable access control list (ACL).
	//
	//     	- **false**: disables VPC transmission encryption. This is the default value.
	//
	// 	- **enable.acl**: specifies whether to enable ACL. Valid values:
	//
	//     	- **true**: enables ACL.
	//
	//     	- **false**: disables the ACL feature. This is the default value.
	//
	// 	- **kafka.log.retention.hours**: the maximum message retention period when the disk capacity is sufficient. Unit: hours. Valid values: 24 to 480. Default value: **72**. When the disk usage reaches 85%, the disk capacity is insufficient. In this case, the system deletes the earliest stored messages to ensure service availability.
	//
	// 	- **kafka.message.max.bytes**: the maximum size of a message that can be sent and received by ApsaraMQ for Kafka. Unit: bytes. Valid values: 1048576 to 10485760. Default value: **1048576**. Before you change the maximum message size to a new value, make sure that the new value matches the configurations of the producers and consumers.
	//
	// example:
	//
	// {"kafka.log.retention.hours":"33"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether cross-zone deployment is required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// false
	CrossZone *bool `json:"CrossZone,omitempty" xml:"CrossZone,omitempty"`
	// The deployment mode. If the instance is an ApsaraMQ for Kafka V2 instance, this parameter is required. If the instance is an ApsaraMQ for Kafka V3 instance or an ApsaraMQ for Confluent instance, this parameter is optional. Valid values:
	//
	// 	- **vpc**: deploys the instance in a virtual private cloud (VPC).
	//
	// 	- **eip**: deploys the instance over the Internet and in the VPC.
	//
	// The deployment mode of the ApsaraMQ for Kafka instance must be consistent with the instance type. If the instance is a VPC-connected instance, set this parameter to **vpc**. If the instance is an Internet- and VPC-connected instance, set this parameter to **eip**.
	//
	// example:
	//
	// vpc
	DeployModule *string `json:"DeployModule,omitempty" xml:"DeployModule,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the instance supports elastic IP addresses (EIPs). Valid values:
	//
	// 	- **true**: supports EIPs and allows access from the Internet and a VPC.
	//
	// 	- **false**: does not support EIPs and allows access only from a VPC.
	//
	// The value of this parameter must match the type of the instance. For example, if the instance allows access only from a VPC, set this parameter to **false**.
	//
	// example:
	//
	// false
	IsEipInner *bool `json:"IsEipInner,omitempty" xml:"IsEipInner,omitempty"`
	// Specifies whether to forcibly deploy the instance in the selected zones.
	//
	// example:
	//
	// false
	IsForceSelectedZones *bool `json:"IsForceSelectedZones,omitempty" xml:"IsForceSelectedZones,omitempty"`
	// Specifies whether to set a new username and password. Valid values:
	//
	// 	- **true**: sets a new username and password.
	//
	// 	- **false**: does not set a new username or password.
	//
	// This parameter is available only if you deploy an instance that allows access from the Internet and a VPC.
	//
	// example:
	//
	// false
	IsSetUserAndPassword *bool `json:"IsSetUserAndPassword,omitempty" xml:"IsSetUserAndPassword,omitempty"`
	// The ID of the key that is used for disk encryption in the region where the instance is deployed. You can obtain the ID of the key in the [Key Management Service (KMS) console](https://kms.console.aliyun.com/?spm=a2c4g.11186623.2.5.336745b8hfiU21) or create a key. For more information, see [Manage CMKs](https://help.aliyun.com/document_detail/181610.html).
	//
	// If this parameter is configured, disk encryption is enabled for the instance. You cannot disable disk encryption after disk encryption is enabled. When you call this operation, the system checks whether the AliyunServiceRoleForAlikafkaInstanceEncryption service-linked role is created. If the role is not created, the system automatically creates the role. For more information, see [Service-linked roles](https://help.aliyun.com/document_detail/190460.html).
	//
	// > When you deploy a serverless ApsaraMQ for Kafka V3 instance, you cannot configure this parameter.
	//
	// example:
	//
	// 0d24xxxx-da7b-4786-b981-9a164dxxxxxx
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The name of the instance.
	//
	// >  If you specify a value for this parameter, make sure that the specified value is unique in the region of the instance.
	//
	// example:
	//
	// newInstanceName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The alert contact.
	//
	// example:
	//
	// Mr. Wang
	Notifier *string `json:"Notifier,omitempty" xml:"Notifier,omitempty"`
	// The instance password.
	//
	// 	- This parameter is available only for Internet- and VPC- connected ApsaraMQ for Kafka V2 and V3 instances.
	//
	// 	- If the instance is an ApsaraMQ for Confluent instance, this parameter is required. The value of this parameter must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported: ! @ # $ % ^ & \\	- () _ + - =
	//
	// example:
	//
	// password
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group of the instance.
	//
	// If you do not specify this parameter, ApsaraMQ for Kafka automatically configures a security group for your instance. If you specify this parameter, you must create a security group in advance. For more information, see [Create a security group](https://help.aliyun.com/document_detail/25468.html).
	//
	// example:
	//
	// sg-bp13wfx7kz9pko****
	SecurityGroup *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	// The two-dimensional arrays that consist of the candidate set for primary zones and the candidate set for secondary zones. Custom code in the `zone {zone}` format and standard code in the `cn-RegionID-{zone}` format are supported.
	//
	// 	- If you set CrossZone to true and specify Zone H and Zone F as the candidate set for primary zones and Zone K as the candidate set for secondary zones, set this parameter to `[["zoneh","zonef"],["zonek"]]`.
	//
	// > If you specify multiple zones as the primary or secondary zones, the system deploys the instance in one of the zones without prioritizing them. For example, if you set this parameter to `[["zoneh","zonef"],["zonek"]]`, the primary zone in which the instance is deployed can be Zone H or Zone F, and the secondary zone is Zone K.
	//
	// 	- If you set CrossZone to false and want to deploy the instance in Zone K, set this parameter to `[["zonek"],[]]`. In this case, the value of this parameter must still be two-dimensional arrays, but the array that specifies the candidate for secondary zones is left empty.
	//
	// example:
	//
	// [[\\"zonel\\"],[\\"zonek\\"]]
	SelectedZones *string `json:"SelectedZones,omitempty" xml:"SelectedZones,omitempty"`
	// The version of the ApsaraMQ for Kafka instance. Valid values:
	//
	// 	- ApsaraMQ for Kafka V2 instances: 2.2.0 and 2.6.2.
	//
	// 	- ApsaraMQ for Kafka V3 instances: 3.3.1.
	//
	// 	- ApsaraMQ for Confluent instances: 7.4.0.
	//
	// Default value:
	//
	// 	- ApsaraMQ for Kafka V2 instances: 2.2.0.
	//
	// 	- ApsaraMQ for Kafka V3 instances: 3.3.1.
	//
	// 	- ApsaraMQ for Confluent instances: 7.4.0.
	//
	// example:
	//
	// ApsaraMQ for Kafka V2 instances: 2.2.0
	//
	// ApsaraMQ for Kafka V3 instances: 3.3.1
	//
	// ApsaraMQ for Confluent instances: 7.4.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The mobile phone number of the alert contact.
	//
	// example:
	//
	// 1581234****
	UserPhoneNum *string `json:"UserPhoneNum,omitempty" xml:"UserPhoneNum,omitempty"`
	// The instance username.
	//
	// 	- This parameter is available only for Internet- and VPC- connected ApsaraMQ for Kafka V2 and V3 instances.
	//
	// 	- If the instance is an ApsaraMQ for Confluent instance, set this parameter to root or leave this parameter empty.
	//
	// Default value for ApsaraMQ for Kafka V2 and V3 instances: username. Default value for ApsaraMQ for Confluent instances: root.
	//
	// example:
	//
	// username
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The ID of the vSwitch to which you want to connect the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1j3sg5979fstnpl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The IDs of the vSwitches with which the instance is associated. If the instance is an ApsaraMQ for Kafka V2 or V3 instance, this parameter is required. If the instance is an ApsaraMQ for Confluent instance, you must configure one of VSwitchIds and VSwitchId. If you configure both of the parameters, the value of VSwitchIds takes effect.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) in which you want to deploy the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1r4eg3yrxmygv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone where you want to deploy the ApsaraMQ for Kafka instance.
	//
	// 	- The zone ID of the ApsaraMQ for Kafka instance must be the same as that of the vSwitch.
	//
	// 	- The value must be in the zoneX or Region ID-X format. Examples: zonea and cn-hangzhou-k.
	//
	// >  If resources in the specified zone is insufficient, the instance may be deployed in another zone.
	//
	// example:
	//
	// cn-hangzhou-k
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

func (s *StartInstanceRequest) SetVSwitchIds(v []*string) *StartInstanceRequest {
	s.VSwitchIds = v
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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

type StopInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetRegionId(v string) *StopInstanceRequest {
	s.RegionId = &v
	return s
}

type StopInstanceResponseBody struct {
	// The returned status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17D425C2-4EA3-4AB8-928D-E10511ECF***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetCode(v int32) *StopInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopInstanceResponseBody) SetMessage(v string) *StopInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBody) SetSuccess(v bool) *StopInstanceResponseBody {
	s.Success = &v
	return s
}

type StopInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The ID of the Message Queue for Apache RocketMQ instance which contains the resource to which you want to attach tags.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region in which the resource is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resources. The value is an enumerated value. Valid values:
	//
	// 	- **INSTANCE**
	//
	// 	- **TOPIC**
	//
	// 	- **CONSUMERGROUP**
	//
	// >  The value of this parameter is not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	//
	// This parameter is required.
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
	// 	- You must specify this parameter.
	//
	// 	- The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// 	- You can leave this parameter empty.
	//
	// 	- The tag value can be up to 128 characters in length and cannot contain http:// or https://. The tag value cannot start with acs: or aliyun.
	//
	// example:
	//
	// FinanceJoshua
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
	//
	// example:
	//
	// C46FF5A8-C5F0-4024-8262-B16B6392****
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
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The ID of the region in which the resource is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources from which you want to detach tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resources. Valid values:
	//
	// 	- **INSTANCE**
	//
	// 	- **TOPIC**
	//
	// 	- **CONSUMERGROUP**
	//
	// >  The value of this parameter is not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the resource tag.
	//
	// example:
	//
	// FinanceDept
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
	//
	// example:
	//
	// C46FF5A8-C5F0-4024-8262-B16B6392****
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
	// 	- If the **UpdateType*	- parameter is set to **add**, specify one or more IP addresses for this parameter. Separate multiple IP addresses with commas (,).
	//
	// 	- If the **UpdateType*	- parameter is set to **delete**, specify only one IP address.
	//
	// 	- Exercise caution when you delete IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	AllowedListIp *string `json:"AllowedListIp,omitempty" xml:"AllowedListIp,omitempty"`
	// The type of the whitelist. Valid values:
	//
	// 	- **vpc**: a whitelist for access from a VPC.
	//
	// 	- **internet**: a whitelist for access from the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	AllowedListType *string `json:"AllowedListType,omitempty" xml:"AllowedListType,omitempty"`
	// The description of the whitelist.
	//
	// example:
	//
	// tf-testAccEcsImageConfigBasic3549descriptionChange
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-0pp1cng20***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The port range. Valid values:
	//
	// 	- **9092/9092**: Messages are transmitted in a virtual private cloud (VPC) by using the PLAINTEXT protocol.
	//
	// 	- **9093/9093**: Messages are transmitted over the Internet by using the SASL_SSL protocol.
	//
	// 	- **9094/9094**: Messages are transmitted in a VPC by using the SASL_PLAINTEXT protocol.
	//
	// 	- **9095/9095**: Messages are transmitted in a VPC by using the SASL_SSL protocol.
	//
	// This parameter must correspond to **AllowdedListType**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9092/9092
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of configuration change. Valid values:
	//
	// 	- **add**
	//
	// 	- **delete**
	//
	// This parameter is required.
	//
	// example:
	//
	// add
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 17D425C2-4EA3-4AB8-928D-E10511ECF***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// 	- The name can contain letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be **3 to 64*	- characters in length. If a name contains more than **64*	- characters, the name is automatically truncated.
	//
	// 	- The name of a consumer group cannot be changed after the consumer group is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp91inkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// If you set resetType to offset, you can use this parameter to reset the consumer offset of each partition of a specific topic in the consumer group.
	//
	// if can be null:
	// true
	Offsets []*UpdateConsumerOffsetRequestOffsets `json:"Offsets,omitempty" xml:"Offsets,omitempty" type:"Repeated"`
	// The region ID of the instance to which the consumer group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method that is used to reset the consumer offsets of the subscribed topics of a consumer group. Valid values:
	//
	// 	- **timestamp*	- (default)
	//
	// 	- **offset**
	//
	// example:
	//
	// timestamp
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The point in time when message consumption starts. The value of this parameter is a UNIX timestamp in milliseconds. The value of this parameter must be **less than 0*	- or **within the retention period of the consumer offset**. This parameter takes effect only if you set resetType to timestamp.
	//
	// 	- If you want to reset the consumer offset to the latest offset, set this parameter to -1.
	//
	// 	- If you want to reset the consumer offset to the earliest offset, set this parameter to -2.
	//
	// example:
	//
	// -1
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The topic name.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be **3 to 64*	- characters in length. If a name contains more than **64*	- characters, the name is automatically truncated.
	//
	// 	- The name of a topic cannot be changed after the topic is created.
	//
	// **If you want to reset the consumer offsets of all topics to which the consumer subscribes, specify an empty string.**
	//
	// This parameter is required.
	//
	// example:
	//
	// topic_name
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
	//
	// example:
	//
	// 1
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The partition ID.
	//
	// example:
	//
	// 0
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
	// 	- The name can contain letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be **3 to 64*	- characters in length. If a name contains more than **64*	- characters, the name is automatically truncated.
	//
	// 	- The name of a consumer group cannot be changed after the consumer group is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// kafka-test
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp91inkw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// If you set resetType to offset, you can use this parameter to reset the consumer offset of each partition of a specific topic in the consumer group.
	//
	// if can be null:
	// true
	OffsetsShrink *string `json:"Offsets,omitempty" xml:"Offsets,omitempty"`
	// The region ID of the instance to which the consumer group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The method that is used to reset the consumer offsets of the subscribed topics of a consumer group. Valid values:
	//
	// 	- **timestamp*	- (default)
	//
	// 	- **offset**
	//
	// example:
	//
	// timestamp
	ResetType *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	// The point in time when message consumption starts. The value of this parameter is a UNIX timestamp in milliseconds. The value of this parameter must be **less than 0*	- or **within the retention period of the consumer offset**. This parameter takes effect only if you set resetType to timestamp.
	//
	// 	- If you want to reset the consumer offset to the latest offset, set this parameter to -1.
	//
	// 	- If you want to reset the consumer offset to the earliest offset, set this parameter to -2.
	//
	// example:
	//
	// -1
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The topic name.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be **3 to 64*	- characters in length. If a name contains more than **64*	- characters, the name is automatically truncated.
	//
	// 	- The name of a topic cannot be changed after the topic is created.
	//
	// **If you want to reset the consumer offsets of all topics to which the consumer subscribes, specify an empty string.**
	//
	// This parameter is required.
	//
	// example:
	//
	// topic_name
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
	// The HTTP status code that is returned. The status code **200*	- indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 56729737-C428-4E1B-AC68-7A8C2D5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
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
	//
	// This parameter is required.
	//
	// example:
	//
	// {"kafka.log.retention.hours":"33"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B6D821D-7F67-4CAA-9E13-A5A997C35***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful.
	//
	// example:
	//
	// true
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
	// The key of the topic configuration.
	//
	// 	- For reserved instances, you can modify only the configurations of the topics that use local storage.
	//
	// 	- For serverless instances, you can modify the configurations of all topics.
	//
	// 	- Reserved instances whose topics use local storage support the following keys: retention.ms, max.message.bytes, replications, message.timestamp.type, and message.timestamp.difference.max.ms.``
	//
	// 	- Serverless instances support the following keys: retention.hours, max.message.bytes, message.timestamp.type, message.timestamp.difference.max.ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// replications
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The topic name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dqc_test2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// The value of the topic configuration.
	//
	// 	- Serverless instances support the following values:
	//
	//     	- `retention.hours`: the message retention period. Value type: string. Valid values: 24 to 8760.
	//
	//     	- `max.message.bytes`: the maximum size of a sent message. Value type: string. Valid values: 1048576 to 10485760.
	//
	//     	- `message.timestamp.type`: the timestamp type of a message. Valid values: CreateTime and LogAppendTime. The value CreateTime specifies the timestamp that is specified by the producer when the message is sent. The value LogAppendTime specifies the time when the broker appends the message to its log. If you do not specify this parameter, the time when the message is created on the client is used.
	//
	//     	- `message.timestamp.difference.max.ms`: the maximum positive or negative difference allowed between the timestamp when the broker receives a message and the timestamp specified in the message. If you set message.timestamp.type to CreateTime, **a message is rejected*	- if the difference in timestamp exceeds the threshold. If you set message.timestamp.type to LogAppendTime, this configuration does not take effect.
	//
	// 	- Reserved instances support the following values:
	//
	//     	- `retention.ms`: the message retention period. Value type: string. Valid values: 3600000 to 31536000000.
	//
	//     	- `max.message.bytes`: the maximum size of a sent message. Value type: string. Valid values: 1048576 to 10485760.
	//
	//     	- `replications`: the number of replicas. Value type: string. Valid values: 1 to 3.
	//
	//     	- `message.timestamp.type`: the timestamp type of a message. Valid values: CreateTime and LogAppendTime. The value CreateTime specifies the timestamp that is specified by the producer when the message is sent. The value LogAppendTime specifies the time when the broker appends the message to its log. If you do not specify this parameter, the time when the message is created on the client is used.
	//
	//     	- `message.timestamp.difference.max.ms`: the maximum positive or negative difference allowed between the timestamp when the broker receives a message and the timestamp specified in the message. If you set message.timestamp.type to CreateTime, **a message is rejected*	- if the difference in timestamp exceeds the threshold. If you set message.timestamp.type to LogAppendTime, this configuration does not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	// The HTTP status code. If the request is successful, 200 is returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// []
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0178A3A7-E87B-5E50-A16F-3E62F534****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The major version to be upgraded to. Valid values:
	//
	// 	- **0.10.2**
	//
	// 	- **2.2.0**
	//
	// If you set this parameter to the current major version, the system upgrades the instance to the latest minor version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.10.2
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	// 	- The disk size that you specify must be greater than or equal to the current disk size of the instance.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The maximum Internet traffic of the instance.
	//
	// 	- The Internet traffic that you specify must be greater than or equal to the current Internet traffic of the instance.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >
	//
	// 	- If you set **EipModel*	- to **true**, set **EipMax*	- to a value that is greater than 0.
	//
	// 	- If you set **EipModel*	- to **false**, set **EipMax*	- to **0**.
	//
	// 	- If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// 	- true: enables Internet access.
	//
	// 	- false: disables Internet access.
	//
	// example:
	//
	// false
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic of the instance. We recommend that you do not configure this parameter.
	//
	// 	- The maximum traffic that you specify must be greater than or equal to the current maximum traffic of the instance.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 60
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.6xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 80
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameters that are configured for the serverless instance. These parameters are required only when you create a serverless instance.
	ServerlessConfig *UpgradePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The instance edition.
	//
	// Valid values for this parameter if you set PaidType to 1:
	//
	// 	- normal: Standard Edition (High Write)
	//
	// 	- professional: Professional Edition (High Write)
	//
	// 	- professionalForHighRead: Professional Edition (High Read)
	//
	// Valid values for this parameter if you set PaidType to 3:
	//
	// 	- normal: Serverless Standard Edition
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you use exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 80
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
	// The reserved capacity for publishing messages. You can specify only an integer for this parameter. Minimum value: 60.
	//
	// >  The maximum capacity that you can reserve for an instance varies based on available resources in a region. The reserved capacity range displayed on the buy page shall prevail.
	//
	// example:
	//
	// 50
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved capacity for subscribing to messages. You can specify only an integer for this parameter. Minimum value: 50.
	//
	// >  The maximum capacity that you can reserve for an instance varies based on available resources in a region. The reserved capacity range displayed on the buy page shall prevail.
	//
	// example:
	//
	// 50
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
	// 	- The disk size that you specify must be greater than or equal to the current disk size of the instance.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The maximum Internet traffic of the instance.
	//
	// 	- The Internet traffic that you specify must be greater than or equal to the current Internet traffic of the instance.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >
	//
	// 	- If you set **EipModel*	- to **true**, set **EipMax*	- to a value that is greater than 0.
	//
	// 	- If you set **EipModel*	- to **false**, set **EipMax*	- to **0**.
	//
	// 	- If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// 	- true: enables Internet access.
	//
	// 	- false: disables Internet access.
	//
	// example:
	//
	// false
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic of the instance. We recommend that you do not configure this parameter.
	//
	// 	- The maximum traffic that you specify must be greater than or equal to the current maximum traffic of the instance.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 60
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.6xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 80
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameters that are configured for the serverless instance. These parameters are required only when you create a serverless instance.
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// The instance edition.
	//
	// Valid values for this parameter if you set PaidType to 1:
	//
	// 	- normal: Standard Edition (High Write)
	//
	// 	- professional: Professional Edition (High Write)
	//
	// 	- professionalForHighRead: Professional Edition (High Read)
	//
	// Valid values for this parameter if you set PaidType to 3:
	//
	// 	- normal: Serverless Standard Edition
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you use exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 80
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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
	ConfluentConfig *UpgradePrePayOrderRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	// The size of the disk.
	//
	// 	- The disk size that you specify must be greater than or equal to the current disk size of the instance.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 900
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The Internet traffic for the instance.
	//
	// 	- The Internet traffic volume that you specify must be greater than or equal to the current Internet traffic volume of the instance.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > - If the **EipModel*	- parameter is set to **true**, set the **EipMax*	- parameter to a value that is greater than 0.
	//
	// > - If the **EipModel*	- parameter is set to **false**, set the **EipMax*	- parameter to **0**.
	//
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// 	- true: enables Internet access.
	//
	// 	- false: disables Internet access.
	//
	// example:
	//
	// true
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic for the instance. We recommend that you do not configure this parameter.
	//
	// 	- The maximum traffic volume that you specify must be greater than or equal to the current maximum traffic volume of the instance.
	//
	// 	- You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you configure only the IoMaxSpec parameter.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 40
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	//
	// 	- You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you configure only the IoMaxSpec parameter.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	PaidType  *int32  `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	//
	// 	- If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 50
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The edition of the instance. Valid values:
	//
	// 	- **normal**: Standard Edition (High Write)
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// You cannot downgrade an instance from the Professional Edition to the Standard Edition. For more information about these instance editions, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	//
	// 	- If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	//
	// 	- The default value of the TopicQuota parameter varies based on the value of the IoMaxSpec parameter. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s UpgradePrePayOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradePrePayOrderRequest) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderRequest) SetConfluentConfig(v *UpgradePrePayOrderRequestConfluentConfig) *UpgradePrePayOrderRequest {
	s.ConfluentConfig = v
	return s
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

func (s *UpgradePrePayOrderRequest) SetPaidType(v int32) *UpgradePrePayOrderRequest {
	s.PaidType = &v
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

type UpgradePrePayOrderRequestConfluentConfig struct {
	ConnectCU             *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	ConnectReplica        *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	ControlCenterCU       *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	ControlCenterReplica  *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	ControlCenterStorage  *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	KafkaCU               *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	KafkaReplica          *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	KafkaRestProxyCU      *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	KafkaStorage          *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	KsqlCU                *int32 `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	KsqlReplica           *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	KsqlStorage           *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	SchemaRegistryCU      *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	ZooKeeperCU           *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	ZooKeeperReplica      *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	ZooKeeperStorage      *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s UpgradePrePayOrderRequestConfluentConfig) String() string {
	return tea.Prettify(s)
}

func (s UpgradePrePayOrderRequestConfluentConfig) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetConnectCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetConnectReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetControlCenterCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetControlCenterReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetControlCenterStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaRestProxyCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaRestProxyReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKsqlCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKsqlReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKsqlStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetSchemaRegistryCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetSchemaRegistryReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetZooKeeperCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetZooKeeperReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetZooKeeperStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

type UpgradePrePayOrderShrinkRequest struct {
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	// The size of the disk.
	//
	// 	- The disk size that you specify must be greater than or equal to the current disk size of the instance.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 900
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The Internet traffic for the instance.
	//
	// 	- The Internet traffic volume that you specify must be greater than or equal to the current Internet traffic volume of the instance.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > - If the **EipModel*	- parameter is set to **true**, set the **EipMax*	- parameter to a value that is greater than 0.
	//
	// > - If the **EipModel*	- parameter is set to **false**, set the **EipMax*	- parameter to **0**.
	//
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// 	- true: enables Internet access.
	//
	// 	- false: disables Internet access.
	//
	// example:
	//
	// true
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum traffic for the instance. We recommend that you do not configure this parameter.
	//
	// 	- The maximum traffic volume that you specify must be greater than or equal to the current maximum traffic volume of the instance.
	//
	// 	- You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you configure only the IoMaxSpec parameter.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 40
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	//
	// 	- You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you configure only the IoMaxSpec parameter.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	PaidType  *int32  `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	//
	// 	- If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 50
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The edition of the instance. Valid values:
	//
	// 	- **normal**: Standard Edition (High Write)
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// You cannot downgrade an instance from the Professional Edition to the Standard Edition. For more information about these instance editions, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	//
	// 	- If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	//
	// 	- The default value of the TopicQuota parameter varies based on the value of the IoMaxSpec parameter. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s UpgradePrePayOrderShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradePrePayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderShrinkRequest) SetConfluentConfigShrink(v string) *UpgradePrePayOrderShrinkRequest {
	s.ConfluentConfigShrink = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetDiskSize(v int32) *UpgradePrePayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetEipMax(v int32) *UpgradePrePayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetEipModel(v bool) *UpgradePrePayOrderShrinkRequest {
	s.EipModel = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetInstanceId(v string) *UpgradePrePayOrderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetIoMax(v int32) *UpgradePrePayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetIoMaxSpec(v string) *UpgradePrePayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetPaidType(v int32) *UpgradePrePayOrderShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetPartitionNum(v int32) *UpgradePrePayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetRegionId(v string) *UpgradePrePayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetSpecType(v string) *UpgradePrePayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetTopicQuota(v int32) *UpgradePrePayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

type UpgradePrePayOrderResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
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

func (s *UpgradePrePayOrderResponseBody) SetOrderId(v string) *UpgradePrePayOrderResponseBody {
	s.OrderId = &v
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

// Summary:
//
// Changes the resource group of an ApsaraMQ for Kafka instance.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the resource group of an ApsaraMQ for Kafka instance.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
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

// Summary:
//
// Changes the billing method of a Message Queue for Apache Kafka instance from pay-as-you-go to subscription.
//
// @param request - ConvertPostPayOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertPostPayOrderResponse
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

	if !tea.BoolValue(util.IsUnset(request.PaidType)) {
		query["PaidType"] = request.PaidType
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ConvertPostPayOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ConvertPostPayOrderResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the billing method of a Message Queue for Apache Kafka instance from pay-as-you-go to subscription.
//
// @param request - ConvertPostPayOrderRequest
//
// @return ConvertPostPayOrderResponse
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

// Summary:
//
// Creates an access control list (ACL).
//
// @param request - CreateAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAclResponse
func (client *Client) CreateAclWithOptions(request *CreateAclRequest, runtime *util.RuntimeOptions) (_result *CreateAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclOperationType)) {
		query["AclOperationType"] = request.AclOperationType
	}

	if !tea.BoolValue(util.IsUnset(request.AclOperationTypes)) {
		query["AclOperationTypes"] = request.AclOperationTypes
	}

	if !tea.BoolValue(util.IsUnset(request.AclPermissionType)) {
		query["AclPermissionType"] = request.AclPermissionType
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

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		query["Host"] = request.Host
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAclResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAclResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates an access control list (ACL).
//
// @param request - CreateAclRequest
//
// @return CreateAclResponse
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

// Summary:
//
// Creates a consumer group.
//
// @param request - CreateConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConsumerGroupResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateConsumerGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateConsumerGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a consumer group.
//
// @param request - CreateConsumerGroupRequest
//
// @return CreateConsumerGroupResponse
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

// Summary:
//
// 创建后付费实例。
//
// @param tmpReq - CreatePostPayInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePostPayInstanceResponse
func (client *Client) CreatePostPayInstanceWithOptions(tmpReq *CreatePostPayInstanceRequest, runtime *util.RuntimeOptions) (_result *CreatePostPayInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePostPayInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ServerlessConfig)) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, tea.String("ServerlessConfig"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.IoMaxSpec)) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !tea.BoolValue(util.IsUnset(request.PaidType)) {
		query["PaidType"] = request.PaidType
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

	if !tea.BoolValue(util.IsUnset(request.ServerlessConfigShrink)) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SpecType)) {
		query["SpecType"] = request.SpecType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePostPayInstance"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePostPayInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePostPayInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建后付费实例。
//
// @param request - CreatePostPayInstanceRequest
//
// @return CreatePostPayInstanceResponse
func (client *Client) CreatePostPayInstance(request *CreatePostPayInstanceRequest) (_result *CreatePostPayInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePostPayInstanceResponse{}
	_body, _err := client.CreatePostPayInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreatePostPayOrderWithOptions(tmpReq *CreatePostPayOrderRequest, runtime *util.RuntimeOptions) (_result *CreatePostPayOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePostPayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ServerlessConfig)) {
		request.ServerlessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServerlessConfig, tea.String("ServerlessConfig"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.PaidType)) {
		query["PaidType"] = request.PaidType
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

	if !tea.BoolValue(util.IsUnset(request.ServerlessConfigShrink)) {
		query["ServerlessConfig"] = request.ServerlessConfigShrink
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePostPayOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePostPayOrderResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a pay-as-you-go ApsaraMQ for Kafka instance. Pay-as-you-go instances allow you to pay after you use the resources. You are charged for pay-as-you-go instances based on the actual resource usage. You can use pay-as-you-go instances in test scenarios or scenarios in which the peak traffic is uncertain.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of pay-as-you-go Message Queue for Apache Kafka instances. For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
//
// @param request - CreatePostPayOrderRequest
//
// @return CreatePostPayOrderResponse
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

// Summary:
//
// 创建预付费实例
//
// @param tmpReq - CreatePrePayInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrePayInstanceResponse
func (client *Client) CreatePrePayInstanceWithOptions(tmpReq *CreatePrePayInstanceRequest, runtime *util.RuntimeOptions) (_result *CreatePrePayInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePrePayInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConfluentConfig)) {
		request.ConfluentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfluentConfig, tea.String("ConfluentConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfluentConfigShrink)) {
		query["ConfluentConfig"] = request.ConfluentConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		query["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.EipMax)) {
		query["EipMax"] = request.EipMax
	}

	if !tea.BoolValue(util.IsUnset(request.IoMaxSpec)) {
		query["IoMaxSpec"] = request.IoMaxSpec
	}

	if !tea.BoolValue(util.IsUnset(request.PaidType)) {
		query["PaidType"] = request.PaidType
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePrePayInstance"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePrePayInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePrePayInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建预付费实例
//
// @param request - CreatePrePayInstanceRequest
//
// @return CreatePrePayInstanceResponse
func (client *Client) CreatePrePayInstance(request *CreatePrePayInstanceRequest) (_result *CreatePrePayInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePrePayInstanceResponse{}
	_body, _err := client.CreatePrePayInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreatePrePayOrderWithOptions(tmpReq *CreatePrePayOrderRequest, runtime *util.RuntimeOptions) (_result *CreatePrePayOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePrePayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConfluentConfig)) {
		request.ConfluentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfluentConfig, tea.String("ConfluentConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfluentConfigShrink)) {
		query["ConfluentConfig"] = request.ConfluentConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		query["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		query["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
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

	if !tea.BoolValue(util.IsUnset(request.PaidType)) {
		query["PaidType"] = request.PaidType
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePrePayOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePrePayOrderResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @param request - CreatePrePayOrderRequest
//
// @return CreatePrePayOrderResponse
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

// Summary:
//
// Creates a Simple Authentication and Security Layer (SASL) user.
//
// @param request - CreateSaslUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSaslUserResponse
func (client *Client) CreateSaslUserWithOptions(request *CreateSaslUserRequest, runtime *util.RuntimeOptions) (_result *CreateSaslUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mechanism)) {
		query["Mechanism"] = request.Mechanism
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateSaslUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateSaslUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a Simple Authentication and Security Layer (SASL) user.
//
// @param request - CreateSaslUserRequest
//
// @return CreateSaslUserResponse
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
func (client *Client) CreateScheduledScalingRuleWithOptions(tmpReq *CreateScheduledScalingRuleRequest, runtime *util.RuntimeOptions) (_result *CreateScheduledScalingRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateScheduledScalingRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.WeeklyTypes)) {
		request.WeeklyTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WeeklyTypes, tea.String("WeeklyTypes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DurationMinutes)) {
		query["DurationMinutes"] = request.DurationMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.FirstScheduledTime)) {
		query["FirstScheduledTime"] = request.FirstScheduledTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatType)) {
		query["RepeatType"] = request.RepeatType
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedPubFlow)) {
		query["ReservedPubFlow"] = request.ReservedPubFlow
	}

	if !tea.BoolValue(util.IsUnset(request.ReservedSubFlow)) {
		query["ReservedSubFlow"] = request.ReservedSubFlow
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleType)) {
		query["ScheduleType"] = request.ScheduleType
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		query["TimeZone"] = request.TimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.WeeklyTypesShrink)) {
		query["WeeklyTypes"] = request.WeeklyTypesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduledScalingRule"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateScheduledScalingRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateScheduledScalingRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a scheduled scaling rule for a serverless ApsaraMQ for Kafka V3 instance.
//
// Description:
//
// ###### [](#-v3-serverless-)This operation is supported only by serverless ApsaraMQ for Kafka V3 instances.
//
// @param request - CreateScheduledScalingRuleRequest
//
// @return CreateScheduledScalingRuleResponse
func (client *Client) CreateScheduledScalingRule(request *CreateScheduledScalingRuleRequest) (_result *CreateScheduledScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduledScalingRuleResponse{}
	_body, _err := client.CreateScheduledScalingRuleWithOptions(request, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTopicResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTopicResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return CreateTopicResponse
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

// Summary:
//
// Deletes an access control list (ACL).
//
// @param request - DeleteAclRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAclResponse
func (client *Client) DeleteAclWithOptions(request *DeleteAclRequest, runtime *util.RuntimeOptions) (_result *DeleteAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclOperationType)) {
		query["AclOperationType"] = request.AclOperationType
	}

	if !tea.BoolValue(util.IsUnset(request.AclOperationTypes)) {
		query["AclOperationTypes"] = request.AclOperationTypes
	}

	if !tea.BoolValue(util.IsUnset(request.AclPermissionType)) {
		query["AclPermissionType"] = request.AclPermissionType
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

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		query["Host"] = request.Host
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAclResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAclResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an access control list (ACL).
//
// @param request - DeleteAclRequest
//
// @return DeleteAclResponse
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

// Summary:
//
// Deletes a consumer group from a specified Message Queue for Apache Kafka instance.
//
// @param request - DeleteConsumerGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConsumerGroupResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteConsumerGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteConsumerGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a consumer group from a specified Message Queue for Apache Kafka instance.
//
// @param request - DeleteConsumerGroupRequest
//
// @return DeleteConsumerGroupResponse
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

// Summary:
//
// Deletes an instance. You can delete subscription and pay-as-you-go instances after you release them.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes an instance. You can delete subscription and pay-as-you-go instances after you release them.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
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

// Summary:
//
// Deletes a Simple Authentication and Security Layer (SASL) user.
//
// @param request - DeleteSaslUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSaslUserResponse
func (client *Client) DeleteSaslUserWithOptions(request *DeleteSaslUserRequest, runtime *util.RuntimeOptions) (_result *DeleteSaslUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Mechanism)) {
		query["Mechanism"] = request.Mechanism
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteSaslUserResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteSaslUserResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a Simple Authentication and Security Layer (SASL) user.
//
// @param request - DeleteSaslUserRequest
//
// @return DeleteSaslUserResponse
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
func (client *Client) DeleteScheduledScalingRuleWithOptions(request *DeleteScheduledScalingRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduledScalingRuleResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScheduledScalingRule"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteScheduledScalingRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteScheduledScalingRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return DeleteScheduledScalingRuleResponse
func (client *Client) DeleteScheduledScalingRule(request *DeleteScheduledScalingRuleRequest) (_result *DeleteScheduledScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduledScalingRuleResponse{}
	_body, _err := client.DeleteScheduledScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteTopicResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteTopicResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Deletes a topic.
//
// @param request - DeleteTopicRequest
//
// @return DeleteTopicResponse
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

// Summary:
//
// 查询acl资源名
//
// @param request - DescribeAclResourceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAclResourceNameResponse
func (client *Client) DescribeAclResourceNameWithOptions(request *DescribeAclResourceNameRequest, runtime *util.RuntimeOptions) (_result *DescribeAclResourceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAclResourceName"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAclResourceNameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAclResourceNameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询acl资源名
//
// @param request - DescribeAclResourceNameRequest
//
// @return DescribeAclResourceNameResponse
func (client *Client) DescribeAclResourceName(request *DescribeAclResourceNameRequest) (_result *DescribeAclResourceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAclResourceNameResponse{}
	_body, _err := client.DescribeAclResourceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeAclsWithOptions(request *DescribeAclsRequest, runtime *util.RuntimeOptions) (_result *DescribeAclsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AclOperationType)) {
		query["AclOperationType"] = request.AclOperationType
	}

	if !tea.BoolValue(util.IsUnset(request.AclPermissionType)) {
		query["AclPermissionType"] = request.AclPermissionType
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

	if !tea.BoolValue(util.IsUnset(request.Host)) {
		query["Host"] = request.Host
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeAclsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeAclsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries access control lists (ACLs).
//
// @param request - DescribeAclsRequest
//
// @return DescribeAclsResponse
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

// Summary:
//
// Queries Simple Authentication and Security Layer (SASL) users.
//
// @param request - DescribeSaslUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSaslUsersResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeSaslUsersResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeSaslUsersResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries Simple Authentication and Security Layer (SASL) users.
//
// @param request - DescribeSaslUsersRequest
//
// @return DescribeSaslUsersResponse
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

// Summary:
//
// Enables and disables the flexible group creation feature.
//
// @param request - EnableAutoGroupCreationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAutoGroupCreationResponse
func (client *Client) EnableAutoGroupCreationWithOptions(request *EnableAutoGroupCreationRequest, runtime *util.RuntimeOptions) (_result *EnableAutoGroupCreationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
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
		Action:      tea.String("EnableAutoGroupCreation"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableAutoGroupCreationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableAutoGroupCreationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Enables and disables the flexible group creation feature.
//
// @param request - EnableAutoGroupCreationRequest
//
// @return EnableAutoGroupCreationResponse
func (client *Client) EnableAutoGroupCreation(request *EnableAutoGroupCreationRequest) (_result *EnableAutoGroupCreationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAutoGroupCreationResponse{}
	_body, _err := client.EnableAutoGroupCreationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EnableAutoTopicCreationWithOptions(request *EnableAutoTopicCreationRequest, runtime *util.RuntimeOptions) (_result *EnableAutoTopicCreationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Operate)) {
		query["Operate"] = request.Operate
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNum)) {
		query["PartitionNum"] = request.PartitionNum
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.UpdatePartition)) {
		query["UpdatePartition"] = request.UpdatePartition
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAutoTopicCreation"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &EnableAutoTopicCreationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &EnableAutoTopicCreationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Enables or disables the automatic topic creation feature, or changes the number of partitions in topics that are automatically created.
//
// @param request - EnableAutoTopicCreationRequest
//
// @return EnableAutoTopicCreationResponse
func (client *Client) EnableAutoTopicCreation(request *EnableAutoTopicCreationRequest) (_result *EnableAutoTopicCreationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAutoTopicCreationResponse{}
	_body, _err := client.EnableAutoTopicCreationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAllInstanceIdListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAllInstanceIdListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the IDs of all instances in the current account.
//
// @param request - GetAllInstanceIdListRequest
//
// @return GetAllInstanceIdListResponse
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

// Summary:
//
// Queries the IP address whitelist.
//
// @param request - GetAllowedIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAllowedIpListResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAllowedIpListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAllowedIpListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the IP address whitelist.
//
// @param request - GetAllowedIpListRequest
//
// @return GetAllowedIpListResponse
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
func (client *Client) GetAutoScalingConfigurationWithOptions(request *GetAutoScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *GetAutoScalingConfigurationResponse, _err error) {
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
		Action:      tea.String("GetAutoScalingConfiguration"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAutoScalingConfigurationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAutoScalingConfigurationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return GetAutoScalingConfigurationResponse
func (client *Client) GetAutoScalingConfiguration(request *GetAutoScalingConfigurationRequest) (_result *GetAutoScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoScalingConfigurationResponse{}
	_body, _err := client.GetAutoScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetConsumerListWithOptions(request *GetConsumerListRequest, runtime *util.RuntimeOptions) (_result *GetConsumerListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerId)) {
		query["ConsumerId"] = request.ConsumerId
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetConsumerListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetConsumerListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries one or more consumer groups in a specified Message Queue for Apache Kafka instance.
//
// @param request - GetConsumerListRequest
//
// @return GetConsumerListResponse
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

// Summary:
//
// Queries the consumer progress of a consumer group.
//
// @param request - GetConsumerProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConsumerProgressResponse
func (client *Client) GetConsumerProgressWithOptions(request *GetConsumerProgressRequest, runtime *util.RuntimeOptions) (_result *GetConsumerProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumerId)) {
		query["ConsumerId"] = request.ConsumerId
	}

	if !tea.BoolValue(util.IsUnset(request.HideLastTimestamp)) {
		query["HideLastTimestamp"] = request.HideLastTimestamp
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetConsumerProgressResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetConsumerProgressResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the consumer progress of a consumer group.
//
// @param request - GetConsumerProgressRequest
//
// @return GetConsumerProgressResponse
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

// Summary:
//
// Queries the information about instances in a specified region.
//
// @param request - GetInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceListResponse
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

	if !tea.BoolValue(util.IsUnset(request.Series)) {
		query["Series"] = request.Series
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about instances in a specified region.
//
// @param request - GetInstanceListRequest
//
// @return GetInstanceListResponse
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
func (client *Client) GetKafkaClientIpWithOptions(request *GetKafkaClientIpRequest, runtime *util.RuntimeOptions) (_result *GetKafkaClientIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["Group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Topic)) {
		query["Topic"] = request.Topic
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetKafkaClientIp"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetKafkaClientIpResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetKafkaClientIpResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return GetKafkaClientIpResponse
func (client *Client) GetKafkaClientIp(request *GetKafkaClientIpRequest) (_result *GetKafkaClientIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKafkaClientIpResponse{}
	_body, _err := client.GetKafkaClientIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetQuotaTipResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetQuotaTipResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the used quota of topics and partitions.
//
// @param request - GetQuotaTipRequest
//
// @return GetQuotaTipResponse
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

// Summary:
//
// Queries the information about a topic.
//
// @param request - GetTopicListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicListResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTopicListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTopicListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about a topic.
//
// @param request - GetTopicListRequest
//
// @return GetTopicListResponse
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

// Summary:
//
// Queries the messaging status of a topic.
//
// @param request - GetTopicStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicStatusResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTopicStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTopicStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the messaging status of a topic.
//
// @param request - GetTopicStatusRequest
//
// @return GetTopicStatusResponse
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

// Summary:
//
// Queries the information about the groups that subscribe to a topic.
//
// @param request - GetTopicSubscribeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicSubscribeStatusResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTopicSubscribeStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTopicSubscribeStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the information about the groups that subscribe to a topic.
//
// @param request - GetTopicSubscribeStatusRequest
//
// @return GetTopicSubscribeStatusResponse
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

// Summary:
//
// Queries the tags that are attached to a specified resource.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the tags that are attached to a specified resource.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Changes the name of an ApsaraMQ for Kafka instance. After you deploy an instance, you can call this operation to change the name of the instance.
//
// @param request - ModifyInstanceNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceNameResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyInstanceNameResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyInstanceNameResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the name of an ApsaraMQ for Kafka instance. After you deploy an instance, you can call this operation to change the name of the instance.
//
// @param request - ModifyInstanceNameRequest
//
// @return ModifyInstanceNameResponse
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

// Summary:
//
// Changes the number of partitions in a topic.
//
// @param request - ModifyPartitionNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPartitionNumResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyPartitionNumResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyPartitionNumResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the number of partitions in a topic.
//
// @param request - ModifyPartitionNumRequest
//
// @return ModifyPartitionNumResponse
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
func (client *Client) ModifyScheduledScalingRuleWithOptions(request *ModifyScheduledScalingRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyScheduledScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScheduledScalingRule"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyScheduledScalingRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyScheduledScalingRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return ModifyScheduledScalingRuleResponse
func (client *Client) ModifyScheduledScalingRule(request *ModifyScheduledScalingRuleRequest) (_result *ModifyScheduledScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScheduledScalingRuleResponse{}
	_body, _err := client.ModifyScheduledScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyTopicRemarkResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyTopicRemarkResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the description of a topic.
//
// @param request - ModifyTopicRemarkRequest
//
// @return ModifyTopicRemarkResponse
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

// Summary:
//
// Queries messages stored in a topic. You can query messages by creation time or offset.
//
// @param request - QueryMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMessageResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &QueryMessageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &QueryMessageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries messages stored in a topic. You can query messages by creation time or offset.
//
// @param request - QueryMessageRequest
//
// @return QueryMessageResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return ReleaseInstanceResponse
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
func (client *Client) ReopenInstanceWithOptions(request *ReopenInstanceRequest, runtime *util.RuntimeOptions) (_result *ReopenInstanceResponse, _err error) {
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
		Action:      tea.String("ReopenInstance"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReopenInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReopenInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return ReopenInstanceResponse
func (client *Client) ReopenInstance(request *ReopenInstanceRequest) (_result *ReopenInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReopenInstanceResponse{}
	_body, _err := client.ReopenInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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

	if !tea.BoolValue(util.IsUnset(request.VSwitchIds)) {
		query["VSwitchIds"] = request.VSwitchIds
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StartInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StartInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return StartInstanceResponse
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
func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
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
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &StopInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &StopInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return StopInstanceResponse
func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Attaches a tag to a resource.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// Detaches tags from a specified resource.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Detaches tags from a specified resource.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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

// Summary:
//
// Updates the IP address whitelist of an ApsaraMQ for Kafka instance. Only IP addresses and ports that are configured in the IP address whitelist of an instance can access the instance.
//
// @param request - UpdateAllowedIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAllowedIpResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateAllowedIpResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateAllowedIpResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Updates the IP address whitelist of an ApsaraMQ for Kafka instance. Only IP addresses and ports that are configured in the IP address whitelist of an instance can access the instance.
//
// @param request - UpdateAllowedIpRequest
//
// @return UpdateAllowedIpResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateConsumerOffsetResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateConsumerOffsetResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @param request - UpdateConsumerOffsetRequest
//
// @return UpdateConsumerOffsetResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return UpdateInstanceConfigResponse
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

// Summary:
//
// Modifies the configurations of a topic. After you create a topic, you can modify the message retention period and maximum message size of the topic.
//
// @param request - UpdateTopicConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTopicConfigResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTopicConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTopicConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Modifies the configurations of a topic. After you create a topic, you can modify the message retention period and maximum message size of the topic.
//
// @param request - UpdateTopicConfigRequest
//
// @return UpdateTopicConfigResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeInstanceVersionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeInstanceVersionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// @return UpgradeInstanceVersionResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradePostPayOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradePostPayOrderResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Upgrades a pay-as-you-go ApsaraMQ for Kafka instance.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of pay-as-you-go Message Queue for Apache Kafka instances. For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
//
// @param request - UpgradePostPayOrderRequest
//
// @return UpgradePostPayOrderResponse
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
func (client *Client) UpgradePrePayOrderWithOptions(tmpReq *UpgradePrePayOrderRequest, runtime *util.RuntimeOptions) (_result *UpgradePrePayOrderResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpgradePrePayOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConfluentConfig)) {
		request.ConfluentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfluentConfig, tea.String("ConfluentConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfluentConfigShrink)) {
		query["ConfluentConfig"] = request.ConfluentConfigShrink
	}

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

	if !tea.BoolValue(util.IsUnset(request.PaidType)) {
		query["PaidType"] = request.PaidType
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradePrePayOrderResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradePrePayOrderResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Upgrades a Message Queue for Apache Kafka instance that uses the subscription billing method.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing method and pricing of subscription Message Queue for Apache Kafka instances. For more information, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
//
// @param request - UpgradePrePayOrderRequest
//
// @return UpgradePrePayOrderResponse
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
