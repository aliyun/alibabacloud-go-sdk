// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationCenterServiceInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationCenterServiceInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationCenterServiceInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationCenterServiceInstancesResponseBody) *ListApplicationCenterServiceInstancesResponse
	GetBody() *ListApplicationCenterServiceInstancesResponseBody
}

type ListApplicationCenterServiceInstancesResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationCenterServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationCenterServiceInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationCenterServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationCenterServiceInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationCenterServiceInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationCenterServiceInstancesResponse) GetBody() *ListApplicationCenterServiceInstancesResponseBody {
	return s.Body
}

func (s *ListApplicationCenterServiceInstancesResponse) SetHeaders(v map[string]*string) *ListApplicationCenterServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponse) SetStatusCode(v int32) *ListApplicationCenterServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponse) SetBody(v *ListApplicationCenterServiceInstancesResponseBody) *ListApplicationCenterServiceInstancesResponse {
	s.Body = v
	return s
}

func (s *ListApplicationCenterServiceInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
