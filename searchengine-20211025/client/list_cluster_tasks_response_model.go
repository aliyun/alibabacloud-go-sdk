// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterTasksResponseBody) *ListClusterTasksResponse
	GetBody() *ListClusterTasksResponseBody
}

type ListClusterTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTasksResponse) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterTasksResponse) GetBody() *ListClusterTasksResponseBody {
	return s.Body
}

func (s *ListClusterTasksResponse) SetHeaders(v map[string]*string) *ListClusterTasksResponse {
	s.Headers = v
	return s
}

func (s *ListClusterTasksResponse) SetStatusCode(v int32) *ListClusterTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterTasksResponse) SetBody(v *ListClusterTasksResponseBody) *ListClusterTasksResponse {
	s.Body = v
	return s
}

func (s *ListClusterTasksResponse) Validate() error {
	return dara.Validate(s)
}
