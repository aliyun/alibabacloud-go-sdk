// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateTeamInput) *CreateTeamRequest
	GetBody() *CreateTeamInput
}

type CreateTeamRequest struct {
	Body *CreateTeamInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamRequest) GoString() string {
	return s.String()
}

func (s *CreateTeamRequest) GetBody() *CreateTeamInput {
	return s.Body
}

func (s *CreateTeamRequest) SetBody(v *CreateTeamInput) *CreateTeamRequest {
	s.Body = v
	return s
}

func (s *CreateTeamRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
