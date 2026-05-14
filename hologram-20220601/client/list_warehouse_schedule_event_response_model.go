// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseScheduleEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWarehouseScheduleEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWarehouseScheduleEventResponse
	GetStatusCode() *int32
	SetBody(v *ListWarehouseScheduleEventResponseBody) *ListWarehouseScheduleEventResponse
	GetBody() *ListWarehouseScheduleEventResponseBody
}

type ListWarehouseScheduleEventResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarehouseScheduleEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarehouseScheduleEventResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleEventResponse) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWarehouseScheduleEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWarehouseScheduleEventResponse) GetBody() *ListWarehouseScheduleEventResponseBody {
	return s.Body
}

func (s *ListWarehouseScheduleEventResponse) SetHeaders(v map[string]*string) *ListWarehouseScheduleEventResponse {
	s.Headers = v
	return s
}

func (s *ListWarehouseScheduleEventResponse) SetStatusCode(v int32) *ListWarehouseScheduleEventResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarehouseScheduleEventResponse) SetBody(v *ListWarehouseScheduleEventResponseBody) *ListWarehouseScheduleEventResponse {
	s.Body = v
	return s
}

func (s *ListWarehouseScheduleEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
