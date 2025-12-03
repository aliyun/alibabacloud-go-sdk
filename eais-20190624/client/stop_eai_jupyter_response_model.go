// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEaiJupyterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopEaiJupyterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopEaiJupyterResponse
	GetStatusCode() *int32
	SetBody(v *StopEaiJupyterResponseBody) *StopEaiJupyterResponse
	GetBody() *StopEaiJupyterResponseBody
}

type StopEaiJupyterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopEaiJupyterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopEaiJupyterResponse) String() string {
	return dara.Prettify(s)
}

func (s StopEaiJupyterResponse) GoString() string {
	return s.String()
}

func (s *StopEaiJupyterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopEaiJupyterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopEaiJupyterResponse) GetBody() *StopEaiJupyterResponseBody {
	return s.Body
}

func (s *StopEaiJupyterResponse) SetHeaders(v map[string]*string) *StopEaiJupyterResponse {
	s.Headers = v
	return s
}

func (s *StopEaiJupyterResponse) SetStatusCode(v int32) *StopEaiJupyterResponse {
	s.StatusCode = &v
	return s
}

func (s *StopEaiJupyterResponse) SetBody(v *StopEaiJupyterResponseBody) *StopEaiJupyterResponse {
	s.Body = v
	return s
}

func (s *StopEaiJupyterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
