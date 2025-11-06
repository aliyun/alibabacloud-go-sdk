// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEurekaInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEurekaInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEurekaInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListEurekaInstancesResponseBody) *ListEurekaInstancesResponse
	GetBody() *ListEurekaInstancesResponseBody
}

type ListEurekaInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEurekaInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEurekaInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEurekaInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListEurekaInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEurekaInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEurekaInstancesResponse) GetBody() *ListEurekaInstancesResponseBody {
	return s.Body
}

func (s *ListEurekaInstancesResponse) SetHeaders(v map[string]*string) *ListEurekaInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListEurekaInstancesResponse) SetStatusCode(v int32) *ListEurekaInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEurekaInstancesResponse) SetBody(v *ListEurekaInstancesResponseBody) *ListEurekaInstancesResponse {
	s.Body = v
	return s
}

func (s *ListEurekaInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
