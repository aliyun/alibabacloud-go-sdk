// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptLogstashTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InterruptLogstashTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InterruptLogstashTaskResponse
	GetStatusCode() *int32
	SetBody(v *InterruptLogstashTaskResponseBody) *InterruptLogstashTaskResponse
	GetBody() *InterruptLogstashTaskResponseBody
}

type InterruptLogstashTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InterruptLogstashTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InterruptLogstashTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s InterruptLogstashTaskResponse) GoString() string {
	return s.String()
}

func (s *InterruptLogstashTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InterruptLogstashTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InterruptLogstashTaskResponse) GetBody() *InterruptLogstashTaskResponseBody {
	return s.Body
}

func (s *InterruptLogstashTaskResponse) SetHeaders(v map[string]*string) *InterruptLogstashTaskResponse {
	s.Headers = v
	return s
}

func (s *InterruptLogstashTaskResponse) SetStatusCode(v int32) *InterruptLogstashTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *InterruptLogstashTaskResponse) SetBody(v *InterruptLogstashTaskResponseBody) *InterruptLogstashTaskResponse {
	s.Body = v
	return s
}

func (s *InterruptLogstashTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
