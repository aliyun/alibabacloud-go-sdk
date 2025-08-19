// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMobile3MetaSimpleStandardVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Mobile3MetaSimpleStandardVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Mobile3MetaSimpleStandardVerifyResponse
	GetStatusCode() *int32
	SetBody(v *Mobile3MetaSimpleStandardVerifyResponseBody) *Mobile3MetaSimpleStandardVerifyResponse
	GetBody() *Mobile3MetaSimpleStandardVerifyResponseBody
}

type Mobile3MetaSimpleStandardVerifyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Mobile3MetaSimpleStandardVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Mobile3MetaSimpleStandardVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s Mobile3MetaSimpleStandardVerifyResponse) GoString() string {
	return s.String()
}

func (s *Mobile3MetaSimpleStandardVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Mobile3MetaSimpleStandardVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Mobile3MetaSimpleStandardVerifyResponse) GetBody() *Mobile3MetaSimpleStandardVerifyResponseBody {
	return s.Body
}

func (s *Mobile3MetaSimpleStandardVerifyResponse) SetHeaders(v map[string]*string) *Mobile3MetaSimpleStandardVerifyResponse {
	s.Headers = v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponse) SetStatusCode(v int32) *Mobile3MetaSimpleStandardVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponse) SetBody(v *Mobile3MetaSimpleStandardVerifyResponseBody) *Mobile3MetaSimpleStandardVerifyResponse {
	s.Body = v
	return s
}

func (s *Mobile3MetaSimpleStandardVerifyResponse) Validate() error {
	return dara.Validate(s)
}
