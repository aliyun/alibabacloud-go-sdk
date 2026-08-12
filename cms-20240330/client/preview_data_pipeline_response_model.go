// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDataPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreviewDataPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreviewDataPipelineResponse
	GetStatusCode() *int32
	SetBody(v *PreviewDataPipelineResponseBody) *PreviewDataPipelineResponse
	GetBody() *PreviewDataPipelineResponseBody
}

type PreviewDataPipelineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreviewDataPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreviewDataPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s PreviewDataPipelineResponse) GoString() string {
	return s.String()
}

func (s *PreviewDataPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreviewDataPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreviewDataPipelineResponse) GetBody() *PreviewDataPipelineResponseBody {
	return s.Body
}

func (s *PreviewDataPipelineResponse) SetHeaders(v map[string]*string) *PreviewDataPipelineResponse {
	s.Headers = v
	return s
}

func (s *PreviewDataPipelineResponse) SetStatusCode(v int32) *PreviewDataPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewDataPipelineResponse) SetBody(v *PreviewDataPipelineResponseBody) *PreviewDataPipelineResponse {
	s.Body = v
	return s
}

func (s *PreviewDataPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
