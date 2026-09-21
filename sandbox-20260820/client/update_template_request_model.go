// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PublicUpdateTemplateInput) *UpdateTemplateRequest
	GetBody() *PublicUpdateTemplateInput
	SetTeamID(v string) *UpdateTemplateRequest
	GetTeamID() *string
}

type UpdateTemplateRequest struct {
	Body *PublicUpdateTemplateInput `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// team-a1b2c3d4e5f6
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) GetBody() *PublicUpdateTemplateInput {
	return s.Body
}

func (s *UpdateTemplateRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *UpdateTemplateRequest) SetBody(v *PublicUpdateTemplateInput) *UpdateTemplateRequest {
	s.Body = v
	return s
}

func (s *UpdateTemplateRequest) SetTeamID(v string) *UpdateTemplateRequest {
	s.TeamID = &v
	return s
}

func (s *UpdateTemplateRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
