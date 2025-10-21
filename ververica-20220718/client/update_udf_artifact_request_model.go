// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UdfArtifact) *UpdateUdfArtifactRequest
	GetBody() *UdfArtifact
}

type UpdateUdfArtifactRequest struct {
	// The details of the JAR file of the UDF.
	//
	// This parameter is required.
	Body *UdfArtifact `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUdfArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfArtifactRequest) GoString() string {
	return s.String()
}

func (s *UpdateUdfArtifactRequest) GetBody() *UdfArtifact {
	return s.Body
}

func (s *UpdateUdfArtifactRequest) SetBody(v *UdfArtifact) *UpdateUdfArtifactRequest {
	s.Body = v
	return s
}

func (s *UpdateUdfArtifactRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
