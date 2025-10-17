// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Mobile3MetaVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Mobile3MetaVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *Mobile3MetaVerifyIntlResponseBody) *Mobile3MetaVerifyIntlResponse
	GetBody() *Mobile3MetaVerifyIntlResponseBody
}

type Mobile3MetaVerifyIntlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Mobile3MetaVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Mobile3MetaVerifyIntlResponse) GetBody() *Mobile3MetaVerifyIntlResponseBody {
	return s.Body
}

func (s *Mobile3MetaVerifyIntlResponse) SetHeaders(v map[string]*string) *Mobile3MetaVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaVerifyIntlResponse) SetStatusCode(v int32) *Mobile3MetaVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaVerifyIntlResponse) SetBody(v *Mobile3MetaVerifyIntlResponseBody) *Mobile3MetaVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *Mobile3MetaVerifyIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
