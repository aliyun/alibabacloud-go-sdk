// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiConfigErrorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDpiConfigErrorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDpiConfigErrorResponse
	GetStatusCode() *int32
	SetBody(v *ListDpiConfigErrorResponseBody) *ListDpiConfigErrorResponse
	GetBody() *ListDpiConfigErrorResponseBody
}

type ListDpiConfigErrorResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDpiConfigErrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDpiConfigErrorResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDpiConfigErrorResponse) GoString() string {
	return s.String()
}

func (s *ListDpiConfigErrorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDpiConfigErrorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDpiConfigErrorResponse) GetBody() *ListDpiConfigErrorResponseBody {
	return s.Body
}

func (s *ListDpiConfigErrorResponse) SetHeaders(v map[string]*string) *ListDpiConfigErrorResponse {
	s.Headers = v
	return s
}

func (s *ListDpiConfigErrorResponse) SetStatusCode(v int32) *ListDpiConfigErrorResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDpiConfigErrorResponse) SetBody(v *ListDpiConfigErrorResponseBody) *ListDpiConfigErrorResponse {
	s.Body = v
	return s
}

func (s *ListDpiConfigErrorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
