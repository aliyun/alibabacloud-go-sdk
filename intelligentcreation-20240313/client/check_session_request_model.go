// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v string) *CheckSessionRequest
	GetProjectId() *string
	SetSessionId(v string) *CheckSessionRequest
	GetSessionId() *string
}

type CheckSessionRequest struct {
	// example:
	//
	// 11111
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 121dlsga4o7golrl1hoja
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CheckSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSessionRequest) GoString() string {
	return s.String()
}

func (s *CheckSessionRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CheckSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CheckSessionRequest) SetProjectId(v string) *CheckSessionRequest {
	s.ProjectId = &v
	return s
}

func (s *CheckSessionRequest) SetSessionId(v string) *CheckSessionRequest {
	s.SessionId = &v
	return s
}

func (s *CheckSessionRequest) Validate() error {
	return dara.Validate(s)
}
