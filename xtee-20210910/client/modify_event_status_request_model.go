// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEventStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ModifyEventStatusRequest
	GetLang() *string
	SetCreateType(v string) *ModifyEventStatusRequest
	GetCreateType() *string
	SetEventCode(v string) *ModifyEventStatusRequest
	GetEventCode() *string
	SetFromEventSatus(v string) *ModifyEventStatusRequest
	GetFromEventSatus() *string
	SetRegId(v string) *ModifyEventStatusRequest
	GetRegId() *string
	SetToEventSatus(v string) *ModifyEventStatusRequest
	GetToEventSatus() *string
}

type ModifyEventStatusRequest struct {
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
	// Creation type
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aamexg3015
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Initial event status, to avoid duplicate submissions or historical replays
	//
	// example:
	//
	// ONLINE
	FromEventSatus *string `json:"fromEventSatus,omitempty" xml:"fromEventSatus,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Updated event status
	//
	// example:
	//
	// OFFLINE
	ToEventSatus *string `json:"toEventSatus,omitempty" xml:"toEventSatus,omitempty"`
}

func (s ModifyEventStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEventStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyEventStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyEventStatusRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *ModifyEventStatusRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *ModifyEventStatusRequest) GetFromEventSatus() *string {
	return s.FromEventSatus
}

func (s *ModifyEventStatusRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModifyEventStatusRequest) GetToEventSatus() *string {
	return s.ToEventSatus
}

func (s *ModifyEventStatusRequest) SetLang(v string) *ModifyEventStatusRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEventStatusRequest) SetCreateType(v string) *ModifyEventStatusRequest {
	s.CreateType = &v
	return s
}

func (s *ModifyEventStatusRequest) SetEventCode(v string) *ModifyEventStatusRequest {
	s.EventCode = &v
	return s
}

func (s *ModifyEventStatusRequest) SetFromEventSatus(v string) *ModifyEventStatusRequest {
	s.FromEventSatus = &v
	return s
}

func (s *ModifyEventStatusRequest) SetRegId(v string) *ModifyEventStatusRequest {
	s.RegId = &v
	return s
}

func (s *ModifyEventStatusRequest) SetToEventSatus(v string) *ModifyEventStatusRequest {
	s.ToEventSatus = &v
	return s
}

func (s *ModifyEventStatusRequest) Validate() error {
	return dara.Validate(s)
}
