// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupNotifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *CreateMonitorGroupNotifyPolicyRequest
	GetEndTime() *int64
	SetGroupId(v string) *CreateMonitorGroupNotifyPolicyRequest
	GetGroupId() *string
	SetPolicyType(v string) *CreateMonitorGroupNotifyPolicyRequest
	GetPolicyType() *string
	SetRegionId(v string) *CreateMonitorGroupNotifyPolicyRequest
	GetRegionId() *string
	SetStartTime(v int64) *CreateMonitorGroupNotifyPolicyRequest
	GetStartTime() *int64
}

type CreateMonitorGroupNotifyPolicyRequest struct {
	// The end time of the validity period for the policy.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1623208500000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7301****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The policy type. Valid value: PauseNotify.
	//
	// This parameter is required.
	//
	// example:
	//
	// PauseNotify
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the validity period for the policy.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1622949300000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateMonitorGroupNotifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupNotifyPolicyRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateMonitorGroupNotifyPolicyRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateMonitorGroupNotifyPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreateMonitorGroupNotifyPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMonitorGroupNotifyPolicyRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetEndTime(v int64) *CreateMonitorGroupNotifyPolicyRequest {
	s.EndTime = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetGroupId(v string) *CreateMonitorGroupNotifyPolicyRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetPolicyType(v string) *CreateMonitorGroupNotifyPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetRegionId(v string) *CreateMonitorGroupNotifyPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyRequest) SetStartTime(v int64) *CreateMonitorGroupNotifyPolicyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
