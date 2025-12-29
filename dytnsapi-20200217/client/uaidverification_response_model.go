// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UAIDVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UAIDVerificationResponse
	GetStatusCode() *int32
	SetBody(v *UAIDVerificationResponseBody) *UAIDVerificationResponse
	GetBody() *UAIDVerificationResponseBody
}

type UAIDVerificationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UAIDVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UAIDVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s UAIDVerificationResponse) GoString() string {
	return s.String()
}

func (s *UAIDVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UAIDVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UAIDVerificationResponse) GetBody() *UAIDVerificationResponseBody {
	return s.Body
}

func (s *UAIDVerificationResponse) SetHeaders(v map[string]*string) *UAIDVerificationResponse {
	s.Headers = v
	return s
}

func (s *UAIDVerificationResponse) SetStatusCode(v int32) *UAIDVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *UAIDVerificationResponse) SetBody(v *UAIDVerificationResponseBody) *UAIDVerificationResponse {
	s.Body = v
	return s
}

func (s *UAIDVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
