// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAvatarSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *StopAvatarSessionRequest
	GetProjectId() *string
	SetSessionId(v string) *StopAvatarSessionRequest
	GetSessionId() *string
}

type StopAvatarSessionRequest struct {
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s StopAvatarSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAvatarSessionRequest) GoString() string {
	return s.String()
}

func (s *StopAvatarSessionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *StopAvatarSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StopAvatarSessionRequest) SetProjectId(v string) *StopAvatarSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *StopAvatarSessionRequest) SetSessionId(v string) *StopAvatarSessionRequest {
	s.SessionId = &v
	return s
}

func (s *StopAvatarSessionRequest) Validate() error {
	return dara.Validate(s)
}
