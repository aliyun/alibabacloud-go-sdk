// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerxNotificationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelTimeRange(v string) *CreateSchedulerxNotificationPolicyRequest
	GetChannelTimeRange() *string
	SetDescription(v string) *CreateSchedulerxNotificationPolicyRequest
	GetDescription() *string
	SetPolicyName(v string) *CreateSchedulerxNotificationPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *CreateSchedulerxNotificationPolicyRequest
	GetRegionId() *string
}

type CreateSchedulerxNotificationPolicyRequest struct {
	// The configuration for the effective time periods of notification channels.
	//
	// >  Please see the detailed explanation of this parameter below.
	//
	// This parameter is required.
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
	// The description of the notification policy.
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

func (s CreateSchedulerxNotificationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerxNotificationPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateSchedulerxNotificationPolicyRequest) GetChannelTimeRange() *string {
	return s.ChannelTimeRange
}

func (s *CreateSchedulerxNotificationPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSchedulerxNotificationPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreateSchedulerxNotificationPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSchedulerxNotificationPolicyRequest) SetChannelTimeRange(v string) *CreateSchedulerxNotificationPolicyRequest {
	s.ChannelTimeRange = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyRequest) SetDescription(v string) *CreateSchedulerxNotificationPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyRequest) SetPolicyName(v string) *CreateSchedulerxNotificationPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyRequest) SetRegionId(v string) *CreateSchedulerxNotificationPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
