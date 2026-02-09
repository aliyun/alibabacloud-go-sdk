// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePptArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePptArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePptArtifactResponse
	GetStatusCode() *int32
	SetBody(v *DeletePptArtifactResponseBody) *DeletePptArtifactResponse
	GetBody() *DeletePptArtifactResponseBody
}

type DeletePptArtifactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePptArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePptArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePptArtifactResponse) GoString() string {
	return s.String()
}

func (s *DeletePptArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePptArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePptArtifactResponse) GetBody() *DeletePptArtifactResponseBody {
	return s.Body
}

func (s *DeletePptArtifactResponse) SetHeaders(v map[string]*string) *DeletePptArtifactResponse {
	s.Headers = v
	return s
}

func (s *DeletePptArtifactResponse) SetStatusCode(v int32) *DeletePptArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePptArtifactResponse) SetBody(v *DeletePptArtifactResponseBody) *DeletePptArtifactResponse {
	s.Body = v
	return s
}

func (s *DeletePptArtifactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
