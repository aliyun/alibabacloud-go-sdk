// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSearchTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSearchTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSearchTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListSearchTasksResponseBody) *ListSearchTasksResponse
	GetBody() *ListSearchTasksResponseBody
}

type ListSearchTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSearchTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSearchTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSearchTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSearchTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSearchTasksResponse) GetBody() *ListSearchTasksResponseBody {
	return s.Body
}

func (s *ListSearchTasksResponse) SetHeaders(v map[string]*string) *ListSearchTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSearchTasksResponse) SetStatusCode(v int32) *ListSearchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchTasksResponse) SetBody(v *ListSearchTasksResponseBody) *ListSearchTasksResponse {
	s.Body = v
	return s
}

func (s *ListSearchTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
