// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventEpoch(v string) *GetAgentSessionEventsRequest
	GetEventEpoch() *string
	SetSessionCode(v string) *GetAgentSessionEventsRequest
	GetSessionCode() *string
}

type GetAgentSessionEventsRequest struct {
	// The epoch identifier of the event stream, obtained from data.eventEpoch in the stream_open control frame of the server. Pass the original value only when Last-Event-Id of the same native session and the same epoch is provided. Omit this parameter when subscribing for the first time or when the epoch is not obtained. Do not construct it yourself or reuse it across native sessions.
	//
	// example:
	//
	// epoch_example
	EventEpoch *string `json:"EventEpoch,omitempty" xml:"EventEpoch,omitempty"`
	// The ID of the target agent session. Set the value to the SessionId returned by CreateAgentSession. The parameter name for this API operation is SessionCode.
	//
	// This parameter is required.
	//
	// example:
	//
	// session_example
	SessionCode *string `json:"SessionCode,omitempty" xml:"SessionCode,omitempty"`
}

func (s GetAgentSessionEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionEventsRequest) GetEventEpoch() *string {
	return s.EventEpoch
}

func (s *GetAgentSessionEventsRequest) GetSessionCode() *string {
	return s.SessionCode
}

func (s *GetAgentSessionEventsRequest) SetEventEpoch(v string) *GetAgentSessionEventsRequest {
	s.EventEpoch = &v
	return s
}

func (s *GetAgentSessionEventsRequest) SetSessionCode(v string) *GetAgentSessionEventsRequest {
	s.SessionCode = &v
	return s
}

func (s *GetAgentSessionEventsRequest) Validate() error {
	return dara.Validate(s)
}
