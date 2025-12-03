// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelineRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopPipelineRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopPipelineRunResponse
	GetStatusCode() *int32
	SetBody(v *StopPipelineRunResponseBody) *StopPipelineRunResponse
	GetBody() *StopPipelineRunResponseBody
}

type StopPipelineRunResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopPipelineRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopPipelineRunResponse) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineRunResponse) GoString() string {
	return s.String()
}

func (s *StopPipelineRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopPipelineRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopPipelineRunResponse) GetBody() *StopPipelineRunResponseBody {
	return s.Body
}

func (s *StopPipelineRunResponse) SetHeaders(v map[string]*string) *StopPipelineRunResponse {
	s.Headers = v
	return s
}

func (s *StopPipelineRunResponse) SetStatusCode(v int32) *StopPipelineRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPipelineRunResponse) SetBody(v *StopPipelineRunResponseBody) *StopPipelineRunResponse {
	s.Body = v
	return s
}

func (s *StopPipelineRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
