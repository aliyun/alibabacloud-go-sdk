// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateEventsRequestBody) *CreateEventsRequest
	GetBody() *CreateEventsRequestBody
}

type CreateEventsRequest struct {
	// The JSON request body.
	//
	// This parameter is required.
	Body *CreateEventsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s CreateEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsRequest) GoString() string {
	return s.String()
}

func (s *CreateEventsRequest) GetBody() *CreateEventsRequestBody {
	return s.Body
}

func (s *CreateEventsRequest) SetBody(v *CreateEventsRequestBody) *CreateEventsRequest {
	s.Body = v
	return s
}

func (s *CreateEventsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventsRequestBody struct {
	// A list of 1 to 100 events. The request body cannot exceed 1 MiB.
	//
	// This parameter is required.
	Events []*CreateEventsRequestBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// The exact session to which all events belong.
	//
	// This parameter is required.
	Scope *CreateEventsRequestBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The information saved when the first event is successfully written. It is not overwritten by subsequent writes.
	Session *CreateEventsRequestBodySession `json:"session,omitempty" xml:"session,omitempty" type:"Struct"`
}

func (s CreateEventsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsRequestBody) GoString() string {
	return s.String()
}

func (s *CreateEventsRequestBody) GetEvents() []*CreateEventsRequestBodyEvents {
	return s.Events
}

func (s *CreateEventsRequestBody) GetScope() *CreateEventsRequestBodyScope {
	return s.Scope
}

func (s *CreateEventsRequestBody) GetSession() *CreateEventsRequestBodySession {
	return s.Session
}

func (s *CreateEventsRequestBody) SetEvents(v []*CreateEventsRequestBodyEvents) *CreateEventsRequestBody {
	s.Events = v
	return s
}

func (s *CreateEventsRequestBody) SetScope(v *CreateEventsRequestBodyScope) *CreateEventsRequestBody {
	s.Scope = v
	return s
}

func (s *CreateEventsRequestBody) SetSession(v *CreateEventsRequestBodySession) *CreateEventsRequestBody {
	s.Session = v
	return s
}

func (s *CreateEventsRequestBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	if s.Session != nil {
		if err := s.Session.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventsRequestBodyEvents struct {
	// The event-specific data. Use the corresponding fields based on the type. For conditional requirements, refer to the description of each field.
	//
	// This parameter is required.
	Data *CreateEventsRequestBodyEventsData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The actual occurrence time provided by the client as a Unix timestamp in milliseconds. This is the only source for business sorting. Identical values do not indicate a known sequence, and the array position does not imply a business order.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1758614400000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// The event type. This determines the applicable fields and conditional requirements for the data.
	//
	// This parameter is required.
	//
	// example:
	//
	// message
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateEventsRequestBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsRequestBodyEvents) GoString() string {
	return s.String()
}

func (s *CreateEventsRequestBodyEvents) GetData() *CreateEventsRequestBodyEventsData {
	return s.Data
}

func (s *CreateEventsRequestBodyEvents) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CreateEventsRequestBodyEvents) GetType() *string {
	return s.Type
}

func (s *CreateEventsRequestBodyEvents) SetData(v *CreateEventsRequestBodyEventsData) *CreateEventsRequestBodyEvents {
	s.Data = v
	return s
}

func (s *CreateEventsRequestBodyEvents) SetTimestamp(v int64) *CreateEventsRequestBodyEvents {
	s.Timestamp = &v
	return s
}

func (s *CreateEventsRequestBodyEvents) SetType(v string) *CreateEventsRequestBodyEvents {
	s.Type = &v
	return s
}

func (s *CreateEventsRequestBodyEvents) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventsRequestBodyEventsData struct {
	// The call ID. This field is required for tool_call and tool_result events, or message events with the role set to tool. It associates the call intent, actual execution, and feedback.
	//
	// example:
	//
	// call-001
	CallId *string `json:"callId,omitempty" xml:"callId,omitempty"`
	// The ordered content blocks. This field is required for message and tool_result events. Use a text block for plain text. This can be an empty array if the tool has no output.
	Content []*CreateEventsRequestBodyEventsDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The list of visible event IDs referenced by the summary. This list cannot be empty.
	EventIds []*string `json:"eventIds,omitempty" xml:"eventIds,omitempty" type:"Repeated"`
	// The anchor event ID of the parent session. This field is required for session_fork events. The child session continues from the valid context at this anchor and does not inherit interactions after the anchor.
	//
	// example:
	//
	// evt-2f5f9c5a8d0e4c3b
	ParentEventId *string `json:"parentEventId,omitempty" xml:"parentEventId,omitempty"`
	// The business scope of the parent session. This field is required for session_fork events. The event is written to the child session, and the requested scope is the child session.
	ParentScope *CreateEventsRequestBodyEventsDataParentScope `json:"parentScope,omitempty" xml:"parentScope,omitempty" type:"Struct"`
	// The message role. This field is required for message events.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The human-readable summary. This field is optional for context_update events.
	//
	// example:
	//
	// The user plans to go hiking on the weekend
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
}

func (s CreateEventsRequestBodyEventsData) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsRequestBodyEventsData) GoString() string {
	return s.String()
}

func (s *CreateEventsRequestBodyEventsData) GetCallId() *string {
	return s.CallId
}

func (s *CreateEventsRequestBodyEventsData) GetContent() []*CreateEventsRequestBodyEventsDataContent {
	return s.Content
}

func (s *CreateEventsRequestBodyEventsData) GetEventIds() []*string {
	return s.EventIds
}

func (s *CreateEventsRequestBodyEventsData) GetParentEventId() *string {
	return s.ParentEventId
}

func (s *CreateEventsRequestBodyEventsData) GetParentScope() *CreateEventsRequestBodyEventsDataParentScope {
	return s.ParentScope
}

func (s *CreateEventsRequestBodyEventsData) GetRole() *string {
	return s.Role
}

func (s *CreateEventsRequestBodyEventsData) GetSummary() *string {
	return s.Summary
}

func (s *CreateEventsRequestBodyEventsData) SetCallId(v string) *CreateEventsRequestBodyEventsData {
	s.CallId = &v
	return s
}

func (s *CreateEventsRequestBodyEventsData) SetContent(v []*CreateEventsRequestBodyEventsDataContent) *CreateEventsRequestBodyEventsData {
	s.Content = v
	return s
}

func (s *CreateEventsRequestBodyEventsData) SetEventIds(v []*string) *CreateEventsRequestBodyEventsData {
	s.EventIds = v
	return s
}

func (s *CreateEventsRequestBodyEventsData) SetParentEventId(v string) *CreateEventsRequestBodyEventsData {
	s.ParentEventId = &v
	return s
}

func (s *CreateEventsRequestBodyEventsData) SetParentScope(v *CreateEventsRequestBodyEventsDataParentScope) *CreateEventsRequestBodyEventsData {
	s.ParentScope = v
	return s
}

func (s *CreateEventsRequestBodyEventsData) SetRole(v string) *CreateEventsRequestBodyEventsData {
	s.Role = &v
	return s
}

func (s *CreateEventsRequestBodyEventsData) SetSummary(v string) *CreateEventsRequestBodyEventsData {
	s.Summary = &v
	return s
}

func (s *CreateEventsRequestBodyEventsData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ParentScope != nil {
		if err := s.ParentScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEventsRequestBodyEventsDataContent struct {
	// The JSON parameter values. This field is required for tool_call blocks. Unparsed raw strings can be retained.
	//
	// example:
	//
	// {"city":"Hangzhou"}
	Arguments interface{} `json:"arguments,omitempty" xml:"arguments,omitempty"`
	// The logical call ID. This field is required for tool_call blocks.
	//
	// example:
	//
	// call-001
	CallId *string `json:"callId,omitempty" xml:"callId,omitempty"`
	// The access format identifier. This field is required for opaque blocks. Saving the format does not mean that its content can be interpreted or used across platforms.
	//
	// example:
	//
	// application/json
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The tool name. This field is required for tool_call blocks. This content block only expresses the call intent.
	//
	// example:
	//
	// get_weather
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The actual text. This field is required for text and reasoning blocks.
	//
	// example:
	//
	// The user likes hiking
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The content block type. This field determines the required fields.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The JSON value. This field is required for json and opaque blocks. The data field is not used for opaque blocks.
	//
	// example:
	//
	// {"preference":"hiking"}
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateEventsRequestBodyEventsDataContent) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsRequestBodyEventsDataContent) GoString() string {
	return s.String()
}

func (s *CreateEventsRequestBodyEventsDataContent) GetArguments() interface{} {
	return s.Arguments
}

func (s *CreateEventsRequestBodyEventsDataContent) GetCallId() *string {
	return s.CallId
}

func (s *CreateEventsRequestBodyEventsDataContent) GetFormat() *string {
	return s.Format
}

func (s *CreateEventsRequestBodyEventsDataContent) GetName() *string {
	return s.Name
}

func (s *CreateEventsRequestBodyEventsDataContent) GetText() *string {
	return s.Text
}

func (s *CreateEventsRequestBodyEventsDataContent) GetType() *string {
	return s.Type
}

func (s *CreateEventsRequestBodyEventsDataContent) GetValue() interface{} {
	return s.Value
}

func (s *CreateEventsRequestBodyEventsDataContent) SetArguments(v interface{}) *CreateEventsRequestBodyEventsDataContent {
	s.Arguments = v
	return s
}

func (s *CreateEventsRequestBodyEventsDataContent) SetCallId(v string) *CreateEventsRequestBodyEventsDataContent {
	s.CallId = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataContent) SetFormat(v string) *CreateEventsRequestBodyEventsDataContent {
	s.Format = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataContent) SetName(v string) *CreateEventsRequestBodyEventsDataContent {
	s.Name = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataContent) SetText(v string) *CreateEventsRequestBodyEventsDataContent {
	s.Text = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataContent) SetType(v string) *CreateEventsRequestBodyEventsDataContent {
	s.Type = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataContent) SetValue(v interface{}) *CreateEventsRequestBodyEventsDataContent {
	s.Value = v
	return s
}

func (s *CreateEventsRequestBodyEventsDataContent) Validate() error {
	return dara.Validate(s)
}

type CreateEventsRequestBodyEventsDataParentScope struct {
	// The business agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The exact session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The business team ID.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The business user ID. This is not the platform tenant identity.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateEventsRequestBodyEventsDataParentScope) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsRequestBodyEventsDataParentScope) GoString() string {
	return s.String()
}

func (s *CreateEventsRequestBodyEventsDataParentScope) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateEventsRequestBodyEventsDataParentScope) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateEventsRequestBodyEventsDataParentScope) GetTeamId() *string {
	return s.TeamId
}

func (s *CreateEventsRequestBodyEventsDataParentScope) GetUserId() *string {
	return s.UserId
}

func (s *CreateEventsRequestBodyEventsDataParentScope) SetAgentId(v string) *CreateEventsRequestBodyEventsDataParentScope {
	s.AgentId = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataParentScope) SetSessionId(v string) *CreateEventsRequestBodyEventsDataParentScope {
	s.SessionId = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataParentScope) SetTeamId(v string) *CreateEventsRequestBodyEventsDataParentScope {
	s.TeamId = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataParentScope) SetUserId(v string) *CreateEventsRequestBodyEventsDataParentScope {
	s.UserId = &v
	return s
}

func (s *CreateEventsRequestBodyEventsDataParentScope) Validate() error {
	return dara.Validate(s)
}

type CreateEventsRequestBodyScope struct {
	// The business agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The exact session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The business team ID.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The business user ID. This is not the platform tenant identity.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateEventsRequestBodyScope) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsRequestBodyScope) GoString() string {
	return s.String()
}

func (s *CreateEventsRequestBodyScope) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateEventsRequestBodyScope) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateEventsRequestBodyScope) GetTeamId() *string {
	return s.TeamId
}

func (s *CreateEventsRequestBodyScope) GetUserId() *string {
	return s.UserId
}

func (s *CreateEventsRequestBodyScope) SetAgentId(v string) *CreateEventsRequestBodyScope {
	s.AgentId = &v
	return s
}

func (s *CreateEventsRequestBodyScope) SetSessionId(v string) *CreateEventsRequestBodyScope {
	s.SessionId = &v
	return s
}

func (s *CreateEventsRequestBodyScope) SetTeamId(v string) *CreateEventsRequestBodyScope {
	s.TeamId = &v
	return s
}

func (s *CreateEventsRequestBodyScope) SetUserId(v string) *CreateEventsRequestBodyScope {
	s.UserId = &v
	return s
}

func (s *CreateEventsRequestBodyScope) Validate() error {
	return dara.Validate(s)
}

type CreateEventsRequestBodySession struct {
	// The JSON metadata of the session.
	//
	// example:
	//
	// {"topic":"hiking"}
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The session title. It can be up to 256 characters in length.
	//
	// example:
	//
	// Travel plan
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateEventsRequestBodySession) String() string {
	return dara.Prettify(s)
}

func (s CreateEventsRequestBodySession) GoString() string {
	return s.String()
}

func (s *CreateEventsRequestBodySession) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *CreateEventsRequestBodySession) GetTitle() *string {
	return s.Title
}

func (s *CreateEventsRequestBodySession) SetMetadata(v map[string]*string) *CreateEventsRequestBodySession {
	s.Metadata = v
	return s
}

func (s *CreateEventsRequestBodySession) SetTitle(v string) *CreateEventsRequestBodySession {
	s.Title = &v
	return s
}

func (s *CreateEventsRequestBodySession) Validate() error {
	return dara.Validate(s)
}
