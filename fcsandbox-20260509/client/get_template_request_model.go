// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTeamID(v string) *GetTemplateRequest
	GetTeamID() *string
}

type GetTemplateRequest struct {
	// The unique identifier of the team.
	//
	// example:
	//
	// 88a4c762-b0ce-4661-9413-578b2309e60f
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *GetTemplateRequest) SetTeamID(v string) *GetTemplateRequest {
	s.TeamID = &v
	return s
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
