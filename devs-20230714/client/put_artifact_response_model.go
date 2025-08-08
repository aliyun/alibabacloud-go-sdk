// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutArtifactResponse
	GetStatusCode() *int32
	SetBody(v *Artifact) *PutArtifactResponse
	GetBody() *Artifact
}

type PutArtifactResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Artifact          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s PutArtifactResponse) GoString() string {
	return s.String()
}

func (s *PutArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutArtifactResponse) GetBody() *Artifact {
	return s.Body
}

func (s *PutArtifactResponse) SetHeaders(v map[string]*string) *PutArtifactResponse {
	s.Headers = v
	return s
}

func (s *PutArtifactResponse) SetStatusCode(v int32) *PutArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *PutArtifactResponse) SetBody(v *Artifact) *PutArtifactResponse {
	s.Body = v
	return s
}

func (s *PutArtifactResponse) Validate() error {
	return dara.Validate(s)
}
