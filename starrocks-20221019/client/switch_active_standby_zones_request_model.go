// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchActiveStandbyZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SwitchActiveStandbyZonesRequest
	GetInstanceId() *string
	SetTargetZoneId(v string) *SwitchActiveStandbyZonesRequest
	GetTargetZoneId() *string
}

type SwitchActiveStandbyZonesRequest struct {
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	TargetZoneId *string `json:"TargetZoneId,omitempty" xml:"TargetZoneId,omitempty"`
}

func (s SwitchActiveStandbyZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchActiveStandbyZonesRequest) GoString() string {
	return s.String()
}

func (s *SwitchActiveStandbyZonesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchActiveStandbyZonesRequest) GetTargetZoneId() *string {
	return s.TargetZoneId
}

func (s *SwitchActiveStandbyZonesRequest) SetInstanceId(v string) *SwitchActiveStandbyZonesRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchActiveStandbyZonesRequest) SetTargetZoneId(v string) *SwitchActiveStandbyZonesRequest {
	s.TargetZoneId = &v
	return s
}

func (s *SwitchActiveStandbyZonesRequest) Validate() error {
	return dara.Validate(s)
}
