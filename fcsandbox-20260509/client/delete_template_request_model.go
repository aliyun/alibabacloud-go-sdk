// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTeamID(v string) *DeleteTemplateRequest
	GetTeamID() *string
}

type DeleteTemplateRequest struct {
	// The unique identifier of the team.
	//
	// example:
	//
	// 88a4c762-b0ce-4661-9413-578b2309e60f
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *DeleteTemplateRequest) SetTeamID(v string) *DeleteTemplateRequest {
	s.TeamID = &v
	return s
}

func (s *DeleteTemplateRequest) Validate() error {
	return dara.Validate(s)
}
