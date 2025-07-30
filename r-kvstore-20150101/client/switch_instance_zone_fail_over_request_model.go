// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceZoneFailOverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SwitchInstanceZoneFailOverRequest
	GetInstanceId() *string
	SetSiteFaultTime(v string) *SwitchInstanceZoneFailOverRequest
	GetSiteFaultTime() *string
	SetTargetZoneId(v string) *SwitchInstanceZoneFailOverRequest
	GetTargetZoneId() *string
}

type SwitchInstanceZoneFailOverRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The duration for which the fault lasts. Unit: minutes.
	//
	// Valid values:
	//
	// 	- 5
	//
	// 	- 10
	//
	// example:
	//
	// 5
	SiteFaultTime *string `json:"SiteFaultTime,omitempty" xml:"SiteFaultTime,omitempty"`
	// The ID of the destination zone.
	//
	// example:
	//
	// cn-hangzhou-j
	TargetZoneId *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
}

func (s SwitchInstanceZoneFailOverRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceZoneFailOverRequest) GoString() string {
	return s.String()
}

func (s *SwitchInstanceZoneFailOverRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchInstanceZoneFailOverRequest) GetSiteFaultTime() *string {
	return s.SiteFaultTime
}

func (s *SwitchInstanceZoneFailOverRequest) GetTargetZoneId() *string {
	return s.TargetZoneId
}

func (s *SwitchInstanceZoneFailOverRequest) SetInstanceId(v string) *SwitchInstanceZoneFailOverRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchInstanceZoneFailOverRequest) SetSiteFaultTime(v string) *SwitchInstanceZoneFailOverRequest {
	s.SiteFaultTime = &v
	return s
}

func (s *SwitchInstanceZoneFailOverRequest) SetTargetZoneId(v string) *SwitchInstanceZoneFailOverRequest {
	s.TargetZoneId = &v
	return s
}

func (s *SwitchInstanceZoneFailOverRequest) Validate() error {
	return dara.Validate(s)
}
