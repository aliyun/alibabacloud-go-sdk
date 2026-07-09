// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumePipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumePipelineResponse
	GetStatusCode() *int32
	SetBody(v *ResumePipelineResponseBody) *ResumePipelineResponse
	GetBody() *ResumePipelineResponseBody
}

type ResumePipelineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumePipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumePipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumePipelineResponse) GoString() string {
	return s.String()
}

func (s *ResumePipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumePipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumePipelineResponse) GetBody() *ResumePipelineResponseBody {
	return s.Body
}

func (s *ResumePipelineResponse) SetHeaders(v map[string]*string) *ResumePipelineResponse {
	s.Headers = v
	return s
}

func (s *ResumePipelineResponse) SetStatusCode(v int32) *ResumePipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumePipelineResponse) SetBody(v *ResumePipelineResponseBody) *ResumePipelineResponse {
	s.Body = v
	return s
}

func (s *ResumePipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
