// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AttachInstancesRequest struct {
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s AttachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesRequest) SetInstanceId(v []*string) *AttachInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *AttachInstancesRequest) SetOwnerAccount(v string) *AttachInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachInstancesRequest) SetOwnerId(v int64) *AttachInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachInstancesRequest) SetResourceOwnerAccount(v string) *AttachInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachInstancesRequest) SetResourceOwnerId(v int64) *AttachInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachInstancesRequest) SetScalingGroupId(v string) *AttachInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type AttachInstancesResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s AttachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponseBody) SetRequestId(v string) *AttachInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachInstancesResponseBody) SetScalingActivityId(v string) *AttachInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

type AttachInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachInstancesResponse) GoString() string {
	return s.String()
}

func (s *AttachInstancesResponse) SetHeaders(v map[string]*string) *AttachInstancesResponse {
	s.Headers = v
	return s
}

func (s *AttachInstancesResponse) SetStatusCode(v int32) *AttachInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachInstancesResponse) SetBody(v *AttachInstancesResponseBody) *AttachInstancesResponse {
	s.Body = v
	return s
}

type CreateScalingConfigurationRequest struct {
	DataDisk   []*CreateScalingConfigurationRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SystemDisk *CreateScalingConfigurationRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// This parameter is required.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	InstanceType             *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType       *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn   *int32  `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut  *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized              *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount     *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingConfigurationName *string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	// This parameter is required.
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// This parameter is required.
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s CreateScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequest) SetDataDisk(v []*CreateScalingConfigurationRequestDataDisk) *CreateScalingConfigurationRequest {
	s.DataDisk = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSystemDisk(v *CreateScalingConfigurationRequestSystemDisk) *CreateScalingConfigurationRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateScalingConfigurationRequest) SetImageId(v string) *CreateScalingConfigurationRequest {
	s.ImageId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInstanceType(v string) *CreateScalingConfigurationRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetChargeType(v string) *CreateScalingConfigurationRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetMaxBandwidthIn(v int32) *CreateScalingConfigurationRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetInternetMaxBandwidthOut(v int32) *CreateScalingConfigurationRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetIoOptimized(v string) *CreateScalingConfigurationRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetOwnerAccount(v string) *CreateScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetOwnerId(v int64) *CreateScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourceOwnerAccount(v string) *CreateScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetResourceOwnerId(v int64) *CreateScalingConfigurationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetScalingConfigurationName(v string) *CreateScalingConfigurationRequest {
	s.ScalingConfigurationName = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetScalingGroupId(v string) *CreateScalingConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingConfigurationRequest) SetSecurityGroupId(v string) *CreateScalingConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

type CreateScalingConfigurationRequestDataDisk struct {
	Category           *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithInstance *string `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Device             *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Size               *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId         *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateScalingConfigurationRequestDataDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestDataDisk) SetCategory(v string) *CreateScalingConfigurationRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetDeleteWithInstance(v string) *CreateScalingConfigurationRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetDevice(v string) *CreateScalingConfigurationRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetSize(v int32) *CreateScalingConfigurationRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateScalingConfigurationRequestDataDisk) SetSnapshotId(v string) *CreateScalingConfigurationRequestDataDisk {
	s.SnapshotId = &v
	return s
}

type CreateScalingConfigurationRequestSystemDisk struct {
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Size     *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateScalingConfigurationRequestSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetCategory(v string) *CreateScalingConfigurationRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateScalingConfigurationRequestSystemDisk) SetSize(v int32) *CreateScalingConfigurationRequestSystemDisk {
	s.Size = &v
	return s
}

type CreateScalingConfigurationResponseBody struct {
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s CreateScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationResponseBody) SetRequestId(v string) *CreateScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingConfigurationResponseBody) SetScalingConfigurationId(v string) *CreateScalingConfigurationResponseBody {
	s.ScalingConfigurationId = &v
	return s
}

type CreateScalingConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationResponse) SetHeaders(v map[string]*string) *CreateScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingConfigurationResponse) SetStatusCode(v int32) *CreateScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingConfigurationResponse) SetBody(v *CreateScalingConfigurationResponseBody) *CreateScalingConfigurationResponse {
	s.Body = v
	return s
}

type CreateScalingGroupRequest struct {
	RemovalPolicy   []*string `json:"RemovalPolicy,omitempty" xml:"RemovalPolicy,omitempty" type:"Repeated"`
	DBInstanceIds   *string   `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	DefaultCooldown *int32    `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	LoadBalancerIds *string   `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty"`
	// This parameter is required.
	MaxSize *int32 `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// This parameter is required.
	MinSize      *int32  `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ScalingGroupName     *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupRequest) SetRemovalPolicy(v []*string) *CreateScalingGroupRequest {
	s.RemovalPolicy = v
	return s
}

func (s *CreateScalingGroupRequest) SetDBInstanceIds(v string) *CreateScalingGroupRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *CreateScalingGroupRequest) SetDefaultCooldown(v int32) *CreateScalingGroupRequest {
	s.DefaultCooldown = &v
	return s
}

func (s *CreateScalingGroupRequest) SetLoadBalancerIds(v string) *CreateScalingGroupRequest {
	s.LoadBalancerIds = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMaxSize(v int32) *CreateScalingGroupRequest {
	s.MaxSize = &v
	return s
}

func (s *CreateScalingGroupRequest) SetMinSize(v int32) *CreateScalingGroupRequest {
	s.MinSize = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOwnerAccount(v string) *CreateScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingGroupRequest) SetOwnerId(v int64) *CreateScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetRegionId(v string) *CreateScalingGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScalingGroupRequest) SetResourceOwnerAccount(v string) *CreateScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingGroupRequest) SetScalingGroupName(v string) *CreateScalingGroupRequest {
	s.ScalingGroupName = &v
	return s
}

func (s *CreateScalingGroupRequest) SetVSwitchId(v string) *CreateScalingGroupRequest {
	s.VSwitchId = &v
	return s
}

type CreateScalingGroupResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CreateScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupResponseBody) SetRequestId(v string) *CreateScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingGroupResponseBody) SetScalingGroupId(v string) *CreateScalingGroupResponseBody {
	s.ScalingGroupId = &v
	return s
}

type CreateScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingGroupResponse) SetHeaders(v map[string]*string) *CreateScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingGroupResponse) SetStatusCode(v int32) *CreateScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingGroupResponse) SetBody(v *CreateScalingGroupResponseBody) *CreateScalingGroupResponse {
	s.Body = v
	return s
}

type CreateScalingRuleRequest struct {
	// This parameter is required.
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// This parameter is required.
	AdjustmentValue      *int32  `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	Cooldown             *int32  `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingGroupId  *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s CreateScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleRequest) SetAdjustmentType(v string) *CreateScalingRuleRequest {
	s.AdjustmentType = &v
	return s
}

func (s *CreateScalingRuleRequest) SetAdjustmentValue(v int32) *CreateScalingRuleRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *CreateScalingRuleRequest) SetCooldown(v int32) *CreateScalingRuleRequest {
	s.Cooldown = &v
	return s
}

func (s *CreateScalingRuleRequest) SetOwnerAccount(v string) *CreateScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetOwnerId(v int64) *CreateScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetResourceOwnerAccount(v string) *CreateScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScalingRuleRequest) SetResourceOwnerId(v int64) *CreateScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingGroupId(v string) *CreateScalingRuleRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateScalingRuleRequest) SetScalingRuleName(v string) *CreateScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

type CreateScalingRuleResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingRuleAri *string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
	ScalingRuleId  *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
}

func (s CreateScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleResponseBody) SetRequestId(v string) *CreateScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetScalingRuleAri(v string) *CreateScalingRuleResponseBody {
	s.ScalingRuleAri = &v
	return s
}

func (s *CreateScalingRuleResponseBody) SetScalingRuleId(v string) *CreateScalingRuleResponseBody {
	s.ScalingRuleId = &v
	return s
}

type CreateScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateScalingRuleResponse) SetHeaders(v map[string]*string) *CreateScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateScalingRuleResponse) SetStatusCode(v int32) *CreateScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScalingRuleResponse) SetBody(v *CreateScalingRuleResponseBody) *CreateScalingRuleResponse {
	s.Body = v
	return s
}

type CreateScheduledTaskRequest struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScheduledAction   *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	ScheduledTaskName *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	TaskEnabled       *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
}

func (s CreateScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) SetDescription(v string) *CreateScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetLaunchExpirationTime(v int32) *CreateScheduledTaskRequest {
	s.LaunchExpirationTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetLaunchTime(v string) *CreateScheduledTaskRequest {
	s.LaunchTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetOwnerAccount(v string) *CreateScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetOwnerId(v int64) *CreateScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceEndTime(v string) *CreateScheduledTaskRequest {
	s.RecurrenceEndTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceType(v string) *CreateScheduledTaskRequest {
	s.RecurrenceType = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRecurrenceValue(v string) *CreateScheduledTaskRequest {
	s.RecurrenceValue = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetRegionId(v string) *CreateScheduledTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetResourceOwnerAccount(v string) *CreateScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetResourceOwnerId(v int64) *CreateScheduledTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScheduledAction(v string) *CreateScheduledTaskRequest {
	s.ScheduledAction = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetScheduledTaskName(v string) *CreateScheduledTaskRequest {
	s.ScheduledTaskName = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetTaskEnabled(v bool) *CreateScheduledTaskRequest {
	s.TaskEnabled = &v
	return s
}

type CreateScheduledTaskResponseBody struct {
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduledTaskId *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetScheduledTaskId(v string) *CreateScheduledTaskResponseBody {
	s.ScheduledTaskId = &v
	return s
}

type CreateScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponse) SetHeaders(v map[string]*string) *CreateScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledTaskResponse) SetStatusCode(v int32) *CreateScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledTaskResponse) SetBody(v *CreateScheduledTaskResponseBody) *CreateScheduledTaskResponse {
	s.Body = v
	return s
}

type DeleteScalingConfigurationRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s DeleteScalingConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationRequest) SetOwnerAccount(v string) *DeleteScalingConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetOwnerId(v int64) *DeleteScalingConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetResourceOwnerAccount(v string) *DeleteScalingConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetResourceOwnerId(v int64) *DeleteScalingConfigurationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteScalingConfigurationRequest) SetScalingConfigurationId(v string) *DeleteScalingConfigurationRequest {
	s.ScalingConfigurationId = &v
	return s
}

type DeleteScalingConfigurationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScalingConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationResponseBody) SetRequestId(v string) *DeleteScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScalingConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScalingConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScalingConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationResponse) SetHeaders(v map[string]*string) *DeleteScalingConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingConfigurationResponse) SetStatusCode(v int32) *DeleteScalingConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingConfigurationResponse) SetBody(v *DeleteScalingConfigurationResponseBody) *DeleteScalingConfigurationResponse {
	s.Body = v
	return s
}

type DeleteScalingGroupRequest struct {
	ForceDelete          *bool   `json:"ForceDelete,omitempty" xml:"ForceDelete,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// This parameter is required.
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DeleteScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupRequest) SetForceDelete(v bool) *DeleteScalingGroupRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetOwnerAccount(v string) *DeleteScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetOwnerId(v int64) *DeleteScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetResourceOwnerAccount(v string) *DeleteScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingGroupRequest) SetScalingGroupId(v string) *DeleteScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

type DeleteScalingGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupResponseBody) SetRequestId(v string) *DeleteScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingGroupResponse) SetHeaders(v map[string]*string) *DeleteScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingGroupResponse) SetStatusCode(v int32) *DeleteScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingGroupResponse) SetBody(v *DeleteScalingGroupResponseBody) *DeleteScalingGroupResponse {
	s.Body = v
	return s
}

type DeleteScalingRuleRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
}

func (s DeleteScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleRequest) SetOwnerAccount(v string) *DeleteScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetOwnerId(v int64) *DeleteScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetResourceOwnerAccount(v string) *DeleteScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetResourceOwnerId(v int64) *DeleteScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteScalingRuleRequest) SetScalingRuleId(v string) *DeleteScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

type DeleteScalingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleResponseBody) SetRequestId(v string) *DeleteScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteScalingRuleResponse) SetStatusCode(v int32) *DeleteScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScalingRuleResponse) SetBody(v *DeleteScalingRuleResponseBody) *DeleteScalingRuleResponse {
	s.Body = v
	return s
}

type DeleteScheduledTaskRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScheduledTaskId *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
}

func (s DeleteScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskRequest) SetOwnerAccount(v string) *DeleteScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteScheduledTaskRequest) SetOwnerId(v int64) *DeleteScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteScheduledTaskRequest) SetResourceOwnerAccount(v string) *DeleteScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteScheduledTaskRequest) SetResourceOwnerId(v int64) *DeleteScheduledTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteScheduledTaskRequest) SetScheduledTaskId(v string) *DeleteScheduledTaskRequest {
	s.ScheduledTaskId = &v
	return s
}

type DeleteScheduledTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskResponseBody) SetRequestId(v string) *DeleteScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledTaskResponse) SetHeaders(v map[string]*string) *DeleteScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledTaskResponse) SetStatusCode(v int32) *DeleteScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledTaskResponse) SetBody(v *DeleteScheduledTaskResponseBody) *DeleteScheduledTaskResponse {
	s.Body = v
	return s
}

type DescribeAccountAttributesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAccountAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesRequest) SetOwnerId(v int64) *DescribeAccountAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountAttributesRequest) SetResourceOwnerAccount(v string) *DescribeAccountAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountAttributesRequest) SetResourceOwnerId(v int64) *DescribeAccountAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeAccountAttributesResponseBody struct {
	MaxNumberOfDBInstances           *int32 `json:"MaxNumberOfDBInstances,omitempty" xml:"MaxNumberOfDBInstances,omitempty"`
	MaxNumberOfLoadBalancers         *int32 `json:"MaxNumberOfLoadBalancers,omitempty" xml:"MaxNumberOfLoadBalancers,omitempty"`
	MaxNumberOfMaxSize               *int32 `json:"MaxNumberOfMaxSize,omitempty" xml:"MaxNumberOfMaxSize,omitempty"`
	MaxNumberOfMinSize               *int32 `json:"MaxNumberOfMinSize,omitempty" xml:"MaxNumberOfMinSize,omitempty"`
	MaxNumberOfScalingConfigurations *int32 `json:"MaxNumberOfScalingConfigurations,omitempty" xml:"MaxNumberOfScalingConfigurations,omitempty"`
	MaxNumberOfScalingGroups         *int32 `json:"MaxNumberOfScalingGroups,omitempty" xml:"MaxNumberOfScalingGroups,omitempty"`
	MaxNumberOfScalingInstances      *int32 `json:"MaxNumberOfScalingInstances,omitempty" xml:"MaxNumberOfScalingInstances,omitempty"`
	MaxNumberOfScalingRules          *int32 `json:"MaxNumberOfScalingRules,omitempty" xml:"MaxNumberOfScalingRules,omitempty"`
	MaxNumberOfScheduledTasks        *int32 `json:"MaxNumberOfScheduledTasks,omitempty" xml:"MaxNumberOfScheduledTasks,omitempty"`
}

func (s DescribeAccountAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfDBInstances(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfDBInstances = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfLoadBalancers(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfLoadBalancers = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfMaxSize(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfMaxSize = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfMinSize(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfMinSize = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfScalingConfigurations(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfScalingConfigurations = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfScalingGroups(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfScalingGroups = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfScalingInstances(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfScalingInstances = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfScalingRules(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfScalingRules = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetMaxNumberOfScheduledTasks(v int32) *DescribeAccountAttributesResponseBody {
	s.MaxNumberOfScheduledTasks = &v
	return s
}

type DescribeAccountAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesResponse) SetHeaders(v map[string]*string) *DescribeAccountAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountAttributesResponse) SetStatusCode(v int32) *DescribeAccountAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountAttributesResponse) SetBody(v *DescribeAccountAttributesResponseBody) *DescribeAccountAttributesResponse {
	s.Body = v
	return s
}

type DescribeCapacityHistoryRequest struct {
	// This parameter is required.
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCapacityHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCapacityHistoryRequest) SetEndTime(v string) *DescribeCapacityHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCapacityHistoryRequest) SetOwnerId(v int64) *DescribeCapacityHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCapacityHistoryRequest) SetPageNumber(v int32) *DescribeCapacityHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCapacityHistoryRequest) SetPageSize(v int32) *DescribeCapacityHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCapacityHistoryRequest) SetResourceOwnerAccount(v string) *DescribeCapacityHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCapacityHistoryRequest) SetResourceOwnerId(v int64) *DescribeCapacityHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCapacityHistoryRequest) SetScalingGroupId(v string) *DescribeCapacityHistoryRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeCapacityHistoryRequest) SetStartTime(v string) *DescribeCapacityHistoryRequest {
	s.StartTime = &v
	return s
}

type DescribeCapacityHistoryResponseBody struct {
	CapacityHistoryItems *DescribeCapacityHistoryResponseBodyCapacityHistoryItems `json:"CapacityHistoryItems,omitempty" xml:"CapacityHistoryItems,omitempty" type:"Struct"`
	PageNumber           *int32                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount           *int32                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCapacityHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapacityHistoryResponseBody) SetCapacityHistoryItems(v *DescribeCapacityHistoryResponseBodyCapacityHistoryItems) *DescribeCapacityHistoryResponseBody {
	s.CapacityHistoryItems = v
	return s
}

func (s *DescribeCapacityHistoryResponseBody) SetPageNumber(v int32) *DescribeCapacityHistoryResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCapacityHistoryResponseBody) SetPageSize(v int32) *DescribeCapacityHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCapacityHistoryResponseBody) SetTotalCount(v int32) *DescribeCapacityHistoryResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCapacityHistoryResponseBodyCapacityHistoryItems struct {
	CapacityHistoryModel []*DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel `json:"CapacityHistoryModel,omitempty" xml:"CapacityHistoryModel,omitempty" type:"Repeated"`
}

func (s DescribeCapacityHistoryResponseBodyCapacityHistoryItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityHistoryResponseBodyCapacityHistoryItems) GoString() string {
	return s.String()
}

func (s *DescribeCapacityHistoryResponseBodyCapacityHistoryItems) SetCapacityHistoryModel(v []*DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel) *DescribeCapacityHistoryResponseBodyCapacityHistoryItems {
	s.CapacityHistoryModel = v
	return s
}

type DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel struct {
	AttachedCapacity    *int32  `json:"AttachedCapacity,omitempty" xml:"AttachedCapacity,omitempty"`
	AutoCreatedCapacity *int32  `json:"AutoCreatedCapacity,omitempty" xml:"AutoCreatedCapacity,omitempty"`
	ScalingGroupId      *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	Timestamp           *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TotalCapacity       *int32  `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
}

func (s DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel) GoString() string {
	return s.String()
}

func (s *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel) SetAttachedCapacity(v int32) *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel {
	s.AttachedCapacity = &v
	return s
}

func (s *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel) SetAutoCreatedCapacity(v int32) *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel {
	s.AutoCreatedCapacity = &v
	return s
}

func (s *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel) SetScalingGroupId(v string) *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel) SetTimestamp(v string) *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel {
	s.Timestamp = &v
	return s
}

func (s *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel) SetTotalCapacity(v int32) *DescribeCapacityHistoryResponseBodyCapacityHistoryItemsCapacityHistoryModel {
	s.TotalCapacity = &v
	return s
}

type DescribeCapacityHistoryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCapacityHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCapacityHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCapacityHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapacityHistoryResponse) SetHeaders(v map[string]*string) *DescribeCapacityHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapacityHistoryResponse) SetStatusCode(v int32) *DescribeCapacityHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapacityHistoryResponse) SetBody(v *DescribeCapacityHistoryResponseBody) *DescribeCapacityHistoryResponse {
	s.Body = v
	return s
}

type DescribeLimitationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLimitationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLimitationRequest) GoString() string {
	return s.String()
}

func (s *DescribeLimitationRequest) SetOwnerId(v int64) *DescribeLimitationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLimitationRequest) SetResourceOwnerAccount(v string) *DescribeLimitationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLimitationRequest) SetResourceOwnerId(v int64) *DescribeLimitationRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeLimitationResponseBody struct {
	MaxNumberOfDBInstances           *int32 `json:"MaxNumberOfDBInstances,omitempty" xml:"MaxNumberOfDBInstances,omitempty"`
	MaxNumberOfLoadBalancers         *int32 `json:"MaxNumberOfLoadBalancers,omitempty" xml:"MaxNumberOfLoadBalancers,omitempty"`
	MaxNumberOfMaxSize               *int32 `json:"MaxNumberOfMaxSize,omitempty" xml:"MaxNumberOfMaxSize,omitempty"`
	MaxNumberOfMinSize               *int32 `json:"MaxNumberOfMinSize,omitempty" xml:"MaxNumberOfMinSize,omitempty"`
	MaxNumberOfScalingConfigurations *int32 `json:"MaxNumberOfScalingConfigurations,omitempty" xml:"MaxNumberOfScalingConfigurations,omitempty"`
	MaxNumberOfScalingGroups         *int32 `json:"MaxNumberOfScalingGroups,omitempty" xml:"MaxNumberOfScalingGroups,omitempty"`
	MaxNumberOfScalingInstances      *int32 `json:"MaxNumberOfScalingInstances,omitempty" xml:"MaxNumberOfScalingInstances,omitempty"`
	MaxNumberOfScalingRules          *int32 `json:"MaxNumberOfScalingRules,omitempty" xml:"MaxNumberOfScalingRules,omitempty"`
	MaxNumberOfScheduledTasks        *int32 `json:"MaxNumberOfScheduledTasks,omitempty" xml:"MaxNumberOfScheduledTasks,omitempty"`
}

func (s DescribeLimitationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLimitationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfDBInstances(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfDBInstances = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfLoadBalancers(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfLoadBalancers = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfMaxSize(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfMaxSize = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfMinSize(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfMinSize = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingConfigurations(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingConfigurations = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingGroups(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingGroups = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingInstances(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingInstances = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScalingRules(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScalingRules = &v
	return s
}

func (s *DescribeLimitationResponseBody) SetMaxNumberOfScheduledTasks(v int32) *DescribeLimitationResponseBody {
	s.MaxNumberOfScheduledTasks = &v
	return s
}

type DescribeLimitationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLimitationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLimitationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLimitationResponse) GoString() string {
	return s.String()
}

func (s *DescribeLimitationResponse) SetHeaders(v map[string]*string) *DescribeLimitationResponse {
	s.Headers = v
	return s
}

func (s *DescribeLimitationResponse) SetStatusCode(v int32) *DescribeLimitationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLimitationResponse) SetBody(v *DescribeLimitationResponseBody) *DescribeLimitationResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeScalingActivitiesRequest struct {
	ScalingActivityId []*string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty" type:"Repeated"`
	OwnerAccount      *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber        *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	StatusCode           *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s DescribeScalingActivitiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesRequest) SetScalingActivityId(v []*string) *DescribeScalingActivitiesRequest {
	s.ScalingActivityId = v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetOwnerAccount(v string) *DescribeScalingActivitiesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetOwnerId(v int64) *DescribeScalingActivitiesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetPageNumber(v int32) *DescribeScalingActivitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetPageSize(v int32) *DescribeScalingActivitiesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetRegionId(v string) *DescribeScalingActivitiesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetResourceOwnerAccount(v string) *DescribeScalingActivitiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetResourceOwnerId(v int64) *DescribeScalingActivitiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetScalingGroupId(v string) *DescribeScalingActivitiesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivitiesRequest) SetStatusCode(v string) *DescribeScalingActivitiesRequest {
	s.StatusCode = &v
	return s
}

type DescribeScalingActivitiesResponseBody struct {
	PageNumber        *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivities *DescribeScalingActivitiesResponseBodyScalingActivities `json:"ScalingActivities,omitempty" xml:"ScalingActivities,omitempty" type:"Struct"`
	TotalCount        *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingActivitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBody) SetPageNumber(v int32) *DescribeScalingActivitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetPageSize(v int32) *DescribeScalingActivitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetRequestId(v string) *DescribeScalingActivitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetScalingActivities(v *DescribeScalingActivitiesResponseBodyScalingActivities) *DescribeScalingActivitiesResponseBody {
	s.ScalingActivities = v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetTotalCount(v int32) *DescribeScalingActivitiesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingActivitiesResponseBodyScalingActivities struct {
	ScalingActivity []*DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity `json:"ScalingActivity,omitempty" xml:"ScalingActivity,omitempty" type:"Repeated"`
}

func (s DescribeScalingActivitiesResponseBodyScalingActivities) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBodyScalingActivities) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetScalingActivity(v []*DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ScalingActivity = v
	return s
}

type DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity struct {
	AttachedCapacity    *string `json:"AttachedCapacity,omitempty" xml:"AttachedCapacity,omitempty"`
	AutoCreatedCapacity *string `json:"AutoCreatedCapacity,omitempty" xml:"AutoCreatedCapacity,omitempty"`
	Cause               *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Progress            *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ScalingActivityId   *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	ScalingGroupId      *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatusCode          *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	StatusMessage       *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	TotalCapacity       *string `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
}

func (s DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetAttachedCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.AttachedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetAutoCreatedCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.AutoCreatedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetCause(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.Cause = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetDescription(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.Description = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetEndTime(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.EndTime = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetProgress(v int32) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.Progress = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetScalingActivityId(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetScalingGroupId(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetStartTime(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.StartTime = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetStatusCode(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetStatusMessage(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.StatusMessage = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity) SetTotalCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesScalingActivity {
	s.TotalCapacity = &v
	return s
}

type DescribeScalingActivitiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingActivitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingActivitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivitiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponse) SetHeaders(v map[string]*string) *DescribeScalingActivitiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingActivitiesResponse) SetStatusCode(v int32) *DescribeScalingActivitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesResponse) SetBody(v *DescribeScalingActivitiesResponseBody) *DescribeScalingActivitiesResponse {
	s.Body = v
	return s
}

type DescribeScalingActivityDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingActivityId    *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeScalingActivityDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivityDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailRequest) SetOwnerId(v int64) *DescribeScalingActivityDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetResourceOwnerAccount(v string) *DescribeScalingActivityDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetResourceOwnerId(v int64) *DescribeScalingActivityDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingActivityDetailRequest) SetScalingActivityId(v string) *DescribeScalingActivityDetailRequest {
	s.ScalingActivityId = &v
	return s
}

type DescribeScalingActivityDetailResponseBody struct {
	Detail            *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DescribeScalingActivityDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivityDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailResponseBody) SetDetail(v string) *DescribeScalingActivityDetailResponseBody {
	s.Detail = &v
	return s
}

func (s *DescribeScalingActivityDetailResponseBody) SetScalingActivityId(v string) *DescribeScalingActivityDetailResponseBody {
	s.ScalingActivityId = &v
	return s
}

type DescribeScalingActivityDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingActivityDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingActivityDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingActivityDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivityDetailResponse) SetHeaders(v map[string]*string) *DescribeScalingActivityDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingActivityDetailResponse) SetStatusCode(v int32) *DescribeScalingActivityDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivityDetailResponse) SetBody(v *DescribeScalingActivityDetailResponseBody) *DescribeScalingActivityDetailResponse {
	s.Body = v
	return s
}

type DescribeScalingConfigurationsRequest struct {
	ScalingConfigurationId   []*string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty" type:"Repeated"`
	ScalingConfigurationName []*string `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty" type:"Repeated"`
	OwnerAccount             *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber               *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                 *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeScalingConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsRequest) SetScalingConfigurationId(v []*string) *DescribeScalingConfigurationsRequest {
	s.ScalingConfigurationId = v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingConfigurationName(v []*string) *DescribeScalingConfigurationsRequest {
	s.ScalingConfigurationName = v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetOwnerAccount(v string) *DescribeScalingConfigurationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetOwnerId(v int64) *DescribeScalingConfigurationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetPageNumber(v int32) *DescribeScalingConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetPageSize(v int32) *DescribeScalingConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetRegionId(v string) *DescribeScalingConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetResourceOwnerAccount(v string) *DescribeScalingConfigurationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetResourceOwnerId(v int64) *DescribeScalingConfigurationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingGroupId(v string) *DescribeScalingConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeScalingConfigurationsResponseBody struct {
	PageNumber            *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingConfigurations *DescribeScalingConfigurationsResponseBodyScalingConfigurations `json:"ScalingConfigurations,omitempty" xml:"ScalingConfigurations,omitempty" type:"Struct"`
	TotalCount            *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBody) SetPageNumber(v int32) *DescribeScalingConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetPageSize(v int32) *DescribeScalingConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetRequestId(v string) *DescribeScalingConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetScalingConfigurations(v *DescribeScalingConfigurationsResponseBodyScalingConfigurations) *DescribeScalingConfigurationsResponseBody {
	s.ScalingConfigurations = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBody) SetTotalCount(v int32) *DescribeScalingConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurations struct {
	ScalingConfiguration []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration `json:"ScalingConfiguration,omitempty" xml:"ScalingConfiguration,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurations) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurations) SetScalingConfiguration(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) *DescribeScalingConfigurationsResponseBodyScalingConfigurations {
	s.ScalingConfiguration = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration struct {
	CreationTime             *string                                                                                      `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DataDisks                *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Struct"`
	ImageId                  *string                                                                                      `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceGeneration       *string                                                                                      `json:"InstanceGeneration,omitempty" xml:"InstanceGeneration,omitempty"`
	InstanceType             *string                                                                                      `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InternetChargeType       *string                                                                                      `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn   *int32                                                                                       `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut  *int32                                                                                       `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	IoOptimized              *string                                                                                      `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	LifecycleState           *string                                                                                      `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	ScalingConfigurationId   *string                                                                                      `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingConfigurationName *string                                                                                      `json:"ScalingConfigurationName,omitempty" xml:"ScalingConfigurationName,omitempty"`
	ScalingGroupId           *string                                                                                      `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	SecurityGroupId          *string                                                                                      `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SystemDiskCategory       *string                                                                                      `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	SystemDiskSize           *int32                                                                                       `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetCreationTime(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetDataDisks(v *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.DataDisks = v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetImageId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ImageId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInstanceGeneration(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InstanceGeneration = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInstanceType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InstanceType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInternetChargeType(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInternetMaxBandwidthIn(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetInternetMaxBandwidthOut(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetIoOptimized(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.IoOptimized = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetLifecycleState(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetScalingConfigurationId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetScalingConfigurationName(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ScalingConfigurationName = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetScalingGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSecurityGroupId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskCategory(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration) SetSystemDiskSize(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfiguration {
	s.SystemDiskSize = &v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks struct {
	DataDisk []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks) SetDataDisk(v []*DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisks {
	s.DataDisk = v
	return s
}

type DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk struct {
	Category   *string `json:"Category,omitempty" xml:"Category,omitempty"`
	Device     *string `json:"Device,omitempty" xml:"Device,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetCategory(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetDevice(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Device = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetSize(v int32) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk) SetSnapshotId(v string) *DescribeScalingConfigurationsResponseBodyScalingConfigurationsScalingConfigurationDataDisksDataDisk {
	s.SnapshotId = &v
	return s
}

type DescribeScalingConfigurationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsResponse) SetHeaders(v map[string]*string) *DescribeScalingConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingConfigurationsResponse) SetStatusCode(v int32) *DescribeScalingConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingConfigurationsResponse) SetBody(v *DescribeScalingConfigurationsResponseBody) *DescribeScalingConfigurationsResponse {
	s.Body = v
	return s
}

type DescribeScalingGroupsRequest struct {
	ScalingGroupId   []*string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty" type:"Repeated"`
	ScalingGroupName []*string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty" type:"Repeated"`
	OwnerAccount     *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber       *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DescribeScalingGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsRequest) SetScalingGroupId(v []*string) *DescribeScalingGroupsRequest {
	s.ScalingGroupId = v
	return s
}

func (s *DescribeScalingGroupsRequest) SetScalingGroupName(v []*string) *DescribeScalingGroupsRequest {
	s.ScalingGroupName = v
	return s
}

func (s *DescribeScalingGroupsRequest) SetOwnerAccount(v string) *DescribeScalingGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetOwnerId(v int64) *DescribeScalingGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetPageNumber(v int32) *DescribeScalingGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetPageSize(v int32) *DescribeScalingGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetRegionId(v string) *DescribeScalingGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupsRequest) SetResourceOwnerAccount(v string) *DescribeScalingGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type DescribeScalingGroupsResponseBody struct {
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingGroups *DescribeScalingGroupsResponseBodyScalingGroups `json:"ScalingGroups,omitempty" xml:"ScalingGroups,omitempty" type:"Struct"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBody) SetPageNumber(v int32) *DescribeScalingGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetPageSize(v int32) *DescribeScalingGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetRequestId(v string) *DescribeScalingGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetScalingGroups(v *DescribeScalingGroupsResponseBodyScalingGroups) *DescribeScalingGroupsResponseBody {
	s.ScalingGroups = v
	return s
}

func (s *DescribeScalingGroupsResponseBody) SetTotalCount(v int32) *DescribeScalingGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroups struct {
	ScalingGroup []*DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup `json:"ScalingGroup,omitempty" xml:"ScalingGroup,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroups) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroups) SetScalingGroup(v []*DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) *DescribeScalingGroupsResponseBodyScalingGroups {
	s.ScalingGroup = v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup struct {
	ActiveCapacity               *int32                                                                     `json:"ActiveCapacity,omitempty" xml:"ActiveCapacity,omitempty"`
	ActiveScalingConfigurationId *string                                                                    `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	CreationTime                 *string                                                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBInstanceIds                *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupDBInstanceIds   `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty" type:"Struct"`
	DefaultCooldown              *int32                                                                     `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	LifecycleState               *string                                                                    `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	LoadBalancerIds              *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupLoadBalancerIds `json:"LoadBalancerIds,omitempty" xml:"LoadBalancerIds,omitempty" type:"Struct"`
	MaxSize                      *int32                                                                     `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	MinSize                      *int32                                                                     `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	PendingCapacity              *int32                                                                     `json:"PendingCapacity,omitempty" xml:"PendingCapacity,omitempty"`
	RegionId                     *string                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemovalPolicies              *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupRemovalPolicies `json:"RemovalPolicies,omitempty" xml:"RemovalPolicies,omitempty" type:"Struct"`
	RemovingCapacity             *int32                                                                     `json:"RemovingCapacity,omitempty" xml:"RemovingCapacity,omitempty"`
	ScalingGroupId               *string                                                                    `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingGroupName             *string                                                                    `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
	TotalCapacity                *int32                                                                     `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	VSwitchId                    *string                                                                    `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetActiveCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.ActiveCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetActiveScalingConfigurationId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetCreationTime(v string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetDBInstanceIds(v *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupDBInstanceIds) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.DBInstanceIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetDefaultCooldown(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.DefaultCooldown = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetLifecycleState(v string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetLoadBalancerIds(v *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupLoadBalancerIds) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.LoadBalancerIds = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetMaxSize(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetMinSize(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetPendingCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.PendingCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetRegionId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetRemovalPolicies(v *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupRemovalPolicies) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.RemovalPolicies = v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetRemovingCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.RemovingCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetScalingGroupId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetScalingGroupName(v string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.ScalingGroupName = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetTotalCapacity(v int32) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup) SetVSwitchId(v string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroup {
	s.VSwitchId = &v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupDBInstanceIds struct {
	DBInstanceId []*string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupDBInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupDBInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupDBInstanceIds) SetDBInstanceId(v []*string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupDBInstanceIds {
	s.DBInstanceId = v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupLoadBalancerIds struct {
	LoadBalancerId []*string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupLoadBalancerIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupLoadBalancerIds) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupLoadBalancerIds) SetLoadBalancerId(v []*string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupLoadBalancerIds {
	s.LoadBalancerId = v
	return s
}

type DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupRemovalPolicies struct {
	RemovalPolicy []*string `json:"RemovalPolicy,omitempty" xml:"RemovalPolicy,omitempty" type:"Repeated"`
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupRemovalPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupRemovalPolicies) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupRemovalPolicies) SetRemovalPolicy(v []*string) *DescribeScalingGroupsResponseBodyScalingGroupsScalingGroupRemovalPolicies {
	s.RemovalPolicy = v
	return s
}

type DescribeScalingGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingGroupsResponse) SetHeaders(v map[string]*string) *DescribeScalingGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingGroupsResponse) SetStatusCode(v int32) *DescribeScalingGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingGroupsResponse) SetBody(v *DescribeScalingGroupsResponseBody) *DescribeScalingGroupsResponse {
	s.Body = v
	return s
}

type DescribeScalingInstancesRequest struct {
	InstanceId     []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	CreationType   *string   `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	HealthStatus   *string   `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	LifecycleState *string   `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	OwnerAccount   *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber     *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingGroupId         *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeScalingInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesRequest) SetInstanceId(v []*string) *DescribeScalingInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeScalingInstancesRequest) SetCreationType(v string) *DescribeScalingInstancesRequest {
	s.CreationType = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetHealthStatus(v string) *DescribeScalingInstancesRequest {
	s.HealthStatus = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetLifecycleState(v string) *DescribeScalingInstancesRequest {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetOwnerAccount(v string) *DescribeScalingInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetOwnerId(v int64) *DescribeScalingInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetPageNumber(v int32) *DescribeScalingInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetPageSize(v int32) *DescribeScalingInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetRegionId(v string) *DescribeScalingInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetResourceOwnerAccount(v string) *DescribeScalingInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetResourceOwnerId(v int64) *DescribeScalingInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingConfigurationId(v string) *DescribeScalingInstancesRequest {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingInstancesRequest) SetScalingGroupId(v string) *DescribeScalingInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeScalingInstancesResponseBody struct {
	PageNumber       *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingInstances *DescribeScalingInstancesResponseBodyScalingInstances `json:"ScalingInstances,omitempty" xml:"ScalingInstances,omitempty" type:"Struct"`
	TotalCount       *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBody) SetPageNumber(v int32) *DescribeScalingInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetPageSize(v int32) *DescribeScalingInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetRequestId(v string) *DescribeScalingInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetScalingInstances(v *DescribeScalingInstancesResponseBodyScalingInstances) *DescribeScalingInstancesResponseBody {
	s.ScalingInstances = v
	return s
}

func (s *DescribeScalingInstancesResponseBody) SetTotalCount(v int32) *DescribeScalingInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingInstancesResponseBodyScalingInstances struct {
	ScalingInstance []*DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance `json:"ScalingInstance,omitempty" xml:"ScalingInstance,omitempty" type:"Repeated"`
}

func (s DescribeScalingInstancesResponseBodyScalingInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponseBodyScalingInstances) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBodyScalingInstances) SetScalingInstance(v []*DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) *DescribeScalingInstancesResponseBodyScalingInstances {
	s.ScalingInstance = v
	return s
}

type DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance struct {
	CreationTime           *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CreationType           *string `json:"CreationType,omitempty" xml:"CreationType,omitempty"`
	HealthStatus           *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LifecycleState         *string `json:"LifecycleState,omitempty" xml:"LifecycleState,omitempty"`
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
	ScalingGroupId         *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetCreationTime(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetCreationType(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.CreationType = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetHealthStatus(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.HealthStatus = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetInstanceId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetLifecycleState(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.LifecycleState = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetScalingConfigurationId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.ScalingConfigurationId = &v
	return s
}

func (s *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance) SetScalingGroupId(v string) *DescribeScalingInstancesResponseBodyScalingInstancesScalingInstance {
	s.ScalingGroupId = &v
	return s
}

type DescribeScalingInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingInstancesResponse) SetHeaders(v map[string]*string) *DescribeScalingInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingInstancesResponse) SetStatusCode(v int32) *DescribeScalingInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingInstancesResponse) SetBody(v *DescribeScalingInstancesResponseBody) *DescribeScalingInstancesResponse {
	s.Body = v
	return s
}

type DescribeScalingRulesRequest struct {
	ScalingRuleAri  []*string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty" type:"Repeated"`
	ScalingRuleId   []*string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty" type:"Repeated"`
	ScalingRuleName []*string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty" type:"Repeated"`
	OwnerAccount    *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber      *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScalingGroupId       *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeScalingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesRequest) SetScalingRuleAri(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleAri = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleId(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleId = v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingRuleName(v []*string) *DescribeScalingRulesRequest {
	s.ScalingRuleName = v
	return s
}

func (s *DescribeScalingRulesRequest) SetOwnerAccount(v string) *DescribeScalingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetOwnerId(v int64) *DescribeScalingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetPageNumber(v int32) *DescribeScalingRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetPageSize(v int32) *DescribeScalingRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetRegionId(v string) *DescribeScalingRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetResourceOwnerAccount(v string) *DescribeScalingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetResourceOwnerId(v int64) *DescribeScalingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingRulesRequest) SetScalingGroupId(v string) *DescribeScalingRulesRequest {
	s.ScalingGroupId = &v
	return s
}

type DescribeScalingRulesResponseBody struct {
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingRules *DescribeScalingRulesResponseBodyScalingRules `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Struct"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBody) SetPageNumber(v int32) *DescribeScalingRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetPageSize(v int32) *DescribeScalingRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetRequestId(v string) *DescribeScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetScalingRules(v *DescribeScalingRulesResponseBodyScalingRules) *DescribeScalingRulesResponseBody {
	s.ScalingRules = v
	return s
}

func (s *DescribeScalingRulesResponseBody) SetTotalCount(v int32) *DescribeScalingRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScalingRulesResponseBodyScalingRules struct {
	ScalingRule []*DescribeScalingRulesResponseBodyScalingRulesScalingRule `json:"ScalingRule,omitempty" xml:"ScalingRule,omitempty" type:"Repeated"`
}

func (s DescribeScalingRulesResponseBodyScalingRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRules) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRules) SetScalingRule(v []*DescribeScalingRulesResponseBodyScalingRulesScalingRule) *DescribeScalingRulesResponseBodyScalingRules {
	s.ScalingRule = v
	return s
}

type DescribeScalingRulesResponseBodyScalingRulesScalingRule struct {
	AdjustmentType  *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue *int32  `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	Cooldown        *int32  `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	MaxSize         *int32  `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	MinSize         *int32  `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	ScalingGroupId  *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleAri  *string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
	ScalingRuleId   *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponseBodyScalingRulesScalingRule) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetAdjustmentType(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.AdjustmentType = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetAdjustmentValue(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.AdjustmentValue = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetCooldown(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.Cooldown = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetMaxSize(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.MaxSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetMinSize(v int32) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.MinSize = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingGroupId(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingRuleAri(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingRuleAri = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingRuleId(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingRuleId = &v
	return s
}

func (s *DescribeScalingRulesResponseBodyScalingRulesScalingRule) SetScalingRuleName(v string) *DescribeScalingRulesResponseBodyScalingRulesScalingRule {
	s.ScalingRuleName = &v
	return s
}

type DescribeScalingRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScalingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeScalingRulesResponse) SetHeaders(v map[string]*string) *DescribeScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeScalingRulesResponse) SetStatusCode(v int32) *DescribeScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingRulesResponse) SetBody(v *DescribeScalingRulesResponseBody) *DescribeScalingRulesResponse {
	s.Body = v
	return s
}

type DescribeScheduledTasksRequest struct {
	ScheduledAction   []*string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty" type:"Repeated"`
	ScheduledTaskId   []*string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty" type:"Repeated"`
	ScheduledTaskName []*string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty" type:"Repeated"`
	OwnerAccount      *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber        *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeScheduledTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksRequest) SetScheduledAction(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledAction = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledTaskId(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledTaskId = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetScheduledTaskName(v []*string) *DescribeScheduledTasksRequest {
	s.ScheduledTaskName = v
	return s
}

func (s *DescribeScheduledTasksRequest) SetOwnerAccount(v string) *DescribeScheduledTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetOwnerId(v int64) *DescribeScheduledTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetPageNumber(v int32) *DescribeScheduledTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetPageSize(v int32) *DescribeScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetRegionId(v string) *DescribeScheduledTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetResourceOwnerAccount(v string) *DescribeScheduledTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScheduledTasksRequest) SetResourceOwnerId(v int64) *DescribeScheduledTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

type DescribeScheduledTasksResponseBody struct {
	PageNumber     *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScheduledTasks *DescribeScheduledTasksResponseBodyScheduledTasks `json:"ScheduledTasks,omitempty" xml:"ScheduledTasks,omitempty" type:"Struct"`
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScheduledTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBody) SetPageNumber(v int32) *DescribeScheduledTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetPageSize(v int32) *DescribeScheduledTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetRequestId(v string) *DescribeScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetScheduledTasks(v *DescribeScheduledTasksResponseBodyScheduledTasks) *DescribeScheduledTasksResponseBody {
	s.ScheduledTasks = v
	return s
}

func (s *DescribeScheduledTasksResponseBody) SetTotalCount(v int32) *DescribeScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeScheduledTasksResponseBodyScheduledTasks struct {
	ScheduledTask []*DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask `json:"ScheduledTask,omitempty" xml:"ScheduledTask,omitempty" type:"Repeated"`
}

func (s DescribeScheduledTasksResponseBodyScheduledTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyScheduledTasks) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasks) SetScheduledTask(v []*DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) *DescribeScheduledTasksResponseBodyScheduledTasks {
	s.ScheduledTask = v
	return s
}

type DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	ScheduledAction      *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	ScheduledTaskId      *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	ScheduledTaskName    *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	TaskEnabled          *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
}

func (s DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetDescription(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.Description = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetLaunchExpirationTime(v int32) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.LaunchExpirationTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetLaunchTime(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.LaunchTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetRecurrenceEndTime(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.RecurrenceEndTime = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetRecurrenceType(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.RecurrenceType = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetRecurrenceValue(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.RecurrenceValue = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetScheduledAction(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.ScheduledAction = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetScheduledTaskId(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.ScheduledTaskId = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetScheduledTaskName(v string) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.ScheduledTaskName = &v
	return s
}

func (s *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask) SetTaskEnabled(v bool) *DescribeScheduledTasksResponseBodyScheduledTasksScheduledTask {
	s.TaskEnabled = &v
	return s
}

type DescribeScheduledTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScheduledTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScheduledTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTasksResponse) SetHeaders(v map[string]*string) *DescribeScheduledTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduledTasksResponse) SetStatusCode(v int32) *DescribeScheduledTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduledTasksResponse) SetBody(v *DescribeScheduledTasksResponseBody) *DescribeScheduledTasksResponse {
	s.Body = v
	return s
}

type DetachInstancesRequest struct {
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DetachInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachInstancesRequest) GoString() string {
	return s.String()
}

func (s *DetachInstancesRequest) SetInstanceId(v []*string) *DetachInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DetachInstancesRequest) SetOwnerAccount(v string) *DetachInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachInstancesRequest) SetOwnerId(v int64) *DetachInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachInstancesRequest) SetResourceOwnerAccount(v string) *DetachInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachInstancesRequest) SetResourceOwnerId(v int64) *DetachInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachInstancesRequest) SetScalingGroupId(v string) *DetachInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type DetachInstancesResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s DetachInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DetachInstancesResponseBody) SetRequestId(v string) *DetachInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachInstancesResponseBody) SetScalingActivityId(v string) *DetachInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

type DetachInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachInstancesResponse) GoString() string {
	return s.String()
}

func (s *DetachInstancesResponse) SetHeaders(v map[string]*string) *DetachInstancesResponse {
	s.Headers = v
	return s
}

func (s *DetachInstancesResponse) SetStatusCode(v int32) *DetachInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachInstancesResponse) SetBody(v *DetachInstancesResponseBody) *DetachInstancesResponse {
	s.Body = v
	return s
}

type DisableScalingGroupRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// This parameter is required.
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DisableScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupRequest) SetOwnerAccount(v string) *DisableScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableScalingGroupRequest) SetOwnerId(v int64) *DisableScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableScalingGroupRequest) SetResourceOwnerAccount(v string) *DisableScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableScalingGroupRequest) SetScalingGroupId(v string) *DisableScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

type DisableScalingGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupResponseBody) SetRequestId(v string) *DisableScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

type DisableScalingGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *DisableScalingGroupResponse) SetHeaders(v map[string]*string) *DisableScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *DisableScalingGroupResponse) SetStatusCode(v int32) *DisableScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableScalingGroupResponse) SetBody(v *DisableScalingGroupResponseBody) *DisableScalingGroupResponse {
	s.Body = v
	return s
}

type EnableScalingGroupRequest struct {
	InstanceId                   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	ActiveScalingConfigurationId *string   `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	OwnerAccount                 *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount         *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// This parameter is required.
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s EnableScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupRequest) SetInstanceId(v []*string) *EnableScalingGroupRequest {
	s.InstanceId = v
	return s
}

func (s *EnableScalingGroupRequest) SetActiveScalingConfigurationId(v string) *EnableScalingGroupRequest {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetOwnerAccount(v string) *EnableScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *EnableScalingGroupRequest) SetOwnerId(v int64) *EnableScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableScalingGroupRequest) SetResourceOwnerAccount(v string) *EnableScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *EnableScalingGroupRequest) SetScalingGroupId(v string) *EnableScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

type EnableScalingGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupResponseBody) SetRequestId(v string) *EnableScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

type EnableScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *EnableScalingGroupResponse) SetHeaders(v map[string]*string) *EnableScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *EnableScalingGroupResponse) SetStatusCode(v int32) *EnableScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableScalingGroupResponse) SetBody(v *EnableScalingGroupResponseBody) *EnableScalingGroupResponse {
	s.Body = v
	return s
}

type ExecuteScalingRuleRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingRuleAri *string `json:"ScalingRuleAri,omitempty" xml:"ScalingRuleAri,omitempty"`
}

func (s ExecuteScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ExecuteScalingRuleRequest) SetClientToken(v string) *ExecuteScalingRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetOwnerAccount(v string) *ExecuteScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetOwnerId(v int64) *ExecuteScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetResourceOwnerAccount(v string) *ExecuteScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetResourceOwnerId(v int64) *ExecuteScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ExecuteScalingRuleRequest) SetScalingRuleAri(v string) *ExecuteScalingRuleRequest {
	s.ScalingRuleAri = &v
	return s
}

type ExecuteScalingRuleResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s ExecuteScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteScalingRuleResponseBody) SetRequestId(v string) *ExecuteScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteScalingRuleResponseBody) SetScalingActivityId(v string) *ExecuteScalingRuleResponseBody {
	s.ScalingActivityId = &v
	return s
}

type ExecuteScalingRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ExecuteScalingRuleResponse) SetHeaders(v map[string]*string) *ExecuteScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ExecuteScalingRuleResponse) SetStatusCode(v int32) *ExecuteScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteScalingRuleResponse) SetBody(v *ExecuteScalingRuleResponseBody) *ExecuteScalingRuleResponse {
	s.Body = v
	return s
}

type ModifyScalingGroupRequest struct {
	RemovalPolicy                []*string `json:"RemovalPolicy,omitempty" xml:"RemovalPolicy,omitempty" type:"Repeated"`
	ActiveScalingConfigurationId *string   `json:"ActiveScalingConfigurationId,omitempty" xml:"ActiveScalingConfigurationId,omitempty"`
	DefaultCooldown              *int32    `json:"DefaultCooldown,omitempty" xml:"DefaultCooldown,omitempty"`
	MaxSize                      *int32    `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	MinSize                      *int32    `json:"MinSize,omitempty" xml:"MinSize,omitempty"`
	OwnerAccount                 *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount         *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// This parameter is required.
	ScalingGroupId   *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingGroupName *string `json:"ScalingGroupName,omitempty" xml:"ScalingGroupName,omitempty"`
}

func (s ModifyScalingGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupRequest) SetRemovalPolicy(v []*string) *ModifyScalingGroupRequest {
	s.RemovalPolicy = v
	return s
}

func (s *ModifyScalingGroupRequest) SetActiveScalingConfigurationId(v string) *ModifyScalingGroupRequest {
	s.ActiveScalingConfigurationId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetDefaultCooldown(v int32) *ModifyScalingGroupRequest {
	s.DefaultCooldown = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMaxSize(v int32) *ModifyScalingGroupRequest {
	s.MaxSize = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetMinSize(v int32) *ModifyScalingGroupRequest {
	s.MinSize = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOwnerAccount(v string) *ModifyScalingGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetOwnerId(v int64) *ModifyScalingGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetResourceOwnerAccount(v string) *ModifyScalingGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetScalingGroupId(v string) *ModifyScalingGroupRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyScalingGroupRequest) SetScalingGroupName(v string) *ModifyScalingGroupRequest {
	s.ScalingGroupName = &v
	return s
}

type ModifyScalingGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScalingGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupResponseBody) SetRequestId(v string) *ModifyScalingGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScalingGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScalingGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScalingGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingGroupResponse) SetHeaders(v map[string]*string) *ModifyScalingGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingGroupResponse) SetStatusCode(v int32) *ModifyScalingGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScalingGroupResponse) SetBody(v *ModifyScalingGroupResponseBody) *ModifyScalingGroupResponse {
	s.Body = v
	return s
}

type ModifyScalingRuleRequest struct {
	AdjustmentType       *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	AdjustmentValue      *int32  `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	Cooldown             *int32  `json:"Cooldown,omitempty" xml:"Cooldown,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingRuleId   *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
}

func (s ModifyScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequest) SetAdjustmentType(v string) *ModifyScalingRuleRequest {
	s.AdjustmentType = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetAdjustmentValue(v int32) *ModifyScalingRuleRequest {
	s.AdjustmentValue = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetCooldown(v int32) *ModifyScalingRuleRequest {
	s.Cooldown = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOwnerAccount(v string) *ModifyScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOwnerId(v int64) *ModifyScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetResourceOwnerAccount(v string) *ModifyScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetResourceOwnerId(v int64) *ModifyScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingRuleId(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingRuleName(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleName = &v
	return s
}

type ModifyScalingRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponseBody) SetRequestId(v string) *ModifyScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScalingRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponse) SetHeaders(v map[string]*string) *ModifyScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyScalingRuleResponse) SetStatusCode(v int32) *ModifyScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScalingRuleResponse) SetBody(v *ModifyScalingRuleResponseBody) *ModifyScalingRuleResponse {
	s.Body = v
	return s
}

type ModifyScheduledTaskRequest struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LaunchExpirationTime *int32  `json:"LaunchExpirationTime,omitempty" xml:"LaunchExpirationTime,omitempty"`
	LaunchTime           *string `json:"LaunchTime,omitempty" xml:"LaunchTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RecurrenceEndTime    *string `json:"RecurrenceEndTime,omitempty" xml:"RecurrenceEndTime,omitempty"`
	RecurrenceType       *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	RecurrenceValue      *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScheduledAction      *string `json:"ScheduledAction,omitempty" xml:"ScheduledAction,omitempty"`
	// This parameter is required.
	ScheduledTaskId   *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	ScheduledTaskName *string `json:"ScheduledTaskName,omitempty" xml:"ScheduledTaskName,omitempty"`
	TaskEnabled       *bool   `json:"TaskEnabled,omitempty" xml:"TaskEnabled,omitempty"`
}

func (s ModifyScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) SetDescription(v string) *ModifyScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetLaunchExpirationTime(v int32) *ModifyScheduledTaskRequest {
	s.LaunchExpirationTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetLaunchTime(v string) *ModifyScheduledTaskRequest {
	s.LaunchTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetOwnerAccount(v string) *ModifyScheduledTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetOwnerId(v int64) *ModifyScheduledTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceEndTime(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceEndTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceType(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceType = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetRecurrenceValue(v string) *ModifyScheduledTaskRequest {
	s.RecurrenceValue = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetResourceOwnerAccount(v string) *ModifyScheduledTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetResourceOwnerId(v int64) *ModifyScheduledTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledAction(v string) *ModifyScheduledTaskRequest {
	s.ScheduledAction = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledTaskId(v string) *ModifyScheduledTaskRequest {
	s.ScheduledTaskId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledTaskName(v string) *ModifyScheduledTaskRequest {
	s.ScheduledTaskName = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetTaskEnabled(v bool) *ModifyScheduledTaskRequest {
	s.TaskEnabled = &v
	return s
}

type ModifyScheduledTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBody) SetRequestId(v string) *ModifyScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

type ModifyScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponse) SetHeaders(v map[string]*string) *ModifyScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduledTaskResponse) SetStatusCode(v int32) *ModifyScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScheduledTaskResponse) SetBody(v *ModifyScheduledTaskResponseBody) *ModifyScheduledTaskResponse {
	s.Body = v
	return s
}

type RemoveInstancesRequest struct {
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s RemoveInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstancesRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstancesRequest) SetInstanceId(v []*string) *RemoveInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *RemoveInstancesRequest) SetOwnerAccount(v string) *RemoveInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveInstancesRequest) SetOwnerId(v int64) *RemoveInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveInstancesRequest) SetResourceOwnerAccount(v string) *RemoveInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveInstancesRequest) SetResourceOwnerId(v int64) *RemoveInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveInstancesRequest) SetScalingGroupId(v string) *RemoveInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

type RemoveInstancesResponseBody struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
}

func (s RemoveInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponseBody) SetRequestId(v string) *RemoveInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInstancesResponseBody) SetScalingActivityId(v string) *RemoveInstancesResponseBody {
	s.ScalingActivityId = &v
	return s
}

type RemoveInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveInstancesResponse) GoString() string {
	return s.String()
}

func (s *RemoveInstancesResponse) SetHeaders(v map[string]*string) *RemoveInstancesResponse {
	s.Headers = v
	return s
}

func (s *RemoveInstancesResponse) SetStatusCode(v int32) *RemoveInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInstancesResponse) SetBody(v *RemoveInstancesResponseBody) *RemoveInstancesResponse {
	s.Body = v
	return s
}

type VerifyAuthenticationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Uid                  *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s VerifyAuthenticationRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyAuthenticationRequest) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationRequest) SetOwnerId(v int64) *VerifyAuthenticationRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetResourceOwnerAccount(v string) *VerifyAuthenticationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetResourceOwnerId(v int64) *VerifyAuthenticationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VerifyAuthenticationRequest) SetUid(v int64) *VerifyAuthenticationRequest {
	s.Uid = &v
	return s
}

type VerifyAuthenticationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyAuthenticationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyAuthenticationResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationResponseBody) SetRequestId(v string) *VerifyAuthenticationResponseBody {
	s.RequestId = &v
	return s
}

type VerifyAuthenticationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyAuthenticationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyAuthenticationResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyAuthenticationResponse) GoString() string {
	return s.String()
}

func (s *VerifyAuthenticationResponse) SetHeaders(v map[string]*string) *VerifyAuthenticationResponse {
	s.Headers = v
	return s
}

func (s *VerifyAuthenticationResponse) SetStatusCode(v int32) *VerifyAuthenticationResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAuthenticationResponse) SetBody(v *VerifyAuthenticationResponseBody) *VerifyAuthenticationResponse {
	s.Body = v
	return s
}

type VerifyUserRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s VerifyUserRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserRequest) GoString() string {
	return s.String()
}

func (s *VerifyUserRequest) SetOwnerId(v int64) *VerifyUserRequest {
	s.OwnerId = &v
	return s
}

func (s *VerifyUserRequest) SetResourceOwnerAccount(v string) *VerifyUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

type VerifyUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s VerifyUserResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyUserResponse) GoString() string {
	return s.String()
}

func (s *VerifyUserResponse) SetHeaders(v map[string]*string) *VerifyUserResponse {
	s.Headers = v
	return s
}

func (s *VerifyUserResponse) SetStatusCode(v int32) *VerifyUserResponse {
	s.StatusCode = &v
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
		"cn-qingdao":                  tea.String("ess.aliyuncs.com"),
		"cn-beijing":                  tea.String("ess.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("ess.aliyuncs.com"),
		"cn-shanghai":                 tea.String("ess.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("ess.aliyuncs.com"),
		"cn-hongkong":                 tea.String("ess.aliyuncs.com"),
		"ap-southeast-1":              tea.String("ess.aliyuncs.com"),
		"us-east-1":                   tea.String("ess.aliyuncs.com"),
		"us-west-1":                   tea.String("ess.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("ess.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("ess.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("ess.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("ess.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("ess.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("ess.aliyuncs.com"),
		"cn-edge-1":                   tea.String("ess.aliyuncs.com"),
		"cn-fujian":                   tea.String("ess.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("ess.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("ess.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("ess.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("ess.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("ess.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("ess.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("ess.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("ess.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("ess.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("ess.aliyuncs.com"),
		"cn-wuhan":                    tea.String("ess.aliyuncs.com"),
		"cn-yushanfang":               tea.String("ess.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("ess.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("ess.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("ess.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("ess.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("ess.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("ess.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ess"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Deprecated: OpenAPI AttachInstances is deprecated, please use Ess::2014-08-28::AttachInstances,Ess::2022-02-22::AttachInstances instead.
//
// @param request - AttachInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachInstancesResponse
// Deprecated
func (client *Client) AttachInstancesWithOptions(request *AttachInstancesRequest, runtime *util.RuntimeOptions) (_result *AttachInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachInstances"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI AttachInstances is deprecated, please use Ess::2014-08-28::AttachInstances,Ess::2022-02-22::AttachInstances instead.
//
// @param request - AttachInstancesRequest
//
// @return AttachInstancesResponse
// Deprecated
func (client *Client) AttachInstances(request *AttachInstancesRequest) (_result *AttachInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachInstancesResponse{}
	_body, _err := client.AttachInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateScalingConfiguration is deprecated, please use Ess::2022-02-22::CreateScalingConfiguration,Ess::2014-08-28::CreateScalingConfiguration instead.
//
// @param request - CreateScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScalingConfigurationResponse
// Deprecated
func (client *Client) CreateScalingConfigurationWithOptions(request *CreateScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateScalingConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetChargeType)) {
		query["InternetChargeType"] = request.InternetChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InternetMaxBandwidthIn)) {
		query["InternetMaxBandwidthIn"] = request.InternetMaxBandwidthIn
	}

	if !tea.BoolValue(util.IsUnset(request.InternetMaxBandwidthOut)) {
		query["InternetMaxBandwidthOut"] = request.InternetMaxBandwidthOut
	}

	if !tea.BoolValue(util.IsUnset(request.IoOptimized)) {
		query["IoOptimized"] = request.IoOptimized
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationName)) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DataDisk)) {
		query["DataDisk"] = request.DataDisk
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDisk)) {
		query["SystemDisk"] = request.SystemDisk
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScalingConfiguration"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScalingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateScalingConfiguration is deprecated, please use Ess::2022-02-22::CreateScalingConfiguration,Ess::2014-08-28::CreateScalingConfiguration instead.
//
// @param request - CreateScalingConfigurationRequest
//
// @return CreateScalingConfigurationResponse
// Deprecated
func (client *Client) CreateScalingConfiguration(request *CreateScalingConfigurationRequest) (_result *CreateScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScalingConfigurationResponse{}
	_body, _err := client.CreateScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateScalingGroup is deprecated, please use Ess::2022-02-22::CreateScalingGroup,Ess::2014-08-28::CreateScalingGroup instead.
//
// @param request - CreateScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScalingGroupResponse
// Deprecated
func (client *Client) CreateScalingGroupWithOptions(request *CreateScalingGroupRequest, runtime *util.RuntimeOptions) (_result *CreateScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DBInstanceIds)) {
		query["DBInstanceIds"] = request.DBInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCooldown)) {
		query["DefaultCooldown"] = request.DefaultCooldown
	}

	if !tea.BoolValue(util.IsUnset(request.LoadBalancerIds)) {
		query["LoadBalancerIds"] = request.LoadBalancerIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSize)) {
		query["MaxSize"] = request.MaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinSize)) {
		query["MinSize"] = request.MinSize
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupName)) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.RemovalPolicy)) {
		query["RemovalPolicy"] = request.RemovalPolicy
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScalingGroup"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateScalingGroup is deprecated, please use Ess::2022-02-22::CreateScalingGroup,Ess::2014-08-28::CreateScalingGroup instead.
//
// @param request - CreateScalingGroupRequest
//
// @return CreateScalingGroupResponse
// Deprecated
func (client *Client) CreateScalingGroup(request *CreateScalingGroupRequest) (_result *CreateScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScalingGroupResponse{}
	_body, _err := client.CreateScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateScalingRule is deprecated, please use Ess::2022-02-22::CreateScalingRule,Ess::2014-08-28::CreateScalingRule instead.
//
// @param request - CreateScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScalingRuleResponse
// Deprecated
func (client *Client) CreateScalingRuleWithOptions(request *CreateScalingRuleRequest, runtime *util.RuntimeOptions) (_result *CreateScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdjustmentType)) {
		query["AdjustmentType"] = request.AdjustmentType
	}

	if !tea.BoolValue(util.IsUnset(request.AdjustmentValue)) {
		query["AdjustmentValue"] = request.AdjustmentValue
	}

	if !tea.BoolValue(util.IsUnset(request.Cooldown)) {
		query["Cooldown"] = request.Cooldown
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScalingRule"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateScalingRule is deprecated, please use Ess::2022-02-22::CreateScalingRule,Ess::2014-08-28::CreateScalingRule instead.
//
// @param request - CreateScalingRuleRequest
//
// @return CreateScalingRuleResponse
// Deprecated
func (client *Client) CreateScalingRule(request *CreateScalingRuleRequest) (_result *CreateScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScalingRuleResponse{}
	_body, _err := client.CreateScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateScheduledTask is deprecated, please use Ess::2022-02-22::CreateScheduledTask,Ess::2014-08-28::CreateScheduledTask instead.
//
// @param request - CreateScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledTaskResponse
// Deprecated
func (client *Client) CreateScheduledTaskWithOptions(request *CreateScheduledTaskRequest, runtime *util.RuntimeOptions) (_result *CreateScheduledTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchExpirationTime)) {
		query["LaunchExpirationTime"] = request.LaunchExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTime)) {
		query["LaunchTime"] = request.LaunchTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceEndTime)) {
		query["RecurrenceEndTime"] = request.RecurrenceEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceType)) {
		query["RecurrenceType"] = request.RecurrenceType
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceValue)) {
		query["RecurrenceValue"] = request.RecurrenceValue
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledAction)) {
		query["ScheduledAction"] = request.ScheduledAction
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskName)) {
		query["ScheduledTaskName"] = request.ScheduledTaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskEnabled)) {
		query["TaskEnabled"] = request.TaskEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduledTask"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateScheduledTask is deprecated, please use Ess::2022-02-22::CreateScheduledTask,Ess::2014-08-28::CreateScheduledTask instead.
//
// @param request - CreateScheduledTaskRequest
//
// @return CreateScheduledTaskResponse
// Deprecated
func (client *Client) CreateScheduledTask(request *CreateScheduledTaskRequest) (_result *CreateScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CreateScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteScalingConfiguration is deprecated, please use Ess::2022-02-22::DeleteScalingConfiguration,Ess::2014-08-28::DeleteScalingConfiguration instead.
//
// @param request - DeleteScalingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScalingConfigurationResponse
// Deprecated
func (client *Client) DeleteScalingConfigurationWithOptions(request *DeleteScalingConfigurationRequest, runtime *util.RuntimeOptions) (_result *DeleteScalingConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationId)) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScalingConfiguration"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScalingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteScalingConfiguration is deprecated, please use Ess::2022-02-22::DeleteScalingConfiguration,Ess::2014-08-28::DeleteScalingConfiguration instead.
//
// @param request - DeleteScalingConfigurationRequest
//
// @return DeleteScalingConfigurationResponse
// Deprecated
func (client *Client) DeleteScalingConfiguration(request *DeleteScalingConfigurationRequest) (_result *DeleteScalingConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScalingConfigurationResponse{}
	_body, _err := client.DeleteScalingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteScalingGroup is deprecated, please use Ess::2022-02-22::DeleteScalingGroup,Ess::2014-08-28::DeleteScalingGroup instead.
//
// @param request - DeleteScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScalingGroupResponse
// Deprecated
func (client *Client) DeleteScalingGroupWithOptions(request *DeleteScalingGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForceDelete)) {
		query["ForceDelete"] = request.ForceDelete
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScalingGroup"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteScalingGroup is deprecated, please use Ess::2022-02-22::DeleteScalingGroup,Ess::2014-08-28::DeleteScalingGroup instead.
//
// @param request - DeleteScalingGroupRequest
//
// @return DeleteScalingGroupResponse
// Deprecated
func (client *Client) DeleteScalingGroup(request *DeleteScalingGroupRequest) (_result *DeleteScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScalingGroupResponse{}
	_body, _err := client.DeleteScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteScalingRule is deprecated, please use Ess::2022-02-22::DeleteScalingRule,Ess::2014-08-28::DeleteScalingRule instead.
//
// @param request - DeleteScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScalingRuleResponse
// Deprecated
func (client *Client) DeleteScalingRuleWithOptions(request *DeleteScalingRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleId)) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScalingRule"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteScalingRule is deprecated, please use Ess::2022-02-22::DeleteScalingRule,Ess::2014-08-28::DeleteScalingRule instead.
//
// @param request - DeleteScalingRuleRequest
//
// @return DeleteScalingRuleResponse
// Deprecated
func (client *Client) DeleteScalingRule(request *DeleteScalingRuleRequest) (_result *DeleteScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScalingRuleResponse{}
	_body, _err := client.DeleteScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteScheduledTask is deprecated, please use Ess::2022-02-22::DeleteScheduledTask,Ess::2014-08-28::DeleteScheduledTask instead.
//
// @param request - DeleteScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScheduledTaskResponse
// Deprecated
func (client *Client) DeleteScheduledTaskWithOptions(request *DeleteScheduledTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteScheduledTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskId)) {
		query["ScheduledTaskId"] = request.ScheduledTaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteScheduledTask"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteScheduledTask is deprecated, please use Ess::2022-02-22::DeleteScheduledTask,Ess::2014-08-28::DeleteScheduledTask instead.
//
// @param request - DeleteScheduledTaskRequest
//
// @return DeleteScheduledTaskResponse
// Deprecated
func (client *Client) DeleteScheduledTask(request *DeleteScheduledTaskRequest) (_result *DeleteScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteScheduledTaskResponse{}
	_body, _err := client.DeleteScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeAccountAttributes is deprecated, please use Ess::2022-02-22::DescribeLimitation,Ess::2014-08-28::DescribeLimitation instead.
//
// @param request - DescribeAccountAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccountAttributesResponse
// Deprecated
func (client *Client) DescribeAccountAttributesWithOptions(request *DescribeAccountAttributesRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAccountAttributes"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAccountAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeAccountAttributes is deprecated, please use Ess::2022-02-22::DescribeLimitation,Ess::2014-08-28::DescribeLimitation instead.
//
// @param request - DescribeAccountAttributesRequest
//
// @return DescribeAccountAttributesResponse
// Deprecated
func (client *Client) DescribeAccountAttributes(request *DescribeAccountAttributesRequest) (_result *DescribeAccountAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountAttributesResponse{}
	_body, _err := client.DescribeAccountAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeCapacityHistory is deprecated, please use Ess::2022-02-22::DescribeScalingActivities,Ess::2014-08-28::DescribeScalingActivities instead.
//
// @param request - DescribeCapacityHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCapacityHistoryResponse
// Deprecated
func (client *Client) DescribeCapacityHistoryWithOptions(request *DescribeCapacityHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeCapacityHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCapacityHistory"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCapacityHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeCapacityHistory is deprecated, please use Ess::2022-02-22::DescribeScalingActivities,Ess::2014-08-28::DescribeScalingActivities instead.
//
// @param request - DescribeCapacityHistoryRequest
//
// @return DescribeCapacityHistoryResponse
// Deprecated
func (client *Client) DescribeCapacityHistory(request *DescribeCapacityHistoryRequest) (_result *DescribeCapacityHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCapacityHistoryResponse{}
	_body, _err := client.DescribeCapacityHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeLimitation is deprecated, please use Ess::2022-02-22::DescribeLimitation,Ess::2014-08-28::DescribeLimitation instead.
//
// @param request - DescribeLimitationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeLimitationResponse
// Deprecated
func (client *Client) DescribeLimitationWithOptions(request *DescribeLimitationRequest, runtime *util.RuntimeOptions) (_result *DescribeLimitationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLimitation"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLimitationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeLimitation is deprecated, please use Ess::2022-02-22::DescribeLimitation,Ess::2014-08-28::DescribeLimitation instead.
//
// @param request - DescribeLimitationRequest
//
// @return DescribeLimitationResponse
// Deprecated
func (client *Client) DescribeLimitation(request *DescribeLimitationRequest) (_result *DescribeLimitationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLimitationResponse{}
	_body, _err := client.DescribeLimitationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeRegions is deprecated, please use Ess::2022-02-22::DescribeRegions,Ess::2014-08-28::DescribeRegions instead.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
// Deprecated
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeRegions is deprecated, please use Ess::2022-02-22::DescribeRegions,Ess::2014-08-28::DescribeRegions instead.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
// Deprecated
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingActivities is deprecated, please use Ess::2022-02-22::DescribeScalingActivities,Ess::2014-08-28::DescribeScalingActivities instead.
//
// @param request - DescribeScalingActivitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingActivitiesResponse
// Deprecated
func (client *Client) DescribeScalingActivitiesWithOptions(request *DescribeScalingActivitiesRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingActivitiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StatusCode)) {
		query["StatusCode"] = request.StatusCode
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingActivityId)) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingActivities"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingActivitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingActivities is deprecated, please use Ess::2022-02-22::DescribeScalingActivities,Ess::2014-08-28::DescribeScalingActivities instead.
//
// @param request - DescribeScalingActivitiesRequest
//
// @return DescribeScalingActivitiesResponse
// Deprecated
func (client *Client) DescribeScalingActivities(request *DescribeScalingActivitiesRequest) (_result *DescribeScalingActivitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingActivitiesResponse{}
	_body, _err := client.DescribeScalingActivitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingActivityDetail is deprecated, please use Ess::2022-02-22::DescribeScalingActivityDetail,Ess::2014-08-28::DescribeScalingActivityDetail instead.
//
// @param request - DescribeScalingActivityDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingActivityDetailResponse
// Deprecated
func (client *Client) DescribeScalingActivityDetailWithOptions(request *DescribeScalingActivityDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingActivityDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingActivityId)) {
		query["ScalingActivityId"] = request.ScalingActivityId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingActivityDetail"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingActivityDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingActivityDetail is deprecated, please use Ess::2022-02-22::DescribeScalingActivityDetail,Ess::2014-08-28::DescribeScalingActivityDetail instead.
//
// @param request - DescribeScalingActivityDetailRequest
//
// @return DescribeScalingActivityDetailResponse
// Deprecated
func (client *Client) DescribeScalingActivityDetail(request *DescribeScalingActivityDetailRequest) (_result *DescribeScalingActivityDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingActivityDetailResponse{}
	_body, _err := client.DescribeScalingActivityDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingConfigurations is deprecated, please use Ess::2022-02-22::DescribeScalingConfigurations,Ess::2014-08-28::DescribeScalingConfigurations instead.
//
// @param request - DescribeScalingConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingConfigurationsResponse
// Deprecated
func (client *Client) DescribeScalingConfigurationsWithOptions(request *DescribeScalingConfigurationsRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingConfigurationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationId)) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationName)) {
		query["ScalingConfigurationName"] = request.ScalingConfigurationName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingConfigurations"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingConfigurations is deprecated, please use Ess::2022-02-22::DescribeScalingConfigurations,Ess::2014-08-28::DescribeScalingConfigurations instead.
//
// @param request - DescribeScalingConfigurationsRequest
//
// @return DescribeScalingConfigurationsResponse
// Deprecated
func (client *Client) DescribeScalingConfigurations(request *DescribeScalingConfigurationsRequest) (_result *DescribeScalingConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingConfigurationsResponse{}
	_body, _err := client.DescribeScalingConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingGroups is deprecated, please use Ess::2022-02-22::DescribeScalingGroups,Ess::2014-08-28::DescribeScalingGroups instead.
//
// @param request - DescribeScalingGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingGroupsResponse
// Deprecated
func (client *Client) DescribeScalingGroupsWithOptions(request *DescribeScalingGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupName)) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingGroups"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingGroups is deprecated, please use Ess::2022-02-22::DescribeScalingGroups,Ess::2014-08-28::DescribeScalingGroups instead.
//
// @param request - DescribeScalingGroupsRequest
//
// @return DescribeScalingGroupsResponse
// Deprecated
func (client *Client) DescribeScalingGroups(request *DescribeScalingGroupsRequest) (_result *DescribeScalingGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingGroupsResponse{}
	_body, _err := client.DescribeScalingGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingInstances is deprecated, please use Ess::2022-02-22::DescribeScalingInstances,Ess::2014-08-28::DescribeScalingInstances instead.
//
// @param request - DescribeScalingInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingInstancesResponse
// Deprecated
func (client *Client) DescribeScalingInstancesWithOptions(request *DescribeScalingInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreationType)) {
		query["CreationType"] = request.CreationType
	}

	if !tea.BoolValue(util.IsUnset(request.HealthStatus)) {
		query["HealthStatus"] = request.HealthStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LifecycleState)) {
		query["LifecycleState"] = request.LifecycleState
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingConfigurationId)) {
		query["ScalingConfigurationId"] = request.ScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingInstances"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingInstances is deprecated, please use Ess::2022-02-22::DescribeScalingInstances,Ess::2014-08-28::DescribeScalingInstances instead.
//
// @param request - DescribeScalingInstancesRequest
//
// @return DescribeScalingInstancesResponse
// Deprecated
func (client *Client) DescribeScalingInstances(request *DescribeScalingInstancesRequest) (_result *DescribeScalingInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingInstancesResponse{}
	_body, _err := client.DescribeScalingInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingRules is deprecated, please use Ess::2022-02-22::DescribeScalingRules,Ess::2014-08-28::DescribeScalingRules instead.
//
// @param request - DescribeScalingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScalingRulesResponse
// Deprecated
func (client *Client) DescribeScalingRulesWithOptions(request *DescribeScalingRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeScalingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleAri)) {
		query["ScalingRuleAri"] = request.ScalingRuleAri
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleId)) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScalingRules"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScalingRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeScalingRules is deprecated, please use Ess::2022-02-22::DescribeScalingRules,Ess::2014-08-28::DescribeScalingRules instead.
//
// @param request - DescribeScalingRulesRequest
//
// @return DescribeScalingRulesResponse
// Deprecated
func (client *Client) DescribeScalingRules(request *DescribeScalingRulesRequest) (_result *DescribeScalingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScalingRulesResponse{}
	_body, _err := client.DescribeScalingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DescribeScheduledTasks is deprecated, please use Ess::2022-02-22::DescribeScheduledTasks,Ess::2014-08-28::DescribeScheduledTasks instead.
//
// @param request - DescribeScheduledTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScheduledTasksResponse
// Deprecated
func (client *Client) DescribeScheduledTasksWithOptions(request *DescribeScheduledTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeScheduledTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledAction)) {
		query["ScheduledAction"] = request.ScheduledAction
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskId)) {
		query["ScheduledTaskId"] = request.ScheduledTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskName)) {
		query["ScheduledTaskName"] = request.ScheduledTaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScheduledTasks"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScheduledTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DescribeScheduledTasks is deprecated, please use Ess::2022-02-22::DescribeScheduledTasks,Ess::2014-08-28::DescribeScheduledTasks instead.
//
// @param request - DescribeScheduledTasksRequest
//
// @return DescribeScheduledTasksResponse
// Deprecated
func (client *Client) DescribeScheduledTasks(request *DescribeScheduledTasksRequest) (_result *DescribeScheduledTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeScheduledTasksResponse{}
	_body, _err := client.DescribeScheduledTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DetachInstances is deprecated, please use Ess::2022-02-22::DetachInstances,Ess::2014-08-28::DetachInstances instead.
//
// @param request - DetachInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachInstancesResponse
// Deprecated
func (client *Client) DetachInstancesWithOptions(request *DetachInstancesRequest, runtime *util.RuntimeOptions) (_result *DetachInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachInstances"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DetachInstances is deprecated, please use Ess::2022-02-22::DetachInstances,Ess::2014-08-28::DetachInstances instead.
//
// @param request - DetachInstancesRequest
//
// @return DetachInstancesResponse
// Deprecated
func (client *Client) DetachInstances(request *DetachInstancesRequest) (_result *DetachInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachInstancesResponse{}
	_body, _err := client.DetachInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DisableScalingGroup is deprecated, please use Ess::2022-02-22::DisableScalingGroup,Ess::2014-08-28::DisableScalingGroup instead.
//
// @param request - DisableScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableScalingGroupResponse
// Deprecated
func (client *Client) DisableScalingGroupWithOptions(request *DisableScalingGroupRequest, runtime *util.RuntimeOptions) (_result *DisableScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableScalingGroup"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DisableScalingGroup is deprecated, please use Ess::2022-02-22::DisableScalingGroup,Ess::2014-08-28::DisableScalingGroup instead.
//
// @param request - DisableScalingGroupRequest
//
// @return DisableScalingGroupResponse
// Deprecated
func (client *Client) DisableScalingGroup(request *DisableScalingGroupRequest) (_result *DisableScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableScalingGroupResponse{}
	_body, _err := client.DisableScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI EnableScalingGroup is deprecated, please use Ess::2014-08-28::EnableScalingGroup,Ess::2022-02-22::EnableScalingGroup instead.
//
// @param request - EnableScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableScalingGroupResponse
// Deprecated
func (client *Client) EnableScalingGroupWithOptions(request *EnableScalingGroupRequest, runtime *util.RuntimeOptions) (_result *EnableScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveScalingConfigurationId)) {
		query["ActiveScalingConfigurationId"] = request.ActiveScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableScalingGroup"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI EnableScalingGroup is deprecated, please use Ess::2014-08-28::EnableScalingGroup,Ess::2022-02-22::EnableScalingGroup instead.
//
// @param request - EnableScalingGroupRequest
//
// @return EnableScalingGroupResponse
// Deprecated
func (client *Client) EnableScalingGroup(request *EnableScalingGroupRequest) (_result *EnableScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableScalingGroupResponse{}
	_body, _err := client.EnableScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ExecuteScalingRule is deprecated, please use Ess::2014-08-28::ExecuteScalingRule,Ess::2022-02-22::ExecuteScalingRule instead.
//
// @param request - ExecuteScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteScalingRuleResponse
// Deprecated
func (client *Client) ExecuteScalingRuleWithOptions(request *ExecuteScalingRuleRequest, runtime *util.RuntimeOptions) (_result *ExecuteScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleAri)) {
		query["ScalingRuleAri"] = request.ScalingRuleAri
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteScalingRule"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ExecuteScalingRule is deprecated, please use Ess::2014-08-28::ExecuteScalingRule,Ess::2022-02-22::ExecuteScalingRule instead.
//
// @param request - ExecuteScalingRuleRequest
//
// @return ExecuteScalingRuleResponse
// Deprecated
func (client *Client) ExecuteScalingRule(request *ExecuteScalingRuleRequest) (_result *ExecuteScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteScalingRuleResponse{}
	_body, _err := client.ExecuteScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyScalingGroup is deprecated, please use Ess::2014-08-28::ModifyScalingGroup,Ess::2022-02-22::ModifyScalingGroup instead.
//
// @param request - ModifyScalingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScalingGroupResponse
// Deprecated
func (client *Client) ModifyScalingGroupWithOptions(request *ModifyScalingGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyScalingGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActiveScalingConfigurationId)) {
		query["ActiveScalingConfigurationId"] = request.ActiveScalingConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultCooldown)) {
		query["DefaultCooldown"] = request.DefaultCooldown
	}

	if !tea.BoolValue(util.IsUnset(request.MaxSize)) {
		query["MaxSize"] = request.MaxSize
	}

	if !tea.BoolValue(util.IsUnset(request.MinSize)) {
		query["MinSize"] = request.MinSize
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupName)) {
		query["ScalingGroupName"] = request.ScalingGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.RemovalPolicy)) {
		query["RemovalPolicy"] = request.RemovalPolicy
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScalingGroup"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScalingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyScalingGroup is deprecated, please use Ess::2014-08-28::ModifyScalingGroup,Ess::2022-02-22::ModifyScalingGroup instead.
//
// @param request - ModifyScalingGroupRequest
//
// @return ModifyScalingGroupResponse
// Deprecated
func (client *Client) ModifyScalingGroup(request *ModifyScalingGroupRequest) (_result *ModifyScalingGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScalingGroupResponse{}
	_body, _err := client.ModifyScalingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyScalingRule is deprecated, please use Ess::2014-08-28::ModifyScalingRule,Ess::2022-02-22::ModifyScalingRule instead.
//
// @param request - ModifyScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScalingRuleResponse
// Deprecated
func (client *Client) ModifyScalingRuleWithOptions(request *ModifyScalingRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdjustmentType)) {
		query["AdjustmentType"] = request.AdjustmentType
	}

	if !tea.BoolValue(util.IsUnset(request.AdjustmentValue)) {
		query["AdjustmentValue"] = request.AdjustmentValue
	}

	if !tea.BoolValue(util.IsUnset(request.Cooldown)) {
		query["Cooldown"] = request.Cooldown
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleId)) {
		query["ScalingRuleId"] = request.ScalingRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingRuleName)) {
		query["ScalingRuleName"] = request.ScalingRuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScalingRule"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyScalingRule is deprecated, please use Ess::2014-08-28::ModifyScalingRule,Ess::2022-02-22::ModifyScalingRule instead.
//
// @param request - ModifyScalingRuleRequest
//
// @return ModifyScalingRuleResponse
// Deprecated
func (client *Client) ModifyScalingRule(request *ModifyScalingRuleRequest) (_result *ModifyScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScalingRuleResponse{}
	_body, _err := client.ModifyScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ModifyScheduledTask is deprecated, please use Ess::2014-08-28::ModifyScheduledTask,Ess::2022-02-22::ModifyScheduledTask instead.
//
// @param request - ModifyScheduledTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduledTaskResponse
// Deprecated
func (client *Client) ModifyScheduledTaskWithOptions(request *ModifyScheduledTaskRequest, runtime *util.RuntimeOptions) (_result *ModifyScheduledTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchExpirationTime)) {
		query["LaunchExpirationTime"] = request.LaunchExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.LaunchTime)) {
		query["LaunchTime"] = request.LaunchTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceEndTime)) {
		query["RecurrenceEndTime"] = request.RecurrenceEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceType)) {
		query["RecurrenceType"] = request.RecurrenceType
	}

	if !tea.BoolValue(util.IsUnset(request.RecurrenceValue)) {
		query["RecurrenceValue"] = request.RecurrenceValue
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledAction)) {
		query["ScheduledAction"] = request.ScheduledAction
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskId)) {
		query["ScheduledTaskId"] = request.ScheduledTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduledTaskName)) {
		query["ScheduledTaskName"] = request.ScheduledTaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskEnabled)) {
		query["TaskEnabled"] = request.TaskEnabled
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScheduledTask"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ModifyScheduledTask is deprecated, please use Ess::2014-08-28::ModifyScheduledTask,Ess::2022-02-22::ModifyScheduledTask instead.
//
// @param request - ModifyScheduledTaskRequest
//
// @return ModifyScheduledTaskResponse
// Deprecated
func (client *Client) ModifyScheduledTask(request *ModifyScheduledTaskRequest) (_result *ModifyScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.ModifyScheduledTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI RemoveInstances is deprecated, please use Ess::2014-08-28::RemoveInstances,Ess::2022-02-22::RemoveInstances instead.
//
// @param request - RemoveInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInstancesResponse
// Deprecated
func (client *Client) RemoveInstancesWithOptions(request *RemoveInstancesRequest, runtime *util.RuntimeOptions) (_result *RemoveInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScalingGroupId)) {
		query["ScalingGroupId"] = request.ScalingGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveInstances"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RemoveInstances is deprecated, please use Ess::2014-08-28::RemoveInstances,Ess::2022-02-22::RemoveInstances instead.
//
// @param request - RemoveInstancesRequest
//
// @return RemoveInstancesResponse
// Deprecated
func (client *Client) RemoveInstances(request *RemoveInstancesRequest) (_result *RemoveInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveInstancesResponse{}
	_body, _err := client.RemoveInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI VerifyAuthentication is deprecated, please use Ess::2014-08-28::VerifyAuthentication instead.
//
// @param request - VerifyAuthenticationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAuthenticationResponse
// Deprecated
func (client *Client) VerifyAuthenticationWithOptions(request *VerifyAuthenticationRequest, runtime *util.RuntimeOptions) (_result *VerifyAuthenticationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyAuthentication"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyAuthenticationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI VerifyAuthentication is deprecated, please use Ess::2014-08-28::VerifyAuthentication instead.
//
// @param request - VerifyAuthenticationRequest
//
// @return VerifyAuthenticationResponse
// Deprecated
func (client *Client) VerifyAuthentication(request *VerifyAuthenticationRequest) (_result *VerifyAuthenticationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyAuthenticationResponse{}
	_body, _err := client.VerifyAuthenticationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI VerifyUser is deprecated, please use Ess::2014-08-28::VerifyUser instead.
//
// @param request - VerifyUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyUserResponse
// Deprecated
func (client *Client) VerifyUserWithOptions(request *VerifyUserRequest, runtime *util.RuntimeOptions) (_result *VerifyUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyUser"),
		Version:     tea.String("2016-07-22"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &VerifyUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI VerifyUser is deprecated, please use Ess::2014-08-28::VerifyUser instead.
//
// @param request - VerifyUserRequest
//
// @return VerifyUserResponse
// Deprecated
func (client *Client) VerifyUser(request *VerifyUserRequest) (_result *VerifyUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyUserResponse{}
	_body, _err := client.VerifyUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
