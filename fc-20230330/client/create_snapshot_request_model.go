// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateSnapshotInput) *CreateSnapshotRequest
	GetBody() *CreateSnapshotInput
	SetQualifier(v string) *CreateSnapshotRequest
	GetQualifier() *string
}

type CreateSnapshotRequest struct {
	Body *CreateSnapshotInput `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// alias
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetBody() *CreateSnapshotInput {
	return s.Body
}

func (s *CreateSnapshotRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *CreateSnapshotRequest) SetBody(v *CreateSnapshotInput) *CreateSnapshotRequest {
	s.Body = v
	return s
}

func (s *CreateSnapshotRequest) SetQualifier(v string) *CreateSnapshotRequest {
	s.Qualifier = &v
	return s
}

func (s *CreateSnapshotRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
