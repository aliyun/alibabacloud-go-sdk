// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySmsCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifySmsCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifySmsCodeResponse
	GetStatusCode() *int32
	SetBody(v *VerifySmsCodeResponseBody) *VerifySmsCodeResponse
	GetBody() *VerifySmsCodeResponseBody
}

type VerifySmsCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifySmsCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifySmsCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifySmsCodeResponse) GoString() string {
	return s.String()
}

func (s *VerifySmsCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifySmsCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifySmsCodeResponse) GetBody() *VerifySmsCodeResponseBody {
	return s.Body
}

func (s *VerifySmsCodeResponse) SetHeaders(v map[string]*string) *VerifySmsCodeResponse {
	s.Headers = v
	return s
}

func (s *VerifySmsCodeResponse) SetStatusCode(v int32) *VerifySmsCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySmsCodeResponse) SetBody(v *VerifySmsCodeResponseBody) *VerifySmsCodeResponse {
	s.Body = v
	return s
}

func (s *VerifySmsCodeResponse) Validate() error {
	return dara.Validate(s)
}
