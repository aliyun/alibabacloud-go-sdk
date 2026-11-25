// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecycleHookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultResult(v string) *CreateLifecycleHookRequest
	GetDefaultResult() *string
	SetHeartbeatTimeout(v int32) *CreateLifecycleHookRequest
	GetHeartbeatTimeout() *int32
	SetLifecycleActionSources(v []*string) *CreateLifecycleHookRequest
	GetLifecycleActionSources() []*string
	SetLifecycleHookName(v string) *CreateLifecycleHookRequest
	GetLifecycleHookName() *string
	SetLifecycleTransition(v string) *CreateLifecycleHookRequest
	GetLifecycleTransition() *string
	SetNotificationArn(v string) *CreateLifecycleHookRequest
	GetNotificationArn() *string
	SetNotificationMetadata(v string) *CreateLifecycleHookRequest
	GetNotificationMetadata() *string
	SetOwnerAccount(v string) *CreateLifecycleHookRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLifecycleHookRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateLifecycleHookRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *CreateLifecycleHookRequest
	GetScalingGroupId() *string
}

type CreateLifecycleHookRequest struct {
	// The action to be performed after the wait state ends. Valid values:
	//
	// - CONTINUE: Continues the response to the scale-out or scale-in activity.
	//
	// - ABANDON: Releases the ECS instances created by the scale-out activity or removes the ECS instances in the scale-in activity from the scaling group.
	//
	// - ROLLBACK: Rejects the release of ECS instances and rolls back the scale-in activity. For a scale-out activity, this value directly releases the created ECS instances, which has the same effect as ABANDON.
	//
	// When a scale-in activity occurs in the scaling group and multiple lifecycle hooks are triggered, if the DefaultResult of one lifecycle hook is set to ABANDON or ROLLBACK and the wait state of the hook ends, the wait states of other lifecycle hooks are ended in advance. In other cases, the action to be performed is subject to the action of the last lifecycle hook whose wait state ends.
	//
	// Default value: CONTINUE.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The wait time configured by the lifecycle hook for the scaling activity. After the wait state times out, the next action is performed. Valid values: 30 to 21600. Unit: seconds.
	//
	// After you create a lifecycle hook, you can call RecordLifecycleActionHeartbeat to extend the wait time of the ECS instance, or call CompleteLifecycleAction to end the wait state of the scaling activity in advance.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The specific scenarios triggered by the lifecycle hook.
	LifecycleActionSources []*string `json:"LifecycleActionSources,omitempty" xml:"LifecycleActionSources,omitempty" type:"Repeated"`
	// The name of the lifecycle hook. The name must be unique among all lifecycle hooks in the current scaling group. It must be 2 to 64 characters in length, start with a digit, an uppercase or lowercase letter, or a Chinese character, and can contain digits, underscores (_), hyphens (-), and periods (.).
	//
	// Default value: the value of LifecycleHookId.
	//
	// example:
	//
	// lifecyclehook****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	// The type of scaling activity to which the lifecycle hook applies. Valid values:
	//
	//
	//
	// - SCALE_OUT: scale-out activity of the scaling group.
	//
	// - SCALE_IN: scale-in activity of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// SCALE_OUT
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The identifier of the notification recipient for the lifecycle hook. If you do not specify this parameter, no notifications are sent. If you specify this parameter, the following notification methods are supported:
	//
	// - MNS queue. Format: acs:mns:{region-id}:{account-id}:queue/{queuename}.
	//
	// - MNS topic. Format: acs:mns:{region-id}:{account-id}:topic/{topicname}.
	//
	// - OOS template. Format: acs:oos:{region-id}:{account-id}:template/{templatename}.
	//
	// - EventBridge. Format: acs:eventbridge:{region-id}:{account-id}:eventbus/default.
	//
	// The variables in the formats have the following meanings:
	//
	// - region-id: the ID of the region where the scaling group resides.
	//
	// - account-id: the Alibaba Cloud account ID. The ID of a Resource Access Management (RAM) user is not supported.
	//
	// - queuename: the name of the MNS queue.
	//
	// - topicname: the name of the MNS topic.
	//
	// - templatename: the name of the OOS template.
	//
	// example:
	//
	// acs:mns:cn-beijing:161456884340****:queue/modifyLifecycleHo****
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The notification message pushed to the NotificationArn (notification recipient) when the lifecycle hook pauses the scaling activity. This parameter helps you manage and mark different types of notification messages. You must specify NotificationArn together with this parameter. The value can be up to 4,096 characters in length.
	//
	// If NotificationArn is set to an OOS template, including a public or custom template, NotificationMetadata must be specified as a JSON string that matches the parameters of the OOS template. For example, `{"dbInstanceId": "dds-bp17661e0135****", "modifyMode": "Append"}`, where `dbInstanceId` and `modifyMode` are parameters defined in the OOS template. Some parameters of the OOS template have default values. You must specify parameters without default values in NotificationMetadata. If you specify parameters with default values, the original default values are overwritten. However, keep the default values for the following parameters to obtain relevant information at runtime:
	//
	// - `regionId`: the region where the scaling activity is executed. Default value: ${regionId}.
	//
	// - `instanceIds`: the list of instances associated with the scaling activity. Default value: ${instanceIds}.
	//
	// - `lifecycleHookId`: the ID of the lifecycle hook. Default value: ${lifecycleHookId}.
	//
	// - `lifecycleActionToken`: the wait state identifier for the scaling activity, which is used to end the current suspended activity in advance. Default value: ${lifecycleActionToken}.
	//
	// - `scalingGroupId`: the ID of the scaling group to which the scaling activity belongs. Default value: ${scalingGroupId}.
	//
	// - `lifecycleActionResult`: the action to be performed after the wait state of the lifecycle hook ends. This parameter is used to end the current suspended activity in advance when the OOS template execution fails. When DefaultResult is set to ROLLBACK, the default value of this parameter is ROLLBACK. In other cases, the default value is ABANDON.
	//
	// >- You can configure the value of lifecycleActionResult to overwrite the current default value as needed. Valid values: ABANDON, CONTINUE, ROLLBACK, and ${lifecycleActionResult}.
	//
	// Among these values, ${lifecycleActionResult} indicates that the value of lifecycleActionResult is the same as that of DefaultResult.
	//
	// - You can obtain detailed information about the parameters of the corresponding template in the [OOS console](https://oos.console.aliyun.com/).
	//
	// example:
	//
	// Test lifecycle hook.
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1eyv4qn8ssgv43****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CreateLifecycleHookRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecycleHookRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecycleHookRequest) GetDefaultResult() *string {
	return s.DefaultResult
}

func (s *CreateLifecycleHookRequest) GetHeartbeatTimeout() *int32 {
	return s.HeartbeatTimeout
}

func (s *CreateLifecycleHookRequest) GetLifecycleActionSources() []*string {
	return s.LifecycleActionSources
}

func (s *CreateLifecycleHookRequest) GetLifecycleHookName() *string {
	return s.LifecycleHookName
}

func (s *CreateLifecycleHookRequest) GetLifecycleTransition() *string {
	return s.LifecycleTransition
}

func (s *CreateLifecycleHookRequest) GetNotificationArn() *string {
	return s.NotificationArn
}

func (s *CreateLifecycleHookRequest) GetNotificationMetadata() *string {
	return s.NotificationMetadata
}

func (s *CreateLifecycleHookRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLifecycleHookRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLifecycleHookRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLifecycleHookRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateLifecycleHookRequest) SetDefaultResult(v string) *CreateLifecycleHookRequest {
	s.DefaultResult = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetHeartbeatTimeout(v int32) *CreateLifecycleHookRequest {
	s.HeartbeatTimeout = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetLifecycleActionSources(v []*string) *CreateLifecycleHookRequest {
	s.LifecycleActionSources = v
	return s
}

func (s *CreateLifecycleHookRequest) SetLifecycleHookName(v string) *CreateLifecycleHookRequest {
	s.LifecycleHookName = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetLifecycleTransition(v string) *CreateLifecycleHookRequest {
	s.LifecycleTransition = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetNotificationArn(v string) *CreateLifecycleHookRequest {
	s.NotificationArn = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetNotificationMetadata(v string) *CreateLifecycleHookRequest {
	s.NotificationMetadata = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetOwnerAccount(v string) *CreateLifecycleHookRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetOwnerId(v int64) *CreateLifecycleHookRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetResourceOwnerAccount(v string) *CreateLifecycleHookRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLifecycleHookRequest) SetScalingGroupId(v string) *CreateLifecycleHookRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateLifecycleHookRequest) Validate() error {
	return dara.Validate(s)
}
