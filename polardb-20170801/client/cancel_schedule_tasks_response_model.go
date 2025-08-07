// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelScheduleTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelScheduleTasksResponse
	GetStatusCode() *int32
	SetBody(v *CancelScheduleTasksResponseBody) *CancelScheduleTasksResponse
	GetBody() *CancelScheduleTasksResponseBody
}

type CancelScheduleTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelScheduleTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelScheduleTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleTasksResponse) GoString() string {
	return s.String()
}

func (s *CancelScheduleTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelScheduleTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelScheduleTasksResponse) GetBody() *CancelScheduleTasksResponseBody {
	return s.Body
}

func (s *CancelScheduleTasksResponse) SetHeaders(v map[string]*string) *CancelScheduleTasksResponse {
	s.Headers = v
	return s
}

func (s *CancelScheduleTasksResponse) SetStatusCode(v int32) *CancelScheduleTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelScheduleTasksResponse) SetBody(v *CancelScheduleTasksResponseBody) *CancelScheduleTasksResponse {
	s.Body = v
	return s
}

func (s *CancelScheduleTasksResponse) Validate() error {
	return dara.Validate(s)
}
