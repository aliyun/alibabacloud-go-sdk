// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id2MetaVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id2MetaVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *Id2MetaVerifyIntlResponseBody) *Id2MetaVerifyIntlResponse
	GetBody() *Id2MetaVerifyIntlResponseBody
}

type Id2MetaVerifyIntlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id2MetaVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id2MetaVerifyIntlResponse) GetBody() *Id2MetaVerifyIntlResponseBody {
	return s.Body
}

func (s *Id2MetaVerifyIntlResponse) SetHeaders(v map[string]*string) *Id2MetaVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaVerifyIntlResponse) SetStatusCode(v int32) *Id2MetaVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaVerifyIntlResponse) SetBody(v *Id2MetaVerifyIntlResponseBody) *Id2MetaVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *Id2MetaVerifyIntlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
