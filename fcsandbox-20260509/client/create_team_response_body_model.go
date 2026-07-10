// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTeamResponseBody
	GetCode() *string
	SetMessage(v string) *CreateTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTeamResponseBody
	GetRequestId() *string
	SetTeam(v *E2BTeam) *CreateTeamResponseBody
	GetTeam() *E2BTeam
}

type CreateTeamResponseBody struct {
	Code      *string  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Team      *E2BTeam `json:"team,omitempty" xml:"team,omitempty"`
}

func (s CreateTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTeamResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTeamResponseBody) GetTeam() *E2BTeam {
	return s.Team
}

func (s *CreateTeamResponseBody) SetCode(v string) *CreateTeamResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTeamResponseBody) SetMessage(v string) *CreateTeamResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTeamResponseBody) SetRequestId(v string) *CreateTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTeamResponseBody) SetTeam(v *E2BTeam) *CreateTeamResponseBody {
	s.Team = v
	return s
}

func (s *CreateTeamResponseBody) Validate() error {
	if s.Team != nil {
		if err := s.Team.Validate(); err != nil {
			return err
		}
	}
	return nil
}
