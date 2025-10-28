// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAsyncTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAsyncTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListAsyncTaskOutput) *ListAsyncTasksResponse
	GetBody() *ListAsyncTaskOutput
}

type ListAsyncTasksResponse struct {
	Headers    map[string]*string   `json:"headers" xml:"headers"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsyncTaskOutput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsyncTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAsyncTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAsyncTasksResponse) GetBody() *ListAsyncTaskOutput {
	return s.Body
}

func (s *ListAsyncTasksResponse) SetHeaders(v map[string]*string) *ListAsyncTasksResponse {
	s.Headers = v
	return s
}

func (s *ListAsyncTasksResponse) SetStatusCode(v int32) *ListAsyncTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsyncTasksResponse) SetBody(v *ListAsyncTaskOutput) *ListAsyncTasksResponse {
	s.Body = v
	return s
}

func (s *ListAsyncTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
