// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWuyingServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopWuyingServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopWuyingServerResponse
	GetStatusCode() *int32
	SetBody(v *StopWuyingServerResponseBody) *StopWuyingServerResponse
	GetBody() *StopWuyingServerResponseBody
}

type StopWuyingServerResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopWuyingServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopWuyingServerResponse) String() string {
	return dara.Prettify(s)
}

func (s StopWuyingServerResponse) GoString() string {
	return s.String()
}

func (s *StopWuyingServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopWuyingServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopWuyingServerResponse) GetBody() *StopWuyingServerResponseBody {
	return s.Body
}

func (s *StopWuyingServerResponse) SetHeaders(v map[string]*string) *StopWuyingServerResponse {
	s.Headers = v
	return s
}

func (s *StopWuyingServerResponse) SetStatusCode(v int32) *StopWuyingServerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopWuyingServerResponse) SetBody(v *StopWuyingServerResponseBody) *StopWuyingServerResponse {
	s.Body = v
	return s
}

func (s *StopWuyingServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
