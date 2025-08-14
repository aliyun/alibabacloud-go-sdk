// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateByPassShuntEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateByPassShuntEventRequest
	GetLang() *string
	SetEventId(v int64) *UpdateByPassShuntEventRequest
	GetEventId() *int64
	SetEventName(v string) *UpdateByPassShuntEventRequest
	GetEventName() *string
	SetRegId(v string) *UpdateByPassShuntEventRequest
	GetRegId() *string
}

type UpdateByPassShuntEventRequest struct {
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
	// 233
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// Event name.
	//
	// example:
	//
	// 用户昵称文本审核检测结果
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s UpdateByPassShuntEventRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateByPassShuntEventRequest) GoString() string {
	return s.String()
}

func (s *UpdateByPassShuntEventRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateByPassShuntEventRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *UpdateByPassShuntEventRequest) GetEventName() *string {
	return s.EventName
}

func (s *UpdateByPassShuntEventRequest) GetRegId() *string {
	return s.RegId
}

func (s *UpdateByPassShuntEventRequest) SetLang(v string) *UpdateByPassShuntEventRequest {
	s.Lang = &v
	return s
}

func (s *UpdateByPassShuntEventRequest) SetEventId(v int64) *UpdateByPassShuntEventRequest {
	s.EventId = &v
	return s
}

func (s *UpdateByPassShuntEventRequest) SetEventName(v string) *UpdateByPassShuntEventRequest {
	s.EventName = &v
	return s
}

func (s *UpdateByPassShuntEventRequest) SetRegId(v string) *UpdateByPassShuntEventRequest {
	s.RegId = &v
	return s
}

func (s *UpdateByPassShuntEventRequest) Validate() error {
	return dara.Validate(s)
}
