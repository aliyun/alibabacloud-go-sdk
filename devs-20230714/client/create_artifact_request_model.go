// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Artifact) *CreateArtifactRequest
	GetBody() *Artifact
}

type CreateArtifactRequest struct {
	Body *Artifact `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequest) GetBody() *Artifact {
	return s.Body
}

func (s *CreateArtifactRequest) SetBody(v *Artifact) *CreateArtifactRequest {
	s.Body = v
	return s
}

func (s *CreateArtifactRequest) Validate() error {
	return dara.Validate(s)
}
