// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPipelineIntegratedTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopPipelineIntegratedTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopPipelineIntegratedTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopPipelineIntegratedTaskResponseBody) *StopPipelineIntegratedTaskResponse
	GetBody() *StopPipelineIntegratedTaskResponseBody
}

type StopPipelineIntegratedTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopPipelineIntegratedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopPipelineIntegratedTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopPipelineIntegratedTaskResponse) GoString() string {
	return s.String()
}

func (s *StopPipelineIntegratedTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopPipelineIntegratedTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopPipelineIntegratedTaskResponse) GetBody() *StopPipelineIntegratedTaskResponseBody {
	return s.Body
}

func (s *StopPipelineIntegratedTaskResponse) SetHeaders(v map[string]*string) *StopPipelineIntegratedTaskResponse {
	s.Headers = v
	return s
}

func (s *StopPipelineIntegratedTaskResponse) SetStatusCode(v int32) *StopPipelineIntegratedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPipelineIntegratedTaskResponse) SetBody(v *StopPipelineIntegratedTaskResponseBody) *StopPipelineIntegratedTaskResponse {
	s.Body = v
	return s
}

func (s *StopPipelineIntegratedTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
