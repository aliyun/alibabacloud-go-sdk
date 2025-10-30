// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSimulationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopSimulationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopSimulationTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopSimulationTaskResponseBody) *StopSimulationTaskResponse
	GetBody() *StopSimulationTaskResponseBody
}

type StopSimulationTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopSimulationTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopSimulationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopSimulationTaskResponse) GoString() string {
	return s.String()
}

func (s *StopSimulationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopSimulationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopSimulationTaskResponse) GetBody() *StopSimulationTaskResponseBody {
	return s.Body
}

func (s *StopSimulationTaskResponse) SetHeaders(v map[string]*string) *StopSimulationTaskResponse {
	s.Headers = v
	return s
}

func (s *StopSimulationTaskResponse) SetStatusCode(v int32) *StopSimulationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopSimulationTaskResponse) SetBody(v *StopSimulationTaskResponseBody) *StopSimulationTaskResponse {
	s.Body = v
	return s
}

func (s *StopSimulationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
