// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTurnOnZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHpAlbZoneDrained(v bool) *TurnOnZoneRequest
	GetHpAlbZoneDrained() *bool
	SetZone(v string) *TurnOnZoneRequest
	GetZone() *string
}

type TurnOnZoneRequest struct {
	HpAlbZoneDrained *bool `json:"hpAlbZoneDrained,omitempty" xml:"hpAlbZoneDrained,omitempty"`
	// The zone of the instance.
	//
	// example:
	//
	// cn-hangzhou-i
	Zone *string `json:"zone,omitempty" xml:"zone,omitempty"`
}

func (s TurnOnZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s TurnOnZoneRequest) GoString() string {
	return s.String()
}

func (s *TurnOnZoneRequest) GetHpAlbZoneDrained() *bool {
	return s.HpAlbZoneDrained
}

func (s *TurnOnZoneRequest) GetZone() *string {
	return s.Zone
}

func (s *TurnOnZoneRequest) SetHpAlbZoneDrained(v bool) *TurnOnZoneRequest {
	s.HpAlbZoneDrained = &v
	return s
}

func (s *TurnOnZoneRequest) SetZone(v string) *TurnOnZoneRequest {
	s.Zone = &v
	return s
}

func (s *TurnOnZoneRequest) Validate() error {
	return dara.Validate(s)
}
