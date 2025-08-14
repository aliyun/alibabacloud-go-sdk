// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveByPassOrShuntEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SaveByPassOrShuntEventRequest
	GetLang() *string
	SetEventId(v int64) *SaveByPassOrShuntEventRequest
	GetEventId() *int64
	SetEventName(v string) *SaveByPassOrShuntEventRequest
	GetEventName() *string
	SetEventType(v string) *SaveByPassOrShuntEventRequest
	GetEventType() *string
	SetRegId(v string) *SaveByPassOrShuntEventRequest
	GetRegId() *string
}

type SaveByPassOrShuntEventRequest struct {
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
	// Event ID.
	//
	// example:
	//
	// 445
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册事件
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Event type
	//
	// example:
	//
	// MAIN
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s SaveByPassOrShuntEventRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveByPassOrShuntEventRequest) GoString() string {
	return s.String()
}

func (s *SaveByPassOrShuntEventRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveByPassOrShuntEventRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *SaveByPassOrShuntEventRequest) GetEventName() *string {
	return s.EventName
}

func (s *SaveByPassOrShuntEventRequest) GetEventType() *string {
	return s.EventType
}

func (s *SaveByPassOrShuntEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *SaveByPassOrShuntEventRequest) SetLang(v string) *SaveByPassOrShuntEventRequest {
	s.Lang = &v
	return s
}

func (s *SaveByPassOrShuntEventRequest) SetEventId(v int64) *SaveByPassOrShuntEventRequest {
	s.EventId = &v
	return s
}

func (s *SaveByPassOrShuntEventRequest) SetEventName(v string) *SaveByPassOrShuntEventRequest {
	s.EventName = &v
	return s
}

func (s *SaveByPassOrShuntEventRequest) SetEventType(v string) *SaveByPassOrShuntEventRequest {
	s.EventType = &v
	return s
}

func (s *SaveByPassOrShuntEventRequest) SetRegId(v string) *SaveByPassOrShuntEventRequest {
	s.RegId = &v
	return s
}

func (s *SaveByPassOrShuntEventRequest) Validate() error {
	return dara.Validate(s)
}
