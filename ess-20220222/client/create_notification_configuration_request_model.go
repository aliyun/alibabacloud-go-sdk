// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNotificationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageEncoding(v string) *CreateNotificationConfigurationRequest
	GetMessageEncoding() *string
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
	// The encoding method for the notification content. Valid values:
	//
	// - PlainText: no encoding. The content is transmitted in plaintext.
	//
	// - Base64: Base64 encoding.
	//
	// Default value: Base64.
	//
	// > The Settings for this parameter are supported only for MNS notification recipients.
	//
	// example:
	//
	// PlainText
	MessageEncoding *string `json:"MessageEncoding,omitempty" xml:"MessageEncoding,omitempty"`
	// The identifier of the notification recipient. The following notification methods are supported when you configure the Settings for this parameter:
	//
	// - Cloud monitoring. Format: `acs:ess:{region-id}:{account-id}:cloudmonitor`.
	//
	// - MNS queue. Format: `acs:mns:{region-id}:{account-id}:queue/{queuename}`.
	//
	// - MNS topic. Format: `acs:mns:{region-id}:{account-id}:topic/{topicname}`.
	//
	// Variables in the formats:
	//
	// - `region-id`: the ID of the region where the scaling group resides.
	//
	// - `account-id`: the Alibaba Cloud account ID. You can invoke the STS GetCallerIdentity operation to obtain the AccountId response parameter, or view it by clicking your profile picture in the upper-right corner of the Alibaba Cloud Management Console.
	//
	// - `queuename`: the name of the MNS queue.
	//
	// - `topicname`: the name of the MNS topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:mns:cn-beijing:161456884340****:queue/modifyLifecycleHo****
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The types of Auto Scaling events and resource changes for which you want to receive notifications. You can specify multiple values by using a repeated list.
	//
	// You can invoke the DescribeNotificationTypes operation to query parameter values.
	//
	// This parameter is required.
	NotificationTypes []*string `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Repeated"`
	OwnerId           *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the scaling group belongs.
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
	// The time zone for notifications.
	//
	// Specify the time zone offset from Coordinated Universal Time (UTC), such as UTC+8 or UTC-7.
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

func (s *CreateNotificationConfigurationRequest) GetMessageEncoding() *string {
	return s.MessageEncoding
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

func (s *CreateNotificationConfigurationRequest) SetMessageEncoding(v string) *CreateNotificationConfigurationRequest {
	s.MessageEncoding = &v
	return s
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
