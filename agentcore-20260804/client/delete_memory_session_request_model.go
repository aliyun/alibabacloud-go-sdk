// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemorySessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DeleteMemorySessionRequest
	GetAgentId() *string
	SetTeamId(v string) *DeleteMemorySessionRequest
	GetTeamId() *string
	SetUserId(v string) *DeleteMemorySessionRequest
	GetUserId() *string
}

type DeleteMemorySessionRequest struct {
	// agentId
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// teamId
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// userId
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s DeleteMemorySessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemorySessionRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemorySessionRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteMemorySessionRequest) GetTeamId() *string {
	return s.TeamId
}

func (s *DeleteMemorySessionRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteMemorySessionRequest) SetAgentId(v string) *DeleteMemorySessionRequest {
	s.AgentId = &v
	return s
}

func (s *DeleteMemorySessionRequest) SetTeamId(v string) *DeleteMemorySessionRequest {
	s.TeamId = &v
	return s
}

func (s *DeleteMemorySessionRequest) SetUserId(v string) *DeleteMemorySessionRequest {
	s.UserId = &v
	return s
}

func (s *DeleteMemorySessionRequest) Validate() error {
	return dara.Validate(s)
}
