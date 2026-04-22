// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionArtifactMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentSessionArtifactMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentSessionArtifactMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentSessionArtifactMetaResponseBody) *GetAgentSessionArtifactMetaResponse
	GetBody() *GetAgentSessionArtifactMetaResponseBody
}

type GetAgentSessionArtifactMetaResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentSessionArtifactMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentSessionArtifactMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionArtifactMetaResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSessionArtifactMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentSessionArtifactMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentSessionArtifactMetaResponse) GetBody() *GetAgentSessionArtifactMetaResponseBody {
	return s.Body
}

func (s *GetAgentSessionArtifactMetaResponse) SetHeaders(v map[string]*string) *GetAgentSessionArtifactMetaResponse {
	s.Headers = v
	return s
}

func (s *GetAgentSessionArtifactMetaResponse) SetStatusCode(v int32) *GetAgentSessionArtifactMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentSessionArtifactMetaResponse) SetBody(v *GetAgentSessionArtifactMetaResponseBody) *GetAgentSessionArtifactMetaResponse {
	s.Body = v
	return s
}

func (s *GetAgentSessionArtifactMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
