// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDebuggingJMeterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDebuggingJMeterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDebuggingJMeterSceneResponse
	GetStatusCode() *int32
	SetBody(v *StopDebuggingJMeterSceneResponseBody) *StopDebuggingJMeterSceneResponse
	GetBody() *StopDebuggingJMeterSceneResponseBody
}

type StopDebuggingJMeterSceneResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDebuggingJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDebuggingJMeterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDebuggingJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *StopDebuggingJMeterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDebuggingJMeterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDebuggingJMeterSceneResponse) GetBody() *StopDebuggingJMeterSceneResponseBody {
	return s.Body
}

func (s *StopDebuggingJMeterSceneResponse) SetHeaders(v map[string]*string) *StopDebuggingJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *StopDebuggingJMeterSceneResponse) SetStatusCode(v int32) *StopDebuggingJMeterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDebuggingJMeterSceneResponse) SetBody(v *StopDebuggingJMeterSceneResponseBody) *StopDebuggingJMeterSceneResponse {
	s.Body = v
	return s
}

func (s *StopDebuggingJMeterSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
