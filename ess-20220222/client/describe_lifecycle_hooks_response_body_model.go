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
	// The list of lifecycle hooks.
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
	if s.LifecycleHooks != nil {
		for _, item := range s.LifecycleHooks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLifecycleHooksResponseBodyLifecycleHooks struct {
	// The action to perform after the wait state of the instance ends.
	//
	// example:
	//
	// CONTINUE
	DefaultResult *string `json:"DefaultResult,omitempty" xml:"DefaultResult,omitempty"`
	// The wait time configured in the lifecycle hook for the scaling activity. The next action is performed after the wait time expires.
	//
	// example:
	//
	// 60
	HeartbeatTimeout *int32 `json:"HeartbeatTimeout,omitempty" xml:"HeartbeatTimeout,omitempty"`
	// The specific scenarios in which the lifecycle hook is triggered.
	LifecycleActionSources []*string `json:"LifecycleActionSources,omitempty" xml:"LifecycleActionSources,omitempty" type:"Repeated"`
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
	// - Active: The lifecycle hook is active.
	//
	// - InActive: The lifecycle hook is inactive.
	//
	// example:
	//
	// Active
	LifecycleHookStatus *string `json:"LifecycleHookStatus,omitempty" xml:"LifecycleHookStatus,omitempty"`
	// The scaling activity type associated with the lifecycle hook.
	//
	// example:
	//
	// SCALE_OUT
	LifecycleTransition *string `json:"LifecycleTransition,omitempty" xml:"LifecycleTransition,omitempty"`
	// The identifier of the notification recipient for the lifecycle hook. The parameter format is as follows:
	//
	// - Notifications not configured: `acs:ess:{region-id}:{account-id}:null/null`.
	//
	// - MNS queue: `acs:mns:{region-id}:{account-id}:queue/{queuename}`.
	//
	// - MNS topic: `acs:mns:{region-id}:{account-id}:topic/{topicname}`.
	//
	// - OOS template: `acs:oos:{region-id}:{account-id}:template/{templatename}`.
	//
	// - EventBridge: `acs:eventbridge:{region-id}:{account-id}:eventbus/default`.
	//
	// The variables in the parameter format have the following meanings:
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
	// acs:ess:cn-beijing:161456884340****:null/null
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The fixed string information for the wait state of the scaling activity.
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

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) GetLifecycleActionSources() []*string {
	return s.LifecycleActionSources
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

func (s *DescribeLifecycleHooksResponseBodyLifecycleHooks) SetLifecycleActionSources(v []*string) *DescribeLifecycleHooksResponseBodyLifecycleHooks {
	s.LifecycleActionSources = v
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
