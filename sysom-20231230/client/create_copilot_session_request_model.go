// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateCopilotSessionRequest
	GetName() *string
}

type CreateCopilotSessionRequest struct {
	// The session name. The name must be 1 to 255 characters in length. Default value: New Session. Empty strings or null values are not supported. If no name is specified, the server uses content generation to create a session name based on the first user message when it is received.
	//
	// example:
	//
	// Troubleshoot CPU load
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateCopilotSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateCopilotSessionRequest) GetName() *string {
	return s.Name
}

func (s *CreateCopilotSessionRequest) SetName(v string) *CreateCopilotSessionRequest {
	s.Name = &v
	return s
}

func (s *CreateCopilotSessionRequest) Validate() error {
	return dara.Validate(s)
}
