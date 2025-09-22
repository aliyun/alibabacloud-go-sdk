// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTicketRecordByticketCodePopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgendaId(v string) *UpdateTicketRecordByticketCodePopRequest
	GetAgendaId() *string
	SetCode(v string) *UpdateTicketRecordByticketCodePopRequest
	GetCode() *string
	SetEvent(v string) *UpdateTicketRecordByticketCodePopRequest
	GetEvent() *string
	SetSceneId(v string) *UpdateTicketRecordByticketCodePopRequest
	GetSceneId() *string
	SetTime(v string) *UpdateTicketRecordByticketCodePopRequest
	GetTime() *string
}

type UpdateTicketRecordByticketCodePopRequest struct {
	// example:
	//
	// 3402
	AgendaId *string `json:"AgendaId,omitempty" xml:"AgendaId,omitempty"`
	// example:
	//
	// 4546-100000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// boarding
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 12345
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 2024-09-25 14:11
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s UpdateTicketRecordByticketCodePopRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTicketRecordByticketCodePopRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRecordByticketCodePopRequest) GetAgendaId() *string {
	return s.AgendaId
}

func (s *UpdateTicketRecordByticketCodePopRequest) GetCode() *string {
	return s.Code
}

func (s *UpdateTicketRecordByticketCodePopRequest) GetEvent() *string {
	return s.Event
}

func (s *UpdateTicketRecordByticketCodePopRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateTicketRecordByticketCodePopRequest) GetTime() *string {
	return s.Time
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetAgendaId(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.AgendaId = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetCode(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.Code = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetEvent(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.Event = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetSceneId(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) SetTime(v string) *UpdateTicketRecordByticketCodePopRequest {
	s.Time = &v
	return s
}

func (s *UpdateTicketRecordByticketCodePopRequest) Validate() error {
	return dara.Validate(s)
}
