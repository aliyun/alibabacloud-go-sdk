// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTurnOffZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetZone(v string) *TurnOffZoneRequest
	GetZone() *string
}

type TurnOffZoneRequest struct {
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

func (s *TurnOffZoneRequest) GetZone() *string {
	return s.Zone
}

func (s *TurnOffZoneRequest) SetZone(v string) *TurnOffZoneRequest {
	s.Zone = &v
	return s
}

func (s *TurnOffZoneRequest) Validate() error {
	return dara.Validate(s)
}
