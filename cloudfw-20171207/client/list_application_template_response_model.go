// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationTemplateResponseBody) *ListApplicationTemplateResponse
	GetBody() *ListApplicationTemplateResponseBody
}

type ListApplicationTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationTemplateResponse) GetBody() *ListApplicationTemplateResponseBody {
	return s.Body
}

func (s *ListApplicationTemplateResponse) SetHeaders(v map[string]*string) *ListApplicationTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationTemplateResponse) SetStatusCode(v int32) *ListApplicationTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationTemplateResponse) SetBody(v *ListApplicationTemplateResponseBody) *ListApplicationTemplateResponse {
	s.Body = v
	return s
}

func (s *ListApplicationTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
