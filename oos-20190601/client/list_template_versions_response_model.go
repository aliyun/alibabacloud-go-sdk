// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplateVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplateVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplateVersionsResponseBody) *ListTemplateVersionsResponse
	GetBody() *ListTemplateVersionsResponseBody
}

type ListTemplateVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplateVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplateVersionsResponse) GetBody() *ListTemplateVersionsResponseBody {
	return s.Body
}

func (s *ListTemplateVersionsResponse) SetHeaders(v map[string]*string) *ListTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateVersionsResponse) SetStatusCode(v int32) *ListTemplateVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateVersionsResponse) SetBody(v *ListTemplateVersionsResponseBody) *ListTemplateVersionsResponse {
	s.Body = v
	return s
}

func (s *ListTemplateVersionsResponse) Validate() error {
	return dara.Validate(s)
}
