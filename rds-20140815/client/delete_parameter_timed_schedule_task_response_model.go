// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterTimedScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteParameterTimedScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteParameterTimedScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteParameterTimedScheduleTaskResponseBody) *DeleteParameterTimedScheduleTaskResponse
	GetBody() *DeleteParameterTimedScheduleTaskResponseBody
}

type DeleteParameterTimedScheduleTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParameterTimedScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParameterTimedScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterTimedScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteParameterTimedScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteParameterTimedScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteParameterTimedScheduleTaskResponse) GetBody() *DeleteParameterTimedScheduleTaskResponseBody {
	return s.Body
}

func (s *DeleteParameterTimedScheduleTaskResponse) SetHeaders(v map[string]*string) *DeleteParameterTimedScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteParameterTimedScheduleTaskResponse) SetStatusCode(v int32) *DeleteParameterTimedScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParameterTimedScheduleTaskResponse) SetBody(v *DeleteParameterTimedScheduleTaskResponseBody) *DeleteParameterTimedScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteParameterTimedScheduleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
