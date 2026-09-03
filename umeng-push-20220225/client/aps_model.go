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
	SetAttributes(v string) *Aps
	GetAttributes() *string
	SetAttributesType(v string) *Aps
	GetAttributesType() *string
	SetBadge(v string) *Aps
	GetBadge() *string
	SetCategory(v string) *Aps
	GetCategory() *string
	SetContentAvailable(v int32) *Aps
	GetContentAvailable() *int32
	SetContentState(v string) *Aps
	GetContentState() *string
	SetDismissalDate(v int32) *Aps
	GetDismissalDate() *int32
	SetEvent(v string) *Aps
	GetEvent() *string
	SetInterruptionLevel(v string) *Aps
	GetInterruptionLevel() *string
	SetMutableContent(v int32) *Aps
	GetMutableContent() *int32
	SetSound(v string) *Aps
	GetSound() *string
	SetThreadID(v string) *Aps
	GetThreadID() *string
	SetTimestamp(v int32) *Aps
	GetTimestamp() *int32
}

type Aps struct {
	Alert          *Alert  `json:"alert,omitempty" xml:"alert,omitempty"`
	Attributes     *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	AttributesType *string `json:"attributesType,omitempty" xml:"attributesType,omitempty"`
	// example:
	//
	// +1(自增)，-1(自减)，4(设置数字)
	Badge            *string `json:"badge,omitempty" xml:"badge,omitempty"`
	Category         *string `json:"category,omitempty" xml:"category,omitempty"`
	ContentAvailable *int32  `json:"contentAvailable,omitempty" xml:"contentAvailable,omitempty"`
	// example:
	//
	// {                  "status": "shippingbox.fill"                 }
	ContentState  *string `json:"contentState,omitempty" xml:"contentState,omitempty"`
	DismissalDate *int32  `json:"dismissalDate,omitempty" xml:"dismissalDate,omitempty"`
	// example:
	//
	// 创建:start , 更新:update,结束:end
	Event             *string `json:"event,omitempty" xml:"event,omitempty"`
	InterruptionLevel *string `json:"interruptionLevel,omitempty" xml:"interruptionLevel,omitempty"`
	// example:
	//
	// 1
	MutableContent *int32  `json:"mutableContent,omitempty" xml:"mutableContent,omitempty"`
	Sound          *string `json:"sound,omitempty" xml:"sound,omitempty"`
	ThreadID       *string `json:"threadID,omitempty" xml:"threadID,omitempty"`
	Timestamp      *int32  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
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

func (s *Aps) GetAttributes() *string {
	return s.Attributes
}

func (s *Aps) GetAttributesType() *string {
	return s.AttributesType
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

func (s *Aps) GetContentState() *string {
	return s.ContentState
}

func (s *Aps) GetDismissalDate() *int32 {
	return s.DismissalDate
}

func (s *Aps) GetEvent() *string {
	return s.Event
}

func (s *Aps) GetInterruptionLevel() *string {
	return s.InterruptionLevel
}

func (s *Aps) GetMutableContent() *int32 {
	return s.MutableContent
}

func (s *Aps) GetSound() *string {
	return s.Sound
}

func (s *Aps) GetThreadID() *string {
	return s.ThreadID
}

func (s *Aps) GetTimestamp() *int32 {
	return s.Timestamp
}

func (s *Aps) SetAlert(v *Alert) *Aps {
	s.Alert = v
	return s
}

func (s *Aps) SetAttributes(v string) *Aps {
	s.Attributes = &v
	return s
}

func (s *Aps) SetAttributesType(v string) *Aps {
	s.AttributesType = &v
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

func (s *Aps) SetContentState(v string) *Aps {
	s.ContentState = &v
	return s
}

func (s *Aps) SetDismissalDate(v int32) *Aps {
	s.DismissalDate = &v
	return s
}

func (s *Aps) SetEvent(v string) *Aps {
	s.Event = &v
	return s
}

func (s *Aps) SetInterruptionLevel(v string) *Aps {
	s.InterruptionLevel = &v
	return s
}

func (s *Aps) SetMutableContent(v int32) *Aps {
	s.MutableContent = &v
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

func (s *Aps) SetTimestamp(v int32) *Aps {
	s.Timestamp = &v
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
