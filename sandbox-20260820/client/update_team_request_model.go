// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateTeamInput) *UpdateTeamRequest
	GetBody() *UpdateTeamInput
}

type UpdateTeamRequest struct {
	Body *UpdateTeamInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamRequest) GoString() string {
	return s.String()
}

func (s *UpdateTeamRequest) GetBody() *UpdateTeamInput {
	return s.Body
}

func (s *UpdateTeamRequest) SetBody(v *UpdateTeamInput) *UpdateTeamRequest {
	s.Body = v
	return s
}

func (s *UpdateTeamRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
