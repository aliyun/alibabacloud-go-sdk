// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeAndroidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AuthorizeAndroidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AuthorizeAndroidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *AuthorizeAndroidInstanceResponseBody) *AuthorizeAndroidInstanceResponse
	GetBody() *AuthorizeAndroidInstanceResponseBody
}

type AuthorizeAndroidInstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AuthorizeAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AuthorizeAndroidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeAndroidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AuthorizeAndroidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AuthorizeAndroidInstanceResponse) GetBody() *AuthorizeAndroidInstanceResponseBody {
	return s.Body
}

func (s *AuthorizeAndroidInstanceResponse) SetHeaders(v map[string]*string) *AuthorizeAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeAndroidInstanceResponse) SetStatusCode(v int32) *AuthorizeAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeAndroidInstanceResponse) SetBody(v *AuthorizeAndroidInstanceResponseBody) *AuthorizeAndroidInstanceResponse {
	s.Body = v
	return s
}

func (s *AuthorizeAndroidInstanceResponse) Validate() error {
	return dara.Validate(s)
}
