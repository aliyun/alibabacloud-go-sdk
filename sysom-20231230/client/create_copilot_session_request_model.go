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
	// 不填为 New Session
	//
	// example:
	//
	// 排查 CPU 负载
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
