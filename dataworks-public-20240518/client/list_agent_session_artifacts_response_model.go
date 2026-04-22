// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionArtifactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentSessionArtifactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentSessionArtifactsResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentSessionArtifactsResponseBody) *ListAgentSessionArtifactsResponse
	GetBody() *ListAgentSessionArtifactsResponseBody
}

type ListAgentSessionArtifactsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentSessionArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentSessionArtifactsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionArtifactsResponse) GoString() string {
	return s.String()
}

func (s *ListAgentSessionArtifactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentSessionArtifactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentSessionArtifactsResponse) GetBody() *ListAgentSessionArtifactsResponseBody {
	return s.Body
}

func (s *ListAgentSessionArtifactsResponse) SetHeaders(v map[string]*string) *ListAgentSessionArtifactsResponse {
	s.Headers = v
	return s
}

func (s *ListAgentSessionArtifactsResponse) SetStatusCode(v int32) *ListAgentSessionArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentSessionArtifactsResponse) SetBody(v *ListAgentSessionArtifactsResponseBody) *ListAgentSessionArtifactsResponse {
	s.Body = v
	return s
}

func (s *ListAgentSessionArtifactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
