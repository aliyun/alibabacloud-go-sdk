// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWarehouseScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWarehouseScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *ListWarehouseScheduleTaskResponseBody) *ListWarehouseScheduleTaskResponse
	GetBody() *ListWarehouseScheduleTaskResponseBody
}

type ListWarehouseScheduleTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarehouseScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarehouseScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWarehouseScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWarehouseScheduleTaskResponse) GetBody() *ListWarehouseScheduleTaskResponseBody {
	return s.Body
}

func (s *ListWarehouseScheduleTaskResponse) SetHeaders(v map[string]*string) *ListWarehouseScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *ListWarehouseScheduleTaskResponse) SetStatusCode(v int32) *ListWarehouseScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarehouseScheduleTaskResponse) SetBody(v *ListWarehouseScheduleTaskResponseBody) *ListWarehouseScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *ListWarehouseScheduleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
