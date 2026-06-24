// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTurnOffZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHpAlbZoneDrained(v bool) *TurnOffZoneRequest
	GetHpAlbZoneDrained() *bool
	SetZone(v string) *TurnOffZoneRequest
	GetZone() *string
}

type TurnOffZoneRequest struct {
	HpAlbZoneDrained *bool `json:"hpAlbZoneDrained,omitempty" xml:"hpAlbZoneDrained,omitempty"`
	// The zone of the instance.
	//
	// example:
	//
	// cn-hangzhou-i
	Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
}

func (s TurnOffZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s TurnOffZoneRequest) GoString() string {
	return s.String()
}

func (s *TurnOffZoneRequest) GetHpAlbZoneDrained() *bool {
	return s.HpAlbZoneDrained
}

func (s *TurnOffZoneRequest) GetZone() *string {
	return s.Zone
}

func (s *TurnOffZoneRequest) SetHpAlbZoneDrained(v bool) *TurnOffZoneRequest {
	s.HpAlbZoneDrained = &v
	return s
}

func (s *TurnOffZoneRequest) SetZone(v string) *TurnOffZoneRequest {
	s.Zone = &v
	return s
}

func (s *TurnOffZoneRequest) Validate() error {
	return dara.Validate(s)
}
