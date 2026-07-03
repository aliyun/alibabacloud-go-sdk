// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdnAuthorityVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IdnAuthorityVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IdnAuthorityVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *IdnAuthorityVerifyIntlResponseBody) *IdnAuthorityVerifyIntlResponse
	GetBody() *IdnAuthorityVerifyIntlResponseBody
}

type IdnAuthorityVerifyIntlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IdnAuthorityVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IdnAuthorityVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s IdnAuthorityVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *IdnAuthorityVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IdnAuthorityVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IdnAuthorityVerifyIntlResponse) GetBody() *IdnAuthorityVerifyIntlResponseBody {
	return s.Body
}

func (s *IdnAuthorityVerifyIntlResponse) SetHeaders(v map[string]*string) *IdnAuthorityVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *IdnAuthorityVerifyIntlResponse) SetStatusCode(v int32) *IdnAuthorityVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *IdnAuthorityVerifyIntlResponse) SetBody(v *IdnAuthorityVerifyIntlResponseBody) *IdnAuthorityVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *IdnAuthorityVerifyIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
