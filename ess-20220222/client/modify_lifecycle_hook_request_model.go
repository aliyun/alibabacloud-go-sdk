// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLifecycleHookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultResult(v string) *ModifyLifecycleHookRequest
	GetDefaultResult() *string
	SetHeartbeatTimeout(v int32) *ModifyLifecycleHookRequest
	GetHeartbeatTimeout() *int32
	SetLifecycleActionSources(v []*string) *ModifyLifecycleHookRequest
	GetLifecycleActionSources() []*string
	SetLifecycleHookId(v string) *ModifyLifecycleHookRequest
	GetLifecycleHookId() *string
	SetLifecycleHookName(v string) *ModifyLifecycleHookRequest
	GetLifecycleHookName() *string
	SetLifecycleHookStatus(v string) *ModifyLifecycleHookRequest
	GetLifecycleHookStatus() *string
	SetLifecycleTransition(v string) *ModifyLifecycleHookRequest
	GetLifecycleTransition() *string
	SetNotificationArn(v string) *ModifyLifecycleHookRequest
	GetNotificationArn() *string
	SetNotificationMetadata(v string) *ModifyLifecycleHookRequest
	GetNotificationMetadata() *string
	SetOwnerAccount(v string) *ModifyLifecycleHookRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyLifecycleHookRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyLifecycleHookRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyLifecycleHookRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *ModifyLifecycleHookRequest
	GetScalingGroupId() *string
}

type ModifyLifecycleHookRequest struct {
	// The subsequent action to perform after the wait state ends. Valid values:
	//
	// - CONTINUE: Continues the response to the scale-out or scale-down activity.
	//
	// - ABANDON: Releases the ECS instances created during the scale-out activity or removes the ECS instances from the scaling group during the scale-down activity.
	//
	// - ROLLBACK: Rejects the release of ECS instances and rolls back the scale-down activity. For a scale-out activity, releases the created ECS instances, which has the same effect as ABANDON.
	//
	// If a scale-down activity is triggered and multiple lifecycle hooks are invoked, and the DefaultResult of one lifecycle hook is set to ABANDON or ROLLBACK, the wait states of other lifecycle hooks are ended in advance when the wait state of this hook ends. In other cases, the subsequent action is determined by the last lifecycle hook that ends its wait state.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The wait time set by the lifecycle hook for the scaling activity. The subsequent action is performed after the wait state times out. Valid values: 30 to 21600. Unit: seconds.
	//
	// You can call the RecordLifecycleActionHeartbeat operation to extend the wait time of the ECS instance, or call the CompleteLifecycleAction operation to end the wait state of the scaling activity in advance.
	//
	// example:
	//
	// 600
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The specific scenarios that trigger the lifecycle hook.
	LifecycleActionSources []*string `json:"LifecycleActionSources,omitempty" xml:"LifecycleActionSources,omitempty" type:"Repeated"`
	// The ID of the lifecycle hook to be modified.
	//
	// example:
	//
	// ash-bp1fxuqyi98w0aib****
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	// The name of the lifecycle hook to be modified.
	//
	// example:
	//
	// test_SCAL****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	// The status of the lifecycle hook. Valid values:
	//
	// - Active: Enables the lifecycle hook.
	//
	// - InActive: Disables the lifecycle hook.
	//
	// If you do not specify this parameter, the status of the lifecycle hook remains unchanged.
	//
	// > After a lifecycle hook is created, it is enabled by default.
	//
	// example:
	//
	// Active
	LifecycleHookStatus *string `json:"LifecycleHookStatus,omitempty" xml:"LifecycleHookStatus,omitempty"`
	// The type of scaling activity to which the lifecycle hook applies. Valid values:
	//
	// - SCALE_OUT: Scale-out activity of the scaling group.
	//
	// - SCALE_IN: Scale-in activity of the scaling group.
	//
	// example:
	//
	// SCALE_IN
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The identifier of the notification recipient for the lifecycle hook. The following notification methods are supported:
	//
	// - Message Service (MNS) queue. Format: acs:mns:{region-id}:{account-id}:queue/{queuename}.
	//
	// - MNS topic. Format: acs:mns:{region-id}:{account-id}:topic/{topicname}.
	//
	// - Operation Orchestration Service (OOS) template. Format: acs:oos:{region-id}:{account-id}:template/{templatename}.
	//
	// - EventBridge. Format: acs:eventbridge:{region-id}:{account-id}:eventbus/default.
	//
	// Variables in the formats:
	//
	// - region-id: The ID of the region where the scaling group resides.
	//
	// - account-id: The Alibaba Cloud account ID.
	//
	// - queuename: The name of the MNS queue.
	//
	// - topicname: The name of the MNS topic.
	//
	// - templatename: The name of the OOS template.
	//
	// example:
	//
	// acs:mns:cn-beijing:161456884340****:queue/modifyLifecycleHo****
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The fixed string information for the wait state of the scaling activity. The value can be up to 4,096 characters in length.
	//
	// When Auto Scaling pushes a message to the notification recipient, it also sends the predefined NotificationMetadata value to help you manage and tag different types of notifications. This parameter takes effect only when the NotificationArn parameter is specified.
	//
	// example:
	//
	// Test
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group to which the lifecycle hook belongs.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s ModifyLifecycleHookRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *ModifyLifecycleHookRequest) GetDefaultResult() *string {
	return s.DefaultResult
}

func (s *ModifyLifecycleHookRequest) GetHeartbeatTimeout() *int32 {
	return s.HeartbeatTimeout
}

func (s *ModifyLifecycleHookRequest) GetLifecycleActionSources() []*string {
	return s.LifecycleActionSources
}

func (s *ModifyLifecycleHookRequest) GetLifecycleHookId() *string {
	return s.LifecycleHookId
}

func (s *ModifyLifecycleHookRequest) GetLifecycleHookName() *string {
	return s.LifecycleHookName
}

func (s *ModifyLifecycleHookRequest) GetLifecycleHookStatus() *string {
	return s.LifecycleHookStatus
}

func (s *ModifyLifecycleHookRequest) GetLifecycleTransition() *string {
	return s.LifecycleTransition
}

func (s *ModifyLifecycleHookRequest) GetNotificationArn() *string {
	return s.NotificationArn
}

func (s *ModifyLifecycleHookRequest) GetNotificationMetadata() *string {
	return s.NotificationMetadata
}

func (s *ModifyLifecycleHookRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyLifecycleHookRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLifecycleHookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLifecycleHookRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyLifecycleHookRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ModifyLifecycleHookRequest) SetDefaultResult(v string) *ModifyLifecycleHookRequest {
	s.DefaultResult = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetHeartbeatTimeout(v int32) *ModifyLifecycleHookRequest {
	s.HeartbeatTimeout = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleActionSources(v []*string) *ModifyLifecycleHookRequest {
	s.LifecycleActionSources = v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleHookId(v string) *ModifyLifecycleHookRequest {
	s.LifecycleHookId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleHookName(v string) *ModifyLifecycleHookRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleHookStatus(v string) *ModifyLifecycleHookRequest {
	s.LifecycleHookStatus = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetLifecycleTransition(v string) *ModifyLifecycleHookRequest {
	s.LifecycleTransition = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetNotificationArn(v string) *ModifyLifecycleHookRequest {
	s.NotificationArn = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetNotificationMetadata(v string) *ModifyLifecycleHookRequest {
	s.NotificationMetadata = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetOwnerAccount(v string) *ModifyLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetOwnerId(v int64) *ModifyLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetRegionId(v string) *ModifyLifecycleHookRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetResourceOwnerAccount(v string) *ModifyLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLifecycleHookRequest) SetScalingGroupId(v string) *ModifyLifecycleHookRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyLifecycleHookRequest) Validate() error {
	return dara.Validate(s)
}
