// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListStackInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListStackInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListStackInstancesResponseBody) *ListStackInstancesResponse
	GetBody() *ListStackInstancesResponseBody
}

type ListStackInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStackInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListStackInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListStackInstancesResponse) GetBody() *ListStackInstancesResponseBody {
	return s.Body
}

func (s *ListStackInstancesResponse) SetHeaders(v map[string]*string) *ListStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListStackInstancesResponse) SetStatusCode(v int32) *ListStackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackInstancesResponse) SetBody(v *ListStackInstancesResponseBody) *ListStackInstancesResponse {
	s.Body = v
	return s
}

func (s *ListStackInstancesResponse) Validate() error {
	return dara.Validate(s)
}
