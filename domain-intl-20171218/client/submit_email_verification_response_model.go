// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEmailVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitEmailVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitEmailVerificationResponse
	GetStatusCode() *int32
	SetBody(v *SubmitEmailVerificationResponseBody) *SubmitEmailVerificationResponse
	GetBody() *SubmitEmailVerificationResponseBody
}

type SubmitEmailVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitEmailVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *SubmitEmailVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitEmailVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitEmailVerificationResponse) GetBody() *SubmitEmailVerificationResponseBody {
	return s.Body
}

func (s *SubmitEmailVerificationResponse) SetHeaders(v map[string]*string) *SubmitEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *SubmitEmailVerificationResponse) SetStatusCode(v int32) *SubmitEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitEmailVerificationResponse) SetBody(v *SubmitEmailVerificationResponseBody) *SubmitEmailVerificationResponse {
	s.Body = v
	return s
}

func (s *SubmitEmailVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
