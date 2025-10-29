// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTurnOnZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetZone(v string) *TurnOnZoneRequest
	GetZone() *string
}

type TurnOnZoneRequest struct {
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

func (s *TurnOnZoneRequest) GetZone() *string {
	return s.Zone
}

func (s *TurnOnZoneRequest) SetZone(v string) *TurnOnZoneRequest {
	s.Zone = &v
	return s
}

func (s *TurnOnZoneRequest) Validate() error {
	return dara.Validate(s)
}
