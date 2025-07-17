// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTaskInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTaskInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListTaskInstancesResponseBody) *ListTaskInstancesResponse
	GetBody() *ListTaskInstancesResponseBody
}

type ListTaskInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTaskInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListTaskInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTaskInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTaskInstancesResponse) GetBody() *ListTaskInstancesResponseBody {
	return s.Body
}

func (s *ListTaskInstancesResponse) SetHeaders(v map[string]*string) *ListTaskInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListTaskInstancesResponse) SetStatusCode(v int32) *ListTaskInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskInstancesResponse) SetBody(v *ListTaskInstancesResponseBody) *ListTaskInstancesResponse {
	s.Body = v
	return s
}

func (s *ListTaskInstancesResponse) Validate() error {
	return dara.Validate(s)
}
