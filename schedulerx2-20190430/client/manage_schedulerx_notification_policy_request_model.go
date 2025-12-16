// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxNotificationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelTimeRange(v string) *ManageSchedulerxNotificationPolicyRequest
	GetChannelTimeRange() *string
	SetDescription(v string) *ManageSchedulerxNotificationPolicyRequest
	GetDescription() *string
	SetPolicyName(v string) *ManageSchedulerxNotificationPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *ManageSchedulerxNotificationPolicyRequest
	GetRegionId() *string
}

type ManageSchedulerxNotificationPolicyRequest struct {
	// The time range configuration for notification channels.
	//
	// >  See the supplementary description of ChannelTimeRange in the request parameters.
	//
	// example:
	//
	// {
	//
	//   "sendChannel": "sms,webhook,mail,phone",
	//
	//   "timezone": "UTC+08:00",
	//
	//   "webhookIsAtAll": "false",
	//
	//   "timeRanges": {
	//
	//     "all": [
	//
	//       {
	//
	//         "startTime": "08:00",
	//
	//         "endTime": "18:00",
	//
	//         "daysOfWeek": [1, 2, 3, 4, 5]
	//
	//       }
	//
	//     ]
	//
	//   }
	//
	// }
	ChannelTimeRange *string `json:"ChannelTimeRange,omitempty" xml:"ChannelTimeRange,omitempty"`
	// The notification policy description.
	//
	// example:
	//
	// Monday-Friday only
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the notification policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-weekdays
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ManageSchedulerxNotificationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxNotificationPolicyRequest) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxNotificationPolicyRequest) GetChannelTimeRange() *string {
	return s.ChannelTimeRange
}

func (s *ManageSchedulerxNotificationPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *ManageSchedulerxNotificationPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ManageSchedulerxNotificationPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ManageSchedulerxNotificationPolicyRequest) SetChannelTimeRange(v string) *ManageSchedulerxNotificationPolicyRequest {
	s.ChannelTimeRange = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyRequest) SetDescription(v string) *ManageSchedulerxNotificationPolicyRequest {
	s.Description = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyRequest) SetPolicyName(v string) *ManageSchedulerxNotificationPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyRequest) SetRegionId(v string) *ManageSchedulerxNotificationPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
