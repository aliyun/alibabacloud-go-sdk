// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMdsCubeTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMdsCubeTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMdsCubeTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListMdsCubeTasksResponseBody) *ListMdsCubeTasksResponse
	GetBody() *ListMdsCubeTasksResponseBody
}

type ListMdsCubeTasksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMdsCubeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMdsCubeTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMdsCubeTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMdsCubeTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMdsCubeTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMdsCubeTasksResponse) GetBody() *ListMdsCubeTasksResponseBody {
	return s.Body
}

func (s *ListMdsCubeTasksResponse) SetHeaders(v map[string]*string) *ListMdsCubeTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMdsCubeTasksResponse) SetStatusCode(v int32) *ListMdsCubeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMdsCubeTasksResponse) SetBody(v *ListMdsCubeTasksResponseBody) *ListMdsCubeTasksResponse {
	s.Body = v
	return s
}

func (s *ListMdsCubeTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
