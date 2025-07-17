// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecycleHooksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLifecycleHooks(v []*DescribeLifecycleHooksResponseBodyLifecycleHooks) *DescribeLifecycleHooksResponseBody
	GetLifecycleHooks() []*DescribeLifecycleHooksResponseBodyLifecycleHooks
	SetPageNumber(v int32) *DescribeLifecycleHooksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLifecycleHooksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLifecycleHooksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLifecycleHooksResponseBody
	GetTotalCount() *int32
}

type DescribeLifecycleHooksResponseBody struct {
	// Details about the lifecycle hooks.
	LifecycleHooks []*DescribeLifecycleHooksResponseBodyLifecycleHooks `json:"LifecycleHooks,omitempty" xml:"LifecycleHooks,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of lifecycle hooks.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLifecycleHooksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecycleHooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponseBody) GetLifecycleHooks() []*DescribeLifecycleHooksResponseBodyLifecycleHooks {
	return s.LifecycleHooks
}

func (s *DescribeLifecycleHooksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLifecycleHooksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLifecycleHooksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLifecycleHooksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLifecycleHooksResponseBody) SetLifecycleHooks(v []*DescribeLifecycleHooksResponseBodyLifecycleHooks) *DescribeLifecycleHooksResponseBody {
	s.LifecycleHooks = v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetPageNumber(v int32) *DescribeLifecycleHooksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetPageSize(v int32) *DescribeLifecycleHooksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetRequestId(v string) *DescribeLifecycleHooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) SetTotalCount(v int32) *DescribeLifecycleHooksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLifecycleHooksResponseBodyLifecycleHooks struct {
	// The next action that is performed after the lifecycle hook times out.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The period of time before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action that is specified by DefaultResult.
	//
	// example:
	//
	// 60
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The ID of the lifecycle hook.
	//
	// example:
	//
	// ash-bp19d1032y9kij96****
	LifecycleHookId *string `json:"LifecycleHookId,omitempty" xml:"LifecycleHookId,omitempty"`
	// The name of the lifecycle hook.
	//
	// example:
	//
	// lifecyclehook****
	LifecycleHookName *string `json:"LifecycleHookName,omitempty" xml:"LifecycleHookName,omitempty"`
	// The status of the lifecycle hook. Valid values:
	//
	// 	- Active: The lifecycle hook is enabled.
	//
	// 	- InActive: The lifecycle hook is disabled.
	//
	// example:
	//
	// Active
	LifecycleHookStatus *string `json:"LifecycleHookStatus,omitempty" xml:"LifecycleHookStatus,omitempty"`
	// The type of the scaling activity to which the lifecycle hook applies.
	//
	// example:
	//
	// SCALE_OUT
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The ARN of the notification recipient when the lifecycle hook takes effect. The value of this parameter must be in one of the following formats:
	//
	// 	- If you do not create a notification rule, specify the value in the `acs:ess:{region-id}:{account-id}:null/null` format.
	//
	// 	- If you specify a Simple Message Queue (SMQ, formerly MNS) queue as the notification recipient, specify the value in the `acs:mns:{region-id}:{account-id}:queue/{queuename}` format.
	//
	// 	- If you specify an SMQ as the notification recipient, specify the value in the `acs:mns:{region-id}:{account-id}:topic/{topicname}` format.
	//
	// 	- If you specify a CloudOps Orchestration Service (OOS) template as the notification recipient, specify the value in the `acs:oos:{region-id}:{account-id}:template/{templatename}` format.
	//
	// 	- If you specify an event bus as the notification recipient, specify the value in the `acs:eventbridge:{region-id}:{account-id}:eventbus/default` format.
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
	// acs:ess:cn-beijing:161456884340****:null/null
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The fixed string that is included in a notification that Auto Scaling sends when the lifecycle hook takes effect.
	//
	// example:
	//
	// Test Lifecycle Hook.
	NotificationMetadata *string `json:"NotificationMetadata,omitempty" xml:"NotificationMetadata,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeLifecycleHooksResponseBodyLifecycleHooks) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecycleHooksResponseBodyLifecycleHooks) GoString() string {
	return s.String()
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetDefaultResult() *string {
	return s.DefaultResult
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetHeartbeatTimeout() *int32 {
	return s.HeartbeatTimeout
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetLifecycleHookId() *string {
	return s.LifecycleHookId
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetLifecycleHookName() *string {
	return s.LifecycleHookName
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetLifecycleHookStatus() *string {
	return s.LifecycleHookStatus
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetLifecycleTransition() *string {
	return s.LifecycleTransition
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetNotificationArn() *string {
	return s.NotificationArn
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetNotificationMetadata() *string {
	return s.NotificationMetadata
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetDefaultResult(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.DefaultResult = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetHeartbeatTimeout(v int32) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.HeartbeatTimeout = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleHookId(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleHookId = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleHookName(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleHookName = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleHookStatus(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleHookStatus = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleTransition(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleTransition = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetNotificationArn(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.NotificationArn = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetNotificationMetadata(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.NotificationMetadata = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetScalingGroupId(v string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) Validate() error {
	return dara.Validate(s)
}
