// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplateResponseBody) *ListTemplateResponse
	GetBody() *ListTemplateResponseBody
}

type ListTemplateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplateResponse) GetBody() *ListTemplateResponseBody {
	return s.Body
}

func (s *ListTemplateResponse) SetHeaders(v map[string]*string) *ListTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateResponse) SetStatusCode(v int32) *ListTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateResponse) SetBody(v *ListTemplateResponseBody) *ListTemplateResponse {
	s.Body = v
	return s
}

func (s *ListTemplateResponse) Validate() error {
	return dara.Validate(s)
}
