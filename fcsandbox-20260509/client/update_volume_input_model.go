// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVolumeInput interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v string) *UpdateVolumeInput
	GetStatus() *string
	SetTeamID(v string) *UpdateVolumeInput
	GetTeamID() *string
}

type UpdateVolumeInput struct {
	// The target status of the volume.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The unique identifier of the team to which the volume belongs.
	//
	// example:
	//
	// 70d1c834-0383-58d8-97ac-5336eb91abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s UpdateVolumeInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateVolumeInput) GoString() string {
	return s.String()
}

func (s *UpdateVolumeInput) GetStatus() *string {
	return s.Status
}

func (s *UpdateVolumeInput) GetTeamID() *string {
	return s.TeamID
}

func (s *UpdateVolumeInput) SetStatus(v string) *UpdateVolumeInput {
	s.Status = &v
	return s
}

func (s *UpdateVolumeInput) SetTeamID(v string) *UpdateVolumeInput {
	s.TeamID = &v
	return s
}

func (s *UpdateVolumeInput) Validate() error {
	return dara.Validate(s)
}
