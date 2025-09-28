// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVerifySchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVerifySchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVerifySchemeResponse
	GetStatusCode() *int32
	SetBody(v *CreateVerifySchemeResponseBody) *CreateVerifySchemeResponse
	GetBody() *CreateVerifySchemeResponseBody
}

type CreateVerifySchemeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVerifySchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVerifySchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVerifySchemeResponse) GetBody() *CreateVerifySchemeResponseBody {
	return s.Body
}

func (s *CreateVerifySchemeResponse) SetHeaders(v map[string]*string) *CreateVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *CreateVerifySchemeResponse) SetStatusCode(v int32) *CreateVerifySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVerifySchemeResponse) SetBody(v *CreateVerifySchemeResponseBody) *CreateVerifySchemeResponse {
	s.Body = v
	return s
}

func (s *CreateVerifySchemeResponse) Validate() error {
	return dara.Validate(s)
}
