// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTeamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTeamRequest
	GetInstanceId() *string
	SetName(v string) *GetTeamRequest
	GetName() *string
}

type GetTeamRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetTeamRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTeamRequest) GoString() string {
	return s.String()
}

func (s *GetTeamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTeamRequest) GetName() *string {
	return s.Name
}

func (s *GetTeamRequest) SetInstanceId(v string) *GetTeamRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTeamRequest) SetName(v string) *GetTeamRequest {
	s.Name = &v
	return s
}

func (s *GetTeamRequest) Validate() error {
	return dara.Validate(s)
}
