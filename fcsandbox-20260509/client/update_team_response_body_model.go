// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateTeamResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateTeamResponseBody
	GetRequestId() *string
	SetTeam(v *E2BTeam) *UpdateTeamResponseBody
	GetTeam() *E2BTeam
}

type UpdateTeamResponseBody struct {
	Code      *string  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Team      *E2BTeam `json:"team,omitempty" xml:"team,omitempty"`
}

func (s UpdateTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTeamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTeamResponseBody) GetTeam() *E2BTeam {
	return s.Team
}

func (s *UpdateTeamResponseBody) SetCode(v string) *UpdateTeamResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTeamResponseBody) SetMessage(v string) *UpdateTeamResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTeamResponseBody) SetRequestId(v string) *UpdateTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTeamResponseBody) SetTeam(v *E2BTeam) *UpdateTeamResponseBody {
	s.Team = v
	return s
}

func (s *UpdateTeamResponseBody) Validate() error {
	if s.Team != nil {
		if err := s.Team.Validate(); err != nil {
			return err
		}
	}
	return nil
}
