// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPipelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunPipelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunPipelineResponse
	GetStatusCode() *int32
	SetBody(v *RunPipelineResponseBody) *RunPipelineResponse
	GetBody() *RunPipelineResponseBody
}

type RunPipelineResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunPipelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPipelineResponse) String() string {
	return dara.Prettify(s)
}

func (s RunPipelineResponse) GoString() string {
	return s.String()
}

func (s *RunPipelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunPipelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunPipelineResponse) GetBody() *RunPipelineResponseBody {
	return s.Body
}

func (s *RunPipelineResponse) SetHeaders(v map[string]*string) *RunPipelineResponse {
	s.Headers = v
	return s
}

func (s *RunPipelineResponse) SetStatusCode(v int32) *RunPipelineResponse {
	s.StatusCode = &v
	return s
}

func (s *RunPipelineResponse) SetBody(v *RunPipelineResponseBody) *RunPipelineResponse {
	s.Body = v
	return s
}

func (s *RunPipelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
