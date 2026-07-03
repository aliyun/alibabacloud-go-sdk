// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataAgentAccuracyTestTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDataAgentAccuracyTestTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDataAgentAccuracyTestTaskResponse
	GetStatusCode() *int32
	SetBody(v *StopDataAgentAccuracyTestTaskResponseBody) *StopDataAgentAccuracyTestTaskResponse
	GetBody() *StopDataAgentAccuracyTestTaskResponseBody
}

type StopDataAgentAccuracyTestTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDataAgentAccuracyTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDataAgentAccuracyTestTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDataAgentAccuracyTestTaskResponse) GoString() string {
	return s.String()
}

func (s *StopDataAgentAccuracyTestTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDataAgentAccuracyTestTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDataAgentAccuracyTestTaskResponse) GetBody() *StopDataAgentAccuracyTestTaskResponseBody {
	return s.Body
}

func (s *StopDataAgentAccuracyTestTaskResponse) SetHeaders(v map[string]*string) *StopDataAgentAccuracyTestTaskResponse {
	s.Headers = v
	return s
}

func (s *StopDataAgentAccuracyTestTaskResponse) SetStatusCode(v int32) *StopDataAgentAccuracyTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDataAgentAccuracyTestTaskResponse) SetBody(v *StopDataAgentAccuracyTestTaskResponseBody) *StopDataAgentAccuracyTestTaskResponse {
	s.Body = v
	return s
}

func (s *StopDataAgentAccuracyTestTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
