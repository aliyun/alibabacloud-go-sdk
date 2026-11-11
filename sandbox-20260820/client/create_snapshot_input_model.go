// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotInput interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSnapshotInput
	GetName() *string
	SetSandboxID(v string) *CreateSnapshotInput
	GetSandboxID() *string
	SetTeamID(v string) *CreateSnapshotInput
	GetTeamID() *string
}

type CreateSnapshotInput struct {
	// The snapshot name.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The source sandbox ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sbx-****
	SandboxID *string `json:"sandboxID,omitempty" xml:"sandboxID,omitempty"`
	// The team ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 70d1****abcd
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
}

func (s CreateSnapshotInput) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotInput) GoString() string {
	return s.String()
}

func (s *CreateSnapshotInput) GetName() *string {
	return s.Name
}

func (s *CreateSnapshotInput) GetSandboxID() *string {
	return s.SandboxID
}

func (s *CreateSnapshotInput) GetTeamID() *string {
	return s.TeamID
}

func (s *CreateSnapshotInput) SetName(v string) *CreateSnapshotInput {
	s.Name = &v
	return s
}

func (s *CreateSnapshotInput) SetSandboxID(v string) *CreateSnapshotInput {
	s.SandboxID = &v
	return s
}

func (s *CreateSnapshotInput) SetTeamID(v string) *CreateSnapshotInput {
	s.TeamID = &v
	return s
}

func (s *CreateSnapshotInput) Validate() error {
	return dara.Validate(s)
}
