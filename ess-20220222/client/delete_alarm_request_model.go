// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlarmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmTaskId(v string) *DeleteAlarmRequest
	GetAlarmTaskId() *string
	SetOwnerId(v int64) *DeleteAlarmRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteAlarmRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteAlarmRequest
	GetResourceOwnerAccount() *string
}

type DeleteAlarmRequest struct {
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

func (s DeleteAlarmRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlarmRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlarmRequest) GetAlarmTaskId() *string {
	return s.AlarmTaskId
}

func (s *DeleteAlarmRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAlarmRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAlarmRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAlarmRequest) SetAlarmTaskId(v string) *DeleteAlarmRequest {
	s.AlarmTaskId = &v
	return s
}

func (s *DeleteAlarmRequest) SetOwnerId(v int64) *DeleteAlarmRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAlarmRequest) SetRegionId(v string) *DeleteAlarmRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlarmRequest) SetResourceOwnerAccount(v string) *DeleteAlarmRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAlarmRequest) Validate() error {
	return dara.Validate(s)
}
