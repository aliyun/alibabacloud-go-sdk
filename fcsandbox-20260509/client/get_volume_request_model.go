// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTeamID(v string) *GetVolumeRequest
	GetTeamID() *string
}

type GetVolumeRequest struct {
	// The unique identifier of the Team.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s GetVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVolumeRequest) GoString() string {
	return s.String()
}

func (s *GetVolumeRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *GetVolumeRequest) SetTeamID(v string) *GetVolumeRequest {
	s.TeamID = &v
	return s
}

func (s *GetVolumeRequest) Validate() error {
	return dara.Validate(s)
}
