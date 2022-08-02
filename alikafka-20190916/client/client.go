// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeResourceGroupRequest struct {
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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
	Code               *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *int64  `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Duration   *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConvertPostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclOperationType       *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	AclResourceName        *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	AclResourceType        *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Username               *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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

type CreateConsumerGroupResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DeployType      *int32  `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DiskSize        *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType        *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	EipMax          *int32  `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	IoMax           *int32  `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxSpec       *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SpecType        *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TopicQuota      *int32  `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
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

func (s *CreatePostPayOrderRequest) SetTopicQuota(v int32) *CreatePostPayOrderRequest {
	s.TopicQuota = &v
	return s
}

type CreatePostPayOrderResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DeployType      *int32  `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DiskSize        *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType        *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	EipMax          *int32  `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	IoMax           *int32  `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxSpec       *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SpecType        *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TopicQuota      *int32  `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
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

func (s *CreatePrePayOrderRequest) SetTopicQuota(v int32) *CreatePrePayOrderRequest {
	s.TopicQuota = &v
	return s
}

type CreatePrePayOrderResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePrePayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Password   *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CompactTopic      *bool                  `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	Config            map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	InstanceId        *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LocalTopic        *bool                  `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	MinInsyncReplicas *int64                 `json:"MinInsyncReplicas,omitempty" xml:"MinInsyncReplicas,omitempty"`
	PartitionNum      *string                `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	RegionId          *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark            *string                `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ReplicationFactor *int64                 `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	Topic             *string                `json:"Topic,omitempty" xml:"Topic,omitempty"`
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

func (s *CreateTopicRequest) SetConfig(v map[string]interface{}) *CreateTopicRequest {
	s.Config = v
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

func (s *CreateTopicRequest) SetTopic(v string) *CreateTopicRequest {
	s.Topic = &v
	return s
}

type CreateTopicShrinkRequest struct {
	CompactTopic      *bool   `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	ConfigShrink      *string `json:"Config,omitempty" xml:"Config,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LocalTopic        *bool   `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	MinInsyncReplicas *int64  `json:"MinInsyncReplicas,omitempty" xml:"MinInsyncReplicas,omitempty"`
	PartitionNum      *string `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ReplicationFactor *int64  `json:"ReplicationFactor,omitempty" xml:"ReplicationFactor,omitempty"`
	Topic             *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s CreateTopicShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicShrinkRequest) SetCompactTopic(v bool) *CreateTopicShrinkRequest {
	s.CompactTopic = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetConfigShrink(v string) *CreateTopicShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetInstanceId(v string) *CreateTopicShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetLocalTopic(v bool) *CreateTopicShrinkRequest {
	s.LocalTopic = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetMinInsyncReplicas(v int64) *CreateTopicShrinkRequest {
	s.MinInsyncReplicas = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetPartitionNum(v string) *CreateTopicShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetRegionId(v string) *CreateTopicShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetRemark(v string) *CreateTopicShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetReplicationFactor(v int64) *CreateTopicShrinkRequest {
	s.ReplicationFactor = &v
	return s
}

func (s *CreateTopicShrinkRequest) SetTopic(v string) *CreateTopicShrinkRequest {
	s.Topic = &v
	return s
}

type CreateTopicResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclOperationType       *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	AclResourceName        *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	AclResourceType        *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Username               *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSaslUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AclResourceName        *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	AclResourceType        *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Username               *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Code         *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	KafkaAclList *DescribeAclsResponseBodyKafkaAclList `json:"KafkaAclList,omitempty" xml:"KafkaAclList,omitempty" type:"Struct"`
	Message      *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AclOperationType       *string `json:"AclOperationType,omitempty" xml:"AclOperationType,omitempty"`
	AclResourceName        *string `json:"AclResourceName,omitempty" xml:"AclResourceName,omitempty"`
	AclResourcePatternType *string `json:"AclResourcePatternType,omitempty" xml:"AclResourcePatternType,omitempty"`
	AclResourceType        *string `json:"AclResourceType,omitempty" xml:"AclResourceType,omitempty"`
	Host                   *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Username               *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAclsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeNodeStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNodeStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeStatusRequest) SetInstanceId(v string) *DescribeNodeStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNodeStatusRequest) SetRegionId(v string) *DescribeNodeStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeNodeStatusResponseBody struct {
	Code       *int32                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatusList *DescribeNodeStatusResponseBodyStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Struct"`
	Success    *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeNodeStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeStatusResponseBody) SetCode(v int32) *DescribeNodeStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNodeStatusResponseBody) SetMessage(v string) *DescribeNodeStatusResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNodeStatusResponseBody) SetRequestId(v string) *DescribeNodeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeStatusResponseBody) SetStatusList(v *DescribeNodeStatusResponseBodyStatusList) *DescribeNodeStatusResponseBody {
	s.StatusList = v
	return s
}

func (s *DescribeNodeStatusResponseBody) SetSuccess(v bool) *DescribeNodeStatusResponseBody {
	s.Success = &v
	return s
}

type DescribeNodeStatusResponseBodyStatusList struct {
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s DescribeNodeStatusResponseBodyStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeStatusResponseBodyStatusList) GoString() string {
	return s.String()
}

func (s *DescribeNodeStatusResponseBodyStatusList) SetStatus(v []*string) *DescribeNodeStatusResponseBodyStatusList {
	s.Status = v
	return s
}

type DescribeNodeStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNodeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeStatusResponse) SetHeaders(v map[string]*string) *DescribeNodeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeStatusResponse) SetStatusCode(v int32) *DescribeNodeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeStatusResponse) SetBody(v *DescribeNodeStatusResponseBody) *DescribeNodeStatusResponse {
	s.Body = v
	return s
}

type DescribeSaslUsersRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code         *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SaslUserList *DescribeSaslUsersResponseBodySaslUserList `json:"SaslUserList,omitempty" xml:"SaslUserList,omitempty" type:"Struct"`
	Success      *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSaslUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Code        *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceIds map[string]interface{} `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Message     *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllInstanceIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	AllowedList *GetAllowedIpListResponseBodyAllowedList `json:"AllowedList,omitempty" xml:"AllowedList,omitempty" type:"Struct"`
	Code        *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
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
	DeployType   *int32                                                 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	InternetList []*GetAllowedIpListResponseBodyAllowedListInternetList `json:"InternetList,omitempty" xml:"InternetList,omitempty" type:"Repeated"`
	VpcList      []*GetAllowedIpListResponseBodyAllowedListVpcList      `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
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
	AllowedIpList []*string `json:"AllowedIpList,omitempty" xml:"AllowedIpList,omitempty" type:"Repeated"`
	PortRange     *string   `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
}

func (s GetAllowedIpListResponseBodyAllowedListInternetList) String() string {
	return tea.Prettify(s)
}

func (s GetAllowedIpListResponseBodyAllowedListInternetList) GoString() string {
	return s.String()
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
	AllowedIpList []*string `json:"AllowedIpList,omitempty" xml:"AllowedIpList,omitempty" type:"Repeated"`
	PortRange     *string   `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
}

func (s GetAllowedIpListResponseBodyAllowedListVpcList) String() string {
	return tea.Prettify(s)
}

func (s GetAllowedIpListResponseBodyAllowedListVpcList) GoString() string {
	return s.String()
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllowedIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code         *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	ConsumerList *GetConsumerListResponseBodyConsumerList `json:"ConsumerList,omitempty" xml:"ConsumerList,omitempty" type:"Struct"`
	Message      *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
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
	ConsumerId *string                                                `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	InstanceId *string                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark     *string                                                `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Tags       *GetConsumerListResponseBodyConsumerListConsumerVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s GetConsumerListResponseBodyConsumerListConsumerVO) String() string {
	return tea.Prettify(s)
}

func (s GetConsumerListResponseBodyConsumerListConsumerVO) GoString() string {
	return s.String()
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConsumerListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code             *int32                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	ConsumerProgress *GetConsumerProgressResponseBodyConsumerProgress `json:"ConsumerProgress,omitempty" xml:"ConsumerProgress,omitempty" type:"Struct"`
	Message          *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId        *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success          *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
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
	LastTimestamp *int64                                                    `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	TopicList     *GetConsumerProgressResponseBodyConsumerProgressTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
	TotalDiff     *int64                                                    `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
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
	LastTimestamp *int64                                                                       `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	OffsetList    *GetConsumerProgressResponseBodyConsumerProgressTopicListTopicListOffsetList `json:"OffsetList,omitempty" xml:"OffsetList,omitempty" type:"Struct"`
	Topic         *string                                                                      `json:"Topic,omitempty" xml:"Topic,omitempty"`
	TotalDiff     *int64                                                                       `json:"TotalDiff,omitempty" xml:"TotalDiff,omitempty"`
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
	BrokerOffset   *int64 `json:"BrokerOffset,omitempty" xml:"BrokerOffset,omitempty"`
	ConsumerOffset *int64 `json:"ConsumerOffset,omitempty" xml:"ConsumerOffset,omitempty"`
	LastTimestamp  *int64 `json:"LastTimestamp,omitempty" xml:"LastTimestamp,omitempty"`
	Partition      *int32 `json:"Partition,omitempty" xml:"Partition,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConsumerProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId      []*string                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OrderId         *string                      `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RegionId        *string                      `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                      `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*GetInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Code         *int32                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	InstanceList *GetInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Struct"`
	Message      *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AllConfig                *string                                                                    `json:"AllConfig,omitempty" xml:"AllConfig,omitempty"`
	CreateTime               *int64                                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployType               *int32                                                                     `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DiskSize                 *int32                                                                     `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType                 *int32                                                                     `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	DomainEndpoint           *string                                                                    `json:"DomainEndpoint,omitempty" xml:"DomainEndpoint,omitempty"`
	EipMax                   *int32                                                                     `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	EndPoint                 *string                                                                    `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	ExpiredTime              *int64                                                                     `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceId               *string                                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoMax                    *int32                                                                     `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	MsgRetain                *int32                                                                     `json:"MsgRetain,omitempty" xml:"MsgRetain,omitempty"`
	Name                     *string                                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	PaidType                 *int32                                                                     `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	RegionId                 *string                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId          *string                                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SaslDomainEndpoint       *string                                                                    `json:"SaslDomainEndpoint,omitempty" xml:"SaslDomainEndpoint,omitempty"`
	SecurityGroup            *string                                                                    `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	ServiceStatus            *int32                                                                     `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	SpecType                 *string                                                                    `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	SslDomainEndpoint        *string                                                                    `json:"SslDomainEndpoint,omitempty" xml:"SslDomainEndpoint,omitempty"`
	SslEndPoint              *string                                                                    `json:"SslEndPoint,omitempty" xml:"SslEndPoint,omitempty"`
	Tags                     *GetInstanceListResponseBodyInstanceListInstanceVOTags                     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TopicNumLimit            *int32                                                                     `json:"TopicNumLimit,omitempty" xml:"TopicNumLimit,omitempty"`
	UpgradeServiceDetailInfo *GetInstanceListResponseBodyInstanceListInstanceVOUpgradeServiceDetailInfo `json:"UpgradeServiceDetailInfo,omitempty" xml:"UpgradeServiceDetailInfo,omitempty" type:"Struct"`
	UsedGroupCount           *int32                                                                     `json:"UsedGroupCount,omitempty" xml:"UsedGroupCount,omitempty"`
	UsedPartitionCount       *int32                                                                     `json:"UsedPartitionCount,omitempty" xml:"UsedPartitionCount,omitempty"`
	UsedTopicCount           *int32                                                                     `json:"UsedTopicCount,omitempty" xml:"UsedTopicCount,omitempty"`
	VSwitchId                *string                                                                    `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                    *string                                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId                   *string                                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetTopicListRequest struct {
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic       *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Code        *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Message     *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize    *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	TopicList   *GetTopicListResponseBodyTopicList `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Struct"`
	Total       *int32                             `json:"Total,omitempty" xml:"Total,omitempty"`
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
	CompactTopic *bool                                         `json:"CompactTopic,omitempty" xml:"CompactTopic,omitempty"`
	CreateTime   *int64                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId   *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LocalTopic   *bool                                         `json:"LocalTopic,omitempty" xml:"LocalTopic,omitempty"`
	PartitionNum *int32                                        `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	RegionId     *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark       *string                                       `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status       *int32                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusName   *string                                       `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	Tags         *GetTopicListResponseBodyTopicListTopicVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Topic        *string                                       `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s GetTopicListResponseBodyTopicListTopicVO) String() string {
	return tea.Prettify(s)
}

func (s GetTopicListResponseBodyTopicListTopicVO) GoString() string {
	return s.String()
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTopicListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Code        *int32                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
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
	LastTimeStamp *int64                                            `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	OffsetTable   *GetTopicStatusResponseBodyTopicStatusOffsetTable `json:"OffsetTable,omitempty" xml:"OffsetTable,omitempty" type:"Struct"`
	TotalCount    *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	LastUpdateTimestamp *int64  `json:"LastUpdateTimestamp,omitempty" xml:"LastUpdateTimestamp,omitempty"`
	MaxOffset           *int64  `json:"MaxOffset,omitempty" xml:"MaxOffset,omitempty"`
	MinOffset           *int64  `json:"MinOffset,omitempty" xml:"MinOffset,omitempty"`
	Partition           *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	Topic               *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTopicStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListTagResourcesRequest struct {
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AddPartitionNum *int32  `json:"AddPartitionNum,omitempty" xml:"AddPartitionNum,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Topic           *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyPartitionNumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark     *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTopicRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ReleaseInstanceRequest struct {
	ForceDeleteInstance *bool   `json:"ForceDeleteInstance,omitempty" xml:"ForceDeleteInstance,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Config               *string `json:"Config,omitempty" xml:"Config,omitempty"`
	DeployModule         *string `json:"DeployModule,omitempty" xml:"DeployModule,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsEipInner           *bool   `json:"IsEipInner,omitempty" xml:"IsEipInner,omitempty"`
	IsForceSelectedZones *bool   `json:"IsForceSelectedZones,omitempty" xml:"IsForceSelectedZones,omitempty"`
	IsSetUserAndPassword *bool   `json:"IsSetUserAndPassword,omitempty" xml:"IsSetUserAndPassword,omitempty"`
	KMSKeyId             *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Notifier             *string `json:"Notifier,omitempty" xml:"Notifier,omitempty"`
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroup        *string `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty"`
	SelectedZones        *string `json:"SelectedZones,omitempty" xml:"SelectedZones,omitempty"`
	ServiceVersion       *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	UserPhoneNum         *string `json:"UserPhoneNum,omitempty" xml:"UserPhoneNum,omitempty"`
	Username             *string `json:"Username,omitempty" xml:"Username,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId   *string                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AllowedListIp   *string `json:"AllowedListIp,omitempty" xml:"AllowedListIp,omitempty"`
	AllowedListType *string `json:"AllowedListType,omitempty" xml:"AllowedListType,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PortRange       *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	UpdateType      *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAllowedIpResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateInstanceConfigRequest struct {
	Config     *string `json:"Config,omitempty" xml:"Config,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpgradeInstanceVersionRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeInstanceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DiskSize   *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	EipMax     *int32  `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	EipModel   *bool   `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoMax      *int32  `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxSpec  *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpecType   *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TopicQuota *int32  `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
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

func (s *UpgradePostPayOrderRequest) SetRegionId(v string) *UpgradePostPayOrderRequest {
	s.RegionId = &v
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

type UpgradePostPayOrderResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradePostPayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DiskSize   *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	EipMax     *int32  `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	EipModel   *bool   `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoMax      *int32  `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxSpec  *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpecType   *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TopicQuota *int32  `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
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
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradePrePayOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("alikafka.aliyuncs.com"),
		"ap-southeast-2":              tea.String("alikafka.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("alikafka.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("alikafka.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("alikafka.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("alikafka.aliyuncs.com"),
		"cn-edge-1":                   tea.String("alikafka.aliyuncs.com"),
		"cn-fujian":                   tea.String("alikafka.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("alikafka.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("alikafka.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("alikafka.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("alikafka.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("alikafka.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("alikafka.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("alikafka.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("alikafka.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("alikafka.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("alikafka.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("alikafka.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("alikafka.aliyuncs.com"),
		"cn-wuhan":                    tea.String("alikafka.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("alikafka.aliyuncs.com"),
		"cn-yushanfang":               tea.String("alikafka.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("alikafka.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("alikafka.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("alikafka.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("alikafka.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("alikafka.aliyuncs.com"),
		"me-east-1":                   tea.String("alikafka.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("alikafka.aliyuncs.com"),
	}
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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

func (client *Client) CreateTopicWithOptions(tmpReq *CreateTopicRequest, runtime *util.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTopicShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Config)) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, tea.String("Config"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompactTopic)) {
		query["CompactTopic"] = request.CompactTopic
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigShrink)) {
		query["Config"] = request.ConfigShrink
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

func (client *Client) DescribeNodeStatusWithOptions(request *DescribeNodeStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeStatusResponse, _err error) {
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
		Action:      tea.String("DescribeNodeStatus"),
		Version:     tea.String("2019-09-16"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeStatus(request *DescribeNodeStatusRequest) (_result *DescribeNodeStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeStatusResponse{}
	_body, _err := client.DescribeNodeStatusWithOptions(request, runtime)
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

func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		query["Config"] = request.Config
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

func (client *Client) UpgradePostPayOrderWithOptions(request *UpgradePostPayOrderRequest, runtime *util.RuntimeOptions) (_result *UpgradePostPayOrderResponse, _err error) {
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
