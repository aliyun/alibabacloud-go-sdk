// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRefDISyncTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRefDISyncTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRefDISyncTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListRefDISyncTasksResponseBody) *ListRefDISyncTasksResponse
	GetBody() *ListRefDISyncTasksResponseBody
}

type ListRefDISyncTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRefDISyncTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRefDISyncTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRefDISyncTasksResponse) GoString() string {
	return s.String()
}

func (s *ListRefDISyncTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRefDISyncTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRefDISyncTasksResponse) GetBody() *ListRefDISyncTasksResponseBody {
	return s.Body
}

func (s *ListRefDISyncTasksResponse) SetHeaders(v map[string]*string) *ListRefDISyncTasksResponse {
	s.Headers = v
	return s
}

func (s *ListRefDISyncTasksResponse) SetStatusCode(v int32) *ListRefDISyncTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRefDISyncTasksResponse) SetBody(v *ListRefDISyncTasksResponseBody) *ListRefDISyncTasksResponse {
	s.Body = v
	return s
}

func (s *ListRefDISyncTasksResponse) Validate() error {
	return dara.Validate(s)
}
