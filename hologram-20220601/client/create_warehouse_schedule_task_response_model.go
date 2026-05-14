// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarehouseScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateWarehouseScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateWarehouseScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateWarehouseScheduleTaskResponseBody) *CreateWarehouseScheduleTaskResponse
	GetBody() *CreateWarehouseScheduleTaskResponseBody
}

type CreateWarehouseScheduleTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWarehouseScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWarehouseScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateWarehouseScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateWarehouseScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateWarehouseScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateWarehouseScheduleTaskResponse) GetBody() *CreateWarehouseScheduleTaskResponseBody {
	return s.Body
}

func (s *CreateWarehouseScheduleTaskResponse) SetHeaders(v map[string]*string) *CreateWarehouseScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateWarehouseScheduleTaskResponse) SetStatusCode(v int32) *CreateWarehouseScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWarehouseScheduleTaskResponse) SetBody(v *CreateWarehouseScheduleTaskResponseBody) *CreateWarehouseScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *CreateWarehouseScheduleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
