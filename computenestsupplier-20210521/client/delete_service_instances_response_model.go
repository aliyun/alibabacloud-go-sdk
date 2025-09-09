// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse
	GetBody() *DeleteServiceInstancesResponseBody
}

type DeleteServiceInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceInstancesResponse) GetBody() *DeleteServiceInstancesResponseBody {
	return s.Body
}

func (s *DeleteServiceInstancesResponse) SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceInstancesResponse) SetStatusCode(v int32) *DeleteServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceInstancesResponse) Validate() error {
	return dara.Validate(s)
}
