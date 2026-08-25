// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTeamID(v string) *DeleteVolumeRequest
	GetTeamID() *string
}

type DeleteVolumeRequest struct {
	// The unique identifier of the Team.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s DeleteVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVolumeRequest) GoString() string {
	return s.String()
}

func (s *DeleteVolumeRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *DeleteVolumeRequest) SetTeamID(v string) *DeleteVolumeRequest {
	s.TeamID = &v
	return s
}

func (s *DeleteVolumeRequest) Validate() error {
	return dara.Validate(s)
}
