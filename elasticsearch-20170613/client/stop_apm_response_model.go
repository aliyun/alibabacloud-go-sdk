// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopApmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopApmResponse
	GetStatusCode() *int32
	SetBody(v *StopApmResponseBody) *StopApmResponse
	GetBody() *StopApmResponseBody
}

type StopApmResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopApmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopApmResponse) String() string {
	return dara.Prettify(s)
}

func (s StopApmResponse) GoString() string {
	return s.String()
}

func (s *StopApmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopApmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopApmResponse) GetBody() *StopApmResponseBody {
	return s.Body
}

func (s *StopApmResponse) SetHeaders(v map[string]*string) *StopApmResponse {
	s.Headers = v
	return s
}

func (s *StopApmResponse) SetStatusCode(v int32) *StopApmResponse {
	s.StatusCode = &v
	return s
}

func (s *StopApmResponse) SetBody(v *StopApmResponseBody) *StopApmResponse {
	s.Body = v
	return s
}

func (s *StopApmResponse) Validate() error {
	return dara.Validate(s)
}
