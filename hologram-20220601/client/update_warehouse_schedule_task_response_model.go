// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWarehouseScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWarehouseScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWarehouseScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWarehouseScheduleTaskResponseBody) *UpdateWarehouseScheduleTaskResponse
	GetBody() *UpdateWarehouseScheduleTaskResponseBody
}

type UpdateWarehouseScheduleTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWarehouseScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWarehouseScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWarehouseScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateWarehouseScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWarehouseScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWarehouseScheduleTaskResponse) GetBody() *UpdateWarehouseScheduleTaskResponseBody {
	return s.Body
}

func (s *UpdateWarehouseScheduleTaskResponse) SetHeaders(v map[string]*string) *UpdateWarehouseScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateWarehouseScheduleTaskResponse) SetStatusCode(v int32) *UpdateWarehouseScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWarehouseScheduleTaskResponse) SetBody(v *UpdateWarehouseScheduleTaskResponseBody) *UpdateWarehouseScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateWarehouseScheduleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
