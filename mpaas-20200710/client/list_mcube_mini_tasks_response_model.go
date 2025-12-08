// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeMiniTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeMiniTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeMiniTasksResponseBody) *ListMcubeMiniTasksResponse
	GetBody() *ListMcubeMiniTasksResponseBody
}

type ListMcubeMiniTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeMiniTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeMiniTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniTasksResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeMiniTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeMiniTasksResponse) GetBody() *ListMcubeMiniTasksResponseBody {
	return s.Body
}

func (s *ListMcubeMiniTasksResponse) SetHeaders(v map[string]*string) *ListMcubeMiniTasksResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeMiniTasksResponse) SetStatusCode(v int32) *ListMcubeMiniTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeMiniTasksResponse) SetBody(v *ListMcubeMiniTasksResponseBody) *ListMcubeMiniTasksResponse {
	s.Body = v
	return s
}

func (s *ListMcubeMiniTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
