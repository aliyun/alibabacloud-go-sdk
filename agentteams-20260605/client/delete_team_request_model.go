// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteTeamRequest
	GetInstanceId() *string
	SetName(v string) *DeleteTeamRequest
	GetName() *string
}

type DeleteTeamRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTeamRequest) GoString() string {
	return s.String()
}

func (s *DeleteTeamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteTeamRequest) GetName() *string {
	return s.Name
}

func (s *DeleteTeamRequest) SetInstanceId(v string) *DeleteTeamRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTeamRequest) SetName(v string) *DeleteTeamRequest {
	s.Name = &v
	return s
}

func (s *DeleteTeamRequest) Validate() error {
	return dara.Validate(s)
}
