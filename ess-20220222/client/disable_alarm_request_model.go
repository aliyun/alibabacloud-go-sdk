// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmTaskId(v string) *DisableAlarmRequest
	GetAlarmTaskId() *string
	SetOwnerId(v int64) *DisableAlarmRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DisableAlarmRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DisableAlarmRequest
	GetResourceOwnerAccount() *string
}

type DisableAlarmRequest struct {
	// The ID of the event-triggered task.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1hvbnmkl10vll5****_f95ce797-dc2e-4bad-9618-14fee7d1****
	AlarmTaskId *string `json:"AlarmTaskId,omitempty" xml:"AlarmTaskId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s DisableAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAlarmRequest) GoString() string {
	return s.String()
}

func (s *DisableAlarmRequest) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *DisableAlarmRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableAlarmRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableAlarmRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableAlarmRequest) SetAlarmTaskId(v string) *DisableAlarmRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *DisableAlarmRequest) SetOwnerId(v int64) *DisableAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableAlarmRequest) SetRegionId(v string) *DisableAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *DisableAlarmRequest) SetResourceOwnerAccount(v string) *DisableAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableAlarmRequest) Validate() error {
	return dara.Validate(s)
}
