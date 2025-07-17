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
	// The action that you want Auto Scaling to perform after the lifecycle hook times out. Valid values:
	//
	// 	- CONTINUE: Auto Scaling continues to respond to scale-in or scale-out requests.
	//
	// 	- ABANDON: Auto Scaling releases Elastic Compute Service (ECS) instances that are created during scale-out activities, or removes ECS instances from the scaling group during scale-in activities.
	//
	// If multiple lifecycle hooks in a scaling group are triggered during scale-in activities and you set the DefaultResult parameter to ABANDON for one of the lifecycle hooks, Auto Scaling immediately performs the action after the lifecycle hook whose DefaultResult is set to ABANDON times out. As a result, other lifecycle hooks time out ahead of schedule. In other cases, Auto Scaling performs the action only after all lifecycle hooks time out.
	//
	// Default value: CONTINUE.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The period of time before the lifecycle hook times out. After the lifecycle hook times out, Auto Scaling performs the default action. Valid values: 30 to 21600. Unit: seconds.
	//
	// After you create a lifecycle hook, you can call the RecordLifecycleActionHeartbeat operation to prolong the timeout period of the lifecycle hook. You can also call the CompleteLifecycleAction operation to end the timeout period of the lifecycle hook ahead of schedule.
	//
	// Default value: 600.
	//
	// example:
	//
	// 600
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The name of the lifecycle hook. Each lifecycle hook name must be unique within a scaling group. The name must be 2 to 64 characters in length, and can contain letters, digits, underscores (_), hyphens (-), and periods (.). It must start with a letter or a digit.
	//
	// If you do not specify this parameter, the value of the LifecycleHookId parameter is used.
	//
	// example:
	//
	// lifecyclehook****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	// The type of the scaling activity to which the lifecycle hook applies. Valid values:
	//
	// 	- SCALE_OUT
	//
	// 	- SCALE_IN
	//
	// This parameter is required.
	//
	// example:
	//
	// SCALE_OUT
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the notification recipient. If you do not specify this parameter, no notification is sent when the lifecycle hook takes effect. If you specify this parameter, the value must be in one of the following formats:
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
	// 	- account-id: the ID of the Alibaba Cloud account. IDs of Resource Access Management (RAM) users are not supported.
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
	// The notification metadata that is sent when the lifecycle hook takes effect. This helps you manage and categorize notifications in an efficient manner. If you specify this parameter, you must specify the NotificationArn parameter. The parameter value cannot exceed 4,096 characters in length.
	//
	// If you use the NotificationArn parameter to specify a public or customOOS template, the value of the NotificationMetadata parameter must be a JSON string that contains the OOS template parameters. For example, your OOS template includes the following parameters: `{"dbInstanceId": "dds-bp17661e0135****", "modifyMode": "Append"}`, `dbInstanceId`, and `modifyMode`. Some parameters defined in your OOS template have default values. When you specify the NotificationMetadata parameter, specify parameters that do not have default values. If you specify parameters that have default values, the default values are overwritten. However, the default values of the following parameters must be retained to obtain information about scaling activities that are in progress:
	//
	// 	- `regionId`: the region ID of the scaling activity that is in progress. Default value: ${regionId}.
	//
	// 	- `instanceIds`: the IDs of ECS instances that are scaled in in the scaling activity. Default value: ${instanceIds}.
	//
	// 	- `lifecycleHookId`: the ID of the lifecycle hook. Default value: ${lifecycleHookId}.
	//
	// 	- `lifecycleActionToken`: the token of the lifecycle action. You can use the token to end the timeout period of the lifecycle hook ahead of schedule. Default value: ${lifecycleActionToken}
	//
	// 	- `scalingGroupId`: the ID of the scaling group in which the scaling activity is executed. Default value: ${scalingGroupId}.
	//
	// > You can obtain template parameter information in the [OOS console](https://oos.console.aliyun.com/).
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
