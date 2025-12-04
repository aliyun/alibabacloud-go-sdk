// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpdsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpdsResponseBody) *ListVpdsResponse
	GetBody() *ListVpdsResponseBody
}

type ListVpdsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpdsResponse) GoString() string {
	return s.String()
}

func (s *ListVpdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpdsResponse) GetBody() *ListVpdsResponseBody {
	return s.Body
}

func (s *ListVpdsResponse) SetHeaders(v map[string]*string) *ListVpdsResponse {
	s.Headers = v
	return s
}

func (s *ListVpdsResponse) SetStatusCode(v int32) *ListVpdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpdsResponse) SetBody(v *ListVpdsResponseBody) *ListVpdsResponse {
	s.Body = v
	return s
}

func (s *ListVpdsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
