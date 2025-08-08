// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetArtifactResponse
	GetStatusCode() *int32
	SetBody(v *Artifact) *GetArtifactResponse
	GetBody() *Artifact
}

type GetArtifactResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Artifact          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetArtifactResponse) GetBody() *Artifact {
	return s.Body
}

func (s *GetArtifactResponse) SetHeaders(v map[string]*string) *GetArtifactResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactResponse) SetStatusCode(v int32) *GetArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactResponse) SetBody(v *Artifact) *GetArtifactResponse {
	s.Body = v
	return s
}

func (s *GetArtifactResponse) Validate() error {
	return dara.Validate(s)
}
