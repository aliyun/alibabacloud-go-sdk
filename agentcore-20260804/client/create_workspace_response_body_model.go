// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkspaceResponseBody
	GetCode() *string
	SetData(v *CreateWorkspaceResponseBodyData) *CreateWorkspaceResponseBody
	GetData() *CreateWorkspaceResponseBodyData
	SetHttpStatusCode(v int32) *CreateWorkspaceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkspaceResponseBody
	GetSuccess() *bool
}

type CreateWorkspaceResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The workspace details.
	Data *CreateWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkspaceResponseBody) GetData() *CreateWorkspaceResponseBodyData {
	return s.Data
}

func (s *CreateWorkspaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkspaceResponseBody) SetCode(v string) *CreateWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetData(v *CreateWorkspaceResponseBodyData) *CreateWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkspaceResponseBody) SetHttpStatusCode(v int32) *CreateWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetMessage(v string) *CreateWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetSuccess(v bool) *CreateWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyData struct {
	// The AI Registry namespace ID. It is returned after related resources are bound and may be empty during initialization.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// namespace-0123456789abcdef
	AiRegistryNamespaceId *string `json:"aiRegistryNamespaceId,omitempty" xml:"aiRegistryNamespaceId,omitempty"`
	// The OSS storage authorization status.
	//
	// example:
	//
	// AUTHORIZED
	AuthorizationStatus *string `json:"authorizationStatus,omitempty" xml:"authorizationStatus,omitempty"`
	// The name of the private OSS bucket.
	//
	// example:
	//
	// bucket-001
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The CloudMonitor workspace ID. It is returned after related resources are bound and may be empty during initialization.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// cms-ws-0123456789abcdef
	CmsWorkspaceId *string `json:"cmsWorkspaceId,omitempty" xml:"cmsWorkspaceId,omitempty"`
	// The time when the workspace was created, in ISO 8601 format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-08-28T10:00:00+08:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// production-agents
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration of the workspace.
	NetworkConfiguration *CreateWorkspaceResponseBodyDataNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The ID of the region where the workspace resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the resource group to which the workspace belongs. It may be empty if no resource group is specified.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// rg-acfm1234567890
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The status of the workspace.
	//
	// example:
	//
	// Initialized
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The supplementary reason for the current status of the workspace. It is used to display the specific reason when initialization fails or when waiting for authorization. It may be empty under normal circumstances.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Waiting for OSS RAM authorization
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The storage type of the workspace.
	//
	// example:
	//
	// PRIVATE
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// The list of tags for the workspace. An empty array is returned if no tags are set.
	//
	// This parameter is required.
	Tags []*CreateWorkspaceResponseBodyDataTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// example:
	//
	// tenant-123456
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyData) GetAiRegistryNamespaceId() *string {
	return s.AiRegistryNamespaceId
}

func (s *CreateWorkspaceResponseBodyData) GetAuthorizationStatus() *string {
	return s.AuthorizationStatus
}

func (s *CreateWorkspaceResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *CreateWorkspaceResponseBodyData) GetCmsWorkspaceId() *string {
	return s.CmsWorkspaceId
}

func (s *CreateWorkspaceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateWorkspaceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceResponseBodyData) GetNetworkConfiguration() *CreateWorkspaceResponseBodyDataNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateWorkspaceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWorkspaceResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateWorkspaceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateWorkspaceResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *CreateWorkspaceResponseBodyData) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateWorkspaceResponseBodyData) GetTags() []*CreateWorkspaceResponseBodyDataTags {
	return s.Tags
}

func (s *CreateWorkspaceResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceResponseBodyData) SetAiRegistryNamespaceId(v string) *CreateWorkspaceResponseBodyData {
	s.AiRegistryNamespaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetAuthorizationStatus(v string) *CreateWorkspaceResponseBodyData {
	s.AuthorizationStatus = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetBucketName(v string) *CreateWorkspaceResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetCmsWorkspaceId(v string) *CreateWorkspaceResponseBodyData {
	s.CmsWorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetCreateTime(v string) *CreateWorkspaceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetName(v string) *CreateWorkspaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetNetworkConfiguration(v *CreateWorkspaceResponseBodyDataNetworkConfiguration) *CreateWorkspaceResponseBodyData {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetRegionId(v string) *CreateWorkspaceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetResourceGroupId(v string) *CreateWorkspaceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetStatus(v string) *CreateWorkspaceResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetStatusReason(v string) *CreateWorkspaceResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetStorageType(v string) *CreateWorkspaceResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetTags(v []*CreateWorkspaceResponseBodyDataTags) *CreateWorkspaceResponseBodyData {
	s.Tags = v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetTenantId(v string) *CreateWorkspaceResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetWorkspaceId(v string) *CreateWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyDataNetworkConfiguration struct {
	// Returns the target configuration value. Updating indicates that the change is still in progress, and UpdateFailed indicates that the current change failed. This parameter only controls whether the latest versions of managed and custom agents use the public network egress of the platform. It does not control fixed historical versions, model or MCP proxy egress, or the public network egress of the user VPC.
	//
	// example:
	//
	// true
	PublicEgressEnabled *bool `json:"publicEgressEnabled,omitempty" xml:"publicEgressEnabled,omitempty"`
	// Returns the target configuration value. Updating indicates that the change is in progress, and UpdateFailed indicates that the change failed. When disabled, the public network ingress for all agent endpoints in the workspace is closed. The original public network ingress is restored when re-enabled. INNER_DIRECT and user VPC ingress are not affected. Collaboration and Element Web ingress are not controlled by this parameter.
	//
	// example:
	//
	// true
	PublicIngressEnabled *bool `json:"publicIngressEnabled,omitempty" xml:"publicIngressEnabled,omitempty"`
	// The user VPC network configuration.
	Vpc *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s CreateWorkspaceResponseBodyDataNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyDataNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) GetPublicEgressEnabled() *bool {
	return s.PublicEgressEnabled
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) GetPublicIngressEnabled() *bool {
	return s.PublicIngressEnabled
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) GetVpc() *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	return s.Vpc
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) SetPublicEgressEnabled(v bool) *CreateWorkspaceResponseBodyDataNetworkConfiguration {
	s.PublicEgressEnabled = &v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) SetPublicIngressEnabled(v bool) *CreateWorkspaceResponseBodyDataNetworkConfiguration {
	s.PublicIngressEnabled = &v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) SetVpc(v *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) *CreateWorkspaceResponseBodyDataNetworkConfiguration {
	s.Vpc = v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) Validate() error {
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyDataNetworkConfigurationVpc struct {
	// Specifies whether to enable the VPC network.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The user VPC ID.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetEnabled(v bool) *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.Enabled = &v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetVSwitchIds(v []*string) *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.VSwitchIds = v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetVpcId(v string) *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.VpcId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) Validate() error {
	return dara.Validate(s)
}

type CreateWorkspaceResponseBodyDataTags struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// environment
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// This parameter is required.
	//
	// example:
	//
	// development
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateWorkspaceResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *CreateWorkspaceResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *CreateWorkspaceResponseBodyDataTags) SetKey(v string) *CreateWorkspaceResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *CreateWorkspaceResponseBodyDataTags) SetValue(v string) *CreateWorkspaceResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *CreateWorkspaceResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
