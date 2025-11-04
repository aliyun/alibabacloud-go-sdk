// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizedVoicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomizedVoicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomizedVoicesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomizedVoicesResponseBody) *ListCustomizedVoicesResponse
	GetBody() *ListCustomizedVoicesResponseBody
}

type ListCustomizedVoicesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomizedVoicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomizedVoicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizedVoicesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomizedVoicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomizedVoicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomizedVoicesResponse) GetBody() *ListCustomizedVoicesResponseBody {
	return s.Body
}

func (s *ListCustomizedVoicesResponse) SetHeaders(v map[string]*string) *ListCustomizedVoicesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomizedVoicesResponse) SetStatusCode(v int32) *ListCustomizedVoicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomizedVoicesResponse) SetBody(v *ListCustomizedVoicesResponseBody) *ListCustomizedVoicesResponse {
	s.Body = v
	return s
}

func (s *ListCustomizedVoicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
