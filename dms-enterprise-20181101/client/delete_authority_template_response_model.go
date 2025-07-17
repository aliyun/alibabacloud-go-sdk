// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAuthorityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAuthorityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAuthorityTemplateResponseBody) *DeleteAuthorityTemplateResponse
	GetBody() *DeleteAuthorityTemplateResponseBody
}

type DeleteAuthorityTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuthorityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAuthorityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorityTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthorityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAuthorityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAuthorityTemplateResponse) GetBody() *DeleteAuthorityTemplateResponseBody {
	return s.Body
}

func (s *DeleteAuthorityTemplateResponse) SetHeaders(v map[string]*string) *DeleteAuthorityTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthorityTemplateResponse) SetStatusCode(v int32) *DeleteAuthorityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuthorityTemplateResponse) SetBody(v *DeleteAuthorityTemplateResponseBody) *DeleteAuthorityTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteAuthorityTemplateResponse) Validate() error {
	return dara.Validate(s)
}
