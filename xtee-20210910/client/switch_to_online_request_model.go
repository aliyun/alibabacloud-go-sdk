// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchToOnlineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SwitchToOnlineRequest
	GetLang() *string
	SetEventId(v int64) *SwitchToOnlineRequest
	GetEventId() *int64
	SetRegId(v string) *SwitchToOnlineRequest
	GetRegId() *string
}

type SwitchToOnlineRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The event ID to switch to.
	//
	// This parameter is required.
	//
	// example:
	//
	// de_ahqhsw7665
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s SwitchToOnlineRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchToOnlineRequest) GoString() string {
	return s.String()
}

func (s *SwitchToOnlineRequest) GetLang() *string {
	return s.Lang
}

func (s *SwitchToOnlineRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *SwitchToOnlineRequest) GetRegId() *string {
	return s.RegId
}

func (s *SwitchToOnlineRequest) SetLang(v string) *SwitchToOnlineRequest {
	s.Lang = &v
	return s
}

func (s *SwitchToOnlineRequest) SetEventId(v int64) *SwitchToOnlineRequest {
	s.EventId = &v
	return s
}

func (s *SwitchToOnlineRequest) SetRegId(v string) *SwitchToOnlineRequest {
	s.RegId = &v
	return s
}

func (s *SwitchToOnlineRequest) Validate() error {
	return dara.Validate(s)
}
