// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComponentTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComponentTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ComponentTemplatePageResult) *ListComponentTemplatesResponse
	GetBody() *ComponentTemplatePageResult
}

type ListComponentTemplatesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ComponentTemplatePageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComponentTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComponentTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListComponentTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComponentTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComponentTemplatesResponse) GetBody() *ComponentTemplatePageResult {
	return s.Body
}

func (s *ListComponentTemplatesResponse) SetHeaders(v map[string]*string) *ListComponentTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListComponentTemplatesResponse) SetStatusCode(v int32) *ListComponentTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComponentTemplatesResponse) SetBody(v *ComponentTemplatePageResult) *ListComponentTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListComponentTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
