// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnregisterDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnregisterDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnregisterDeviceResponse
	GetStatusCode() *int32
	SetBody(v *UnregisterDeviceResponseBody) *UnregisterDeviceResponse
	GetBody() *UnregisterDeviceResponseBody
}

type UnregisterDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnregisterDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnregisterDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s UnregisterDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnregisterDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnregisterDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnregisterDeviceResponse) GetBody() *UnregisterDeviceResponseBody {
	return s.Body
}

func (s *UnregisterDeviceResponse) SetHeaders(v map[string]*string) *UnregisterDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnregisterDeviceResponse) SetStatusCode(v int32) *UnregisterDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnregisterDeviceResponse) SetBody(v *UnregisterDeviceResponseBody) *UnregisterDeviceResponse {
	s.Body = v
	return s
}

func (s *UnregisterDeviceResponse) Validate() error {
	return dara.Validate(s)
}
