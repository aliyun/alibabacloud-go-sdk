// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAps interface {
	dara.Model
	String() string
	GoString() string
	SetAlert(v *Alert) *Aps
	GetAlert() *Alert
	SetBadge(v string) *Aps
	GetBadge() *string
	SetCategory(v string) *Aps
	GetCategory() *string
	SetContentAvailable(v int32) *Aps
	GetContentAvailable() *int32
	SetInterruptionLevel(v string) *Aps
	GetInterruptionLevel() *string
	SetSound(v string) *Aps
	GetSound() *string
	SetThreadID(v string) *Aps
	GetThreadID() *string
}

type Aps struct {
	Alert *Alert `json:"alert,omitempty" xml:"alert,omitempty"`
	// example:
	//
	// +1(自增)，-1(自减)，4(设置数字)
	Badge             *string `json:"badge,omitempty" xml:"badge,omitempty"`
	Category          *string `json:"category,omitempty" xml:"category,omitempty"`
	ContentAvailable  *int32  `json:"contentAvailable,omitempty" xml:"contentAvailable,omitempty"`
	InterruptionLevel *string `json:"interruptionLevel,omitempty" xml:"interruptionLevel,omitempty"`
	Sound             *string `json:"sound,omitempty" xml:"sound,omitempty"`
	ThreadID          *string `json:"threadID,omitempty" xml:"threadID,omitempty"`
}

func (s Aps) String() string {
	return dara.Prettify(s)
}

func (s Aps) GoString() string {
	return s.String()
}

func (s *Aps) GetAlert() *Alert {
	return s.Alert
}

func (s *Aps) GetBadge() *string {
	return s.Badge
}

func (s *Aps) GetCategory() *string {
	return s.Category
}

func (s *Aps) GetContentAvailable() *int32 {
	return s.ContentAvailable
}

func (s *Aps) GetInterruptionLevel() *string {
	return s.InterruptionLevel
}

func (s *Aps) GetSound() *string {
	return s.Sound
}

func (s *Aps) GetThreadID() *string {
	return s.ThreadID
}

func (s *Aps) SetAlert(v *Alert) *Aps {
	s.Alert = v
	return s
}

func (s *Aps) SetBadge(v string) *Aps {
	s.Badge = &v
	return s
}

func (s *Aps) SetCategory(v string) *Aps {
	s.Category = &v
	return s
}

func (s *Aps) SetContentAvailable(v int32) *Aps {
	s.ContentAvailable = &v
	return s
}

func (s *Aps) SetInterruptionLevel(v string) *Aps {
	s.InterruptionLevel = &v
	return s
}

func (s *Aps) SetSound(v string) *Aps {
	s.Sound = &v
	return s
}

func (s *Aps) SetThreadID(v string) *Aps {
	s.ThreadID = &v
	return s
}

func (s *Aps) Validate() error {
	if s.Alert != nil {
		if err := s.Alert.Validate(); err != nil {
			return err
		}
	}
	return nil
}
