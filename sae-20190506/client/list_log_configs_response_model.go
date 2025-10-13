// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListLogConfigsResponseBody) *ListLogConfigsResponse
	GetBody() *ListLogConfigsResponseBody
}

type ListLogConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListLogConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogConfigsResponse) GetBody() *ListLogConfigsResponseBody {
	return s.Body
}

func (s *ListLogConfigsResponse) SetHeaders(v map[string]*string) *ListLogConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListLogConfigsResponse) SetStatusCode(v int32) *ListLogConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogConfigsResponse) SetBody(v *ListLogConfigsResponseBody) *ListLogConfigsResponse {
	s.Body = v
	return s
}

func (s *ListLogConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
