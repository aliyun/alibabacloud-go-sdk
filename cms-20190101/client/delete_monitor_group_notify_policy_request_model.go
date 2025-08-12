// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupNotifyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteMonitorGroupNotifyPolicyRequest
	GetGroupId() *string
	SetPolicyType(v string) *DeleteMonitorGroupNotifyPolicyRequest
	GetPolicyType() *string
	SetRegionId(v string) *DeleteMonitorGroupNotifyPolicyRequest
	GetRegionId() *string
}

type DeleteMonitorGroupNotifyPolicyRequest struct {
	// The ID of the application group.
	//
	// example:
	//
	// 6780****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The type of the policy.
	//
	// Valid value: PauseNotify.
	//
	// This parameter is required.
	//
	// example:
	//
	// PauseNotify
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteMonitorGroupNotifyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupNotifyPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) SetGroupId(v string) *DeleteMonitorGroupNotifyPolicyRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) SetPolicyType(v string) *DeleteMonitorGroupNotifyPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) SetRegionId(v string) *DeleteMonitorGroupNotifyPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
