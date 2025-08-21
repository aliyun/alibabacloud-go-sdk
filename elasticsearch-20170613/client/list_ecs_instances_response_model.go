// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEcsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEcsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEcsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *ListEcsInstancesResponseBody) *ListEcsInstancesResponse
	GetBody() *ListEcsInstancesResponseBody
}

type ListEcsInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEcsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEcsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEcsInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListEcsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEcsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEcsInstancesResponse) GetBody() *ListEcsInstancesResponseBody {
	return s.Body
}

func (s *ListEcsInstancesResponse) SetHeaders(v map[string]*string) *ListEcsInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListEcsInstancesResponse) SetStatusCode(v int32) *ListEcsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEcsInstancesResponse) SetBody(v *ListEcsInstancesResponseBody) *ListEcsInstancesResponse {
	s.Body = v
	return s
}

func (s *ListEcsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
