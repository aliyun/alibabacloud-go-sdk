// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTasksResponse
	GetStatusCode() *int32
	SetBody(v []*Task) *ListTasksResponse
	GetBody() []*Task
}

type ListTasksResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*Task            `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTasksResponse) GetBody() []*Task {
	return s.Body
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v []*Task) *ListTasksResponse {
	s.Body = v
	return s
}

func (s *ListTasksResponse) Validate() error {
	return dara.Validate(s)
}
