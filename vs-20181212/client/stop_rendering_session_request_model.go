// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRenderingSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *StopRenderingSessionRequest
	GetClientId() *string
	SetProjectId(v string) *StopRenderingSessionRequest
	GetProjectId() *string
	SetSessionId(v string) *StopRenderingSessionRequest
	GetSessionId() *string
}

type StopRenderingSessionRequest struct {
	// example:
	//
	// 04c30850-1d91-4da1-b811-66d0ee94af7d
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// session-i205217481741918129226
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StopRenderingSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRenderingSessionRequest) GoString() string {
	return s.String()
}

func (s *StopRenderingSessionRequest) GetClientId() *string {
	return s.ClientId
}

func (s *StopRenderingSessionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *StopRenderingSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopRenderingSessionRequest) SetClientId(v string) *StopRenderingSessionRequest {
	s.ClientId = &v
	return s
}

func (s *StopRenderingSessionRequest) SetProjectId(v string) *StopRenderingSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *StopRenderingSessionRequest) SetSessionId(v string) *StopRenderingSessionRequest {
	s.SessionId = &v
	return s
}

func (s *StopRenderingSessionRequest) Validate() error {
	return dara.Validate(s)
}
