// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Artifact) *PutArtifactRequest
	GetBody() *Artifact
	SetForce(v bool) *PutArtifactRequest
	GetForce() *bool
}

type PutArtifactRequest struct {
	Body *Artifact `json:"body,omitempty" xml:"body,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s PutArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s PutArtifactRequest) GoString() string {
	return s.String()
}

func (s *PutArtifactRequest) GetBody() *Artifact {
	return s.Body
}

func (s *PutArtifactRequest) GetForce() *bool {
	return s.Force
}

func (s *PutArtifactRequest) SetBody(v *Artifact) *PutArtifactRequest {
	s.Body = v
	return s
}

func (s *PutArtifactRequest) SetForce(v bool) *PutArtifactRequest {
	s.Force = &v
	return s
}

func (s *PutArtifactRequest) Validate() error {
	return dara.Validate(s)
}
