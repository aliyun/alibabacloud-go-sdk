// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRunningTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRunningTasksResponse
	GetStatusCode() *int32
	SetBody(v *GetRunningTasksResponseBody) *GetRunningTasksResponse
	GetBody() *GetRunningTasksResponseBody
}

type GetRunningTasksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRunningTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRunningTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRunningTasksResponse) GoString() string {
	return s.String()
}

func (s *GetRunningTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRunningTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRunningTasksResponse) GetBody() *GetRunningTasksResponseBody {
	return s.Body
}

func (s *GetRunningTasksResponse) SetHeaders(v map[string]*string) *GetRunningTasksResponse {
	s.Headers = v
	return s
}

func (s *GetRunningTasksResponse) SetStatusCode(v int32) *GetRunningTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunningTasksResponse) SetBody(v *GetRunningTasksResponseBody) *GetRunningTasksResponse {
	s.Body = v
	return s
}

func (s *GetRunningTasksResponse) Validate() error {
	return dara.Validate(s)
}
