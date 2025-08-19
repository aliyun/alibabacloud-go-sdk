// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id2MetaVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id2MetaVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Id2MetaVerifyResponseBody) *Id2MetaVerifyResponse
	GetBody() *Id2MetaVerifyResponseBody
}

type Id2MetaVerifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id2MetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id2MetaVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id2MetaVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id2MetaVerifyResponse) GetBody() *Id2MetaVerifyResponseBody {
	return s.Body
}

func (s *Id2MetaVerifyResponse) SetHeaders(v map[string]*string) *Id2MetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *Id2MetaVerifyResponse) SetStatusCode(v int32) *Id2MetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Id2MetaVerifyResponse) SetBody(v *Id2MetaVerifyResponseBody) *Id2MetaVerifyResponse {
	s.Body = v
	return s
}

func (s *Id2MetaVerifyResponse) Validate() error {
	return dara.Validate(s)
}
