// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiUserInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultiUserInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultiUserInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListMultiUserInstancesResponseBody) *ListMultiUserInstancesResponse
	GetBody() *ListMultiUserInstancesResponseBody
}

type ListMultiUserInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiUserInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiUserInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultiUserInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListMultiUserInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultiUserInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultiUserInstancesResponse) GetBody() *ListMultiUserInstancesResponseBody {
	return s.Body
}

func (s *ListMultiUserInstancesResponse) SetHeaders(v map[string]*string) *ListMultiUserInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListMultiUserInstancesResponse) SetStatusCode(v int32) *ListMultiUserInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiUserInstancesResponse) SetBody(v *ListMultiUserInstancesResponseBody) *ListMultiUserInstancesResponse {
	s.Body = v
	return s
}

func (s *ListMultiUserInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
