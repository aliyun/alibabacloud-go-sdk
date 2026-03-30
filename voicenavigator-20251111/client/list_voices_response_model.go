// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVoicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVoicesResponse
	GetStatusCode() *int32
	SetBody(v *ListVoicesResponseBody) *ListVoicesResponse
	GetBody() *ListVoicesResponseBody
}

type ListVoicesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVoicesResponse) GoString() string {
	return s.String()
}

func (s *ListVoicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVoicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVoicesResponse) GetBody() *ListVoicesResponseBody {
	return s.Body
}

func (s *ListVoicesResponse) SetHeaders(v map[string]*string) *ListVoicesResponse {
	s.Headers = v
	return s
}

func (s *ListVoicesResponse) SetStatusCode(v int32) *ListVoicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoicesResponse) SetBody(v *ListVoicesResponseBody) *ListVoicesResponse {
	s.Body = v
	return s
}

func (s *ListVoicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
