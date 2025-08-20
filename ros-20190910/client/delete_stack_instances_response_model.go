// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStackInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStackInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStackInstancesResponseBody) *DeleteStackInstancesResponse
	GetBody() *DeleteStackInstancesResponseBody
}

type DeleteStackInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStackInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStackInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStackInstancesResponse) GetBody() *DeleteStackInstancesResponseBody {
	return s.Body
}

func (s *DeleteStackInstancesResponse) SetHeaders(v map[string]*string) *DeleteStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackInstancesResponse) SetStatusCode(v int32) *DeleteStackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStackInstancesResponse) SetBody(v *DeleteStackInstancesResponseBody) *DeleteStackInstancesResponse {
	s.Body = v
	return s
}

func (s *DeleteStackInstancesResponse) Validate() error {
	return dara.Validate(s)
}
