// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerxNotificationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *DeleteSchedulerxNotificationPolicyRequest
	GetPolicyName() *string
	SetRegionId(v string) *DeleteSchedulerxNotificationPolicyRequest
	GetRegionId() *string
}

type DeleteSchedulerxNotificationPolicyRequest struct {
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

func (s DeleteSchedulerxNotificationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerxNotificationPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerxNotificationPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DeleteSchedulerxNotificationPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSchedulerxNotificationPolicyRequest) SetPolicyName(v string) *DeleteSchedulerxNotificationPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyRequest) SetRegionId(v string) *DeleteSchedulerxNotificationPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
