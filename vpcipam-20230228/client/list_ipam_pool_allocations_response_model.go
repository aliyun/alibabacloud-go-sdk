// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamPoolAllocationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIpamPoolAllocationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIpamPoolAllocationsResponse
	GetStatusCode() *int32
	SetBody(v *ListIpamPoolAllocationsResponseBody) *ListIpamPoolAllocationsResponse
	GetBody() *ListIpamPoolAllocationsResponseBody
}

type ListIpamPoolAllocationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIpamPoolAllocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIpamPoolAllocationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIpamPoolAllocationsResponse) GoString() string {
	return s.String()
}

func (s *ListIpamPoolAllocationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIpamPoolAllocationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIpamPoolAllocationsResponse) GetBody() *ListIpamPoolAllocationsResponseBody {
	return s.Body
}

func (s *ListIpamPoolAllocationsResponse) SetHeaders(v map[string]*string) *ListIpamPoolAllocationsResponse {
	s.Headers = v
	return s
}

func (s *ListIpamPoolAllocationsResponse) SetStatusCode(v int32) *ListIpamPoolAllocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIpamPoolAllocationsResponse) SetBody(v *ListIpamPoolAllocationsResponseBody) *ListIpamPoolAllocationsResponse {
	s.Body = v
	return s
}

func (s *ListIpamPoolAllocationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
