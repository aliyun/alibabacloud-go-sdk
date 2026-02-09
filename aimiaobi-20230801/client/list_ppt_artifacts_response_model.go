// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPptArtifactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPptArtifactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPptArtifactsResponse
	GetStatusCode() *int32
	SetBody(v *ListPptArtifactsResponseBody) *ListPptArtifactsResponse
	GetBody() *ListPptArtifactsResponseBody
}

type ListPptArtifactsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPptArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPptArtifactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPptArtifactsResponse) GoString() string {
	return s.String()
}

func (s *ListPptArtifactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPptArtifactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPptArtifactsResponse) GetBody() *ListPptArtifactsResponseBody {
	return s.Body
}

func (s *ListPptArtifactsResponse) SetHeaders(v map[string]*string) *ListPptArtifactsResponse {
	s.Headers = v
	return s
}

func (s *ListPptArtifactsResponse) SetStatusCode(v int32) *ListPptArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPptArtifactsResponse) SetBody(v *ListPptArtifactsResponseBody) *ListPptArtifactsResponse {
	s.Body = v
	return s
}

func (s *ListPptArtifactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
