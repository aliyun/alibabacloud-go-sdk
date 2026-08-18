// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSnapshotInput
	GetDescription() *string
	SetSessionId(v string) *CreateSnapshotInput
	GetSessionId() *string
}

type CreateSnapshotInput struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// custom-test-session-id
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CreateSnapshotInput) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotInput) GoString() string {
	return s.String()
}

func (s *CreateSnapshotInput) GetDescription() *string {
	return s.Description
}

func (s *CreateSnapshotInput) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateSnapshotInput) SetDescription(v string) *CreateSnapshotInput {
	s.Description = &v
	return s
}

func (s *CreateSnapshotInput) SetSessionId(v string) *CreateSnapshotInput {
	s.SessionId = &v
	return s
}

func (s *CreateSnapshotInput) Validate() error {
	return dara.Validate(s)
}
