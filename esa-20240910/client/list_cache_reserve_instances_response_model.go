// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheReserveInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCacheReserveInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCacheReserveInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListCacheReserveInstancesResponseBody) *ListCacheReserveInstancesResponse
	GetBody() *ListCacheReserveInstancesResponseBody
}

type ListCacheReserveInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCacheReserveInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCacheReserveInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCacheReserveInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListCacheReserveInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCacheReserveInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCacheReserveInstancesResponse) GetBody() *ListCacheReserveInstancesResponseBody {
	return s.Body
}

func (s *ListCacheReserveInstancesResponse) SetHeaders(v map[string]*string) *ListCacheReserveInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListCacheReserveInstancesResponse) SetStatusCode(v int32) *ListCacheReserveInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCacheReserveInstancesResponse) SetBody(v *ListCacheReserveInstancesResponseBody) *ListCacheReserveInstancesResponse {
	s.Body = v
	return s
}

func (s *ListCacheReserveInstancesResponse) Validate() error {
	return dara.Validate(s)
}
