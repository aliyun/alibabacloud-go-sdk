// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationHostsResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationHostsResponseBody) *ListOperationHostsResponse
	GetBody() *ListOperationHostsResponseBody
}

type ListOperationHostsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHostsResponse) GoString() string {
	return s.String()
}

func (s *ListOperationHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationHostsResponse) GetBody() *ListOperationHostsResponseBody {
	return s.Body
}

func (s *ListOperationHostsResponse) SetHeaders(v map[string]*string) *ListOperationHostsResponse {
	s.Headers = v
	return s
}

func (s *ListOperationHostsResponse) SetStatusCode(v int32) *ListOperationHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationHostsResponse) SetBody(v *ListOperationHostsResponseBody) *ListOperationHostsResponse {
	s.Body = v
	return s
}

func (s *ListOperationHostsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
