// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEmailVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEmailVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEmailVerificationResponse
	GetStatusCode() *int32
	SetBody(v *ListEmailVerificationResponseBody) *ListEmailVerificationResponse
	GetBody() *ListEmailVerificationResponseBody
}

type ListEmailVerificationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEmailVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *ListEmailVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEmailVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEmailVerificationResponse) GetBody() *ListEmailVerificationResponseBody {
	return s.Body
}

func (s *ListEmailVerificationResponse) SetHeaders(v map[string]*string) *ListEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *ListEmailVerificationResponse) SetStatusCode(v int32) *ListEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEmailVerificationResponse) SetBody(v *ListEmailVerificationResponseBody) *ListEmailVerificationResponse {
	s.Body = v
	return s
}

func (s *ListEmailVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
