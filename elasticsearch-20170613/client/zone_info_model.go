// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iZoneInfo interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *ZoneInfo
	GetStatus() *string
	SetZoneId(v string) *ZoneInfo
	GetZoneId() *string
}

type ZoneInfo struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s ZoneInfo) String() string {
	return dara.Prettify(s)
}

func (s ZoneInfo) GoString() string {
	return s.String()
}

func (s *ZoneInfo) GetStatus() *string {
	return s.Status
}

func (s *ZoneInfo) GetZoneId() *string {
	return s.ZoneId
}

func (s *ZoneInfo) SetStatus(v string) *ZoneInfo {
	s.Status = &v
	return s
}

func (s *ZoneInfo) SetZoneId(v string) *ZoneInfo {
	s.ZoneId = &v
	return s
}

func (s *ZoneInfo) Validate() error {
	return dara.Validate(s)
}
