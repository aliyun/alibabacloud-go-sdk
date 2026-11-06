// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *GetAgentSessionEventsResponseBody
	GetData() interface{}
	SetEvent(v string) *GetAgentSessionEventsResponseBody
	GetEvent() *string
	SetId(v string) *GetAgentSessionEventsResponseBody
	GetId() *string
	SetRequestId(v string) *GetAgentSessionEventsResponseBody
	GetRequestId() *string
}

type GetAgentSessionEventsResponseBody struct {
	// The event data of the current frame. The content varies with the Event type and should be parsed according to the event type. Preserve the original session, turn, and tool invocation identifiers.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The SSE event type, such as session_update. Process control frames and business frames separately.
	//
	// example:
	//
	// session_update
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The remote SSE event cursor. Save it after consumption for subsequent resumption using Last-Event-Id. Some control frames may not have a cursor. Do not generate a local cursor for them.
	//
	// example:
	//
	// 42
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID of the current OpenAPI subscription request, used to troubleshoot subscription failures. It is not an event cursor or an interactive question ID.
	//
	// example:
	//
	// 01A0B6E0-E005-1465-B5DE-5CD17677426F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentSessionEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSessionEventsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetAgentSessionEventsResponseBody) GetEvent() *string {
	return s.Event
}

func (s *GetAgentSessionEventsResponseBody) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSessionEventsResponseBody) SetData(v interface{}) *GetAgentSessionEventsResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentSessionEventsResponseBody) SetEvent(v string) *GetAgentSessionEventsResponseBody {
	s.Event = &v
	return s
}

func (s *GetAgentSessionEventsResponseBody) SetId(v string) *GetAgentSessionEventsResponseBody {
	s.Id = &v
	return s
}

func (s *GetAgentSessionEventsResponseBody) SetRequestId(v string) *GetAgentSessionEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSessionEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
