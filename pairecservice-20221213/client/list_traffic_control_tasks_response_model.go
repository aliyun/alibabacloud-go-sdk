// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficControlTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrafficControlTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrafficControlTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListTrafficControlTasksResponseBody) *ListTrafficControlTasksResponse
	GetBody() *ListTrafficControlTasksResponseBody
}

type ListTrafficControlTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrafficControlTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrafficControlTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficControlTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficControlTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrafficControlTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrafficControlTasksResponse) GetBody() *ListTrafficControlTasksResponseBody {
	return s.Body
}

func (s *ListTrafficControlTasksResponse) SetHeaders(v map[string]*string) *ListTrafficControlTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficControlTasksResponse) SetStatusCode(v int32) *ListTrafficControlTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrafficControlTasksResponse) SetBody(v *ListTrafficControlTasksResponseBody) *ListTrafficControlTasksResponse {
	s.Body = v
	return s
}

func (s *ListTrafficControlTasksResponse) Validate() error {
	return dara.Validate(s)
}
