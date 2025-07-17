// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNotificationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationArn(v string) *CreateNotificationConfigurationRequest
	GetNotificationArn() *string
	SetNotificationTypes(v []*string) *CreateNotificationConfigurationRequest
	GetNotificationTypes() []*string
	SetOwnerId(v int64) *CreateNotificationConfigurationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateNotificationConfigurationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateNotificationConfigurationRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *CreateNotificationConfigurationRequest
	GetScalingGroupId() *string
	SetTimeZone(v string) *CreateNotificationConfigurationRequest
	GetTimeZone() *string
}

type CreateNotificationConfigurationRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the notification recipient. The following list describes the value formats of this parameter:
	//
	// 	- If you specify CloudMonitor as the notification recipient, specify the value in the `acs:ess:{region-id}:{account-id}:cloudmonitor` format.
	//
	// 	- If you specify an SMQ queue as the notification recipient, specify the value in the `acs:mns:{region-id}:{account-id}:queue/{queuename}` format.
	//
	// 	- If you specify an SMQ topic as the notification recipient, specify the value in the `acs:mns:{region-id}:{account-id}:topic/{topicname}` format.
	//
	// The variables in the preceding formats have the following meanings:
	//
	// 	- `region-id`: the region ID of the scaling group.
	//
	// 	- `account-id`: the ID of the Alibaba Cloud account.
	//
	// 	- `queuename`: the name of the SMQ queue.
	//
	// 	- `topicname`: the name of the SMQ topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:mns:cn-beijing:161456884340****:queue/modifyLifecycleHo****
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The notification types. Specify multiple IDs in the repeated list form.
	//
	// You can call the DescribeNotificationTypes operation to query the values of this parameter.
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

func (s CreateNotificationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateNotificationConfigurationRequest) GetNotificationArn() *string {
	return s.NotificationArn
}

func (s *CreateNotificationConfigurationRequest) GetNotificationTypes() []*string {
	return s.NotificationTypes
}

func (s *CreateNotificationConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNotificationConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNotificationConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNotificationConfigurationRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateNotificationConfigurationRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateNotificationConfigurationRequest) SetNotificationArn(v string) *CreateNotificationConfigurationRequest {
	s.NotificationArn = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetNotificationTypes(v []*string) *CreateNotificationConfigurationRequest {
	s.NotificationTypes = v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetOwnerId(v int64) *CreateNotificationConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetRegionId(v string) *CreateNotificationConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetResourceOwnerAccount(v string) *CreateNotificationConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetScalingGroupId(v string) *CreateNotificationConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) SetTimeZone(v string) *CreateNotificationConfigurationRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateNotificationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
