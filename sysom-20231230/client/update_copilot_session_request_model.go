// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCopilotSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateCopilotSessionRequest
	GetName() *string
	SetStatus(v string) *UpdateCopilotSessionRequest
	GetStatus() *string
}

type UpdateCopilotSessionRequest struct {
	// The session title. The length is 1 to 255 characters.
	//
	// example:
	//
	// CPU load analysis record
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The session status.
	//
	// example:
	//
	// archived
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateCopilotSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCopilotSessionRequest) GoString() string {
	return s.String()
}

func (s *UpdateCopilotSessionRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCopilotSessionRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateCopilotSessionRequest) SetName(v string) *UpdateCopilotSessionRequest {
	s.Name = &v
	return s
}

func (s *UpdateCopilotSessionRequest) SetStatus(v string) *UpdateCopilotSessionRequest {
	s.Status = &v
	return s
}

func (s *UpdateCopilotSessionRequest) Validate() error {
	return dara.Validate(s)
}
