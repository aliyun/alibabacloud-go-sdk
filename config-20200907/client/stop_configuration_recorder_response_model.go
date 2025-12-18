// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopConfigurationRecorderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopConfigurationRecorderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopConfigurationRecorderResponse
	GetStatusCode() *int32
	SetBody(v *StopConfigurationRecorderResponseBody) *StopConfigurationRecorderResponse
	GetBody() *StopConfigurationRecorderResponseBody
}

type StopConfigurationRecorderResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopConfigurationRecorderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopConfigurationRecorderResponse) String() string {
	return dara.Prettify(s)
}

func (s StopConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *StopConfigurationRecorderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopConfigurationRecorderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopConfigurationRecorderResponse) GetBody() *StopConfigurationRecorderResponseBody {
	return s.Body
}

func (s *StopConfigurationRecorderResponse) SetHeaders(v map[string]*string) *StopConfigurationRecorderResponse {
	s.Headers = v
	return s
}

func (s *StopConfigurationRecorderResponse) SetStatusCode(v int32) *StopConfigurationRecorderResponse {
	s.StatusCode = &v
	return s
}

func (s *StopConfigurationRecorderResponse) SetBody(v *StopConfigurationRecorderResponseBody) *StopConfigurationRecorderResponse {
	s.Body = v
	return s
}

func (s *StopConfigurationRecorderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
