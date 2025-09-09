// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseArtifactResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseArtifactResponseBody) *ReleaseArtifactResponse
	GetBody() *ReleaseArtifactResponseBody
}

type ReleaseArtifactResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseArtifactResponse) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseArtifactResponse) GetBody() *ReleaseArtifactResponseBody {
	return s.Body
}

func (s *ReleaseArtifactResponse) SetHeaders(v map[string]*string) *ReleaseArtifactResponse {
	s.Headers = v
	return s
}

func (s *ReleaseArtifactResponse) SetStatusCode(v int32) *ReleaseArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseArtifactResponse) SetBody(v *ReleaseArtifactResponseBody) *ReleaseArtifactResponse {
	s.Body = v
	return s
}

func (s *ReleaseArtifactResponse) Validate() error {
	return dara.Validate(s)
}
