// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComputeInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComputeInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListComputeInstancesResponseBody) *ListComputeInstancesResponse
	GetBody() *ListComputeInstancesResponseBody
}

type ListComputeInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComputeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComputeInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComputeInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComputeInstancesResponse) GetBody() *ListComputeInstancesResponseBody {
	return s.Body
}

func (s *ListComputeInstancesResponse) SetHeaders(v map[string]*string) *ListComputeInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListComputeInstancesResponse) SetStatusCode(v int32) *ListComputeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComputeInstancesResponse) SetBody(v *ListComputeInstancesResponseBody) *ListComputeInstancesResponse {
	s.Body = v
	return s
}

func (s *ListComputeInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
