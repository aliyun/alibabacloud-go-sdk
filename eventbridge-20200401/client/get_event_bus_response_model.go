// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventBusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEventBusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEventBusResponse
	GetStatusCode() *int32
	SetBody(v *GetEventBusResponseBody) *GetEventBusResponse
	GetBody() *GetEventBusResponseBody
}

type GetEventBusResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventBusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEventBusResponse) GoString() string {
	return s.String()
}

func (s *GetEventBusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEventBusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEventBusResponse) GetBody() *GetEventBusResponseBody {
	return s.Body
}

func (s *GetEventBusResponse) SetHeaders(v map[string]*string) *GetEventBusResponse {
	s.Headers = v
	return s
}

func (s *GetEventBusResponse) SetStatusCode(v int32) *GetEventBusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventBusResponse) SetBody(v *GetEventBusResponseBody) *GetEventBusResponse {
	s.Body = v
	return s
}

func (s *GetEventBusResponse) Validate() error {
	return dara.Validate(s)
}
