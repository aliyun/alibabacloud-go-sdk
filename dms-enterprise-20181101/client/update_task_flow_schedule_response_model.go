// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowScheduleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTaskFlowScheduleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTaskFlowScheduleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTaskFlowScheduleResponseBody) *UpdateTaskFlowScheduleResponse
	GetBody() *UpdateTaskFlowScheduleResponseBody
}

type UpdateTaskFlowScheduleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTaskFlowScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTaskFlowScheduleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowScheduleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTaskFlowScheduleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTaskFlowScheduleResponse) GetBody() *UpdateTaskFlowScheduleResponseBody {
	return s.Body
}

func (s *UpdateTaskFlowScheduleResponse) SetHeaders(v map[string]*string) *UpdateTaskFlowScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateTaskFlowScheduleResponse) SetStatusCode(v int32) *UpdateTaskFlowScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTaskFlowScheduleResponse) SetBody(v *UpdateTaskFlowScheduleResponseBody) *UpdateTaskFlowScheduleResponse {
	s.Body = v
	return s
}

func (s *UpdateTaskFlowScheduleResponse) Validate() error {
	return dara.Validate(s)
}
