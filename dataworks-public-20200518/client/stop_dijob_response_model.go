// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDIJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDIJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDIJobResponse
	GetStatusCode() *int32
	SetBody(v *StopDIJobResponseBody) *StopDIJobResponse
	GetBody() *StopDIJobResponseBody
}

type StopDIJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDIJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDIJobResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDIJobResponse) GoString() string {
	return s.String()
}

func (s *StopDIJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDIJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDIJobResponse) GetBody() *StopDIJobResponseBody {
	return s.Body
}

func (s *StopDIJobResponse) SetHeaders(v map[string]*string) *StopDIJobResponse {
	s.Headers = v
	return s
}

func (s *StopDIJobResponse) SetStatusCode(v int32) *StopDIJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDIJobResponse) SetBody(v *StopDIJobResponseBody) *StopDIJobResponse {
	s.Body = v
	return s
}

func (s *StopDIJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
