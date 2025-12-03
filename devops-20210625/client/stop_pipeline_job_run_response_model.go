// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelineJobRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopPipelineJobRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopPipelineJobRunResponse
	GetStatusCode() *int32
	SetBody(v *StopPipelineJobRunResponseBody) *StopPipelineJobRunResponse
	GetBody() *StopPipelineJobRunResponseBody
}

type StopPipelineJobRunResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopPipelineJobRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopPipelineJobRunResponse) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineJobRunResponse) GoString() string {
	return s.String()
}

func (s *StopPipelineJobRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopPipelineJobRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopPipelineJobRunResponse) GetBody() *StopPipelineJobRunResponseBody {
	return s.Body
}

func (s *StopPipelineJobRunResponse) SetHeaders(v map[string]*string) *StopPipelineJobRunResponse {
	s.Headers = v
	return s
}

func (s *StopPipelineJobRunResponse) SetStatusCode(v int32) *StopPipelineJobRunResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPipelineJobRunResponse) SetBody(v *StopPipelineJobRunResponseBody) *StopPipelineJobRunResponse {
	s.Body = v
	return s
}

func (s *StopPipelineJobRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
