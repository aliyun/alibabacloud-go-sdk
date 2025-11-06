// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsoleProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsoleProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsoleProductResponse
	GetStatusCode() *int32
	SetBody(v *ListConsoleProductResponseBody) *ListConsoleProductResponse
	GetBody() *ListConsoleProductResponseBody
}

type ListConsoleProductResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsoleProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsoleProductResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsoleProductResponse) GoString() string {
	return s.String()
}

func (s *ListConsoleProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsoleProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsoleProductResponse) GetBody() *ListConsoleProductResponseBody {
	return s.Body
}

func (s *ListConsoleProductResponse) SetHeaders(v map[string]*string) *ListConsoleProductResponse {
	s.Headers = v
	return s
}

func (s *ListConsoleProductResponse) SetStatusCode(v int32) *ListConsoleProductResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsoleProductResponse) SetBody(v *ListConsoleProductResponseBody) *ListConsoleProductResponse {
	s.Body = v
	return s
}

func (s *ListConsoleProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
