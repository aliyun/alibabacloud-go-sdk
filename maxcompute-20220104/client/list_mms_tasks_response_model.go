// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMmsTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMmsTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListMmsTasksResponseBody) *ListMmsTasksResponse
	GetBody() *ListMmsTasksResponseBody
}

type ListMmsTasksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMmsTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMmsTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMmsTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMmsTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMmsTasksResponse) GetBody() *ListMmsTasksResponseBody {
	return s.Body
}

func (s *ListMmsTasksResponse) SetHeaders(v map[string]*string) *ListMmsTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMmsTasksResponse) SetStatusCode(v int32) *ListMmsTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMmsTasksResponse) SetBody(v *ListMmsTasksResponseBody) *ListMmsTasksResponse {
	s.Body = v
	return s
}

func (s *ListMmsTasksResponse) Validate() error {
	return dara.Validate(s)
}
