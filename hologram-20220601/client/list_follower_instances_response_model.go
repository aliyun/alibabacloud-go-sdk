// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFollowerInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFollowerInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFollowerInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListFollowerInstancesResponseBody) *ListFollowerInstancesResponse
	GetBody() *ListFollowerInstancesResponseBody
}

type ListFollowerInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFollowerInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFollowerInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFollowerInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListFollowerInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFollowerInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFollowerInstancesResponse) GetBody() *ListFollowerInstancesResponseBody {
	return s.Body
}

func (s *ListFollowerInstancesResponse) SetHeaders(v map[string]*string) *ListFollowerInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListFollowerInstancesResponse) SetStatusCode(v int32) *ListFollowerInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFollowerInstancesResponse) SetBody(v *ListFollowerInstancesResponseBody) *ListFollowerInstancesResponse {
	s.Body = v
	return s
}

func (s *ListFollowerInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
