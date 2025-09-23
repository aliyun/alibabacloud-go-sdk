// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *ModifyScheduleTaskResponseBody) *ModifyScheduleTaskResponse
	GetBody() *ModifyScheduleTaskResponseBody
}

type ModifyScheduleTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyScheduleTaskResponse) GetBody() *ModifyScheduleTaskResponseBody {
	return s.Body
}

func (s *ModifyScheduleTaskResponse) SetHeaders(v map[string]*string) *ModifyScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduleTaskResponse) SetStatusCode(v int32) *ModifyScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScheduleTaskResponse) SetBody(v *ModifyScheduleTaskResponseBody) *ModifyScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *ModifyScheduleTaskResponse) Validate() error {
	return dara.Validate(s)
}
