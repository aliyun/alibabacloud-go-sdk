// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPptArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPptArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPptArtifactResponse
	GetStatusCode() *int32
	SetBody(v *GetPptArtifactResponseBody) *GetPptArtifactResponse
	GetBody() *GetPptArtifactResponseBody
}

type GetPptArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPptArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPptArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPptArtifactResponse) GoString() string {
	return s.String()
}

func (s *GetPptArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPptArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPptArtifactResponse) GetBody() *GetPptArtifactResponseBody {
	return s.Body
}

func (s *GetPptArtifactResponse) SetHeaders(v map[string]*string) *GetPptArtifactResponse {
	s.Headers = v
	return s
}

func (s *GetPptArtifactResponse) SetStatusCode(v int32) *GetPptArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPptArtifactResponse) SetBody(v *GetPptArtifactResponseBody) *GetPptArtifactResponse {
	s.Body = v
	return s
}

func (s *GetPptArtifactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
