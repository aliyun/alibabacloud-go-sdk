// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTestingJMeterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTestingJMeterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTestingJMeterSceneResponse
	GetStatusCode() *int32
	SetBody(v *StopTestingJMeterSceneResponseBody) *StopTestingJMeterSceneResponse
	GetBody() *StopTestingJMeterSceneResponseBody
}

type StopTestingJMeterSceneResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTestingJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTestingJMeterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTestingJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *StopTestingJMeterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTestingJMeterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTestingJMeterSceneResponse) GetBody() *StopTestingJMeterSceneResponseBody {
	return s.Body
}

func (s *StopTestingJMeterSceneResponse) SetHeaders(v map[string]*string) *StopTestingJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *StopTestingJMeterSceneResponse) SetStatusCode(v int32) *StopTestingJMeterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTestingJMeterSceneResponse) SetBody(v *StopTestingJMeterSceneResponseBody) *StopTestingJMeterSceneResponse {
	s.Body = v
	return s
}

func (s *StopTestingJMeterSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
