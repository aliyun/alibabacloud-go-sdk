// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVerifySchemeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVerifySchemeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVerifySchemeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVerifySchemeResponseBody) *DeleteVerifySchemeResponse
	GetBody() *DeleteVerifySchemeResponseBody
}

type DeleteVerifySchemeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVerifySchemeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVerifySchemeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVerifySchemeResponse) GoString() string {
	return s.String()
}

func (s *DeleteVerifySchemeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVerifySchemeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVerifySchemeResponse) GetBody() *DeleteVerifySchemeResponseBody {
	return s.Body
}

func (s *DeleteVerifySchemeResponse) SetHeaders(v map[string]*string) *DeleteVerifySchemeResponse {
	s.Headers = v
	return s
}

func (s *DeleteVerifySchemeResponse) SetStatusCode(v int32) *DeleteVerifySchemeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVerifySchemeResponse) SetBody(v *DeleteVerifySchemeResponseBody) *DeleteVerifySchemeResponse {
	s.Body = v
	return s
}

func (s *DeleteVerifySchemeResponse) Validate() error {
	return dara.Validate(s)
}
