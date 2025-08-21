// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScheduleTaskResponseBody) *DeleteScheduleTaskResponse
	GetBody() *DeleteScheduleTaskResponseBody
}

type DeleteScheduleTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScheduleTaskResponse) GetBody() *DeleteScheduleTaskResponseBody {
	return s.Body
}

func (s *DeleteScheduleTaskResponse) SetHeaders(v map[string]*string) *DeleteScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduleTaskResponse) SetStatusCode(v int32) *DeleteScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduleTaskResponse) SetBody(v *DeleteScheduleTaskResponseBody) *DeleteScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteScheduleTaskResponse) Validate() error {
	return dara.Validate(s)
}
