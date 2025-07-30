// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyCredentialResponse
	GetStatusCode() *int32
	SetBody(v *VerifyCredentialResponseBody) *VerifyCredentialResponse
	GetBody() *VerifyCredentialResponseBody
}

type VerifyCredentialResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyCredentialResponse) GoString() string {
	return s.String()
}

func (s *VerifyCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyCredentialResponse) GetBody() *VerifyCredentialResponseBody {
	return s.Body
}

func (s *VerifyCredentialResponse) SetHeaders(v map[string]*string) *VerifyCredentialResponse {
	s.Headers = v
	return s
}

func (s *VerifyCredentialResponse) SetStatusCode(v int32) *VerifyCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCredentialResponse) SetBody(v *VerifyCredentialResponseBody) *VerifyCredentialResponse {
	s.Body = v
	return s
}

func (s *VerifyCredentialResponse) Validate() error {
	return dara.Validate(s)
}
