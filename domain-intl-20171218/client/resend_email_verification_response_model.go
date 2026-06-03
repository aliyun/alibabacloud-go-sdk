// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendEmailVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResendEmailVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResendEmailVerificationResponse
	GetStatusCode() *int32
	SetBody(v *ResendEmailVerificationResponseBody) *ResendEmailVerificationResponse
	GetBody() *ResendEmailVerificationResponseBody
}

type ResendEmailVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendEmailVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s ResendEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *ResendEmailVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResendEmailVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResendEmailVerificationResponse) GetBody() *ResendEmailVerificationResponseBody {
	return s.Body
}

func (s *ResendEmailVerificationResponse) SetHeaders(v map[string]*string) *ResendEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *ResendEmailVerificationResponse) SetStatusCode(v int32) *ResendEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendEmailVerificationResponse) SetBody(v *ResendEmailVerificationResponseBody) *ResendEmailVerificationResponse {
	s.Body = v
	return s
}

func (s *ResendEmailVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
