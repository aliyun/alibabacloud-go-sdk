// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopApplicationResponse
	GetStatusCode() *int32
	SetBody(v *StopApplicationResponseBody) *StopApplicationResponse
	GetBody() *StopApplicationResponseBody
}

type StopApplicationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s StopApplicationResponse) GoString() string {
	return s.String()
}

func (s *StopApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopApplicationResponse) GetBody() *StopApplicationResponseBody {
	return s.Body
}

func (s *StopApplicationResponse) SetHeaders(v map[string]*string) *StopApplicationResponse {
	s.Headers = v
	return s
}

func (s *StopApplicationResponse) SetStatusCode(v int32) *StopApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopApplicationResponse) SetBody(v *StopApplicationResponseBody) *StopApplicationResponse {
	s.Body = v
	return s
}

func (s *StopApplicationResponse) Validate() error {
	return dara.Validate(s)
}
