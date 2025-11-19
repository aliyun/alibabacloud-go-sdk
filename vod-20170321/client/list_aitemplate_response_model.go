// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAITemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAITemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListAITemplateResponseBody) *ListAITemplateResponse
	GetBody() *ListAITemplateResponseBody
}

type ListAITemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAITemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAITemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAITemplateResponse) GoString() string {
	return s.String()
}

func (s *ListAITemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAITemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAITemplateResponse) GetBody() *ListAITemplateResponseBody {
	return s.Body
}

func (s *ListAITemplateResponse) SetHeaders(v map[string]*string) *ListAITemplateResponse {
	s.Headers = v
	return s
}

func (s *ListAITemplateResponse) SetStatusCode(v int32) *ListAITemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAITemplateResponse) SetBody(v *ListAITemplateResponseBody) *ListAITemplateResponse {
	s.Body = v
	return s
}

func (s *ListAITemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
