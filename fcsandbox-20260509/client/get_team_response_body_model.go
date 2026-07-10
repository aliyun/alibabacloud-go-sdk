// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTeamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTeamResponseBody
	GetCode() *string
	SetMessage(v string) *GetTeamResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTeamResponseBody
	GetRequestId() *string
	SetTeam(v *E2BTeam) *GetTeamResponseBody
	GetTeam() *E2BTeam
}

type GetTeamResponseBody struct {
	Code      *string  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Team      *E2BTeam `json:"team,omitempty" xml:"team,omitempty"`
}

func (s GetTeamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTeamResponseBody) GoString() string {
	return s.String()
}

func (s *GetTeamResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTeamResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTeamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTeamResponseBody) GetTeam() *E2BTeam {
	return s.Team
}

func (s *GetTeamResponseBody) SetCode(v string) *GetTeamResponseBody {
	s.Code = &v
	return s
}

func (s *GetTeamResponseBody) SetMessage(v string) *GetTeamResponseBody {
	s.Message = &v
	return s
}

func (s *GetTeamResponseBody) SetRequestId(v string) *GetTeamResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTeamResponseBody) SetTeam(v *E2BTeam) *GetTeamResponseBody {
	s.Team = v
	return s
}

func (s *GetTeamResponseBody) Validate() error {
	if s.Team != nil {
		if err := s.Team.Validate(); err != nil {
			return err
		}
	}
	return nil
}
