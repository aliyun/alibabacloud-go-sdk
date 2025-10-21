// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUdfArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UdfArtifact) *CreateUdfArtifactRequest
	GetBody() *UdfArtifact
}

type CreateUdfArtifactRequest struct {
	// The resource file of the UDF.
	//
	// This parameter is required.
	Body *UdfArtifact `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUdfArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUdfArtifactRequest) GoString() string {
	return s.String()
}

func (s *CreateUdfArtifactRequest) GetBody() *UdfArtifact {
	return s.Body
}

func (s *CreateUdfArtifactRequest) SetBody(v *UdfArtifact) *CreateUdfArtifactRequest {
	s.Body = v
	return s
}

func (s *CreateUdfArtifactRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
