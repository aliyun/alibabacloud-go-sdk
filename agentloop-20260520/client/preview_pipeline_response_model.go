// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewPipelineResponse
	GetStatusCode() *int32
	SetBody(v *PreviewPipelineResponseBody) *PreviewPipelineResponse
	GetBody() *PreviewPipelineResponseBody
}

type PreviewPipelineResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewPipelineResponse) GoString() string {
	return s.String()
}

func (s *PreviewPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewPipelineResponse) GetBody() *PreviewPipelineResponseBody {
	return s.Body
}

func (s *PreviewPipelineResponse) SetHeaders(v map[string]*string) *PreviewPipelineResponse {
	s.Headers = v
	return s
}

func (s *PreviewPipelineResponse) SetStatusCode(v int32) *PreviewPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewPipelineResponse) SetBody(v *PreviewPipelineResponseBody) *PreviewPipelineResponse {
	s.Body = v
	return s
}

func (s *PreviewPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
