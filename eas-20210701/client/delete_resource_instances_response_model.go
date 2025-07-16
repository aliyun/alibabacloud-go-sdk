// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceInstancesResponseBody) *DeleteResourceInstancesResponse
	GetBody() *DeleteResourceInstancesResponseBody
}

type DeleteResourceInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceInstancesResponse) GetBody() *DeleteResourceInstancesResponseBody {
	return s.Body
}

func (s *DeleteResourceInstancesResponse) SetHeaders(v map[string]*string) *DeleteResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceInstancesResponse) SetStatusCode(v int32) *DeleteResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceInstancesResponse) SetBody(v *DeleteResourceInstancesResponseBody) *DeleteResourceInstancesResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceInstancesResponse) Validate() error {
	return dara.Validate(s)
}
