// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorityTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAuthorityTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAuthorityTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListAuthorityTemplateResponseBody) *ListAuthorityTemplateResponse
	GetBody() *ListAuthorityTemplateResponseBody
}

type ListAuthorityTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAuthorityTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAuthorityTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorityTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListAuthorityTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAuthorityTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAuthorityTemplateResponse) GetBody() *ListAuthorityTemplateResponseBody {
	return s.Body
}

func (s *ListAuthorityTemplateResponse) SetHeaders(v map[string]*string) *ListAuthorityTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListAuthorityTemplateResponse) SetStatusCode(v int32) *ListAuthorityTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAuthorityTemplateResponse) SetBody(v *ListAuthorityTemplateResponseBody) *ListAuthorityTemplateResponse {
	s.Body = v
	return s
}

func (s *ListAuthorityTemplateResponse) Validate() error {
	return dara.Validate(s)
}
