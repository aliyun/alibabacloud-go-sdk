// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFunctionTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFunctionTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListFunctionTasksResponseBody) *ListFunctionTasksResponse
	GetBody() *ListFunctionTasksResponseBody
}

type ListFunctionTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionTasksResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFunctionTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFunctionTasksResponse) GetBody() *ListFunctionTasksResponseBody {
	return s.Body
}

func (s *ListFunctionTasksResponse) SetHeaders(v map[string]*string) *ListFunctionTasksResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionTasksResponse) SetStatusCode(v int32) *ListFunctionTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionTasksResponse) SetBody(v *ListFunctionTasksResponseBody) *ListFunctionTasksResponse {
	s.Body = v
	return s
}

func (s *ListFunctionTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
