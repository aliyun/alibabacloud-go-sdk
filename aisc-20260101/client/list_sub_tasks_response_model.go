// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListSubTasksResponseBody) *ListSubTasksResponse
	GetBody() *ListSubTasksResponseBody
}

type ListSubTasksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSubTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubTasksResponse) GetBody() *ListSubTasksResponseBody {
	return s.Body
}

func (s *ListSubTasksResponse) SetHeaders(v map[string]*string) *ListSubTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSubTasksResponse) SetStatusCode(v int32) *ListSubTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubTasksResponse) SetBody(v *ListSubTasksResponseBody) *ListSubTasksResponse {
	s.Body = v
	return s
}

func (s *ListSubTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
