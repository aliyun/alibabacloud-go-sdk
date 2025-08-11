// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotificationConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotificationConfigurationModels(v []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) *DescribeNotificationConfigurationsResponseBody
	GetNotificationConfigurationModels() []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels
	SetRequestId(v string) *DescribeNotificationConfigurationsResponseBody
	GetRequestId() *string
}

type DescribeNotificationConfigurationsResponseBody struct {
	// The notification settings.
	NotificationConfigurationModels []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels `json:"NotificationConfigurationModels,omitempty" xml:"NotificationConfigurationModels,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNotificationConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponseBody) GetNotificationConfigurationModels() []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	return s.NotificationConfigurationModels
}

func (s *DescribeNotificationConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNotificationConfigurationsResponseBody) SetNotificationConfigurationModels(v []*DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) *DescribeNotificationConfigurationsResponseBody {
	s.NotificationConfigurationModels = v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBody) SetRequestId(v string) *DescribeNotificationConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels struct {
	MessageEncoding *string `json:"MessageEncoding,omitempty" xml:"MessageEncoding,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the notification recipient. The value is in one of the following formats:
	//
	// 	- If you specify CloudMonitor as the notification recipient, the value is in the acs:ess:{region-id}:{account-id}:cloudmonitor format.
	//
	// 	- If you specify a Simple Message Queue (SMQ, formerly MNS) as the notification recipient, the value is in the acs:mns:{region-id}:{account-id}:queue/{queuename} format.
	//
	// 	- If you specify an SMQ topic as the notification recipient, the value is in the acs:mns:{region-id}:{account-id}:topic/{topicname} format.
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
	// example:
	//
	// acs:mns:cn-beijing:161456884340****:topic/modifyLifecycleHo****
	NotificationArn *string `json:"NotificationArn,omitempty" xml:"NotificationArn,omitempty"`
	// The types of the notifications.
	NotificationTypes []*string `json:"NotificationTypes,omitempty" xml:"NotificationTypes,omitempty" type:"Repeated"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// The time zone of the notification. The value must be in UTC. For example, a value of UTC+8 indicates that the time is 8 hours ahead of Coordinated Universal Time, and a value of UTC-7 indicates that the time is 7 hours behind Coordinated Universal Time.
	//
	// example:
	//
	// UTC+8
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) GetMessageEncoding() *string {
	return s.MessageEncoding
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) GetNotificationArn() *string {
	return s.NotificationArn
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) GetNotificationTypes() []*string {
	return s.NotificationTypes
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) GetTimeZone() *string {
	return s.TimeZone
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetMessageEncoding(v string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.MessageEncoding = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetNotificationArn(v string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.NotificationArn = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetNotificationTypes(v []*string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.NotificationTypes = v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetScalingGroupId(v string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) SetTimeZone(v string) *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels {
	s.TimeZone = &v
	return s
}

func (s *DescribeNotificationConfigurationsResponseBodyNotificationConfigurationModels) Validate() error {
	return dara.Validate(s)
}
