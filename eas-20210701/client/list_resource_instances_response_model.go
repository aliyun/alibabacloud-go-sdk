// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceInstancesResponseBody) *ListResourceInstancesResponse
	GetBody() *ListResourceInstancesResponseBody
}

type ListResourceInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceInstancesResponse) GetBody() *ListResourceInstancesResponseBody {
	return s.Body
}

func (s *ListResourceInstancesResponse) SetHeaders(v map[string]*string) *ListResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceInstancesResponse) SetStatusCode(v int32) *ListResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceInstancesResponse) SetBody(v *ListResourceInstancesResponseBody) *ListResourceInstancesResponse {
	s.Body = v
	return s
}

func (s *ListResourceInstancesResponse) Validate() error {
	return dara.Validate(s)
}
