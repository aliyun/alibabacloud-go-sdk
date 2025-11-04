// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomTemplatesResponseBody) *ListCustomTemplatesResponse
	GetBody() *ListCustomTemplatesResponseBody
}

type ListCustomTemplatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomTemplatesResponse) GetBody() *ListCustomTemplatesResponseBody {
	return s.Body
}

func (s *ListCustomTemplatesResponse) SetHeaders(v map[string]*string) *ListCustomTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomTemplatesResponse) SetStatusCode(v int32) *ListCustomTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomTemplatesResponse) SetBody(v *ListCustomTemplatesResponseBody) *ListCustomTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListCustomTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
