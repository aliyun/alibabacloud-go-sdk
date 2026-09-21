// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVolumeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateVolumeInput) *UpdateVolumeRequest
	GetBody() *UpdateVolumeInput
	SetTeamID(v string) *UpdateVolumeRequest
	GetTeamID() *string
}

type UpdateVolumeRequest struct {
	Body *UpdateVolumeInput `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// 9f5a1fe9-****
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s UpdateVolumeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVolumeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVolumeRequest) GetBody() *UpdateVolumeInput {
	return s.Body
}

func (s *UpdateVolumeRequest) GetTeamID() *string {
	return s.TeamID
}

func (s *UpdateVolumeRequest) SetBody(v *UpdateVolumeInput) *UpdateVolumeRequest {
	s.Body = v
	return s
}

func (s *UpdateVolumeRequest) SetTeamID(v string) *UpdateVolumeRequest {
	s.TeamID = &v
	return s
}

func (s *UpdateVolumeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
