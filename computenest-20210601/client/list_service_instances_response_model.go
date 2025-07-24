// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse
	GetBody() *ListServiceInstancesResponseBody
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceInstancesResponse) GetBody() *ListServiceInstancesResponseBody {
	return s.Body
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

func (s *ListServiceInstancesResponse) Validate() error {
	return dara.Validate(s)
}
