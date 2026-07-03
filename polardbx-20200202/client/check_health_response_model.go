// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHealthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckHealthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckHealthResponse
	GetStatusCode() *int32
	SetBody(v *CheckHealthResponseBody) *CheckHealthResponse
	GetBody() *CheckHealthResponseBody
}

type CheckHealthResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckHealthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckHealthResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckHealthResponse) GoString() string {
	return s.String()
}

func (s *CheckHealthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckHealthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckHealthResponse) GetBody() *CheckHealthResponseBody {
	return s.Body
}

func (s *CheckHealthResponse) SetHeaders(v map[string]*string) *CheckHealthResponse {
	s.Headers = v
	return s
}

func (s *CheckHealthResponse) SetStatusCode(v int32) *CheckHealthResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckHealthResponse) SetBody(v *CheckHealthResponseBody) *CheckHealthResponse {
	s.Body = v
	return s
}

func (s *CheckHealthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
