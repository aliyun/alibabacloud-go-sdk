// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduleTaskResponseBody) *CreateScheduleTaskResponse
	GetBody() *CreateScheduleTaskResponseBody
}

type CreateScheduleTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduleTaskResponse) GetBody() *CreateScheduleTaskResponseBody {
	return s.Body
}

func (s *CreateScheduleTaskResponse) SetHeaders(v map[string]*string) *CreateScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduleTaskResponse) SetStatusCode(v int32) *CreateScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduleTaskResponse) SetBody(v *CreateScheduleTaskResponseBody) *CreateScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *CreateScheduleTaskResponse) Validate() error {
	return dara.Validate(s)
}
