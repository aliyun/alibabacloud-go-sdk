// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComponentsVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckComponentsVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckComponentsVersionResponse
	GetStatusCode() *int32
	SetBody(v *CheckComponentsVersionResponseBody) *CheckComponentsVersionResponse
	GetBody() *CheckComponentsVersionResponseBody
}

type CheckComponentsVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckComponentsVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckComponentsVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckComponentsVersionResponse) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckComponentsVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckComponentsVersionResponse) GetBody() *CheckComponentsVersionResponseBody {
	return s.Body
}

func (s *CheckComponentsVersionResponse) SetHeaders(v map[string]*string) *CheckComponentsVersionResponse {
	s.Headers = v
	return s
}

func (s *CheckComponentsVersionResponse) SetStatusCode(v int32) *CheckComponentsVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckComponentsVersionResponse) SetBody(v *CheckComponentsVersionResponseBody) *CheckComponentsVersionResponse {
	s.Body = v
	return s
}

func (s *CheckComponentsVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
