// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterTimedScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyParameterTimedScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyParameterTimedScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyParameterTimedScheduleTaskResponseBody) *ModifyParameterTimedScheduleTaskResponse
	GetBody() *ModifyParameterTimedScheduleTaskResponseBody
}

type ModifyParameterTimedScheduleTaskResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyParameterTimedScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyParameterTimedScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterTimedScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyParameterTimedScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyParameterTimedScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyParameterTimedScheduleTaskResponse) GetBody() *ModifyParameterTimedScheduleTaskResponseBody {
	return s.Body
}

func (s *ModifyParameterTimedScheduleTaskResponse) SetHeaders(v map[string]*string) *ModifyParameterTimedScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyParameterTimedScheduleTaskResponse) SetStatusCode(v int32) *ModifyParameterTimedScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParameterTimedScheduleTaskResponse) SetBody(v *ModifyParameterTimedScheduleTaskResponseBody) *ModifyParameterTimedScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyParameterTimedScheduleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
