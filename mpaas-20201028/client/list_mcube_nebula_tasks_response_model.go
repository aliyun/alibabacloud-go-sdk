// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeNebulaTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeNebulaTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeNebulaTasksResponseBody) *ListMcubeNebulaTasksResponse
	GetBody() *ListMcubeNebulaTasksResponseBody
}

type ListMcubeNebulaTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeNebulaTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeNebulaTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeNebulaTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeNebulaTasksResponse) GetBody() *ListMcubeNebulaTasksResponseBody {
	return s.Body
}

func (s *ListMcubeNebulaTasksResponse) SetHeaders(v map[string]*string) *ListMcubeNebulaTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeNebulaTasksResponse) SetStatusCode(v int32) *ListMcubeNebulaTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeNebulaTasksResponse) SetBody(v *ListMcubeNebulaTasksResponseBody) *ListMcubeNebulaTasksResponse {
	s.Body = v
	return s
}

func (s *ListMcubeNebulaTasksResponse) Validate() error {
	return dara.Validate(s)
}
