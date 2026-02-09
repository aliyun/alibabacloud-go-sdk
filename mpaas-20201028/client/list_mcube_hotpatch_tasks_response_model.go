// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeHotpatchTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeHotpatchTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeHotpatchTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeHotpatchTasksResponseBody) *ListMcubeHotpatchTasksResponse
	GetBody() *ListMcubeHotpatchTasksResponseBody
}

type ListMcubeHotpatchTasksResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeHotpatchTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeHotpatchTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeHotpatchTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeHotpatchTasksResponse) GetBody() *ListMcubeHotpatchTasksResponseBody {
	return s.Body
}

func (s *ListMcubeHotpatchTasksResponse) SetHeaders(v map[string]*string) *ListMcubeHotpatchTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeHotpatchTasksResponse) SetStatusCode(v int32) *ListMcubeHotpatchTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeHotpatchTasksResponse) SetBody(v *ListMcubeHotpatchTasksResponseBody) *ListMcubeHotpatchTasksResponse {
	s.Body = v
	return s
}

func (s *ListMcubeHotpatchTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
