// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteArtifactResponse
	GetStatusCode() *int32
	SetBody(v *DeleteArtifactResponseBody) *DeleteArtifactResponse
	GetBody() *DeleteArtifactResponseBody
}

type DeleteArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactResponse) GoString() string {
	return s.String()
}

func (s *DeleteArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteArtifactResponse) GetBody() *DeleteArtifactResponseBody {
	return s.Body
}

func (s *DeleteArtifactResponse) SetHeaders(v map[string]*string) *DeleteArtifactResponse {
	s.Headers = v
	return s
}

func (s *DeleteArtifactResponse) SetStatusCode(v int32) *DeleteArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteArtifactResponse) SetBody(v *DeleteArtifactResponseBody) *DeleteArtifactResponse {
	s.Body = v
	return s
}

func (s *DeleteArtifactResponse) Validate() error {
	return dara.Validate(s)
}
