// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *DeleteEventsRequestBody) *DeleteEventsRequest
	GetBody() *DeleteEventsRequestBody
}

type DeleteEventsRequest struct {
	// The JSON request body.
	//
	// This parameter is required.
	Body *DeleteEventsRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s DeleteEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventsRequest) GetBody() *DeleteEventsRequestBody {
	return s.Body
}

func (s *DeleteEventsRequest) SetBody(v *DeleteEventsRequestBody) *DeleteEventsRequest {
	s.Body = v
	return s
}

func (s *DeleteEventsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteEventsRequestBody struct {
	// A list of 1 to 100 unique event IDs. Only events owned by the current session are deleted.
	//
	// This parameter is required.
	EventIds []*string `json:"eventIds,omitempty" xml:"eventIds,omitempty" type:"Repeated"`
	// The scope for precise session location.
	//
	// This parameter is required.
	Scope *DeleteEventsRequestBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
}

func (s DeleteEventsRequestBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsRequestBody) GoString() string {
	return s.String()
}

func (s *DeleteEventsRequestBody) GetEventIds() []*string {
	return s.EventIds
}

func (s *DeleteEventsRequestBody) GetScope() *DeleteEventsRequestBodyScope {
	return s.Scope
}

func (s *DeleteEventsRequestBody) SetEventIds(v []*string) *DeleteEventsRequestBody {
	s.EventIds = v
	return s
}

func (s *DeleteEventsRequestBody) SetScope(v *DeleteEventsRequestBodyScope) *DeleteEventsRequestBody {
	s.Scope = v
	return s
}

func (s *DeleteEventsRequestBody) Validate() error {
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteEventsRequestBodyScope struct {
	// The business agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The ID of the precisely located session.
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

func (s DeleteEventsRequestBodyScope) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventsRequestBodyScope) GoString() string {
	return s.String()
}

func (s *DeleteEventsRequestBodyScope) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteEventsRequestBodyScope) GetSessionId() *string {
	return s.SessionId
}

func (s *DeleteEventsRequestBodyScope) GetTeamId() *string {
	return s.TeamId
}

func (s *DeleteEventsRequestBodyScope) GetUserId() *string {
	return s.UserId
}

func (s *DeleteEventsRequestBodyScope) SetAgentId(v string) *DeleteEventsRequestBodyScope {
	s.AgentId = &v
	return s
}

func (s *DeleteEventsRequestBodyScope) SetSessionId(v string) *DeleteEventsRequestBodyScope {
	s.SessionId = &v
	return s
}

func (s *DeleteEventsRequestBodyScope) SetTeamId(v string) *DeleteEventsRequestBodyScope {
	s.TeamId = &v
	return s
}

func (s *DeleteEventsRequestBodyScope) SetUserId(v string) *DeleteEventsRequestBodyScope {
	s.UserId = &v
	return s
}

func (s *DeleteEventsRequestBodyScope) Validate() error {
	return dara.Validate(s)
}
