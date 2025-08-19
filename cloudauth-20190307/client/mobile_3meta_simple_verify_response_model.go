// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaSimpleVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Mobile3MetaSimpleVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Mobile3MetaSimpleVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Mobile3MetaSimpleVerifyResponseBody) *Mobile3MetaSimpleVerifyResponse
	GetBody() *Mobile3MetaSimpleVerifyResponseBody
}

type Mobile3MetaSimpleVerifyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaSimpleVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaSimpleVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaSimpleVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Mobile3MetaSimpleVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Mobile3MetaSimpleVerifyResponse) GetBody() *Mobile3MetaSimpleVerifyResponseBody {
	return s.Body
}

func (s *Mobile3MetaSimpleVerifyResponse) SetHeaders(v map[string]*string) *Mobile3MetaSimpleVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponse) SetStatusCode(v int32) *Mobile3MetaSimpleVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponse) SetBody(v *Mobile3MetaSimpleVerifyResponseBody) *Mobile3MetaSimpleVerifyResponse {
	s.Body = v
	return s
}

func (s *Mobile3MetaSimpleVerifyResponse) Validate() error {
	return dara.Validate(s)
}
