// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduledTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduledTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduledTasksResponseBody) *ListScheduledTasksResponse
	GetBody() *ListScheduledTasksResponseBody
}

type ListScheduledTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledTasksResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduledTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduledTasksResponse) GetBody() *ListScheduledTasksResponseBody {
	return s.Body
}

func (s *ListScheduledTasksResponse) SetHeaders(v map[string]*string) *ListScheduledTasksResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledTasksResponse) SetStatusCode(v int32) *ListScheduledTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledTasksResponse) SetBody(v *ListScheduledTasksResponseBody) *ListScheduledTasksResponse {
	s.Body = v
	return s
}

func (s *ListScheduledTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
