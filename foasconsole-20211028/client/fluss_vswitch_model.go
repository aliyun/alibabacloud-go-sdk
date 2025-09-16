// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlussVswitch interface {
	dara.Model
	String() string
	GoString() string
	SetVSwitchIds(v []*string) *FlussVswitch
	GetVSwitchIds() []*string
	SetZoneId(v string) *FlussVswitch
	GetZoneId() *string
}

type FlussVswitch struct {
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	ZoneId     *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s FlussVswitch) String() string {
	return dara.Prettify(s)
}

func (s FlussVswitch) GoString() string {
	return s.String()
}

func (s *FlussVswitch) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *FlussVswitch) GetZoneId() *string {
	return s.ZoneId
}

func (s *FlussVswitch) SetVSwitchIds(v []*string) *FlussVswitch {
	s.VSwitchIds = v
	return s
}

func (s *FlussVswitch) SetZoneId(v string) *FlussVswitch {
	s.ZoneId = &v
	return s
}

func (s *FlussVswitch) Validate() error {
	return dara.Validate(s)
}
