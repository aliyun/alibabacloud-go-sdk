// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNotificationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationArn(v string) *DeleteNotificationConfigurationRequest
	GetNotificationArn() *string
	SetOwnerId(v int64) *DeleteNotificationConfigurationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteNotificationConfigurationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteNotificationConfigurationRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DeleteNotificationConfigurationRequest
	GetScalingGroupId() *string
}

type DeleteNotificationConfigurationRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the notification recipient. Specify the value in one of the following formats:
	//
	// 	- If you specify CloudMonitor as the notification recipient, specify the value in the acs:ess:{region-id}:{account-id}:cloudmonitor format.
	//
	// 	- If you specify a Simple Message Queue (SMQ, formerly MNS) queue as the notification recipient, specify the value in the acs:mns:{region-id}:{account-id}:queue/{queuename} format.
	//
	// 	- If you specify an SMQ queue as the notification recipient, specify the value in the acs:mns:{region-id}:{account-id}:topic/{topicname} format.
	//
	// The variables in the preceding value formats have the following meanings:
	//
	// 	- region-id: the region ID of the scaling group.
	//
	// 	- account-id: the ID of your Alibaba Cloud cloud.
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
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DeleteNotificationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNotificationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteNotificationConfigurationRequest) GetNotificationArn() *string {
	return s.NotificationArn
}

func (s *DeleteNotificationConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteNotificationConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteNotificationConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteNotificationConfigurationRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DeleteNotificationConfigurationRequest) SetNotificationArn(v string) *DeleteNotificationConfigurationRequest {
	s.NotificationArn = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) SetOwnerId(v int64) *DeleteNotificationConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) SetRegionId(v string) *DeleteNotificationConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) SetResourceOwnerAccount(v string) *DeleteNotificationConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) SetScalingGroupId(v string) *DeleteNotificationConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DeleteNotificationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
