// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataSourceTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataSourceTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListDataSourceTasksResponseBody) *ListDataSourceTasksResponse
	GetBody() *ListDataSourceTasksResponseBody
}

type ListDataSourceTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataSourceTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTasksResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataSourceTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataSourceTasksResponse) GetBody() *ListDataSourceTasksResponseBody {
	return s.Body
}

func (s *ListDataSourceTasksResponse) SetHeaders(v map[string]*string) *ListDataSourceTasksResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTasksResponse) SetStatusCode(v int32) *ListDataSourceTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTasksResponse) SetBody(v *ListDataSourceTasksResponseBody) *ListDataSourceTasksResponse {
	s.Body = v
	return s
}

func (s *ListDataSourceTasksResponse) Validate() error {
	return dara.Validate(s)
}
