// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventBusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventBusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventBusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventBusResponseBody) *UpdateEventBusResponse
	GetBody() *UpdateEventBusResponseBody
}

type UpdateEventBusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventBusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventBusResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventBusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventBusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventBusResponse) GetBody() *UpdateEventBusResponseBody {
	return s.Body
}

func (s *UpdateEventBusResponse) SetHeaders(v map[string]*string) *UpdateEventBusResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventBusResponse) SetStatusCode(v int32) *UpdateEventBusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventBusResponse) SetBody(v *UpdateEventBusResponseBody) *UpdateEventBusResponse {
	s.Body = v
	return s
}

func (s *UpdateEventBusResponse) Validate() error {
	return dara.Validate(s)
}
