// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcResponseBody) *ListVpcResponse
	GetBody() *ListVpcResponseBody
}

type ListVpcResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcResponse) GoString() string {
	return s.String()
}

func (s *ListVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcResponse) GetBody() *ListVpcResponseBody {
	return s.Body
}

func (s *ListVpcResponse) SetHeaders(v map[string]*string) *ListVpcResponse {
	s.Headers = v
	return s
}

func (s *ListVpcResponse) SetStatusCode(v int32) *ListVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcResponse) SetBody(v *ListVpcResponseBody) *ListVpcResponse {
	s.Body = v
	return s
}

func (s *ListVpcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
