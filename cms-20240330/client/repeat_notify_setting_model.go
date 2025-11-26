// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepeatNotifySetting interface {
	dara.Model
	String() string
	GoString() string
	SetEndIncidentState(v string) *RepeatNotifySetting
	GetEndIncidentState() *string
	SetRepeatInterval(v string) *RepeatNotifySetting
	GetRepeatInterval() *string
}

type RepeatNotifySetting struct {
	EndIncidentState *string `json:"endIncidentState,omitempty" xml:"endIncidentState,omitempty"`
	RepeatInterval   *string `json:"repeatInterval,omitempty" xml:"repeatInterval,omitempty"`
}

func (s RepeatNotifySetting) String() string {
	return dara.Prettify(s)
}

func (s RepeatNotifySetting) GoString() string {
	return s.String()
}

func (s *RepeatNotifySetting) GetEndIncidentState() *string {
	return s.EndIncidentState
}

func (s *RepeatNotifySetting) GetRepeatInterval() *string {
	return s.RepeatInterval
}

func (s *RepeatNotifySetting) SetEndIncidentState(v string) *RepeatNotifySetting {
	s.EndIncidentState = &v
	return s
}

func (s *RepeatNotifySetting) SetRepeatInterval(v string) *RepeatNotifySetting {
	s.RepeatInterval = &v
	return s
}

func (s *RepeatNotifySetting) Validate() error {
	return dara.Validate(s)
}
