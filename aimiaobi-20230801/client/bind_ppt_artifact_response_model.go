// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPptArtifactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindPptArtifactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindPptArtifactResponse
	GetStatusCode() *int32
	SetBody(v *BindPptArtifactResponseBody) *BindPptArtifactResponse
	GetBody() *BindPptArtifactResponseBody
}

type BindPptArtifactResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindPptArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindPptArtifactResponse) String() string {
	return dara.Prettify(s)
}

func (s BindPptArtifactResponse) GoString() string {
	return s.String()
}

func (s *BindPptArtifactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindPptArtifactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindPptArtifactResponse) GetBody() *BindPptArtifactResponseBody {
	return s.Body
}

func (s *BindPptArtifactResponse) SetHeaders(v map[string]*string) *BindPptArtifactResponse {
	s.Headers = v
	return s
}

func (s *BindPptArtifactResponse) SetStatusCode(v int32) *BindPptArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *BindPptArtifactResponse) SetBody(v *BindPptArtifactResponseBody) *BindPptArtifactResponse {
	s.Body = v
	return s
}

func (s *BindPptArtifactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
