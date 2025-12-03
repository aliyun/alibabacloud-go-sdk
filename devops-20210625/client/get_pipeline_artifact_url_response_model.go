// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineArtifactUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineArtifactUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineArtifactUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineArtifactUrlResponseBody) *GetPipelineArtifactUrlResponse
	GetBody() *GetPipelineArtifactUrlResponseBody
}

type GetPipelineArtifactUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineArtifactUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineArtifactUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineArtifactUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineArtifactUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineArtifactUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineArtifactUrlResponse) GetBody() *GetPipelineArtifactUrlResponseBody {
	return s.Body
}

func (s *GetPipelineArtifactUrlResponse) SetHeaders(v map[string]*string) *GetPipelineArtifactUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineArtifactUrlResponse) SetStatusCode(v int32) *GetPipelineArtifactUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineArtifactUrlResponse) SetBody(v *GetPipelineArtifactUrlResponseBody) *GetPipelineArtifactUrlResponse {
	s.Body = v
	return s
}

func (s *GetPipelineArtifactUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
