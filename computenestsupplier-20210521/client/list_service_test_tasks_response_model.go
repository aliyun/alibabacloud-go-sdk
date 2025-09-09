// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceTestTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceTestTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceTestTasksResponseBody) *ListServiceTestTasksResponse
	GetBody() *ListServiceTestTasksResponseBody
}

type ListServiceTestTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceTestTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceTestTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTasksResponse) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceTestTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceTestTasksResponse) GetBody() *ListServiceTestTasksResponseBody {
	return s.Body
}

func (s *ListServiceTestTasksResponse) SetHeaders(v map[string]*string) *ListServiceTestTasksResponse {
	s.Headers = v
	return s
}

func (s *ListServiceTestTasksResponse) SetStatusCode(v int32) *ListServiceTestTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceTestTasksResponse) SetBody(v *ListServiceTestTasksResponseBody) *ListServiceTestTasksResponse {
	s.Body = v
	return s
}

func (s *ListServiceTestTasksResponse) Validate() error {
	return dara.Validate(s)
}
