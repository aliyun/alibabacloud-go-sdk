// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventBusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventBusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventBusResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventBusResponseBody) *CreateEventBusResponse
	GetBody() *CreateEventBusResponseBody
}

type CreateEventBusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventBusResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventBusResponse) GoString() string {
	return s.String()
}

func (s *CreateEventBusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventBusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventBusResponse) GetBody() *CreateEventBusResponseBody {
	return s.Body
}

func (s *CreateEventBusResponse) SetHeaders(v map[string]*string) *CreateEventBusResponse {
	s.Headers = v
	return s
}

func (s *CreateEventBusResponse) SetStatusCode(v int32) *CreateEventBusResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventBusResponse) SetBody(v *CreateEventBusResponseBody) *CreateEventBusResponse {
	s.Body = v
	return s
}

func (s *CreateEventBusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
