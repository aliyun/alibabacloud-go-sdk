// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateInstanceZoneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetZoneId(v string) *MigrateInstanceZoneRequest
	GetZoneId() *string
}

type MigrateInstanceZoneRequest struct {
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s MigrateInstanceZoneRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateInstanceZoneRequest) GoString() string {
	return s.String()
}

func (s *MigrateInstanceZoneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *MigrateInstanceZoneRequest) SetZoneId(v string) *MigrateInstanceZoneRequest {
	s.ZoneId = &v
	return s
}

func (s *MigrateInstanceZoneRequest) Validate() error {
	return dara.Validate(s)
}
