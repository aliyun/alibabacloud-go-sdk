// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartOrStopByPassShuntEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *StartOrStopByPassShuntEventRequest
	GetLang() *string
	SetEventId(v int64) *StartOrStopByPassShuntEventRequest
	GetEventId() *int64
	SetRegId(v string) *StartOrStopByPassShuntEventRequest
	GetRegId() *string
	SetStatus(v string) *StartOrStopByPassShuntEventRequest
	GetStatus() *string
}

type StartOrStopByPassShuntEventRequest struct {
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
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Status.
	//
	// example:
	//
	// 1
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StartOrStopByPassShuntEventRequest) String() string {
	return dara.Prettify(s)
}

func (s StartOrStopByPassShuntEventRequest) GoString() string {
	return s.String()
}

func (s *StartOrStopByPassShuntEventRequest) GetLang() *string {
	return s.Lang
}

func (s *StartOrStopByPassShuntEventRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *StartOrStopByPassShuntEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *StartOrStopByPassShuntEventRequest) GetStatus() *string {
	return s.Status
}

func (s *StartOrStopByPassShuntEventRequest) SetLang(v string) *StartOrStopByPassShuntEventRequest {
	s.Lang = &v
	return s
}

func (s *StartOrStopByPassShuntEventRequest) SetEventId(v int64) *StartOrStopByPassShuntEventRequest {
	s.EventId = &v
	return s
}

func (s *StartOrStopByPassShuntEventRequest) SetRegId(v string) *StartOrStopByPassShuntEventRequest {
	s.RegId = &v
	return s
}

func (s *StartOrStopByPassShuntEventRequest) SetStatus(v string) *StartOrStopByPassShuntEventRequest {
	s.Status = &v
	return s
}

func (s *StartOrStopByPassShuntEventRequest) Validate() error {
	return dara.Validate(s)
}
