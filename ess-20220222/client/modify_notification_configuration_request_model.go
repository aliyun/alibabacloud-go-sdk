// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNotificationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationArn(v string) *ModifyNotificationConfigurationRequest
	GetNotificationArn() *string
	SetNotificationTypes(v []*string) *ModifyNotificationConfigurationRequest
	GetNotificationTypes() []*string
	SetOwnerId(v int64) *ModifyNotificationConfigurationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyNotificationConfigurationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyNotificationConfigurationRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *ModifyNotificationConfigurationRequest
	GetScalingGroupId() *string
	SetTimeZone(v string) *ModifyNotificationConfigurationRequest
	GetTimeZone() *string
}

type ModifyNotificationConfigurationRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the notification recipient. The following list describes the value formats of this parameter:
	//
	// 	- If you specify CloudMonitor as the notification recipient, specify the value in the `acs:ess:{region-id}:{account-id}:cloudmonitor` format.
	//
	// 	- If you specify a Simple Message Queue (SMQ) queue as the notification recipient, specify the value in the `acs:mns:{region-id}:{account-id}:queue/{queuename}` format.
	//
	// 	- If you specify an SMQ topic as the notification recipient, specify the value in the `acs:mns:{region-id}:{account-id}:topic/{topicname}` format.
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
	// This parameter is required.
	//
	// example:
	//
	// acs:ess:cn-beijing:161456884340****:cloudmonitor
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The event types.
	//
	// This parameter is required.
	NotificationTypes []*string `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Repeated"`
	OwnerId           *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The time zone of the notification. Specify the value in UTC. For example, a value of UTC+8 specifies that the time is 8 hours ahead of Coordinated Universal Time, and a value of UTC-7 specifies that the time is 7 hours behind Coordinated Universal Time.
	//
	// example:
	//
	// UTC+8
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s ModifyNotificationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyNotificationConfigurationRequest) GetNotificationArn() *string {
	return s.NotificationArn
}

func (s *ModifyNotificationConfigurationRequest) GetNotificationTypes() []*string {
	return s.NotificationTypes
}

func (s *ModifyNotificationConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyNotificationConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyNotificationConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyNotificationConfigurationRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ModifyNotificationConfigurationRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ModifyNotificationConfigurationRequest) SetNotificationArn(v string) *ModifyNotificationConfigurationRequest {
	s.NotificationArn = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetNotificationTypes(v []*string) *ModifyNotificationConfigurationRequest {
	s.NotificationTypes = v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetOwnerId(v int64) *ModifyNotificationConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetRegionId(v string) *ModifyNotificationConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetResourceOwnerAccount(v string) *ModifyNotificationConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetScalingGroupId(v string) *ModifyNotificationConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) SetTimeZone(v string) *ModifyNotificationConfigurationRequest {
	s.TimeZone = &v
	return s
}

func (s *ModifyNotificationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
