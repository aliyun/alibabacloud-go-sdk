// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile2MetaVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Mobile2MetaVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Mobile2MetaVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Mobile2MetaVerifyResponseBody) *Mobile2MetaVerifyResponse
	GetBody() *Mobile2MetaVerifyResponseBody
}

type Mobile2MetaVerifyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile2MetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile2MetaVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Mobile2MetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile2MetaVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Mobile2MetaVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Mobile2MetaVerifyResponse) GetBody() *Mobile2MetaVerifyResponseBody {
	return s.Body
}

func (s *Mobile2MetaVerifyResponse) SetHeaders(v map[string]*string) *Mobile2MetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile2MetaVerifyResponse) SetStatusCode(v int32) *Mobile2MetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile2MetaVerifyResponse) SetBody(v *Mobile2MetaVerifyResponseBody) *Mobile2MetaVerifyResponse {
	s.Body = v
	return s
}

func (s *Mobile2MetaVerifyResponse) Validate() error {
	return dara.Validate(s)
}
