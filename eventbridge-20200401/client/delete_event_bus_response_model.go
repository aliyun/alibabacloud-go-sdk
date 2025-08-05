// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventBusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventBusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventBusResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventBusResponseBody) *DeleteEventBusResponse
	GetBody() *DeleteEventBusResponseBody
}

type DeleteEventBusResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventBusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventBusResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventBusResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventBusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventBusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventBusResponse) GetBody() *DeleteEventBusResponseBody {
	return s.Body
}

func (s *DeleteEventBusResponse) SetHeaders(v map[string]*string) *DeleteEventBusResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventBusResponse) SetStatusCode(v int32) *DeleteEventBusResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventBusResponse) SetBody(v *DeleteEventBusResponseBody) *DeleteEventBusResponse {
	s.Body = v
	return s
}

func (s *DeleteEventBusResponse) Validate() error {
	return dara.Validate(s)
}
