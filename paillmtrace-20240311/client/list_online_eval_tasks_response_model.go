// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnlineEvalTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOnlineEvalTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOnlineEvalTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListOnlineEvalTasksResponseBody) *ListOnlineEvalTasksResponse
	GetBody() *ListOnlineEvalTasksResponseBody
}

type ListOnlineEvalTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOnlineEvalTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOnlineEvalTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOnlineEvalTasksResponse) GoString() string {
	return s.String()
}

func (s *ListOnlineEvalTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOnlineEvalTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOnlineEvalTasksResponse) GetBody() *ListOnlineEvalTasksResponseBody {
	return s.Body
}

func (s *ListOnlineEvalTasksResponse) SetHeaders(v map[string]*string) *ListOnlineEvalTasksResponse {
	s.Headers = v
	return s
}

func (s *ListOnlineEvalTasksResponse) SetStatusCode(v int32) *ListOnlineEvalTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOnlineEvalTasksResponse) SetBody(v *ListOnlineEvalTasksResponseBody) *ListOnlineEvalTasksResponse {
	s.Body = v
	return s
}

func (s *ListOnlineEvalTasksResponse) Validate() error {
	return dara.Validate(s)
}
