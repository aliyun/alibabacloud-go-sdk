// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchTemplateResponse
	GetStatusCode() *int32
	SetBody(v *SearchTemplateResponseBody) *SearchTemplateResponse
	GetBody() *SearchTemplateResponseBody
}

type SearchTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchTemplateResponse) GoString() string {
	return s.String()
}

func (s *SearchTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchTemplateResponse) GetBody() *SearchTemplateResponseBody {
	return s.Body
}

func (s *SearchTemplateResponse) SetHeaders(v map[string]*string) *SearchTemplateResponse {
	s.Headers = v
	return s
}

func (s *SearchTemplateResponse) SetStatusCode(v int32) *SearchTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTemplateResponse) SetBody(v *SearchTemplateResponseBody) *SearchTemplateResponse {
	s.Body = v
	return s
}

func (s *SearchTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
