// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateArtifactResponse
	GetStatusCode() *int32
	SetBody(v *Artifact) *CreateArtifactResponse
	GetBody() *Artifact
}

type CreateArtifactResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Artifact          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateArtifactResponse) GetBody() *Artifact {
	return s.Body
}

func (s *CreateArtifactResponse) SetHeaders(v map[string]*string) *CreateArtifactResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactResponse) SetStatusCode(v int32) *CreateArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactResponse) SetBody(v *Artifact) *CreateArtifactResponse {
	s.Body = v
	return s
}

func (s *CreateArtifactResponse) Validate() error {
	return dara.Validate(s)
}
