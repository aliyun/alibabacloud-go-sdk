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
	// The action that you want Auto Scaling to perform after the lifecycle hook ends. Valid values:
	//
	// 	- CONTINUE: Auto Scaling continues to respond to scaling requests.
	//
	// 	- ABANDON: Auto Scaling releases Elastic Compute Service (ECS) instances that are created during scale-out activities, or removes ECS instances from the scaling group during scale-in activities.
	//
	// If multiple lifecycle hooks in a scaling group are triggered during scale-in activities and you set the DefaultResult parameter to ABANDON for the lifecycle hook that you want to modify, Auto Scaling immediately performs the action after the lifecycle hook that you want to modify ends. As a result, other lifecycle hooks end ahead of schedule. In other cases, Auto Scaling performs the action only after all lifecycle hooks end.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The period of time before the lifecycle hook ends. Auto Scaling performs the specified action after the lifecycle hook ends. Valid values: 30 to 21600. Unit: seconds.
	//
	// You can call the RecordLifecycleActionHeartbeat operation to prolong the length of a lifecycle hook. You can also call the CompleteLifecycleAction operation to end a lifecycle hook ahead of schedule.
	//
	// example:
	//
	// 600
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The ID of the lifecycle hook that you want to modify.
	//
	// example:
	//
	// ash-bp1fxuqyi98w0aib****
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	// The name of the lifecycle hook that you want to modify.
	//
	// example:
	//
	// test_SCALE_IN
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	// The status into which you want to put the lifecycle hook. Valid values:
	//
	// 	- Active
	//
	// 	- InActive
	//
	// If you do not specify this parameter, the status of the lifecycle hook remains unchanged after you call this operation.
	//
	// > By default, a lifecycle hook is in the Active state after you create it.
	//
	// example:
	//
	// Active
	LifecycleHookStatus *string `json:"LifecycleHookStatus,omitempty" xml:"LifecycleHookStatus,omitempty"`
	// The type of scaling activity to which the lifecycle hook applies. Valid values:
	//
	// 	- SCALE_OUT
	//
	// 	- SCALE_IN
	//
	// example:
	//
	// SCALE_IN
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the notification recipient. Specify the value in one of the following formats:
	//
	// 	- If you specify a Simple Message Queue (SMQ, formerly MNS) as the notification recipient, specify the value in the acs:mns:{region-id}:{account-id}:queue/{queuename} format.
	//
	// 	- If you specify an SMQ topic as the notification recipient, specify the value in the acs:mns:{region-id}:{account-id}:topic/{topicname} format.
	//
	// 	- If you specify a CloudOps Orchestration Service (OOS) template as the notification recipient, specify the value in the acs:oos:{region-id}:{account-id}:template/{templatename} format.
	//
	// 	- If you specify an event bus as the notification recipient, specify the value in the acs:eventbridge:{region-id}:{account-id}:eventbus/default format.
	//
	// The variables in the preceding value formats have the following meanings:
	//
	// 	- region-id: the region ID of your scaling group.
	//
	// 	- account-id: the ID of your Alibaba Cloud account.
	//
	// 	- queuename: the name of the SMQ queue.
	//
	// 	- topicname: the name of the SMQ topic.
	//
	// 	- templatename: the name of the OOS template.
	//
	// example:
	//
	// acs:mns:cn-beijing:161456884340****:queue/modifyLifecycleHo****
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The fixed string that is included in a notification. Auto Scaling sends the notification when the lifecycle hook takes effect. The value of this parameter cannot exceed 4,096 characters in length.
	//
	// Auto Scaling sends the value specified for the NotificationMetadata parameter together with the notification. This helps you categorize your notifications. The NotificationMetadata parameter takes effect only after you specify the NotificationArn parameter.
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
