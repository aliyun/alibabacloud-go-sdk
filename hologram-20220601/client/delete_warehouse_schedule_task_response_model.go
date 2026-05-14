// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWarehouseScheduleTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWarehouseScheduleTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWarehouseScheduleTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWarehouseScheduleTaskResponseBody) *DeleteWarehouseScheduleTaskResponse
	GetBody() *DeleteWarehouseScheduleTaskResponseBody
}

type DeleteWarehouseScheduleTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWarehouseScheduleTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWarehouseScheduleTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWarehouseScheduleTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteWarehouseScheduleTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWarehouseScheduleTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWarehouseScheduleTaskResponse) GetBody() *DeleteWarehouseScheduleTaskResponseBody {
	return s.Body
}

func (s *DeleteWarehouseScheduleTaskResponse) SetHeaders(v map[string]*string) *DeleteWarehouseScheduleTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteWarehouseScheduleTaskResponse) SetStatusCode(v int32) *DeleteWarehouseScheduleTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWarehouseScheduleTaskResponse) SetBody(v *DeleteWarehouseScheduleTaskResponseBody) *DeleteWarehouseScheduleTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteWarehouseScheduleTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
