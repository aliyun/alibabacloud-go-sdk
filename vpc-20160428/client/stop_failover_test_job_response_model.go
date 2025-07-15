// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopFailoverTestJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopFailoverTestJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopFailoverTestJobResponse
	GetStatusCode() *int32
	SetBody(v *StopFailoverTestJobResponseBody) *StopFailoverTestJobResponse
	GetBody() *StopFailoverTestJobResponseBody
}

type StopFailoverTestJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopFailoverTestJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopFailoverTestJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopFailoverTestJobResponse) GoString() string {
	return s.String()
}

func (s *StopFailoverTestJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopFailoverTestJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopFailoverTestJobResponse) GetBody() *StopFailoverTestJobResponseBody {
	return s.Body
}

func (s *StopFailoverTestJobResponse) SetHeaders(v map[string]*string) *StopFailoverTestJobResponse {
	s.Headers = v
	return s
}

func (s *StopFailoverTestJobResponse) SetStatusCode(v int32) *StopFailoverTestJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopFailoverTestJobResponse) SetBody(v *StopFailoverTestJobResponseBody) *StopFailoverTestJobResponse {
	s.Body = v
	return s
}

func (s *StopFailoverTestJobResponse) Validate() error {
	return dara.Validate(s)
}
