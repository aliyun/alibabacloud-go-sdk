// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRestoreTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRestoreTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRestoreTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListRestoreTasksResponseBody) *ListRestoreTasksResponse
	GetBody() *ListRestoreTasksResponseBody
}

type ListRestoreTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRestoreTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRestoreTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRestoreTasksResponse) GoString() string {
	return s.String()
}

func (s *ListRestoreTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRestoreTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRestoreTasksResponse) GetBody() *ListRestoreTasksResponseBody {
	return s.Body
}

func (s *ListRestoreTasksResponse) SetHeaders(v map[string]*string) *ListRestoreTasksResponse {
	s.Headers = v
	return s
}

func (s *ListRestoreTasksResponse) SetStatusCode(v int32) *ListRestoreTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRestoreTasksResponse) SetBody(v *ListRestoreTasksResponseBody) *ListRestoreTasksResponse {
	s.Body = v
	return s
}

func (s *ListRestoreTasksResponse) Validate() error {
	return dara.Validate(s)
}
