// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirtualResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVirtualResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVirtualResourceResponse
	GetStatusCode() *int32
	SetBody(v *ListVirtualResourceResponseBody) *ListVirtualResourceResponse
	GetBody() *ListVirtualResourceResponseBody
}

type ListVirtualResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVirtualResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVirtualResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVirtualResourceResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVirtualResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVirtualResourceResponse) GetBody() *ListVirtualResourceResponseBody {
	return s.Body
}

func (s *ListVirtualResourceResponse) SetHeaders(v map[string]*string) *ListVirtualResourceResponse {
	s.Headers = v
	return s
}

func (s *ListVirtualResourceResponse) SetStatusCode(v int32) *ListVirtualResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVirtualResourceResponse) SetBody(v *ListVirtualResourceResponseBody) *ListVirtualResourceResponse {
	s.Body = v
	return s
}

func (s *ListVirtualResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
