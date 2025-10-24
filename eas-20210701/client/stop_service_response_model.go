// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopServiceResponse
	GetStatusCode() *int32
	SetBody(v *StopServiceResponseBody) *StopServiceResponse
	GetBody() *StopServiceResponseBody
}

type StopServiceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopServiceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopServiceResponse) GetBody() *StopServiceResponseBody {
	return s.Body
}

func (s *StopServiceResponse) SetHeaders(v map[string]*string) *StopServiceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceResponse) SetStatusCode(v int32) *StopServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceResponse) SetBody(v *StopServiceResponseBody) *StopServiceResponse {
	s.Body = v
	return s
}

func (s *StopServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
