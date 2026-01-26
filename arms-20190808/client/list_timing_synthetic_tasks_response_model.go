// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTimingSyntheticTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTimingSyntheticTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTimingSyntheticTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListTimingSyntheticTasksResponseBody) *ListTimingSyntheticTasksResponse
	GetBody() *ListTimingSyntheticTasksResponseBody
}

type ListTimingSyntheticTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTimingSyntheticTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTimingSyntheticTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTimingSyntheticTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTimingSyntheticTasksResponse) GetBody() *ListTimingSyntheticTasksResponseBody {
	return s.Body
}

func (s *ListTimingSyntheticTasksResponse) SetHeaders(v map[string]*string) *ListTimingSyntheticTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTimingSyntheticTasksResponse) SetStatusCode(v int32) *ListTimingSyntheticTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTimingSyntheticTasksResponse) SetBody(v *ListTimingSyntheticTasksResponseBody) *ListTimingSyntheticTasksResponse {
	s.Body = v
	return s
}

func (s *ListTimingSyntheticTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
