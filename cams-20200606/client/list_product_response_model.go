// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductResponse
	GetStatusCode() *int32
	SetBody(v *ListProductResponseBody) *ListProductResponse
	GetBody() *ListProductResponseBody
}

type ListProductResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductResponse) GoString() string {
	return s.String()
}

func (s *ListProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductResponse) GetBody() *ListProductResponseBody {
	return s.Body
}

func (s *ListProductResponse) SetHeaders(v map[string]*string) *ListProductResponse {
	s.Headers = v
	return s
}

func (s *ListProductResponse) SetStatusCode(v int32) *ListProductResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductResponse) SetBody(v *ListProductResponseBody) *ListProductResponse {
	s.Body = v
	return s
}

func (s *ListProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
