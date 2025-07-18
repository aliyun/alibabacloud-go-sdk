// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile2MetaVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Mobile2MetaVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Mobile2MetaVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *Mobile2MetaVerifyIntlResponseBody) *Mobile2MetaVerifyIntlResponse
	GetBody() *Mobile2MetaVerifyIntlResponseBody
}

type Mobile2MetaVerifyIntlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile2MetaVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile2MetaVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s Mobile2MetaVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Mobile2MetaVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Mobile2MetaVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Mobile2MetaVerifyIntlResponse) GetBody() *Mobile2MetaVerifyIntlResponseBody {
	return s.Body
}

func (s *Mobile2MetaVerifyIntlResponse) SetHeaders(v map[string]*string) *Mobile2MetaVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Mobile2MetaVerifyIntlResponse) SetStatusCode(v int32) *Mobile2MetaVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile2MetaVerifyIntlResponse) SetBody(v *Mobile2MetaVerifyIntlResponseBody) *Mobile2MetaVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *Mobile2MetaVerifyIntlResponse) Validate() error {
	return dara.Validate(s)
}
