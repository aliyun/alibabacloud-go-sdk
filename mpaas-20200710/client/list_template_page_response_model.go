// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatePageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTemplatePageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTemplatePageResponse
	GetStatusCode() *int32
	SetBody(v *ListTemplatePageResponseBody) *ListTemplatePageResponse
	GetBody() *ListTemplatePageResponseBody
}

type ListTemplatePageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplatePageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplatePageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatePageResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatePageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTemplatePageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTemplatePageResponse) GetBody() *ListTemplatePageResponseBody {
	return s.Body
}

func (s *ListTemplatePageResponse) SetHeaders(v map[string]*string) *ListTemplatePageResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatePageResponse) SetStatusCode(v int32) *ListTemplatePageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplatePageResponse) SetBody(v *ListTemplatePageResponseBody) *ListTemplatePageResponse {
	s.Body = v
	return s
}

func (s *ListTemplatePageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
