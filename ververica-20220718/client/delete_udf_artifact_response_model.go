// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdfArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUdfArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUdfArtifactResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUdfArtifactResponseBody) *DeleteUdfArtifactResponse
	GetBody() *DeleteUdfArtifactResponseBody
}

type DeleteUdfArtifactResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUdfArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUdfArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdfArtifactResponse) GoString() string {
	return s.String()
}

func (s *DeleteUdfArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUdfArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUdfArtifactResponse) GetBody() *DeleteUdfArtifactResponseBody {
	return s.Body
}

func (s *DeleteUdfArtifactResponse) SetHeaders(v map[string]*string) *DeleteUdfArtifactResponse {
	s.Headers = v
	return s
}

func (s *DeleteUdfArtifactResponse) SetStatusCode(v int32) *DeleteUdfArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUdfArtifactResponse) SetBody(v *DeleteUdfArtifactResponseBody) *DeleteUdfArtifactResponse {
	s.Body = v
	return s
}

func (s *DeleteUdfArtifactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
