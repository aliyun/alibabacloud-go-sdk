// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchUpdateTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchUpdateTasksResponse
	GetStatusCode() *int32
	SetBody(v *BatchUpdateTasksResponseBody) *BatchUpdateTasksResponse
	GetBody() *BatchUpdateTasksResponseBody
}

type BatchUpdateTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchUpdateTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchUpdateTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateTasksResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchUpdateTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchUpdateTasksResponse) GetBody() *BatchUpdateTasksResponseBody {
	return s.Body
}

func (s *BatchUpdateTasksResponse) SetHeaders(v map[string]*string) *BatchUpdateTasksResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateTasksResponse) SetStatusCode(v int32) *BatchUpdateTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateTasksResponse) SetBody(v *BatchUpdateTasksResponseBody) *BatchUpdateTasksResponse {
	s.Body = v
	return s
}

func (s *BatchUpdateTasksResponse) Validate() error {
	return dara.Validate(s)
}
