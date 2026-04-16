// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataZoneSupportCompactionServiceValue interface {
	dara.Model
	String() string
	GoString() string
	SetZoneId(v string) *DataZoneSupportCompactionServiceValue
	GetZoneId() *string
	SetResourceLevel(v string) *DataZoneSupportCompactionServiceValue
	GetResourceLevel() *string
	SetRecommended(v bool) *DataZoneSupportCompactionServiceValue
	GetRecommended() *bool
}

type DataZoneSupportCompactionServiceValue struct {
	ZoneId        *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
	ResourceLevel *string `json:"resourceLevel,omitempty" xml:"resourceLevel,omitempty"`
	Recommended   *bool   `json:"recommended,omitempty" xml:"recommended,omitempty"`
}

func (s DataZoneSupportCompactionServiceValue) String() string {
	return dara.Prettify(s)
}

func (s DataZoneSupportCompactionServiceValue) GoString() string {
	return s.String()
}

func (s *DataZoneSupportCompactionServiceValue) GetZoneId() *string {
	return s.ZoneId
}

func (s *DataZoneSupportCompactionServiceValue) GetResourceLevel() *string {
	return s.ResourceLevel
}

func (s *DataZoneSupportCompactionServiceValue) GetRecommended() *bool {
	return s.Recommended
}

func (s *DataZoneSupportCompactionServiceValue) SetZoneId(v string) *DataZoneSupportCompactionServiceValue {
	s.ZoneId = &v
	return s
}

func (s *DataZoneSupportCompactionServiceValue) SetResourceLevel(v string) *DataZoneSupportCompactionServiceValue {
	s.ResourceLevel = &v
	return s
}

func (s *DataZoneSupportCompactionServiceValue) SetRecommended(v bool) *DataZoneSupportCompactionServiceValue {
	s.Recommended = &v
	return s
}

func (s *DataZoneSupportCompactionServiceValue) Validate() error {
	return dara.Validate(s)
}
