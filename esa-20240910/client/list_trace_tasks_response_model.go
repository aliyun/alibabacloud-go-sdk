// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTraceTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTraceTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTraceTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListTraceTasksResponseBody) *ListTraceTasksResponse
	GetBody() *ListTraceTasksResponseBody
}

type ListTraceTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTraceTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTraceTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTraceTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTraceTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTraceTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTraceTasksResponse) GetBody() *ListTraceTasksResponseBody {
	return s.Body
}

func (s *ListTraceTasksResponse) SetHeaders(v map[string]*string) *ListTraceTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTraceTasksResponse) SetStatusCode(v int32) *ListTraceTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTraceTasksResponse) SetBody(v *ListTraceTasksResponseBody) *ListTraceTasksResponse {
	s.Body = v
	return s
}

func (s *ListTraceTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
