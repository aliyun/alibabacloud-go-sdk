// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteByPassShuntEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteByPassShuntEventRequest
	GetLang() *string
	SetEventId(v int64) *DeleteByPassShuntEventRequest
	GetEventId() *int64
	SetRegId(v string) *DeleteByPassShuntEventRequest
	GetRegId() *string
}

type DeleteByPassShuntEventRequest struct {
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
	// 25
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteByPassShuntEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteByPassShuntEventRequest) GoString() string {
	return s.String()
}

func (s *DeleteByPassShuntEventRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteByPassShuntEventRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *DeleteByPassShuntEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteByPassShuntEventRequest) SetLang(v string) *DeleteByPassShuntEventRequest {
	s.Lang = &v
	return s
}

func (s *DeleteByPassShuntEventRequest) SetEventId(v int64) *DeleteByPassShuntEventRequest {
	s.EventId = &v
	return s
}

func (s *DeleteByPassShuntEventRequest) SetRegId(v string) *DeleteByPassShuntEventRequest {
	s.RegId = &v
	return s
}

func (s *DeleteByPassShuntEventRequest) Validate() error {
	return dara.Validate(s)
}
