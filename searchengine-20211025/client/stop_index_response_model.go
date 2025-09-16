// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopIndexResponse
	GetStatusCode() *int32
	SetBody(v *StopIndexResponseBody) *StopIndexResponse
	GetBody() *StopIndexResponseBody
}

type StopIndexResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s StopIndexResponse) GoString() string {
	return s.String()
}

func (s *StopIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopIndexResponse) GetBody() *StopIndexResponseBody {
	return s.Body
}

func (s *StopIndexResponse) SetHeaders(v map[string]*string) *StopIndexResponse {
	s.Headers = v
	return s
}

func (s *StopIndexResponse) SetStatusCode(v int32) *StopIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *StopIndexResponse) SetBody(v *StopIndexResponseBody) *StopIndexResponse {
	s.Body = v
	return s
}

func (s *StopIndexResponse) Validate() error {
	return dara.Validate(s)
}
