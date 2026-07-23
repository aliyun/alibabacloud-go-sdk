// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptForArbitrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InterruptForArbitrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InterruptForArbitrationResponse
	GetStatusCode() *int32
	SetBody(v *InterruptForArbitrationResponseBody) *InterruptForArbitrationResponse
	GetBody() *InterruptForArbitrationResponseBody
}

type InterruptForArbitrationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InterruptForArbitrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InterruptForArbitrationResponse) String() string {
	return dara.Prettify(s)
}

func (s InterruptForArbitrationResponse) GoString() string {
	return s.String()
}

func (s *InterruptForArbitrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InterruptForArbitrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InterruptForArbitrationResponse) GetBody() *InterruptForArbitrationResponseBody {
	return s.Body
}

func (s *InterruptForArbitrationResponse) SetHeaders(v map[string]*string) *InterruptForArbitrationResponse {
	s.Headers = v
	return s
}

func (s *InterruptForArbitrationResponse) SetStatusCode(v int32) *InterruptForArbitrationResponse {
	s.StatusCode = &v
	return s
}

func (s *InterruptForArbitrationResponse) SetBody(v *InterruptForArbitrationResponseBody) *InterruptForArbitrationResponse {
	s.Body = v
	return s
}

func (s *InterruptForArbitrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
