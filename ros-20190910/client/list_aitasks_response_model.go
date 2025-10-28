// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAITasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAITasksResponse
	GetStatusCode() *int32
	SetBody(v *ListAITasksResponseBody) *ListAITasksResponse
	GetBody() *ListAITasksResponseBody
}

type ListAITasksResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAITasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAITasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAITasksResponse) GoString() string {
	return s.String()
}

func (s *ListAITasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAITasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAITasksResponse) GetBody() *ListAITasksResponseBody {
	return s.Body
}

func (s *ListAITasksResponse) SetHeaders(v map[string]*string) *ListAITasksResponse {
	s.Headers = v
	return s
}

func (s *ListAITasksResponse) SetStatusCode(v int32) *ListAITasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAITasksResponse) SetBody(v *ListAITasksResponseBody) *ListAITasksResponse {
	s.Body = v
	return s
}

func (s *ListAITasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
