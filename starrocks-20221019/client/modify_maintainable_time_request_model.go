// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaintainableTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyMaintainableTimeRequest
	GetInstanceId() *string
	SetMaintainableTimePeriod(v string) *ModifyMaintainableTimeRequest
	GetMaintainableTimePeriod() *string
}

type ModifyMaintainableTimeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 08:00-09:00
	MaintainableTimePeriod *string `json:"MaintainableTimePeriod,omitempty" xml:"MaintainableTimePeriod,omitempty"`
}

func (s ModifyMaintainableTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaintainableTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaintainableTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyMaintainableTimeRequest) GetMaintainableTimePeriod() *string {
	return s.MaintainableTimePeriod
}

func (s *ModifyMaintainableTimeRequest) SetInstanceId(v string) *ModifyMaintainableTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyMaintainableTimeRequest) SetMaintainableTimePeriod(v string) *ModifyMaintainableTimeRequest {
	s.MaintainableTimePeriod = &v
	return s
}

func (s *ModifyMaintainableTimeRequest) Validate() error {
	return dara.Validate(s)
}
