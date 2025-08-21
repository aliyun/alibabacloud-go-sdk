// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartDevicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStartDevicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStartDevicesResponse
	GetStatusCode() *int32
	SetBody(v *BatchStartDevicesResponseBody) *BatchStartDevicesResponse
	GetBody() *BatchStartDevicesResponseBody
}

type BatchStartDevicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStartDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStartDevicesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStartDevicesResponse) GoString() string {
	return s.String()
}

func (s *BatchStartDevicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStartDevicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStartDevicesResponse) GetBody() *BatchStartDevicesResponseBody {
	return s.Body
}

func (s *BatchStartDevicesResponse) SetHeaders(v map[string]*string) *BatchStartDevicesResponse {
	s.Headers = v
	return s
}

func (s *BatchStartDevicesResponse) SetStatusCode(v int32) *BatchStartDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStartDevicesResponse) SetBody(v *BatchStartDevicesResponseBody) *BatchStartDevicesResponse {
	s.Body = v
	return s
}

func (s *BatchStartDevicesResponse) Validate() error {
	return dara.Validate(s)
}
