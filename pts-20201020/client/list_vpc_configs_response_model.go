// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcConfigsResponseBody) *ListVpcConfigsResponse
	GetBody() *ListVpcConfigsResponseBody
}

type ListVpcConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcConfigsResponse) GetBody() *ListVpcConfigsResponseBody {
	return s.Body
}

func (s *ListVpcConfigsResponse) SetHeaders(v map[string]*string) *ListVpcConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcConfigsResponse) SetStatusCode(v int32) *ListVpcConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcConfigsResponse) SetBody(v *ListVpcConfigsResponseBody) *ListVpcConfigsResponse {
	s.Body = v
	return s
}

func (s *ListVpcConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
