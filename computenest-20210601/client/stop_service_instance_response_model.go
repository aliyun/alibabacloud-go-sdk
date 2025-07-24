// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StopServiceInstanceResponseBody) *StopServiceInstanceResponse
	GetBody() *StopServiceInstanceResponseBody
}

type StopServiceInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StopServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopServiceInstanceResponse) GetBody() *StopServiceInstanceResponseBody {
	return s.Body
}

func (s *StopServiceInstanceResponse) SetHeaders(v map[string]*string) *StopServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceInstanceResponse) SetStatusCode(v int32) *StopServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceInstanceResponse) SetBody(v *StopServiceInstanceResponseBody) *StopServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *StopServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
