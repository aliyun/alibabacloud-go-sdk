// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyEmailResponse
	GetStatusCode() *int32
	SetBody(v *VerifyEmailResponseBody) *VerifyEmailResponse
	GetBody() *VerifyEmailResponseBody
}

type VerifyEmailResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyEmailResponse) GoString() string {
	return s.String()
}

func (s *VerifyEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyEmailResponse) GetBody() *VerifyEmailResponseBody {
	return s.Body
}

func (s *VerifyEmailResponse) SetHeaders(v map[string]*string) *VerifyEmailResponse {
	s.Headers = v
	return s
}

func (s *VerifyEmailResponse) SetStatusCode(v int32) *VerifyEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyEmailResponse) SetBody(v *VerifyEmailResponseBody) *VerifyEmailResponse {
	s.Body = v
	return s
}

func (s *VerifyEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
