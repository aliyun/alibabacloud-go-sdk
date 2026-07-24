// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopComputeJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopComputeJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopComputeJobResponse
	GetStatusCode() *int32
	SetBody(v *StopComputeJobResponseBody) *StopComputeJobResponse
	GetBody() *StopComputeJobResponseBody
}

type StopComputeJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopComputeJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopComputeJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopComputeJobResponse) GoString() string {
	return s.String()
}

func (s *StopComputeJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopComputeJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopComputeJobResponse) GetBody() *StopComputeJobResponseBody {
	return s.Body
}

func (s *StopComputeJobResponse) SetHeaders(v map[string]*string) *StopComputeJobResponse {
	s.Headers = v
	return s
}

func (s *StopComputeJobResponse) SetStatusCode(v int32) *StopComputeJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopComputeJobResponse) SetBody(v *StopComputeJobResponseBody) *StopComputeJobResponse {
	s.Body = v
	return s
}

func (s *StopComputeJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
