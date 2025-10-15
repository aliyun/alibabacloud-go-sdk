// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopAppResponse
	GetStatusCode() *int32
	SetBody(v *StopAppResponseBody) *StopAppResponse
	GetBody() *StopAppResponseBody
}

type StopAppResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopAppResponse) String() string {
	return dara.Prettify(s)
}

func (s StopAppResponse) GoString() string {
	return s.String()
}

func (s *StopAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopAppResponse) GetBody() *StopAppResponseBody {
	return s.Body
}

func (s *StopAppResponse) SetHeaders(v map[string]*string) *StopAppResponse {
	s.Headers = v
	return s
}

func (s *StopAppResponse) SetStatusCode(v int32) *StopAppResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAppResponse) SetBody(v *StopAppResponseBody) *StopAppResponse {
	s.Body = v
	return s
}

func (s *StopAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
