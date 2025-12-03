// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineEmasArtifactUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPipelineEmasArtifactUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPipelineEmasArtifactUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetPipelineEmasArtifactUrlResponseBody) *GetPipelineEmasArtifactUrlResponse
	GetBody() *GetPipelineEmasArtifactUrlResponseBody
}

type GetPipelineEmasArtifactUrlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPipelineEmasArtifactUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPipelineEmasArtifactUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineEmasArtifactUrlResponse) GoString() string {
	return s.String()
}

func (s *GetPipelineEmasArtifactUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPipelineEmasArtifactUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPipelineEmasArtifactUrlResponse) GetBody() *GetPipelineEmasArtifactUrlResponseBody {
	return s.Body
}

func (s *GetPipelineEmasArtifactUrlResponse) SetHeaders(v map[string]*string) *GetPipelineEmasArtifactUrlResponse {
	s.Headers = v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponse) SetStatusCode(v int32) *GetPipelineEmasArtifactUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponse) SetBody(v *GetPipelineEmasArtifactUrlResponseBody) *GetPipelineEmasArtifactUrlResponse {
	s.Body = v
	return s
}

func (s *GetPipelineEmasArtifactUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
