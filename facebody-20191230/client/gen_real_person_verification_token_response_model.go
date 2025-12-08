// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenRealPersonVerificationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenRealPersonVerificationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenRealPersonVerificationTokenResponse
	GetStatusCode() *int32
	SetBody(v *GenRealPersonVerificationTokenResponseBody) *GenRealPersonVerificationTokenResponse
	GetBody() *GenRealPersonVerificationTokenResponseBody
}

type GenRealPersonVerificationTokenResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenRealPersonVerificationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenRealPersonVerificationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GenRealPersonVerificationTokenResponse) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenRealPersonVerificationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenRealPersonVerificationTokenResponse) GetBody() *GenRealPersonVerificationTokenResponseBody {
	return s.Body
}

func (s *GenRealPersonVerificationTokenResponse) SetHeaders(v map[string]*string) *GenRealPersonVerificationTokenResponse {
	s.Headers = v
	return s
}

func (s *GenRealPersonVerificationTokenResponse) SetStatusCode(v int32) *GenRealPersonVerificationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GenRealPersonVerificationTokenResponse) SetBody(v *GenRealPersonVerificationTokenResponseBody) *GenRealPersonVerificationTokenResponse {
	s.Body = v
	return s
}

func (s *GenRealPersonVerificationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
