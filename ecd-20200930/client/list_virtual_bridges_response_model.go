// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualBridgesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirtualBridgesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirtualBridgesResponse
	GetStatusCode() *int32
	SetBody(v *ListVirtualBridgesResponseBody) *ListVirtualBridgesResponse
	GetBody() *ListVirtualBridgesResponseBody
}

type ListVirtualBridgesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualBridgesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualBridgesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualBridgesResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualBridgesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirtualBridgesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirtualBridgesResponse) GetBody() *ListVirtualBridgesResponseBody {
	return s.Body
}

func (s *ListVirtualBridgesResponse) SetHeaders(v map[string]*string) *ListVirtualBridgesResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualBridgesResponse) SetStatusCode(v int32) *ListVirtualBridgesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualBridgesResponse) SetBody(v *ListVirtualBridgesResponseBody) *ListVirtualBridgesResponse {
	s.Body = v
	return s
}

func (s *ListVirtualBridgesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
