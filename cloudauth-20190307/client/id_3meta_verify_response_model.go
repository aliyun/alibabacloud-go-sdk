// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Id3MetaVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Id3MetaVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Id3MetaVerifyResponseBody) *Id3MetaVerifyResponse
	GetBody() *Id3MetaVerifyResponseBody
}

type Id3MetaVerifyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Id3MetaVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Id3MetaVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyResponse) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Id3MetaVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Id3MetaVerifyResponse) GetBody() *Id3MetaVerifyResponseBody {
	return s.Body
}

func (s *Id3MetaVerifyResponse) SetHeaders(v map[string]*string) *Id3MetaVerifyResponse {
	s.Headers = v
	return s
}

func (s *Id3MetaVerifyResponse) SetStatusCode(v int32) *Id3MetaVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Id3MetaVerifyResponse) SetBody(v *Id3MetaVerifyResponseBody) *Id3MetaVerifyResponse {
	s.Body = v
	return s
}

func (s *Id3MetaVerifyResponse) Validate() error {
	return dara.Validate(s)
}
